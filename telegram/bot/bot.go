package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lcpuz/finance-tracker-bot/telegram/repository"
	"github.com/lcpuz/finance-tracker-bot/telegram/request"
)

type TelegramBot struct {
	bot                 *tgbotapi.BotAPI
	repository          *repository.Repository
	language            string
	UncategorizedIncome map[int64]*request.AddUncategorizedIncomeRequest
}

func NewTelegramBot(bot *tgbotapi.BotAPI, repository *repository.Repository) *TelegramBot {
	return &TelegramBot{
		bot:                 bot,
		repository:          repository,
		language:            "en",
		UncategorizedIncome: make(map[int64]*request.AddUncategorizedIncomeRequest),
	}
}

func (b *TelegramBot) StartTelegramBot() error {
	updates, err := b.InitUpdatesChannel()
	if err != nil {
		return err
	}

	b.HandleUpdates(updates)

	//change language

	return nil
}

func (b *TelegramBot) InitUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u), nil
}

func (b *TelegramBot) HandleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		b.HandleMessage(update.Message)
	}
}
