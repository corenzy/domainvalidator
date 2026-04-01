# --- Build Stage ---
FROM golang:1.22-alpine AS builder

# Set working directory
WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
# CGO_ENABLED=0 results in a static binary that can run on alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# --- Final Stage ---
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Install ca-certificates for secure DNS lookups
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 3000

# Run the application
CMD ["./main"]
