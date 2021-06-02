package study_0603_2

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	a := []int{2, 1, 3, 2}
	b := 2
	assert.Equal(t, 1, solution(a, b))
}

func Test2(t *testing.T) {
	a := []int{1, 1, 9, 1, 1, 1}
	b := 0
	assert.Equal(t, 5, solution(a, b))
}

func Test3(t *testing.T) {
	a := []int{1, 2, 3}
	b := 1
	assert.Equal(t, 2, solution(a, b))
}

func Test4(t *testing.T) {
	a := []int{1}
	b := 0
	assert.Equal(t, 1, solution(a, b))
}

func Test5(t *testing.T) {
	a := []int{2, 2, 2, 1, 3, 4}
	b := 3
	assert.Equal(t, 6, solution(a, b))
}

