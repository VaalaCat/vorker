package secret

import "testing"

func TestMD5(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "md5",
			args: args{
				"hello world",
			},
			want: "5eb63bbbe01eeed093cb22bb8f5acdc3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.args.content); got != tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}
