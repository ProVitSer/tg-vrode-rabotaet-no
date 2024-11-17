package bot

import (
	"log"

	"gopkg.in/telebot.v4"
)

type Bot struct {
	Bot            *telebot.Bot
	searchQuery    string
	searchType     string
	keyword        string
	state          string
	subscriptionId string
}

func NewBot(token string) (*Bot, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{},
	})
	if err != nil {
		return nil, err
	}

	return &Bot{
		Bot: bot,
	}, nil
}

func (b *Bot) Start() {
	log.Println("Бот запущен...")

	mainMenu, btnSearch, btnSubscribe, btnSubscriptions, btnUnsubscribe := createMainMenu()
	typeMenu, btnChannel, btnChat, btnAll := createTypeMenu()

	b.registerHandlers(mainMenu, btnSearch, btnSubscribe, btnSubscriptions, btnUnsubscribe, typeMenu, btnChannel, btnChat, btnAll)

	b.Bot.Start()
}
