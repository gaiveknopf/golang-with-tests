package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	e := NewEchoServer()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, rec.Code)
	}

	if rec.Body.String() != "Hello World!" {
		t.Errorf("Expected response body to be 'Hello World!'. Got '%s'\n", rec.Body.String())
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	e := NewEchoServer()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		e.ServeHTTP(rec, req)
	}
}
