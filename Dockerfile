FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o ./kscale .

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/kscale .

ENTRYPOINT ["/app/kscale"]