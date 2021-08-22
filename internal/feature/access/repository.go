package access

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v10"
)

var ErrItemNotFound = errors.New("item not found")

type IRepository interface {
	GetAccessForUser(ctx context.Context, userID uint) ([]*Access, error)
}

type Repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAccessForUser(ctx context.Context, userID uint) ([]*Access, error) {
	var access []*Access

	userAccess := r.db.ModelContext(ctx, (*UserAccess)(nil)).ColumnExpr("access_id").Where("user_id = ?", userID)
	err := r.db.ModelContext(ctx, &access).Where("id IN (?)", userAccess).Select()

	if len(access) == 0 {
		return nil, ErrItemNotFound
	}

	return access, err
}
