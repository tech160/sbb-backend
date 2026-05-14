# ── Stage 1: builder ──────────────────────────────────────────
FROM golang:1.23-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
      -ldflags="-s -w" \
      -o server ./cmd/main.go

# ── Stage 2: runner (distroless — no shell, minimal attack surface) ──
FROM gcr.io/distroless/static-debian12 AS runner
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/server"]
