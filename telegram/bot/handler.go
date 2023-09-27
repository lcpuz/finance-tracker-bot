package telegrambot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lcpuz/finance-tracker-bot/telegram/lang"
)

func (b *TelegramBot) HandleMessage(message *tgbotapi.Message) error {
	userID, err := b.repository.GetUserID(message.From.ID)
	if err != nil {
		return err
	}

	incomeCategoryState, err := b.repository.GetIncomeCategoryState(userID)
	if err != nil {
		return err
	}

	switch {
	case incomeCategoryState == 1:
		return b.HandleIncomeCategory(message)
	case message.Text == lang.Income[b.language]:
		return b.HandleIncome(message)
	case message.Text == lang.AddIncomeCategory[b.language]:
		return b.AddIncomeCategory(message)
	case message.Text == lang.BackToStartFromIncome[b.language]:
		return b.HandleBack(message)
	default:
		b.HandleUnknown(message)
	}

	return nil
}

func (b *TelegramBot) HandleUnknown(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, lang.UnknownMessage[b.language])
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (b *TelegramBot) HandleBack(message *tgbotapi.Message) error {
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

	msg := tgbotapi.NewMessage(message.Chat.ID, lang.ChooseAction[b.language])
	msg.ReplyMarkup = buttons
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
