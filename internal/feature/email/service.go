package email

import (
	"crypto/tls"
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

func NewService(host, port, user, password, from, fromMessage string) *Service {
	mailer := mail.NewSMTPClient()

	mailer.Host = host
	mailer.Port, _ = strconv.Atoi(port)
	mailer.Username = user
	mailer.Password = password
	mailer.Encryption = mail.EncryptionSTARTTLS
	mailer.KeepAlive = false
	mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return &Service{Mailer: mailer, FromAddress: from, FromMessage: fromMessage}
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
