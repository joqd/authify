package domain

import "time"

type AccessToken struct {
	Token     string
	ExpiresAt time.Time
	UserID    int64
}
