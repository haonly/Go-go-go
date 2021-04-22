package Solution42586

import (
	"fmt"
	"testing"
)

func Test_Solution(t *testing.T) {
	type args struct {
		progresses []int
		speeds     []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case 0", args{[]int{93, 30, 55}, []int{1, 30, 5}}, []int{2, 1}},
		{"case 1", args{[]int{95, 90, 99, 99, 80, 99}, []int{1, 1, 1, 1, 1, 1}}, []int{1, 3, 2}},
		{"case 2", args{[]int{}, []int{}}, []int{}},
		{"case 3", args{[]int{99, 99, 99}, []int{1, 1, 1}}, []int{3}},
		{"case 4", args{[]int{99}, []int{1}}, []int{1}},
		{"case 5", args{[]int{99, 1, 99, 1}, []int{1, 1, 1, 1}}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("test %v\n", tt.name)
			got := Solution(tt.args.progresses, tt.args.speeds)
			for i := 0; i < len(tt.want); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("solution() = #{%v}, want #{%v}", got, tt.want)
				}
			}
		})
	}
}
