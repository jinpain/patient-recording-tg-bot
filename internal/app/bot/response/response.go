package response

import "log/slog"

type Response struct {
	log             *slog.Logger
	RegistrarChatId int64
	photoPath       string
}

func New(log *slog.Logger, registrarChatId int64, photoPath string) *Response {
	return &Response{
		log:             log,
		RegistrarChatId: registrarChatId,
		photoPath:       photoPath,
	}
}
