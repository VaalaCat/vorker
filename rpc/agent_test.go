package rpc

import "testing"

func TestSyncAgent(t *testing.T) {
	type args struct {
		endpoint string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				endpoint: "http://localhost:8888",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := SyncAgent(tt.args.endpoint); (err != nil) != tt.wantErr {
				t.Errorf("SyncAgent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
