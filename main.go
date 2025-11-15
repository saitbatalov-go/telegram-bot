package main

import (
	"context"
	"flag"
	"log"

	_ "github.com/mattn/go-sqlite3"

	tgClient "telegram-bot/clients/telegram"
	eventConsumer "telegram-bot/consumer/event-consumer"
	"telegram-bot/events/telegram"
	"telegram-bot/storage/sqlite"
)

const (
	tgBotHost         = "api.telegram.org"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	//storage := files.New(storagePath)

	storage, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatalf("Error connect sqlite storage: %v", err)
	}

	err = storage.Init(context.TODO())
	if err != nil {
		log.Fatalf("Error initializing storage: %v", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		storage,
	)

	log.Println("Telegram bot started")

	consumer := eventConsumer.New(eventsProcessor, eventsProcessor, batchSize)

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
