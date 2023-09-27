package telegrambot

import (
	"fmt"
	"log"
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

func (b *TelegramBot) AddIncomeCategory(message *tgbotapi.Message) error {
	//Remove the keyboard
	RemoveKeyboard := tgbotapi.NewRemoveKeyboard(true)

	UserID, err := b.repository.GetUserID(message.From.ID)
	if err != nil {
		return err
	}

	err = b.repository.IncomeCategoryStateRepository.UpdateIncomeCategoryState(UserID, 1)
	if err != nil {
		fmt.Println("error in CreateIncomeCategoryState")
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, lang.AddIncomeCategoryText[b.language])
	msg.ReplyMarkup = RemoveKeyboard
	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *TelegramBot) HandleIncomeCategory(message *tgbotapi.Message) error {
	// Get User ID
	UserID, err := b.repository.GetUserID(message.From.ID)
	if err != nil {
		return err
	}

	// Create a category
	CreateCategoryRequest := request.CreateCategoryRequest{
		CategoryName:   message.Text,
		CategoryUserID: UserID,
	}

	// Attempt to create the category
	err = b.repository.IncomeCategoryRepository.CreateCategory(CreateCategoryRequest)
	if err != nil {
		// If category already exists - log and send a message from lang.CategoryAlreadyExists
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			log.Printf("Category already exists: %v", err)
			msg := tgbotapi.NewMessage(message.Chat.ID, lang.CategoryAlreadyExists[b.language])
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Println("Error sending message:", err)
			}
		} else {
			// If there is another error - log and send a message from lang.FailAddIncomeCategory
			log.Printf("Error adding category: %v", err)
			msg := tgbotapi.NewMessage(message.Chat.ID, lang.FailAddIncomeCategory[b.language])
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Println("Error sending message:", err)
			}
		}
	} else {
		// Category added successfully - log and send a message from lang.SuccessAddIncomeCategory
		log.Println("Category added successfully.")
		msg := tgbotapi.NewMessage(message.Chat.ID, lang.SuccessAddIncomeCategory[b.language])
		_, err = b.bot.Send(msg)
		if err != nil {
			log.Println("Error sending message:", err)
		}

		// Delete the state
		err = b.repository.IncomeCategoryStateRepository.DeleteIncomeCategoryState(UserID)
		if err != nil {
			log.Println("Error deleting category state:", err)
		}
	}

	// Handle income after category is added or if an error occurred
	return b.HandleIncome(message)
}
