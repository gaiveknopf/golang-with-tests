package main

import (
	"fmt"
	"net/http"
	"time"
)

var limitTime = 10 * time.Second

func Runner(a, b string) (winner string, err error) {
	return Configurable(a, b, limitTime)
}

func Configurable(a, b string, limitTime time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(limitTime):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		_, err := http.Get(url)
		if err != nil {
			return
		}
		ch <- true
	}()
	return ch
}

func main() {
	//winner, _ := Runner("https://www.databoxsistemas.com.br", "ge.com")
	winner, _ := Runner("http://www.google.com", "http://www.facebook.com")
	fmt.Println(winner)
}
