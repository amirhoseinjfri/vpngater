package main

import (
	"log"
	"mbu_vpngater_bot/telegram"
)

func main() {
	log.Println("mbu gate has been started...")
	telegram.StartBot()
}
