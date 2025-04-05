package main

import (
	"log/slog"

	"github.com/jinpain/patient-recording-tg-bot/internal/app"
	"github.com/jinpain/patient-recording-tg-bot/internal/config"
	"github.com/jinpain/patient-recording-tg-bot/internal/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)

	log.Info("launch a application...")

	application := app.New(log, cfg.Token)

	log.Info("application started", slog.Any("main", application))

	log.Info("bot running...", slog.Any("main", application.Bot))

	application.Bot.MustRun(&cfg.Registrars)
}
