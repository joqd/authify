package service

import (
	"context"

	"github.com/joqd/authify/internal/core/domain"
	"github.com/joqd/authify/internal/core/port"
	"github.com/joqd/authify/internal/core/util"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo   port.UserRepository
	logger *zap.SugaredLogger
}

func NewUserService(repo port.UserRepository, logger *zap.SugaredLogger) port.UserService {
	return &userService{
		repo:   repo,
		logger: logger,
	}
}

func (us *userService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashed, err := util.HashSecret(user.PasswordHash)
	if err != nil {
		us.logger.Errorf("failed to hash password: %v", err)
		return nil, err
	}

	user.PasswordHash = string(hashed)

	created, err := us.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (us *userService) Retrieve(ctx context.Context, id uint64) (*domain.User, error) {
	user, err := us.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userService) Delete(ctx context.Context, id uint64) error {
	// We need to check user permission for delete users
	// for example only sudo can delete superusers
	return us.repo.DeleteByID(ctx, id)
}

func (us *userService) List(ctx context.Context) ([]domain.User, error) {
	return us.repo.List(ctx)
}

func (us *userService) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	return us.repo.Update(ctx, user)
}

func (us *userService) LoginSuperuser(ctx context.Context, username string, password string) (*domain.User, error) {
	user, err := us.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if !user.IsSuperuser {
		return nil, domain.ErrForbidden
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	if err := us.repo.UpdateLastLoginByID(ctx, user.ID); err != nil {
		us.logger.Errorf("failed to update last_login: id: %d; err: %v", user.ID, err)
	}

	return user, nil
}
