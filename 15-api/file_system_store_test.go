package api

import (
	"os"
	"testing"
)

func TestFileDBPlayer(t *testing.T) {
	t.Run("should returns ordered league data from reader", func(t *testing.T) {
		database, clearDatabase := createTempFile(t, `[
			{"Name": "Maria", "Wins": 10},
			{"Name": "Pedro", "Wins": 20}
		]`)
		defer clearDatabase()

		storage, err := NewFileSystemPlayerStore(database)
		checkNoError(t, err)

		got := storage.GetLeague()

		want := []Player{
			{"Pedro", 20},
			{"Maria", 10},
		}

		checkLeague(t, got, want)
	})

	t.Run("should return player score", func(t *testing.T) {
		database, clearDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
		defer clearDatabase()

		storage, err := NewFileSystemPlayerStore(database)
		checkNoError(t, err)

		got := storage.GetPlayerScore("Chris")
		want := 33
		checkEqualScore(t, got, want)
	})

	t.Run("should save victory of existing player", func(t *testing.T) {
		database, clearDatabase := createTempFile(t, `[
			{"Name": "Maria", "Wins": 10},
			{"Name": "Pedro", "Wins": 20}
		]`)
		defer clearDatabase()

		storage, err := NewFileSystemPlayerStore(database)
		checkNoError(t, err)

		storage.RecordWin("Maria")

		got := storage.GetPlayerScore("Maria")
		want := 11

		checkEqualScore(t, got, want)
	})

	t.Run("should save victory of new player", func(t *testing.T) {
		database, clearDatabase := createTempFile(t, `[
			{"Name": "Maria", "Wins": 10},
			{"Name": "Pedro", "Wins": 20}
		]`)
		defer clearDatabase()

		storage, err := NewFileSystemPlayerStore(database)
		checkNoError(t, err)

		storage.RecordWin("Cleo")

		got := storage.GetPlayerScore("Cleo")
		want := 1

		checkEqualScore(t, got, want)
	})
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tempFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file, %v", err)
	}

	_, err = tempFile.Write([]byte(initialData))
	if err != nil {
		return nil, nil
	}

	removeFile := func() {
		err := tempFile.Close()
		if err != nil {
			return
		}
		err = os.Remove(tempFile.Name())
		if err != nil {
			return
		}
	}

	return tempFile, removeFile
}

func checkEqualScore(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func checkNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
