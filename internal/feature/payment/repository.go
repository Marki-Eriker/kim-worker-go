package payment

import (
	"context"
	"github.com/go-pg/pg/v10"
)

type IRepository interface {
	SaveInvoice(ctx context.Context, entity *Invoice) error
	SaveConfirmation(ctx context.Context, entity *Confirmation) error
	GetConfirmationByInvoiceID(ctx context.Context, ids []uint) ([]*Confirmation, error)
	GetInvoiceByContractID(ctx context.Context, ids []uint) ([]*Invoice, error)
	UpdateConfirmationProven(ctx context.Context, id uint) (*Confirmation, error)
}

type Repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SaveInvoice(ctx context.Context, entity *Invoice) error {
	_, err := r.db.ModelContext(ctx, entity).Insert()
	return err
}

func (r *Repository) SaveConfirmation(ctx context.Context, entity *Confirmation) error {
	_, err := r.db.ModelContext(ctx, entity).Insert()
	return err
}

func (r *Repository) GetConfirmationByInvoiceID(ctx context.Context, ids []uint) ([]*Confirmation, error) {
	var confirmation []*Confirmation
	err := r.db.ModelContext(ctx, &confirmation).Where("contract_payment_invoice_id IN (?)", pg.In(ids)).Select()
	return confirmation, err
}

func (r *Repository) GetInvoiceByContractID(ctx context.Context, ids []uint) ([]*Invoice, error) {
	var invoice []*Invoice
	err := r.db.ModelContext(ctx, &invoice).Where("contract_id IN (?)", pg.In(ids)).Order("created_at DESC").Select()
	return invoice, err
}

func (r *Repository) UpdateConfirmationProven(ctx context.Context, id uint) (*Confirmation, error) {
	var confirmation Confirmation
	_, err := r.db.ModelContext(ctx, &confirmation).Set("proven = ?", true).Where("id = ?", id).Returning("*").Update()
	return &confirmation, err
}
