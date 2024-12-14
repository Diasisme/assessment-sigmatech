FROM golang:1.23.2-alpine

# Install tool for auto compile and build local program
RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

ENTRYPOINT CompileDaemon --build="go build -o svc-acc" -command="./svc-acc" -build-dir=/app