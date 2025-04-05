package bot

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	log *slog.Logger
}

func New(log *slog.Logger, token string) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Error("Recovered from panic", slog.Any("panic", err))
	}

	return &Bot{
		bot: bot,
		log: log,
	}
}
