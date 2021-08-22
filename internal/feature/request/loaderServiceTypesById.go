package request

import (
	"context"
	"fmt"
	"time"
)

type LoaderServiceTypesByIDKey struct {
	ServiceTypeID uint
}

func NewConfiguredLoaderServiceTypesByID(repo IRepository, maxBatch int) *LoaderServiceTypesByID {
	return NewLoaderServiceTypesByID(LoaderServiceTypesByIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderServiceTypesByIDKey) ([]*ServiceType, []error) {
			items := make([]*ServiceType, len(keys))
			errors := make([]error, len(keys))
			ctx := context.Background()
			ids := getUniqueServiceTypeIDs(keys)

			serviceTypes, err := repo.GetServiceTypeByIDs(ctx, ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
				return nil, errors
			}

			group := groupServiceType(serviceTypes)
			for i, key := range keys {
				if st, ok := group[key.ServiceTypeID]; ok {
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

func getUniqueServiceTypeIDs(keys []LoaderServiceTypesByIDKey) []uint {
	mapping := make(map[uint]bool)

	for _, key := range keys {
		mapping[key.ServiceTypeID] = true
	}

	ids := make([]uint, len(mapping))

	i := 0
	for key := range mapping {
		ids[i] = key
		i++
	}

	return ids
}

func groupServiceType(st []*ServiceType) map[uint]*ServiceType {
	group := make(map[uint]*ServiceType)

	for _, v := range st {
		group[v.ID] = v
	}

	return group
}
