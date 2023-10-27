package config

import (
	"errors"
	"os"
)

type Config struct {
	TelegramBotToken string
}

func GetConfig() (*Config, error) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		return nil, errors.New("TELEGRAM_BOT_TOKEN is not set")
	}

	return &Config{
		TelegramBotToken: botToken,
	}, nil

}
