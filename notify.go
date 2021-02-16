package main

import (
	"fmt"
	"log"
	"time"

	"github.com/emaele/rss-telegram-notifier/entities"
	tb "gopkg.in/tucnak/telebot.v2"
)

func notificationRoutine() {

	for range time.NewTicker(10 * time.Minute).C {
		var elements []entities.RssItem
		rows := db.Where("sent = ?", false).Find(&elements).RowsAffected

		if rows == 0 {
			continue
		}

		for _, element := range elements {
			message := fmt.Sprintf("%s \n %s", element.Title, element.URL)

			_, err := bot.Send(&tb.Chat{ID: telegramChatID}, message)
			if err != nil {
				log.Println(err)
				continue
			}

			// setting as sent
			db.Model(&element).Update("sent", true)
		}
	}
}
