package Solution12918

import (
	"fmt"
	"testing"
)

func Test_Solution(t *testing.T) {
	type args struct {
		s  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case 0", args{"a234"}, false},
		{"case 1", args{"1234"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("test %v\n", tt.name)
			if got := Solution(tt.args.s); got != tt.want {
				t.Errorf("solution() = #{%v}, want #{%v}", got, tt.want)
			}
		})
	}
}