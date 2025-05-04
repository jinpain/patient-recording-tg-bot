package app

import (
	"log/slog"

	"github.com/jinpain/patient-recording-tg-bot/internal/app/bot"
	"github.com/jinpain/patient-recording-tg-bot/internal/config"
)

type App struct {
	Bot *bot.Bot
	log *slog.Logger
}

func New(log *slog.Logger, cfg *config.Config) *App {
	bot := bot.New(log, cfg)

	return &App{
		Bot: bot,
		log: log,
	}
}
