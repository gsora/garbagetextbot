package api

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

var apiKey string

var url string

func init() {
	apiKey = getVariableOrFatal("TG_APIKEY")
	url = getVariableOrFatal("NOWSH_URL")

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url + "/api/garbage"))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/api/garbage")

	for u := range updates {
		go handleUpdate(u, bot)
	}
}

func getVariableOrFatal(varName string) string {
	v := os.Getenv(varName)
	if v == "" {
		log.Fatal(varName, " environment key required")
	}

	return v
}
