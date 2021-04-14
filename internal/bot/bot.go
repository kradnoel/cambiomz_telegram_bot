package bot

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
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

func (bt *Bot) Run() {
	b, err := tb.NewBot(tb.Settings{
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		Token:  bt.Token,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		VerifyPrivateChannel(m)
		b.Send(m.Sender, "Hello from cambiomz_telegram_bot!")
	})

	log.Printf("Cambiomz Telegram bot Started...")

	b.Start()
}
