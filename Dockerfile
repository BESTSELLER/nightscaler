FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY ./app /app
ENTRYPOINT /app