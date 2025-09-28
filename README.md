Simple Go HTTP server

This repository contains a tiny Go web server located in `cmd/server`.

To run:

1. Build

   go build -o bin/server ./cmd/server

2. Run

   ./bin/server

The server listens on :8080 with one endpoint:

- GET / -> "hello world"
