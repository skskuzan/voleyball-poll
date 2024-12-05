package workers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
	"volleyball_bot/cache"
	"volleyball_bot/helpers"
)

func RunFetchUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := helpers.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.PollAnswer != nil {
			cache.C.AddPollAnswer(update.PollAnswer)
		}

		if update.Message != nil {
			cache.C.AddMessage(update.Message)
		}

		if update.Poll != nil {
			cache.C.AddPoll(update.Poll.ID, time.Now())
		}
	}

}
