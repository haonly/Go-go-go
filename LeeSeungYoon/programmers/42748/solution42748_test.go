package Solution42748

import (
	"fmt"
	"testing"
)

func Test_Solution(t *testing.T) {
	type args struct {
		array    []int
		commands [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case 0", args{[]int{1, 5, 2, 6, 3, 7, 4}, [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}}, []int{5, 6, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("test %v\n", tt.name)
			got := Solution(tt.args.array, tt.args.commands)
			for i := 0; i < len(tt.want); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("solution() = #{%v}, want #{%v}", got, tt.want)
				}
			}
		})
	}
}
