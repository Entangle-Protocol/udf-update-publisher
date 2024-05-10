package main

import (
	"fmt"
	"time"
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"

	"gitlab.ent-dx.com/entangle/pull-update-publisher/config"
)

var pullUpdatePublisherCmd = &cobra.Command{
	Use:   "pull-update-publisher",
	Short: "Starts service that publishes prices to PullOracle contract on destination chains",
	Long: `Starts service that publishes prices to PullOracle contract on destination chains.`,
	Run: func(cmd *cobra.Command, args []string) {
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

		fmt.Println("Config loaded: %+v", config)
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
