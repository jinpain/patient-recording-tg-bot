package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) NewPhoto(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	h.response.PhotoUser(bot, message)
}
