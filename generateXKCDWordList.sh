#!/bin/bash

set -eu
set -o pipefail

VERSION="8744120d"
SOURCE="https://cdn.rawgit.com/leonardr/olipy/${VERSION}/data/more-corpora/scribblenauts_words.txt"

WORDS=$(curl -sL "${SOURCE}" | awk '/^[a-z]{4}[a-z]*$/{ print "\""$1"\"," }')

cat -s <<EOF > lib/xkcd_words.go
package securepassword

// xkcdWordList contains a list of words derived from the scribblenauts
// word list inside the olipy library by leonardr
// https://github.com/leonardr/olipy
var xkcdWordList = []string{
  ${WORDS}
}
EOF

gofmt -s -w lib/xkcd_words.go
