package api

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

var dummySpyAlerter = &SpyBlindAlerter{}

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStorage := &StubStoragePlayer{}

		cli := NewCLI(playerStorage, in, nil)
		cli.PlayPoker()

		CheckPlayerWin(t, playerStorage, "Chris")
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStorage := &StubStoragePlayer{}

		cli := NewCLI(playerStorage, in, nil)
		cli.PlayPoker()

		CheckPlayerWin(t, playerStorage, "Cleo")
	})

	t.Run("should schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStorage := &StubStoragePlayer{}
		blindAlerter := &SpyBlindAlerter{}

		cli := NewCLI(playerStorage, in, blindAlerter)
		cli.PlayPoker()

		cases := []scheduledAlert{
			{0, 100},
			{10 * time.Second, 200},
			{20 * time.Second, 300},
			{30 * time.Second, 400},
			{40 * time.Second, 500},
			{50 * time.Second, 600},
			{60 * time.Second, 800},
			{70 * time.Second, 1000},
			{80 * time.Second, 2000},
			{90 * time.Second, 4000},
			{100 * time.Second, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}
				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})
}
