# ---------- Build Stage ----------
FROM golang:1.26.1 AS builder

WORKDIR /app

# Copy dependency files first
COPY go.mod go.sum* ./

RUN go mod download

# Copy source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -o api-gateway \
    ./cmd/server

# ---------- Runtime Stage ----------
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /

COPY --from=builder /app/api-gateway .

EXPOSE 8084

USER nonroot:nonroot

ENTRYPOINT ["/api-gateway"]