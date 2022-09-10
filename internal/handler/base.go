package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageHandler interface {
	HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update)
}

var botCommandEntityType string = "bot_command"

func IsUpdateMessage(update tgbotapi.Update) bool {
	return update.Message != nil
}

func IsUpdateBotMessage(update tgbotapi.Update) bool {
	if !IsUpdateMessage(update) ||
		len(update.Message.Entities) == 0 {
		return false
	}

	if update.Message.Entities[0].Type == botCommandEntityType {
		return true
	}

	return false
}
