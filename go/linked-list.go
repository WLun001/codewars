package codewar

func linkedList(A []int) int {
	return linkedListLength(A, A[0], 1)
}

func linkedListLength(a []int, index, length int) int {
	if a[index] == -1 {
		return length + 1
	}
	return linkedListLength(a, a[index], length+1)
}
