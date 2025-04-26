package domain

import (
	"time"
)

type UserActivity struct {
	ID        uint64
	UserID    uint64
	Event     string
	CreatedAt time.Time
}
