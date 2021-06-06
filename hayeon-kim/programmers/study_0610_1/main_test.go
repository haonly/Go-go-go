package study_0610_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func Test1(t *testing.T) {
	a := []int{2, 1, 3, 2}
	b := []bool{true, false, true, true}
	assert.Equal(t, 6, solution(a, b))
}

func Test2(t *testing.T) {
	a := []int{4, 7, 12}
	b := []bool{true, false, true}
	assert.Equal(t, 9, solution(a, b))
}

func Test3(t *testing.T) {
	a := []int{1, 2, 3}
	b := []bool{false, false, true}
	assert.Equal(t, 0, solution(a, b))
}