package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/lcpuz/finance-tracker-bot/pkg"
	telegrambot "github.com/lcpuz/finance-tracker-bot/telegram/bot"
	"github.com/lcpuz/finance-tracker-bot/telegram/repository"
	_ "github.com/lib/pq"
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

	// Connect to db
	db, err := pkg.NewPostgresDB(pkg.PostgresConfig{
		Host:     os.Getenv("DBHOST"),
		Port:     os.Getenv("DBPORT"),
		Username: os.Getenv("DBUSERNAME"),
		DBName:   os.Getenv("DBNAME"),
		Password: os.Getenv("DBPASSWORD"),
		SSLMode:  os.Getenv("SSLMODE"),
	})

	//Initialize new repository
	repository := repository.NewRepository(db)

	//Check if there is an error while connecting to db
	if err != nil {
		logrus.Fatalf("Error occured while connecting to db: %s", err.Error())
	}

	//Set bot debug mode
	botToken.Debug = false

	//Create new telegram bot
	telegramBot := telegrambot.NewTelegramBot(botToken, repository)

	//Start telegram bot
	if err := telegramBot.StartTelegramBot(); err != nil {
		logrus.Fatalf("Error occured while starting telegram bot: %s", err.Error())
	}
}
