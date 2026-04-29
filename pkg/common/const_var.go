package common

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	// Menu texts
	FirstMenu  = "<b>Menu 1</b>\n\nA beautiful menu with a shiny inline button."
	SecondMenu = "<b>Menu 2</b>\n\nA better menu with even more shiny inline buttons."

	// Button texts
	NextButton     = "Next"
	BackButton     = "Back"
	TutorialButton = "Tutorial"

	// Store bot screaming status
	Screaming = false

	// Keyboard layout for the first menu. One button, one row
	FirstMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(NextButton, NextButton),
		),
	)

	// Keyboard layout for the second menu. Two buttons, one per row
	SecondMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(BackButton, BackButton),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(TutorialButton, "https://core.telegram.org/bots/api"),
		),
	)
)
