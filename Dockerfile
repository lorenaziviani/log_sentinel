# Build stage
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN cd cmd/collector && CGO_ENABLED=0 go build -o /collector main.go

# Run stage
FROM gcr.io/distroless/static
COPY --from=builder /collector /collector
COPY --from=builder /app/cmd/collector/.env.sample /app/.env.sample
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/collector"] 