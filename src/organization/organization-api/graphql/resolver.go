package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Connectilly/connectilly/src/organization/organization-api/graphql/generated"
	"github.com/Connectilly/connectilly/src/organization/organization-api/services"
	"github.com/Connectilly/connectilly/src/organization/shared/entities"
)

type Resolver struct {
	client              *entities.Client
	organizationService services.OrganizationService
}

type GraphQLServer interface {
}

func NewGraphQLServer(
	entgoClient entities.EntgoClient,
	organizationService services.OrganizationService) (*handler.Server, error) {
	executableSchema := generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			client:              entgoClient.GetClient(),
			organizationService: organizationService,
		},
	})

	return handler.NewDefaultServer(executableSchema), nil
}
