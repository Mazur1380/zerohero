package main

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const token = "6349483943:AAGEN47LBBQbnsdjeuelvsC7j5RpZR025Vc"

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		txt := "Hello " + time.Now().Format("2006-Junary-02 15:04")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, txt)

		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {

			log.Fatal(err)
		}
	}

}
