package api

import (
	"fmt"
	"os"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type BlindAlerterFunc func(duration time.Duration, amount int)

func (f BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	f(duration, amount)
}

func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		_, err := fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
		if err != nil {
			return
		}
	})
}
