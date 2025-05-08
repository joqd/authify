package port

import (
	"context"

	"github.com/joqd/authify/internal/core/domain"
	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	GetByID(ctx context.Context, id uint64) (*domain.User, error)
	DeleteByID(ctx context.Context, id uint64) error
	List(ctx context.Context) ([]domain.User, error)
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	UpdateLastLoginByID(ctx context.Context, id uint64) error
}

type UserService interface {
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
	Retrieve(ctx context.Context, id uint64) (*domain.User, error)
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context) ([]domain.User, error)
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	Authenticate(ctx context.Context, username string, password string) (*domain.User, error)
}

type UserHandler interface {
	Register(c echo.Context) error
	Retrieve(c echo.Context) error
	Delete(c echo.Context) error
	List(c echo.Context) error
	Update(c echo.Context) error
	Login(c echo.Context) error
}
