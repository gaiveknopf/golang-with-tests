package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllRest(numbers ...[]int) (sums []int) {
	for _, number := range numbers {
		if len(number) == 0 {
			sums = append(sums, 0)
			continue
		}
		final := number[1:]
		sums = append(sums, Sum(final))
	}
	return
}
