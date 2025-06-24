FROM golang:1.23.1-alpine

# Install necessary packages
RUN apk add --no-cache bash curl mysql-client

# Install migrate tool
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Set environment for testing
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64