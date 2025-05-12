package server

import (
	_ "github.com/joqd/authify/docs"
	"github.com/joqd/authify/internal/adapter/server/handler/apikey"
	"github.com/joqd/authify/internal/adapter/server/handler/user"
	"github.com/joqd/authify/internal/adapter/server/mware"

	// "github.com/joqd/authify/internal/adapter/server/middleware"
	"github.com/joqd/authify/internal/adapter/storage/postgres/repository"
	"github.com/joqd/authify/internal/core/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) SetupRoutes() {
	s.app.GET("/swagger/*", echoSwagger.WrapHandler)
	s.app.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	// Public api v1
	api := s.app.Group("/api/v1")
	s.SetupUserRouter(api)
	s.SetupAPIKeyRouter(api)
}

func (s *Server) SetupUserRouter(api *echo.Group) {
	userRepo := repository.NewUserRepository(s.db.GetDB(), s.logger)
	userService := service.NewUserService(userRepo, s.logger)
	userHandler := user.NewUserHandler(userService, s.logger, s.conf)

	apiKeyRepo := repository.NewApiKeyRepository(s.db.GetDB(), s.logger)
	apiKeyService := service.NewApiKeyService(apiKeyRepo, s.logger)
	apiKeyMiddleware := mware.NewApiKeyMiddleware(apiKeyService, s.logger)

	api.POST("/login", userHandler.Login)

	users := api.Group("/users", middleware.KeyAuth(apiKeyMiddleware.Validator))
	users.GET("", userHandler.List)
	users.GET("/:id", userHandler.Retrieve)
	users.PUT("/:id", userHandler.Update)
	users.DELETE("/:id", userHandler.Delete)
	users.POST("", userHandler.Register)
}

func (s *Server) SetupAPIKeyRouter(api *echo.Group) {
	apiKeyRepo := repository.NewApiKeyRepository(s.db.GetDB(), s.logger)
	apiKeyService := service.NewApiKeyService(apiKeyRepo, s.logger)
	apiKeyHandler := apikey.NewApiKeyHandler(apiKeyService, s.logger)

	apiKeys := api.Group("/api-keys")
	apiKeys.POST("", apiKeyHandler.Create)
}
