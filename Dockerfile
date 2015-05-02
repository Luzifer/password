FROM golang

MAINTAINER Knut Ahlers <knut@ahlers.me>

RUN go get github.com/Luzifer/password && \
    go install github.com/Luzifer/password

EXPOSE 3000

ENTRYPOINT ["/go/bin/password"]
CMD ["serve", "--port", "3000"]
