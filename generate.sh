#!/bin/bash
for x in Blocks.txt UnicodeData.txt; do
  if [ ! -f "$x" ]; then
    echo "downloading $x"
    LANG=C wget -q "https://www.unicode.org/Public/UCD/latest/ucd/${x}" -O "$x"
  else
    echo "skipping $x, already exists"
  fi
done

go generate
