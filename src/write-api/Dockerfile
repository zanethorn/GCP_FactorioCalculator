ARG GO_VERSION=1.26
# ---- Build Stage ----
FROM golang:${GO_VERSION} AS builder

WORKDIR /src

# Copy go.mod and go.sum first for caching
COPY src/write-api/go.mod src/write-api/go.sum /src/write-api/

# Shared source
COPY src/shared/. /src/shared/
RUN cd /src/write-api && go mod download

# Copy the rest of the source
COPY src/write-api/. /src/write-api/

# Build static binary
RUN cd /src/write-api && \
    CGO_ENABLED=0 \
    GOOS=linux GOARCH=amd64 \
    go build -o /src/server ./cmd/server

# ---- Runtime Stage ----
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /src/server .

# Cloud Run listens on $PORT
ENV PORT=8080
EXPOSE 8080

CMD ["/app/server"]
