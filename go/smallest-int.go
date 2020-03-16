package codewar

import (
	"sort"
)

// Find the smallest positive integer that does not occur in a given sequence.
func SmallestInt(A []int) int {
	sort.Ints(A)
	smallest := A[0]
	uniqueSlice := unique(A)

	for _, v := range uniqueSlice {
		if v == smallest {
			continue
		}
		if v < smallest {
			break
		}
		smallest++
	}

	for {
		if smallest <= 0 {
			smallest++
			continue
		}
		break
	}

	if contains(uniqueSlice, smallest) {
		smallest++
	}
	return smallest
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func contains(intSlice []int, a int) bool {
	for _, v := range intSlice {
		if v == a {
			return true
		}
	}
	return false
}
