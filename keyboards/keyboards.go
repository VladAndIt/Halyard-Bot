package keyboards

import tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var MainKeyboard = tgapi.NewReplyKeyboard(
	tgapi.NewKeyboardButtonRow(
		tgapi.NewKeyboardButton("1"),
		tgapi.NewKeyboardButton("2"),
		tgapi.NewKeyboardButton("3"),
	),
	tgapi.NewKeyboardButtonRow(
		tgapi.NewKeyboardButton("4"),
		tgapi.NewKeyboardButton("5"),
		tgapi.NewKeyboardButton("6"),
	),
)

var InlineKeboard = tgapi.NewInlineKeyboardMarkup(
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
