package main

import (
	"errors"
	"log_sentinel/cmd/collector/internal/parser"
	"testing"
	"time"
)

// Mocks globais
var savedAlert *parser.LogEntry
var saveLogCalled bool
var savedAnomaly *parser.LogEntry
var saveAnomalyCalled bool

func mockSaveLog(entry *parser.LogEntry) error {
	saveLogCalled = true
	savedAlert = entry
	if entry.Level == "ALERT" {
		return nil
	}
	return errors.New("not an alert")
}

func mockSaveAnomaly(entry *parser.LogEntry, score float64) error {
	saveAnomalyCalled = true
	savedAnomaly = entry
	return nil
}

func resetMocks() {
	saveLogCalled = false
	savedAlert = nil
	saveAnomalyCalled = false
	savedAnomaly = nil
}

func TestAlertLogCreation_Mock(t *testing.T) {
	resetMocks()
	origSaveLog := saveLog
	defer func() { saveLog = origSaveLog }()
	saveLog = mockSaveLog

	anomalyCount = alertRule.Threshold
	alertRule.LastAlert = time.Now().Add(-2 * alertRule.Window)

	entry := &parser.LogEntry{
		Timestamp: time.Now(),
		Level:     "ERROR",
		Message:   "Brute force detected",
		Source:    "auth-service",
	}
	processLogWithAnomaly(entry)

	if !saveLogCalled {
		t.Fatalf("saveLog not called for ALERT")
	}
	if savedAlert == nil || savedAlert.Level != "ALERT" {
		t.Fatalf("Alert log not persisted correctly: %+v", savedAlert)
	}
	if savedAlert.Source != "log-collector" {
		t.Errorf("Expected source 'log-collector', got: %s", savedAlert.Source)
	}
	if savedAlert.Message == "" {
		t.Errorf("Alert message not filled")
	}
}

func TestNoAlertBelowThreshold(t *testing.T) {
	resetMocks()
	origSaveLog := saveLog
	defer func() { saveLog = origSaveLog }()
	saveLog = mockSaveLog

	anomalyCount = alertRule.Threshold - 1
	alertRule.LastAlert = time.Now().Add(-2 * alertRule.Window)

	entry := &parser.LogEntry{
		Timestamp: time.Now(),
		Level:     "ERROR",
		Message:   "Brute force detected",
		Source:    "auth-service",
	}
	processLogWithAnomaly(entry)

	if saveLogCalled {
		t.Errorf("saveLog should not be called for ALERT below threshold")
	}
}

func TestNoAlertIfWindowNotExpired(t *testing.T) {
	resetMocks()
	origSaveLog := saveLog
	defer func() { saveLog = origSaveLog }()
	saveLog = mockSaveLog

	anomalyCount = alertRule.Threshold
	alertRule.LastAlert = time.Now() // window not expired

	entry := &parser.LogEntry{
		Timestamp: time.Now(),
		Level:     "ERROR",
		Message:   "Brute force detected",
		Source:    "auth-service",
	}
	processLogWithAnomaly(entry)

	if saveLogCalled {
		t.Errorf("saveLog should not be called for ALERT if window not expired")
	}
}

func TestAnomalyPersistence_Mock(t *testing.T) {
	resetMocks()
	origSaveAnomaly := saveAnomaly
	defer func() { saveAnomaly = origSaveAnomaly }()
	saveAnomaly = mockSaveAnomaly

	entry := &parser.LogEntry{
		Timestamp: time.Now(),
		Level:     "ERROR",
		Message:   "Brute force detected",
		Source:    "auth-service",
	}
	// Simulate checkAnomaly returning true (anomaly)
	anomalyTotal.Inc()
	anomalyCount++
	err := saveAnomaly(entry, 0.99)
	if err != nil {
		t.Fatalf("Error saving anomaly: %v", err)
	}
	if !saveAnomalyCalled {
		t.Fatalf("saveAnomaly not called")
	}
	if savedAnomaly == nil || savedAnomaly.Message == "" {
		t.Errorf("Anomaly not persisted correctly: %+v", savedAnomaly)
	}
}

func TestSuite(t *testing.T) {
	t.Run("TestAlertLogCreation_Mock", TestAlertLogCreation_Mock)
	t.Run("TestNoAlertBelowThreshold", TestNoAlertBelowThreshold)
	t.Run("TestNoAlertIfWindowNotExpired", TestNoAlertIfWindowNotExpired)
	t.Run("TestAnomalyPersistence_Mock", TestAnomalyPersistence_Mock)
}
