package main

import (
    "encoding/json"
    "fmt"
    "log/slog"
    "net/http"
    "os"
    "time"
)

var startTime time.Time
var logger *slog.Logger

type InfoResponse struct {
    Service   string  `json:"service"`
    Message   string  `json:"message"`
    Timestamp string  `json:"timestamp"`
    PodName   string  `json:"pod_name"`
    Uptime    float64 `json:"uptime"`
}

func init() {
    startTime = time.Now()
    logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}

func main() {
    port := getEnv("PORT", "5000")
    podName := getEnv("HOSTNAME", "unknown")
    
    logger.Info("Starting backend server", "port", port, "pod_name", podName)
    
    http.HandleFunc("/api/info", infoHandler)
    http.HandleFunc("/health", healthHandler)
    
    logger.Info("Server started", "address", fmt.Sprintf(":%s", port))
    http.ListenAndServe(":"+port, nil)
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
    response := InfoResponse{
        Service:   "backend",
        Message:   "Hello from Kubernetes!",
        Timestamp: time.Now().Format(time.RFC3339),
        PodName:   getEnv("HOSTNAME", "unknown"),
        Uptime:    time.Since(startTime).Seconds(),
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}