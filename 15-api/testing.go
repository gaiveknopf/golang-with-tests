package api

import (
	"fmt"
	"testing"
	"time"
)

type StubPlayerStore struct {
	score    map[string]int
	winCalls []string
	league   []Player
}

type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduledAlert{duration, amount})
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.score[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func assertScheduledAlert(t *testing.T, got, want ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got scheduled alert %v, want %v", got, want)
	}
}

func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}

	got := store.winCalls[0]
	if got != winner {
		t.Errorf("didn't store correct winner got '%s', want '%s'", got, winner)
	}
}
