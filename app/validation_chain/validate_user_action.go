package validation_chain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"web-diet-bot/dto"
	"web-diet-bot/respository"
)

type ValidateUserAction struct {
	next Action
	Bot  *tgbotapi.BotAPI
	Repo respository.Repository
}

func (v *ValidateUserAction) Execute(event *dto.Event) {
	var user int64
	var messageId int64
	if event.Message.MessageID != 0 {
		user = event.Message.From.ID
		messageId = event.Message.MessageID
	} else {
		user = event.CallbackQuery.From.ID
		messageId = event.CallbackQuery.Message.MessageID
	}

	active, err := v.Repo.UserIsActive(strconv.FormatInt(user, 10))
	var msg tgbotapi.MessageConfig
	if err != nil {
		msg = tgbotapi.NewMessage(messageId, "Erro processando seu username 😵")
		v.Bot.Send(msg)
		return
	}
	if !active {
		msg = tgbotapi.NewMessage(messageId, "Vocë ainda não está cadastrado em nossa base! 🥹")
		v.Bot.Send(msg)
		return
	}
	v.next.Execute(event)
}

func (v *ValidateUserAction) SetNext(action Action) {
	v.next = action
}
