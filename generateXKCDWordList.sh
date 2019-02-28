#!/bin/bash
set -euo pipefail

version="19af63af"
source="https://raw.githubusercontent.com/leonardr/olipy/${version}/olipy/data/corpora-olipy/words/scribblenauts.json"

words=$(curl -sL "${source}" | jq -r '.nouns | .[]' | sort | awk '/^[a-z]{4}[a-z]*$/{ print "\""$1"\"," }')

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
