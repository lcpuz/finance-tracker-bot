package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	telegrambot "github.com/lcpuz/finance-tracker-bot/telegram/bot"
	"github.com/sirupsen/logrus"
)

func main() {
	//Set logrus to JSON formatter
	logrus.SetFormatter(&logrus.JSONFormatter{})

	//Load environment variables
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error occured while loading environment variables: %s", err.Error())
	}

	//Bot token
	botToken, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		logrus.Fatalf("Error occured while creating bot: %s", err.Error())
	}

	//Set bot debug mode
	botToken.Debug = true

	//Create new telegram bot
	telegramBot := telegrambot.NewTelegramBot(botToken)

	//Start telegram bot
	if err := telegramBot.StartTelegramBot(); err != nil {
		logrus.Fatalf("Error occured while starting telegram bot: %s", err.Error())
	}
}
