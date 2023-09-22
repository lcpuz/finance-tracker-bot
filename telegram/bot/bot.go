package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lcpuz/finance-tracker-bot/telegram/repository"
)

type TelegramBot struct {
	bot        *tgbotapi.BotAPI
	repository *repository.Repository
	language   string
}

func NewTelegramBot(bot *tgbotapi.BotAPI, repository *repository.Repository) *TelegramBot {
	return &TelegramBot{
		bot:        bot,
		repository: repository,
		language:   "en",
	}
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
