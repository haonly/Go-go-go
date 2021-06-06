package main

import (
	"github.com/stretchr/testify/suite"
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

func Test_solution(t *testing.T) {
	type args struct {
		priorities []int
		location   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "tc1", args: args{
			priorities: []int{3, 2, 1, 3},
			location:   0,
		}, want: 1},
		{name: "tc2", args: args{
			priorities: []int{3, 2, 1, 1},
			location:   1,
		}, want: 2},
		{name: "tc3", args: args{
			priorities: []int{2, 1, 3, 2},
			location:   2,
		}, want: 1},
		{name: "tc4", args: args{
			priorities: []int{1, 1, 9, 1, 1, 1},
			location:   0,
		}, want: 5},
		{name: "tc5", args: args{
			priorities: []int{1, 1, 9, 1, 1, 1},
			location:   1,
		}, want: 6},
		{name: "tc6", args: args{
			priorities: []int{2, 2, 1, 2, 1},
			location:   2,
		}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.priorities, tt.args.location); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
