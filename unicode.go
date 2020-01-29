package strutils

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/gob"
	"log"
)

//go:generate go run ./unicodegen/unicodegen.go

type unicodeBlock struct {
	start rune
	end   rune
	name  string
}

type unicodeRow struct {
	R       rune
	Name    string
	OldName string
}

var unicodeData []unicodeRow

func init() {
	b, err := base64.StdEncoding.DecodeString(gobbedUnicode)
	if err != nil {
		log.Fatalf("failed decoding unicode data: %s", err)
	}
	r := bytes.NewReader(b)
	rd, err := zlib.NewReader(r)
	if err != nil {
		log.Fatalf("failed zlib reader: %s", err)
	}
	dec := gob.NewDecoder(rd)
	if err := dec.Decode(&unicodeData); err != nil {
		log.Fatalf("failed decoding gob unicode data: %s", err)
	}
}