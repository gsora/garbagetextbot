package main

import (
	"log"
	"strings"

	"github.com/gsora/garbagetextbot/meme"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	inlineMsg string = "you are horrible"
)

var (
	apiKey         string
	webHookURL     string
	debug          bool
	numberOfSpaces int
	fillingSpace   string
)

func main() {
	if !(numberOfSpaces <= 0 || numberOfSpaces == 1) {
		generateSpaces()
	}

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = debug

	log.Printf("Authorized on account %s, garbagetextbot is running!", bot.Self.UserName)

	_, err = bot.RemoveWebhook()
	if err != nil {
		log.Fatal(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 1

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		go meme.HandleUpdate(update, bot)
	}

}

func generateSpaces() {
	fillingSpace = strings.Repeat(" ", numberOfSpaces)
}
