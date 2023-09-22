package telegrambot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	bot      *tgbotapi.BotAPI
	language string
}

func NewTelegramBot(bot *tgbotapi.BotAPI) *TelegramBot {
	return &TelegramBot{bot: bot}
}

func (b *TelegramBot) StartTelegramBot() error {
	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(updates)

	//change language

	return nil
}

func (b *TelegramBot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u), nil
}

func (b *TelegramBot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		b.handleMessage(update.Message)
	}
}

func (b *TelegramBot) handleMessage(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Salom!")
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
