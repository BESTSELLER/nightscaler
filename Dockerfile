FROM golang:1.20-alpine AS builder

ARG UPX_VERSION=4.0.2-r0

RUN apk update && apk add --no-cache ca-certificates upx=$UPX_VERSION tzdata && update-ca-certificates

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 64000 \
  kscale

COPY . .
RUN go build -buildvcs=false -tags netgo -trimpath -tags netgo -ldflags="-w -s" -o ./kscale .

RUN upx --best --lzma kscale

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder --chown=kscale:kscale /app/kscale /app/kscale

USER kscale:kscale

ENTRYPOINT ["/app/kscale"]