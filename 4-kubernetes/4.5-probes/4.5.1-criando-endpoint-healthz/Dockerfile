FROM golang:1.22.2-alpine
RUN go mod init server
COPY . .
RUN go build -o server .
CMD ["./server"]
