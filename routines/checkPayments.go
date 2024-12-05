package routines

import (
	"fmt"
	"time"
	"volleyball_bot/cache"
	"volleyball_bot/helpers"
)

func CheckPayments() {
	var participantsFrom, participantsTo, payersFrom, payersTo time.Time
	if time.Now().Weekday() == time.Tuesday {
		payersFrom = time.Now().Add(-time.Hour * 6)
		payersTo = time.Now()
	} else if time.Now().Weekday() == time.Sunday {
		payersFrom = time.Now().Add(-time.Hour * 10)
		payersTo = time.Now()
	}

	participantsFrom = time.Now().Add(-time.Hour * 24 * 5)
	participantsTo = time.Now().Add(-time.Hour * 24 * 4)

	participants := cache.HandlePollAnswers(participantsFrom, participantsTo)
	payers := cache.HandlePaymentsMessages(payersFrom, payersTo)

	// Compare and send reminders
	for _, user := range participants {
		if _, ok := payers[user.UserName]; !ok {
			_ = helpers.SendMessage(chatID, fmt.Sprintf("@%s, please pay!", user.UserName))
		} else {
			payers[user.UserName]--
			if payers[user.UserName] < 0 {
				_ = helpers.SendMessage(chatID, fmt.Sprintf("@%s, please pay!", user.UserName))
			}
		}
	}
}
