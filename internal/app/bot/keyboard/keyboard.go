package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var RegistrarMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Активные записи"),
		tgbotapi.NewKeyboardButton("Закрытые записи"),
		tgbotapi.NewKeyboardButton("Все записи"),
	),
)

var UserMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Записаться"),
		tgbotapi.NewKeyboardButton("Мои активные записи"),
		tgbotapi.NewKeyboardButton("Все записи"),
	),
)
