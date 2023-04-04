package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 7, 8, 9}

		result := Sum(numbers)
		expect := 30

		if result != expect {
			t.Errorf("got %d want %d given, %v", result, expect, numbers)
		}
	})
}
