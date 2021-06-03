package main

import (
	"sync"
)

func solution(a []int, b []int) int {
	cnt := 0
	for i := 0; i < len(a); i++ {
		cnt += a[i] * b[i]
	}
	return cnt
}

func Solution(a, b []int) int {
	return solution(a, b)
}

func fastSolution(a []int, b []int) int {
	const NumTasks = 100
	workSize := len(a) / NumTasks

	ch := make(chan int)

	wg := &sync.WaitGroup{}

	// Fan-Out
	wg.Add(NumTasks)
	for i := 0; i < NumTasks; i++ {
		go func(c chan<- int, startIdx, size int) {
			defer wg.Done()
			if len(a) <= startIdx {
				return
			}
			cnt := 0
			for j := startIdx; j < startIdx+size; j++ {
				cnt += a[j] * b[j]
			}
			c <- cnt
		}(ch, i*workSize, workSize)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	cnt := 0
	for v := range ch {
		cnt += v
	}
	return cnt
}

func FastSolution(a, b []int) int {
	return fastSolution(a, b)
}
