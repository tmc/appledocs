package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestResolveURL(t *testing.T) {
	tests := []struct {
		name     string
		base     string
		relative string
		expected string
	}{
		{
			name:     "absolute URL with http",
			base:     "https://developer.apple.com",
			relative: "https://example.com/test",
			expected: "https://example.com/test",
		},
		{
			name:     "absolute URL with https",
			base:     "https://developer.apple.com",
			relative: "http://example.com/test",
			expected: "http://example.com/test",
		},
		{
			name:     "documentation path",
			base:     "https://developer.apple.com",
			relative: "documentation/SwiftUI",
			expected: "https://developer.apple.com/tutorials/data/documentation/SwiftUI",
		},
		{
			name:     "relative path with leading slash",
			base:     "https://developer.apple.com",
			relative: "/tutorials/data/documentation/UIKit.json",
			expected: "https://developer.apple.com/tutorials/data/documentation/UIKit.json",
		},
		{
			name:     "relative path without leading slash",
			base:     "https://developer.apple.com",
			relative: "tutorials/data/documentation/Foundation.json",
			expected: "https://developer.apple.com/tutorials/data/documentation/Foundation.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := resolveURL(tt.base, tt.relative)
			if result != tt.expected {
				t.Errorf("resolveURL() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestExtractJSONURLs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name: "basic JSON with URL field",
			input: `{
				"title": "Test Framework",
				"url": "documentation/TestFramework.json"
			}`,
			expected: []string{"documentation/TestFramework.json"},
		},
		{
			name: "nested JSON with destination identifiers",
			input: `{
				"references": {
					"ref1": {
						"type": "reference",
						"identifier": "doc://com.apple.documentation/documentation/SwiftUI/View"
					}
				}
			}`,
			expected: []string{"documentation/SwiftUI/View.json"},
		},
		{
			name: "array with path fields",
			input: `{
				"technologies": [
					{
						"title": "SwiftUI",
						"path": "/documentation/swiftui/view/button"
					},
					{
						"title": "UIKit", 
						"path": "/documentation/uikit/view/controller"
					}
				]
			}`,
			expected: []string{"/documentation/swiftui/view/button", "/documentation/uikit/view/controller"},
		},
		{
			name: "mixed URL patterns",
			input: `{
				"mainURL": "documentation/Framework.json",
				"destination": {
					"type": "reference",
					"identifier": "doc://com.apple.documentation/documentation/Foundation/NSObject"
				},
				"items": [
					{
						"path": "/documentation/core/basics/advanced"
					}
				]
			}`,
			expected: []string{
				"/documentation/core/basics/advanced",
				"documentation/Framework.json",
				"documentation/Foundation/NSObject.json",
			},
		},
		{
			name:     "invalid JSON",
			input:    `{invalid json}`,
			expected: nil,
		},
		{
			name:     "empty JSON",
			input:    `{}`,
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractJSONURLs([]byte(tt.input))
			
			if len(result) != len(tt.expected) {
				t.Errorf("extractJSONURLs() returned %d URLs, expected %d. Got: %v", len(result), len(tt.expected), result)
				return
			}
			
			// Sort both to compare regardless of order since extraction order may vary
			sort.Strings(result)
			expected := append([]string{}, tt.expected...)
			sort.Strings(expected)
			
			for i, url := range result {
				if url != expected[i] {
					t.Errorf("extractJSONURLs()[%d] = %v, want %v", i, url, expected[i])
				}
			}
		})
	}
}

func TestIsSymbolURL(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected bool
	}{
		{
			name:     "method with numeric identifier",
			url:      "/documentation/uikit/uiview/1622418-alpha",
			expected: true,
		},
		{
			name:     "deep nested symbol",
			url:      "/documentation/foundation/nsstring/1411946-stringwithformat",
			expected: true,
		},
		{
			name:     "framework documentation",
			url:      "/documentation/swiftui",
			expected: false,
		},
		{
			name:     "class documentation",
			url:      "/documentation/uikit/uiview",
			expected: false,
		},
		{
			name:     "protocol documentation",
			url:      "/documentation/swiftui/view",
			expected: false,
		},
		{
			name:     "very deep symbol path",
			url:      "/documentation/foundation/nsstring/methods/formatting/stringwithformat",
			expected: true,
		},
		{
			name:     "root documentation",
			url:      "/documentation",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSymbolURL(tt.url)
			if result != tt.expected {
				t.Errorf("isSymbolURL(%q) = %v, want %v", tt.url, result, tt.expected)
			}
		})
	}
}

