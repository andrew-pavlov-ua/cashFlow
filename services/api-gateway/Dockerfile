#base go image

FROM golang:1.25.7-alpine as builder

RUN mkdir /gateway

COPY ./ /gateway

WORKDIR /gateway

RUN CGO_ENABLED=0 go build -o zero ./cmd

RUN chmod +x ./zero

#build a tiny docker image
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

RUN mkdir /gateway

COPY --from=builder /gateway /gateway

WORKDIR /gateway

CMD ["/gateway/zero"]
