package port

import (
	"context"

	"github.com/joqd/authify/internal/core/domain"
	"github.com/labstack/echo/v4"
)

type ApiKeyRepository interface {
	List(ctx context.Context) ([]domain.APIKey, error)
	Create(ctx context.Context, apiKey *domain.APIKey) (*domain.APIKey, error)
	GetByKey(ctx context.Context, key string) (*domain.APIKey, error)
}

type ApiKeyService interface {
	List(ctx context.Context) ([]domain.APIKey, error)
	Create(ctx context.Context, apiKey *domain.APIKey) (*domain.APIKey, error)
	IsValid(ctx context.Context, key string) (bool, error)
}

type ApiKeyHandler interface {
	Create(c echo.Context) error
}
