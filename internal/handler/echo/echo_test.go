package echo

import (
	"testing"

	"github.com/austinsantoso/timebot/internal/mocks/client/mock_telegram"
	"github.com/austinsantoso/timebot/internal/util"
	"github.com/golang/mock/gomock"
)

func TestHandleUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	chatId := 1
	message := "hello"
	testUpdate := util.NewUpdateMessage(int64(chatId), message)
	testMessageConfig := util.NewMessageConfig(int64(chatId), message)
	mockBot := mock_telegram.NewMockTelegramBotClient(mockCtrl)

	mockBot.EXPECT().Send(testMessageConfig).Times(1)

	echoBot := NewEchoBot()
	echoBot.HandleUpdate(mockBot, testUpdate)

}
