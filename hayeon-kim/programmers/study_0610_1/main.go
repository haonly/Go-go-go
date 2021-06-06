package study_0610_1

func solution(absolutes []int, signs []bool) int {
	var ans = 0
	for idx, _ := range absolutes{
		if signs[idx] == false {
			ans += absolutes[idx] * -1
		} else {
			ans += absolutes[idx]
		}
	}
	return ans
}