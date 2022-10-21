package telegram

import (
	"log"
	"mbu_vpngater_bot/vpngate"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	TOKEN        = "5445719929:AAH_eE5WE_Ak2W3FIioaTfqteREIqr7cZ6M"
	NO_COMMAND   = "🚫 داداش چی میگی نزن! این دستور اصلا وجود نداره اگر میخواهی به سازندش بگو. بای بای"
	FIND_COMMAND = "find"
)

func StartBot() {
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		panic(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if !update.Message.IsCommand() {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ReplyToMessageID = update.Message.MessageID

		switch update.Message.Command() {
		case FIND_COMMAND:
			for _, v := range vpngate.GetServerList() {
				msg.Text = v
				if _, err := bot.Send(msg); err != nil {
					log.Println(err)
				}
			}

		default:
			msg.Text = NO_COMMAND
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
		}
	}
}
