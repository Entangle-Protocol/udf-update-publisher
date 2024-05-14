package utils

import (
	"testing"
)

func TestAsciiToPaddedHex(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		expected [32]byte
		wantErr bool
	}{
		{
			name:    "Test 1",
			str:     "hello",
			expected: [32]byte{0x68, 0x65, 0x6c, 0x6c, 0x6f},
			wantErr: false,
		},
		{
			name:    "Test 2",
			str:     "NGL/USDT",
			expected: [32]byte{0x4e, 0x47, 0x4c, 0x2f, 0x55, 0x53, 0x44, 0x54},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			paddedHex, err := AsciiToPaddedHex(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("AsciiToPaddedHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if paddedHex != tt.expected {
				t.Errorf("AsciiToPaddedHex() = %v, want %v", paddedHex, tt.expected)
			}
		})
	}
}
