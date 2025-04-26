package mapper

import (
	"github.com/joqd/authify/internal/adapter/storage/postgres/model"
	"github.com/joqd/authify/internal/core/domain"
)

func GroupModelToGroupDomain(in *model.GroupModel) *domain.Group {
	return &domain.Group{
		ID:          in.ID,
		Name:        in.Name,
		Permissions: PermissionsModelToPermissionsDomain(in.Permissions),
	}
}

func GroupsModelToGroupsDomain(in []model.GroupModel) []domain.Group {
	groups := make([]domain.Group, len(in))
	for i, g := range in {
		groups[i] = *GroupModelToGroupDomain(&g)
	}

	return groups
}

func GroupDomainToGroupModel(in *domain.Group) *model.GroupModel {
	return &model.GroupModel{
		ID:          in.ID,
		Name:        in.Name,
		Permissions: PermissionsDomainToPermissionsModel(in.Permissions),
	}
}

func GroupsDomainToGroupsModel(in []domain.Group) []model.GroupModel {
	groups := make([]model.GroupModel, len(in))
	for i, g := range in {
		groups[i] = *GroupDomainToGroupModel(&g)
	}

	return groups
}
