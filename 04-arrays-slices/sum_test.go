package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 7, 8, 9}

		result := Sum(numbers)
		expect := 30

		if result != expect {
			t.Errorf("result %d expect %d given, %v", result, expect, numbers)
		}
	})
}

func TestSumAllRest(t *testing.T) {
	checkSum := func(t *testing.T, result, expect []int, numbers [][]int) {
		t.Helper()
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("result %d expect %d given %v", result, expect, numbers)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		numbers := [][]int{[]int{1, 2}, []int{0, 9}}
		result := SumAllRest(numbers...)
		expect := []int{2, 9}

		checkSum(t, result, expect, numbers)
	})

	t.Run("safely sum empty slice", func(t *testing.T) {
		numbers := [][]int{[]int{}, []int{3, 4, 5}}
		result := SumAllRest(numbers...)
		expect := []int{0, 9}

		checkSum(t, result, expect, numbers)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		numbers := [][]int{[]int{}, []int{}}
		result := SumAllRest(numbers...)
		expect := []int{0, 0}

		checkSum(t, result, expect, numbers)
	})
}
