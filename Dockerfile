
FROM golang:1.23 AS build

WORKDIR /app

COPY . .

RUN go build -o /go-docker

FROM alpine:latest

WORKDIR /app

COPY --from=build /go-docker .

CMD ["/app/go-docker"]
