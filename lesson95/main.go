package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/redis/go-redis/v9"
)

const token = "6349483943:AAGEN47LBBQbnsdjeuelvsC7j5RpZR025Vc"

var keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Санкт-Петербург"),
	),
)

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "red:6379",
		Password: "",
		DB:       0,
	})

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

		switch update.Message.Text {
		case "/start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберете город")
			msg.ReplyMarkup = keyboard
			if _, err := bot.Send(msg); err != nil {

				log.Fatal(err)
			}
		case "Санкт-Петербург":
			text := ""
			day := time.Now().Day()
			key := fmt.Sprintf("spb:%v", day)
			val, err := rdb.Get(context.Background(), key).Result()
			if err != nil && !errors.Is(err, redis.Nil) {
				log.Fatal(err)
			} else {
				text = val
			}
			if text == "" {
				w, err := getWeather()
				if err != nil {
					log.Fatal(err)
				}
				for i := range w.Hourly.Time {
					text = text + fmt.Sprintln(w.Hourly.Time[i], " - ", w.Hourly.Temp[i], "°C")
				}
				err = rdb.Set(context.Background(), key, text, time.Hour).Err()
				if err != nil {
					log.Fatal(err)
				}
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			if _, err := bot.Send(msg); err != nil {

				log.Fatal(err)
			}

		}

	}
}
