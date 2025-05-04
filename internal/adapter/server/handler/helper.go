package handler

import (
	"github.com/joqd/authify/internal/adapter/server/dto/response"
	"github.com/labstack/echo/v4"
)

func RespondError(c echo.Context, status int, desc string) error {
	return c.JSON(status, response.ErrorResponse{
		BaseResponse: response.BaseResponse{Ok: false},
		ErrorCode:    status,
		Description:  desc,
	})
}

func RespondSuccess[T any](c echo.Context, status int, result T) error {
	return c.JSON(status, response.Response[T]{
		BaseResponse: response.BaseResponse{Ok: true},
		Result:       &result,
	})
}
