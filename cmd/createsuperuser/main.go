package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/joqd/authify/internal/adapter/config"
	"github.com/joqd/authify/internal/adapter/logger"
	"github.com/joqd/authify/internal/adapter/storage/postgres"
	"github.com/joqd/authify/internal/adapter/storage/postgres/repository"
	"github.com/joqd/authify/internal/core/domain"
	"github.com/joqd/authify/internal/core/service"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(raw string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	return string(hashed), err
}

func main() {
	logger, err := logger.NewLogger()
	if err != nil {
		fmt.Printf("failed to make logger: %v", err)
		os.Exit(1)
	}

	if len(os.Args) != 3 {
		logger.Error("usage: go run main.go <username> <password>")
		os.Exit(1)
	}

	username := os.Args[1]
	password := os.Args[2]

	conf := config.LoadConfig()
	db := postgres.NewPostgresDatabase(conf)

	hashedPassword, err := hashPassword(password)
	if err != nil {
		logger.Errorf("failed to hash password: %v", err)
		os.Exit(1)
	}

	userRepo := repository.NewUserRepository(db.GetDB(), logger)
	userServ := service.NewUserService(userRepo, logger)

	superuser := &domain.User{
		Username:     username,
		PasswordHash: hashedPassword,
		IsSuperuser:  true,
		IsStaff:      true,
	}

	ctx := context.Background()

	_, err = userServ.Register(ctx, superuser)
	if err != nil {
		if errors.Is(err, domain.ErrConflictingData) {
			logger.Errorf("user already exists: %v", err)
		} else {
			logger.Errorf("failed to create superuser: %v", err)
		}

		os.Exit(1)
	}

	logger.Infof("✅ Superuser '%s' created successfully.", username)
}
