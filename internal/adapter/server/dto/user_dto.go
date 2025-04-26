package dto

type RegisterRequestDTO struct {
	Username    string  `json:"username" validate:"required,alphanum,max=30,min=3"`
	FirstName   *string `json:"first_name" validate:"omitempty,alpha,max=30,min=2"`
	LastName    *string `json:"last_name" validate:"omitempty,alpha,max=30,min=2"`
	Email       *string `json:"email" validate:"omitempty,email"`
	Password    string  `json:"password" validate:"required,max=64,min=8"`
	IsStaff     *bool   `json:"is_staff" validate:"omitempty"`
	IsActive    *bool   `json:"is_active" validate:"omitempty"`
	IsSuperuser *bool   `json:"is_superuser" validate:"omitempty"`
}

func (ru *RegisterRequestDTO) SetDefaults() {
	if ru.FirstName == nil {
		def := ""
		ru.FirstName = &def
	}
	if ru.LastName == nil {
		def := ""
		ru.LastName = &def
	}
	if ru.Email == nil {
		def := ""
		ru.Email = &def
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

type RegisterResponseDTO struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Message   string `json:"message"`
}
