package domain

import "time"

type PasswordResetToken struct {
	Token     string
	UserID    uint64
	ExpiresAt time.Time
	CreatedAt time.Time
}
