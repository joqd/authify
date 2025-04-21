package domain

import "time"

type User struct {
	ID              int64
	Username        string
	FirstName       *string
	LastName        *string
	Email           *string
	PasswordHash    string
	Groups          []Group
	Permissions     []Permission
	IsStaff         bool
	IsActive        bool
	IsSuperuser     bool
	IsEmailVerified bool
	LastLogin       *time.Time
	DateJoined      time.Time
}
