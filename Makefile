default: pack

compile_js:
	rm -f ./frontend/assets/*
	docker run --rm -i \
		-v "$(CURDIR):$(CURDIR)" \
		-w "$(CURDIR)/js" \
		-u $(shell id -u) \
		node:10-alpine \
		sh -c "yarn && npx webpack"

debug:
	go-bindata --debug frontend/...
	go run *.go serve

pack: compile_js
	go-bindata -modtime 1 frontend/...
	bash generateXKCDWordList.sh

publish:
	curl -sSLo golang.sh https://raw.githubusercontent.com/Luzifer/github-publish/master/golang.sh
	bash golang.sh

workflow:
	bash build-workflow.sh
