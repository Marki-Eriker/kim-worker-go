package request

import (
	"context"
	"fmt"
	"time"
)

type LoaderOrganizationContactByIDKey struct {
	OrganizationContactId uint
}

func NewConfiguredLoaderOrganizationContactById(repo IRepository, maxBatch int) *LoaderOrganizationContactByID {
	return NewLoaderOrganizationContactByID(LoaderOrganizationContactByIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderOrganizationContactByIDKey) ([]*OrganizationContact, []error) {
			items := make([]*OrganizationContact, len(keys))
			errors := make([]error, len(keys))
			ctx := context.Background()
			ids := getUniqueOrganizationContactIDs(keys)

			organizationContacts, err := repo.GetOrganizationContactsByIDs(ctx, ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
				return nil, errors
			}

			group := groupOrganizationContact(organizationContacts)
			for i, key := range keys {
				if c, ok := group[key.OrganizationContactId]; ok {
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

func getUniqueOrganizationContactIDs(keys []LoaderOrganizationContactByIDKey) []uint {
	mapping := make(map[uint]bool)

	for _, key := range keys {
		mapping[key.OrganizationContactId] = true
	}

	ids := make([]uint, len(mapping))

	i := 0
	for key := range mapping {
		ids[i] = key
		i++
	}

	return ids
}

func groupOrganizationContact(c []*OrganizationContact) map[uint]*OrganizationContact {
	group := make(map[uint]*OrganizationContact)

	for _, v := range c {
		group[v.ID] = v
	}

	return group
}
