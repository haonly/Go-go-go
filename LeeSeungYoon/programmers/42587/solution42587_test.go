package Solution42587

import (
	"fmt"
	"testing"
)

type args struct {
	priorities []int
	location   int
}

func Test_Solution(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 0", args{[]int{2, 1, 3, 2}, 2}, 1},
		{"case 1", args{[]int{1, 1, 9, 1, 1, 1}, 0}, 5},
		{"case 1-1", args{[]int{1, 1, 9, 1, 1, 1}, 1}, 6},
		{"case 2", args{[]int{1}, 0}, 1},
		{"case 3", args{[]int{1, 1}, 1}, 2},
		{"case 3", args{[]int{1, 1, 1}, 1}, 2},
		{"case 4", args{[]int{1, 2, 3, 4, 5, 6}, 0}, 6},
		{"case 4", args{[]int{2, 2, 2, 1, 3, 4}, 3}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("test %v\n", tt.name)
			got := Solution(tt.args.priorities, tt.args.location)
			if got != tt.want {
				t.Errorf("solution() = #{%v}, want #{%v}", got, tt.want)
			}
		})
	}
}