func TestShouldExcludePath(t *testing.T) {
	// Store original value to restore later
	originalExcludePaths := *excludePaths
	defer func() {
		*excludePaths = originalExcludePaths
	}()

	tests := []struct {
		name         string
		excludePaths string
		path         string
		expected     bool
	}{
		{
			name:         "empty exclude paths",
			excludePaths: "",
			path:         "/documentation/swiftui",
			expected:     false,
		},
		{
			name:         "single exclude pattern match",
			excludePaths: "en-US/docs/Mozilla",
			path:         "/en-US/docs/Mozilla/test",
			expected:     true,
		},
		{
			name:         "single exclude pattern no match",
			excludePaths: "en-US/docs/Mozilla",
			path:         "/documentation/swiftui",
			expected:     false,
		},
		{
			name:         "multiple exclude patterns with match",
			excludePaths: "Mozilla,test-path,unwanted",
			path:         "/documentation/test-path/file.json",
			expected:     true,
		},
		{
			name:         "multiple exclude patterns no match",
			excludePaths: "Mozilla,test-path,unwanted",
			path:         "/documentation/swiftui/view.json",
			expected:     false,
		},
		{
			name:         "whitespace in patterns",
			excludePaths: " Mozilla , test-path , unwanted ",
			path:         "/documentation/test-path/file.json",
			expected:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			*excludePaths = tt.excludePaths
			result := shouldExcludePath(tt.path)
			if result != tt.expected {
				t.Errorf("shouldExcludePath(%q) = %v, want %v", tt.path, result, tt.expected)
			}
		})
	}
}

