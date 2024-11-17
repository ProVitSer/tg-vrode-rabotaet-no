package bot

import (
	"fmt"

	"github.com/ProVitSer/tg-vrode-rabotaet-no/internal/logger"
	tgstat_api "github.com/ProVitSer/tg-vrode-rabotaet-no/internal/tgstat-api"
	"gopkg.in/telebot.v4"
)

func (b *Bot) unsubscribe(c telebot.Context) error {

	logger.GetLogger().Printf("Отписка от ID: %v", b.subscriptionId)

	var err = tgstat_api.Unsubscribe(b.subscriptionId)
	if err != nil {

		message := fmt.Sprintf(
			"Произошла ошибка при попытке отписаться от ID  %s возмжно вы ввели некорректное значение",
			b.subscriptionId,
		)

		return c.Send(message, &telebot.ReplyMarkup{
			RemoveKeyboard: true,
		})

	}

	return c.Send("Вы успешно отписались", &telebot.ReplyMarkup{
		RemoveKeyboard: true,
	})

}
