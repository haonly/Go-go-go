package study_0610_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func Test1(t *testing.T) {
	a := []int{1,2,3,4,5}
	ret := []int{1}
	assert.Equal(t, ret, solution(a))
}

func Test2(t *testing.T) {
	a := []int{1,3,2,4,2}
	ret := []int{1,2,3}
	assert.Equal(t, ret, solution(a))
}

func Test3(t *testing.T) {
	a := []int{1,1,1,1}
	ret := []int{3}
	assert.Equal(t, ret, solution(a))
}