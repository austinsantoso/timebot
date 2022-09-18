package timebot

import (
	"testing"

	"github.com/austinsantoso/timebot/internal/mocks/client/mock_telegram"
	"github.com/austinsantoso/timebot/internal/time"
	"github.com/austinsantoso/timebot/internal/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// 1662797475288 == Saturday 10 September 2022 at 8:11:15.288 AM (+00:00)
func newSetTimeModuleProvider(timestamp int64) currentTimeProvider {
	return func() *time.TimeModule {
		return time.NewSetTimeModule(timestamp)
	}
}

func TestNewSetTimeModuleProvider(t *testing.T) {
	a := newSetTimeModuleProvider(123123123123)
	b := newSetTimeModuleProvider(123123123123)
	assert.Equal(t, a().String(), b().String())
}

func TestHandleNowCommand(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	chatId := 1
	message := "/now"
	timeProvider := newSetTimeModuleProvider(1662797475288)
	testUpdate := util.NewUpdateBotMessage(int64(chatId), message)
	mockBot := mock_telegram.NewMockTelegramBotClient(mockCtrl)

	testTime := timeProvider()
	testMessageConfig := util.NewMessageConfig(int64(chatId), testTime.String())

	mockBot.EXPECT().Send(testMessageConfig).Times(1)

	bot := NewCustomTimeBot(timeProvider)
	bot.HandleUpdate(mockBot, testUpdate)
}

func TestHandleHelpCommand(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	chatId := 1
	message := "/help"
	result := `
	Hello I am timebot 
	<required field type>

	/help - get help message\n

	/now  - Get the current time
	/msb <number> - Get the time number of days ago 
	/sb <number> - Get the time number of days ago
	/mb <number> - Get the time number of days ago
	/hb <number> - Get the time number of days ago
	/db <number> - Get the time number of days ago
	/wb <number> - Get the time number of week ago
	`
	timeProvider := newSetTimeModuleProvider(1662797475288)
	testUpdate := util.NewUpdateBotMessage(int64(chatId), message)
	mockBot := mock_telegram.NewMockTelegramBotClient(mockCtrl)
	testMessageConfig := util.NewMessageConfig(int64(chatId), result)

	mockBot.EXPECT().Send(testMessageConfig).Times(1)

	bot := NewCustomTimeBot(timeProvider)
	bot.HandleUpdate(mockBot, testUpdate)
}

func TestTimeBefore(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testCases := map[string]map[string]struct {
		message string
		result  string
	}{
		"milliseconds": {
			"valid input": {
				message: "/msb 1",
				result:  "Saturday 10 September 2022 at 8:11:15.287 AM (+00:00)",
			},
			"no args": {
				message: "/msb",
				result:  "invalid syntax, please use /msb <number>",
			},
			"invalid args": {
				message: "/msb abc",
				result:  "invalid syntax, please use /msb <number>",
			},
		},
		"seconds": {
			"valid input": {
				message: "/sb 1",
				result:  "Saturday 10 September 2022 at 8:11:14.288 AM (+00:00)",
			},
			"no args": {
				message: "/sb",
				result:  "invalid syntax, please use /sb <number>",
			},
			"invalid args": {
				message: "/sb abc",
				result:  "invalid syntax, please use /sb <number>",
			},
		},
		"minutes": {
			"valid input": {
				message: "/mb 1",
				result:  "Saturday 10 September 2022 at 8:10:15.288 AM (+00:00)",
			},
			"no args": {
				message: "/mb",
				result:  "invalid syntax, please use /mb <number>",
			},
			"invalid args": {
				message: "/mb abc",
				result:  "invalid syntax, please use /mb <number>",
			},
		},
		"hours": {
			"valid input": {
				message: "/hb 1",
				result:  "Saturday 10 September 2022 at 7:11:15.288 AM (+00:00)",
			},
			"no args": {
				message: "/hb",
				result:  "invalid syntax, please use /hb <number>",
			},
			"invalid args": {
				message: "/hb abc",
				result:  "invalid syntax, please use /hb <number>",
			},
		},
		"days": {
			"valid input": {
				message: "/db 1",
				result:  "Friday 09 September 2022 at 8:11:15.288 AM (+00:00)",
			},
			"no args": {
				message: "/db",
				result:  "invalid syntax, please use /db <number>",
			},
			"invalid args": {
				message: "/db abc",
				result:  "invalid syntax, please use /db <number>",
			},
		},
		"weeks": {
			"valid input": {
				message: "/wb 1",
				result:  "Saturday 03 September 2022 at 8:11:15.288 AM (+00:00)",
			},
			"no args": {
				message: "/wb",
				result:  "invalid syntax, please use /wb <number>",
			},
			"invalid args": {
				message: "/wb abc",
				result:  "invalid syntax, please use /wb <number>",
			},
		},
	}

	for name, subGroup := range testCases {
		for subName, tc := range subGroup {
			t.Run(name+"_"+subName, func(t *testing.T) {
				chatId := 1
				message := tc.message
				timeProvider := newSetTimeModuleProvider(1662797475288)
				testUpdate := util.NewUpdateBotMessage(int64(chatId), message)
				testMessageConfig := util.NewMessageConfig(int64(chatId), tc.result)

				mockBot := mock_telegram.NewMockTelegramBotClient(mockCtrl)
				mockBot.EXPECT().Send(testMessageConfig).Times(1)

				bot := NewCustomTimeBot(timeProvider)
				bot.HandleUpdate(mockBot, testUpdate)
			})
		}
	}

}
