package services

import (
	"context"

	"github.com/Connectilly/connectilly/src/organization/shared/models"
	"github.com/Connectilly/connectilly/src/organization/shared/repositories"
	"go.uber.org/zap"
)

type OrganizationService interface {
	Submit(ctx context.Context, organization models.Organization) (*models.Organization, error)
}

type organizationService struct {
	logger                 *zap.SugaredLogger
	organizationRepository repositories.OrganizationRepository
}

func NewOrganizationService(
	logger *zap.SugaredLogger,
	organizationRepository repositories.OrganizationRepository) (OrganizationService, error) {
	return &organizationService{
		logger:                 logger,
		organizationRepository: organizationRepository,
	}, nil
}

func (os *organizationService) Submit(
	ctx context.Context,
	organization models.Organization) (*models.Organization, error) {
	return os.organizationRepository.CreateOrganization(ctx, organization, nil)
}
