package mappers

import (
	graphqlmodels "github.com/Connectilly/connectilly/src/organization/organization-api/graphql/models"
	"github.com/Connectilly/connectilly/src/organization/shared/models"
)

func CreateOrganizationInputToOrganization(src graphqlmodels.CreateOrganizationInput) models.Organization {
	organization := models.Organization{
		Name: src.Name,
	}

	return organization
}
