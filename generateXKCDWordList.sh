#!/bin/bash
set -euo pipefail

version="8744120d"
source="https://raw.githubusercontent.com/leonardr/olipy/${version}/data/more-corpora/scribblenauts_words.txt"

words=$(curl -sL "${source}" | awk '/^[a-z]{4}[a-z]*$/{ print "\""$1"\"," }')

cat -s <<EOF >lib/xkcd_words.go
package securepassword

// xkcdWordList contains a list of words derived from the scribblenauts
// word list inside the olipy library by leonardr
// https://github.com/leonardr/olipy
var xkcdWordList = []string{
  ${words}
}
EOF

gofmt -s -w lib/xkcd_words.go
