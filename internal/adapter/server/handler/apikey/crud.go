package apikey

import (
	"errors"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/joqd/authify/internal/adapter/server/dto/request"
	"github.com/joqd/authify/internal/adapter/server/dto/response"
	"github.com/joqd/authify/internal/adapter/server/handler"
	"github.com/joqd/authify/internal/core/domain"
	"github.com/labstack/echo/v4"
)

func (akh *apiKeyHandler) Create(c echo.Context) error {
	// Bind request into DTO
	var req request.CreateApiKeyRequest
	if err := c.Bind(&req); err != nil {
		akh.logger.Errorf("failed to bind data: %v", err)
		return handler.RespondError(c, http.StatusBadRequest, response.DescBadRequest)
	}

	// Validate
	if err := c.Validate(&req); err != nil {
		akh.logger.Errorf("failed to validate data: %v", err)
		return handler.RespondError(c, http.StatusUnprocessableEntity, response.DescValidationFailed)
	}

	// Copy req to domainApiKey
	var domainApiKey domain.APIKey
	if err := copier.CopyWithOption(&domainApiKey, &req, copier.Option{IgnoreEmpty: true}); err != nil {
		return handler.RespondError(c, http.StatusInternalServerError, response.DescInternalError)
	}

	// Call service
	created, err := akh.apiKeyService.Create(c.Request().Context(), &domainApiKey)
	if err != nil {
		if errors.Is(err, domain.ErrConflictingData) {
			return handler.RespondError(c, http.StatusConflict, response.DescConflict)
		}

		return handler.RespondError(c, http.StatusInternalServerError, response.DescInternalError)
	}

	// Copy domainApiKey to CreatedApiKey
	var createdApiKey response.CreatedApiKey
	if err := copier.Copy(&createdApiKey, &created); err != nil {
		return handler.RespondError(c, http.StatusInternalServerError, response.DescInternalError)
	}

	return handler.RespondSuccess(c, http.StatusCreated, createdApiKey)
}
