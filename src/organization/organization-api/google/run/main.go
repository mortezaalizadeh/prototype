package main

import (
	"os"

	"github.com/Connectilly/connectilly/src/organization/organization-api/appsetup"
	"github.com/Connectilly/connectilly/src/organization/organization-api/configuration"
	organizationappsetup "github.com/Connectilly/connectilly/src/organization/shared/appsetup"
	enterpriseappsetup "github.com/Connectilly/connectilly/src/shared/enterprise/appsetup"
	"github.com/Connectilly/connectilly/src/shared/enterprise/logger"
)

func main() {
	sugarLogger := logger.CreateLogger()

	configurationHelper, err := enterpriseappsetup.NewConfigurationHelper(sugarLogger)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	var config configuration.Config
	if err := configurationHelper.LoadYaml("config.yaml", &config); err != nil {
		sugarLogger.Fatal(err)
	}

	if val, ok := os.LookupEnv("PORT"); ok {
		config.App.ListeningInterface = ":" + val
	} else {
		config.App.ListeningInterface = ":8080"
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
}
