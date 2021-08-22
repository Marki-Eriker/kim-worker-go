package access

import "github.com/marki-eriker/kim-worker-go/internal/gql/model"

func MapOneToGqlModel(a Access) *model.Access {
	return &model.Access{
		ID:   a.ID,
		Name: a.Name,
	}
}

func MapManyToGqlModel(a []*Access) []*model.Access {
	items := make([]*model.Access, len(a))
	for i, v := range a {
		items[i] = MapOneToGqlModel(*v)
	}

	return items
}
