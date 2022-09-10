package main

import (
	"encoding/json"
	"log"

	"github.com/austinsantoso/timebot/internal/config"
	messageHandler "github.com/austinsantoso/timebot/internal/handler"

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

	var response string
	if messageHandler.IsUpdateBotMessage(update) {
		response = "bot command"
	} else {
		response = update.Message.Text
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)

	bot.Send(msg)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

/*
"message":{"message_id":26,"from":{"id":682521999,"is_bot":false,"first_name":"Austin","last_name":"Santoso","username":"austinsantoso","language_code":"en"},"chat":{"id":682521999,"first_name":"Austin","last_name":"Santoso","username":"austinsantoso","type":"private"},"date":1662792710,"text":"adfasdf"}} IsBase64Encoded:false}
*/
