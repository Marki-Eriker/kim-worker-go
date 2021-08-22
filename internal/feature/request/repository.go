package request

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

type IRepository interface {
	GetRequest(ctx context.Context, requestID uint) (*Request, error)
	GetRequestByID(ctx context.Context, ids []uint) ([]*Request, error)
	GetAllRequest(ctx context.Context, input *model.RequestListInput, serviceTypes []uint) ([]*Request, int, error)
	GetServiceTypeByIDs(ctx context.Context, ids []uint) ([]*ServiceType, error)
	GetContractorByIDs(ctx context.Context, ids []uint) ([]*Contractor, error)
	GetOrganizationContactsByIDs(ctx context.Context, ids []uint) ([]*OrganizationContact, error)
	GetBankAccountByIDs(ctx context.Context, ids []uint) ([]*BankAccount, error)
	GetSignatoryByIDs(ctx context.Context, ids []uint) ([]*Signatory, error)
	GetShipByRequestIDs(ctx context.Context, ids []uint) ([]*Ship, error)
	GetRequestToShip(ctx context.Context, ids []uint) ([]*ShipRequest, error)
	UpdateRequestStatus(ctx context.Context, id uint, status model.RequestStatus) (*Request, error)
	FindPersonByID(ctx context.Context, id uint) (*Person, error)
}

type Repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetRequest(ctx context.Context, requestID uint) (*Request, error) {
	var request Request
	err := r.db.ModelContext(ctx, &request).Where("id = ?", requestID).Select()
	return &request, err
}

func (r *Repository) GetRequestByID(ctx context.Context, ids []uint) ([]*Request, error) {
	var request []*Request
	err := r.db.ModelContext(ctx, &request).Where("id IN (?)", pg.In(ids)).Select()
	return request, err
}

func (r *Repository) GetAllRequest(ctx context.Context, input *model.RequestListInput, serviceTypes []uint) ([]*Request, int, error) {
	var requests []*Request
	limit := *input.Filter.PageSize
	offset := (*input.Filter.Page - 1) * *input.Filter.PageSize
	order := fmt.Sprintf("%v %v", *input.Filter.OrderField, input.Filter.OrderBy)

	q := r.db.ModelContext(ctx, &requests).Order(order).Limit(limit).Offset(offset)
	q.Where("service_type_id IN (?)", pg.In(serviceTypes))

	if input.Status != nil {
		q.Where("status = (?)", input.Status)
	}

	count, err := q.SelectAndCount()

	return requests, count, err
}

func (r *Repository) GetServiceTypeByIDs(ctx context.Context, ids []uint) ([]*ServiceType, error) {
	var types []*ServiceType
	err := r.db.ModelContext(ctx, &types).Where("id IN (?)", pg.In(ids)).Select()
	return types, err
}

func (r *Repository) GetContractorByIDs(ctx context.Context, ids []uint) ([]*Contractor, error) {
	var contractors []*Contractor
	err := r.db.ModelContext(ctx, &contractors).Where("id IN (?)", pg.In(ids)).Select()
	return contractors, err
}

func (r *Repository) GetOrganizationContactsByIDs(ctx context.Context, ids []uint) ([]*OrganizationContact, error) {
	var contacts []*OrganizationContact
	err := r.db.ModelContext(ctx, &contacts).Where("organization_id IN (?)", pg.In(ids)).Select()
	return contacts, err
}

func (r *Repository) GetBankAccountByIDs(ctx context.Context, ids []uint) ([]*BankAccount, error) {
	var accounts []*BankAccount
	err := r.db.ModelContext(ctx, &accounts).Where("id IN (?)", pg.In(ids)).Select()
	return accounts, err
}

func (r *Repository) GetSignatoryByIDs(ctx context.Context, ids []uint) ([]*Signatory, error) {
	var signatory []*Signatory
	err := r.db.ModelContext(ctx, &signatory).Where("id IN (?)", pg.In(ids)).Select()
	return signatory, err
}

func (r *Repository) GetShipByRequestIDs(ctx context.Context, ids []uint) ([]*Ship, error) {
	var ships []*Ship
	shipsIDs := r.db.ModelContext(ctx, (*ShipRequest)(nil)).ColumnExpr("ship_id").Where("service_request_id IN (?)", pg.In(ids))
	err := r.db.ModelContext(ctx, &ships).Where("id IN (?)", shipsIDs).Select()
	return ships, err
}

func (r *Repository) GetRequestToShip(ctx context.Context, ids []uint) ([]*ShipRequest, error) {
	var shipRequest []*ShipRequest
	err := r.db.ModelContext(ctx, &shipRequest).Where("service_request_id IN (?)", pg.In(ids)).Select()
	return shipRequest, err
}

func (r *Repository) UpdateRequestStatus(ctx context.Context, id uint, status model.RequestStatus) (*Request, error) {
	var request Request
	request.Status = status
	_, err := r.db.ModelContext(ctx, &request).Column("status").Where("id = ?", id).Returning("*").Update()
	return &request, err
}

func (r *Repository) FindPersonByID(ctx context.Context, id uint) (*Person, error) {
	var person Person
	err := r.db.ModelContext(ctx, &person).Where("id = ?", id).First()
	return &person, err
}
