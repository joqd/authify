package service

import (
	"context"

	"github.com/joqd/authify/internal/core/domain"
	"github.com/joqd/authify/internal/core/port"
	"github.com/joqd/authify/internal/core/util"
	"go.uber.org/zap"
)

type apiKeyService struct {
	repo   port.ApiKeyRepository
	logger *zap.SugaredLogger
}

func NewApiKeyService(repo port.ApiKeyRepository, logger *zap.SugaredLogger) port.ApiKeyService {
	return &apiKeyService{
		repo:   repo,
		logger: logger,
	}
}

func (aks *apiKeyService) Create(ctx context.Context, apiKey *domain.APIKey) (*domain.APIKey, error) {
	// Initialize key hash
	newKey := util.GenerateRandomAPIKey()
	hashedKey, err := util.HashSecret(newKey)
	if err != nil {
		return nil, domain.ErrInternal
	}

	apiKey.KeyHash = hashedKey

	// Call repo
	created, err := aks.repo.Create(ctx, apiKey)
	if err != nil {
		return nil, err
	}

	created.Key = &newKey
	return created, nil
}

func (aks *apiKeyService) List(ctx context.Context) ([]domain.APIKey, error) {
	return aks.repo.List(ctx)
}

func (aks *apiKeyService) IsValid(ctx context.Context, key string) (bool, error) {
	_, err := aks.repo.GetByKey(ctx, key)
	if err != nil {
		return false, err
	}

	return true, nil
}
