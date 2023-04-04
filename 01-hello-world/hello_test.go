package main

import "testing"

func TestHelloPassingAName(t *testing.T) {
	result := Hello("Charmaine")
	expected := "Hello, Charmaine!"
	if result != expected {
		t.Errorf("Result '%s', expected '%s'", result, expected)
	}
}
