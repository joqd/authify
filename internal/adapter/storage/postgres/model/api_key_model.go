package model

import "time"

type APIKeyModel struct {
	Key       string    `gorm:"primaryKey;size:255;not null"`
	Secret    string    `gorm:"not null;size:255"`
	Owner     string    `gorm:"not null;index"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (APIKeyModel) TableName() string {
	return "api_keys"
}
