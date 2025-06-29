package main

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestMetricsSnapshot(t *testing.T) {
	app := &appledocs{
		startTime:   time.Now().Add(-5 * time.Second), // Started 5 seconds ago
		httpErrors:  make(map[int]int),
		rateLimiter: rate.NewLimiter(rate.Inf, 0),
	}

	// Simulate some activity
	app.recordResponseTime(100 * time.Millisecond)
	app.recordResponseTime(200 * time.Millisecond)
	app.recordBytesDownloaded(1024 * 1024) // 1 MB
	app.recordBytesFromCache(512 * 1024)   // 512 KB
	app.recordHTTPError(404)
	app.recordHTTPError(500)
	app.recordHTTPError(404) // Another 404
	app.incrementRetryCount()
	app.incrementCacheHits()
	app.incrementCacheMisses()
	app.recordContentType("tutorials/data/documentation/SwiftUI.json")       // Framework
	app.recordContentType("tutorials/data/documentation/SwiftUI/View.json")  // Class
	app.recordContentType("tutorials/data/documentation/SwiftUI/View/init.json") // Method

	metrics := app.getEnhancedMetrics()

	// Verify basic metrics structure
	if metrics.StartTime.IsZero() {
		t.Error("StartTime should be set")
	}

	if metrics.RuntimeDuration <= 0 {
		t.Error("RuntimeDuration should be positive")
	}

	if metrics.AvgResponseTime != 150*time.Millisecond {
		t.Errorf("Expected avg response time 150ms, got %v", metrics.AvgResponseTime)
	}

	if metrics.TotalBytesDownloaded != 1024*1024 {
		t.Errorf("Expected 1MB downloaded, got %d", metrics.TotalBytesDownloaded)
	}

	if metrics.TotalBytesFromCache != 512*1024 {
		t.Errorf("Expected 512KB from cache, got %d", metrics.TotalBytesFromCache)
	}

	if metrics.RetryCount != 1 {
		t.Errorf("Expected 1 retry, got %d", metrics.RetryCount)
	}

	if metrics.FrameworkCount != 1 {
		t.Errorf("Expected 1 framework, got %d", metrics.FrameworkCount)
	}

	if metrics.ClassCount != 1 {
		t.Errorf("Expected 1 class, got %d", metrics.ClassCount)
	}

	if metrics.MethodCount != 1 {
		t.Errorf("Expected 1 method, got %d", metrics.MethodCount)
	}

	// Check HTTP errors
	if len(metrics.HTTPErrors) != 2 {
		t.Errorf("Expected 2 different HTTP error types, got %d", len(metrics.HTTPErrors))
	}

	if metrics.HTTPErrors[404] != 2 {
		t.Errorf("Expected 2 x 404 errors, got %d", metrics.HTTPErrors[404])
	}

	if metrics.HTTPErrors[500] != 1 {
		t.Errorf("Expected 1 x 500 error, got %d", metrics.HTTPErrors[500])
	}

	// Check cache hit rate
	expectedHitRate := float64(1) / float64(2) * 100 // 1 hit / (1 hit + 1 miss) * 100
	if metrics.CacheHitRate != expectedHitRate {
		t.Errorf("Expected cache hit rate %.1f%%, got %.1f%%", expectedHitRate, metrics.CacheHitRate)
	}
}

func TestMetricsJSONSerialization(t *testing.T) {
	app := &appledocs{
		startTime:   time.Now(),
		httpErrors:  make(map[int]int),
		rateLimiter: rate.NewLimiter(rate.Inf, 0),
	}

	app.recordResponseTime(100 * time.Millisecond)
	app.recordBytesDownloaded(1024)
	app.recordHTTPError(404)

	metrics := app.getEnhancedMetrics()

	// Test JSON serialization
	jsonData, err := json.MarshalIndent(metrics, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal metrics to JSON: %v", err)
	}

	// Test JSON deserialization
	var deserializedMetrics MetricsSnapshot
	if err := json.Unmarshal(jsonData, &deserializedMetrics); err != nil {
		t.Fatalf("Failed to unmarshal metrics from JSON: %v", err)
	}

	// Verify key fields are preserved
	if deserializedMetrics.TotalBytesDownloaded != metrics.TotalBytesDownloaded {
		t.Error("TotalBytesDownloaded not preserved in JSON serialization")
	}

	if deserializedMetrics.AvgResponseTime != metrics.AvgResponseTime {
		t.Error("AvgResponseTime not preserved in JSON serialization")
	}

	if len(deserializedMetrics.HTTPErrors) != len(metrics.HTTPErrors) {
		t.Error("HTTPErrors not preserved in JSON serialization")
	}
}

func TestMetricsExport(t *testing.T) {
	app := &appledocs{
		startTime:   time.Now(),
		httpErrors:  make(map[int]int),
		rateLimiter: rate.NewLimiter(rate.Inf, 0),
	}

	app.recordBytesDownloaded(1024)
	app.recordHTTPError(404)

	metrics := app.getEnhancedMetrics()

	// Test exporting to a temporary file
	tempFile := t.TempDir() + "/test_metrics.json"

	metricsJSON, err := json.MarshalIndent(metrics, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal metrics: %v", err)
	}

	if err := os.WriteFile(tempFile, metricsJSON, 0644); err != nil {
		t.Fatalf("Failed to write metrics file: %v", err)
	}

	// Verify file was created and has content
	fileInfo, err := os.Stat(tempFile)
	if err != nil {
		t.Fatalf("Metrics file was not created: %v", err)
	}

	if fileInfo.Size() == 0 {
		t.Error("Metrics file is empty")
	}

	// Verify content can be read back
	fileContent, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("Failed to read metrics file: %v", err)
	}

	var readBackMetrics MetricsSnapshot
	if err := json.Unmarshal(fileContent, &readBackMetrics); err != nil {
		t.Fatalf("Failed to parse metrics file: %v", err)
	}

	if readBackMetrics.TotalBytesDownloaded != 1024 {
		t.Error("Metrics not correctly saved to file")
	}
}