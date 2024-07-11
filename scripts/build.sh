#!/bin/bash

set -e

# Set the target OS and architecture
GOOS=linux GOARCH=amd64 go build -o main cmd/lambda/main.go
