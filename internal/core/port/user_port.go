package port

import (
	"context"

	"github.com/joqd/authify/internal/core/domain"
	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// GetUserByID(ctx context.Context, id uint64) (*domain.User, error)
	// GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	// GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	// ListUsers(ctx context.Context, skip, limit uint64) ([]domain.User, error)
	// UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// DeleteUser(ctx context.Context, id uint64) error
}

type UserService interface {
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
	// GetUser(ctx context.Context, id uint64) (*domain.User, error)
	// ListUsers(ctx context.Context, skip, limit uint64) ([]domain.User, error)
	// UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// DeleteUser(ctx context.Context, id uint64) error
}

type UserHandler interface {
	Register(c echo.Context) error
}
