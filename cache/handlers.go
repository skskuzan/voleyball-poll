package cache

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	ID       int64
	UserName string
	FullName string
}

func HandlePaymentsMessages(from time.Time, to time.Time) map[string]int {
	var (
		payersLock sync.Mutex
		payers     = make(map[string]int) // Key: Username, Value: User struct
	)

	for _, message := range C.GetMessagesByTimeframe(from, to) {
		payersLock.Lock()
		if _, ok := payers[message.From.UserName]; !ok {
			payers[message.From.UserName] = 1
		} else {
			payers[message.From.UserName]++
		}

		payersLock.Unlock()

	}

	return payers

}

// Handler for poll answers
func HandlePollAnswers(from time.Time, to time.Time) map[string]User {
	var (
		participantLock sync.Mutex
		participants    = make(map[string]User) // Key: Username, Value: User struct
	)

	for _, pollAnswer := range C.GetPollAnswersByTimeframe(from, to) {

		userID := pollAnswer.User.ID
		userName := pollAnswer.User.UserName

		// Fetch the poll to get options
		// (Assuming you have stored the poll ID and options mapping)
		pollOptions := []string{"Yes", "No", "Maybe", "Will bring plus one"}

		selectedOptions := pollAnswer.OptionIDs
		for _, optionID := range selectedOptions {
			selectedOption := pollOptions[optionID]

			if selectedOption == "Yes" {
				participantLock.Lock()
				participants[userName] = User{
					ID:       userID,
					UserName: userName,
					FullName: fmt.Sprint(pollAnswer.User.FirstName, " ", pollAnswer.User.LastName),
				}
				participantLock.Unlock()
				break
			}

			if selectedOption == "Will bring plus one" {
				participantLock.Lock()
				participants[userName] = User{
					ID:       userID,
					UserName: userName,
					FullName: fmt.Sprint("Плюс: ", pollAnswer.User.FirstName, " ", pollAnswer.User.LastName),
				}
				participantLock.Unlock()
				break
			}
		}
	}

	return participants
}
