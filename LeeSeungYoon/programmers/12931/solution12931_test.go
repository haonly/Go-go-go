package Solution12931

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
		want int
	}{
		{"case 0", args{123}, 6},
		{"case 1", args{987}, 24},
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