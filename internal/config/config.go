package config

import (
	"flag"
)

type Config struct {
	Token string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Init() {
	flag.StringVar(&(c.Token), "token", "INSERT_TELEGRAM_BOT_TOKEN", "Telegram bot token from BotFather")

	flag.Parse()
}

func (c *Config) GetBotToken() string {
	return c.Token
}
