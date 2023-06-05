package validation_chain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"regexp"
	"strconv"
	"web-diet-bot/client"
	"web-diet-bot/dto"
	"web-diet-bot/respository"
)

type MealTypeAction struct {
	next   Action
	Bot    *tgbotapi.BotAPI
	Repo   respository.Repository
	Client client.WebDiet
}

func (m *MealTypeAction) Execute(event *dto.Event) {
	data := event.CallbackQuery.Data
	if data == "" || !regexp.MustCompile(`^[1-6]$`).MatchString(data) {
		m.next.Execute(event)
		return
	}

	meals := map[string]string{
		"1": "Café da Manhã",
		"2": "Lanche da Manhã",
		"3": "Almoço",
		"4": "Lanche da Tarde",
		"5": "Jantar",
		"6": "Ceia",
	}

	editMessage := tgbotapi.NewEditMessageText(event.CallbackQuery.Message.Chat.ID, int(event.CallbackQuery.Message.MessageID), meals[data])
	m.Bot.Send(editMessage)

	user, err := m.Repo.FindUser(strconv.FormatInt(event.CallbackQuery.From.ID, 10))
	if err != nil {
		log.Println("error finding user")
		return
	}

	m.Client.UploadPhoto(user.LastPhoto, meals[data], "", user.Auth.N, user.Auth.P)
	m.Repo.UpdateUserPhoto(strconv.FormatInt(event.CallbackQuery.From.ID, 10), "")

	msg := tgbotapi.NewMessage(event.CallbackQuery.Message.Chat.ID, "Enviado ✅🍏")
	m.Bot.Send(msg)
}

func (m *MealTypeAction) SetNext(action Action) {
	m.next = action
}
