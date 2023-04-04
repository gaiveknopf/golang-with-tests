package main

import "testing"

func TestHello(t *testing.T) {
	checkRightMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("Result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("Should say hello to people", func(t *testing.T) {
		result := Hello("Charmaine", "")
		expected := "Hello, Charmaine!"
		checkRightMessage(t, result, expected)
	})

	t.Run("Should say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, world!"
		checkRightMessage(t, result, expected)
	})

	t.Run("Should say Hola, Rosita!", func(t *testing.T) {
		result := Hello("Rosita", "Spanish")
		expected := "Hola, Rosita!"
		checkRightMessage(t, result, expected)
	})

	t.Run("Should say Bonjour Pietra", func(t *testing.T) {
		result := Hello("Pietra", "French")
		expected := "Bonjour, Pietra!"
		checkRightMessage(t, result, expected)
	})
}
