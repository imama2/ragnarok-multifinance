FROM golang:1.22-alpine AS builder
RUN apk add --no-cache git build-base
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -o ragnarok .

FROM alpine:3.19
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /app/ragnarok /usr/local/bin/
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=3s \
  CMD curl -f http://localhost:8080/health || exit 1
CMD ["ragnarok"] 