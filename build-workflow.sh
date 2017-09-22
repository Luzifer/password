#!/bin/bash

set -ex
set -o pipefail

# Check for a publishable version
VERSION=$(git describe --tags --exact-match || echo "notag")
if ( test "${VERSION}" == "notag" ); then
	echo "No exact tag found, no publishing required."
	exit 0
fi

# Collect assets to pack
mkdir -p dist
cp -r password_darwin_amd64 \
	alfred-workflow/exec.py \
	alfred-workflow/icon.png \
	alfred-workflow/info.plist \
	alfred-workflow/lib/workflow \
	dist
echo -n "${VERSION}" > dist/version

# Create ZIP
cd dist
zip -r -9 ../PasswordGenerator.alfredworkflow *
cd -

# Upload to Github releases
github-release upload --user luzifer --repo password --tag ${VERSION} \
	--name PasswordGenerator.alfredworkflow \
	--file PasswordGenerator.alfredworkflow
