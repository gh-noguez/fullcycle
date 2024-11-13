FROM golang:1.18-alpine AS builder

WORKDIR /app

RUN go mod init main

COPY main.go .

RUN go build -o main .

FROM scratch

COPY --from=builder /app .

CMD ["./main"]