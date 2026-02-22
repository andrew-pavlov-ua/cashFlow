# Build stage
FROM golang:1.25.7-alpine AS builder
RUN apk add --no-cache git

WORKDIR /build
ENV GOWORK=off

# deps cache
COPY services/api-gateway/go.mod services/api-gateway/go.sum ./services/api-gateway/
COPY pkg/go.mod pkg/go.sum ./pkg/
COPY proto/go.mod proto/go.sum ./proto/
RUN cd services/api-gateway && go mod download

# source
COPY services/api-gateway ./services/api-gateway
COPY pkg ./pkg
COPY proto ./proto

# build
WORKDIR /build/services/api-gateway
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags="-w -s" -o /out/api-gateway ./cmd/main.go


# Runtime stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=builder /out/api-gateway ./api-gateway

EXPOSE 8080 50051
CMD ["./api-gateway"]