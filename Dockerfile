FROM golang:1.25-alpine@sha256:72567335df90b4ed71c01bf91fb5f8cc09fc4d5f6f21e183a085bafc7ae1bec8 AS builder

COPY . /src/github.com/Luzifer/password
WORKDIR /src/github.com/Luzifer/password

RUN set -ex \
 && apk add --no-cache \
      build-base \
      git \
      nodejs-current \
 && make frontend_prod

RUN set -ex \
 && go install -ldflags "-X github.com/Luzifer/password/v2/pkg/cli.version=$(git describe --tags --exclude 'lib/*' --always || git rev-parse --short HEAD || echo dev)"


FROM alpine:3.23@sha256:865b95f46d98cf867a156fe4a135ad3fe50d2056aa3f25ed31662dff6da4eb62

LABEL org.opencontainers.image.authors="Knut Ahlers <knut@ahlers.me>" \
      org.opencontainers.image.documentation="https://github.com/Luzifer/password"

RUN set -ex \
 && apk --no-cache add \
      ca-certificates \
      mailcap

COPY --from=builder /go/bin/password /usr/local/bin/password

EXPOSE 3000

ENTRYPOINT ["/usr/local/bin/password"]
CMD ["--"]

# vim: set ft=Dockerfile:
