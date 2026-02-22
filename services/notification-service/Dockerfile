# Build stage
FROM golang:1.25.7-alpine AS builder
RUN apk add --no-cache git

WORKDIR /build
ENV GOWORK=off

# Cache deps (copy only go.mod/go.sum first)
COPY services/notification-service/go.mod services/notification-service/go.sum ./services/notification-service/
COPY pkg/go.mod pkg/go.sum ./pkg/
COPY proto/go.mod proto/go.sum ./proto/

RUN cd services/notification-service && go mod download

# Copy source
COPY services/notification-service ./services/notification-service
COPY pkg ./pkg
COPY proto ./proto

# Optional: keep module files consistent (can be removed if you don't want it in docker build)
RUN cd services/notification-service && go mod tidy

# Build binary
WORKDIR /build/services/notification-service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags="-w -s" -o /out/notification-service ./cmd/main.go


# Runtime stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=builder /out/notification-service ./notification-service

CMD ["./notification-service"]