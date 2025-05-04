package server

import (
	"github.com/joqd/authify/internal/adapter/server/handler"
	"github.com/joqd/authify/internal/adapter/storage/postgres/repository"
	"github.com/joqd/authify/internal/core/service"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/joqd/authify/docs"
)


func (s *Server) SetupRoutes() {
	s.app.GET("/swagger/*", echoSwagger.WrapHandler)
	s.app.GET("health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	s.SetupUserRouter()
}

func (s *Server) SetupUserRouter() {
	userRepo := repository.NewUserRepository(s.db.GetDB(), s.logger)
	userService := service.NewUserService(userRepo, s.logger)
	userHandler := handler.NewUserHandler(userService, s.logger)

	users := s.app.Group("/users")
	users.GET("/", userHandler.List)
	users.GET("/:id", userHandler.Retrieve)
	users.PUT("/:id", userHandler.Update)
	users.DELETE("/:id", userHandler.Delete)
	users.POST("/", userHandler.Register)
}
