package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var GlobalConfig *Config

type Config struct {
	TelegramBotToken         string
	TGStatToken              string
	ExternServerId           string
	SubscribeWordCallbackUrl string
}

func LoadConfig() error {
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
		log.Fatalf("TGStat token is not set in config")
	}

	var ExternServerId = os.Getenv("EXTERNAL_IP")

	if ExternServerId == "" {
		log.Fatalf("Extern IP is not set in config")
	}

	var SubscribeWordCallbackUrl = os.Getenv("SUBSCRIBE_WORD_CALLBACK_URL")

	if SubscribeWordCallbackUrl == "" {
		log.Fatalf("SubscribeWordCallbackUrl is not set in config")
	}

	GlobalConfig = &Config{
		TelegramBotToken:         tgToken,
		TGStatToken:              tgStatToken,
		ExternServerId:           ExternServerId,
		SubscribeWordCallbackUrl: SubscribeWordCallbackUrl,
	}

	return nil
}
