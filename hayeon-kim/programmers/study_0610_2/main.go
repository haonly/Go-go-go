package study_0610_2

func solution(answers []int) []int {

	var case1 = []int{1,2,3,4,5}
	var case2 = []int{2,1,2,3,2,4,2,5}
	var case3 = []int{3,3,1,1,2,2,4,4,5,5}

	var solved = make([]int, 3, 3)

	for idx, ans := range answers{
		if ans == case1[idx% (len(case1))]{
			solved[0] += 1
		}
		if ans == case2[idx % (len(case2))]{
			solved[1] += 1
		}
		if ans == case3[idx % (len(case3))]{
			solved[2] += 1
		}
	}
	return biggest(solved)
}

func biggest(scores []int) []int {
	var ret = make([]int, 0, 0)

	var max = 0
	for _, score := range scores{
		if max < score{
			max = score
		}
	}
	for i, score := range scores{
		if max == score {
			ret = append(ret, i+1)
		}
	}

	return ret
}