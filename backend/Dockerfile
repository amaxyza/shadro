FROM golang:1.23 AS builder
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/server .
COPY templates/ templates/
COPY static/ static/

EXPOSE 8081
ENTRYPOINT ["./server"]