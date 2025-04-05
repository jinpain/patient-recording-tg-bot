package bot

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinpain/patient-recording-tg-bot/internal/app/bot/handler"
	"github.com/jinpain/patient-recording-tg-bot/internal/config"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	log     *slog.Logger
	handler *handler.Handler
}

func New(log *slog.Logger, token string) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Error("Recovered from panic", slog.Any("panic", err))
	}

	handler := handler.New(log)

	return &Bot{
		bot:     bot,
		log:     log,
		handler: handler,
	}
}

func (b *Bot) MustRun(registrars *[]config.Registrar) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 1

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			b.handler.NewMessage(b.bot, update.Message, registrars)
		}
	}
}
