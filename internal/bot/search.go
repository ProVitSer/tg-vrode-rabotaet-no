package bot

import (
	"fmt"
	"time"

	"github.com/ProVitSer/tg-vrode-rabotaet-no/internal/logger"
	tgstat_api "github.com/ProVitSer/tg-vrode-rabotaet-no/internal/tgstat-api"
	"gopkg.in/telebot.v4"
)

const textSnippetLimit = 500

func (b *Bot) performSearch(c telebot.Context) error {

	posts, err := tgstat_api.PostSearch(b.searchQuery)
	if err != nil {
		logger.GetLogger().Printf("Ошибка поиска: %v", err)
		return c.Send("Произошла ошибка при поиске. Попробуйте позже.")
	}

	if len(posts) == 0 {
		return c.Send("По вашему запросу ничего не найдено.")
	}

	for _, postInfo := range posts {
		logger.GetLogger().Printf("Обрабатываем пост: ID=%d, Views=%d, Link=%s", postInfo.ID, postInfo.Views, postInfo.Link)

		date := time.Unix(int64(postInfo.Date), 0).Format("2006-01-02 15:04:05")

		// Обрезаем текст
		snippet := trimText(postInfo.Text, textSnippetLimit)

		// Формируем строку результата
		message := fmt.Sprintf(
			"Дата: %s\nСсылка: %s\nПросмотры: %d\nТекст: %s",
			date, postInfo.Link, postInfo.Views, snippet,
		)

		// Отправляем каждое сообщение отдельно
		if err := c.Send(message); err != nil {
			logger.GetLogger().Printf("Ошибка отправки сообщения: %v", err)
		}
	}

	b.searchQuery = ""
	return nil
}

func trimText(text string, limit int) string {
	if len(text) > limit {
		return text[:limit] + "..."
	}
	return text
}
