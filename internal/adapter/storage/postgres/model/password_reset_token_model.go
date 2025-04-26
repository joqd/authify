package model

import "time"

type PasswordResetTokenModel struct {
	Token     string    `gorm:"primaryKey;size:255;not null"`
	UserID    uint64    `gorm:"not null;index"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (PasswordResetTokenModel) TableName() string {
	return "password_reset_tokens"
}
