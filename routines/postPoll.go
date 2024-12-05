package routins

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func postPoll() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

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
