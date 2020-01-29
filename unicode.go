package strutils

//go:generate go run ./unicodegen/unicodegen.go

type unicodeBlock struct {
	start rune
	end   rune
	name  string
}

type unicodeRow struct {
	value   rune
	name    string
	oldName string
}
