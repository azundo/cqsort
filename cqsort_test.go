package cqsort

import (
	"testing"
	"math/rand"
)

func TestBasicSort(t *testing.T) {
	unsorted := rand.Perm(100)
	Cqsort(unsorted)
	for i := 0; i < 100; i++ {
		if unsorted[i] != i {
			t.Errorf("expecting sorted slice")
		}
	}
}
