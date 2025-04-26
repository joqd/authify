package model

import (
	"time"
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
