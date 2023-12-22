package main

import (
	"fmt"
	"log"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgapi.NewInlineKeyboardMarkup(
	tgapi.NewInlineKeyboardRow(
		tgapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgapi.NewInlineKeyboardButtonSwitch("2sw", "open 2"),
		tgapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgapi.NewInlineKeyboardRow(
		tgapi.NewInlineKeyboardButtonData("4", "4"),
		tgapi.NewInlineKeyboardButtonData("5", "5"),
		tgapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

func main() {
	bot, err := tgapi.NewBotAPI("6728338149:AAFvWf0zpe_D2aRs2robHzBLMYsNm8YdO3k")
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true

	log.Printf("Authorized on account %s \n", bot.Self.UserName)

	u := tgapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	fmt.Println("--------------------------------------------")
	for update := range updates {
		if update.CallbackQuery != nil {
			fmt.Println(update)

			bot.Send(tgapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data))
		}
		if update.Message != nil {

			msg := tgapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			fmt.Println(msg)
			switch update.Message.Text {
			case "/start":
				{
					msg.ReplyMarkup = numericKeyboard
				}
			case "/open":
				{
					msg.ReplyMarkup = numericKeyboard
				}
			}

			bot.Send(msg)
		}
	}
}
