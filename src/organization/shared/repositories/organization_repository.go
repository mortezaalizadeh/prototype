package repositories

import (
	"context"

	"github.com/Connectilly/connectilly/src/organization/shared/entities"
	"github.com/Connectilly/connectilly/src/organization/shared/models"
)

type OrganizationRepository interface {
	CreateOrganization(
		ctx context.Context,
		organization models.Organization,
		tx *entities.Tx) (*models.Organization, error)
}

type organizationRepository struct {
	entgoClient entities.EntgoClient
}

func NewOrganizationRepository(
	entgoClient entities.EntgoClient,
) (OrganizationRepository, error) {
	return &organizationRepository{
		entgoClient: entgoClient,
	}, nil
}

func (or *organizationRepository) CreateOrganization(
	ctx context.Context,
	organization models.Organization,
	tx *entities.Tx) (*models.Organization, error) {
	var savedOrganization *entities.Organization
	var err error

	if tx == nil {
		client := or.entgoClient.GetClient()
		savedOrganization, err = client.Organization.
			Create().
			SetName(organization.Name).
			Save(ctx)

	} else {
		savedOrganization, err = tx.Organization.
			Create().
			SetName(organization.Name).
			Save(ctx)
	}

	if err != nil {
		return nil, err
	}

	organization.Id = savedOrganization.ID

	return &organization, nil
}
