package domain

import "time"

type APIKey struct {
	Key       string
	Secret    string
	Owner     string
	ExpiresAt time.Time
	CreatedAt time.Time
}
