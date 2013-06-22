package cqsort

import (
	"math/rand"
	"fmt"
)

var _ = fmt.Println // comment out after debugging

const NCORES = 1

func Qsort(s []int) {
	if len(s) <= 1 {
		return
	}
	pivot := partition(s)
	Qsort(s[:pivot+1])
	Qsort(s[pivot+1:])
	return
}

func Cqsort(s []int) {

	workerQueue := make(chan int, NCORES)
	for i := 0 ; i < NCORES; i++ {
		workerQueue <- 1
	}
	doneChan := make(chan int)
	var innerCqsort func([]int, chan int, chan int)
	innerCqsort = func(s []int, workerQueue chan int, dChan chan int) {
		if len(s) <= 1 {
			dChan <- 1
			return
		}
		// all the work happens in partition so pull from the workerQueue here
		<-workerQueue
		pivot := partition(s)
		// the hard work is done so allow another worker to go
		workerQueue <- 1
		// channels for recursive calls to signal they are done
		lChan := make(chan int)
		rChan := make(chan int)
		go innerCqsort(s[:pivot+1], workerQueue, lChan)
		go innerCqsort(s[pivot+1:], workerQueue, rChan)
		// block until the l and r channels report they are finished
		for i := 0; i < 2; i++ {
			select {
			case <-lChan:
				// left is done
			case <-rChan:
				// right is done
			}
		}
		// report to caller that we're finished
		dChan <- 1
		return
	}
	go innerCqsort(s, workerQueue, doneChan)
	<-doneChan
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
