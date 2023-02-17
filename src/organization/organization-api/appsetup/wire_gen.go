// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appsetup

import (
	"github.com/Connectilly/connectilly/src/organization/organization-api/graphql"
	"github.com/Connectilly/connectilly/src/organization/organization-api/http"
	"github.com/Connectilly/connectilly/src/organization/organization-api/services"
	"github.com/Connectilly/connectilly/src/organization/shared/entities"
	"github.com/Connectilly/connectilly/src/organization/shared/repositories"
	"github.com/Connectilly/connectilly/src/shared/enterprise/configuration"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewHttpServer(logger *zap.SugaredLogger, appConfig configuration.AppConfig, entgoClient entities.EntgoClient) (http.HttpServer, error) {
	organizationRepository, err := repositories.NewOrganizationRepository()
	if err != nil {
		return nil, err
	}
	organizationService, err := services.NewOrganizationService(logger, entgoClient, organizationRepository)
	if err != nil {
		return nil, err
	}
	server, err := graphql.NewGraphQLServer(entgoClient, organizationService)
	if err != nil {
		return nil, err
	}
	httpServer, err := http.NewHttpServer(logger, appConfig, server)
	if err != nil {
		return nil, err
	}
	return httpServer, nil
}