package user

import (
	"github.com/joqd/authify/internal/adapter/config"
	"github.com/joqd/authify/internal/core/port"
	"go.uber.org/zap"
)

type userHandler struct {
	userService port.UserService
	logger      *zap.SugaredLogger
	conf        *config.Config
}

func NewUserHandler(userService port.UserService, logger *zap.SugaredLogger, conf *config.Config) port.UserHandler {
	return &userHandler{
		userService: userService,
		logger:      logger,
		conf:        conf,
	}
}
