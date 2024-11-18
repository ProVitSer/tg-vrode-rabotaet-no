package bot

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"gopkg.in/telebot.v4"
)

type Bot struct {
	Bot            *telebot.Bot
	searchQuery    string
	searchType     string
	keyword        string
	state          string
	subscriptionId string

	FilePath string
	mu       sync.Mutex
}

func NewBot(token string, filePath string) (*Bot, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{},
	})
	if err != nil {
		return nil, err
	}

	return &Bot{
		Bot:      bot,
		FilePath: filePath,
	}, nil
}

func (b *Bot) Start() {
	log.Println("Бот запущен...")

	mainMenu, btnSearch, btnSubscribe, btnSubscriptions, btnUnsubscribe := createMainMenu()
	typeMenu, btnChannel, btnChat, btnAll := createTypeMenu()

	b.registerHandlers(mainMenu, btnSearch, btnSubscribe, btnSubscriptions, btnUnsubscribe, typeMenu, btnChannel, btnChat, btnAll)
	b.Bot.Start()
}

func (b *Bot) saveChatID(chatID int64) error {

	existingChatIDs, err := b.loadChatIDs()
	if err != nil {
		return fmt.Errorf("ошибка загрузки chat_id: %v", err)
	}

	for _, id := range existingChatIDs {
		if id == chatID {
			return nil
		}
	}

	file, err := os.OpenFile(b.FilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%d\n", chatID))
	return err
}

func (b *Bot) loadChatIDs() ([]int64, error) {

	file, err := os.Open(b.FilePath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var chatIDs []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		id, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err == nil {
			chatIDs = append(chatIDs, id)
		}
	}
	return chatIDs, scanner.Err()
}

func (b *Bot) BroadcastMessage(message string) error {
	chatIDs, err := b.loadChatIDs()
	if err != nil {
		return fmt.Errorf("ошибка загрузки chat_id: %v", err)
	}

	for _, chatID := range chatIDs {
		_, err := b.Bot.Send(&telebot.Chat{ID: chatID}, message)
		if err != nil {
			log.Printf("Ошибка отправки сообщения в чат %d: %v\n", chatID, err)
		}
	}

	return nil
}
