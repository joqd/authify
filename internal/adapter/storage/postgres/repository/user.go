package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"github.com/joqd/authify/internal/adapter/storage/postgres/model"
	"github.com/joqd/authify/internal/core/domain"
	"github.com/joqd/authify/internal/core/port"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type userRepository struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewUserRepository(db *gorm.DB, logger *zap.SugaredLogger) port.UserRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}

func (ur *userRepository) Create(ctx context.Context, userDomain *domain.User) (*domain.User, error) {
	// Copy userDomain to userModel
	var userModel model.UserModel
	err := copier.CopyWithOption(&userModel, &userDomain, copier.Option{DeepCopy: true, IgnoreEmpty: true})
	if err != nil {
		ur.logger.Errorf("failed to copy userDomain to userModel: %v", err)
		return nil, domain.ErrCopier
	}

	// DB action
	if err := ur.db.WithContext(ctx).Create(&userModel).Error; err != nil {
		ur.logger.Errorf("failed to create user; username: %s, err: %v", userDomain.Username, err)

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, domain.ErrConflictingData
		}

		return nil, err
	}

	// Copy userModel to userDomain
	var result domain.User
	err = copier.CopyWithOption(&result, &userModel, copier.Option{DeepCopy: true, IgnoreEmpty: true})
	if err != nil {
		ur.logger.Errorf("failed to copy userModel to userDomain: %v", err)
		return nil, domain.ErrCopier
	}

	return &result, nil
}

func (ur *userRepository) GetByID(ctx context.Context, id uint64) (*domain.User, error) {
	// DB action
	var userModel model.UserModel
	if err := ur.db.WithContext(ctx).First(&userModel, id).Error; err != nil {
		ur.logger.Errorf("failed to find user; id: %d, err: %v", id, err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}

		return nil, err
	}

	// Copy userModel to userDomain
	var userDomain domain.User
	if err := copier.CopyWithOption(&userDomain, &userModel, copier.Option{DeepCopy: true}); err != nil {
		ur.logger.Errorf("failed to map userModel to userDomain: %v", err)
		return nil, domain.ErrCopier
	}

	return &userDomain, nil
}

func (ur *userRepository) DeleteByID(ctx context.Context, id uint64) error {
	// DB action
	var userModel model.UserModel
	result := ur.db.WithContext(ctx).Delete(&userModel, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrDataNotFound
	}

	return nil
}

func (ur *userRepository) List(ctx context.Context) ([]domain.User, error) {
	var usersModel []model.UserModel
	var usersDomain []domain.User

	if err := ur.db.WithContext(ctx).Find(&usersModel).Error; err != nil {
		return nil, err
	}

	for _, userModel := range usersModel {
		var userDomain domain.User
		if err := copier.CopyWithOption(&userDomain, &userModel, copier.Option{DeepCopy: true}); err != nil {
			ur.logger.Errorf("failed to map userModel to userDomain: %v", err)
			return nil, domain.ErrCopier
		}

		usersDomain = append(usersDomain, userDomain)
	}

	return usersDomain, nil
}

func (ur *userRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	var existing model.UserModel

	if err := ur.db.WithContext(ctx).First(&existing, user.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}

		ur.logger.Errorf("failed to find user by ID: %d: %v", user.ID, err)
		return nil, err
	}

	if err := copier.CopyWithOption(&existing, &user, copier.Option{DeepCopy: true, IgnoreEmpty: true}); err != nil {
		ur.logger.Errorf("failed to copy userDomain to existing userModel: %v", err)
		return nil, domain.ErrCopier
	}

	if err := ur.db.WithContext(ctx).Save(&existing).Error; err != nil {
		ur.logger.Errorf("failed to update user: %v", err)
		return nil, err
	}

	var updated domain.User
	if err := copier.CopyWithOption(&updated, &existing, copier.Option{DeepCopy: true}); err != nil {
		ur.logger.Errorf("failed to copy existing userModel to userDomain: %v", err)
		return nil, domain.ErrCopier
	}

	return &updated, nil
}

func (ur *userRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var userModel model.UserModel

	if err := ur.db.WithContext(ctx).Where("username = ?", username).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}

		return nil, err
	}

	var userDomain domain.User
	if err := copier.CopyWithOption(&userDomain, &userModel, copier.Option{DeepCopy: true}); err != nil {
		ur.logger.Errorf("failed to copy existing userModel to userDomain: %v", err)
		return nil, domain.ErrCopier
	}

	return &userDomain, nil
}


func (ur *userRepository) UpdateLastLoginByID(ctx context.Context, id uint64) error {
	if err := ur.db.WithContext(ctx).Model(&model.UserModel{}).Where("id = ?", id).Update("last_login", time.Now()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrDataNotFound
		}

		ur.logger.Errorf("failed to find user by ID: %d: %v", id, err)
		return err
	}

	return nil
}