package util

import (
	"strings"
	"unicode"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewMessageConfig(chatID int64, text string) tgbotapi.MessageConfig {
	return tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           chatID,
			ReplyToMessageID: 0,
		},
		Text:                  text,
		DisableWebPagePreview: false,
	}
}

func NewUpdateMessage(chatID int64, text string) tgbotapi.Update {
	return tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				ID: chatID,
			},
			Text: text,
		},
	}
}

func isWhiteSpace(r rune) bool {
	return unicode.IsSpace(r)
}

func NewUpdateBotMessage(chatID int64, text string) tgbotapi.Update {
	// split tokens from text
	// a

	args := strings.TrimSpace(text)

	argsSplit := strings.FieldsFunc(args, isWhiteSpace)

	tokens := []string{}

	// remove length 0 many spaces
	for i := range argsSplit {
		if len(argsSplit[i]) > 0 {
			tokens = append(tokens, argsSplit[i])
		}
	}

	// get length of first word
	commandLength := len(tokens[0])

	return tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				ID: chatID,
			},
			Text: text,
			Entities: []tgbotapi.MessageEntity{
				{
					Offset: 0,
					Length: commandLength,
					Type:   "bot_command",
				},
			},
		},
	}
}
