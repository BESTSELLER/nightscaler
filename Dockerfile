FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

RUN go mod download
RUN go build -o ./kscale .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/kscale /go/bin/kscale

ENTRYPOINT ["/go/bin/kscale"]