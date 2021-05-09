package primefactor

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
		want []int
	}{
		{"case 0", args{6}, []int{2, 3}},
		{"case 1", args{8}, []int{2, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("test %v\n", tt.name)
			got := PrimeFactor(tt.args.n)
			for i := 0; i < len(tt.want); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("solution() = #{%v}, want #{%v}", got, tt.want)
				}
			}
		})
	}
}
