FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY . .
RUN go build -buildvcs=false -tags netgo -trimpath -tags netgo -ldflags="-w -s" -o ./kscale .

FROM scratch

COPY --from=builder --chown=nonroot:nonroot /app/kscale /app/kscale
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/app/kscale"]