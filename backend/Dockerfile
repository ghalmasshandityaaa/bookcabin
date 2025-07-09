# ================================================
# BUILD STAGE - Compile the Go application
# ================================================
FROM golang:1.24-alpine AS builder

# Install minimal required dependencies
RUN apk add --no-cache git ca-certificates

# Create unprivileged user for security (will be used in runtime)
RUN adduser -D -g '' appuser

WORKDIR /app

# Copy dependency files first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy all source files
COPY . .

# Build a statically linked binary with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-s -w -extldflags "-static"' \
    -trimpath \
    -tags netgo \
    -o /main cmd/main.go

# ================================================
# RUNTIME STAGE - Minimal scratch image
# ================================================
FROM scratch

# Import from builder:
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /main .
COPY --from=builder /app/config/ ./config/
COPY --from=builder --chown=appuser:appuser /app/database/ ./database/

# Runtime environment
ENV TZ=Asia/Jakarta

# Run as unprivileged user
USER appuser

# Expose default port (adjust if your app uses different port)
EXPOSE 3000

# Entrypoint (use exec form for proper signal handling)
ENTRYPOINT ["/main"]