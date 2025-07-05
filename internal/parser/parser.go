package parser

import (
	"encoding/json"
	"errors"
	"time"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Source    string    `json:"source"`
}

func ParseLog(data []byte) (*LogEntry, error) {
	var entry LogEntry
	err := json.Unmarshal(data, &entry)
	if err != nil {
		return nil, err
	}
	if entry.Timestamp.IsZero() || entry.Level == "" || entry.Message == "" || entry.Source == "" {
		return nil, errors.New("missing required log fields")
	}
	return &entry, nil
}
