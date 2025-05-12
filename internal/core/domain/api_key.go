package domain

import (
	"time"
)

type APIKey struct {
	ID        uint
	Key       *string
	KeyHash   string
	Owner     string
	Name      string
	IsActive  bool
	ExpiresAt time.Time
	CreatedAt time.Time
}
