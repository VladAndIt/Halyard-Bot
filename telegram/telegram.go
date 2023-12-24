package telegram

import (
	"context"
	"fmt"

	"halyard.bot/config"
	"halyard.bot/keyboards"
	"halyard.bot/log"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func startBot(ctx context.Context) {
	log := log.LoggerFromContext(ctx)
	config := config.ConfigFromContext(ctx)
	bot, err := tgapi.NewBotAPI(config.Telegram.Tocken)
	if err != nil {
		log.Error("Error")
	}

	log.Info("Authorized on account %s \n", bot.Self.UserName)

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
