FROM golang:1.26.0-alpine AS builder

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY src/go.mod src/go.sum* ./
RUN go mod download

COPY src/ ./

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s" \
    -o /app/bin/api \
    ./cmd/api/main.go

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s" \
    -o /app/bin/migrator \
    ./internal/postgres/migrator.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

USER 10001:10001

WORKDIR /app

COPY --from=builder /app/bin/api /app/api
COPY --from=builder /app/bin/migrator /app/migrator

EXPOSE 3000

ENTRYPOINT ["/app/api"]