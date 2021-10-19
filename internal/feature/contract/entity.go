package contract

import "time"

type Contract struct {
	tableName struct{} `pg:"contract"`

	ID                uint
	ServiceRequestID  uint
	Number            string
	CreatedAt         time.Time
	ContractorID      uint
	FileStorageItemID uint
}
