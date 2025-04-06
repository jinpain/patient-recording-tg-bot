package handler

import (
	"log/slog"

	"github.com/jinpain/patient-recording-tg-bot/internal/app/bot/response"
)

type Handler struct {
	log      *slog.Logger
	response *response.Response
}

func New(log *slog.Logger, registrarChatId int64) *Handler {
	return &Handler{
		log:      log,
		response: response.New(log, registrarChatId),
	}
}
