package updater

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	handler "github.com/austinsantoso/timebot/internal/handler"
)

type Updater interface {
	Start(*tgbotapi.BotAPI, handler.MessageHandler)
}
