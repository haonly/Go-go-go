package FizzBuzz

import (
	"fmt"
	"testing"
)

func Test_Shuffle(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case 0", args{1}, ""},
		{"case 1", args{2}, ""},
		{"case 2", args{3}, "Fizz"},
		{"case 3", args{6}, "Fizz"},
		{"case 4", args{5}, "Buzz"},
		{"case 5", args{10}, "Buzz"},
		{"case 6", args{15}, "FizzBuzz"},
		{"case 7", args{30}, "FizzBuzz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("test %v\n", tt.name)
			got := FizBuz(tt.args.n)
			for i := 0; i < len(tt.want); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("solution() = #{%v}, want #{%v}", got, tt.want)
				}
			}
		})
	}
}
