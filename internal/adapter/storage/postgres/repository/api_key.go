package repository

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"
	"github.com/joqd/authify/internal/adapter/storage/postgres/model"
	"github.com/joqd/authify/internal/core/domain"
	"github.com/joqd/authify/internal/core/port"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type apiKeyRepository struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewApiKeyRepository(db *gorm.DB, logger *zap.SugaredLogger) port.ApiKeyRepository {
	return &apiKeyRepository{
		db:     db,
		logger: logger,
	}
}

func (akr *apiKeyRepository) Create(ctx context.Context, apiKey *domain.APIKey) (*domain.APIKey, error) {
	// copy domainApiKey to modelApiKey
	var modelApiKey model.APIKeyModel
	if err := copier.CopyWithOption(&modelApiKey, &apiKey, copier.Option{IgnoreEmpty: true}); err != nil {
		akr.logger.Errorf("failed to copy domainApiKey to modelApiKey: %v", err)
		return nil, domain.ErrCopier
	}

	// DB action
	if err := akr.db.WithContext(ctx).Create(&modelApiKey).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, domain.ErrConflictingData
		}

		return nil, err
	}

	// copy modelApiKey to domainApiKey
	var domainApiKey domain.APIKey
	if err := copier.CopyWithOption(&domainApiKey, &modelApiKey, copier.Option{DeepCopy: true}); err != nil {
		akr.logger.Errorf("failed to copy modelApiKey to domainApiKey: %v", err)
		return nil, domain.ErrCopier
	}

	return &domainApiKey, nil
}

func (akr *apiKeyRepository) List(ctx context.Context) ([]domain.APIKey, error) {
	var modelApiKeys []model.APIKeyModel
	var domainApiKeys []domain.APIKey

	if err := akr.db.WithContext(ctx).Find(&modelApiKeys).Error; err != nil {
		akr.logger.Errorf("failed to get api key list: %v", err)
		return nil, err
	}

	for _, apiKeyModel := range modelApiKeys {
		var apiKeyDomain domain.APIKey
		if err := copier.CopyWithOption(&apiKeyDomain, &apiKeyModel, copier.Option{DeepCopy: true}); err != nil {
			akr.logger.Errorf("failed to map apiKeyModel to apiKeyDomain: %v", err)
			return nil, err
		}

		domainApiKeys = append(domainApiKeys, apiKeyDomain)
	}

	return domainApiKeys, nil
}

func (akr *apiKeyRepository) GetByKey(ctx context.Context, key string) (*domain.APIKey, error) {
	var modelApiKeys []model.APIKeyModel
	var domainApiKey domain.APIKey

	if err := akr.db.WithContext(ctx).Find(&modelApiKeys).Error; err != nil {
		return nil, err
	}

	for _, apiKey := range modelApiKeys {
		err := bcrypt.CompareHashAndPassword([]byte(apiKey.KeyHash), []byte(key))
		if err == nil {
			if err := copier.CopyWithOption(&domainApiKey, &apiKey, copier.Option{IgnoreEmpty: true}); err != nil {
				return nil, domain.ErrCopier
			}

			return &domainApiKey, nil
		}
	}

	return nil, domain.ErrDataNotFound
}
