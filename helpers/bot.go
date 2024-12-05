package helpers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

var Bot *tgbotapi.BotAPI

func init() {
	var err error
	Bot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
}

// SendMessage sends a message to a specific chat ID
func SendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := Bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message to chat %d: %v", chatID, err)
	}
	return err
}

func PostPoll(chatID int64) {
	poll := tgbotapi.SendPollConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: chatID, // Replace with your Polls channel ID
		},
		Question:              "Will you attend the upcoming training?",
		Options:               []string{"Yes", "No", "Maybe", "Will bring plus one"},
		IsAnonymous:           false,
		AllowsMultipleAnswers: true,
	}

	if _, err := Bot.Send(poll); err != nil {
		log.Println(err)
	}
}
