default:

compile_js: js/node_modules
	cd js && node build.mjs

js/node_modules:
	cd js && npm ci

publish:
	curl -sSLo golang.sh https://raw.githubusercontent.com/Luzifer/github-publish/master/golang.sh
	bash golang.sh
