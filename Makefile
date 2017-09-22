default: pack

compile_coffee:
		coffee -c frontend/assets/application.coffee

debug: compile_coffee
		go-bindata --debug frontend/...
		go run *.go serve
		open http://127.0.0.1:3000/

pack: compile_coffee
		go-bindata frontend/...

publish:
	curl -sSLo golang.sh https://raw.githubusercontent.com/Luzifer/github-publish/master/golang.sh
	bash golang.sh

workflow:
	mkdir -p dist
	cp -r password_darwin_amd64 \
		alfred-workflow/exec.py \
		alfred-workflow/icon.png \
		alfred-workflow/info.plist \
		alfred-workflow/lib/workflow \
		dist
	echo -n $(shell git describe --tags --exact-match) > dist/version
	cd dist && zip -r -9 ../PasswordGenerator.alfredworkflow *
	github-release upload --user luzifer --repo password --tag $(shell git describe --tags --exact-match) --name PasswordGenerator.alfredworkflow --file PasswordGenerator.alfredworkflow
