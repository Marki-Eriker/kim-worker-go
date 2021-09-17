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
	SMTPHost        string
	SMTPPort        string
	SMTPUser        string
	SMTPPassword    string
	SMTPFrom        string
	SMTPFromMessage string
}

func NewService(SMTPHost string, SMTPPort string, SMTPUser string, SMTPPassword string, SMTPFrom string, SMTPFromMessage string) *Service {
	return &Service{SMTPHost: SMTPHost, SMTPPort: SMTPPort, SMTPUser: SMTPUser, SMTPPassword: SMTPPassword, SMTPFrom: SMTPFrom, SMTPFromMessage: SMTPFromMessage}
}

func (s *Service) SendMail(address, message string) error {
	server := mail.NewSMTPClient()

	server.Host = s.SMTPHost
	server.Port, _ = strconv.Atoi(s.SMTPPort)
	server.Username = s.SMTPUser
	server.Password = s.SMTPPassword
	server.Encryption = mail.EncryptionSTARTTLS

	server.KeepAlive = false
	server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}
	defer smtpClient.Close()

	email := mail.NewMSG()
	email.SetFrom(s.SMTPFrom).AddTo(address).SetSubject(s.SMTPFromMessage).SetBody(mail.TextPlain, message)

	if email.Error != nil {
		return email.Error
	}

	err = email.Send(smtpClient)
	if err != nil {
		return err
	}

	return nil
}
