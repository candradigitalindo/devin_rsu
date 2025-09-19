package config

import (
	"os"
)

type Config struct {
	DatabaseURL      string
	AppDomain        string
	LogLevel         string
	MidtransBaseURL  string
	MidtransServerKey string
	MidtransClientKey string
	AuthIssuer       string
	AuthAudience     string
}

func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}

func Load() Config {
	return Config{
		DatabaseURL:       getenv("DATABASE_URL", ""),
		AppDomain:         getenv("APP_DOMAIN", "localhost"),
		LogLevel:          getenv("LOG_LEVEL", "info"),
		MidtransBaseURL:   getenv("MIDTRANS_BASE_URL", "https://app.sandbox.midtrans.com"),
		MidtransServerKey: getenv("MIDTRANS_SERVER_KEY", ""),
		MidtransClientKey: getenv("MIDTRANS_CLIENT_KEY", ""),
		AuthIssuer:        getenv("AUTH_ISSUER", ""),
		AuthAudience:      getenv("AUTH_AUDIENCE", ""),
	}
}
