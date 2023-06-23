package gost

import (
	_ "net/http/pprof"
	"testing"

	_ "vorker/services/proxy"
)

func TestInitGost(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "TestInitGost",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitGost()
		})
	}
}
