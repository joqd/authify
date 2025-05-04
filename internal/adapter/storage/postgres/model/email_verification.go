package model

import "time"

type EmailVerificationTokenModel struct {
	Token     string    `gorm:"primaryKey;size:255;not null"`
	UserID    uint64    `gorm:"not null;index"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (EmailVerificationTokenModel) TableName() string {
	return "email_verification_tokens"
}
