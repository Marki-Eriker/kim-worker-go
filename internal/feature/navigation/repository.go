package navigation

import (
	"context"
	"errors"
	"github.com/marki-eriker/kim-worker-go/internal/feature/access"

	"github.com/go-pg/pg/v10"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

var ErrItemNotFound = errors.New("item not found")

type IRepository interface {
	GetNavigationForAccess(ctx context.Context, access *model.Access) ([]*Navigation, error)
	GetNavigationForAccessIDs(ctx context.Context, accessIDs []uint) ([]Navigation, error)
	GetAccessToNavigation(ctx context.Context) ([]*AccessNavigation, error)
	GetNavigationForUserID(ctx context.Context, userID uint) ([]*Navigation, error)
	GetNavigationForUserIDs(ctx context.Context, userIDs []uint) ([]*Navigation, error)
	GetUserToNavigation(ctx context.Context, userIDs []uint) ([]*UserNavigation, error)
}

type Repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetNavigationForAccess(ctx context.Context, access *model.Access) ([]*Navigation, error) {
	var navigation []*Navigation

	accessNavigation := r.db.ModelContext(ctx, (*AccessNavigation)(nil)).ColumnExpr("navigation_id").Where("access_id = ?", access.ID)
	err := r.db.ModelContext(ctx, &navigation).Where("id IN (?)", accessNavigation).Select()

	return navigation, err
}

func (r *Repository) GetNavigationForAccessIDs(ctx context.Context, accessIDs []uint) ([]Navigation, error) {
	var navigation []Navigation

	accessNavigation := r.db.ModelContext(ctx, (*AccessNavigation)(nil)).ColumnExpr("navigation_id").Where("access_id IN (?)", pg.In(accessIDs))
	err := r.db.ModelContext(ctx, &navigation).Where("id IN (?)", accessNavigation).Select()

	return navigation, err
}

func (r *Repository) GetAccessToNavigation(ctx context.Context) ([]*AccessNavigation, error) {
	var accessToNavigation []*AccessNavigation
	err := r.db.ModelContext(ctx, &accessToNavigation).Select()
	return accessToNavigation, err
}

func (r *Repository) GetNavigationForUserID(ctx context.Context, userID uint) ([]*Navigation, error) {
	var navigation []*Navigation

	userAccess := r.db.ModelContext(ctx, (*access.UserAccess)(nil)).ColumnExpr("access_id").Where("user_id = ?", userID)
	accessNavigation := r.db.ModelContext(ctx, (*AccessNavigation)(nil)).ColumnExpr("navigation_id").Where("access_id IN (?)", userAccess)
	err := r.db.ModelContext(ctx, &navigation).Where("id IN (?)", accessNavigation).Select()

	if len(navigation) == 0 {
		return nil, ErrItemNotFound
	}

	return navigation, err
}

func (r *Repository) GetNavigationForUserIDs(ctx context.Context, userIDs []uint) ([]*Navigation, error) {
	var navigation []*Navigation

	userAccess := r.db.ModelContext(ctx, (*access.UserAccess)(nil)).ColumnExpr("access_id").Where("user_id IN (?)", pg.In(userIDs))
	accessNavigation := r.db.ModelContext(ctx, (*AccessNavigation)(nil)).ColumnExpr("navigation_id").Where("access_id IN (?)", userAccess)
	err := r.db.ModelContext(ctx, &navigation).Where("id IN (?)", accessNavigation).Select()

	if len(navigation) == 0 {
		return nil, ErrItemNotFound
	}

	return navigation, err
}

func (r *Repository) GetUserToNavigation(ctx context.Context, userIDs []uint) ([]*UserNavigation, error) {
	var userToNavigation []*UserNavigation

	q := `SELECT navigation_id, users_access.user_id as user_id
       FROM access_navigation
			 LEFT JOIN users_access on users_access.access_id = access_navigation.access_id
			 WHERE access_navigation.access_id IN (
    	 	SELECT access_id
     		FROM users_access
    		WHERE user_id IN (?)
				)`

	_, err := r.db.Query(&userToNavigation, q, pg.In(userIDs))

	return userToNavigation, err
}
