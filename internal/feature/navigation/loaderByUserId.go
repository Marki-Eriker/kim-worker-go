package navigation

import (
	"context"
	"fmt"
	"time"
)

type LoaderByUserIDKey struct {
	UserID uint
}

func NewConfiguredLoaderByUserID(repo IRepository, maxBatch int) *LoaderByUserID {
	return NewLoaderByUserID(LoaderByUserIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderByUserIDKey) ([][]*Navigation, []error) {
			items := make([][]*Navigation, len(keys))
			errors := make([]error, len(keys))
			ctx := context.Background()
			ids := getUniqueUserIDs(keys)

			navigation, err := repo.GetNavigationForUserIDs(ctx, ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
				return nil, errors
			}

			rawUserToNavigation, err := repo.GetUserToNavigation(ctx, ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
				return nil, errors
			}

			userToNavigation := getUniqueUserToNavigation(rawUserToNavigation)

			groups := groupNavigationByUserID(navigation, userToNavigation)
			for i, key := range keys {
				if p, ok := groups[key.UserID]; ok {
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

func getUniqueUserToNavigation(un []*UserNavigation) []*UserNavigation {
	keys := make(map[UserNavigation]bool)
	var list []*UserNavigation

	for _, v := range un {
		if _, ok := keys[*v]; !ok {
			keys[*v] = true
			list = append(list, v)
		}
	}

	return list
}

func getUniqueUserIDs(keys []LoaderByUserIDKey) []uint {
	mapping := make(map[uint]bool)

	for _, key := range keys {
		mapping[key.UserID] = true
	}

	ids := make([]uint, len(mapping))

	i := 0
	for key := range mapping {
		ids[i] = key
		i++
	}

	return ids
}

func groupNavigationByUserID(n []*Navigation, un []*UserNavigation) map[uint][]*Navigation {
	groups := make(map[uint][]*Navigation)

	for _, v := range n {
		for _, j := range un {
			if v.ID == j.NavigationID {
				groups[j.UserID] = append(groups[j.UserID], v)
			}
		}
	}

	return groups
}
