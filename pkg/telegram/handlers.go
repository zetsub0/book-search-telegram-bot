package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zetsub0/book-search-telegram-bot/internal/stringToLink"
	"log"
	"time"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know that command")
	switch message.Command() {
	case commandStart:
		msg.Text = "Отправь мне название книги или ISBN"
		tm := time.Unix(int64(message.Date), 0)
		log.Printf("[%s] %s \t[%s]", message.From.UserName, message.Text, tm)
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}
	default:
		tm := time.Unix(int64(message.Date), 0)
		log.Printf("[%s] %s \t[%s]", message.From.UserName, message.Text, tm)

		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil

}

func (b *Bot) handleMessage(message *tgbotapi.Message) {

	tm := time.Unix(int64(message.Date), 0)
	b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Поиск книги "+message.Text))
	time.Sleep(200 * time.Millisecond)
	log.Printf("[%s] %s \t[%s]", message.From.UserName, message.Text, tm)
	links := stringToLink.StringToLink(message.Text)
	for key, val := range links {
		message.Text = "[" + key + "]" + "(" + val + ")"
		time.Sleep(50 * time.Millisecond)
		msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
		msg.ParseMode = tgbotapi.ModeMarkdown

		b.bot.Send(msg)
	}

}
