package strutils

type unicodeBlock struct {
	start rune
	end   rune
	name  string
}

type unicodeRow struct {
	value   rune   // 0
	name    string // 1
	oldName string // 10
}
