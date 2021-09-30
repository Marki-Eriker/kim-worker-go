package contract

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/marki-eriker/kim-worker-go/internal/feature/payment"
	"github.com/marki-eriker/kim-worker-go/internal/feature/request"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

type IRepository interface {
	GetAllContract(ctx context.Context, input *model.ContractListInput, serviceTypes []uint) ([]*Contract, int, error)
	GetContractByRequestID(ctx context.Context, ids []uint) ([]*Contract, error)
	Save(ctx context.Context, entity *Contract) error
	FindByRequestID(ctx context.Context, id uint) (*Contract, error)
	UpdateFileID(ctx context.Context, contract *Contract) error
	FindById(ctx context.Context, id uint) (*Contract, error)
}

type Repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllContract(ctx context.Context, input *model.ContractListInput, serviceTypes []uint) ([]*Contract, int, error) {
	var contract []*Contract
	limit := *input.Filter.PageSize
	offset := (*input.Filter.Page - 1) * *input.Filter.PageSize
	order := fmt.Sprintf("%v %v", *input.Filter.OrderField, input.Filter.OrderBy)

	requestIDs := r.db.ModelContext(ctx, (*request.Request)(nil)).ColumnExpr("id")
	requestIDs.Where("service_type_id IN (?)", pg.In(serviceTypes))
	requestIDs.Where("status = ?", model.RequestStatusCompleted)

	var contractIDs *orm.Query

	q := r.db.ModelContext(ctx, &contract).Order(order).Limit(limit).Offset(offset)

	if *input.PaymentFilter != model.PaymentFilterAll {
		contractIDs = r.db.ModelContext(ctx, (*Contract)(nil)).ColumnExpr("id")

		switch *input.PaymentFilter {
		case model.PaymentFilterNotPaid:
			invoiceWithConfirmIDs := r.db.ModelContext(ctx, (*payment.Confirmation)(nil))
			invoiceWithConfirmIDs.ColumnExpr("contract_payment_invoice_id")

			contractWithoutPaymentConfirmationIDs := r.db.ModelContext(ctx, (*payment.Invoice)(nil))
			contractWithoutPaymentConfirmationIDs.ColumnExpr("contract_id")
			contractWithoutPaymentConfirmationIDs.Where("id NOT IN (?)", invoiceWithConfirmIDs)

			contractIDs.Where("id IN (?)", contractWithoutPaymentConfirmationIDs)

			q.Where("id IN (?)", contractIDs)
		case model.PaymentFilterNotVerified:
			contractWitNotVerifiedPaymentConfirmationIDS := r.db.ModelContext(ctx, (*payment.Confirmation)(nil))
			contractWitNotVerifiedPaymentConfirmationIDS.ColumnExpr("contract_id")
			contractWitNotVerifiedPaymentConfirmationIDS.Where("proven = ?", false)

			contractIDs.Where("id IN (?)", contractWitNotVerifiedPaymentConfirmationIDS)
			q.Where("id IN (?)", contractIDs)
		}
	}

	q.Where("service_request_id IN (?)", requestIDs)
	count, err := q.SelectAndCount()

	return contract, count, err
}

func (r *Repository) GetContractByRequestID(ctx context.Context, ids []uint) ([]*Contract, error) {
	var contracts []*Contract
	err := r.db.ModelContext(ctx, &contracts).Where("service_request_id IN (?)", pg.In(ids)).Select()
	return contracts, err
}

func (r *Repository) Save(ctx context.Context, entity *Contract) error {
	_, err := r.db.ModelContext(ctx, entity).Returning("id").Insert()
	return err
}

func (r *Repository) FindByRequestID(ctx context.Context, id uint) (*Contract, error) {
	var contract Contract
	err := r.db.ModelContext(ctx, &contract).Where("service_request_id = ?", id).First()
	return &contract, err
}

func (r *Repository) UpdateFileID(ctx context.Context, contract *Contract) error {
	_, err := r.db.ModelContext(ctx, contract).Column("file_storage_item_id", "number").Where("id = ?", contract.ID).Update()
	return err
}

func (r *Repository) FindById(ctx context.Context, id uint) (*Contract, error) {
	var contract Contract
	err := r.db.ModelContext(ctx, &contract).Where("id = ?", id).First()
	return &contract, err
}
