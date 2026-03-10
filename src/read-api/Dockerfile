ARG GO_VERSION=1.26
# ---- Build Stage ----
FROM golang:${GO_VERSION} AS builder

WORKDIR /src

# Copy go.mod and go.sum first for caching
COPY src/read-api/go.mod src/read-api/go.sum /src/read-api/

# Shared source
COPY src/shared/. /src/shared/
RUN cd /src/read-api && go mod download

# Copy the rest of the source
COPY src/read-api/. /src/read-api/

# Build static binary
RUN cd /src/read-api && \
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
