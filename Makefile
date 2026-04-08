default:

build:
	go build -ldflags "-X github.com/Luzifer/password/v2/pkg/cli.version=$(shell git describe --tags --exclude 'lib/*' --always || echo dev)"

frontend_prod: export NODE_ENV=production
frontend_prod: frontend

frontend: node_modules
	pnpm node ci/build.mjs

frontend_lint: node_modules
	pnpm eslint --fix src

node_modules:
	pnpm install --frozen-lockfile

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
