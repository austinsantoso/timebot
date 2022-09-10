package main

import (
	"encoding/json"
	"log"

	"github.com/austinsantoso/timebot/internal/config"
	"github.com/austinsantoso/timebot/internal/handler/timebot"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Printf("Received value from telegram: %+v\n", request.Body)

	conf := config.NewConfig()
	bot, err := tgbotapi.NewBotAPI(conf.GetBotTokenFromEnvVar())

	if err != nil {
		log.Fatal(err)
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, nil
	}

	// parse the request
	update := tgbotapi.Update{}
	json.Unmarshal([]byte(request.Body), &update)

	handler := timebot.NewBot()

	handler.HandleUpdate(bot, update)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
