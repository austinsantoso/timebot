package timebot

import (
	"github.com/austinsantoso/timebot/internal/handler"
	"github.com/austinsantoso/timebot/internal/time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type EchoBot struct{}

func NewBot() *EchoBot {
	return &EchoBot{}
}

func (b *EchoBot) HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var response string
	if handler.IsUpdateBotMessage(update) {
		handleBotCommand(bot, update)
	} else {
		response = update.Message.Text
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)

		bot.Send(msg)
	}
}

var commandsToFunction = map[string]func(*tgbotapi.BotAPI, tgbotapi.Update){
	"now": handleNowCommand,
}

func handleBotCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	command := handler.ExtractBotCommand(update)

	f, ok := commandsToFunction[command]
	if !ok {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "command not found")

		bot.Send(msg)
		return
	}

	f(bot, update)
}

func handleNowCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, time.NewTimeModule().String())

	bot.Send(msg)
}
