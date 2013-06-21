package cqsort

import "math/rand"

var _ = rand.Intn // comment out after debugging

func Cqsort(s []int) {
	if len(s) <= 1 {
		return
	}
	pivot := partition(s)
	Cqsort(s[:pivot+1])
	Cqsort(s[pivot+1:])
	return
}

func partition(s []int) (swapIdx int) {
	pivotIdx, pivot := pickPivot(s)
	// swap right-most element and pivot
	s[len(s)-1], s[pivotIdx] = s[pivotIdx], s[len(s)-1]
	// sort elements keeping track of pivot's idx
	for i := 0; i < len(s) - 1; i++ {
		if s[i] < pivot {
			s[i], s[swapIdx] = s[swapIdx], s[i]
			swapIdx++
		}
	}
	// swap pivot back to its place and return
	s[swapIdx], s[len(s)-1] = s[len(s)-1], s[swapIdx]
	return
}

func pickPivot(s []int)(pivotIdx int, pivot int) {
	pivotIdx = rand.Intn(len(s))
	pivot = s[pivotIdx]
	return
}
