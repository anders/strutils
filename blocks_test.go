package strutils

import (
	"fmt"
	"testing"
)

func TestBlock(t *testing.T) {
	tests := []struct {
		r    rune
		want string
	}{
		{0x0, "Basic Latin"},
		{0x1, "Basic Latin"},
		{0x80, "Latin-1 Supplement"},
		{0x187, "Latin Extended-B"},
		{0x110000, "No_Block"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("U+%04X", tt.r), func(t *testing.T) {
			if got := Block(tt.r); got != tt.want {
				t.Errorf("Block(U+%04x) = %v, want %v", tt.r, got, tt.want)
			}
		})
	}
}
