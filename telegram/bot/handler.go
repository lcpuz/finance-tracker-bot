package telegrambot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lcpuz/finance-tracker-bot/telegram/lang"
)

func (b *TelegramBot) handleMessage(message *tgbotapi.Message) {
	switch message.Text {
	case lang.Income[b.language]:
		err := b.HandleIncome(message)
		if err != nil {
			log.Println(err)
		}
	default:
		b.HandleUnknown(message)
	}
}

func (b *TelegramBot) HandleUnknown(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, lang.UnknownMessage[b.language])
	_, err := b.bot.Send(msg)
	if err != nil {
		return
	}
}
