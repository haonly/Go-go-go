// https://programmers.co.kr/learn/courses/30/lessons/70128

package Solution70128

import "sync"

func Solution(a []int, b []int) int {

	answer := 0
	for i := 0; i < len(a); i++ {
		answer += a[i] * b[i]
	}

	return answer
}

func multiplyAndSum(n1 int, n2 int, sum *int, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()

	// get multiplied number
	multiNum := n1 * n2

	// mutex lock and unlock
	m.Lock()
	*sum += multiNum
	m.Unlock()
}

func SolutionGoroutine(a []int, b []int) int {

	var wg sync.WaitGroup
	var m sync.Mutex

	sum := 0

	wg.Add(len(a))
	for i := 0; i < len(a); i++ {
		go multiplyAndSum(a[i], b[i], &sum, &wg, &m)
	}
	wg.Wait()

	return sum
}
