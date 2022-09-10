package echo

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type EchoBot struct{}

func NewEchoBot() *EchoBot {
	return &EchoBot{}
}

func (b *EchoBot) HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}
