package main

import (
	"flag"
	"log"
	tgClient "telegram-bot/clients/telegram"
	event_consumer "telegram-bot/consumer/event-consumer"
	"telegram-bot/events/telegram"
	"telegram-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	storage := files.New(storagePath)

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		storage,
	)

	log.Println("Telegram bot started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("Telegram bot service is stopped:", err)
	}

}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token got access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is required")
	}

	return *token
}
