package repositories

import (
	"context"

	"github.com/Connectilly/connectilly/src/organization/shared/entities"
	"github.com/Connectilly/connectilly/src/organization/shared/models"
)

type OrganizationRepository interface {
	CreateOrganization(
		ctx context.Context,
		tx *entities.Tx,
		organization models.Organization) (models.Organization, error)
}

type organizationRepository struct {
}

func NewOrganizationRepository() (OrganizationRepository, error) {
	return &organizationRepository{}, nil
}

func (or *organizationRepository) CreateOrganization(
	ctx context.Context,
	tx *entities.Tx,
	organization models.Organization) (models.Organization, error) {
	savedOrganization, err := tx.Organization.
		Create().
		SetName(organization.Name).
		Save(ctx)
	if err != nil {
		return models.Organization{}, err
	}

	organization.Id = savedOrganization.ID

	return organization, nil
}
