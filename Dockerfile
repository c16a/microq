# Stage 1: Build the Go application (compiler and dependencies)
FROM golang:1.22.2 AS builder

WORKDIR /app

# Copy your Go source code
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w" -o microq github.com/c16a/microq

# Stage 2: Production image (slim and optimized)
FROM scratch

WORKDIR /app

# Copy only the built binary from the builder stage
COPY --from=builder /app/microq ./

# Expose the port your application listens on (replace 8080 with your actual port)
EXPOSE 8080

# Set the command to execute your binary
CMD ["microq"]
