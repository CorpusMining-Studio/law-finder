FROM golang:1.23-alpine3.20 AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest

COPY . .

COPY --from=builder /app/main .

EXPOSE 8001

CMD ["./main"]
