package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type SleeperDefault struct{}

func (d *SleeperDefault) Sleep() {
	time.Sleep(1 * time.Second)
}

type SleeperConfigurable struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *SleeperConfigurable) Sleep() {
	c.sleep(c.duration)
}

func Counter(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		_, err := writer.Write([]byte(fmt.Sprintf("%d\n", i)))
		if err != nil {
			panic(err)
		}
	}

	sleeper.Sleep()
	_, err := writer.Write([]byte(finalWord))
	if err != nil {
		panic(err)
	}
}

func main() {
	sleeper := &SleeperConfigurable{duration: 1 * time.Second, sleep: time.Sleep}
	Counter(os.Stdout, sleeper)
}
