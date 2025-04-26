package handler

import (
	"net/http"

	"github.com/joqd/authify/internal/adapter/server/dto"
	"github.com/joqd/authify/internal/adapter/server/mapper"
	"github.com/joqd/authify/internal/core/port"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) port.UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (uh *userHandler) Register(c echo.Context) error {
	// 1. Bind request into DTO
	var req dto.RegisterRequestDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "INVALID_PAYLOAD",
			Message: err.Error(),
		})
	}

	// 2. Validate
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, dto.ErrorResponse{
			Error:   "VALIDATION_FAILED",
			Message: err.Error(),
		})
	}

	// 3. Set default for nullable fields
	req.SetDefaults()

	// 4. Map Resquest DTO -> Domain
	userDomain := mapper.RegisterRequestDTOToUserDomain(&req)

	// 5. Call service
	created, err := uh.userService.Register(c.Request().Context(), userDomain)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "INTERNAL_ERROR",
			Message: err.Error(),
		})
	}

	// 6. Map Domain -> Response DTO
	res := mapper.UserDomainToRegisterResponse(created, "User registered successfully")

	// 7. Return success
	return c.JSON(http.StatusCreated, res)
}