func TestIsJSON(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected bool
	}{
		{
			name:     "valid JSON object",
			data:     []byte(`{"key": "value"}`),
			expected: true,
		},
		{
			name:     "valid JSON array",
			data:     []byte(`[1, 2, 3]`),
			expected: true,
		},
		{
			name:     "valid JSON string",
			data:     []byte(`"simple string"`),
			expected: true,
		},
		{
			name:     "invalid JSON",
			data:     []byte(`{invalid json}`),
			expected: false,
		},
		{
			name:     "empty data",
			data:     []byte(``),
			expected: false,
		},
		{
			name:     "malformed JSON",
			data:     []byte(`{"key": value}`),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isJSON(tt.data)
			if result != tt.expected {
				t.Errorf("isJSON() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPrettyPrintJSON(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		shouldError bool
	}{
		{
			name:        "valid JSON object",
			input:       []byte(`{"name":"test","value":123}`),
			shouldError: false,
		},
		{
			name:        "valid JSON array",
			input:       []byte(`[1,2,3]`),
			shouldError: false,
		},
		{
			name:        "invalid JSON",
			input:       []byte(`{invalid}`),
			shouldError: true,
		},
		{
			name:        "empty JSON object",
			input:       []byte(`{}`),
			shouldError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := prettyPrintJSON(tt.input)
			
			if tt.shouldError {
				if err == nil {
					t.Errorf("prettyPrintJSON() should have returned an error")
				}
			} else {
				if err != nil {
					t.Errorf("prettyPrintJSON() unexpected error: %v", err)
				}
				
				if len(result) == 0 {
					t.Errorf("prettyPrintJSON() returned empty result")
				}
				
				// Verify the result is valid JSON
				if !isJSON(result) {
					t.Errorf("prettyPrintJSON() result is not valid JSON")
				}
			}
		})
	}
}

func TestBuildFileTree(t *testing.T) {
	files := []JSONFileEntry{
		{Path: "tutorials/data/documentation/SwiftUI.json"},
		{Path: "tutorials/data/documentation/UIKit.json"},
		{Path: "tutorials/data/documentation/SwiftUI/View.json"},
		{Path: "tutorials/data/documentation/UIKit/UIView.json"},
		{Path: "tutorials/data/index/swiftui"},
	}

	root := buildFileTree(files)

	if root.Name != "root" {
		t.Errorf("Root node name should be 'root', got %q", root.Name)
	}

	if !root.IsDir {
		t.Errorf("Root node should be a directory")
	}

	if len(root.Children) == 0 {
		t.Errorf("Root should have children")
	}

	// Verify tutorials directory exists
	var tutorialsNode *TreeNode
	for _, child := range root.Children {
		if child.Name == "tutorials" {
			tutorialsNode = child
			break
		}
	}

	if tutorialsNode == nil {
		t.Errorf("Should have tutorials directory")
	}

	if !tutorialsNode.IsDir {
		t.Errorf("Tutorials should be a directory")
	}
}

func TestCountDirectories(t *testing.T) {
	root := &TreeNode{
		Name:  "root",
		IsDir: true,
		Children: []*TreeNode{
			{
				Name:  "dir1",
				IsDir: true,
				Children: []*TreeNode{
					{Name: "file1.json", IsDir: false},
					{Name: "subdir", IsDir: true, Children: []*TreeNode{}},
				},
			},
			{Name: "file2.json", IsDir: false},
		},
	}

	count := countDirectories(root)
	
	// Should count dir1 and subdir (not root)
	expected := 2
	if count != expected {
		t.Errorf("countDirectories() = %d, want %d", count, expected)
	}
}

func TestAddBrowserLikeHeaders(t *testing.T) {
	req := httptest.NewRequest("GET", "https://developer.apple.com/documentation/SwiftUI.json", nil)
	
	addBrowserLikeHeaders(req)

	// Check essential headers are set
	requiredHeaders := []string{
		"User-Agent",
		"Accept",
		"Accept-Language",
		"Cache-Control",
		"Connection",
	}

	for _, header := range requiredHeaders {
		if req.Header.Get(header) == "" {
			t.Errorf("Header %q should be set", header)
		}
	}

	// Check User-Agent contains browser-like information
	userAgent := req.Header.Get("User-Agent")
	if !strings.Contains(userAgent, "Mozilla") {
		t.Errorf("User-Agent should contain 'Mozilla', got %q", userAgent)
	}

	// Check Referer is set appropriately for JSON requests
	referer := req.Header.Get("Referer")
	if !strings.Contains(referer, "developer.apple.com") {
		t.Errorf("Referer should contain developer.apple.com, got %q", referer)
	}
}

func TestSaveToOutputDir(t *testing.T) {
	// Store original values to restore later
	originalOutputDir := *outputDir
	originalPrettyJSON := *prettyJSON
	defer func() {
		*outputDir = originalOutputDir
		*prettyJSON = originalPrettyJSON
	}()

	tempDir := t.TempDir()
	*outputDir = tempDir

	tests := []struct {
		name         string
		path         string
		content      []byte
		prettyJSON   bool
		expectPretty bool
	}{
		{
			name:         "regular text file",
			path:         "test.txt",
			content:      []byte("plain text content"),
			prettyJSON:   false,
			expectPretty: false,
		},
		{
			name:         "JSON file with pretty printing disabled",
			path:         "test.json",
			content:      []byte(`{"key":"value"}`),
			prettyJSON:   false,
			expectPretty: false,
		},
		{
			name:         "JSON file with pretty printing enabled",
			path:         "nested/test.json",
			content:      []byte(`{"name":"test","items":[1,2,3]}`),
			prettyJSON:   true,
			expectPretty: true,
		},
		{
			name:         "invalid JSON with pretty printing enabled",
			path:         "invalid.json",
			content:      []byte(`{invalid json}`),
			prettyJSON:   true,
			expectPretty: false, // Should fall back to original content
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			*prettyJSON = tt.prettyJSON

			err := saveToOutputDir(tt.path, tt.content)
			if err != nil {
				t.Errorf("saveToOutputDir() error = %v", err)
			}

			// Verify file was created
			outputPath := filepath.Join(tempDir, tt.path)
			savedContent, err := os.ReadFile(outputPath)
			if err != nil {
				t.Errorf("Failed to read saved file: %v", err)
			}

			// Check if content was pretty-printed when expected
			if tt.expectPretty && strings.HasSuffix(tt.path, ".json") {
				// Pretty-printed JSON should have newlines and indentation
				if !strings.Contains(string(savedContent), "\n") || !strings.Contains(string(savedContent), "  ") {
					t.Errorf("Expected pretty-printed JSON but got compact format")
				}
			} else {
				// Content should match original
				if string(savedContent) != string(tt.content) && !tt.expectPretty {
					t.Errorf("Saved content doesn't match original")
				}
			}
		})
	}
}

