FROM golang:1.17-alpine AS buildenv
WORKDIR /src
ADD . /src

RUN GOOS=linux go build -o ./out/sha256sum ./cmd/feature_fifth/main.go

RUN chmod +x ./out/sha256sum

FROM alpine:latest
WORKDIR /app
COPY --from=buildenv /src/out/sha256sum .

#### Local application port
EXPOSE 9090

ENTRYPOINT ["/app/sha256sum"]