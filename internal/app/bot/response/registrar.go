package response

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/jinpain/patient-recording-tg-bot/internal/app/bot/keyboard"
	"github.com/jinpain/patient-recording-tg-bot/internal/common"
)

func (r *Response) RecordingRegistrar(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewPhoto(r.RegistrarChatId, tgbotapi.FileID(message.Photo[0].FileID))

	msg.Caption = fmt.Sprintf("Отправьте ответ пользователю №%v", message.Chat.ID)

	msg.ReplyMarkup = keyboard.RegistrarConfirmRecordingKeyboard

	if _, err := bot.Send(msg); err != nil {
		r.log.Error("PhotoUser", slog.Any("error", err))
	}
}

func (r *Response) ResponseReg(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	if common.MessageInProgress == nil {
		forwardMsg := tgbotapi.NewForward(message.Chat.ID, message.Chat.ID, message.MessageID)

		sentForwardMsg, err := bot.Send(forwardMsg)
		if err != nil {
			r.log.Error("PhotoUser", slog.Any("error", err))
		}

		str := strings.Split(message.Caption, "№")[1]

		chatId, _ := strconv.ParseInt(str, 10, 64)

		fmt.Println(chatId, message.Caption, str)

		common.MessageInProgress = &common.Message{
			ChatId: chatId,
		}

		common.MessageInProgress.MessagesId = append(common.MessageInProgress.MessagesId, message.MessageID)

		common.MessageInProgress.MessagesId = append(common.MessageInProgress.MessagesId, sentForwardMsg.MessageID)
	}

	fmt.Println(common.MessageInProgress)
}

func (r *Response) ResponseConfirmationReg(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	msg.ReplyMarkup = keyboard.RegistrarSendResponseRecordingKeyboard

	if _, err := bot.Send(msg); err != nil {
		r.log.Error("ResponseReg", slog.Any("error", err))
	}

	common.MessageInProgress.MessagesId = append(common.MessageInProgress.MessagesId, message.MessageID)
}

func (r *Response) ConfirmResponseReg(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(common.MessageInProgress.ChatId, message.Text)

	if _, err := bot.Send(msg); err != nil {
		r.log.Error("ConfirmResponseReg", slog.Any("error", err))
	}
	r.DeleteMessagesCache(bot, message, false)
}

func (r *Response) DeleteMessagesCache(bot *tgbotapi.BotAPI, message *tgbotapi.Message, isCancel bool) {
	for i := range common.MessageInProgress.MessagesId {
		if isCancel && i == 0 {
			continue
		}

		msg := tgbotapi.NewDeleteMessage(message.Chat.ID, common.MessageInProgress.MessagesId[i])
		if _, err := bot.Request(msg); err != nil {
			r.log.Error("DeleteMessagesCache", slog.Any("error", err))
		}
	}

	msg := tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID)
	if _, err := bot.Request(msg); err != nil {
		r.log.Error("DeleteMessagesCache", slog.Any("error", err))
	}

	common.MessageInProgress = nil
}
