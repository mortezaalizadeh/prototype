package commands

import (
	"github.com/Connectilly/connectilly/src/organization/organization-api/appsetup"
	"github.com/Connectilly/connectilly/src/organization/organization-api/configuration"
	organizationappsetup "github.com/Connectilly/connectilly/src/organization/shared/appsetup"
	enterpriseappsetup "github.com/Connectilly/connectilly/src/shared/enterprise/appsetup"
	"github.com/Connectilly/connectilly/src/shared/enterprise/logger"
	"github.com/spf13/cobra"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start organization-api",
		Long:  "Start organization-api",
		Run: func(cmd *cobra.Command, args []string) {
			sugarLogger := logger.CreateLogger()

			configurationHelper, err := enterpriseappsetup.NewConfigurationHelper(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			var config configuration.Config
			if err := configurationHelper.LoadYaml("config.yaml", &config); err != nil {
				sugarLogger.Fatal(err)
			}

			entgoClient, err := organizationappsetup.NewEntgoClient(
				sugarLogger,
				config.Postgres)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			httpServer, err := appsetup.NewHttpServer(
				sugarLogger,
				config.App,
				entgoClient)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			err = httpServer.ListenAndServe()
			if err != nil {
				sugarLogger.Fatal(err)
			}
		},
	}

	return cmd
}
