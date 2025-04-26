package model

import (
	"time"

	"github.com/joqd/authify/internal/core/domain"
)

type AccessTokenModel struct {
	Token     string    `gorm:"primaryKey;size:255;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	UserID    uint64    `gorm:"not null;index"`
}

func (AccessTokenModel) TableName() string {
	return "access_tokens"
}

func (a *AccessTokenModel) FromDomain(data *domain.AccessToken) *AccessTokenModel {
	a.Token = data.Token
	a.ExpiresAt = data.ExpiresAt
	a.UserID = data.UserID
	return a
}
