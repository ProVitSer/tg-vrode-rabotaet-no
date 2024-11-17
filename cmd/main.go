package main

import (
	"log"

	"github.com/ProVitSer/tg-vrode-rabotaet-no/config"
	"github.com/ProVitSer/tg-vrode-rabotaet-no/internal/bot"
	"github.com/ProVitSer/tg-vrode-rabotaet-no/internal/logger"
	"github.com/ProVitSer/tg-vrode-rabotaet-no/internal/server"
	tgstat_api "github.com/ProVitSer/tg-vrode-rabotaet-no/internal/tgstat-api"
	tgstat "github.com/helios-ag/tgstat-go"
)

func main() {

	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	err := logger.InitLogger("output.log")
	if err != nil {
		log.Fatalf("Ошибка инициализации логгера: %v\n", err)
	}

	tgstat.Token = config.GlobalConfig.TGStatToken

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic in server.StartServer: %v", r)
			}
		}()
		server.StartServer()

	}()

	b, err := bot.NewBot(config.GlobalConfig.TelegramBotToken)
	if err != nil {
		log.Fatalf("Ошибка инициализации бота: %v", err)
	}

	var _, e = tgstat_api.GetCallbackInfo()

	if e != nil {

		tgstat_api.SetCallbackSubscribeWord()
	}

	b.Start()

	select {}
}
