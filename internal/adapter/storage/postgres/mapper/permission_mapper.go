package mapper

import (
	"github.com/joqd/authify/internal/adapter/storage/postgres/model"
	"github.com/joqd/authify/internal/core/domain"
)

func PermissionModelToPermissionDomain(in *model.PermissionModel) *domain.Permission {
	return &domain.Permission{
		ID:          in.ID,
		Name:        in.Name,
		CodeName:    in.CodeName,
		Description: in.Description,
	}
}

func PermissionsModelToPermissionsDomain(in []model.PermissionModel) []domain.Permission {
	permissions := make([]domain.Permission, len(in))
	for i, p := range in {
		permissions[i] = *PermissionModelToPermissionDomain(&p)
	}

	return permissions
}

func PermissionDomainToPermissionModel(in *domain.Permission) *model.PermissionModel {
	return &model.PermissionModel{
		ID:          in.ID,
		Name:        in.Name,
		CodeName:    in.CodeName,
		Description: in.Description,
	}
}

func PermissionsDomainToPermissionsModel(in []domain.Permission) []model.PermissionModel {
	permissions := make([]model.PermissionModel, len(in))
	for i, p := range in {
		permissions[i] = *PermissionDomainToPermissionModel(&p)
	}

	return permissions
}
