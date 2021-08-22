package request

import (
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"time"
)

type Request struct {
	tableName struct{} `pg:"lk_dev.service_request,discard_unknown_columns"`

	ID                       uint
	ServiceTypeID            uint
	ContractorID             uint
	OrganizationContactID    uint
	ContractMediumType       model.ContractMediumType
	ContractFilledTemplateID uint
	Status                   model.RequestStatus
	CreatedAt                time.Time
	BankAccountID            uint
	SignatoryID              uint
}

type ServiceType struct {
	tableName struct{} `pg:"lk_dev.service_type"`

	ID   uint
	Name string
}

type Contractor struct {
	tableName struct{} `pg:"lk_dev.contractor"`

	ID             uint
	ContractorType model.ContractorType
	FullName       string
	ShortName      string
	PersonID       uint
}

type Person struct {
	tableName struct{} `pg:"lk_dev.person"`

	ID    uint
	Email string
	Phone string
}

type OrganizationContact struct {
	tableName struct{} `pg:"lk_dev.organization_contact"`

	ID    uint `pg:"organization_id"`
	Phone string
	Email string
}

type BankAccount struct {
	tableName struct{} `pg:"lk_dev.bank_account"`

	ID                  uint
	AccountNumber       string
	CorrespondentNumber string `pg:"correspondent_account_number"`
	Bik                 string
	BankName            string
}

type Signatory struct {
	tableName struct{} `pg:"lk_dev.contractor_signatory"`

	ID            uint
	Name          string
	ActingBasis   string
	WarrantNumber string
	WarrantDate   time.Time
}

type Ship struct {
	tableName struct{} `pg:"lk_dev.ship"`

	ID                              uint
	Name                            string
	HullNumber                      string
	ProjectNumber                   string
	Length                          float64
	Width                           float64
	HullHeight                      float64
	Cubic                           float64
	Flag                            string
	ShipConfirmParamsCertificateIds []uint `pg:",array"`
	OwnerShipRightsCertificateIds   []uint `pg:",array"`
}

type ShipRequest struct {
	tableName struct{} `pg:"lk_dev.service_request_ship"`

	ServiceRequestID uint `pg:"service_request_id"`
	ShipID           uint `pg:"ship_id"`
}
