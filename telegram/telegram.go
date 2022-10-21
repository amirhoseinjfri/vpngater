package telegram

import (
	"log"
	"mbu_vpngater_bot/vpngate"
	"os"

	"github.com/yanzay/tbot/v2"
)

const (
	TOKEN        = "5445719929:AAH_eE5WE_Ak2W3FIioaTfqteREIqr7cZ6M"
	NO_COMMAND   = "🚫 داداش چی میگی نزن! این دستور اصلا وجود نداره اگر میخواهی به سازندش بگو. بای بای"
	FIND_COMMAND = "find"
)

func StartBot() {
	bot := tbot.New(TOKEN, tbot.WithWebhook("https://mbugate.herokuapp.com", ":"+os.Getenv("PORT")))
	c := bot.Client()
	bot.HandleMessage("/find", func(m *tbot.Message) {
		for _, v := range vpngate.GetServerList() {
			c.SendMessage(m.Chat.ID, v)
		}
	})
	bot.HandleMessage(".*.*", func(m *tbot.Message) {
		c.SendMessage(m.Chat.ID, NO_COMMAND)
	})
	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
