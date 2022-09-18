package echo

import (
	botClient "github.com/austinsantoso/timebot/internal/client/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type EchoBot struct{}

func NewEchoBot() *EchoBot {
	return &EchoBot{}
}

func (b *EchoBot) HandleUpdate(bot botClient.TelegramBotClient, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}
