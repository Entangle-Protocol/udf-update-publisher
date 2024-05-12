package main

import (
	"fmt"
	"time"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	"gitlab.ent-dx.com/entangle/pull-update-publisher/config"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/publisher"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/transactor"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/keystore"
)

var pullUpdatePublisherCmd = &cobra.Command{
	Use:   "pull-update-publisher",
	Short: "Starts service that publishes prices to PullOracle contract on destination chains",
	Long: `Starts service that publishes prices to PullOracle contract on destination chains.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		verbosity, err := cmd.Flags().GetInt("verbosity")
		if err != nil {
			log.Fatalf("Failed to parse verbosity flag: %v", err)
			panic(err)
		}

		enableLines, err := cmd.Flags().GetBool("lines")
		if err != nil {
			log.Fatalf("Failed to parse lines flag: %v", err)
			panic(err)
		}
		SetupLogger(verbosity, enableLines)

		fmt.Println("Starting price oracle publisher")

		config, err := config.LoadConfigFromEnv()
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
			panic(err)
		}
		config.Verify()

		fmt.Printf("Config loaded: %+v\n", config)

		// Create fetcher
		restFetcher, err := fetcher.NewRestFetcher(
			config.FinalizeSnapshotUrl, &http.Client{},
		)
		if err != nil {
			log.WithFields(log.Fields{
				"finalizeSnapshotUrl": config.FinalizeSnapshotUrl,
			}).Fatalf("Failed to create fetcher: %v", err)
			panic(err)
		}

		//
		// Create transactor
		// 

		client, err := ethclient.Dial(config.TargetChainUrl)
		if err != nil {
			log.WithFields(log.Fields{
				"targetChainUrl": config.TargetChainUrl,
			}).Fatalf("Failed to dial target chain: %v", err)
			panic(err)
		}

		chainID, err := client.ChainID(ctx)
		if err != nil {
			log.WithFields(log.Fields{
				"targetChainUrl": config.TargetChainUrl,
			}).Fatalf("Failed to get chain ID: %v", err)
			panic(err)
		}

		// Load private key
		key, err := keystore.ParseKeyFromHex(config.PrivateKey)

		transactor, err := transactor.NewTransactor(ctx, client, key, chainID, config.PullOracleAddress)
		if err != nil {
			log.Fatalf("Failed to create transactor: %v", err)
			panic(err)
		}

		// Create publisher
		publisher := publisher.NewUpdatePublisher(transactor, restFetcher)

		for {
			// Publish latest feed
			err := publisher.PublishUpdate()
			if err != nil {
				log.Errorf("Failed to publish update: %v", err)
			}

			time.Sleep(10 * time.Second)
		}
	},
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
}

func main() {
	Execute()
}
