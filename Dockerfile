FROM golang:alpine as builder

COPY . /src/github.com/Luzifer/password
WORKDIR /src/github.com/Luzifer/password

RUN set -ex \
 && apk add --no-cache \
      build-base \
      git \
      nodejs \
      npm \
 && make compile_js

WORKDIR /src/github.com/Luzifer/password/cmd/password

RUN set -ex \
 && go install -ldflags "-X main.version=$(git describe --tags || git rev-parse --short HEAD || echo dev)"


FROM alpine:latest

LABEL maintainer "Knut Ahlers <knut@ahlers.me>"

RUN set -ex \
 && apk --no-cache add \
      ca-certificates \
      mailcap

COPY --from=builder /go/bin/password /usr/local/bin/password

EXPOSE 3000

ENTRYPOINT ["/usr/local/bin/password"]
CMD ["--"]

# vim: set ft=Dockerfile:
