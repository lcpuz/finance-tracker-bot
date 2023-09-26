package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lcpuz/finance-tracker-bot/telegram/lang"
)

func (b *TelegramBot) HandleIncome(message *tgbotapi.Message) error {
	//Get User ID
	UserID, err := b.repository.GetUserID(message.From.ID)
	if err != nil {
		return err
	}

	//Get Categories
	Categories, err := b.repository.GetCategories(UserID)
	if err != nil {
		return err
	}

	//Make buttons
	if len(Categories) == 0 {
		buttons := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(lang.AddIncomeCategory[b.language]),
				tgbotapi.NewKeyboardButton(lang.AddIncome[b.language]),
			),
		)
		msg := tgbotapi.NewMessage(message.Chat.ID, lang.NoCategories[b.language])
		msg.ReplyMarkup = buttons
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}
		return nil
	}

	// If categories is greater than 4 then make 2 rows
	if len(Categories) > 4 {
		buttons := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(Categories[0].CategoryName),
				tgbotapi.NewKeyboardButton(Categories[1].CategoryName),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(Categories[2].CategoryName),
				tgbotapi.NewKeyboardButton(Categories[3].CategoryName),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(lang.AddIncome[b.language]),
			),
		)
		msg := tgbotapi.NewMessage(message.Chat.ID, lang.IncomeCategoriesMessage[b.language])
		msg.ReplyMarkup = buttons
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}
