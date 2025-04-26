package domain

import "time"

type RefreshToken struct {
	Token     string
	ExpiresAt time.Time
	UserID    uint64
}
