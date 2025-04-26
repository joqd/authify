package model

type GroupModel struct {
	ID          uint              `gorm:"primaryKey;autoIncrement"`
	Name        string            `gorm:"size:255;not null;unique"`
	Permissions []PermissionModel `gorm:"many2many:group_permissions;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}

func (GroupModel) TableName() string {
	return "groups"
}
