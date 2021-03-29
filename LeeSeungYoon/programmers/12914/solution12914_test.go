package Solution12914

import (
	"fmt"
	"testing"
)

func Test_Solution(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"case 0", args{4}, 5},
		{"case 1", args{3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("test %v\n", tt.name)
			if got := Solution(tt.args.n); got != tt.want {
				t.Errorf("solution() = #{%v}, want #{%v}", got, tt.want)
			}
		})
	}
}