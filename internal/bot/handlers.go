package bot

import (
	"gopkg.in/telebot.v4"
)

func (b *Bot) registerHandlers(mainMenu *telebot.ReplyMarkup, btnSearch, btnSubscribe, btnSubscriptions, btnUnsubscribe telebot.Btn, typeMenu *telebot.ReplyMarkup, btnChannel, btnChat, btnAll telebot.Btn) {

	b.Bot.Handle("/start", func(c telebot.Context) error {
		chatID := c.Chat().ID
		err := b.saveChatID(chatID)
		if err != nil {
			return c.Send("Ошибка сохранения вашего чата. Попробуйте позже.")
		}
		return c.Send("Добро пожаловать! Выберите действие:", mainMenu)
	})

	b.Bot.Handle(&btnSearch, func(c telebot.Context) error {
		b.state = "search"
		return c.Send("Введите текст для поиска:")
	})

	b.Bot.Handle(&btnChannel, func(c telebot.Context) error {
		b.searchType = "channel"
		return b.subscribeWord(c)
	})

	b.Bot.Handle(&btnChat, func(c telebot.Context) error {
		b.searchType = "chat"
		return b.subscribeWord(c)
	})

	b.Bot.Handle(&btnAll, func(c telebot.Context) error {
		b.searchType = "all"
		return b.subscribeWord(c)
	})

	b.Bot.Handle(&btnSubscribe, func(c telebot.Context) error {
		b.state = "subscribe"
		return c.Send("Введите ключевое слово для подписки:")
	})

	b.Bot.Handle(&btnSubscriptions, func(c telebot.Context) error {
		return b.subscribeList(c)
	})

	b.Bot.Handle(&btnUnsubscribe, func(c telebot.Context) error {
		b.state = "unsubscribe"
		return c.Send("Введите id подписки от которой хотите отписаться")
	})

	b.Bot.Handle(telebot.OnText, func(c telebot.Context) error {
		switch b.state {
		case "subscribe":
			b.keyword = c.Text()
			b.state = ""
			return c.Send("Выберите тип источника:", typeMenu)
		case "search":
			b.searchQuery = c.Text()
			b.state = ""
			c.Send("Ищем")
			return b.performSearch(c)

		case "unsubscribe":
			b.subscriptionId = c.Text()
			return b.unsubscribe(c)
		default:
			return c.Send("Неизвестная команда. Пожалуйста, выберите действие из меню.")
		}
	})

}
