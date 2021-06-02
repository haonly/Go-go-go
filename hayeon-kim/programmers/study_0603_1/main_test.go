package study_0603_1

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{-3, -1, 0, 2}
	assert.Equal(t, 3, solution(a, b))
}

func Test2(t *testing.T) {
	a := []int{-1,0,1}
	b := []int{1, 0, -1}
	assert.Equal(t, -2, solution(a, b))
}
