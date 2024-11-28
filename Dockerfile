# Build stage
FROM golang:1.22 AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o fortanix-csi-provider

# Final stage
FROM alpine:3.20
RUN apk update && \
    apk upgrade --no-cache libcrypto3
COPY --from=builder /build/fortanix-csi-provider /bin/
ENTRYPOINT ["/bin/fortanix-csi-provider"]
