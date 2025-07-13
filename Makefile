default:

build:
	go build -ldflags "-X main.version=$(git describe --tags --always || echo dev)"

frontend_prod: export NODE_ENV=production
frontend_prod: frontend

frontend: node_modules
	corepack yarn@1 node ci/build.mjs

frontend_lint: node_modules
	corepack yarn@1 eslint --fix src

node_modules:
	corepack yarn@1 install --production=false --frozen-lockfile

publish: frontend
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
