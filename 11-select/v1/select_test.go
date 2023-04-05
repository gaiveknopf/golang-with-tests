package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	slowServer := createDelayedServer(20 * time.Millisecond)
	fastServer := createDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	URLSlow := slowServer.URL
	URLFast := fastServer.URL

	want := URLFast
	got := Racer(URLSlow, URLFast)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func createDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
