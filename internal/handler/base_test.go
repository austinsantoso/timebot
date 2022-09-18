package handler

import (
	"testing"

	"github.com/austinsantoso/timebot/internal/mocks/client/mock_telegram"
	"github.com/austinsantoso/timebot/internal/util"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSendMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	chatId := 1
	message := "hello"
	testUpdate := util.NewUpdateMessage(int64(chatId), "random teext")
	testMessageConfig := util.NewMessageConfig(int64(chatId), message)
	mockBot := mock_telegram.NewMockTelegramBotClient(mockCtrl)

	mockBot.EXPECT().Send(testMessageConfig).Times(1)

	SendMessage(mockBot, testUpdate, message)
}

func TestIsUpdateTextMessage(t *testing.T) {
	testCases := map[string]struct {
		input tgbotapi.Update
		want  bool
	}{
		"TrueCase":  {input: tgbotapi.Update{Message: &tgbotapi.Message{Text: "ASf"}}, want: true},
		"FalseCase": {input: tgbotapi.Update{Message: &tgbotapi.Message{Sticker: &tgbotapi.Sticker{}}}, want: false},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			res := IsUpdateTextMessage(tc.input)
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

func TestExtractBotCommand(t *testing.T) {
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

func TestExtractBotCommandArguments(t *testing.T) {
	testCases := map[string]struct {
		input tgbotapi.Update
		want  []string
	}{
		"no arguments": {input: tgbotapi.Update{Message: &tgbotapi.Message{Text: "/asdf", Entities: []tgbotapi.MessageEntity{{Offset: 0, Length: 5, Type: "bot_command"}}}}, want: []string{}},
		"one argument": {input: tgbotapi.Update{Message: &tgbotapi.Message{Text: "/AsDf hello", Entities: []tgbotapi.MessageEntity{{Offset: 0, Length: 5, Type: "bot_command"}}}}, want: []string{"hello"}},
		"2 arguments multiple spaces": {input: tgbotapi.Update{Message: &tgbotapi.Message{Text: "/AsDf    hello\t 	  world", Entities: []tgbotapi.MessageEntity{{Offset: 0, Length: 5, Type: "bot_command"}}}}, want: []string{"hello", "world"}},
		"with back slash": {input: tgbotapi.Update{Message: &tgbotapi.Message{Text: "/AsDf    hello\\t 	  world", Entities: []tgbotapi.MessageEntity{{Offset: 0, Length: 5, Type: "bot_command"}}}}, want: []string{"hello\\t", "world"}},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			res := ExtractBotCommandArguments(tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}
