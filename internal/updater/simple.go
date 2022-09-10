package updater

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	handler "github.com/austinsantoso/timebot/internal/handler"
)

type SimpleUpdater struct {
}

func NewSimpleUpdater() *SimpleUpdater {
	return &SimpleUpdater{}
}

func (b *SimpleUpdater) Start(bot *tgbotapi.BotAPI, h handler.MessageHandler) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if handler.IsUpdateMessage(update) { // If we got a message
			log.Printf("[Update][%s] %+v", update.Message.From.UserName, update)
			log.Printf("[Message][%s] %+v", update.Message.From.UserName, update.Message)
			log.Printf("[Text][%s] %s", update.Message.From.UserName, update.Message.Text)

			h.HandleUpdate(bot, update)
		}
	}
}
