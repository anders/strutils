package utils

import (
	"reflect"
	"testing"
)

func TestUTF8Len(t *testing.T) {
	tests := []struct {
		input  string
		length int
	}{
		{"A", 1},            // LATIN CAPITAL LETTER A
		{"\u00c5", 2},       // LATIN CAPITAL LETTER A WITH RING ABOVE
		{"\u3042", 3},       // HIRAGANA LETTER A
		{"\U0001F393", 4},   // GRADUATION CAP
		{"\xff\xff\xff", 0}, // Invalid UTF-8
	}

	for _, test := range tests {
		if got := utf8len(test.input[0]); got != test.length {
			t.Errorf("length of %s wrong, expected %d but got %d", test.input, test.length, got)
		}
	}
}

func TestTruncateString(t *testing.T) {
	tests := []struct {
		input  string
		want   string
		length int
	}{
		{"Anders", "And", 3},
		// Hiragana a times 4
		{"\u3042\u3042\u3042\u3042", "\u3042", 4},
		// Invalid starting bytes
		{"\xff\xff\xff\xff\xff", "", 100},
		// Continuation bytes
		{"\x80\x80\x80\x80\x80", "", 100},
	}
	for _, test := range tests {
		if got := TruncateString(test.input, test.length); got != test.want {
			t.Errorf("expected %q, got %q", test.want, got)
		}
	}
}

func TestSplitLength(t *testing.T) {
	tests := []struct {
		input string
		len   int
		want  []string
	}{
		{
			"AndersSrednaFoobarBazbarX",
			6,
			[]string{"Anders", "Sredna", "Foobar", "Bazbar", "X"},
		},
		{
			"AAAA\u00c5\u00c5\u00c5\u00c5\u3042\u3042\u3042\u3042\U0001F393\U0001F393\U0001F393\U0001F393",
			4,
			[]string{
				"AAAA",                         // 1 byte
				"\u00c5\u00c5", "\u00c5\u00c5", // 2 bytes
				"\u3042", "\u3042", "\u3042", "\u3042", // 3 bytes
				"\U0001F393", "\U0001F393", "\U0001F393", "\U0001F393", // 4 bytes
			},
		},
		{
			"\u3042\u3042\u3042\u3042",
			8,
			[]string{"\u3042\u3042", "\u3042\u3042"},
		},
	}
	for _, tt := range tests {
		if got := SplitLength(tt.input, tt.len); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SplitLength() = %v, want %v", got, tt.want)
		}
	}
}
