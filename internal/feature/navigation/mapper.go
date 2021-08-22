package navigation

import "github.com/marki-eriker/kim-worker-go/internal/gql/model"

func MapOneToGqlModel(n Navigation) *model.Navigation {
	return &model.Navigation{
		ID:          n.ID,
		Path:        n.Path,
		Title:       n.Title,
		Description: &n.Description,
		Icon:        &n.Icon,
		ParentID:    n.ParentId,
		Order:       n.Order,
		Node:        n.Node,
		Dev:         n.Dev,
	}
}

func MapManyToGqlModels(n []*Navigation) []*model.Navigation {
	items := make([]*model.Navigation, len(n))
	for i, v := range n {
		items[i] = MapOneToGqlModel(*v)
	}

	return items
}
