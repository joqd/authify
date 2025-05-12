package apikey

import (
	"github.com/joqd/authify/internal/adapter/config"
	"github.com/joqd/authify/internal/core/port"
	"go.uber.org/zap"
)

type apiKeyHandler struct {
	apiKeyService port.ApiKeyService
	logger        *zap.SugaredLogger
	conf          *config.Config
}

func NewApiKeyHandler(apiKeyService port.ApiKeyService, logger *zap.SugaredLogger) port.ApiKeyHandler {
	return &apiKeyHandler{
		apiKeyService: apiKeyService,
		logger:        logger,
	}
}
