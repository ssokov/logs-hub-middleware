package middleware

import (
	"context"
	"fmt"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type LogMiddleware struct {
	client      Client
	endpoint    string
	serviceName string
	// [TODO] api key
}

type BadStatusCodeError struct {
	Code int
}

func (e *BadStatusCodeError) Error() string {
	return fmt.Sprintf("bad status code: %d", e.Code)
}
func (e *BadStatusCodeError) StatusCode() int {
	return e.Code
}
func (e *BadStatusCodeError) Is(target error) bool {
	if target == nil {
		return e == nil
	}
	_, ok := target.(*BadStatusCodeError)
	return ok
}

// endpoint - URL бэкенда для отправки логов
func NewLogMiddleware(endpoint string, ServiceName string) *LogMiddleware {
	return &LogMiddleware{
		endpoint:    endpoint,
		serviceName: ServiceName,
		client:      *NewClient(endpoint, nil),
	}
}

func (l *LogMiddleware) TelegramLogMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		if update.Message != nil {

			go func(ctx context.Context, endpoint string, update *models.Update) {
				defer func() {
					if r := recover(); r != nil {
						log.Printf("panic log middleware: %v", r)
					}
				}()
				convertedUpdate := ToModelsUpdate(*update, l.serviceName)
				err := l.client.Logs.AddTelegramLog(ctx, convertedUpdate)
				if err != nil {
					log.Printf("Failed send log to backend. err: %v", err)
				}
			}(ctx, l.endpoint, update)

		}
		next(ctx, bot, update)
	}
}
