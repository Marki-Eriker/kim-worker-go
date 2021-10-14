package application

import (
	"github.com/marki-eriker/kim-worker-go/internal/feature/contract"
	"github.com/marki-eriker/kim-worker-go/internal/feature/email"
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
	EmailService        email.IService
}

func NewServices(repos *Repositories, mail *EmailCredentials) *Services {
	return &Services{
		UserService:         user.NewService(repos.UserRepository),
		RefreshTokenService: refreshtoken.NewService(repos.RefreshTokenRepository),
		RequestService:      request.NewService(repos.RequestRepository),
		ContractService:     contract.NewService(repos.ContractRepository),
		FileService:         file.NewService(repos.FileRepository),
		PaymentService:      payment.NewService(repos.PaymentRepository),
		EmailService: email.NewService(
			mail.SMTPHost,
			mail.SMTPPort,
			mail.SMTPUser,
			mail.SMTPPassword,
			mail.SMTPFrom,
			mail.SMTPFromMessage,
		),
	}
}
