// Solution2
package MaximumSubarray

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func MaxSubArray2(nums []int) int {
	currentMax := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currentMax = max(nums[i], currentMax+nums[i])
		maxSum = max(maxSum, currentMax)
	}
	return maxSum
}
