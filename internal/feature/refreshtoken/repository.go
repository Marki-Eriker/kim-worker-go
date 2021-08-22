package refreshtoken

import (
	"context"

	"github.com/go-pg/pg/v10"
)

type IRepository interface {
	Save(ctx context.Context, entity RefreshToken) error
	GetByValue(ctx context.Context, value string) (RefreshToken, error)
	GetByUserId(ctx context.Context, id uint) ([]RefreshToken, error)
	DeleteByIds(ctx context.Context, ids []uint) error
	DeleteByValue(ctx context.Context, value string) error
}

type Repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(ctx context.Context, entity RefreshToken) error {
	_, err := r.db.ModelContext(ctx, &entity).Insert()
	return err
}

func (r *Repository) GetByValue(ctx context.Context, value string) (RefreshToken, error) {
	var token RefreshToken
	err := r.db.ModelContext(ctx, &token).Where("token = ?", value).First()
	return token, err
}

func (r *Repository) GetByUserId(ctx context.Context, id uint) ([]RefreshToken, error) {
	var tokens []RefreshToken
	err := r.db.ModelContext(ctx, &tokens).Where("user_id = ?", id).OrderExpr("created_at DESC").Select()
	return tokens, err
}

func (r *Repository) DeleteByIds(ctx context.Context, ids []uint) error {
	_, err := r.db.ModelContext(ctx, (*RefreshToken)(nil)).Where("id IN (?)", pg.In(ids)).Delete()
	return err
}

func (r *Repository) DeleteByValue(ctx context.Context, value string) error {
	_, err := r.db.ModelContext(ctx, (*RefreshToken)(nil)).Where("token = ?", value).Delete()
	return err
}
