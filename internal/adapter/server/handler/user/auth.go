package user

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joqd/authify/internal/adapter/server/dto/request"
	"github.com/joqd/authify/internal/adapter/server/dto/response"
	"github.com/joqd/authify/internal/adapter/server/handler"
	"github.com/joqd/authify/internal/core/domain"
	"github.com/labstack/echo/v4"
)

func (uh *userHandler) Login(c echo.Context) error {
	// Bind request into DTO
	var req request.LoginRequest
	if err := c.Bind(&req); err != nil {
		return handler.RespondError(c, http.StatusBadRequest, response.DescBadRequest)
	}

	// Validate
	if err := c.Validate(&req); err != nil {
		uh.logger.Errorf("failed to validate RegisterUserRequest: %v", err)
		return handler.RespondError(c, http.StatusUnprocessableEntity, response.DescValidationFailed)
	}

	user, err := uh.userService.Authenticate(c.Request().Context(), req.Username, req.Password)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			return handler.RespondError(c, http.StatusUnauthorized, response.DescUnauthorized)
		}

		return handler.RespondError(c, http.StatusInternalServerError, response.DescInternalError)
	}


	// Issue JWT
	claims := jwt.MapClaims{
		"sub": strconv.Itoa(int(user.ID)),
		"sup": user.IsSuperuser,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(uh.conf.JWTSecretKey))
	if err != nil {
		return handler.RespondError(c, http.StatusInternalServerError, response.DescInternalError)
	}

	var loginedUser response.LoginedUser
	loginedUser.AccessToken = signedToken

	return handler.RespondSuccess(c, http.StatusOK, loginedUser)
}
