package cache

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sync"
	"time"
)

type Cache struct {
	Messages    []*tgbotapi.Message
	messageLock sync.Mutex

	Polls     map[string]time.Time
	pollsLock sync.Mutex

	PollAnswers     map[string][]*tgbotapi.PollAnswer
	pollAnswersLock sync.Mutex
}

var C *Cache

func init() {
	C = NewCache()
}

func NewCache() *Cache {
	return &Cache{
		Messages:    []*tgbotapi.Message{},
		Polls:       make(map[string]time.Time),
		PollAnswers: make(map[string][]*tgbotapi.PollAnswer),
	}
}

func (c *Cache) AddMessage(message *tgbotapi.Message) {
	c.messageLock.Lock()
	defer c.messageLock.Unlock()

	c.Messages = append(c.Messages, message)
	if len(c.Messages) > 100 {
		c.Messages = c.Messages[1:]
	}
}

func (c *Cache) GetMessagesByTimeframe(from, to time.Time) []*tgbotapi.Message {
	c.messageLock.Lock()
	defer c.messageLock.Unlock()

	var messagesInTimeframe []*tgbotapi.Message
	for _, msg := range c.Messages {
		msgTime := time.Unix(int64(msg.Date), 0)
		if msgTime.After(from) && msgTime.Before(to) {
			messagesInTimeframe = append(messagesInTimeframe, msg)
		}
	}
	return messagesInTimeframe
}

func (c *Cache) AddPoll(pollID string, created time.Time) {
	c.pollsLock.Lock()
	defer c.pollsLock.Unlock()

	c.Polls[pollID] = created
}

func (c *Cache) GetPoll(pollID string) time.Time {
	c.pollsLock.Lock()
	defer c.pollsLock.Unlock()

	return c.Polls[pollID]
}

func (c *Cache) AddPollAnswer(answer *tgbotapi.PollAnswer) {
	c.pollAnswersLock.Lock()
	defer c.pollAnswersLock.Unlock()

	// Initialize the slice if it doesn't exist
	if _, exists := c.Polls[answer.PollID]; !exists {
		c.AddPoll(answer.PollID, time.Now())
	}

	// Initialize the slice if it doesn't exist
	if _, exists := c.PollAnswers[answer.PollID]; !exists {
		c.PollAnswers[answer.PollID] = []*tgbotapi.PollAnswer{}
	}

	c.PollAnswers[answer.PollID] = append(c.PollAnswers[answer.PollID], answer)
	if len(c.PollAnswers) > 100 {
		deprecatedAnswers := c.GetPollAnswersByTimeframe(time.UnixMilli(0), time.Now().Add(-time.Hour*7500))
		for _, deprecatedAnswer := range deprecatedAnswers {
			delete(c.PollAnswers, deprecatedAnswer.PollID)
		}
	}
}

func (c *Cache) GetPollAnswersByTimeframe(from, to time.Time) []*tgbotapi.PollAnswer {
	c.pollAnswersLock.Lock()
	defer c.pollAnswersLock.Unlock()

	var answersInTimeframe []*tgbotapi.PollAnswer
	for pollID, answers := range c.PollAnswers {
		pollCreatedTime, exists := c.Polls[pollID]
		if !exists {
			continue
		}

		// Check if the poll was created within the timeframe
		if pollCreatedTime.After(from) && pollCreatedTime.Before(to) {
			answersInTimeframe = append(answersInTimeframe, answers...)
		}
	}
	return answersInTimeframe
}
