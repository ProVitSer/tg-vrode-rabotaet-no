package server

import (
	"encoding/json"
	"net/http"

	"github.com/ProVitSer/tg-vrode-rabotaet-no/internal/logger"
)

func SubscribeWordCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:

			subInfo, err := json.MarshalIndent(r.Response, "", "  ")
			if err != nil {
				logger.GetLogger().Printf("Ошибка при сериализации JSON: %v\n", err)
			}

			logger.GetLogger().Printf("Subscribe Info:\n%s\n", subInfo)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
