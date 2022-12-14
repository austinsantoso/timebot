// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/austinsantoso/timebot/internal/client/telegram (interfaces: TelegramBotClient)

// Package mock_telegram is a generated GoMock package.
package mock_telegram

import (
	reflect "reflect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	gomock "github.com/golang/mock/gomock"
)

// MockTelegramBotClient is a mock of TelegramBotClient interface.
type MockTelegramBotClient struct {
	ctrl     *gomock.Controller
	recorder *MockTelegramBotClientMockRecorder
}

// MockTelegramBotClientMockRecorder is the mock recorder for MockTelegramBotClient.
type MockTelegramBotClientMockRecorder struct {
	mock *MockTelegramBotClient
}

// NewMockTelegramBotClient creates a new mock instance.
func NewMockTelegramBotClient(ctrl *gomock.Controller) *MockTelegramBotClient {
	mock := &MockTelegramBotClient{ctrl: ctrl}
	mock.recorder = &MockTelegramBotClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelegramBotClient) EXPECT() *MockTelegramBotClientMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockTelegramBotClient) Send(arg0 tgbotapi.Chattable) (tgbotapi.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(tgbotapi.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockTelegramBotClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTelegramBotClient)(nil).Send), arg0)
}
