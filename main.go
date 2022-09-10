package main

import (
	"log"

	"github.com/austinsantoso/timebot/internal/config"
	handler "github.com/austinsantoso/timebot/internal/handler"

	// echoHandler "github.com/austinsantoso/timebot/internal/handler/echo"
	timebot "github.com/austinsantoso/timebot/internal/handler/timebot"
	botUpdater "github.com/austinsantoso/timebot/internal/updater"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	conf := config.NewConfig()
	conf.Init()

	bot, err := tgbotapi.NewBotAPI(conf.GetBotToken())

	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	var botServer botUpdater.Updater
	botServer = botUpdater.NewSimpleUpdater()

	var h handler.MessageHandler
	// h = echoHandler.NewEchoBot()
	h = timebot.NewBot()

	botServer.Start(bot, h)

}
