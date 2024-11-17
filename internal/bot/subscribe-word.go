package bot

import (
	"fmt"

	"github.com/ProVitSer/tg-vrode-rabotaet-no/internal/logger"
	tgstat_api "github.com/ProVitSer/tg-vrode-rabotaet-no/internal/tgstat-api"
	"gopkg.in/telebot.v4"
)

func (b *Bot) subscribeWord(c telebot.Context) error {

	logger.GetLogger().Printf("Подписка на слово: %v", b.keyword)
	logger.GetLogger().Printf("Подписка на тип: %v", b.searchType)

	var subId, err = tgstat_api.SubscribeWord(b.keyword, b.searchType)
	if err != nil {
		return c.Send("Произошла ошибка при подписке, попробуйте позже", &telebot.ReplyMarkup{
			RemoveKeyboard: true,
		})

	}

	b.searchType = ""
	b.keyword = ""

	message := fmt.Sprintf(
		"Подписка оформлена. ID=%d",
		*subId,
	)
	return c.Send(message, &telebot.ReplyMarkup{
		RemoveKeyboard: true,
	})

}
