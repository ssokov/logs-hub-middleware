package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type LogMiddleware struct {
	Endpoint string
}

var ErrBadStatusCode = errors.New("bad status code")

func NewLogMiddleware(endpoint string) *LogMiddleware {
	return &LogMiddleware{
		Endpoint: endpoint,
	}
}

func (l *LogMiddleware) TelegramLogMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		if update.Message != nil {

			err := SendTelegramLog(ctx, l.Endpoint, update)

			if err != nil {
				log.Printf("Failed send log to backend. err: %v", err)
			}
		}
		next(ctx, bot, update)
	}
}

func SendTelegramLog(ctx context.Context, endpoint string, log *models.Update) error {
	data, err := json.Marshal(log)
	if err != nil {
		return fmt.Errorf("failed marshal telegram log. err: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed create request. err: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed send request. err: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("%w: status %d,",
			ErrBadStatusCode, resp.StatusCode)
	}

	return nil
}
