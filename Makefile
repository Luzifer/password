default:

compile_js:
	rm -f ./frontend/assets/*
	docker run --rm -i \
		-v "$(CURDIR):$(CURDIR)" \
		-w "$(CURDIR)/js" \
		-u $(shell id -u) \
		node:18-alpine \
		sh -c "node build.mjs"

publish: compile_js
	curl -sSLo golang.sh https://raw.githubusercontent.com/Luzifer/github-publish/master/golang.sh
	bash golang.sh
