package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Connectilly/connectilly/src/organization/organization-api/graphql/generated"
	"github.com/Connectilly/connectilly/src/organization/shared/entities"
)

func (r *queryResolver) Organization(ctx context.Context, where *entities.OrganizationWhereInput) (*entities.Organization, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Organizations(ctx context.Context, after *entities.Cursor, first *int, before *entities.Cursor, last *int, orderBy []*entities.OrganizationOrder, where *entities.OrganizationWhereInput) (*entities.OrganizationConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
