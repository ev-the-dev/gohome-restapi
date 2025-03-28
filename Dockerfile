FROM docker.io/golang:1.24 AS builder
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o restapi .

FROM docker.io/alpine:latest
RUN mkdir /app && adduser -h /app -D restapi
WORKDIR /app
COPY --chown=restapi --from=builder /app/restapi .
EXPOSE 4000
CMD ["/app/restapi"]
