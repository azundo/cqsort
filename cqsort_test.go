package cqsort

import (
	"testing"
	"math/rand"
	"sort"
)

var _ = sort.Ints
func TestBasicSort(t *testing.T) {
	sortSize := 1000000
	MAXGOROUTINES = 10000
	unsorted := make([]int, 0, sortSize)
	unsorted = rand.Perm(sortSize)
	Cqsort(unsorted)
	for i := 0; i < sortSize; i++ {
		if unsorted[i] != i {
			t.Errorf("expecting sorted slice")
			return
		}
	}
}

func BenchmarkCqsort1000(b *testing.B) {
	sortSize := 1000000
	MAXGOROUTINES = 1000
	unsorted := make([]int, 0, sortSize)
	for i:= 0; i < b.N; i++ {
		unsorted = rand.Perm(sortSize)
		Cqsort(unsorted)
	}
}

func BenchmarkCqsort2(b *testing.B) {
	sortSize := 1000000
	MAXGOROUTINES = 2
	unsorted := make([]int, 0, sortSize)
	for i:= 0; i < b.N; i++ {
		unsorted = rand.Perm(sortSize)
		Cqsort(unsorted)
	}
}

func BenchmarkCqsort10000(b *testing.B) {
	sortSize := 1000000
	MAXGOROUTINES = 10000
	unsorted := make([]int, 0, sortSize)
	for i:= 0; i < b.N; i++ {
		unsorted = rand.Perm(sortSize)
		Cqsort(unsorted)
	}
}

func BenchmarkQsort(b *testing.B) {
	sortSize := 1000000
	unsorted := make([]int, 0, sortSize)
	for i:= 0; i < b.N; i++ {
		unsorted = rand.Perm(sortSize)
		Qsort(unsorted)
	}
}

func BenchmarkSortInts(b *testing.B) {
	sortSize := 1000000
	unsorted := make([]int, 0, sortSize)
	for i:= 0; i < b.N; i++ {
		unsorted = rand.Perm(sortSize)
		sort.Ints(unsorted)
	}
}
