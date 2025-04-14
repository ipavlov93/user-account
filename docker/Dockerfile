# Start from a small Go base image to build the app
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o bookly-api ./cmd

# --- Runtime image ---
FROM alpine:latest

# Optional: Install ca-certificates if your app makes HTTPS calls
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the built binary from builder
COPY --from=builder /app/bookly-api .

# Copy auth private key
COPY account.json .
ENV GOOGLE_APPLICATION_CREDENTIALS=/app/account.json

# Run the app
ENTRYPOINT ["/app/bookly-api"]
