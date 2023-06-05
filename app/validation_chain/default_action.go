package validation_chain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"web-diet-bot/dto"
)

type DefaultAction struct {
	next Action
	Bot  *tgbotapi.BotAPI
}

func (d *DefaultAction) Execute(event *dto.Event) {
	msg := tgbotapi.NewMessage(event.Message.Chat.ID, "Olá! Envie a foto da sua refeicão que eu irie encaminhar para o web diet")
	d.Bot.Send(msg)
}

func (d *DefaultAction) SetNext(action Action) {
	d.next = action
}
