package telegrambot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lcpuz/finance-tracker-bot/telegram/lang"
)

func (b *TelegramBot) HandleMessage(message *tgbotapi.Message) error {
	UserId, err := b.repository.GetUserID(message.From.ID)
	if err != nil {
		return err
	}

	GetIncomeCategoryState, err := b.repository.GetIncomeCategoryState(UserId)
	if err != nil {
		return err
	}

	if GetIncomeCategoryState == 1 {
		err := b.HandleIncomeCategory(message)
		if err != nil {
			log.Println(err)
		}
	} else {
		switch message.Text {
		case lang.Income[b.language]:
			err := b.HandleIncome(message)
			if err != nil {
				log.Println(err)
			}
		case lang.AddIncomeCategory[b.language]:
			err := b.AddIncomeCategory(message)
			if err != nil {
				log.Println(err)
			}
		case lang.BackToStartFromIncome[b.language]:
			err := b.HandleBack(message)
			if err != nil {
				log.Println(err)
			}
		default:
			b.HandleUnknown(message)
		}
	}

	return nil
}

func (b *TelegramBot) HandleUnknown(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, lang.UnknownMessage[b.language])
	_, err := b.bot.Send(msg)
	if err != nil {
		return
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
