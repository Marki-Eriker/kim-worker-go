package refreshtoken

import "time"

type RefreshToken struct {
	tableName struct{} `pg:"refresh_token"`

	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Token     string
	Agent     string
	ExpiresIn time.Time
	UserId    uint
}
