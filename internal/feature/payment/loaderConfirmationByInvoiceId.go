package payment

import (
	"context"
	"fmt"
	"time"
)

type LoaderConfirmationByInvoiceIDKey struct {
	InvoiceID uint
}

func NewConfiguredLoaderConfirmationByInvoiceID(repo IRepository, maxBatch int) *LoaderConfirmationByInvoiceID {
	return NewLoaderConfirmationByInvoiceID(LoaderConfirmationByInvoiceIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderConfirmationByInvoiceIDKey) ([]*Confirmation, []error) {
			items := make([]*Confirmation, len(keys))
			errors := make([]error, len(keys))
			ctx := context.Background()
			ids := getUniqueInvoiceIDs(keys)

			confirmation, err := repo.GetConfirmationByInvoiceID(ctx, ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
				return nil, errors
			}

			group := groupConfirmation(confirmation)
			for i, key := range keys {
				if c, ok := group[key.InvoiceID]; ok {
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

func getUniqueInvoiceIDs(keys []LoaderConfirmationByInvoiceIDKey) []uint {
	mapping := make(map[uint]bool)

	for _, key := range keys {
		mapping[key.InvoiceID] = true
	}

	ids := make([]uint, len(mapping))

	i := 0
	for key := range mapping {
		ids[i] = key
		i++
	}

	return ids
}

func groupConfirmation(c []*Confirmation) map[uint]*Confirmation {
	group := make(map[uint]*Confirmation)

	for _, v := range c {
		group[v.ContractPaymentInvoiceID] = v
	}

	return group
}
