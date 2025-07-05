package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"log_sentinel/internal/parser"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	esClient *elasticsearch.Client
	esIndex  string
	logDir   string
)

func initConfig() {
	esIndex = os.Getenv("ELASTIC_INDEX")
	if esIndex == "" {
		esIndex = "logs-sentinel"
	}
	logDir = os.Getenv("LOG_SENTINEL_DIR")
	if logDir == "" {
		logDir = "/var/log/log_sentinel"
	}
}

func initElastic() {
	addr := os.Getenv("ELASTIC_ADDR")
	if addr == "" {
		addr = "http://localhost:9200"
	}
	cfg := elasticsearch.Config{
		Addresses: []string{addr},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("[WARN] ElasticSearch unavailable: %v. Using fallback local.", err)
		esClient = nil
		return
	}
	esClient = client
}

func parseLog(r *http.Request) (*parser.LogEntry, error) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return parser.ParseLog(data)
}

func saveLogToFile(entry *parser.LogEntry) error {
	f, err := os.OpenFile("logs.jsonl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	_, err = f.Write(append(b, '\n'))
	return err
}

func saveLog(entry *parser.LogEntry) error {
	if esClient != nil {
		data, err := json.Marshal(entry)
		if err != nil {
			return err
		}
		res, err := esClient.Index(esIndex, bytes.NewReader(data), esClient.Index.WithContext(context.Background()))
		if err == nil && !res.IsError() {
			defer res.Body.Close()
			return nil
		}
		log.Printf("[WARN] Failed to send log to Elastic: %v. Saving locally.", err)
	}
	return saveLogToFile(entry)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}
	entry, err := parseLog(r)
	if err != nil {
		http.Error(w, "Invalid log format", http.StatusBadRequest)
		return
	}
	err = saveLog(entry)
	if err != nil {
		http.Error(w, "Failed to save log", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func processLogFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("[ERROR] Failed to open log file %s: %v", path, err)
		return
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	for dec.More() {
		var entry parser.LogEntry
		if err := dec.Decode(&entry); err != nil {
			log.Printf("[WARN] Failed to parse line of file %s: %v", path, err)
			continue
		}
		err = saveLog(&entry)
		if err != nil {
			log.Printf("[WARN] Failed to save log of file %s: %v", path, err)
		}
	}
}

func watchLogFiles(dir string) {
	for {
		files, err := filepath.Glob(filepath.Join(dir, "*.jsonl"))
		if err != nil {
			log.Printf("[ERROR] Failed to list files in %s: %v", dir, err)
			continue
		}
		for _, f := range files {
			processLogFile(f)
		}
		time.Sleep(30 * time.Second)
	}
}

func main() {
	initConfig()
	initElastic()
	go watchLogFiles(logDir) // Example of local log directory
	http.HandleFunc("/logs", logHandler)
	log.Println("LogCollector listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
