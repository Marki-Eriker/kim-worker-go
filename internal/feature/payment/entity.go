package payment

import "time"

type Invoice struct {
	tableName struct{} `pg:"contract_payment_invoice"`

	ID                uint
	ContractID        uint
	FileStorageItemID uint
	CreatedAt         time.Time
}

type Confirmation struct {
	tableName struct{} `pg:"contract_payment_confirmation"`

	ID                       uint
	FileStorageItemID        uint
	ContractPaymentInvoiceID uint
	Proven                   bool
	ContractID               uint
	CreatedAt                time.Time
	UpdatedAt                time.Time
}
