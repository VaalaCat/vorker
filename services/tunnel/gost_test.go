package tunnel

import "testing"

func Test_genConf(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genConf(); got != tt.want {
				t.Errorf("genConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