func TestIsDataURL(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected bool
	}{
		{
			name:     "valid data URL",
			url:      "https://developer.apple.com/tutorials/data/documentation/SwiftUI.json",
			expected: true,
		},
		{
			name:     "another valid data URL",
			url:      "/tutorials/data/index/swift",
			expected: true,
		},
		{
			name:     "non-data URL",
			url:      "https://developer.apple.com/documentation/SwiftUI",
			expected: false,
		},
		{
			name:     "marketing URL",
			url:      "https://developer.apple.com/swift/",
			expected: false,
		},
		{
			name:     "empty URL",
			url:      "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isDataURL(tt.url)
			if result != tt.expected {
				t.Errorf("isDataURL(%q) = %v, want %v", tt.url, result, tt.expected)
			}
		})
	}
}

func TestBuildFrameworkURLs(t *testing.T) {
	tests := []struct {
		name      string
		framework string
		expected  []string
	}{
		{
			name:      "simple framework name",
			framework: "SwiftUI",
			expected: []string{
				"tutorials/data/documentation/SwiftUI.json",
				"tutorials/data/index/swiftui",
			},
		},
		{
			name:      "framework with .json suffix",
			framework: "UIKit.json",
			expected: []string{
				"tutorials/data/documentation/UIKit.json",
				"tutorials/data/index/uikit",
			},
		},
		{
			name:      "mixed case framework",
			framework: "EndpointSecurity",
			expected: []string{
				"tutorials/data/documentation/EndpointSecurity.json",
				"tutorials/data/index/endpointsecurity",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildFrameworkURLs(tt.framework)
			
			if len(result) != len(tt.expected) {
				t.Errorf("buildFrameworkURLs() returned %d URLs, expected %d", len(result), len(tt.expected))
				return
			}
			
			for i, url := range result {
				if url != tt.expected[i] {
					t.Errorf("buildFrameworkURLs()[%d] = %v, want %v", i, url, tt.expected[i])
				}
			}
		})
	}
}

func TestAppledocsStatsMethods(t *testing.T) {
	app := &crawler{
		badURLs:     make(map[string]bool),
		urlDepths:   make(map[string]int),
		rateLimiter: rate.NewLimiter(rate.Inf, 0), // No limit for tests
	}
	// visitedURLs is a sync.Map, no initialization needed

	// Test initial state
	processed, cacheHits, cacheMisses, errors, skipped := app.getStats()
	if processed != 0 || cacheHits != 0 || cacheMisses != 0 || errors != 0 || skipped != 0 {
		t.Errorf("Initial stats should all be zero")
	}

	// Test increment methods
	app.incrementCacheHits()
	app.incrementCacheHits()
	app.incrementCacheMisses()
	app.incrementErrors()
	app.incrementSkippedURLs()
	app.incrementSkippedSymbols()

	processed, cacheHits, cacheMisses, errors, skipped = app.getStats()
	
	if cacheHits != 2 {
		t.Errorf("Expected 2 cache hits, got %d", cacheHits)
	}
	if cacheMisses != 1 {
		t.Errorf("Expected 1 cache miss, got %d", cacheMisses)
	}
	if errors != 1 {
		t.Errorf("Expected 1 error, got %d", errors)
	}
	if skipped != 2 { // incrementSkippedSymbols also increments skippedURLs
		t.Errorf("Expected 2 skipped URLs, got %d", skipped)
	}
}

func TestQueueNewURLs(t *testing.T) {
	// Store original values
	originalBaseURL := *baseURL
	defer func() {
		*baseURL = originalBaseURL
	}()

	*baseURL = "https://example.com"

	app := &crawler{
		badURLs:     make(map[string]bool),
		urlDepths:   make(map[string]int),
		rateLimiter: rate.NewLimiter(rate.Inf, 0), // No limit for tests
	}
	// visitedURLs is a sync.Map, no initialization needed

	urlQueue := make(chan string, 10)
	defer close(urlQueue)

	newURLs := []string{
		"documentation/SwiftUI.json",
		"documentation/UIKit.json",
		"documentation/SwiftUI.json", // Duplicate
	}

	added := app.queueNewURLs(newURLs, urlQueue)

	// Should add 2 unique URLs
	if added != 2 {
		t.Errorf("Expected 2 URLs to be added, got %d", added)
	}

	// Verify URLs were queued
	queuedCount := 0
	for {
		select {
		case <-urlQueue:
			queuedCount++
		default:
			goto done
		}
	}
done:

	if queuedCount != 2 {
		t.Errorf("Expected 2 URLs in queue, got %d", queuedCount)
	}

	// Test that subsequent call with same URLs adds nothing
	added2 := app.queueNewURLs(newURLs, urlQueue)
	if added2 != 0 {
		t.Errorf("Expected 0 URLs to be added on second call, got %d", added2)
	}
}

