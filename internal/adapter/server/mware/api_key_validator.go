package mware

import (
	"github.com/joqd/authify/internal/core/port"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type apiKeyMiddleware struct {
	apiKeyService port.ApiKeyService
	logger        *zap.SugaredLogger
}

func NewApiKeyMiddleware(apiKeyService port.ApiKeyService, logger *zap.SugaredLogger) *apiKeyMiddleware {
	return &apiKeyMiddleware{
		apiKeyService: apiKeyService,
		logger:        logger,
	}
}

func (apm *apiKeyMiddleware) Validator(key string, c echo.Context) (bool, error) {
	isValid, err := apm.apiKeyService.IsValid(c.Request().Context(), key)
	if err != nil {
		apm.logger.Errorf("failed to validate api key: %v", err)
		return false, err
	}

	return isValid, nil
}
