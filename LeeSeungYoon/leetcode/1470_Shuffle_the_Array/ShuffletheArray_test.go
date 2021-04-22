package ShuffletheArray

import (
	"fmt"
	"testing"
)

func Test_Shuffle(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case 0", args{[]int{2, 5, 1, 3, 4, 7}, 3}, []int{2, 3, 5, 4, 1, 7}},
		{"case 1", args{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4}, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{"case 2", args{[]int{1, 1, 2, 2}, 2}, []int{1, 2, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("test %v\n", tt.name)
			got := Shuffle(tt.args.nums, tt.args.n)
			for i := 0; i < len(tt.want); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("solution() = #{%v}, want #{%v}", got, tt.want)
				}
			}
		})
	}
}
