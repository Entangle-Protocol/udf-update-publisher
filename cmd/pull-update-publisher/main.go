package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"gitlab.ent-dx.com/entangle/pull-update-publisher/config"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/keystore"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/publisher"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/transactor"
)

var pullUpdatePublisherCmd = &cobra.Command{
	Use:   "pull-update-publisher",
	Short: "Starts service that publishes prices to PullOracle contract on destination chains",
	Long:  `Starts service that publishes prices to PullOracle contract on destination chains.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		verbosity, err := cmd.Flags().GetInt("verbosity")
		if err != nil {
			log.Fatalf("Failed to parse verbosity flag: %v", err)
			panic(err)
		}

		configPath, err := cmd.Flags().GetString("config")
		if err != nil {
			log.Fatalf("Failed to parse config flag: %v", err)
			panic(err)
		}

		enableLines, err := cmd.Flags().GetBool("lines")
		if err != nil {
			log.Fatalf("Failed to parse lines flag: %v", err)
			panic(err)
		}
		SetupLogger(verbosity, enableLines)

		log.Info("Starting price oracle publisher")

		config, err := config.LoadConfig(configPath)
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
			panic(err)
		}

		fmt.Printf("Config loaded: %+v\n", config)

		// Create fetcher
		restFetcher := fetcher.NewRestFetcher(
			http.DefaultClient, config.FinalizeSnapshotURL,
		)

		transactors := CreateTransactors(ctx, config.Networks)
		// Create publisher
		publisher := publisher.NewUpdatePublisher(config.Publisher, transactors, restFetcher, config.DataKeys, config.Assets)

		ticker := time.NewTicker(time.Duration(config.Publisher.UpdateInterval) * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				// Publish latest feed
				// err := publisher.PublishUpdate(ctx)
				err := publisher.PublishMultipleUpdate(ctx)
				if err != nil {
					log.Errorf("Failed to publish update: %v", err)
				}
			}
		}
	},
}

func CreateTransactors(ctx context.Context, networks map[string]config.NetworkConfig) []transactor.ITransactor {
	transactors := make([]transactor.ITransactor, 0, len(networks))
	for name, net := range networks {
		log := log.WithFields(log.Fields{
			"network": name,
		})
		if net.Type == "nonevm" {
			log.Infof("Register solana transactor")
			transactor, err := transactor.NewSolanaTransactor()
			if err != nil {
				log.Fatalf("Failed to create Solana transactor: %v", err)
				panic(err)
			}
			transactors = append(transactors, transactor)
		} else if net.Type == "evm" {
			log.Infof("Register evm transactor: %s, url: %s, pull oracle: %s", name, net.TargetChainURL, net.PullOracleAddress)
			transactor, err := CreateEVMTransactor(ctx, name, net)
			if err != nil {
				log.Fatalf("Failed to create EVM transactor: %v", err)
				panic(err)
			}
			transactors = append(transactors, transactor)
		}
	}
	return transactors
}

func CreateEVMTransactor(ctx context.Context, networkName string, net config.NetworkConfig) (transactor.ITransactor, error) {
	log := log.WithFields(log.Fields{
		"network":        networkName,
		"targetChainUrl": net.TargetChainURL,
	})

	client, err := ethclient.Dial(net.TargetChainURL)
	if err != nil {
		log.Fatalf("Failed to dial target chain: %v", err)
		return nil, err
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
		return nil, err
	}

	// Load private key
	key, err := keystore.ParseKeyFromHex(net.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
		return nil, err
	}

	transactor, err := transactor.NewTransactor(ctx, client, key, chainID, net.PullOracleAddress)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
		return nil, err
	}

	return transactor, nil
}

func Execute() {
	if err := pullUpdatePublisherCmd.Execute(); err != nil {
		panic(err)
	}
}

func SetupLogger(verbosity int, enableLines bool) {
	switch verbosity {
	case 0:
		fmt.Println("Verbosity set to panic(0) level")
		log.SetLevel(log.PanicLevel)
	case 1:
		log.SetLevel(log.FatalLevel)
		fmt.Println("Verbosity set to fatal(1) level")
	case 2:
		log.SetLevel(log.ErrorLevel)
		fmt.Println("Verbosity set to error(2) level")
	case 3:
		log.SetLevel(log.WarnLevel)
		fmt.Println("Verbosity set to warn(3) level")
	case 4:
		log.SetLevel(log.InfoLevel)
		fmt.Println("Verbosity set to info(4) level")
	case 5:
		log.SetLevel(log.DebugLevel)
		fmt.Println("Verbosity set to debug(5) level")
	case 6:
		log.SetLevel(log.TraceLevel)
		fmt.Println("Verbosity set to trace(6) level")
	default:
		panic("Invalid verbosity level")
	}

	// Set the logger formatter
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// Include line numbers and file names in logs. Sometimes useful for debugging.
	if enableLines {
		log.SetReportCaller(true)
	}

	log.WithFields(log.Fields{
		"verbosity": verbosity,
	}).Debug("Logger initialized")
}

func init() {
	time.Local = time.UTC

	pullUpdatePublisherCmd.PersistentFlags().IntP("verbosity", "v", 4, "Verbosity level")
	pullUpdatePublisherCmd.PersistentFlags().Bool("lines", false, "Include line numbers in logs")
	pullUpdatePublisherCmd.PersistentFlags().String("config", "config.yaml", "Configuration file path")
}

func main() {
	Execute()
}
