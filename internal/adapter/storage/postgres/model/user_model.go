package model

import (
	"time"

	"github.com/joqd/authify/internal/core/domain"
)

type UserModel struct {
	ID              uint64            `gorm:"primaryKey;autoIncrement"`
	Username        string            `gorm:"size:255;unique;not null"`
	FirstName       *string           `gorm:"size:255"`
	LastName        *string           `gorm:"size:255"`
	Email           *string           `gorm:"size:255;unique"`
	PasswordHash    string            `gorm:"not null"`
	Groups          []GroupModel      `gorm:"many2many:user_groups;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
	Permissions     []PermissionModel `gorm:"many2many:user_permissions;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
	IsStaff         bool              `gorm:"default:false"`
	IsActive        bool              `gorm:"default:true"`
	IsSuperuser     bool              `gorm:"default:false"`
	IsEmailVerified bool              `gorm:"default:false"`
	LastLogin       *time.Time        `gorm:"type:timestamp"`
	DateJoined      time.Time         `gorm:"autoCreateTime"`
}

func (UserModel) TableName() string {
	return "users"
}

func DBToDomainUser(user *UserModel) *domain.User {
	permissions := make([]domain.Permission, len(user.Permissions))
	groups := make([]domain.Group, len(user.Groups))

	for i, p := range user.Permissions {
		permissions[i] = *DBToDomainPermission(&p)
	}

	for i, g := range user.Groups {
		groups[i] = *DBToDomainGroup(&g)
	}

	return &domain.User{
		ID:              user.ID,
		Username:        user.Username,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		PasswordHash:    user.PasswordHash,
		Groups:          groups,
		Permissions:     permissions,
		IsStaff:         user.IsStaff,
		IsActive:        user.IsActive,
		IsSuperuser:     user.IsSuperuser,
		IsEmailVerified: user.IsEmailVerified,
		LastLogin:       user.LastLogin,
		DateJoined:      user.DateJoined,
	}
}

func DomainToDBUser(user *domain.User) *UserModel {
	permissions := make([]PermissionModel, len(user.Permissions))
	groups := make([]GroupModel, len(user.Groups))

	for i, p := range user.Permissions {
		permissions[i] = *DomainToDBPermission(&p)
	}

	for i, g := range user.Groups {
		groups[i] = *DomainToDBGroup(&g)
	}

	return &UserModel{
		ID:              user.ID,
		Username:        user.Username,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		PasswordHash:    user.PasswordHash,
		Groups:          groups,
		Permissions:     permissions,
		IsStaff:         user.IsStaff,
		IsActive:        user.IsActive,
		IsSuperuser:     user.IsSuperuser,
		IsEmailVerified: user.IsEmailVerified,
		LastLogin:       user.LastLogin,
		DateJoined:      user.DateJoined,
	}
}
