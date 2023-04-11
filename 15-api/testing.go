package api

import (
	"fmt"
	"testing"
	"time"
)

type StubStoragePlayer struct {
	score    map[string]int
	winCalls []string
	league   []Player
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}

func (s *StubStoragePlayer) GetPlayerScore(name string) int {
	score := s.score[name]
	return score
}

func (s *StubStoragePlayer) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubStoragePlayer) GetLeague() League {
	return s.league
}

func assertScheduledAlert(t *testing.T, got, want scheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got scheduled alert %v, want %v", got, want)
	}
}

func CheckPlayerWin(t *testing.T, store *StubStoragePlayer, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}

	got := store.winCalls[0]
	if got != winner {
		t.Errorf("didn't store correct winner got '%s', want '%s'", got, winner)
	}
}
