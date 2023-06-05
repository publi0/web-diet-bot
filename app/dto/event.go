package dto

type Event struct {
	UpdateID      int64         `json:"update_id"`
	Message       Message       `json:"message"`
	CallbackQuery CallbackQuery `json:"callback_query"`
}

type CallbackQuery struct {
	ID           string               `json:"id"`
	From         CallbackQueryFrom    `json:"from"`
	Message      CallbackQueryMessage `json:"message"`
	ChatInstance string               `json:"chat_instance"`
	Data         string               `json:"data"`
}

type CallbackQueryFrom struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	LanguageCode string `json:"language_code"`
}

type CallbackQueryMessage struct {
	MessageID   int64       `json:"message_id"`
	From        From        `json:"from"`
	Chat        Chat        `json:"chat"`
	Date        int64       `json:"date"`
	Text        string      `json:"text"`
	ReplyMarkup ReplyMarkup `json:"reply_markup"`
}

type Chat struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
}

type From struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

type ReplyMarkup struct {
	InlineKeyboard [][]InlineKeyboard `json:"inline_keyboard"`
}

type InlineKeyboard struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type Message struct {
	Text      string            `json:"text"`
	MessageID int64             `json:"message_id"`
	From      CallbackQueryFrom `json:"from"`
	Chat      Chat              `json:"chat"`
	Date      int64             `json:"date"`
	Photo     []Photo           `json:"photo"`
}

type Photo struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
}
