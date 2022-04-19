FROM golang:1.18-alpine AS buildenv
WORKDIR /src
ADD . /src

RUN GOOS=linux go build -o sha256sum cmd/feature_fifth/main.go

RUN chmod +x sha256sum

FROM alpine:latest
WORKDIR /app
COPY --from=buildenv /src/sha256sum .
COPY --from=buildenv /src/config.yaml ./
#### Local application port
EXPOSE 9090

ENTRYPOINT ["/app/sha256sum"]