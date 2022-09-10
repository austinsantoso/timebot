package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBotToken(t *testing.T) {
	conf := NewConfig()
	conf.Init()

	var expected string = "INSERT_TELEGRAM_BOT_TOKEN"

	assert.Equal(t, expected, conf.GetBotToken(), "wrong expected bot token")
}
