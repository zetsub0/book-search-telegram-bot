package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zetsub0/book-search-telegram-bot/pkg/telegram"
	"log"
	"os"
)

type Conf struct {
	Token string
}

var conf Conf

func init() {
	conf.Token = getEnv("BOT_API", "")

}

func main() {
	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		log.Panic("the api key is not valid")
	}

	//bot.Debug = true
	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
