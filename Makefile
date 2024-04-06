default:

compile_js: js/node_modules
	cd js && node build.mjs

js/node_modules:
	cd js && npm ci

publish: compile_js
	bash ci/build.sh

trivy:
	trivy fs . \
		--dependency-tree \
		--exit-code 1 \
		--format table \
		--ignore-unfixed \
		--quiet \
		--scanners misconfig,license,secret,vuln \
		--severity HIGH,CRITICAL \
		--skip-dirs docs
