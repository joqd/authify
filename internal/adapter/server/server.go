package server

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joqd/authify/internal/adapter/config"
	"github.com/joqd/authify/internal/adapter/logger"
	"github.com/joqd/authify/internal/adapter/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type Server struct {
	app    *echo.Echo
	logger *zap.SugaredLogger
	db     storage.PostgresDatabase
	conf   *config.Config
}

func NewServer(conf *config.Config, db storage.PostgresDatabase) *Server {
	app := echo.New()

	app.Logger.SetLevel(log.DEBUG)
	app.Validator = &CustomValidator{validator: validator.New()}

	logger, err := logger.NewLogger()
	if err != nil {
		log.Errorf("Failed to make new logger: %v", err)
		os.Exit(1)
	}

	return &Server{
		app:    app,
		logger: logger,
		db:     db,
		conf:   conf,
	}
}

func (s *Server) Start() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(20))))

	s.SetupRoutes()

	serverURL := fmt.Sprintf(":%d", s.conf.ServerPort)
	s.app.Logger.Fatal(s.app.Start(serverURL))
}
