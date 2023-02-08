package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server   server
	Mail     mail
	Telegram telegram
}

type telegram struct {
	BotKey string
}

type mail struct {
	Smtp     string
	Port     int
	Mail     string
	Password string
}

type server struct {
	Url string
}

func GetConfig() *Config {
	return &Config{
		Server: server{
			Url: getEnv("SERVER", ":3000"),
		},
		Mail: mail{
			Smtp:     getEnv("SMTP_SERVER", ""),
			Port:     getEnvInt("SMTP_PORT", 465),
			Mail:     getEnv("SMTP_EMAIL", ""),
			Password: getEnv("SMTP_PASSWORD", ""),
		},
		Telegram: telegram{
			BotKey: getEnv("TELEGRAM_BOT", ""),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultValue
}
