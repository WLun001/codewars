package codewar

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func ShortestBinary(A []int, B []int) []int {
	aStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(A)), ""), "[]")
	bStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(B)), ""), "[]")
	a, _ := strconv.ParseInt(aStr, 2, 64)
	b, _ := strconv.ParseInt(bStr, 2, 64)
	binary := strconv.FormatInt(a+b, 2)
	arr := make([]int, len(binary))
	log.Println(a)
	log.Println(b)
	for i := range arr {
		arr[i], _ = strconv.Atoi(string(binary[i]))
	}
	return arr

	//return append(bits(-decimal(A)), bits(-decimal(B))...)
}

func decimal(a []int) int {
	decimal := 0

	for i, b := range a {
		decimal += b * int(math.Pow(float64(-2), float64(i)))
	}

	return decimal
}

func bits(a int) []int {
	num := -a
	var bits []int
	for num != 0 {
		bits = append(bits, num%2)
		num /= -2
	}

	return bits
}
