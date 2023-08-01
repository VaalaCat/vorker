package utils

import "testing"

func Test_getFlagValue(t *testing.T) {
	type args struct {
		flag  string
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				flag:  "-L",
				input: "-a -L rtcp://:0/ -F 123123123",
			},
			want: "rtcp://:0/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFlagValue(tt.args.flag, tt.args.input); got != tt.want {
				t.Errorf("getFlagValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
