func shuffle(nums []int, n int) []int {
    answer := []int{}
    for i:= 0; i < n; i++ {
        answer = append(answer, nums[i], nums[i+n])
    }
    return answer
}