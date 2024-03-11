FROM golang:1.21
WORKDIR /app
COPY main.go .
RUN go build -o hello-go main.go
CMD ["./hello-go"]