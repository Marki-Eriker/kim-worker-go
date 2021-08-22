package user

import (
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func MapOneToGqlModel(u *User) *model.User {
	return &model.User{
		ID:           u.ID,
		Email:        u.Email,
		FullName:     u.FullName,
		BaseRole:     u.BaseRole,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		ServiceTypes: u.ServiceTypes,
	}
}

func MapManyToGqlModels(u []*User) []*model.User {
	items := make([]*model.User, len(u))
	for i, v := range u {
		items[i] = MapOneToGqlModel(v)
	}

	return items
}
