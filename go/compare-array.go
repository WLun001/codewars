package codewar

import (
	"reflect"
)

// Given two arrays a and b write a function comp(a, b) (compSame(a, b) in Clojure) that checks whether the two arrays have the "same" elements, with the same multiplicities. "Same" means, here, that the elements in b are the elements in a squared, regardless of the order.

// Examples
// Valid arrays
// a = [121, 144, 19, 161, 19, 144, 19, 11]
// b = [121, 14641, 20736, 361, 25921, 361, 20736, 361]
// comp(a, b) returns true because in b 121 is the square of 11, 14641 is the square of 121, 20736 the square of 144, 361 the square of 19, 25921 the square of 161, and so on. It gets obvious if we write b's elements in terms of squares:

// a = [121, 144, 19, 161, 19, 144, 19, 11]
// b = [11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19]

// Invalid arrays
// If we change the first number to something else, comp may not return true anymore:
//
// a = [121, 144, 19, 161, 19, 144, 19, 11]
// b = [132, 14641, 20736, 361, 25921, 361, 20736, 361]
// comp(a,b) returns false because in b 132 is not the square of any number of a.
//
// a = [121, 144, 19, 161, 19, 144, 19, 11]
// b = [121, 14641, 20736, 36100, 25921, 361, 20736, 361]
// comp(a,b) returns false because in b 36100 is not the square of any number of a.
func CompArr(array1, array2 []int) bool {
	if len(array1) < 0 || len(array2) < 0 {
		return false
	}

	if len(array1) != len(array2) {
		return false
	}

	compArr := make([]int, 0)

	for i := 0; i < len(array1); i++ {
		if i == 0 {
			last := array1[len(array1)-1]
			compArr = append(compArr, last*last)
		} else {
			item := array1[i-1]
			compArr = append(compArr, item*item)
		}
		continue
	}

	return reflect.DeepEqual(compArr, array2)
}
