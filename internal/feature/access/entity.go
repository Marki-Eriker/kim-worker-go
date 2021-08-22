package access

import "time"

type Access struct {
	tableName struct{} `pg:"access"`

	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

type UserAccess struct {
	tableName struct{} `pg:"users_access"`

	UserID   uint `pg:"user_id"`
	AccessID uint `pg:"access_id"`
}
