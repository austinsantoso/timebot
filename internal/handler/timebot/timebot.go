package timebot

import (
	"log"
	"strconv"
	gotime "time"

	botClient "github.com/austinsantoso/timebot/internal/client/telegram"
	"github.com/austinsantoso/timebot/internal/handler"
	"github.com/austinsantoso/timebot/internal/time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type currentTimeProvider func() *time.TimeModule

type EchoBot struct {
	timeProvider currentTimeProvider
}

func getCurrentTime() *time.TimeModule {
	return time.Now()
}

func NewBot() *EchoBot {
	return &EchoBot{
		timeProvider: getCurrentTime,
	}
}

func NewCustomTimeBot(timeProvider func() *time.TimeModule) *EchoBot {
	return &EchoBot{
		timeProvider: timeProvider,
	}
}

func (b *EchoBot) HandleUpdate(bot botClient.TelegramBotClient, update tgbotapi.Update) {
	var response string
	if handler.IsUpdateBotMessage(update) {
		handleBotCommand(bot, update, b.timeProvider)
	} else {
		response = update.Message.Text
		handler.SendMessage(bot, update, response)
	}
}

var commandsToFunction = map[string]func(botClient.TelegramBotClient, tgbotapi.Update, currentTimeProvider){
	"now":  handleNowCommand,
	"help": handleHelpCommand,
	"msb":  handleMillisecondsBefore,
	"sb":   handleSecondsBefore,
	"mb":   handleMinutesBefore,
	"hb":   handleHoursBefore,
	"db":   handleDaysBefore,
	"wb":   handleWeeksBefore,
}

func handleBotCommand(bot botClient.TelegramBotClient, update tgbotapi.Update, t currentTimeProvider) {
	command := handler.ExtractBotCommand(update)

	f, ok := commandsToFunction[command]
	if !ok {
		handler.SendMessage(bot, update, "command not found")
		return
	}

	f(bot, update, t)
}

func handleHelpCommand(bot botClient.TelegramBotClient, update tgbotapi.Update, t currentTimeProvider) {
	text := `
	Hello I am timebot 
	<required field type>

	/help - get help message\n

	/now  - Get the current time
	/msb <number> - Get the time number of days ago 
	/sb <number> - Get the time number of days ago
	/mb <number> - Get the time number of days ago
	/hb <number> - Get the time number of days ago
	/db <number> - Get the time number of days ago
	/wb <number> - Get the time number of week ago
	`

	handler.SendMessage(bot, update, text)
}

func handleNowCommand(bot botClient.TelegramBotClient, update tgbotapi.Update, t currentTimeProvider) {
	handler.SendMessage(bot, update, t().String())
}

func handleGenericTimeAgo(bot botClient.TelegramBotClient, update tgbotapi.Update, t currentTimeProvider, errorMessage string, dur gotime.Duration) {
	args := handler.ExtractBotCommandArguments(update)
	if len(args) == 0 {
		handler.SendMessage(bot, update, errorMessage)
		return
	}
	d, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		handler.SendMessage(bot, update, errorMessage)
		log.Println(err)
		return
	}
	handler.SendMessage(bot, update, t().Add(dur*gotime.Duration(-1*d)).String())
}

func handleMillisecondsBefore(bot botClient.TelegramBotClient, update tgbotapi.Update, t currentTimeProvider) {
	message := "invalid syntax, please use /msb <number>"
	handleGenericTimeAgo(bot, update, t, message, gotime.Millisecond)
}

func handleSecondsBefore(bot botClient.TelegramBotClient, update tgbotapi.Update, t currentTimeProvider) {
	message := "invalid syntax, please use /sb <number>"
	handleGenericTimeAgo(bot, update, t, message, gotime.Second)
}

func handleMinutesBefore(bot botClient.TelegramBotClient, update tgbotapi.Update, t currentTimeProvider) {
	message := "invalid syntax, please use /mb <number>"
	handleGenericTimeAgo(bot, update, t, message, gotime.Minute)
}

func handleHoursBefore(bot botClient.TelegramBotClient, update tgbotapi.Update, t currentTimeProvider) {
	message := "invalid syntax, please use /hb <number>"
	handleGenericTimeAgo(bot, update, t, message, gotime.Hour)
}

func handleDaysBefore(bot botClient.TelegramBotClient, update tgbotapi.Update, t currentTimeProvider) {
	message := "invalid syntax, please use /db <number>"
	handleGenericTimeAgo(bot, update, t, message, gotime.Hour*gotime.Duration(24))
}

func handleWeeksBefore(bot botClient.TelegramBotClient, update tgbotapi.Update, t currentTimeProvider) {
	message := "invalid syntax, please use /wb <number>"
	handleGenericTimeAgo(bot, update, t, message, gotime.Hour*gotime.Duration(24*7))
}
