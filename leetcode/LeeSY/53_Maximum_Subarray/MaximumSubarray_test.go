package MaximumSubarray_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	MaximumSubarray "leetcode/53_Maximum_Subarray"
)

func TestCase1(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	answer := 6
	testResult := MaximumSubarray.MaxSubArray(nums)
	assert.Equal(t, testResult, answer, "Wrong Answer")
}

func TestCase2(t *testing.T) {
	nums := []int{1}
	answer := 1
	testResult := MaximumSubarray.MaxSubArray(nums)
	assert.Equal(t, testResult, answer, "Wrong Answer")
}
func TestCase3(t *testing.T) {
	nums := []int{5, 4, -1, 7, 8}
	answer := 23
	testResult := MaximumSubarray.MaxSubArray(nums)
	assert.Equal(t, testResult, answer, "Wrong Answer")
}
