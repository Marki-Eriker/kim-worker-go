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
	LogLevel           string
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

	logLevel := os.Getenv("LOG_LEVEL")
	if !checkLogLevel(logLevel) {
		return nil, fmt.Errorf("invalid logLevel: %s: alloved values: %v", logLevel, getAllowedLogLevels())
	}

	return &Args{
		ListenPort:         listenPort,
		PrimaryDatabaseURL: primaryDatabaseURL,
		LKDatabaseURL:      lkDatabaseURL,
		LogLevel:           logLevel,
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
