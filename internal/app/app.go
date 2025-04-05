package app

import (
	"log/slog"

	"github.com/jinpain/patient-recording-tg-bot/internal/app/bot"
)

type App struct {
	Bot *bot.Bot
	log *slog.Logger
}

func New(log *slog.Logger, token string) *App {
	bot := bot.New(log, token)

	return &App{
		Bot: bot,
		log: log,
	}
}
