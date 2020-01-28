// +build !gen

package strutils

import (
	"testing"
)

func TestName(t *testing.T) {
	tests := []struct {
		r    rune
		name string
	}{
		{0xc5, "LATIN CAPITAL LETTER A WITH RING ABOVE"},
		{0x1f4a9, "PILE OF POO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnicodeName(tt.r); got != tt.name {
				t.Errorf("UnicodeName(U+%04X) = %#v, want %#v", tt.r, got, tt.name)
			}
		})
	}
}
