package command

import (
	"oat431/try-go-telegram-bot/pkg/common"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// When we get a command, we react accordingly
func CommandList(bot *tgbotapi.BotAPI, chatId int64, command string) error {
	var err error

	switch command {
	case "/scream":
		common.Screaming = true
		break

	case "/whisper":
		common.Screaming = false
		break

	case "/menu":
		err = SendMenu(bot, chatId)
		break
	}

	return err
}
