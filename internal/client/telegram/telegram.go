package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//go:generate mockgen -destination=../../mocks/client/mock_telegram/mock_telegram.go -package=mock_telegram github.com/austinsantoso/timebot/internal/client/telegram TelegramBotClient
type TelegramBotClient interface {
	Send(tgbotapi.Chattable) (tgbotapi.Message, error)
}
