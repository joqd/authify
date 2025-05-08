package model

import (
	"time"

	"github.com/google/uuid"
)

type APIKeyModel struct {
	ID         uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	Key        string    `gorm:"size:255;not null"`
	SecretHash string    `gorm:"size:255;not null"`
	Owner      string    `gorm:"not null;index"`
	Name       string    `gorm:"not null"`
	IsActive   bool      `gorm:"default:true"`
	ExpiresAt  time.Time `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

func (APIKeyModel) TableName() string {
	return "api_keys"
}
