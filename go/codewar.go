package codewar

import (
	"math"
	"strconv"
)

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

// Some numbers have funny properties. For example:
//
// 89 --> 8¹ + 9² = 89 * 1
//
// 695 --> 6² + 9³ + 5⁴= 1390 = 695 * 2
//
// 46288 --> 4³ + 6⁴+ 2⁵ + 8⁶ + 8⁷ = 2360688 = 46288 * 51
//
// Given a positive integer n written as abcd... (a, b, c, d... being digits) and a positive integer p
//
// we want to find a positive integer k, if it exists, such as the sum of the digits of n taken to the successive powers of p is equal to k * n.
// In other words:
//
// Is there an integer k such as : (a ^ p + b ^ (p+1) + c ^(p+2) + d ^ (p+3) + ...) = n * k
//
// If it is the case we will return k, if not return -1.
//
// Note: n and p will always be given as strictly positive integers
//
//digPow(89, 1) should return 1 since 8¹ + 9² = 89 = 89 * 1
//digPow(92, 1) should return -1 since there is no k such as 9¹ + 2² equals 92 * k
//digPow(695, 2) should return 2 since 6² + 9³ + 5⁴= 1390 = 695 * 2
//digPow(46288, 3) should return 51 since 4³ + 6⁴+ 2⁵ + 8⁶ + 8⁷ = 2360688 = 46288 * 51
func DigPow(n, p int) int {
	if n <= 0 || p <= 0 {
		return -1
	}

	sum := 0

	digitStr := strconv.Itoa(n)
	for i := 0; i < len(digitStr); i++ {
		current := string([]rune(digitStr)[i])
		currentDigit, _ := strconv.Atoi(current)
		sum += int(math.Pow(
			float64(currentDigit),
			float64(p),
		))
		p++
	}

	k := sum / n
	if k == 0 {
		return -1
	}
	return k
}
