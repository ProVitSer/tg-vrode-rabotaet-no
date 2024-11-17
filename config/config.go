package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	TelegramBotToken string
	TGStatToken      string
	ExternServerId   string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error loading env file on config package")
	}

	var tgToken = os.Getenv("TELEGRAM_BOT_TOKEN")

	if tgToken == "" {
		log.Fatalf("Telegram bot token is not set in config")
	}

	var tgStatToken = os.Getenv("TGSTAT_TOKEN")

	if tgStatToken == "" {
		log.Fatalf("Telegram bot token is not set in config")
	}

	var ExternServerId = os.Getenv("EXTERNAL_IP")

	if ExternServerId == "" {
		log.Fatalf("Telegram bot token is not set in config")
	}

	return &Config{
		TelegramBotToken: tgToken,
		TGStatToken:      tgStatToken,
		ExternServerId:   ExternServerId,
	}, nil
}
