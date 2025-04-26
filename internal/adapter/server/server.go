package server

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/joqd/authify/internal/adapter/config"
	"github.com/joqd/authify/internal/adapter/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)



type Server struct {
	app  *echo.Echo
	db   storage.PostgresDatabase
	conf *config.Config
}

func NewServer(conf *config.Config, db storage.PostgresDatabase) *Server {
	app := echo.New()

	app.Logger.SetLevel(log.DEBUG)
	app.Validator = &CustomValidator{validator: validator.New()}

	return &Server{
		app:  app,
		db:   db,
		conf: conf,
	}
}

func (server *Server) Start() {
	server.app.Use(middleware.Recover())
	server.app.Use(middleware.Logger())

	server.SetupRoutes()

	serverURL := fmt.Sprintf(":%d", server.conf.ServerPort)
	server.app.Logger.Fatal(server.app.Start(serverURL))
}
