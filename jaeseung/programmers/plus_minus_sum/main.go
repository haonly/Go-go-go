package main

func Solution(absolutes []int, signs []bool) int {
	sum := 0
	for i, n := range absolutes {
		if signs[i] {
			sum += n
			continue
		}
		sum -= n
	}
	return sum
}
