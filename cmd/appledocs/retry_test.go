package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestFetchWithCacheRetry(t *testing.T) {
	// Test retry logic for server errors
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 3 {
			// Fail first two attempts
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Succeed on third attempt
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"test": "data"}`))
	}))
	defer server.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	app := &crawler{
		client:      client,
		visitedURLs: sync.Map{},
		badURLs:     make(map[string]bool),
		rateLimiter: rate.NewLimiter(rate.Inf, 0), // No limit for tests
	}

	ctx := context.Background()
	data, err := fetchWithCache(ctx, client, server.URL, app)

	if err != nil {
		t.Fatalf("Expected successful fetch after retries, got error: %v", err)
	}

	if string(data) != `{"test": "data"}` {
		t.Errorf("Expected test data, got: %s", string(data))
	}

	if attempts != 3 {
		t.Errorf("Expected 3 attempts, got %d", attempts)
	}
}

func TestFetchWithCacheRetryTimeout(t *testing.T) {
	// Test that non-retryable errors fail immediately
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	app := &crawler{
		client:      client,
		visitedURLs: sync.Map{},
		badURLs:     make(map[string]bool),
		rateLimiter: rate.NewLimiter(rate.Inf, 0), // No limit for tests
	}

	ctx := context.Background()
	_, err := fetchWithCache(ctx, client, server.URL, app)

	if err == nil {
		t.Fatal("Expected error for 404 response")
	}

	// URL should be added to bad URLs
	if !app.badURLs[server.URL] {
		t.Error("Expected URL to be added to bad URLs list")
	}
}

func TestWriteFileAtomic(t *testing.T) {
	// Test atomic file writing
	tmpDir := t.TempDir()
	testFile := tmpDir + "/test.txt"
	testData := []byte("test data")

	err := writeFileAtomic(testFile, testData)
	if err != nil {
		t.Fatalf("writeFileAtomic failed: %v", err)
	}

	// Verify file was written correctly
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read written file: %v", err)
	}

	if string(data) != string(testData) {
		t.Errorf("Expected %q, got %q", string(testData), string(data))
	}
}

func TestExtractJSONURLsPool(t *testing.T) {
	// Test that URL slice pool works correctly
	jsonData := []byte(`{
		"references": {
			"test1": {
				"url": "test1.json"
			},
			"test2": {
				"url": "test2.json"
			}
		}
	}`)

	urls1 := extractJSONURLs(jsonData)
	urls2 := extractJSONURLs(jsonData)

	// Should get the same results
	if len(urls1) != len(urls2) {
		t.Errorf("Expected same length, got %d and %d", len(urls1), len(urls2))
	}

	// Results should be independent (not sharing underlying array)
	if len(urls1) > 0 && len(urls2) > 0 {
		// Modify one slice
		urls1[0] = "modified"
		// Other slice should be unchanged
		if urls2[0] == "modified" {
			t.Error("URL slices are sharing memory when they shouldn't")
		}
	}
}