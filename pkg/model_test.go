package pkg

import (
	"fmt"
	"testing"
)

func TestDomain_Valid(t *testing.T) {
	tests := []struct {
		name   string
		domain Domain
		want   bool
	}{
		{
			name:   "valid",
			domain: "example.com",
			want:   true,
		},
		{
			name:   "valid",
			domain: "server.example.com",
			want:   true,
		},
		{
			name:   "valid",
			domain: "server",
			want:   true,
		},
		{
			name:   "valid",
			domain: "localhost",
			want:   true,
		},
		{
			name:   "invalid",
			domain: "",
			want:   false,
		},
		{
			name:   "invalid",
			domain: "3245 2345  13245",
			want:   false,
		},
		{
			name:   "invalid",
			domain: "iubsir fgonse[oif p[ijoef ioj[oisef [oi[e",
			want:   false,
		},
	}
	for n, tt := range tests {
		t.Run(fmt.Sprintf("%s-%d", tt.name, n+1), func(t *testing.T) {
			if got := tt.domain.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIp_Valid(t *testing.T) {
	tests := []struct {
		name string
		ip   Ip
		want bool
	}{
		{name: "valid", ip: "127.0.0.1", want: true},
		{name: "valid", ip: "244.178.44.111", want: true},
		{name: "valid", ip: "0.0.0.0", want: true},

		{name: "invalid", ip: "", want: false},
		{name: "invalid", ip: "3245 2345  13245", want: false},
		{name: "invalid", ip: "256.235.235.2", want: false},
		{name: "invalid", ip: "234.333.235.678", want: false},
		{name: "invalid", ip: "iubsir fgonse[oif p[ijoef ioj[oisef [oi[e", want: false},
	}
	for n, tt := range tests {
		t.Run(fmt.Sprintf("%s-%d", tt.name, n+1), func(t *testing.T) {
			if got := tt.ip.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
