package validation_chain

import (
	"encoding/base64"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"web-diet-bot/client"
	"web-diet-bot/dto"
	"web-diet-bot/respository"
)

type PhotoAction struct {
	next   Action
	Bot    *tgbotapi.BotAPI
	Client client.TelegramClient
	Repo   respository.Repository
}

func (p *PhotoAction) Execute(event *dto.Event) {
	if event.Message.Photo == nil && event.CallbackQuery.ID == "" {
		p.next.Execute(event)
		return
	}

	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Café da Manhã", "1"),
			tgbotapi.NewInlineKeyboardButtonData("Lanche da Manhã", "2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Almoço", "3"),
			tgbotapi.NewInlineKeyboardButtonData("Lanche da Tarde", "4"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Jantar", "5"),
			tgbotapi.NewInlineKeyboardButtonData("Ceia", "6"),
		),
	)

	image, err := p.Client.DownloadImage(event.Message.Photo[len(event.Message.Photo)-1].FileID)
	if err != nil {
		log.Println(err)
		msg := tgbotapi.NewMessage(event.Message.Chat.ID, "Error processando sua imagem 😵")
		p.Bot.Send(msg)
		return
	}

	base64String := toBase64String(image)
	formatInt := strconv.FormatInt(event.Message.From.ID, 10)
	err = p.Repo.UpdateUserPhoto(formatInt, base64String)
	if err != nil {
		log.Println("Error saving user photo")
		msg := tgbotapi.NewMessage(event.Message.Chat.ID, "Error processando sua imagem 😵")
		p.Bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(event.Message.Chat.ID, "Qual a refeição?")
	msg.ReplyMarkup = numericKeyboard
	p.Bot.Send(msg)
}

func (p *PhotoAction) SetNext(action Action) {
	p.next = action
}

func toBase64String(image []byte) string {
	encodeToString := base64.StdEncoding.EncodeToString(image)
	return "data:image/jpeg;base64," + encodeToString
}
