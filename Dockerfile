FROM golang:1.27.1-alpine@sha256:8a5910f31396cd4d89662f56c68b3ae31d374308270a1c3bd96672ee5ed43414 AS builder

COPY --from=ghcr.io/luzifer-docker/pnpm:v11.28.3@sha256:a1e0b2517e1d1271cb46f36614fc39bec24f47b59a79bb67756bcfd4cfb98143 . /

COPY . /src/password
WORKDIR /src/password

RUN set -ex \
 && apk add --no-cache \
      build-base \
      git \
      nodejs-current \
 && make frontend_prod

RUN set -ex \
 && go install -ldflags "-X github.com/Luzifer/password/v2/pkg/cli.version=$(git describe --tags --exclude 'lib/*' --always || git rev-parse --short HEAD || echo dev)"


FROM alpine:3.24.2@sha256:294b683cb724975bec92580e1e685676bd4b50bda910ddb8c51d4cabeaec77e6

LABEL org.opencontainers.image.authors="Knut Ahlers <knut@ahlers.me>" \
      org.opencontainers.image.documentation="https://github.com/Luzifer/password"

RUN set -ex \
 && apk --no-cache add \
      ca-certificates \
      mailcap

COPY --from=builder /go/bin/password /usr/local/bin/password

EXPOSE 3000

USER 1000:1000

ENTRYPOINT ["/usr/local/bin/password"]
