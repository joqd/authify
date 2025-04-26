package mapper

import (
	"github.com/joqd/authify/internal/adapter/server/dto"
	"github.com/joqd/authify/internal/core/domain"
)

func RegisterRequestDTOToUserDomain(in *dto.RegisterRequestDTO) *domain.User {
	return &domain.User{
		Username:     in.Username,
		FirstName:    in.FirstName,
		LastName:     in.LastName,
		Email:        in.Email,
		PasswordHash: in.Password, // still not hashed
		IsStaff:      *in.IsStaff,
		IsActive:     *in.IsActive,
		IsSuperuser:  *in.IsSuperuser,
	}
}

func UserDomainToRegisterResponse(in *domain.User, message string) *dto.RegisterResponseDTO {
	return &dto.RegisterResponseDTO{
		ID:        in.ID,
		Username:  in.Username,
		FirstName: *in.FirstName,
		LastName:  *in.LastName,
		Email:     *in.Email,
		Message:   message,
	}
}
