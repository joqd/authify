package model

type PermissionModel struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"size:255;not null;unique"`
	CodeName    string  `gorm:"size:255;not null;unique"`
	Description *string `gorm:"type:text"`
}

func (PermissionModel) TableName() string {
	return "permissions"
}
