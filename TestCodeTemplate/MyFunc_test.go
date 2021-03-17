package myFunc

import ("testing")

func Test_MyFunc(t *testing.T){
	type args struct {
		input bool
	}
	tests := []struct{
		name string
		args args
		want bool
	}{
	// TODO : Add test cases 
		{"0", args{true}, true},
		{"1", args{true}, false},
		{"2", args{true}, true},
	}

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			if got := MyFunc(tt.args.input); got != tt.want{
				t.Errorf("IsLongPressedName() = #{%v}, want #{%v}", got, tt.want)
			}
		})
	}
}
