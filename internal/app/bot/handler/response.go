package handler

import (
	"fmt"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinpain/patient-recording-tg-bot/internal/app/bot/command"
	"github.com/jinpain/patient-recording-tg-bot/internal/app/bot/keyboard"
	"github.com/jinpain/patient-recording-tg-bot/internal/config"
)

func (h *Handler) start(bot *tgbotapi.BotAPI, message *tgbotapi.Message, registrars *[]config.Registrar) {
	msg := tgbotapi.NewMessage(message.Chat.ID, command.Start)

	fmt.Println(message.Chat.ID, registrars)

	if isRegistrar(message.Chat.ID, registrars) {
		msg.ReplyMarkup = keyboard.RegistrarMenuKeyboard
	} else {
		msg.ReplyMarkup = keyboard.UserMenuKeyboard
	}

	if _, err := bot.Send(msg); err != nil {
		h.log.Error("start response", slog.Any("error", err))
	}
}

func isRegistrar(chatId int64, registrars *[]config.Registrar) bool {
	for _, v := range *registrars {
		if v.ChatId == chatId {
			return true
		}
	}

	return false
}
