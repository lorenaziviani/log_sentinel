# Log normal
curl -X POST http://localhost:8080/logs \
  -H 'Content-Type: application/json' \
  -d '{"timestamp":"2024-07-05T12:00:00Z","level":"INFO","message":"User login successful","source":"auth-service"}'

# Log anomaly (simulating attack)
  for i in {1..6}; do
    ts=$(date -u -v+${i}S +"%Y-%m-%dT%H:%M:%SZ")
    curl -X POST http://localhost:8080/logs \
      -H 'Content-Type: application/json' \
      -d "{\"timestamp\":\"$ts\",\"level\":\"ERROR\",\"message\":\"Brute force detected\",\"source\":\"auth-service\"}"
    sleep 0.5
  done