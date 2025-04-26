package domain

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              uint64
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

func (u *User) GetFullName() string {
	if u.FirstName == nil && u.LastName == nil {
		return u.Username
	}

	return fmt.Sprintf("%s %s", *u.FirstName, *u.LastName)
}

func (u *User) CheckPassword(rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(rawPassword))
	return err == nil
}
