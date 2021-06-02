package Solution70128

import (
	"fmt"
	"testing"
)

type args struct {
	a []int
	b []int
}

func Test_Solution(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 0", args{[]int{1, 2, 3, 4}, []int{-3, -1, 0, 2}}, 3},
		{"case 1", args{[]int{-1, 0, 1}, []int{1, 0, -1}}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("test %v\n", tt.name)
			got := Solution(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("solution() = #{%v}, want #{%v}", got, tt.want)
			}
		})
	}
}

func Test_SolutionGoroutine(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 0", args{[]int{1, 2, 3, 4}, []int{-3, -1, 0, 2}}, 3},
		{"case 1", args{[]int{-1, 0, 1}, []int{1, 0, -1}}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("test %v\n", tt.name)
			got := SolutionGoroutine(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("solution() = #{%v}, want #{%v}", got, tt.want)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 0", args{[]int{1, 2, 3, 4}, []int{-3, -1, 0, 2}}, 3},
		{"case 1", args{[]int{-1, 0, 1}, []int{1, 0, -1}}, -2},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			Solution(tt.args.a, tt.args.b)
		}
	}
}

func BenchmarkSolutionGoroutine(b *testing.B) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 0", args{[]int{1, 2, 3, 4}, []int{-3, -1, 0, 2}}, 3},
		{"case 1", args{[]int{-1, 0, 1}, []int{1, 0, -1}}, -2},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			SolutionGoroutine(tt.args.a, tt.args.b)
		}
	}
}
