package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("should add 3 times and the result will be 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Add()
		counter.Add()
		counter.Add()

		checkCounter(t, counter, 3)
	})

	t.Run("should runs concurrently safely", func(t *testing.T) {
		countExpected := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(countExpected)

		for i := 0; i < countExpected; i++ {
			go func(w *sync.WaitGroup) {
				counter.Add()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		checkCounter(t, counter, countExpected)
	})
}

func checkCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got: %d, want: %d", got.Value(), want)
	}
}
