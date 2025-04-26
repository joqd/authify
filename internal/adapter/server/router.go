package server

import (
	"github.com/joqd/authify/internal/adapter/server/handler"
	"github.com/joqd/authify/internal/adapter/storage/postgres/repository"
	"github.com/joqd/authify/internal/core/service"
	"github.com/labstack/echo/v4"
)

func (server *Server) SetupRoutes() {
	server.app.GET("health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	// User Router
	userRepo := repository.NewUserRepository(server.db.GetDB())
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	users := server.app.Group("/users")
	users.POST("/register", userHandler.Register)
}
