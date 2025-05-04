package response

import (
	"fmt"
	"log"
	"log/slog"
	"time"

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
	photo := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FilePath(h.photoPath))
	photo.Caption = command.Recording
	photo.ParseMode = tgbotapi.ModeHTML

	if _, err := bot.Send(photo); err != nil {
		log.Fatalln(err)
	}
}

func (r *Response) PhotoUser(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewCopyMessage(message.Chat.ID, message.Chat.ID, message.MessageID)
	msg.ReplyMarkup = &keyboard.UserConfirmRecordingKeyboard

	msgDel := tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID)

	if _, err := bot.Send(msg); err != nil {
		r.log.Error("CopyMessage", slog.Any("error", err))
	}

	if _, err := bot.Request(msgDel); err != nil {
		r.log.Error("CopyMessage", slog.Any("error", err))
	}
}

func (h *Response) ConfirmRecordingUser(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msgDelRM := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, keyboard.RemoveKeyboard)

	newMsg, err := bot.Send(msgDelRM)
	if err != nil {
		h.log.Error("ConfirmRecordingUser", slog.Any("error", err))
	}

	number := time.Now().Unix()

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(command.ConfirmRecording, number))
	msg.ParseMode = tgbotapi.ModeHTML

	if _, err := bot.Send(msg); err != nil {
		h.log.Error("ConfirmRecordingUser", slog.Any("error", err))
	}

	h.RecordingRegistrar(bot, &newMsg, number, message.Chat.ID)
}

func (h *Response) CancelRecordingUser(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msgDel := tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID)

	if _, err := bot.Request(msgDel); err != nil {
		h.log.Error("CancelRecordingUser", slog.Any("error", err))
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, command.CancelRecording)

	if _, err := bot.Send(msg); err != nil {
		h.log.Error("ConfirmRecordingUser", slog.Any("error", err))
	}
}
