package domain

import "time"

type UserActivity struct {
	ID        int64
	UserID    int64
	Event     string
	CreatedAt time.Time
}
