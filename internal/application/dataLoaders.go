package application

import (
	"github.com/marki-eriker/kim-worker-go/internal/feature/contract"
	"github.com/marki-eriker/kim-worker-go/internal/feature/navigation"
	"github.com/marki-eriker/kim-worker-go/internal/feature/payment"
	"github.com/marki-eriker/kim-worker-go/internal/feature/request"
)

type DataLoaders struct {
	NavigationLoaderByUserID *navigation.LoaderByUserID
	ServiceTypeLoaderByID    *request.LoaderServiceTypesByID
	ContractorLoaderByID     *request.LoaderContractorByID
	OrganizationContactByID  *request.LoaderOrganizationContactByID
	BankAccountLoaderByID    *request.LoaderBankAccountByID
	SignatoryLoaderByID      *request.LoaderSignatoryByID
	ShipLoaderByRequestID    *request.LoaderShipByRequestID
	ContractByRequestID      *contract.LoaderContractByRequestID
	ConfirmationByInvoiceID  *payment.LoaderConfirmationByInvoiceID
	InvoiceByContractID      *payment.LoaderInvoiceByContractID
	RequestLoaderByID        *request.LoaderRequestByID
}

func NewDataLoaders(repos *Repositories) *DataLoaders {
	return &DataLoaders{
		NavigationLoaderByUserID: navigation.NewConfiguredLoaderByUserID(repos.NavigationRepository, 50),
		ServiceTypeLoaderByID:    request.NewConfiguredLoaderServiceTypesByID(repos.RequestRepository, 50),
		ContractorLoaderByID:     request.NewConfiguredLoaderContractorByID(repos.RequestRepository, 50),
		OrganizationContactByID:  request.NewConfiguredLoaderOrganizationContactById(repos.RequestRepository, 50),
		BankAccountLoaderByID:    request.NewConfiguredLoaderBankAccountByID(repos.RequestRepository, 50),
		SignatoryLoaderByID:      request.NewConfiguredLoaderSignatoryByID(repos.RequestRepository, 50),
		ShipLoaderByRequestID:    request.NewConfiguredLoaderShipByRequestID(repos.RequestRepository, 50),
		ContractByRequestID:      contract.NewConfiguredLoaderContractByRequestID(repos.ContractRepository, 50),
		ConfirmationByInvoiceID:  payment.NewConfiguredLoaderConfirmationByInvoiceID(repos.PaymentRepository, 50),
		InvoiceByContractID:      payment.NewConfiguredLoaderInvoiceByContractID(repos.PaymentRepository, 50),
		RequestLoaderByID:        request.NewConfiguredLoaderRequestByID(repos.RequestRepository, 50),
	}
}
