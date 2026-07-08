FROM golang:1.26.4-alpine@sha256:3ad57304ad93bbec8548a0437ad9e06a455660655d9af011d58b993f6f615648 AS builder

COPY --from=ghcr.io/luzifer-docker/pnpm:v11.10.0@sha256:8abf5ad6d97ee752f9f4fdaf3cb74aeb1fec70585deaa7f5e9538043af4c94fe . /

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


FROM alpine:3.24.1@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b

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
