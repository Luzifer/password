#!/bin/bash
set -euxo pipefail

# Check for a publishable version
VERSION=$(git describe --tags --exact-match || echo "notag")

# Collect assets to pack
mkdir -p dist
unzip ./.build/password_darwin_amd64.zip -d dist
cp -r \
	alfred-workflow/exec.py \
	alfred-workflow/icon.png \
	alfred-workflow/info.plist \
	alfred-workflow/lib/workflow \
	dist

# In order to have a valid version number on no tags for testing
git describe --tags >dist/version

# Create ZIP
cd dist
zip -r -9 ../PasswordGenerator.alfredworkflow *
cd -

# If there is no version tag the artifact is not to be published
if (test "${VERSION}" == "notag"); then
	echo "No exact tag found, no publishing required."
	exit 0
fi

# Upload to Github releases
github-release upload --user luzifer --repo password --tag ${VERSION} \
	--name PasswordGenerator.alfredworkflow \
	--file PasswordGenerator.alfredworkflow
