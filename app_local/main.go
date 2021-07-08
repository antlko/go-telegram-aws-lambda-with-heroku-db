package main

import (
	"github.com/joho/godotenv"
	"github.com/tbot/s_tgapp/botapp"
	"github.com/tbot/s_tgapp/database"
	"log"
)

func init() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env.local file")
	}
}

func main() {
	tgBot, err := botapp.CreateBot(database.NewPostgresDB())
	if err != nil {
		log.Fatal(tgBot)
	}
	tgBot.Start()
}
