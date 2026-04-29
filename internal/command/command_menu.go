package command

import (
	"errors"
	"oat431/try-go-telegram-bot/pkg/common"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMenu(bot *tgbotapi.BotAPI, chatId int64) error {
	if bot == nil {
		return errors.New("telegram bot is not initialized")
	}

	msg := tgbotapi.NewMessage(chatId, common.FirstMenu)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = common.FirstMenuMarkup
	_, err := bot.Send(msg)
	return err
}
