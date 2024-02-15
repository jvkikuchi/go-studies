package main

func Sum(array []int) int {
	var sum int = 0

	for _, number := range array {
		sum += number
	}

	return sum
}