package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

type IRepository interface {
	Save(ctx context.Context, entity *User) (uint, error)
	GetAll(ctx context.Context, filter *model.PaginationInput) ([]*User, int, error)
	GetById(ctx context.Context, id uint) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	DeleteByID(ctx context.Context, id uint) error
	Update(ctx context.Context, entity *User) error
}

type Repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(ctx context.Context, entity *User) (uint, error) {
	_, err := r.db.ModelContext(ctx, entity).Returning("id").Insert()
	return 0, err
}

func (r *Repository) GetAll(ctx context.Context, filter *model.PaginationInput) ([]*User, int, error) {
	var users []*User
	limit := *filter.PageSize
	offset := (*filter.Page - 1) * *filter.PageSize
	order := fmt.Sprintf("%v %v", *filter.OrderField, filter.OrderBy)

	count, err := r.db.ModelContext(ctx, &users).Order(order).Limit(limit).Offset(offset).SelectAndCount()
	return users, count, err
}

func (r *Repository) GetById(ctx context.Context, id uint) (User, error) {
	var user User
	err := r.db.ModelContext(ctx, &user).Where("id = ?", id).First()
	return user, err
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := r.db.ModelContext(ctx, &user).Where("email = ?", email).First()
	return user, err
}

func (r *Repository) DeleteByID(ctx context.Context, id uint) error {
	res, err := r.db.ModelContext(ctx, (*User)(nil)).Where("id = ?", id).Delete()
	if res != nil && res.RowsAffected() != 1 {
		return errors.New("no user to delete")
	}
	return err
}

func (r *Repository) Update(ctx context.Context, entity *User) error {
	_, err := r.db.ModelContext(ctx, entity).Where("id = ?", entity.ID).Update()
	return err
}
