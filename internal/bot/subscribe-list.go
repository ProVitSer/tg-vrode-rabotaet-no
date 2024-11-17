package bot

import (
	"fmt"

	tgstat_api "github.com/ProVitSer/tg-vrode-rabotaet-no/internal/tgstat-api"
	"gopkg.in/telebot.v4"
)

func (b *Bot) subscribeList(c telebot.Context) error {
	var subs, err = tgstat_api.SubscriptionsList()
	if err != nil {
		return c.Send("Произошла ошибка при запросе списка подписок", &telebot.ReplyMarkup{
			RemoveKeyboard: true,
		})

	}
	if len(subs) == 0 {
		return c.Send("У вас нет активных подписок", &telebot.ReplyMarkup{
			RemoveKeyboard: true,
		})
	}

	var message string
	for _, sub := range subs {
		message += fmt.Sprintf(
			"ID: %d\nТип: %s\nКлючевое слово: %s\n\n",
			sub.SubscriptionId, sub.Type, sub.Keyword.Q,
		)
	}

	return c.Send(message, &telebot.ReplyMarkup{
		RemoveKeyboard: true,
	})
}
