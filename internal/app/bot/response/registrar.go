package response

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinpain/patient-recording-tg-bot/internal/app/bot/keyboard"
	"github.com/jinpain/patient-recording-tg-bot/internal/common"
)

func (r *Response) RecordingRegistrar(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewPhoto(r.RegistrarChatId, tgbotapi.FileID(message.Photo[0].FileID))

	msg.Caption = "Отправьте ответ пользователю"

	msg.ReplyMarkup = keyboard.RegistrarConfirmRecordingKeyboard

	if _, err := bot.Send(msg); err != nil {
		r.log.Error("PhotoUser", slog.Any("error", err))
	}
}

func (r *Response) ConfirmRecordingReg(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	if common.RecordingInProgress == nil {
		forwardMsg := tgbotapi.NewForward(message.Chat.ID, message.Chat.ID, message.MessageID)
		if _, err := bot.Send(forwardMsg); err != nil {
			r.log.Error("PhotoUser", slog.Any("error", err))
		}

		common.RecordingInProgress = &common.Recording{}

		common.RecordingInProgress.Main = message.MessageID

		common.RecordingInProgress.Forward = forwardMsg.MessageID
	}
}

func (r *Response) ResponseReg(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

}
