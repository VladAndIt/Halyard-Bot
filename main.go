package main

import (
	"fmt"
	"log"

	mainconfig "halyard.bot/config"
	keyboards "halyard.bot/keyboards"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	var cfg mainconfig.Config
	mainconfig.MustLoad(&cfg)

	bot, err := tgapi.NewBotAPI(cfg.Telegram.Tocken)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s \n", bot.Self.UserName)

	u := tgapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	fmt.Println("--------------------------------------------")
	for update := range updates {
		if update.CallbackQuery != nil {
			fmt.Println(update)

			tgapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

			bot.Send(tgapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data))
		}
		if update.Message != nil {

			msg := tgapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			log.Println(msg)
			switch update.Message.Text {
			case "/start":
				{
					msg.ReplyMarkup = keyboards.MainKeyboard
				}
			case "/open":
				{
					msg.ReplyMarkup = keyboards.InlineKeboard
				}
				bot.Send(msg)
			}
		}
	}
}
