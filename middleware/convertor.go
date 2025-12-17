package middleware

import (
	"strconv"

	"github.com/go-telegram/bot/models"
)

func ToModelsUpdate(update models.Update, serviceName string) Log {
	strTimestamp := strconv.Itoa(update.Message.Date)
	result := Log{
		Type:       "Telegram Update",
		Message: update.Message.Text,
		Tg_nickname: update.Message.From.Username,
		Error_code: 200,
		Tg_user_id: int(update.Message.From.ID),
		Timestamp:  strTimestamp,
		Service_name: serviceName,
	}

	return result
}














































