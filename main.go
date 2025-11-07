package main

import (
	"flag"
	"log"
	"telegram-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())

}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token got access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is required")
	}

	return *token
}
