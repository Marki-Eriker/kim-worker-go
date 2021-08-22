package payment

import (
	"context"
	"fmt"
	"time"
)

type LoaderInvoiceByContractIDKey struct {
	ContractId uint
}

func NewConfiguredLoaderInvoiceByContractID(repo IRepository, maxBatch int) *LoaderInvoiceByContractID {
	return NewLoaderInvoiceByContractID(LoaderInvoiceByContractIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderInvoiceByContractIDKey) ([][]*Invoice, []error) {
			items := make([][]*Invoice, len(keys))
			errors := make([]error, len(keys))
			ctx := context.Background()
			ids := getUniqueContractId(keys)

			invoice, err := repo.GetInvoiceByContractID(ctx, ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
				return nil, errors
			}

			group := groupInvoice(invoice)
			for i, key := range keys {
				if c, ok := group[key.ContractId]; ok {
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

func getUniqueContractId(keys []LoaderInvoiceByContractIDKey) []uint {
	mapping := make(map[uint]bool)

	for _, key := range keys {
		mapping[key.ContractId] = true
	}

	ids := make([]uint, len(mapping))

	i := 0
	for key := range mapping {
		ids[i] = key
		i++
	}

	return ids
}

func groupInvoice(c []*Invoice) map[uint][]*Invoice {
	group := make(map[uint][]*Invoice)

	for _, v := range c {
		group[v.ContractID] = append(group[v.ContractID], v)
	}

	return group
}
