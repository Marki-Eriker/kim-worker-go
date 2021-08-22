package navigation

import "time"

type Navigation struct {
	tableName struct{} `pg:"navigation"`

	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Path        string
	Title       string
	Description string
	Icon        string
	ParentId    uint
	Order       uint
	Node        bool
	Dev         bool
}

type AccessNavigation struct {
	tableName struct{} `pg:"access_navigation"`

	AccessID     uint `pg:"access_id"`
	NavigationID uint `pg:"navigation_id"`
}

type UserNavigation struct {
	UserID       uint `pg:"user_id"`
	NavigationID uint `pg:"navigation_id"`
}
