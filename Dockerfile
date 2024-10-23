FROM golang:1.23 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o go-docker ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/go-docker .

CMD ["/app/go-docker"]
