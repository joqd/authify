package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/joqd/authify/internal/adapter/server/dto/request"
	"github.com/joqd/authify/internal/adapter/server/dto/response"
	"github.com/joqd/authify/internal/core/domain"
	"github.com/joqd/authify/internal/core/port"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type userHandler struct {
	userService port.UserService
	logger      *zap.SugaredLogger
}

func NewUserHandler(userService port.UserService, logger *zap.SugaredLogger) port.UserHandler {
	return &userHandler{
		userService: userService,
		logger:      logger,
	}
}

// @Summary Register user
// @Description Register new user
// @Tags user
// @Produce json
// @Success 201 {object} response.RegisterUserResponseWrapper
// @Failure 400 {object} response.ErrorResponse
// @Failure 422 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users [post]
func (uh *userHandler) Register(c echo.Context) error {
	// Bind request into DTO
	var req request.RegisterUserRequest
	if err := c.Bind(&req); err != nil {
		return RespondError(c, http.StatusBadRequest, response.DescBadRequest)
	}

	// Validate
	if err := c.Validate(&req); err != nil {
		uh.logger.Errorf("failed to validate RegisterUserRequest: %v", err)
		return RespondError(c, http.StatusUnprocessableEntity, response.DescValidationFailed)
	}

	// Set default for nullable fields
	req.SetDefaults()

	// Copy req to userDomain
	var userDomain domain.User
	if err := copier.Copy(&userDomain, &req); err != nil {
		return RespondError(c, http.StatusInternalServerError, response.DescInternalError)
	}

	// Call service
	created, err := uh.userService.Register(c.Request().Context(), &userDomain)
	if err != nil {
		if errors.Is(err, domain.ErrConflictingData) {
			return RespondError(c, http.StatusConflict, response.DescObjectExists)
		}

		return RespondError(c, http.StatusInternalServerError, response.DescInternalError)
	}

	// Copy userDomain to RegisteredUser
	var registeredUser response.RegisteredUser
	if err := copier.Copy(&registeredUser, &created); err != nil {
		return RespondError(c, http.StatusInternalServerError, response.DescInternalError)
	}

	// Return success
	return RespondSuccess(c, http.StatusCreated, registeredUser)
}

// @Summary Retrive user
// @Description Retrive user by id
// @Tags user
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.RetrieveUserResponseWrapper
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users/{id} [get]
func (uh *userHandler) Retrieve(c echo.Context) error {
	// Bind request into DTO
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return RespondError(c, http.StatusBadRequest, response.DescBadRequest)
	}

	// Call service
	user, err := uh.userService.Retrieve(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrDataNotFound) {
			return RespondError(c, http.StatusNotFound, response.DescNotFound)
		}

		return RespondError(c, http.StatusInternalServerError, response.DescInternalError)
	}

	// Copy userDomain to RetrievedUser
	var retrievedUser response.RetrievedUser
	if err := copier.Copy(&retrievedUser, &user); err != nil {
		return RespondError(c, http.StatusInternalServerError, response.DescInternalError)
	}

	// Return success
	return RespondSuccess(c, http.StatusOK, retrievedUser)
}

// @Summary Delete user
// @Description Delete user by id
// @Tags user
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.DeleteUserResponseWrapper
// @Failure 400 {object} response.ErrorResponse
// @Failure 204 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users/{id} [delete]
func (uh *userHandler) Delete(c echo.Context) error {
	// Bind request into DTO
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return RespondError(c, http.StatusBadRequest, response.DescBadRequest)
	}

	// Call service
	if err := uh.userService.Delete(c.Request().Context(), id); err != nil {
		if errors.Is(err, domain.ErrDataNotFound) {
			return RespondError(c, http.StatusNoContent, response.DescNoContent)
		}

		return RespondError(c, http.StatusInternalServerError, response.DescInternalError)
	}

	// Copy req to DeletedUser
	var deletedUser response.DeletedUser
	deletedUser.ID = id

	// Return success
	return RespondSuccess(c, http.StatusOK, deletedUser)
}

// @Summary List of users
// @Description List of users
// @Tags user
// @Produce json
// @Success 200 {object} response.ListUserResponseWrapper
// @Failure 500 {object} response.ErrorResponse
// @Router /users/ [get]
func (uh *userHandler) List(c echo.Context) error {
	// Call service
	usersDomain, err := uh.userService.List(c.Request().Context())
	if err != nil {
		return RespondError(c, http.StatusInternalServerError, response.DescInternalError)
	}

	// Copy usersDomain to RetrievedUsers
	var retrievedUsers response.RetrievedUsers
	retrievedUsers.Count = len(usersDomain)

	for _, userDomain := range usersDomain {
		var retrievedUser response.RetrievedUser
		if err := copier.Copy(&retrievedUser, &userDomain); err != nil {
			return RespondError(c, http.StatusInternalServerError, response.DescInternalError)
		}

		retrievedUsers.Users = append(retrievedUsers.Users, retrievedUser)
	}

	// Return success
	return RespondSuccess(c, http.StatusOK, retrievedUsers)
}

// @Summary Update user
// @Description Update user
// @Tags user
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.UpdateUserResponseWrapper
// @Failure 400 {object} response.ErrorResponse
// @Failure 422 {object} response.ErrorResponse
// @Failure 204 {object} response.ErrorResponse
// @Failure 409 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users/{id} [put]
func (uh *userHandler) Update(c echo.Context) error {
	// Bind request into DTO
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return RespondError(c, http.StatusBadRequest, response.DescBadRequest)
	}

	var req request.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return RespondError(c, http.StatusBadRequest, response.DescBadRequest)
	}

	// Validate
	if err := c.Validate(&req); err != nil {
		uh.logger.Errorf("failed to validate UpdateUserRequest: %v", err)
		return RespondError(c, http.StatusUnprocessableEntity, response.DescValidationFailed)
	}

	// Copy req to userDomain
	var userDomain domain.User
	if err := copier.Copy(&userDomain, &req); err != nil {
		return RespondError(c, http.StatusInternalServerError, response.DescInternalError)
	}

	userDomain.ID = id

	// Call service
	updated, err := uh.userService.Update(c.Request().Context(), &userDomain)
	if err != nil {
		if errors.Is(err, domain.ErrDataNotFound) {
			return RespondError(c, http.StatusNoContent, response.DescNoContent)
		}

		if errors.Is(err, domain.ErrConflictingData) {
			return RespondError(c, http.StatusConflict, response.DescConflict)
		}

		return RespondError(c, http.StatusInternalServerError, response.DescInternalError)
	}

	// Copy userDomain to UpdatedUserResponse
	var updatedUser response.UpdatedUser
	if err := copier.Copy(&updatedUser, &updated); err != nil {
		return RespondError(c, http.StatusInternalServerError, response.DescInternalError)
	}

	// Return success
	return RespondSuccess(c, http.StatusOK, updatedUser)
}
