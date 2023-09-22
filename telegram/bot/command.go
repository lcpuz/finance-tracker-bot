package telegrambot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lcpuz/finance-tracker-bot/telegram/lang"
)

const (
	StartCommand = "start"
	HelpCommand  = "help"
	AboutCommand = "about"
)

func (b *TelegramBot) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case StartCommand:
		b.handleStart(message)
	case HelpCommand:
		b.handleHelp(message)
	case AboutCommand:
		b.handleAbout(message)
	}
}

func (b *TelegramBot) handleStart(message *tgbotapi.Message) {
	//get user language
	language := message.From.LanguageCode
	b.language = language

	msg := tgbotapi.NewMessage(message.Chat.ID, lang.WelcomeMessage[b.language])
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (b *TelegramBot) handleHelp(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Yordam")
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (b *TelegramBot) handleAbout(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Biz haqimizda")
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
