package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	app := NewFiberServer()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", 200, resp.StatusCode)
	}

	want := "Hello World!"
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Errorf("Expected response body to be '%s'. Got '%s'\n", want, string(got))
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	app := NewFiberServer()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	for i := 0; i < b.N; i++ {
		_, err := app.Test(req)
		if err != nil {
			return
		}
	}
}
