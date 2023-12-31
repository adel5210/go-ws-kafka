FROM golang:1.21.5-alpine
WORKDIR /go/src/app
COPY . .
RUN go build -o main ./cmd
EXPOSE 8090
CMD ["./main"]