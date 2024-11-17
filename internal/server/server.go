package server

import (
	"log"
	"net/http"

	"github.com/ProVitSer/tg-vrode-rabotaet-no/config"
)

func StartServer() {

	http.HandleFunc("/subscribe-word", SubscribeWordCallback())

	if err := http.ListenAndServe(config.GlobalConfig.ExternServerId, nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
