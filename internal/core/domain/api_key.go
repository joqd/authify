package domain

import "time"

type APIKey struct {
	Key       string
	Secret    string
	Owner     string
	CreatedAt time.Time
	ExpiresAt time.Time
}
