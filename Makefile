default: pack

compile_js:
	rm -f ./frontend/assets/*
	docker run --rm -i \
		-v "$(CURDIR):$(CURDIR)" \
		-w "$(CURDIR)/js" \
		-u $(shell id -u) \
		node:10-alpine \
		sh -c "npx npm@lts ci && npx webpack"

debug:
	go-bindata --debug -o cmd/password/bindata.go --pkg=main frontend/...
	go run *.go serve

pack: compile_js
	go-bindata -modtime 1 -o cmd/password/bindata.go --pkg=main frontend/...
	bash generateXKCDWordList.sh

publish:
	curl -sSLo cmd/password/golang.sh https://raw.githubusercontent.com/Luzifer/github-publish/master/golang.sh
	cd cmd/password && bash golang.sh
