package response

import (
	"time"
)

// ================================== Register User
type RegisteredUser struct {
	ID         uint64    `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email,omitempty"`
	FirstName  string    `json:"first_name,omitempty"`
	LastName   string    `json:"last_name,omitempty"`
	DateJoined time.Time `json:"date_joined"`
}

type RegisterUserResponseWrapper struct {
	BaseResponse
	Result RegisteredUser `json:"result"`
}

// ================================== Retrieve User
type RetrievedUser struct {
	ID              uint64     `json:"id"`
	Username        string     `json:"username"`
	Email           string     `json:"email,omitempty"`
	FirstName       string     `json:"first_name,omitempty"`
	LastName        string     `json:"last_name,omitempty"`
	IsStaff         bool       `json:"is_staff"`
	IsActive        bool       `json:"is_active"`
	IsSuperuser     bool       `json:"is_superuser"`
	IsEmailVerified bool       `json:"is_email_verified"`
	LastLogin       *time.Time `json:"last_login,omitempty"`
	DateJoined      time.Time  `json:"date_joined"`
}

type RetrieveUserResponseWrapper struct {
	BaseResponse
	Result RetrievedUser `json:"result"`
}

// ================================== Delete User
type DeletedUser struct {
	ID uint64 `json:"id"`
}

type DeleteUserResponseWrapper struct {
	BaseResponse
	Result DeletedUser `json:"result"`
}

// ================================== List User
type RetrievedUsers struct {
	Count int             `json:"count"`
	Users []RetrievedUser `json:"users"`
}

type ListUserResponseWrapper struct {
	BaseResponse
	Result RetrievedUsers `json:"result"`
}

// ================================== Update User
type UpdatedUser struct {
	ID          uint64     `json:"id"`
	Username    string     `json:"username"`
	Email       string     `json:"email,omitempty"`
	FirstName   string     `json:"first_name,omitempty"`
	LastName    string     `json:"last_name,omitempty"`
	IsStaff     bool       `json:"is_staff"`
	IsActive    bool       `json:"is_active"`
	IsSuperuser bool       `json:"is_superuser"`
	LastLogin   *time.Time `json:"last_login,omitempty"`
	DateJoined  time.Time  `json:"date_joined"`
}

type UpdateUserResponseWrapper struct {
	BaseResponse
	Result UpdatedUser `json:"result"`
}
