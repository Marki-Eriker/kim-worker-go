package payment

import (
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func MapOneInvoiceToGqlModel(i *Invoice) *model.PaymentInvoice {
	return &model.PaymentInvoice{
		ID:         i.ID,
		ContractID: i.ContractID,
		FileID:     i.FileStorageItemID,
		CreatedAt:  i.CreatedAt,
	}
}

func MapOneConfirmationToGqlModel(c *Confirmation) *model.PaymentConfirmation {
	return &model.PaymentConfirmation{
		ID:               c.ID,
		FileID:           c.FileStorageItemID,
		PaymentInvoiceID: c.ContractPaymentInvoiceID,
		Proven:           c.Proven,
		ContractID:       c.ContractID,
		CreatedAt:        c.CreatedAt,
		UpdatedAt:        &c.UpdatedAt,
	}
}

func MapManyInvoiceToGqlModel(in []*Invoice) []*model.PaymentInvoice {
	items := make([]*model.PaymentInvoice, len(in))
	for i, v := range in {
		items[i] = MapOneInvoiceToGqlModel(v)
	}

	return items
}

func MapManyConfirmationToGqlModel(c []*Confirmation) []*model.PaymentConfirmation {
	items := make([]*model.PaymentConfirmation, len(c))
	for i, v := range c {
		items[i] = MapOneConfirmationToGqlModel(v)
	}

	return items
}
