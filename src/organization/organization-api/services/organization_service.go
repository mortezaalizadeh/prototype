package services

import (
	"context"
	"fmt"

	"github.com/Connectilly/connectilly/src/organization/shared/entities"
	"github.com/Connectilly/connectilly/src/organization/shared/models"
	"github.com/Connectilly/connectilly/src/organization/shared/repositories"
	"go.uber.org/zap"
)

type OrganizationService interface {
	Submit(ctx context.Context, organization models.Organization) (models.Organization, error)
}

type organizationService struct {
	logger                 *zap.SugaredLogger
	entgoClient            entities.EntgoClient
	organizationRepository repositories.OrganizationRepository
}

func NewOrganizationService(
	logger *zap.SugaredLogger,
	entgoClient entities.EntgoClient,
	organizationRepository repositories.OrganizationRepository) (OrganizationService, error) {
	return &organizationService{
		logger:                 logger,
		entgoClient:            entgoClient,
		organizationRepository: organizationRepository,
	}, nil
}

func (os *organizationService) Submit(
	ctx context.Context,
	organization models.Organization) (models.Organization, error) {
	return organization, fmt.Errorf("not implemented")
}
