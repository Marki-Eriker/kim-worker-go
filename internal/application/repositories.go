package application

import (
	"github.com/marki-eriker/kim-worker-go/internal/database/postgres"
	"github.com/marki-eriker/kim-worker-go/internal/feature/access"
	"github.com/marki-eriker/kim-worker-go/internal/feature/contract"
	"github.com/marki-eriker/kim-worker-go/internal/feature/file"
	"github.com/marki-eriker/kim-worker-go/internal/feature/navigation"
	"github.com/marki-eriker/kim-worker-go/internal/feature/payment"
	"github.com/marki-eriker/kim-worker-go/internal/feature/refreshtoken"
	"github.com/marki-eriker/kim-worker-go/internal/feature/request"
	"github.com/marki-eriker/kim-worker-go/internal/feature/user"
)

type Repositories struct {
	UserRepository         user.IRepository
	RefreshTokenRepository refreshtoken.IRepository
	AccessRepository       access.IRepository
	NavigationRepository   navigation.IRepository
	RequestRepository      request.IRepository
	ContractRepository     contract.IRepository
	FileRepository         file.IRepository
	PaymentRepository      payment.IRepository
}

func NewRepositories(dbs *postgres.Databases) *Repositories {
	return &Repositories{
		UserRepository:         user.NewRepository(dbs.PrimaryDB),
		RefreshTokenRepository: refreshtoken.NewRepository(dbs.PrimaryDB),
		AccessRepository:       access.NewRepository(dbs.PrimaryDB),
		NavigationRepository:   navigation.NewRepository(dbs.PrimaryDB),
		RequestRepository:      request.NewRepository(dbs.LKDB),
		ContractRepository:     contract.NewRepository(dbs.LKDB),
		FileRepository:         file.NewRepository(dbs.LKDB),
		PaymentRepository:      payment.NewRepository(dbs.LKDB),
	}
}
