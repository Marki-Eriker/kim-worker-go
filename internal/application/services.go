package application

import (
	"github.com/marki-eriker/kim-worker-go/internal/feature/contract"
	"github.com/marki-eriker/kim-worker-go/internal/feature/file"
	"github.com/marki-eriker/kim-worker-go/internal/feature/payment"
	"github.com/marki-eriker/kim-worker-go/internal/feature/refreshtoken"
	"github.com/marki-eriker/kim-worker-go/internal/feature/request"
	"github.com/marki-eriker/kim-worker-go/internal/feature/user"
)

type Services struct {
	UserService         user.IService
	RefreshTokenService refreshtoken.IService
	RequestService      request.IService
	ContractService     contract.IService
	FileService         file.IService
	PaymentService      payment.IService
}

func NewServices(repos *Repositories) *Services {
	return &Services{
		UserService:         user.NewService(repos.UserRepository),
		RefreshTokenService: refreshtoken.NewService(repos.RefreshTokenRepository),
		RequestService:      request.NewService(repos.RequestRepository),
		ContractService:     contract.NewService(repos.ContractRepository),
		FileService:         file.NewService(repos.FileRepository),
		PaymentService:      payment.NewService(repos.PaymentRepository),
	}
}