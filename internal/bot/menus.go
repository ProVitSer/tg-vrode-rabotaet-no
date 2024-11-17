package bot

import "gopkg.in/telebot.v4"

func createMainMenu() (*telebot.ReplyMarkup, telebot.Btn, telebot.Btn, telebot.Btn, telebot.Btn) {
	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	btnSearch := menu.Text("🔍 Поиск публикаций")
	btnSubscribe := menu.Text("🔔 Подписаться на слово")
	btnSubscriptions := menu.Text("📜 Мои подписки")
	btnUnsubscribe := menu.Text("❌ Отписаться")

	menu.Reply(menu.Row(btnSearch, btnSubscribe, btnSubscriptions, btnUnsubscribe))
	return menu, btnSearch, btnSubscribe, btnSubscriptions, btnUnsubscribe
}

func createTypeMenu() (*telebot.ReplyMarkup, telebot.Btn, telebot.Btn, telebot.Btn) {
	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	btnChannel := menu.Text("Каналы")
	btnChat := menu.Text("Чаты")
	btnAll := menu.Text("Все")
	menu.Reply(menu.Row(btnChannel, btnChat, btnAll))
	return menu, btnChannel, btnChat, btnAll
}
