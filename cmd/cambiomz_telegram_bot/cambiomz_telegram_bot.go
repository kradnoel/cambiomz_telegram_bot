package main

import (
	b "github.com/kradnoel/cambiomz_telegram_bot/internal/bot"
)

func main() {
	bot := b.New()
	bot.Default()
	bot.Run()
}
