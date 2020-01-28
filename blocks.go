// +build !gen

package strutils

// Block returns the name of the Unicode block that contains the rune r.
func Block(r rune) string {
	for _, b := range unicodeBlocks {
		if r >= b.start && r <= b.end {
			return b.name
		}
	}

	return "No_Block"
}
