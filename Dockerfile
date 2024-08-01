FROM golang:1.22-alpine
WORKDIR /usr/local/src
COPY go.* .
RUN go mod download
COPY main.go .
RUN go build -v ./...
EXPOSE 80
ENTRYPOINT ["./adafruit-sht4x-trinkey-exporter"]
