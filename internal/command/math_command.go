package command

import (
	"fmt"
	"strconv"

	utils "oat431/try-go-telegram-bot/pkg/util"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleIsPrimeCommand(bot *tgbotapi.BotAPI, chatID int64, args []string) error {
	if len(args) != 1 {
		return sendText(bot, chatID, "Usage: /is-prime <number>")
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		return sendText(bot, chatID, "Invalid number. Usage: /is-prime <number>")
	}

	if utils.IsPrime(n) {
		return sendText(bot, chatID, fmt.Sprintf("%d is prime", n))
	}

	return sendText(bot, chatID, fmt.Sprintf("%d is not prime", n))
}

func HandleGCDCommand(bot *tgbotapi.BotAPI, chatID int64, args []string) error {
	if len(args) != 2 {
		return sendText(bot, chatID, "Usage: /gcd <a> <b>")
	}

	a, err := strconv.Atoi(args[0])
	if err != nil {
		return sendText(bot, chatID, "Invalid first number. Usage: /gcd <a> <b>")
	}

	b, err := strconv.Atoi(args[1])
	if err != nil {
		return sendText(bot, chatID, "Invalid second number. Usage: /gcd <a> <b>")
	}

	result := utils.GCD(a, b)
	return sendText(bot, chatID, fmt.Sprintf("gcd(%d, %d) = %d", a, b, result))
}

func HandleLCMCommand(bot *tgbotapi.BotAPI, chatID int64, args []string) error {
	if len(args) != 2 {
		return sendText(bot, chatID, "Usage: /lcm <a> <b>")
	}

	a, err := strconv.Atoi(args[0])
	if err != nil {
		return sendText(bot, chatID, "Invalid first number. Usage: /lcm <a> <b>")
	}

	b, err := strconv.Atoi(args[1])
	if err != nil {
		return sendText(bot, chatID, "Invalid second number. Usage: /lcm <a> <b>")
	}

	if a == 0 && b == 0 {
		return sendText(bot, chatID, "lcm(0, 0) is undefined")
	}

	result := utils.LCM(a, b)
	return sendText(bot, chatID, fmt.Sprintf("lcm(%d, %d) = %d", a, b, result))
}

func sendText(bot *tgbotapi.BotAPI, chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(msg)
	return err
}
