package contract

import (
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func MapOneContractToGqlModel(c *Contract) *model.Contract {
	return &model.Contract{
		ID:                c.ID,
		ServiceRequestID:  c.ServiceRequestID,
		Number:            c.Number,
		CreatedAt:         c.CreatedAt,
		ContractorID:      c.ContractorID,
		FileStorageItemID: c.FileStorageItemID,
	}
}

func MapManyContractToGqlModel(c []*Contract) []*model.Contract {
	items := make([]*model.Contract, len(c))
	for i, v := range c {
		items[i] = MapOneContractToGqlModel(v)
	}

	return items
}
