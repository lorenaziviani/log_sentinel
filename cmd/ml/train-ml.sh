#!/bin/bash

# Script to train the ML service of Log Sentinel

curl -X POST http://localhost:8000/train \
  -H 'Content-Type: application/json' \
  -d '[
  {"timestamp": "...", "level": "INFO", "message": "User login successful", "source": "auth-service"},
  {"timestamp": "...", "level": "INFO", "message": "User login successful", "source": "auth-service"},
  {"timestamp": "...", "level": "INFO", "message": "User logout", "source": "auth-service"},
  {"timestamp": "...", "level": "INFO", "message": "User login successful", "source": "auth-service"},
  {"timestamp": "...", "level": "INFO", "message": "User login successful", "source": "auth-service"},
  {"timestamp": "...", "level": "INFO", "message": "User login successful", "source": "auth-service"},
  {"timestamp": "...", "level": "INFO", "message": "User login successful", "source": "auth-service"},
  {"timestamp": "...", "level": "INFO", "message": "User login successful", "source": "auth-service"},
  {"timestamp": "...", "level": "INFO", "message": "User login successful", "source": "auth-service"},
  {"timestamp": "...", "level": "INFO", "message": "User login successful", "source": "auth-service"},
  {"timestamp": "...", "level": "ERROR", "message": "Brute force detected", "source": "auth-service"},
  {"timestamp": "...", "level": "ERROR", "message": "Brute force detected", "source": "auth-service"}
]' 