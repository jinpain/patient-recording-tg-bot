package response

import "log/slog"

type Response struct {
	log             *slog.Logger
	RegistrarChatId int64
}

func New(log *slog.Logger, registrarChatId int64) *Response {
	return &Response{
		log:             log,
		RegistrarChatId: registrarChatId,
	}
}
