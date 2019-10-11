package main

import (
	"log"
	"math/rand"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gsora/garbagetextbot/garbage"
)

func handleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.InlineQuery != nil {
		log.Printf("[INLINE] new query sent in by %s -> %s\n", update.InlineQuery.From.UserName, update.InlineQuery.Query)
		payload := []interface{}{}

		data := garbageString(update.InlineQuery.Query)
		article := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, "moderate garbage", data)
		article.Description = data
		payload = append(payload, article)

		worseData := letterGarbageString(update.InlineQuery.Query)
		worseArticle := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID+"worse", "a lot of fucking garbage", worseData)
		worseArticle.Description = worseData
		payload = append(payload, worseArticle)

		inlineConf := tgbotapi.InlineConfig{
			InlineQueryID: update.InlineQuery.ID,
			IsPersonal:    true,
			CacheTime:     0,
			Results:       payload,
		}

		if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
			log.Println(err)
		}
	} else if update.Message != nil {
		log.Printf("[MESSAGE] new message sent in by %s -> %s\n", update.Message.From.UserName, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, garbageString(update.Message.Text))
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}

}

func garbageString(str string) (ret string) {
	rand.Seed(time.Now().UnixNano())

	wordsVector := strings.Split(str, " ")
	for index, substr := range wordsVector {
		n := randomInt(0, 5)
		newWord := ""
		switch n {
		case 0:
			newWord += garbage.Nibber(substr)
		case 1:
			newWord += garbage.UpDown(substr)
		case 2:
			newWord += garbage.Uwu(substr)
		case 3:
			newWord += garbage.Breath(substr)
		case 4:
			newWord += garbage.Reverse(substr)
		}

		if index+1 != len(wordsVector) {
			newWord += " "
		}

		ret += newWord
	}

	return
}

func letterGarbageString(str string) (ret string) {
	for _, c := range str {
		ret += garbageString(string(c))
	}

	return
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
