package bot

import (
	"fmt"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/jinpain/patient-recording-tg-bot/internal/app/bot/handler"
)

type Bot struct {
	bot       *tgbotapi.BotAPI
	log       *slog.Logger
	handler   *handler.Handler
	registrar int64
}

func New(log *slog.Logger, token string, registrar int64) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(fmt.Sprintf("Recovered from panic: %v", err))
	}

	handler := handler.New(log, registrar)

	return &Bot{
		bot:       bot,
		log:       log,
		handler:   handler,
		registrar: registrar,
	}
}

func (b *Bot) MustRun() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 1

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				b.handler.NewCommand(b.bot, update.Message)
			} else if update.Message.Photo != nil {
				b.handler.NewPhoto(b.bot, update.Message)
			} else {
				b.handler.NewMessage(b.bot, update.Message)
			}
		} else if update.CallbackQuery != nil {
			b.handler.NewCallback(b.bot, update.CallbackQuery)
		}
	}
}
