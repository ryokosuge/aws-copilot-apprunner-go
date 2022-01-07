FROM golang:1.17-buster as bulder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY cmd/app/main.go ./
RUN go clean -cache && go build -o /server

FROM alpine:3

WORKDIR /app
COPY --from=bulder /server /app/server

RUN set -eux && \
	mkdir /lib64 &&  \
	ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

RUN set -eux && \
	addgroup app && \
	adduser -D -G app app && \
	chown -R app:app /app && \
	chmod 500 /app/server

USER app

CMD ["/app/server"]
