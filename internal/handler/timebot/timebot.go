package timebot

import (
	"github.com/austinsantoso/timebot/internal/handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type EchoBot struct{}

func NewBot() *EchoBot {
	return &EchoBot{}
}

func (b *EchoBot) HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var response string
	if handler.IsUpdateBotMessage(update) {
		response = "bot command"
	} else {
		response = update.Message.Text
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)

	bot.Send(msg)
}
