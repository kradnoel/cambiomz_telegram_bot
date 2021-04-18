package bot

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kradnoel/cambiomz_telegram_bot/internal/http"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Bot struct {
	Token string
}

func init() {
	godotenv.Load(".env")
}

func New() Bot {
	b := Bot{}
	return b
}

func (bt *Bot) Default() {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("$TELEGRAM_TOKEN must be set")
	}

	if token != "" {
		bt.Token = token
	}
}

func VerifyPrivateChannel(m *tb.Message) {
	if !m.Private() {
		return
	}
}

func registerCommands(b *tb.Bot) {
	commands := []tb.Command{
		{Text: "help", Description: "Show help"},
		{Text: "currencies", Description: "List available currencies"},
	}

	err := b.SetCommands(commands)

	if err != nil {
		log.Fatal(err)
		return
	}
}

func helpHandler(b *tb.Bot) string {
	var response string

	commands, err := b.GetCommands()

	if err != nil {
		log.Fatal(err)
		return ""
	}

	response = "I can help you get exchange rate from CambioMZ API.\n\nYou can control me by sending these commands:\n\n"

	for i, cmd := range commands {
		if i < len(commands) {
			response = response + fmt.Sprintf("/%s - %s\n", cmd.Text, cmd.Description)
		}
	}

	return response
}

func (bt *Bot) Run() {
	b, err := tb.NewBot(tb.Settings{
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		Token:  bt.Token,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	registerCommands(b)

	b.Handle("/start", func(m *tb.Message) {
		VerifyPrivateChannel(m)
		b.Send(m.Sender, fmt.Sprintf("Hello from cambiomz_telegram_bot! %s", helpHandler(b)))
	})

	b.Handle("/help", func(m *tb.Message) {
		VerifyPrivateChannel(m)
		b.Send(m.Sender, fmt.Sprintf("%s", helpHandler(b)))
	})

	b.Handle("/currencies", func(m *tb.Message) {
		VerifyPrivateChannel(m)
		currencies := http.GetCurrencies()
		b.Send(m.Sender, fmt.Sprintf("List of availables currencies:\n\n [%s]", currencies))
	})

	log.Printf("Cambiomz Telegram bot Started...")

	b.Start()
}
