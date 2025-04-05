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

	log.Info("launch a bot...")

	bot := app.New(log, cfg.Token)

	log.Info("bot started", slog.Any("main", bot))
}
