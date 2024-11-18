# Stage 1: Build
FROM golang:1.23.3-bullseye AS builder

# Set environment variables for Go
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules manifests first and download dependencies
# This helps in caching dependencies if they don't change
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main

# Stage 2: Run
FROM scratch

WORKDIR /app
# Copy the statically compiled binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/data.yaml .
COPY --from=builder /app/template ./template

# Expose the port your app runs on
EXPOSE 8000

# Command to run the application
CMD ["./main"]
