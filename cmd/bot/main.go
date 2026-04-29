package main

import "oat431/try-go-telegram-bot/internal/config"

func main() {
	config.LoadEnvConfig()
	config.StartTelegramBot()
}
