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
		handler.SendMessage(bot, update, response)
	}
}

var commandsToFunction = map[string]func(*tgbotapi.BotAPI, tgbotapi.Update){
	"now":  handleNowCommand,
	"help": handleHelpCommand,
}

func handleBotCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	command := handler.ExtractBotCommand(update)

	f, ok := commandsToFunction[command]
	if !ok {
		handler.SendMessage(bot, update, "command not found")
		return
	}

	f(bot, update)
}

func handleHelpCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := `
	Hello I am timebot 
	<required field type>

	/help - get help message\n

	/now  - Get the current time
	/millisecondsago <number> - Get the time number of days ago
	/secondsago <number> - Get the time number of days ago
	/minutessago <number> - Get the time number of days ago
	/hourssago <number> - Get the time number of days ago
	/daysago <number> - Get the time number of days ago
	/weeksago <number> - Get the time number of week ago
	`

	handler.SendMessage(bot, update, text)
}

func handleNowCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	handler.SendMessage(bot, update, time.Now().String())
}
