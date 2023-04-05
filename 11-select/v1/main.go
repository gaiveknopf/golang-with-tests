package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	durationA := MeasureResponseTime(a)
	durationB := MeasureResponseTime(b)

	if durationA < durationB {
		return a
	}

	return b
}

func MeasureResponseTime(url string) time.Duration {
	init := time.Now()
	_, err := http.Get(url)
	if err != nil {
		return time.Duration(0)
	}
	return time.Since(init)
}

func main() {
	print(Racer("https://www.google.com", "https://www.facebook.com"))
	fmt.Println()
	print(Racer("https://www.facebook.com", "https://www.google.com"))
}
