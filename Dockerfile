FROM golang:1.26-alpine@sha256:c2a1f7b2095d046ae14b286b18413a05bb82c9bca9b25fe7ff5efef0f0826166 AS builder

COPY --from=ghcr.io/luzifer-docker/pnpm:v10.33.0@sha256:37421bf6d0c9bb40c8fc3471b2ed1f2e7947b4668692713c7db8877feb840d8a . /

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


FROM alpine:3.23@sha256:5b10f432ef3da1b8d4c7eb6c487f2f5a8f096bc91145e68878dd4a5019afde11

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
