package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"context"

	"entgo.io/contrib/entgql"
	"github.com/Connectilly/connectilly/src/organization/organization-api/graphql/generated"
	"github.com/Connectilly/connectilly/src/organization/organization-api/graphql/models"
	"github.com/Connectilly/connectilly/src/organization/organization-api/mappers"
	"github.com/Connectilly/connectilly/src/organization/shared/entities"
)

// CreateOrganization is the resolver for the createOrganization field.
func (r *mutationResolver) CreateOrganization(ctx context.Context, input models.CreateOrganizationInput) (*models.OrganizationPayload, error) {
	organization := mappers.CreateOrganizationInputToOrganization(input)

	createdOrganization, err := r.organizationService.Submit(ctx, organization)
	if err != nil {
		return nil, err
	}

	return &models.OrganizationPayload{
		ClientMutationID: input.ClientMutationID,
		Organization: &entities.Organization{
			ID:   createdOrganization.Id,
			Name: createdOrganization.Name,
		},
	}, nil
}

// Organization is the resolver for the organization field.
func (r *queryResolver) Organization(ctx context.Context, where *entities.OrganizationWhereInput) (*entities.Organization, error) {
	query, err := where.Filter(r.client.Organization.Query())
	if err != nil {
		return nil, err
	}

	result, err := query.First(ctx)
	if _, ok := err.(*entities.NotFoundError); ok {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return result, nil
}

// Organizations is the resolver for the organizations field.
func (r *queryResolver) Organizations(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy []*entities.OrganizationOrder, where *entities.OrganizationWhereInput) (*entities.OrganizationConnection, error) {
	return r.client.Organization.
		Query().
		Paginate(
			ctx,
			after,
			first,
			before,
			last,
			entities.WithOrganizationFilter(where.Filter))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
