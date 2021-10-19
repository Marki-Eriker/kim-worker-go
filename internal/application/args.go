package application

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
	"os"
)

type Args struct {
	ListenPort         string
	PrimaryDatabaseURL string
	LKDatabaseURL      string
	LKSchema           string
	LogLevel           string
	EmailCredentials   *EmailCredentials
}

type EmailCredentials struct {
	SMTPHost        string
	SMTPPort        string
	SMTPUser        string
	SMTPPassword    string
	SMTPFrom        string
	SMTPFromMessage string
}

func NewArgs() (*Args, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load environment: %v", err)
	}

	listenPort := os.Getenv("PORT")
	if listenPort == "" {
		return nil, fmt.Errorf("port must be set")
	}

	primaryDatabaseURL := os.Getenv("DB_URL")
	if !checkPostgresURL(primaryDatabaseURL) {
		return nil, fmt.Errorf("invalid database URL: %v", primaryDatabaseURL)
	}

	lkDatabaseURL := os.Getenv("LK_DB_URL")
	if !checkPostgresURL(lkDatabaseURL) {
		return nil, fmt.Errorf("invalid database URL: %v", lkDatabaseURL)
	}

	lkSchema := os.Getenv("LK_DB_SCHEMA")
	if lkSchema == "" {
		return nil, fmt.Errorf("LK_DB_SCHEMA must be set, lk or lk_dev")
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if !checkLogLevel(logLevel) {
		return nil, fmt.Errorf("invalid logLevel: %s: alloved values: %v", logLevel, getAllowedLogLevels())
	}

	SMTPHost := os.Getenv("SMTP_HOST")
	if SMTPHost == "" {
		return nil, fmt.Errorf("SMTP host must be set")
	}

	SMTPPort := os.Getenv("SMTP_PORT")
	if SMTPPort == "" {
		return nil, fmt.Errorf("SMTP port must be set")
	}

	SMTPUser := os.Getenv("SMTP_USER")
	if SMTPUser == "" {
		return nil, fmt.Errorf("SMTP user must be set")
	}
	SMTPPassword := os.Getenv("SMTP_PASSWORD")
	if SMTPPassword == "" {
		return nil, fmt.Errorf("SMTP password must be set")
	}
	SMTPFrom := os.Getenv("SMTP_FROM_EMAIL")
	if SMTPFrom == "" {
		return nil, fmt.Errorf("SMTP form email must be set")
	}
	SMTPFromMessage := os.Getenv("SMTP_FROM_MESSAGE")
	if SMTPFromMessage == "" {
		return nil, fmt.Errorf("SMTP from message must be set")
	}

	return &Args{
		ListenPort:         listenPort,
		PrimaryDatabaseURL: primaryDatabaseURL,
		LKDatabaseURL:      lkDatabaseURL,
		LKSchema:           lkSchema,
		LogLevel:           logLevel,
		EmailCredentials: &EmailCredentials{
			SMTPHost:        SMTPHost,
			SMTPPort:        SMTPPort,
			SMTPUser:        SMTPUser,
			SMTPPassword:    SMTPPassword,
			SMTPFrom:        SMTPFrom,
			SMTPFromMessage: SMTPFromMessage,
		},
	}, nil
}

func checkPostgresURL(url string) bool {
	_, err := pg.ParseURL(url)

	return err == nil
}

func checkLogLevel(level string) bool {
	for _, l := range getAllowedLogLevels() {
		if level == l {
			return true
		}
	}

	return false
}

func getAllowedLogLevels() []string {
	return []string{
		zapcore.DebugLevel.String(),
		zapcore.InfoLevel.String(),
		zapcore.WarnLevel.String(),
		zapcore.ErrorLevel.String(),
	}
}
