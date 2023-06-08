package validation_chain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"web-diet-bot/client"
	"web-diet-bot/dto"
	"web-diet-bot/respository"
)

type AuthAction struct {
	next   Action
	Bot    *tgbotapi.BotAPI
	Repo   respository.Repository
	Client client.WebDiet
}

func (a AuthAction) Execute(event *dto.Event) {
	var user int64
	var messageId int64
	if event.Message.MessageID != 0 {
		user = event.Message.From.ID
		messageId = event.Message.MessageID
	} else {
		user = event.CallbackQuery.From.ID
		messageId = event.CallbackQuery.Message.MessageID
	}

	userFound, err := a.Repo.FindUser(strconv.FormatInt(user, 10))
	if err != nil {
		return
	}

	if userFound.Auth.P != "" {
		a.next.Execute(event)
	}

	msg = tgbotapi.NewMessage(messageId, "Você ainda não está logado no web diet, digite seu ano de nascimento")
	v.Bot.Send(msg)
}

func (a AuthAction) SetNext(action Action) {
	a.next = action
}
