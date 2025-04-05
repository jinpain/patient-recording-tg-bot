package main

import (
	"github.com/jinpain/patient-recording-tg-bot/internal/config"
	"github.com/jinpain/patient-recording-tg-bot/internal/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)

	log.Info("launch a bot...")
}
