FROM golang:1.13-alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /go/bin/main
#expose port
EXPOSE 9090
ENTRYPOINT ["/go/bin/main"]