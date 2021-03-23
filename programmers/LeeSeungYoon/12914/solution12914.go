package Solution12914

func Solution(n int) int64{

	dpMap := map[int]int{}

	var dp func(int)int
	dp = func(num int) int{
		if num == 1 || num == 2 {
			return num
		}
		// num과 같은 key 가 map에 저장되어 있는지 비교
		for k, v := range dpMap{
			if k == num {
				return v
			}
		}
		// 매 step마다 1234567로 나눠준 몫을 저장
		dpMap[num] = (dp(num-1) + dp(num-2))%1234567
		return dpMap[num]
	}

	return int64(dp(n)%1234567)
}



// 1칸 - 1 
// 2칸 - (1,1), 2 -> 2
// 3칸 - (1,1,1), (1,2), (2,1) -> 3
// 4칸 - (1,1,1,1), (1,2,1) (2,1,1), (1,1,2), (2,2) -> 5
// 5칸 - (1,1,1,1,1) -> 8개
// (1,2,2), (2,1,2), (2,2,1)
// (1,1,1,2), (1,1,2,1), (1,2,1,1), (2,1,1,1)
