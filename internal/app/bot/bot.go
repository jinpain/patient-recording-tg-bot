package bot

import (
	"fmt"
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

func New(log *slog.Logger, cfg *config.Config) *Bot {
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		panic(fmt.Sprintf("Recovered from panic: %v", err))
	}

	handler := handler.New(log, cfg.Registrar.ChatId, cfg.PhotoPath)

	return &Bot{
		bot:     bot,
		log:     log,
		handler: handler,
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
				b.log.Info("New command", slog.Any("command", update.Message.Text),
					slog.Any("userID", update.Message.Chat.ID))
			} else if update.Message.Photo != nil {
				b.handler.NewPhoto(b.bot, update.Message)
				b.log.Info("New photo", slog.Any("userID", update.Message.Chat.ID))
			} else {
				b.handler.NewMessage(b.bot, update.Message)
				b.log.Info("New message", slog.Any("message", update.Message.Text),
					slog.Any("userID", update.Message.Chat.ID))
			}
		} else if update.CallbackQuery != nil {
			b.handler.NewCallback(b.bot, update.CallbackQuery)
			b.log.Info("New callback")
		}
	}
}
