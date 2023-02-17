// Copyright 2018 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package appsetup

import (
	"github.com/Connectilly/connectilly/src/organization/organization-api/graphql"
	"github.com/Connectilly/connectilly/src/organization/organization-api/http"
	"github.com/Connectilly/connectilly/src/organization/organization-api/services"
	"github.com/Connectilly/connectilly/src/organization/shared/entities"
	"github.com/Connectilly/connectilly/src/organization/shared/repositories"
	"github.com/Connectilly/connectilly/src/shared/enterprise/configuration"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func NewHttpServer(
	logger *zap.SugaredLogger,
	appConfig configuration.AppConfig,
	entgoClient entities.EntgoClient) (http.HttpServer, error) {
	wire.Build(
		http.NewHttpServer,
		graphql.NewGraphQLServer,
		services.NewOrganizationService,
		repositories.NewOrganizationRepository)

	return nil, nil
}
