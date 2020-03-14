package codewar

import "math"

// Complete the square sum function so that it squares each number passed into it and then sums the results together.
//
// For example, for [1, 2, 2] it should return 9 because 1^2 + 2^2 + 2^2 = 9
func SquareSum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += int(math.Pow(float64(number), float64(2)))
	}
	return sum
}
