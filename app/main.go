package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"os"
	"web-diet-bot/client"
	"web-diet-bot/dto"
	"web-diet-bot/respository"
	"web-diet-bot/validation_chain"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if !validateRequest(request.Headers["x-telegram-bot-api-secret-token"]) {
		log.Println("Invalid secret token")
		return events.APIGatewayProxyResponse{StatusCode: http.StatusForbidden}, nil
	}

	event := dto.Event{}
	json.Unmarshal([]byte(request.Body), &event)

	dynamo := dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})))
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Println("Con`t connect to telegram")
		log.Panic(err)
	}

	repo := respository.Repository{
		Db: dynamo,
	}

	clientHttp := http.Client{}

	telegramClient := client.TelegramClient{
		Client: &clientHttp,
	}
	webDiet := client.WebDiet{
		Client: &clientHttp,
	}

	defaultAction := validation_chain.DefaultAction{
		Bot: bot,
	}

	photoAction := validation_chain.PhotoAction{
		Bot:    bot,
		Client: telegramClient,
		Repo:   repo,
	}

	mealTypeAction := validation_chain.MealTypeAction{
		Bot:    bot,
		Repo:   repo,
		Client: webDiet,
	}

	validateUserAction := validation_chain.ValidateUserAction{
		Bot:  bot,
		Repo: repo,
	}

	photoAction.SetNext(&defaultAction)
	mealTypeAction.SetNext(&photoAction)
	validateUserAction.SetNext(&mealTypeAction)
	validateUserAction.Execute(&event)

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}

func validateRequest(token string) bool {
	webhookKey := os.Getenv("AUTH_KEY_WEHOOK")
	return webhookKey == token
}
