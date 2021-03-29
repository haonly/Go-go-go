package Solution42576

import (
	"fmt"
	"testing"
)

func Test_Solution(t *testing.T) {
	type args struct {
		participant []string
		completion []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case 0", args{[]string{"leo", "kiki", "eden"}, []string{"eden", "kiki"}}, "leo"},
		{"case 1", args{[]string{"marina", "josipa", "nikola", "vinko", "filipa"}, []string{"josipa", "filipa", "marina", "nikola"}}, "vinko"},
		{"case 2", args{[]string{"mislav", "stanko", "mislav", "ana"}, []string{"stanko", "ana", "mislav"}}, "mislav"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("test %v\n", tt.name)
			if got := Solution(tt.args.participant, tt.args.completion); got != tt.want {
				t.Errorf("solution() = #{%v}, want #{%v}", got, tt.want)
			}
		})
	}
}