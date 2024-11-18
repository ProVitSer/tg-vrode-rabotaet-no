package server

import (
	"log"
	"net/http"

	"github.com/ProVitSer/tg-vrode-rabotaet-no/config"
	"github.com/ProVitSer/tg-vrode-rabotaet-no/internal/bot"
)

func StartServer(b *bot.Bot) {

	http.HandleFunc("/subscribe-word", SubscribeWordCallback(b))

	if err := http.ListenAndServe(config.GlobalConfig.ExternServerId, nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
