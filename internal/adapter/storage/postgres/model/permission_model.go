package model

import "github.com/joqd/authify/internal/core/domain"

type PermissionModel struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"size:255;not null;unique"`
	CodeName    string  `gorm:"size:255;not null;unique"`
	Description *string `gorm:"type:text"`
}

func (PermissionModel) TableName() string {
	return "permissions"
}

func DBToDomainPermission(permission *PermissionModel) *domain.Permission {
	return &domain.Permission{
		ID:          permission.ID,
		Name:        permission.Name,
		CodeName:    permission.CodeName,
		Description: permission.Description,
	}
}

func DomainToDBPermission(permission *domain.Permission) *PermissionModel {
	return &PermissionModel{
		ID:          permission.ID,
		Name:        permission.Name,
		CodeName:    permission.CodeName,
		Description: permission.Description,
	}
}
