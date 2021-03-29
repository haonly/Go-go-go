package Solution12948

import (
	"fmt"
	"testing"
)

func Test_Solution(t *testing.T) {
	type args struct {
		phone_number string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case 0", args{"01033334444"}, "*******4444"},
		{"case 1", args{"027778888"}, "*****8888"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("test %v\n", tt.name)
			if got := Solution(tt.args.phone_number); got != tt.want {
				t.Errorf("solution() = #{%v}, want #{%v}", got, tt.want)
			}
		})
	}
}