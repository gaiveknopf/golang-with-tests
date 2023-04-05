package _1_io_writer

import (
	"io"
	"os"
)

func main() {
	w := io.Writer(os.Stdout)
	_, err := w.Write([]byte("Hello, world"))
	if err != nil {
		return
	}

	_, err = w.Write([]byte("\n"))
	if err != nil {
		return
	}

	red := "\033[31m"

	blue := "\033[34m"

	reset := "\033[0m"

	_, err = w.Write([]byte(red))
	if err != nil {
		return
	}

	_, err = w.Write([]byte("Hello, world"))
	if err != nil {
		return
	}

	_, err = w.Write([]byte("\n"))
	if err != nil {
		return
	}

	_, err = w.Write([]byte(blue))
	if err != nil {
		return
	}

	_, err = w.Write([]byte("Hello, world"))
	if err != nil {
		return
	}

	_, err = w.Write([]byte("\n"))
	if err != nil {
		return
	}

	_, err = w.Write([]byte(reset))
	if err != nil {
		return
	}

	_, err = w.Write([]byte("Hello, world"))
	if err != nil {
		return
	}
}
