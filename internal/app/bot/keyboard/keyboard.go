package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var UserMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Записаться на прием"),
	),
)

var UserConfirmRecordingKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Подтвердить", "confirm_recording"),
		tgbotapi.NewInlineKeyboardButtonData("Отменить", "cancel_recording"),
	),
)

var RegistrarConfirmRecordingKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Ответить", "reg_confirm_recording"),
	),
)
