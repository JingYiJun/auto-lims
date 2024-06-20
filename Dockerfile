FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN apk add --no-cache --virtual .build-deps \
        ca-certificates \
    go mod download

COPY . .

RUN go build -ldflags "-s -w" -o service

FROM alpine

WORKDIR /app

COPY --from=builder /app/service /app/

EXPOSE 8000

ENTRYPOINT ["./service"]