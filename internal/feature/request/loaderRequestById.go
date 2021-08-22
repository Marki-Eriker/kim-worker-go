package request

import (
	"context"
	"fmt"
	"time"
)

type LoaderRequestByIDKey struct {
	RequestID uint
}

func NewConfiguredLoaderRequestByID(repo IRepository, maxBatch int) *LoaderRequestByID {
	return NewLoaderRequestByID(LoaderRequestByIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderRequestByIDKey) ([]*Request, []error) {
			items := make([]*Request, len(keys))
			errors := make([]error, len(keys))
			ctx := context.Background()
			ids := getUniqueRequestIDs(keys)

			request, err := repo.GetRequestByID(ctx, ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
				return nil, errors
			}

			group := groupRequest(request)
			for i, key := range keys {
				if st, ok := group[key.RequestID]; ok {
					items[i] = st
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

func getUniqueRequestIDs(keys []LoaderRequestByIDKey) []uint {
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

func groupRequest(r []*Request) map[uint]*Request {
	group := make(map[uint]*Request)

	for _, v := range r {
		group[v.ID] = v
	}

	return group
}
