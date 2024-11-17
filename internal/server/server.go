package server

import (
	"log"
	"net/http"
)

func StartServer(addr string) {

	//http.HandleFunc("/callback", SetleCallback)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
