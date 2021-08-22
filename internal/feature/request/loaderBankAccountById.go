package request

import (
	"context"
	"fmt"
	"time"
)

type LoaderBankAccountByIDKey struct {
	BankAccountID uint
}

func NewConfiguredLoaderBankAccountByID(repo IRepository, maxBatch int) *LoaderBankAccountByID {
	return NewLoaderBankAccountByID(LoaderBankAccountByIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderBankAccountByIDKey) ([]*BankAccount, []error) {
			items := make([]*BankAccount, len(keys))
			errors := make([]error, len(keys))
			ctx := context.Background()
			ids := getUniqueBankAccountIDs(keys)

			accounts, err := repo.GetBankAccountByIDs(ctx, ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
				return nil, errors
			}

			group := groupBankAccount(accounts)
			for i, key := range keys {
				if c, ok := group[key.BankAccountID]; ok {
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

func getUniqueBankAccountIDs(keys []LoaderBankAccountByIDKey) []uint {
	mapping := make(map[uint]bool)

	for _, key := range keys {
		mapping[key.BankAccountID] = true
	}

	ids := make([]uint, len(mapping))

	i := 0
	for key := range mapping {
		ids[i] = key
		i++
	}

	return ids
}

func groupBankAccount(c []*BankAccount) map[uint]*BankAccount {
	group := make(map[uint]*BankAccount)

	for _, v := range c {
		group[v.ID] = v
	}

	return group
}
