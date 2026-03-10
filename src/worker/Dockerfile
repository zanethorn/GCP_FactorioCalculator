ARG GO_VERSION=1.26
# ---- Build Stage ----
FROM golang:${GO_VERSION} AS builder

WORKDIR /src

# Copy go.mod and go.sum first for caching
COPY src/worker/go.mod src/worker/go.sum /src/worker/

# Shared source
COPY src/shared/. /src/shared/
RUN cd /src/worker && go mod download

# Copy the rest of the source
COPY src/worker/. /src/worker/

# Build static binary
RUN cd /src/worker && \
    CGO_ENABLED=0 \
    GOOS=linux GOARCH=amd64 \
    go build -o /src/worker ./cmd/worker

# ---- Runtime Stage ----
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /src/worker .

# Cloud Run listens on $PORT
ENV PORT=8080
EXPOSE 8080

CMD ["/app/worker"]
