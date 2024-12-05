package helpers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

// SendMessage sends a message to a specific chat ID
func SendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message to chat %d: %v", chatID, err)
	}
	return err
}

func PostPoll(bot *tgbotapi.BotAPI, chatID int64) {
	poll := tgbotapi.SendPollConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: chatID, // Replace with your Polls channel ID
		},
		Question:              "Will you attend the upcoming training?",
		Options:               []string{"Yes", "No", "Maybe", "Will bring plus one"},
		IsAnonymous:           false,
		AllowsMultipleAnswers: true,
	}

	if _, err := bot.Send(poll); err != nil {
		log.Println(err)
	}
}
