package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		absolutes []int
		signs     []bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "tc1", args: struct {
			absolutes []int
			signs     []bool
		}{absolutes: []int{4, 7, 12}, signs: []bool{true, false, true}}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.absolutes, tt.args.signs); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
