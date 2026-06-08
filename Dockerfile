FROM golang:1.26.4-alpine@sha256:f23e8b227fb4493eabe03bede4d5a32d04092da71962f1fb79b5f7d1e6c2a17f AS builder

COPY --from=ghcr.io/luzifer-docker/pnpm:v11.5.2@sha256:46d369a8379ee3cd52fca61147e5be818073896c25127407df2b2aca1ec10a9f . /

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
