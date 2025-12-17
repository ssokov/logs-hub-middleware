package middleware

import (
	"github.com/go-telegram/bot/models"
)

func ToModelsUpdate(update models.Update, serviceName string) LogReq {

	return LogReq{
		Type:         "Telegram Update",
		Message:      update.Message.Text,
		Tg_nickname:  update.Message.From.Username,
		Tg_user_id:   int(update.Message.From.ID),
		Service_name: serviceName,
	}
}
