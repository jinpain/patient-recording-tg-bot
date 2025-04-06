package response

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinpain/patient-recording-tg-bot/internal/app/bot/command"
	"github.com/jinpain/patient-recording-tg-bot/internal/app/bot/keyboard"
)

func (r *Response) StartUser(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, command.Start)

	msg.ReplyMarkup = keyboard.UserMenuKeyboard

	if _, err := bot.Send(msg); err != nil {
		r.log.Error("StartUser", slog.Any("error", err))
	}
}

func (h *Response) RecordingUser(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, command.Recording)

	if _, err := bot.Send(msg); err != nil {
		h.log.Error("RecordingUser", slog.Any("error", err))
	}
}

func (r *Response) PhotoUser(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FileID(message.Photo[0].FileID))

	msg.Caption = "Подтвердите запись на прием"

	msg.ReplyMarkup = keyboard.UserConfirmRecordingKeyboard

	if _, err := bot.Send(msg); err != nil {
		r.log.Error("PhotoUser", slog.Any("error", err))
	}
}

func (h *Response) ConfirmRecordingUser(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msgDel := tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID)

	if _, err := bot.Request(msgDel); err != nil {
		h.log.Error("ConfirmRecordingUser", slog.Any("error", err))
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, command.ConfirmRecording)

	if _, err := bot.Send(msg); err != nil {
		h.log.Error("ConfirmRecordingUser", slog.Any("error", err))
	}
}

func (h *Response) CancelRecordingUser(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	replyMarkup := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, tgbotapi.InlineKeyboardMarkup{InlineKeyboard: make([][]tgbotapi.InlineKeyboardButton, 0)})

	if _, err := bot.Send(replyMarkup); err != nil {
		h.log.Error("ConfirmRecordingUser", slog.Any("error", err))
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, command.CancelRecording)

	if _, err := bot.Send(msg); err != nil {
		h.log.Error("ConfirmRecordingUser", slog.Any("error", err))
	}
}
