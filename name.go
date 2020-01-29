package strutils

import (
	"sort"
	"strings"
)

// UnicodeName returns the Unicode name of the given rune.
func UnicodeName(r rune) string {
	idx := sort.Search(len(unicodeData), func(i int) bool { return unicodeData[i].R >= r })
	if idx < 0 || idx >= len(unicodeData) {
		return ""
	}
	return unicodeData[idx].Name
}

// FindRunesByName returns a list of runes whose name contains needle.
func FindRunesByName(needle string) []rune {
	needle = strings.ToUpper(needle)
	ret := []rune{}
	for _, u := range unicodeData {
		if strings.Contains(u.Name, needle) || strings.Contains(u.OldName, needle) {
			ret = append(ret, u.R)
		}
	}
	return ret
}

// OldUnicodeName returns the Unicode 1.0 name of the given rune.
func OldUnicodeName(r rune) string {
	idx := sort.Search(len(unicodeData), func(i int) bool { return unicodeData[i].R >= r })
	if idx < 0 || idx >= len(unicodeData) {
		return ""
	}
	return unicodeData[idx].OldName
}

// RuneName returns a name for the given rune (either new Unicode or Unicode 1.0).
func RuneName(r rune) string {
	idx := sort.Search(len(unicodeData), func(i int) bool { return unicodeData[i].R >= r })
	if idx < 0 || idx >= len(unicodeData) {
		return ""
	}
	name := unicodeData[idx].Name
	oldName := unicodeData[idx].OldName
	if name != "" {
		return name
	}
	return oldName
}
