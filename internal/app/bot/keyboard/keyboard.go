package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var UserMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Записаться на прием"),
	),
)

var UserConfirmRecordingKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Подтвердить запись", "confirm_recording"),
		tgbotapi.NewInlineKeyboardButtonData("Отменить запись", "cancel_recording"),
	),
)

var RegistrarConfirmRecordingKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Ответить", "reg_response"),
	),
)

var RegistrarSendResponseRecordingKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Подтвердить", "reg_resp_confirm_recording"),
		tgbotapi.NewInlineKeyboardButtonData("Отменить", "reg_resp_cancel_recording"),
	),
)

var RemoveKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	make([]tgbotapi.InlineKeyboardButton, 0),
)
