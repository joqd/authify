package domain

import "time"

type EmailVerificationToken struct {
	Token     string
	UserID    uint64
	ExpiresAt time.Time
	CreatedAt time.Time
}
