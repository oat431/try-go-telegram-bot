package command

import (
	"oat431/try-go-telegram-bot/pkg/common"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// When we get a command, we react accordingly
func CommandList(bot *tgbotapi.BotAPI, chatId int64, commandText string) error {
	var err error

	parts := strings.Fields(commandText)
	if len(parts) == 0 {
		return nil
	}

	command := strings.SplitN(parts[0], "@", 2)[0]
	args := parts[1:]

	switch command {
	case "/scream":
		common.Screaming = true
		return nil

	case "/whisper":
		common.Screaming = false
		return nil

	case "/menu":
		err = SendMenu(bot, chatId)

	case "/is-prime":
		err = HandleIsPrimeCommand(bot, chatId, args)

	case "/gcd":
		err = HandleGCDCommand(bot, chatId, args)

	case "/lcm":
		err = HandleLCMCommand(bot, chatId, args)
	}

	return err
}
