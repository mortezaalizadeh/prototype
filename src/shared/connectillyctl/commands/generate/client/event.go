package client

import (
	"github.com/Connectilly/connectilly/src/shared/connectillyctl/appsetup"
	"github.com/Connectilly/connectilly/src/shared/enterprise/logger"
	"github.com/spf13/cobra"
)

type eventOptions struct {
	protobufFilePath string
	packageName      string
	eventType        string
	topicName        string
	outputPath       string
}

func EventCommand() *cobra.Command {
	options := eventOptions{}
	sugarLogger := logger.CreateLogger()

	cmd := &cobra.Command{
		Use:   "event",
		Short: "Event",
		Long:  "Event",
		Run: func(cmd *cobra.Command, args []string) {
			clientEventSchemaGeneratorService, err := appsetup.NewClientEventSchemaGeneratorService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			clientEventMetadataGeneratorService, err := appsetup.NewClientEventMetadataGeneratorService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			clientEventHandlerGeneratorService, err := appsetup.NewClientEventHandlerGeneratorService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			clientEventConsumerGeneratorService, err := appsetup.NewClientEventConsumerGeneratorService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			clientEventProducerGeneratorService, err := appsetup.NewClientEventProducerGeneratorService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			clientEventGenerateGeneratorService, err := appsetup.NewClientEventGenerateGeneratorService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			if err = clientEventSchemaGeneratorService.Generate(
				options.protobufFilePath,
				options.outputPath); err != nil {
				sugarLogger.Fatal(err)
			}

			if err = clientEventMetadataGeneratorService.Generate(
				options.packageName,
				options.topicName,
				options.outputPath); err != nil {
				sugarLogger.Fatal(err)
			}

			if err = clientEventHandlerGeneratorService.Generate(
				options.packageName,
				options.eventType,
				options.outputPath); err != nil {
				sugarLogger.Fatal(err)
			}

			if err = clientEventConsumerGeneratorService.Generate(
				options.packageName,
				options.eventType,
				options.outputPath); err != nil {
				sugarLogger.Fatal(err)
			}

			if err = clientEventProducerGeneratorService.Generate(
				options.packageName,
				options.eventType,
				options.outputPath); err != nil {
				sugarLogger.Fatal(err)
			}

			if err = clientEventGenerateGeneratorService.Generate(
				options.packageName,
				options.outputPath); err != nil {
				sugarLogger.Fatal(err)
			}
		},
	}

	cmd.Flags().StringVar(&options.protobufFilePath, "protobufFilePath", "", "Specify the protobuf file path")
	cmd.Flags().StringVar(&options.packageName, "packageName", "", "Specify the package name")
	cmd.Flags().StringVar(&options.eventType, "eventType", "", "Specify the event type")
	cmd.Flags().StringVar(&options.topicName, "topicName", "", "Specify the topic name")
	cmd.Flags().StringVar(&options.outputPath, "outputPath", "", "Specify the output path")

	if err := cmd.MarkFlagRequired("protobufFilePath"); err != nil {
		sugarLogger.Fatal(err)
	}

	if err := cmd.MarkFlagRequired("packageName"); err != nil {
		sugarLogger.Fatal(err)
	}

	if err := cmd.MarkFlagRequired("eventType"); err != nil {
		sugarLogger.Fatal(err)
	}

	if err := cmd.MarkFlagRequired("topicName"); err != nil {
		sugarLogger.Fatal(err)
	}

	if err := cmd.MarkFlagRequired("outputPath"); err != nil {
		sugarLogger.Fatal(err)
	}

	return cmd
}
