package mapper

import (
	"github.com/joqd/authify/internal/adapter/storage/postgres/model"
	"github.com/joqd/authify/internal/core/domain"
)

func UserModelToUserDomain(in *model.UserModel) *domain.User {
	return &domain.User{
		ID:              in.ID,
		Username:        in.Username,
		FirstName:       in.FirstName,
		LastName:        in.LastName,
		Email:           in.Email,
		PasswordHash:    in.PasswordHash,
		Groups:          GroupsModelToGroupsDomain(in.Groups),
		Permissions:     PermissionsModelToPermissionsDomain(in.Permissions),
		IsStaff:         in.IsStaff,
		IsActive:        in.IsActive,
		IsSuperuser:     in.IsSuperuser,
		IsEmailVerified: in.IsEmailVerified,
		LastLogin:       in.LastLogin,
		DateJoined:      in.DateJoined,
	}
}

func UserDomainToUserModel(in *domain.User) *model.UserModel {
	return &model.UserModel{
		ID:              in.ID,
		Username:        in.Username,
		FirstName:       in.FirstName,
		LastName:        in.LastName,
		Email:           in.Email,
		PasswordHash:    in.PasswordHash,
		Groups:          GroupsDomainToGroupsModel(in.Groups),
		Permissions:     PermissionsDomainToPermissionsModel(in.Permissions),
		IsStaff:         in.IsStaff,
		IsActive:        in.IsActive,
		IsSuperuser:     in.IsSuperuser,
		IsEmailVerified: in.IsEmailVerified,
		LastLogin:       in.LastLogin,
		DateJoined:      in.DateJoined,
	}
}
