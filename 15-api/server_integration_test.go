package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterVictoriesAndGetThisVictories(t *testing.T) {
	storage := NewMemoryPlayerStorage()
	server := NewPlayerServer(storage)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newRequestVictoryRegisterPost(player))
	server.ServeHTTP(httptest.NewRecorder(), newRequestVictoryRegisterPost(player))
	server.ServeHTTP(httptest.NewRecorder(), newRequestVictoryRegisterPost(player))

	t.Run("should returns score 3", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newRequestGetPoints(player))
		checkResponseStatusCode(t, response.Code, http.StatusOK)
		checkRequestBody(t, response.Body.String(), "3")
	})

	t.Run("should returns league table", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newRequestGetLeagueTable())
		checkResponseStatusCode(t, response.Code, http.StatusOK)

		got := getResponseLeague(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}
		checkLeague(t, got, want)
	})

}
