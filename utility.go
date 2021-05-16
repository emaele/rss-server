package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/emaele/rss-telegram-notifier/entities"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func writeHTTPResponse(statusCode int, body string, writer http.ResponseWriter) {

	writer.WriteHeader(statusCode)
	_, err := writer.Write([]byte(body))
	if err != nil {
		log.Println(err)
	}
}

func createTelegramKeyboard(URL string) tg.InlineKeyboardMarkup {
	var keyboard tg.InlineKeyboardMarkup
	row := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonURL("🔗 Link", URL))
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)

	//We finally append the lower row to the keyboard
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard)

	return keyboard
}

func createTelegramMessage(element entities.RssItem) tg.MessageConfig {

	feedTitle := retrieveFeedTitle(element.Feed)

	text := fmt.Sprintf("*%s*\n\n%s", feedTitle, element.Title)
	text = strings.ReplaceAll(text, ".", "\\.")
	text = strings.ReplaceAll(text, "-", "\\-")

	message := tg.NewMessage(telegramChatID, text)
	message.ParseMode = "MarkdownV2"
	message.ReplyMarkup = createTelegramKeyboard(element.URL)

	return message
}
