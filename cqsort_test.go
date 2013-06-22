package cqsort

import (
	"testing"
	"math/rand"
	"sort"
)

var _ = sort.Ints
func TestBasicSort(t *testing.T) {
	sortSize := 100000
	unsorted := rand.Perm(sortSize)
	Cqsort(unsorted)
	// Qsort(unsorted)
	// sort.Ints(unsorted)
	for i := 0; i < sortSize; i++ {
		if unsorted[i] != i {
			t.Errorf("expecting sorted slice")
		}
	}
}
