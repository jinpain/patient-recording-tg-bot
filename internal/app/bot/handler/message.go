package handler

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinpain/patient-recording-tg-bot/internal/config"
)

type Handler struct {
	log *slog.Logger
}

func New(log *slog.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

func (h *Handler) NewMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message, registrars *[]config.Registrar) {
	switch message.Text {
	case "/start":
		h.start(bot, message, registrars)
	default:

	}
}
