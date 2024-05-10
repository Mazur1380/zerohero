FROM golang:1.22

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main lesson95/*.go

FROM alpine:latest

WORKDIR /app

COPY --from=0 /app/main /app/main

CMD ["/app/main"]