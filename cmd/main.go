package main

import (
	"github.com/ProVitSer/tg-vrode-rabotaet-no/config"
	"github.com/ProVitSer/tg-vrode-rabotaet-no/internal/server"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgstat "github.com/helios-ag/tgstat-go"
	"log"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	tgstat.Token = config.TGStatToken

	bot, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic in server.StartServer: %v", r)
			}
		}()
		server.StartServer(config.ExternServerId)

	}()

	select {}
}
