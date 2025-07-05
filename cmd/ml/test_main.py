import pytest
from fastapi.testclient import TestClient
from main import app, extract_features, LogEntry, train_model

client = TestClient(app)

def test_extract_features():
    log = LogEntry(timestamp="2024-07-05T12:00:00Z", level="ERROR", message="Brute force detected", source="auth")
    features = extract_features(log)
    assert features.shape == (1, 4)
    assert features[0][0] == 2  # level ERROR
    assert features[0][3] == 1  # brute_force flag

def test_train_and_predict():
    logs = [
        LogEntry(timestamp="2024-07-05T12:00:00Z", level="INFO", message="Login ok", source="auth"),
        LogEntry(timestamp="2024-07-05T12:01:00Z", level="ERROR", message="Brute force detected", source="auth"),
    ]
    model = train_model(logs)
    assert model is not None

def test_predict_endpoint_model_not_trained(monkeypatch):
    # Force model=None
    from main import model
    monkeypatch.setattr("main.model", None)
    log = {
        "timestamp": "2024-07-05T12:00:00Z",
        "level": "ERROR",
        "message": "Brute force detected",
        "source": "auth"
    }
    response = client.post("/predict", json=log)
    assert response.status_code == 503

def test_train_and_predict_endpoint():
    logs = [
        {
            "timestamp": "2024-07-05T12:00:00Z",
            "level": "INFO",
            "message": "Login ok",
            "source": "auth"
        },
        {
            "timestamp": "2024-07-05T12:01:00Z",
            "level": "ERROR",
            "message": "Brute force detected",
            "source": "auth"
        }
    ]
    # Train the model
    response = client.post("/train", json=logs)
    assert response.status_code == 200
    # Predict
    response = client.post("/predict", json=logs[1])
    assert response.status_code == 200
    data = response.json()
    assert "anomaly_score" in data
    assert "is_anomaly" in data