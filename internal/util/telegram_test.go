package util

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

func TestNewMessageConfig(t *testing.T) {
	testCases := map[string]struct {
		inputChatID int64
		inputText   string
		result      tgbotapi.MessageConfig
	}{
		"simple case": {
			inputChatID: 1,
			inputText:   "Hello",
			result: tgbotapi.MessageConfig{
				BaseChat: tgbotapi.BaseChat{
					ChatID:           1,
					ReplyToMessageID: 0,
				},
				Text:                  "Hello",
				DisableWebPagePreview: false,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.result, NewMessageConfig(tc.inputChatID, tc.inputText))
		})
	}
}

func TestNewUpdateMessage(t *testing.T) {
	testCases := map[string]struct {
		inputChatID int64
		inputText   string
		result      tgbotapi.Update
	}{
		"simple case": {
			inputChatID: 1,
			inputText:   "Hello",
			result: tgbotapi.Update{
				Message: &tgbotapi.Message{
					Chat: &tgbotapi.Chat{
						ID: 1,
					},
					Text: "Hello",
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.result, NewUpdateMessage(tc.inputChatID, tc.inputText))
		})
	}
}

func TestNewUpdateBotMessage(t *testing.T) {
	testCases := map[string]struct {
		inputChatID int64
		inputText   string
		result      tgbotapi.Update
	}{
		"simple case": {
			inputChatID: 1,
			inputText:   "/now",
			result: tgbotapi.Update{
				Message: &tgbotapi.Message{
					Chat: &tgbotapi.Chat{
						ID: 1,
					},
					Text: "/now",
					Entities: []tgbotapi.MessageEntity{
						{
							Offset: 0,
							Length: 4,
							Type:   "bot_command",
						},
					},
				},
			},
		},
		"one arg": {
			inputChatID: 1,
			inputText:   "/now\t  123abc",
			result: tgbotapi.Update{
				Message: &tgbotapi.Message{
					Chat: &tgbotapi.Chat{
						ID: 1,
					},
					Text: "/now\t  123abc",
					Entities: []tgbotapi.MessageEntity{
						{
							Offset: 0,
							Length: 4,
							Type:   "bot_command",
						},
					},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.result, NewUpdateBotMessage(tc.inputChatID, tc.inputText))
		})
	}
}
