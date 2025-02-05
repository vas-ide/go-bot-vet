package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	routerPkg "github.com/vas-ide/go-bot-vet/internal/app/router"

)

func main() {
	_ = godotenv.Load()

	token, found := os.LookupEnv("TELEGRAM_APITOKEN")
	if !found {
		log.Panic("environment variable TOKEN not found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// Uncomment if you want debugging
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	

	routerHandler := routerPkg.NewRouter(bot)

	for update := range updates {
		routerHandler.HandleUpdate(update)
	}
}
