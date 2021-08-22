package request

import (
	"context"
	"fmt"
	"time"
)

type LoaderShipByRequestIDKey struct {
	RequestID uint
}

func NewConfiguredLoaderShipByRequestID(repo IRepository, maxBatch int) *LoaderShipByRequestID {
	return NewLoaderShipByRequestID(LoaderShipByRequestIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderShipByRequestIDKey) ([][]*Ship, []error) {
			items := make([][]*Ship, len(keys))
			errors := make([]error, len(keys))
			ctx := context.Background()
			ids := getUniqueRequestID(keys)

			ships, err := repo.GetShipByRequestIDs(ctx, ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
				return nil, errors
			}

			requestToShip, err := repo.GetRequestToShip(ctx, ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
				return nil, errors
			}

			groups := groupShipByRequestID(ships, requestToShip)
			for i, key := range keys {
				if p, ok := groups[key.RequestID]; ok {
					items[i] = p
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

func getUniqueRequestID(keys []LoaderShipByRequestIDKey) []uint {
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

func groupShipByRequestID(s []*Ship, sr []*ShipRequest) map[uint][]*Ship {
	groups := make(map[uint][]*Ship)

	for _, v := range s {
		for _, j := range sr {
			if v.ID == j.ShipID {
				groups[j.ServiceRequestID] = append(groups[j.ServiceRequestID], v)
			}
		}
	}

	return groups
}
