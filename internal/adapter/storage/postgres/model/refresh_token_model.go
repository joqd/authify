package model

import "time"

type RefreshTokenModel struct {
	Token     string    `gorm:"primaryKey;size:255;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	UserID    uint64    `gorm:"not null;index"`
}

func (RefreshTokenModel) TableName() string {
	return "refresh_tokens"
}
