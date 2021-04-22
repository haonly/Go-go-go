package ShuffletheArray

func Shuffle(nums []int, n int) []int {
	answer := []int{}

	for i := 0; i < n; i++ {
		answer = append(answer, nums[i], nums[i+n])
	}

	return answer
}

// 0 1 2 3 4 5 6 7
// 0 4 / 1 5 / 2 6 / 3 7
