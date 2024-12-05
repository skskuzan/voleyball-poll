package voleyball_poll

import (
	"github.com/robfig/cron/v3"
	"volleyball_bot/routines"
	"volleyball_bot/workers"
)

func main() {
	c := cron.New()

	c.AddFunc("0 19 * * TUE", routines.PostPoll) // Tuesday at 20:00 UKR, 19:00 FR
	c.AddFunc("0 19 * * THU", routines.PostPoll) // Thursday at 20:00 UKR, 19:00 FR

	c.AddFunc("0 0 * * TUE", routines.CheckPayments) // Tuesday at 01:00 UKR, 00:00 FR
	c.AddFunc("0 0 * * SUN", routines.CheckPayments) // Sunday at 01:00 UKR, 00:00 FR

	c.AddFunc("0 17 * * MON", routines.DivideAndPostTeams) // Monday at 18:00 UKR, 17:00 FR
	c.AddFunc("0 13 * * SAT", routines.DivideAndPostTeams) // Saturday at 14:00 UKR, 13:00 FR
	c.Start()

	workers.RunFetchUpdates()
}
