package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppPort          string
	DBDSN            string
	JWTSecret        string
	MPAccessToken    string
	MPNotificationURL string
}

func Load() *Config {
	cfg := &Config{
		AppPort:           getenv("APP_PORT", "8081"),
		DBDSN:             getenv("DB_DSN", "root:root@tcp(127.0.0.1:3306)/shopdb?parseTime=true&charset=utf8mb4"),
		JWTSecret:         getenv("JWT_SECRET", "supersecret"),
		MPAccessToken:     getenv("MP_ACCESS_TOKEN", ""),
		MPNotificationURL: getenv("MP_NOTIFICATION_URL", "http://localhost:8081"),
	}
	return cfg
}

func (c *Config) Port() string { return fmt.Sprintf(":%s", c.AppPort) }

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
