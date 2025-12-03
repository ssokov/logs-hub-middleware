package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type LogMiddleware struct {
	client   Client
	endpoint string
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
func NewLogMiddleware(endpoint string) *LogMiddleware {
	return &LogMiddleware{
		endpoint: endpoint,
		client:   *NewClient(endpoint, nil),
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
				convertedUpdate := ToModelsUpdate(*update)
				err := l.client.Logs.AddTelegramLog(ctx, convertedUpdate)
				if err != nil {
					log.Printf("Failed send log to backend. err: %v", err)
				}
			}(ctx, l.endpoint, update)

		}
		next(ctx, bot, update)
	}
}

func sendTelegramLog(ctx context.Context, endpoint string, log *models.Update) error {
	data, err := json.Marshal(log)
	if err != nil {
		return fmt.Errorf("failed marshal telegram log. err: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed create request. err: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed send request. err: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return &BadStatusCodeError{Code: resp.StatusCode}
	}

	return nil
}
