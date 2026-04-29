package service

import (
	"context"
	"oat431/try-go-telegram-bot/internal/handler"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ReceiveUpdates(ctx context.Context, bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	// `for {` means the loop is infinite until we manually stop it
	for {
		select {
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			return
		// receive update from channel and then handle it
		case update, ok := <-updates:
			if !ok {
				return
			}
			handler.HandleUpdate(bot, update)
		}
	}
}
