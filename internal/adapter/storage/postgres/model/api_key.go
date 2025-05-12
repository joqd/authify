package model

import (
	"time"
)

type APIKeyModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	KeyHash   string    `gorm:"size:255;not null"`
	Owner     string    `gorm:"not null;index"`
	Name      string    `gorm:"not null"`
	IsActive  bool      `gorm:"default:true"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (APIKeyModel) TableName() string {
	return "api_keys"
}
