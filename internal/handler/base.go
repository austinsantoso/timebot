package handler

import (
	"strings"
	"unicode"

	botClient "github.com/austinsantoso/timebot/internal/client/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageHandler interface {
	HandleUpdate(bot botClient.TelegramBotClient, update tgbotapi.Update)
}

var botCommandEntityType string = "bot_command"

func SendMessage(bot botClient.TelegramBotClient, update tgbotapi.Update, text string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	bot.Send(msg)
}

func IsUpdateTextMessage(update tgbotapi.Update) bool {
	return update.Message != nil && len(update.Message.Text) > 0
}

func IsUpdateBotMessage(update tgbotapi.Update) bool {
	if !IsUpdateTextMessage(update) ||
		len(update.Message.Entities) == 0 {
		return false
	}

	if update.Message.Entities[0].Type == botCommandEntityType {
		return true
	}

	return false
}

func ExtractBotCommand(update tgbotapi.Update) string {
	offset := update.Message.Entities[0].Offset
	length := update.Message.Entities[0].Length
	command := update.Message.Text[offset+1 : offset+length]
	return command
}

func isWhiteSpace(r rune) bool {
	return unicode.IsSpace(r)
}

func ExtractBotCommandArguments(update tgbotapi.Update) []string {
	offset := update.Message.Entities[0].Offset
	length := update.Message.Entities[0].Length
	// command := update.Message.Text[offset+1 : offset+length]
	args := strings.TrimSpace(update.Message.Text[offset+length:])

	argsSplit := strings.FieldsFunc(args, isWhiteSpace)

	filtered := []string{}

	// remove length 0 many spaces
	for i := range argsSplit {
		if len(argsSplit[i]) > 0 {
			filtered = append(filtered, argsSplit[i])
		}
	}

	return filtered
}
