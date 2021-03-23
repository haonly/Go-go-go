// Solution1 - DP
package MaximumSubarray

import "fmt"

func maxValue(dict map[string]int) int {
	max := dict["0_0"]
	for _, v := range dict {
		if max < v {
			max = v
		}
	}
	return max
}

func MaxSubArray(nums []int) int {

	dict := make(map[string]int)

	window := 0
	max := nums[0]
	for window < len(nums) {
		for i := 0; i+window < len(nums); i++ {
			key := fmt.Sprintf("%d_%d", i, i+window) // 0_1, 0_2

			if window == 0 {
				dict[key] = nums[i]
			}

			key1 := fmt.Sprintf("%d_%d", i, i+window-1)      // 0_0, 0_1
			key2 := fmt.Sprintf("%d_%d", i+window, i+window) // 1_1, 2_2

			dict[key] = dict[key1] + dict[key2]

			if max < dict[key] {
				max = dict[key]
			}

		}
		window++
	}
	return max
}
