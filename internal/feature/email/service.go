package email

import (
	"crypto/tls"
	"github.com/marki-eriker/kim-worker-go/internal/application"
	mail "github.com/xhit/go-simple-mail/v2"
	"strconv"
)

type IService interface {
	SendMail(address, message string) error
}

type Service struct {
	Mailer      *mail.SMTPServer
	FromAddress string
	FromMessage string
}

func NewService(args *application.EmailCredentials) *Service {
	mailer := mail.NewSMTPClient()

	mailer.Host = args.SMTPHost
	mailer.Port, _ = strconv.Atoi(args.SMTPPort)
	mailer.Username = args.SMTPUser
	mailer.Password = args.SMTPPassword
	mailer.Encryption = mail.EncryptionSTARTTLS
	mailer.KeepAlive = false
	mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return &Service{Mailer: mailer, FromAddress: args.SMTPFrom, FromMessage: args.SMTPFromMessage}
}

func (s *Service) SendMail(address, message string) error {
	smtpClient, err := s.Mailer.Connect()
	if err != nil {
		return err
	}
	defer smtpClient.Close()

	email := mail.NewMSG()
	email.SetFrom(s.FromAddress).AddTo(address).SetSubject(s.FromMessage).SetBody(mail.TextPlain, message)

	if email.Error != nil {
		return email.Error
	}

	err = email.Send(smtpClient)
	if err != nil {
		return err
	}

	return nil
}
