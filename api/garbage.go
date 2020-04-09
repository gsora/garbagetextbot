package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gsora/garbagetextbot/meme"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var apiKey string

var url string

var bot *tgbotapi.BotAPI

func init() {
	apiKey = meme.GetVariableOrFatal("TG_APIKEY")
	url = meme.GetVariableOrFatal("NOWSH_URL")

	var err error
	bot, err = tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Fatal(err)
	}

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

}

func Handler(w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadAll(r.Body)

	var update tgbotapi.Update
	json.Unmarshal(bytes, &update)

	meme.HandleUpdate(update, bot)
}
