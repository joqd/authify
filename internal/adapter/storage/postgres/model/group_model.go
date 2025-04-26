package model

import "github.com/joqd/authify/internal/core/domain"

type GroupModel struct {
	ID          uint              `gorm:"primaryKey;autoIncrement"`
	Name        string            `gorm:"size:255;not null;unique"`
	Permissions []PermissionModel `gorm:"many2many:group_permissions;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}

func (GroupModel) TableName() string {
	return "groups"
}

func DBToDomainGroup(group *GroupModel) *domain.Group {
	permissions := make([]domain.Permission, len(group.Permissions))
	for i, p := range group.Permissions {
		permissions[i] = *DBToDomainPermission(&p)
	}

	return &domain.Group{
		ID:          group.ID,
		Name:        group.Name,
		Permissions: permissions,
	}
}

func DomainToDBGroup(group *domain.Group) *GroupModel {
	permissions := make([]PermissionModel, len(group.Permissions))
	for i, p := range group.Permissions {
		permissions[i] = *DomainToDBPermission(&p)
	}

	return &GroupModel{
		ID:          group.ID,
		Name:        group.Name,
		Permissions: permissions,
	}
}
