// +build !gen

package strutils

import "sort"

// Block returns the name of the Unicode block that contains the rune r.
func Block(r rune) string {
	idx := sort.Search(len(unicodeBlocks), func(i int) bool {
		b := unicodeBlocks[i]
		return b.end >= r
	})

	if idx < 0 || idx >= len(unicodeBlocks) {
		return "No_Block"
	}

	return unicodeBlocks[idx].name
}
