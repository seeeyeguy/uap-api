# ── Stage 1: base Go environment ──────────────────────────────
FROM golang:1.24-alpine AS base
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download

# ── Stage 2: dev (hot reload with air) ────────────────────────
FROM base AS dev
RUN go install github.com/air-verse/air@v1.61.1
COPY . .
CMD ["air", "-c", ".air.toml"]

# ── Stage 3: build the binary ─────────────────────────────────
FROM base AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server

# ── Stage 4: prod (minimal image, just the binary) ────────────
FROM alpine:3.19 AS prod
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
