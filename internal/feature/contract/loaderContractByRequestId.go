package contract

import (
	"context"
	"fmt"
	"time"
)

type LoaderContractByRequestIDKey struct {
	RequestID uint
}

func NewConfiguredLoaderContractByRequestID(repo IRepository, maxBatch int) *LoaderContractByRequestID {
	return NewLoaderContractByRequestID(LoaderContractByRequestIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderContractByRequestIDKey) ([][]*Contract, []error) {
			items := make([][]*Contract, len(keys))
			errors := make([]error, len(keys))
			ctx := context.Background()
			ids := getUniqueRequestIDs(keys)

			contracts, err := repo.GetContractByRequestID(ctx, ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
				return nil, errors
			}

			group := groupContracts(contracts)
			for i, key := range keys {
				if c, ok := group[key.RequestID]; ok {
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

func getUniqueRequestIDs(keys []LoaderContractByRequestIDKey) []uint {
	mapping := make(map[uint]bool)

	for _, key := range keys {
		mapping[key.RequestID] = true
	}

	ids := make([]uint, len(mapping))

	i := 0
	for key := range mapping {
		ids[i] = key
		i++
	}

	return ids
}

func groupContracts(c []*Contract) map[uint][]*Contract {
	group := make(map[uint][]*Contract)

	for _, v := range c {
		group[v.ServiceRequestID] = append(group[v.ServiceRequestID], v)
	}

	return group
}
