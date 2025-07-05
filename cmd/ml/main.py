import os
import joblib
import numpy as np
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from sklearn.ensemble import IsolationForest

MODEL_PATH = os.getenv("MODEL_PATH", "model.joblib")

app = FastAPI(title="Log Sentinel AnomalyDetector")

class LogEntry(BaseModel):
    timestamp: str
    level: str
    message: str
    source: str

class AnomalyResponse(BaseModel):
    anomaly_score: float
    is_anomaly: bool

# --- Model and features ---
model = None

def extract_features(log: LogEntry):
    # Simple example: message length, encoded level, hour of day
    level_map = {"INFO": 0, "WARN": 1, "ERROR": 2, "DEBUG": 3}
    level = level_map.get(log.level.upper(), -1)
    msg_len = len(log.message)
    hour = int(log.timestamp[11:13]) if len(log.timestamp) > 12 else 0
    return np.array([[level, msg_len, hour]])

def train_model(logs):
    X = np.vstack([extract_features(log) for log in logs])
    clf = IsolationForest(contamination=0.05, random_state=42)
    clf.fit(X)
    joblib.dump(clf, MODEL_PATH)
    return clf

def load_model():
    global model
    if os.path.exists(MODEL_PATH):
        model = joblib.load(MODEL_PATH)
    else:
        model = None

@app.on_event("startup")
def startup_event():
    load_model()

@app.post("/predict", response_model=AnomalyResponse)
def predict(log: LogEntry):
    if model is None:
        raise HTTPException(status_code=503, detail="Model not trained")
    X = extract_features(log)
    score = -model.decision_function(X)[0]  # the higher, the more anomalous
    is_anomaly = model.predict(X)[0] == -1
    return AnomalyResponse(anomaly_score=score, is_anomaly=is_anomaly)

@app.post("/train")
def train(logs: list[LogEntry]):
    global model
    model = train_model(logs)
    return {"status": "trained", "n_logs": len(logs)}

# --- Manual usage example ---
if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True) 