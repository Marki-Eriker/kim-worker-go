package request

import (
	"context"
	"fmt"
	"time"
)

type LoaderContractorByIDKey struct {
	ContractorID uint
}

func NewConfiguredLoaderContractorByID(repo IRepository, maxBatch int) *LoaderContractorByID {
	return NewLoaderContractorByID(LoaderContractorByIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderContractorByIDKey) ([]*Contractor, []error) {
			items := make([]*Contractor, len(keys))
			errors := make([]error, len(keys))
			ctx := context.Background()
			ids := getUniqueContractorIDs(keys)

			contractors, err := repo.GetContractorByIDs(ctx, ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
				return nil, errors
			}

			group := groupContractor(contractors)
			for i, key := range keys {
				if c, ok := group[key.ContractorID]; ok {
					items[i] = c
				}
			}

			for i, v := range items {
				if v == nil {
					errors[i] = fmt.Errorf("item not found")
				}
			}

			return items, errors
		},
	})
}

func getUniqueContractorIDs(keys []LoaderContractorByIDKey) []uint {
	mapping := make(map[uint]bool)

	for _, key := range keys {
		mapping[key.ContractorID] = true
	}

	ids := make([]uint, len(mapping))

	i := 0
	for key := range mapping {
		ids[i] = key
		i++
	}

	return ids
}

func groupContractor(c []*Contractor) map[uint]*Contractor {
	group := make(map[uint]*Contractor)

	for _, v := range c {
		group[v.ID] = v
	}

	return group
}
