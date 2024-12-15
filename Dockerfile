FROM golang:1.23.2-alpine

# Install tool for auto compile and build local program
RUN go install github.com/githubnemo/CompileDaemon@latest

# Tambahkan dukungan debugging
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Gunakan perintah untuk menjalankan dengan debugger
CMD ["dlv", "debug", "--headless", "--listen=:40000", "--api-version=2", "--accept-multiclient", "./main.go"]

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

ENTRYPOINT CompileDaemon --build="go build -o svc-acc" -command="./svc-acc" -build-dir=/app