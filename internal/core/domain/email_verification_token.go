package domain

import "time"

type EmailVerificationToken struct {
	Token     string
	UserID    int64
	ExpiresAt time.Time
	CreatedAt time.Time
}
