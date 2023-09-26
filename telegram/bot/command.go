package telegrambot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lcpuz/finance-tracker-bot/telegram/lang"
	"github.com/lcpuz/finance-tracker-bot/telegram/request"
)

const (
	StartCommand = "start"
	HelpCommand  = "help"
	AboutCommand = "about"
)

func (b *TelegramBot) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case StartCommand:
		b.HandleStartCommand(message)
	case HelpCommand:
		b.HandleHelpCommand(message)
	case AboutCommand:
		b.HandleAboutCommand(message)
	}
}

func (b *TelegramBot) HandleStartCommand(message *tgbotapi.Message) {

	//Create user
	UserRequest := request.CreateUsersRequest{
		UserChatID:   message.From.ID,
		UserName:     &message.From.UserName,
		UserNickName: message.From.FirstName,
	}

	err := b.repository.CreateUser(UserRequest)
	//if error contains "duplicate key value violates unique constraint" then user already exists
	if err != nil && err.Error() != "duplicate key value violates unique constraint " {
		log.Println("User exists")
	} else {
		log.Println("User created")
	}

	//get user language
	language := message.From.LanguageCode

	//check if language is supported
	if language == "uz" || language == "ru" || language == "en" || language == "ar" {
		b.language = language
	}

	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(lang.Income[b.language]),
			tgbotapi.NewKeyboardButton(lang.Spendings[b.language]),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(lang.Report[b.language]),
			tgbotapi.NewKeyboardButton(lang.Settings[b.language]),
		),
	)
	//Make message
	msg := tgbotapi.NewMessage(message.Chat.ID, lang.WelcomeMessage[b.language])
	buttons.ResizeKeyboard = true
	msg.ReplyMarkup = buttons

	//Send message
	_, err = b.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (b *TelegramBot) HandleHelpCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Yordam")
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (b *TelegramBot) HandleAboutCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Biz haqimizda")
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
