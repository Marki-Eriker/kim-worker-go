package payment

import (
	"context"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"time"
)

type IService interface {
	CreateInvoice(input *model.PaymentInvoiceCreateInput) (*Invoice, error)
	CreateConfirmation(input *model.PaymentConfirmationCreateInput) (*Confirmation, error)
	ApproveConfirmation(input *model.PaymentConfirmationApproveInput) (*Confirmation, error)
}

type Service struct {
	repository IRepository
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateInvoice(input *model.PaymentInvoiceCreateInput) (*Invoice, error) {
	invoice := Invoice{
		ContractID:        input.ContractID,
		FileStorageItemID: input.FileID,
		CreatedAt:         time.Now(),
	}

	if err := s.repository.SaveInvoice(context.Background(), &invoice); err != nil {
		return nil, err
	}

	return &invoice, nil
}

func (s *Service) CreateConfirmation(input *model.PaymentConfirmationCreateInput) (*Confirmation, error) {
	confirmation := Confirmation{
		FileStorageItemID:        input.FileID,
		ContractPaymentInvoiceID: input.InvoiceID,
		Proven:                   false,
		ContractID:               input.ContractID,
		CreatedAt:                time.Now(),
	}

	if err := s.repository.SaveConfirmation(context.Background(), &confirmation); err != nil {
		return nil, err
	}

	return &confirmation, nil
}

func (s *Service) ApproveConfirmation(input *model.PaymentConfirmationApproveInput) (*Confirmation, error) {
	return s.repository.UpdateConfirmationProven(context.Background(), input.ConfirmationID)
}
