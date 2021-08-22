package file

import (
	"context"
	"github.com/go-pg/pg/v10"
)

type IRepository interface {
	Save(ctx context.Context, entity *File) error
}

type Repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(ctx context.Context, entity *File) error {
	_, err := r.db.ModelContext(ctx, entity).Returning("id").Insert()
	return err
}
