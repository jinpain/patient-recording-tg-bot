package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinpain/patient-recording-tg-bot/internal/common"
)

func (h *Handler) NewMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	if h.response.RegistrarChatId == message.Chat.ID {
		if common.RecordingInProgress != nil {

		}
	} else {
		switch message.Text {
		case "Записаться на прием":
			h.response.RecordingUser(bot, message)
		default:

		}
	}
}
