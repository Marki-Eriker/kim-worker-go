package request

import (
	"context"
	"fmt"
	"time"
)

type LoaderSignatoryByIDKey struct {
	SignatoryID uint
}

func NewConfiguredLoaderSignatoryByID(repo IRepository, maxBatch int) *LoaderSignatoryByID {
	return NewLoaderSignatoryByID(LoaderSignatoryByIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderSignatoryByIDKey) ([]*Signatory, []error) {
			items := make([]*Signatory, len(keys))
			errors := make([]error, len(keys))
			ctx := context.Background()
			ids := getUniqueSignatoryIDs(keys)

			signatory, err := repo.GetSignatoryByIDs(ctx, ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
				return nil, errors
			}

			group := groupSignatory(signatory)
			for i, key := range keys {
				if c, ok := group[key.SignatoryID]; ok {
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

func getUniqueSignatoryIDs(keys []LoaderSignatoryByIDKey) []uint {
	mapping := make(map[uint]bool)

	for _, key := range keys {
		mapping[key.SignatoryID] = true
	}

	ids := make([]uint, len(mapping))

	i := 0
	for key := range mapping {
		ids[i] = key
		i++
	}

	return ids
}

func groupSignatory(c []*Signatory) map[uint]*Signatory {
	group := make(map[uint]*Signatory)

	for _, v := range c {
		group[v.ID] = v
	}

	return group
}