func TestFetchWithCacheIntegration(t *testing.T) {
	// Store original values
	originalCacheDir := *cacheDir
	originalForceRefresh := *forceRefresh
	defer func() {
		*cacheDir = originalCacheDir
		*forceRefresh = originalForceRefresh
	}()

	tempDir := t.TempDir()
	*cacheDir = tempDir
	*forceRefresh = false

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/test.json" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"test": "data"}`)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	app := &crawler{
		badURLs:     make(map[string]bool),
		urlDepths:   make(map[string]int),
		rateLimiter: rate.NewLimiter(rate.Inf, 0), // No limit for tests
	}
	// visitedURLs is a sync.Map, no initialization needed

	client := &http.Client{Timeout: 5 * time.Second}

	// Test successful fetch
	testURL := server.URL + "/test.json"
	ctx := context.Background()
	data, err := fetchWithCache(ctx, client, testURL, app)
	if err != nil {
		t.Errorf("fetchWithCache() error = %v", err)
	}

	if string(data) != `{"test": "data"}` {
		t.Errorf("fetchWithCache() returned unexpected data: %s", string(data))
	}

	// Verify cache hit on second call
	app.cacheHits = 0 // Reset counter
	data2, err := fetchWithCache(ctx, client, testURL, app)
	if err != nil {
		t.Errorf("fetchWithCache() second call error = %v", err)
	}

	if string(data2) != string(data) {
		t.Errorf("Cached data doesn't match original")
	}

	if app.cacheHits != 1 {
		t.Errorf("Expected 1 cache hit, got %d", app.cacheHits)
	}

	// Test 404 handling
	notFoundURL := server.URL + "/nonexistent.json"
	_, err = fetchWithCache(ctx, client, notFoundURL, app)
	if err == nil {
		t.Errorf("fetchWithCache() should return error for 404")
	}

	if !app.badURLs[notFoundURL] {
		t.Errorf("404 URL should be added to bad URLs")
	}
}

func TestProcessURLContextCancellation(t *testing.T) {
	app := &crawler{
		client:      &http.Client{Timeout: 1 * time.Second},
		badURLs:     make(map[string]bool),
		urlDepths:   make(map[string]int),
		rateLimiter: rate.NewLimiter(rate.Inf, 0), // No limit for tests
	}
	// visitedURLs is a sync.Map, no initialization needed

	// Create a context that's already cancelled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	urlQueue := make(chan string, 1)
	defer close(urlQueue)

	err := app.processURL(ctx, "https://example.com/test.json", urlQueue)
	
	// Should handle cancellation gracefully
	if err != nil && err != context.Canceled && !strings.Contains(err.Error(), "context deadline exceeded") && !strings.Contains(err.Error(), "context canceled") {
		t.Errorf("processURL() should handle context cancellation gracefully, got error: %v", err)
	}
}

// Benchmark tests for performance
func BenchmarkExtractJSONURLs(b *testing.B) {
	jsonData := []byte(`{
		"technologies": [
			{
				"title": "SwiftUI",
				"path": "/documentation/swiftui",
				"url": "documentation/swiftui.json"
			},
			{
				"title": "UIKit",
				"path": "/documentation/uikit",
				"url": "documentation/uikit.json"
			}
		],
		"references": {
			"ref1": {
				"type": "reference",
				"identifier": "doc://com.apple.documentation/documentation/Foundation/NSObject"
			}
		}
	}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = extractJSONURLs(jsonData)
	}
}

func BenchmarkResolveURL(b *testing.B) {
	base := "https://developer.apple.com"
	relative := "documentation/SwiftUI/View.json"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = resolveURL(base, relative)
	}
}

func BenchmarkBuildFileTree(b *testing.B) {
	// Create a large set of files for benchmarking
	files := make([]JSONFileEntry, 1000)
	for i := 0; i < 1000; i++ {
		files[i] = JSONFileEntry{
			Path: fmt.Sprintf("tutorials/data/documentation/framework%d/class%d.json", i%10, i),
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = buildFileTree(files)
	}
}