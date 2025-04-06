package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) NewCallback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	switch callback.Data {
	case "confirm_recording":
		h.response.ConfirmRecordingUser(bot, callback.Message)
		h.response.RecordingRegistrar(bot, callback.Message)
	case "cancel_recording":
		h.response.CancelRecordingUser(bot, callback.Message)
	case "reg_confirm_recording":
		h.response.ConfirmRecordingReg(bot, callback.Message)
	default:

	}
}
