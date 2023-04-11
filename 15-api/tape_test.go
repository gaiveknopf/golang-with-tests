package api

import (
	"io"
	"testing"
)

func TestTape_Write(t *testing.T) {
	t.Run("should write to file", func(t *testing.T) {
		file, clean := createTempFile(t, "hello")
		defer clean()

		tape := &tape{file}

		_, err := tape.Write([]byte("world"))
		if err != nil {
			t.Fatalf("could not write to tape, %v", err)
		}
		_, err = file.Seek(0, 0)
		if err != nil {
			t.Fatalf("could not seek to beginning of file, %v", err)
		}
		newContent, _ := io.ReadAll(file)

		got := string(newContent)
		want := "world"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
