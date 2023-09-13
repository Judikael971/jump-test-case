FROM golang:1.21 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o technical_case cmd/technical_case/*.go


FROM alpine:3.18.3 as app
WORKDIR /app

RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Europe/Paris /etc/localtime \
    && echo "Europe/Paris" > /etc/timezone \
    && apk del tzdata \
    && mkdir -p /app/log \
    && chown -R nobody:nobody /app

COPY --from=builder /app/technical_case /app

USER nobody:nobody
ENTRYPOINT ["./technical_case"]