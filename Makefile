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
	cp password_darwin_amd64 password
	zip -9 -j PasswordGenerator.alfredworkflow alfred-workflow/* password
	github-release upload --user luzifer --repo password --tag $(shell git describe --tags --exact-match) --name PasswordGenerator.alfredworkflow --file PasswordGenerator.alfredworkflow
