package study_0603_1

func solution(a []int, b []int) int {
	ret := 0
	for idx, _ := range a{
		ret += a[idx] * b[idx]
	}
	return ret
}