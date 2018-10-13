package utils

// SplitLength returns a slice of strings, each string is at most length bytes long.
// Does not break UTF-8 codepoints.
func SplitLength(s string, length int) []string {
	var ret []string

	for len(s) > 0 {
		tmp := TruncateString(s, length)
		if len(tmp) == 0 {
			// too short length, or invalid UTF-8 string
			break
		}
		ret = append(ret, tmp)
		s = s[len(tmp):]
	}

	return ret
}

// TruncateString truncates s to a maximum of length bytes without breaking UTF-8 codepoints.
func TruncateString(s string, length int) string {
	if len(s) <= length {
		return s
	}

	cutoff := length
	for s[cutoff]&0xc0 == 0x80 {
		cutoff--
		if cutoff < 0 {
			cutoff = 0
			break
		}
	}

	return s[0:cutoff]
}

// utf8len returns the length of a UTF-8 encoded codepoint based on the starting byte.
func utf8len(b byte) (length int) {
	if b&0x80 == 0 {
		// 0xxxxxxx
		length = 1
	} else if b&0xe0 == 0xc0 {
		// 110xxxxx
		length = 2
	} else if b&0xf0 == 0xe0 {
		// 1110xxxx
		length = 3
	} else if b&0xf8 == 0xf0 {
		// 11110xxx
		length = 4
	}
	return
}
