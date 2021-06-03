package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
	"time"
)

func Test(t *testing.T) {
	suite.Run(t, &Suite{})
}

type Suite struct {
	suite.Suite
	a     []int
	b     []int
	start time.Time
}

func (s *Suite) SetupTest() {
	s.a = make([]int, 1000000000)
	s.b = make([]int, 1000000000)
	s.a[1000000] = 1
	s.b[1000000] = 1
	s.start = time.Now()
}

func (s Suite) TearDownTest() {
	elapsed := time.Since(s.start)
	log.Printf("[%s] %f(sec)", s.T().Name(), elapsed.Seconds())

}

// xUnit test pattern: 3A
func (s Suite) TestSolution() {
	// Arrange

	// Act
	actual := Solution(s.a, s.b)

	// Assert
	assert.Equal(s.T(), 1, actual)

	/// Output: [Test/TestSolution] 9.102127(sec)
}

// xUnit test pattern: 3A
func (s Suite) TestFastSolution() {
	// Arrange

	// Act
	actual := FastSolution(s.a, s.b)

	// Assert
	assert.Equal(s.T(), 1, actual)

	/// Output: [Test/TestFastSolution] 5.989922(sec)
}

// Table-driven test
func TestSolution(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "tc1", args: struct {
			a []int
			b []int
		}{a: []int{1, 2, 3, 4}, b: []int{-3, -1, 0, 2}}, want: 3},
		{name: "tc2", args: struct {
			a []int
			b []int
		}{a: []int{-1, 0, 1}, b: []int{1, 0, -1}}, want: -2},
		{name: "tc3", args: struct {
			a []int
			b []int
		}{a: []int{0, 0, 0}, b: []int{0, 0, 0}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
