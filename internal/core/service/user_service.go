package service

import (
	"context"

	"github.com/joqd/authify/internal/core/domain"
	"github.com/joqd/authify/internal/core/port"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) port.UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = string(hashed)

	created, err := us.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return created, nil
}
