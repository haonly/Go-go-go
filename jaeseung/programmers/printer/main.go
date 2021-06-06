package main

func solution(priorities []int, location int) int {
	printed := make([]bool, len(priorities))
	i := -1
	cnt := 0
	for {
		i = findCandidate(priorities, i, printed)
		if canPrint(i, priorities, printed) {
			printed[i] = true
			if i == location {
				return cnt + 1
			}
			cnt++
		}
	}
}

func findCandidate(priorities []int, i int, printed []bool) int {
	for {
		i++
		i = i % len(priorities)
		if !printed[i] {
			break
		}
	}
	return i
}

func canPrint(idx int, priorities []int, printed []bool) bool {
	pri := priorities[idx]
	for i := 0; i < len(priorities); i++ {
		idx = (idx + 1) % len(priorities)
		if printed[idx] {
			i--
			continue
		}
		if pri < priorities[idx] {
			return false
		}
	}
	return true
}
