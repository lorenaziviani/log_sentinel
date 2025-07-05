#!/bin/bash

# Script para testar o endpoint /predict do serviço de ML do Log Sentinel

curl -X POST http://localhost:8000/predict \
  -H 'Content-Type: application/json' \
  -d '{"timestamp":"2024-07-05T12:01:00Z","level":"ERROR","message":"Brute force detected","source":"auth-service"}' 