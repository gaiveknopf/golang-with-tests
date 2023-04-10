package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubStoragePlayer struct {
	scores           map[string]int
	victoryRegisters []string
	league           []Player
}

func (s *StubStoragePlayer) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubStoragePlayer) VictoryRegister(name string) {
	s.victoryRegisters = append(s.victoryRegisters, name)
}

func (s *StubStoragePlayer) GetLeagueTable() []Player {
	return s.league
}

func TestGetPlayers(t *testing.T) {
	storage := StubStoragePlayer{
		map[string]int{
			"Maria": 20,
			"Pedro": 10,
		},
		nil, nil}

	server := NewPlayerServer(&storage)

	t.Run("should returns Maria's score", func(t *testing.T) {
		request := newRequestGetPoints("Maria")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		checkRequestBody(t, response.Body.String(), "20")
	})

	t.Run("should returns Pedro's score", func(t *testing.T) {
		request := newRequestGetPoints("Pedro")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		checkRequestBody(t, response.Body.String(), "10")
	})
}

func TestVictoryRegister(t *testing.T) {
	storage := StubStoragePlayer{
		map[string]int{},
		nil, nil,
	}
	server := NewPlayerServer(&storage)

	t.Run("should record victory in HTTP POST method", func(t *testing.T) {
		player := "Maria"

		request := newRequestVictoryRegisterPost(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		checkResponseStatusCode(t, response.Code, http.StatusAccepted)

		if len(storage.victoryRegisters) != 1 {
			t.Fatalf("checked %d call on VictoryRegister, want %d", len(storage.victoryRegisters), 1)
		}

		if storage.victoryRegisters[0] != player {
			t.Errorf("did not register the winner correctly, got '%s', want '%s'", storage.victoryRegisters[0], player)
		}
	})
}

func newRequestVictoryRegisterPost(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newRequestGetPoints(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newRequestGetLeagueTable() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func checkRequestBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("invalid request body, got '%s' want '%s'", got, want)
	}
}

func checkResponseStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not receive the expected HTTP status code, got %d, want %d", got, want)
	}
}

func TestLeague(t *testing.T) {
	storage := StubStoragePlayer{}
	server := NewPlayerServer(&storage)

	t.Run("should returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []Player

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}

		checkResponseStatusCode(t, response.Code, http.StatusOK)
	})

	t.Run("should returns the league table as JSON", func(t *testing.T) {
		want := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		storage := StubStoragePlayer{nil, nil, want}
		server := NewPlayerServer(&storage)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []Player

		got = getResponseLeague(t, response.Body)
		checkResponseStatusCode(t, response.Code, http.StatusOK)
		checkLeague(t, got, want)
		checkContentType(t, response, jsonContentType)
	})
}

const jsonContentType = "application/json"

func checkContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
	}
}

func getResponseLeague(t *testing.T, body io.Reader) (got []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&got)
	if err != nil {
		t.Fatalf("unable to parse response from server %q into slice of Player, '%v'", body, err)
	}
	return got
}

func checkLeague(t *testing.T, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}
