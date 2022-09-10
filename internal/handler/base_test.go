package handler

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

func TestIsUpdateMessage(t *testing.T) {
	testCases := map[string]struct {
		input tgbotapi.Update
		want  bool
	}{
		"TrueCase": {input: tgbotapi.Update{Message: &tgbotapi.Message{Text: "ASf"}}, want: true},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			res := IsUpdateMessage(tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}

func TestIsUpdateBotMessage(t *testing.T) {
	testCases := map[string]struct {
		input tgbotapi.Update
		want  bool
	}{
		"Bot Command":  {input: tgbotapi.Update{Message: &tgbotapi.Message{Text: "/asdf", Entities: []tgbotapi.MessageEntity{{Offset: 0, Length: 5, Type: "bot_command"}}}}, want: true},
		"Text Message": {input: tgbotapi.Update{Message: &tgbotapi.Message{Text: "ASf"}}, want: false},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			res := IsUpdateBotMessage(tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}

func TestExtactBotCommand(t *testing.T) {
	testCases := map[string]struct {
		input tgbotapi.Update
		want  string
	}{
		"only bot command": {input: tgbotapi.Update{Message: &tgbotapi.Message{Text: "/asdf", Entities: []tgbotapi.MessageEntity{{Offset: 0, Length: 5, Type: "bot_command"}}}}, want: "asdf"},
		"mixed case":       {input: tgbotapi.Update{Message: &tgbotapi.Message{Text: "/AsDf", Entities: []tgbotapi.MessageEntity{{Offset: 0, Length: 5, Type: "bot_command"}}}}, want: "AsDf"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			res := ExtractBotCommand(tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}
