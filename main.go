package main

import (
	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

func main() {

	key, _ := os.LookupEnv("API_KEY")
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		URL:    "https://api.telegram.org",
		Token:  key,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "DAROVA "+m.Sender.FirstName+" "+m.Sender.LastName)
	})

	b.Handle("hello", func(m *tb.Message) {
		b.Send(m.Sender, "priveti "+m.Sender.FirstName+" "+m.Sender.LastName)
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		b.Send(m.Sender, m.Text)
	})
	b.Handle(tb.OnSticker, func(m *tb.Message) {
		file := &tb.Sticker{
			File: tb.FromDisk("sticker.webp"),
		}
		b.Send(m.Sender, "I DON'T UNDERSTAND STICKERS")
		b.Send(m.Sender, file)
	})

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Hello, please tell me how can i call you. Enter your name after /name")
	})

	b.Handle("/name", func(m *tb.Message) {

	})

	b.Start()
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
