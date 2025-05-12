package request

import "time"

// ================================== Register User
// RegisterUserRequest defines register payload
// swagger:model RegisterUserRequest
type RegisterUserRequest struct {
	// Username of the user
	// required: true
	Username string `json:"username" validate:"required,alphanum,max=30,min=3" example:"admin"`

	FirstName *string `json:"first_name" validate:"omitempty,alpha,max=30,min=2" example:"john"`
	LastName  *string `json:"last_name" validate:"omitempty,alpha,max=30,min=2" example:"doe"`

	// Email of the user
	// required: true
	Email string `json:"email" validate:"required,email" example:"some@domain.com"`

	// Password of the user
	// required: true
	Password string `json:"password" validate:"required,max=64,min=8" example:"strongpassword"`

	IsStaff     *bool `json:"is_staff" validate:"omitempty" example:"false"`
	IsActive    *bool `json:"is_active" validate:"omitempty" example:"true"`
	IsSuperuser *bool `json:"is_superuser" validate:"omitempty" example:"false"`
}

func (ru *RegisterUserRequest) SetDefaults() {
	if ru.FirstName == nil {
		def := ""
		ru.FirstName = &def
	}
	if ru.LastName == nil {
		def := ""
		ru.LastName = &def
	}
	if ru.IsStaff == nil {
		def := false
		ru.IsStaff = &def
	}
	if ru.IsActive == nil {
		def := true
		ru.IsActive = &def
	}
	if ru.IsSuperuser == nil {
		def := false
		ru.IsSuperuser = &def
	}
}

// ================================== Delete User
type DeleteUserRequest struct {
	ID uint64 `json:"id" validate:"required"`
}

// ================================== Update User
type UpdateUserRequest struct {
	Username        *string    `json:"username" validate:"omitempty,alphanum,max=30,min=3"`
	FirstName       *string    `json:"first_name" validate:"omitempty,alpha,max=30,min=2"`
	LastName        *string    `json:"last_name" validate:"omitempty,alpha,max=30,min=2"`
	Email           *string    `json:"email" validate:"omitempty,email"`
	IsStaff         *bool      `json:"is_staff" validate:"omitempty"`
	IsActive        *bool      `json:"is_active" validate:"omitempty"`
	IsSuperuser     *bool      `json:"is_superuser" validate:"omitempty"`
	IsEmailVerified *bool      `json:"is_email_verified" validate:"omitempty"`
	LastLogin       *time.Time `json:"last_login" validate:"omitempty"`
	DateJoined      *time.Time `json:"date_joined" validate:"omitempty"`
}

// ================================== Login User
type LoginRequest struct {
	Username string `json:"username" validate:"required,alphanum,max=30,min=3"`
	Password string `json:"password" validate:"required,max=64,min=8"`
}
