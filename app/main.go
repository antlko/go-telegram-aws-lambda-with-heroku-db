package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tbot/s_tgapp/botapp"
	"github.com/tbot/s_tgapp/database"
	bot "gopkg.in/tucnak/telebot.v2"
	"log"
)

var (
	tgBot *bot.Bot
)

func main() {
	var err error
	tgBot, err = botapp.CreateBot(database.NewPostgresDB())
	if err != nil {
		log.Fatal(err)
	}
	lambda.Start(HandlerBotWebhook)
}

func HandlerBotWebhook(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var u bot.Update
	if err := json.Unmarshal([]byte(req.Body), &u); err == nil {
		tgBot.ProcessUpdate(u)
	}
	return events.APIGatewayProxyResponse{Body: "ok", StatusCode: 200}, nil
}
