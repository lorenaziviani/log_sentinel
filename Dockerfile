# Build stage
FROM golang:1.24.3 AS builder
WORKDIR /app

# Copy the entire project structure first
COPY . .

# Set working directory to the collector
WORKDIR /app/cmd/collector

# Download dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o collector .

# Run stage
FROM gcr.io/distroless/static-debian12
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/cmd/collector/collector /app/collector

# Copy environment sample file
COPY cmd/collector/.env.sample /app/.env.sample

# Expose port
EXPOSE 8080

# Set entrypoint
ENTRYPOINT ["/app/collector"]