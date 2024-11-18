package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ProVitSer/tg-vrode-rabotaet-no/internal/bot"
	"github.com/ProVitSer/tg-vrode-rabotaet-no/internal/logger"
)

type Media struct {
	MediaType string `json:"media_type"`
	Caption   string `json:"caption"`
}

type Post struct {
	ID            int64  `json:"id"`
	Date          int64  `json:"date"`
	Views         int    `json:"views"`
	Link          string `json:"link"`
	ChannelID     int    `json:"channel_id"`
	ForwardedFrom int    `json:"forwarded_from"`
	IsDeleted     int    `json:"is_deleted"`
	Text          string `json:"text"`
	Media         Media  `json:"media"`
}

type Channel struct {
	ID                int    `json:"id"`
	Link              string `json:"link"`
	Username          string `json:"username"`
	Title             string `json:"title"`
	About             string `json:"about"`
	Image100          string `json:"image100"`
	Image640          string `json:"image640"`
	ParticipantsCount int    `json:"participants_count"`
}

type CallbackData struct {
	SubscriptionID   int       `json:"subscription_id"`
	SubscriptionType string    `json:"subscription_type"`
	EventID          int       `json:"event_id"`
	EventType        string    `json:"event_type"`
	Post             Post      `json:"post"`
	Channels         []Channel `json:"channels"`
}

func SubscribeWordCallback(b *bot.Bot) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:

			var bodyBuffer bytes.Buffer
			tee := io.TeeReader(r.Body, &bodyBuffer)
			rawBody, err := io.ReadAll(tee)

			logger.GetLogger().Printf("Получено тело запроса:\n%s\n", string(rawBody))

			r.Body = io.NopCloser(&bodyBuffer)

			var callbackData CallbackData
			err = json.NewDecoder(r.Body).Decode(&callbackData)
			if err != nil {
				logger.GetLogger().Printf("Ошибка при парсинге JSON: %v\n", err)
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			postDate := time.Unix(callbackData.Post.Date, 0).Format("2006-01-02 15:04:05")

			var channelsInfo string
			for _, channel := range callbackData.Channels {
				channelsInfo += fmt.Sprintf("Link: %s, Username: %s\n", channel.Link, channel.Username)
			}

			message := fmt.Sprintf(
				"Подписка: %d\nДата: %s\nСсылка: %s\nТекст:\n%s\nКаналы:\n%s",
				callbackData.SubscriptionID, postDate, callbackData.Post.Link, callbackData.Post.Text, channelsInfo,
			)

			err = b.BroadcastMessage(message)
			if err != nil {
				logger.GetLogger().Printf("Ошибка отправки сообщения: %v\n", err)
			}

			w.WriteHeader(http.StatusOK)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
