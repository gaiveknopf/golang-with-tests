package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DatabasePlayer interface {
	GetPlayerScore(name string) int
	VictoryRegister(name string)
	GetLeagueTable() []Player
}

type Player struct {
	Name string
	Wins int
}

type PlayerServer struct {
	storage DatabasePlayer
	http.Handler
}

func NewPlayerServer(storage DatabasePlayer) *PlayerServer {
	p := new(PlayerServer)
	p.storage = storage
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	p.Handler = router
	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	err := json.NewEncoder(w).Encode(p.storage.GetLeagueTable())
	if err != nil {
		return
	}
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.victoryRegister(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.storage.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	_, err := fmt.Fprint(w, score)
	if err != nil {
		return
	}
}

func (p *PlayerServer) victoryRegister(w http.ResponseWriter, player string) {
	p.storage.VictoryRegister(player)
	w.WriteHeader(http.StatusAccepted)
}
