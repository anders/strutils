package strutils

import (
	"sort"
	"strings"
)

// UnicodeName returns the Unicode name of the given rune.
func UnicodeName(r rune) string {
	idx := sort.Search(len(unicodeData), func(i int) bool { return unicodeData[i].value >= r })
	if idx < 0 || idx >= len(unicodeData) {
		return ""
	}
	return unicodeData[idx].name
}

// FindRunesByName returns a list of runes whose name contains needle.
func FindRunesByName(needle string) []rune {
	needle = strings.ToUpper(needle)
	ret := []rune{}
	for _, u := range unicodeData {
		if strings.Contains(u.name, needle) || strings.Contains(u.oldName, needle) {
			ret = append(ret, u.value)
		}
	}
	return ret
}

// OldUnicodeName returns the Unicode 1.0 name of the given rune.
func OldUnicodeName(r rune) string {
	idx := sort.Search(len(unicodeData), func(i int) bool { return unicodeData[i].value >= r })
	if idx < 0 || idx >= len(unicodeData) {
		return ""
	}
	return unicodeData[idx].oldName
}

// RuneName returns a name for the given rune (either new Unicode or Unicode 1.0).
func RuneName(r rune) string {
	idx := sort.Search(len(unicodeData), func(i int) bool { return unicodeData[i].value >= r })
	if idx < 0 || idx >= len(unicodeData) {
		return ""
	}
	name := unicodeData[idx].name
	oldName := unicodeData[idx].oldName
	if name != "" {
		return name
	}
	return oldName
}
