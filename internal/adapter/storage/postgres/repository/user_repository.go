package repository

import (
	"context"

	"github.com/joqd/authify/internal/adapter/storage/postgres/mapper"
	"github.com/joqd/authify/internal/core/domain"
	"github.com/joqd/authify/internal/core/port"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	userModel := mapper.UserDomainToUserModel(user)

	if err := ur.db.WithContext(ctx).Create(userModel).Error; err != nil {
		return nil, err
	}

	return mapper.UserModelToUserDomain(userModel), nil
}
