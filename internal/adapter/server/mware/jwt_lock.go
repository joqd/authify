package mware

import (
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func JWT(secretKey string) echo.MiddlewareFunc {
	return echojwt.JWT([]byte(secretKey))
}
