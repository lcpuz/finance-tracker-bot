package telegrambot

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lcpuz/finance-tracker-bot/telegram/lang"
	"github.com/lcpuz/finance-tracker-bot/telegram/request"
)

func (b *TelegramBot) HandleIncome(message *tgbotapi.Message) error {
	// Get User ID
	UserID, err := b.repository.GetUserID(message.From.ID)
	if err != nil {
		return err
	}

	// Get Categories
	Categories, err := b.repository.GetCategories(UserID)
	if err != nil {
		return err
	}

	// Create a keyboard for categories
	var keyboardRows [][]tgbotapi.KeyboardButton
	for _, category := range Categories {
		keyboardRows = append(keyboardRows, tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(category.CategoryName),
		))
	}

	// Add buttons for "Add Income" and "Back to Start"
	keyboardRows = append(keyboardRows, tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(lang.AddIncome[b.language]),
		tgbotapi.NewKeyboardButton(lang.AddIncomeCategory[b.language]),
	),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(lang.BackToStartFromIncome[b.language]),
		),
	)

	buttons := tgbotapi.NewReplyKeyboard(keyboardRows...)

	// Create a message
	var msgText string

	if len(Categories) == 0 {
		msgText = lang.NoCategories[b.language]
	} else {
		msgText = lang.IncomeCategoriesMessage[b.language]
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	msg.ReplyMarkup = buttons

	// Send the message
	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *TelegramBot) HandleAddIncomeCategory(message *tgbotapi.Message) error {
	UserID, err := b.repository.GetUserID(message.From.ID)
	if err != nil {
		return err
	}
	
	CreateCategoryRequest := request.CreateCategoryRequest{
		CategoryName:   message.Text,
		CategoryUserID: UserID,
	}

	// Add Category
	err = b.repository.CreateCategory(CreateCategoryRequest)
	if err != nil {
		//if err contains "duplicate key value violates unique constraint" then send message "Category already exists"
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			msg := tgbotapi.NewMessage(message.Chat.ID, lang.CategoryAlreadyExists[b.language])
			_, err = b.bot.Send(msg)
			if err != nil {
				return err
			}
			b.HandleIncome(message)
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, lang.FailAddIncomeCategory[b.language])
		_, err = b.bot.Send(msg)
		if err != nil {
			return err
		}

		b.HandleIncome(message)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, lang.SuccessAddIncomeCategory[b.language])
	// Send the message
	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}

	b.HandleIncome(message)
	return nil
}
