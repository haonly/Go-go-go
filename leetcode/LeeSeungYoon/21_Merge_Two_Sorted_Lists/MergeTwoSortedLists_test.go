package MaximumSubarray_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	answer := 6
	testResult := MaximumSubarray.MaxSubArray(nums)
	assert.Equal(t, testResult, answer, "Wrong Answer")
}
