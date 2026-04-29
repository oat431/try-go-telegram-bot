package handler

import (
	"log"
	"oat431/try-go-telegram-bot/internal/command"
	"oat431/try-go-telegram-bot/pkg/common"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch {
	// Handle messages
	case update.Message != nil:
		HandleMessage(bot, update.Message)
		break

	// Handle button clicks
	case update.CallbackQuery != nil:
		HandleButton(bot, update.CallbackQuery)
		break
	}
}

func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	user := message.From
	text := message.Text

	if user == nil {
		return
	}

	if bot == nil {
		log.Printf("An error occured: telegram bot is not initialized")
		return
	}

	// Print to console
	log.Printf("%s wrote %s", user.FirstName, text)

	var err error
	if strings.HasPrefix(text, "/") {
		err = command.CommandList(bot, message.Chat.ID, text)
	} else if common.Screaming && len(text) > 0 {
		msg := tgbotapi.NewMessage(message.Chat.ID, strings.ToUpper(text))
		// To preserve markdown, we attach entities (bold, italic..)
		msg.Entities = message.Entities
		_, err = bot.Send(msg)
	} else {
		// This is equivalent to forwarding, without the sender's name
		copyMsg := tgbotapi.NewCopyMessage(message.Chat.ID, message.Chat.ID, message.MessageID)
		_, err = bot.CopyMessage(copyMsg)
	}

	if err != nil {
		log.Printf("An error occured: %s", err.Error())
	}
}

func HandleButton(bot *tgbotapi.BotAPI, query *tgbotapi.CallbackQuery) {
	if bot == nil {
		log.Printf("An error occured: telegram bot is not initialized")
		return
	}

	var text string

	markup := tgbotapi.NewInlineKeyboardMarkup()
	message := query.Message

	if query.Data == common.NextButton {
		text = common.SecondMenu
		markup = common.SecondMenuMarkup
	} else if query.Data == common.BackButton {
		text = common.FirstMenu
		markup = common.FirstMenuMarkup
	}

	callbackCfg := tgbotapi.NewCallback(query.ID, "")
	bot.Send(callbackCfg)

	// Replace menu text and keyboard
	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, text, markup)
	msg.ParseMode = tgbotapi.ModeHTML
	bot.Send(msg)
}
