package main

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

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
			name:     "invalid JSON",
			data:     []byte(`{key: value`),
			expected: false,
		},
		{
			name:     "empty data",
			data:     []byte(``),
			expected: false,
		},
		{
			name:     "plain text",
			data:     []byte(`hello world`),
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
		expectError bool
	}{
		{
			name:        "valid JSON",
			input:       []byte(`{"name":"test","value":123}`),
			expectError: false,
		},
		{
			name:        "invalid JSON",
			input:       []byte(`{name:test`),
			expectError: true,
		},
		{
			name:        "nested JSON",
			input:       []byte(`{"outer":{"inner":"value"},"array":[1,2,3]}`),
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := prettyPrintJSON(tt.input)
			if tt.expectError {
				if err == nil {
					t.Errorf("prettyPrintJSON() expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("prettyPrintJSON() error = %v", err)
				}
				if len(result) == 0 {
					t.Errorf("prettyPrintJSON() returned empty result")
				}
				// Verify it's still valid JSON
				if !isJSON(result) {
					t.Errorf("prettyPrintJSON() result is not valid JSON")
				}
			}
		})
	}
}

func TestResolveURL(t *testing.T) {
	baseURL := "https://developer.apple.com"
	tests := []struct {
		name     string
		base     string
		relative string
		expected string
	}{
		{
			name:     "absolute HTTP URL",
			base:     baseURL,
			relative: "https://example.com/path",
			expected: "https://example.com/path",
		},
		{
			name:     "absolute HTTPS URL",
			base:     baseURL,
			relative: "http://example.com/path",
			expected: "http://example.com/path",
		},
		{
			name:     "documentation path",
			base:     baseURL,
			relative: "documentation/SwiftUI",
			expected: "https://developer.apple.com/tutorials/data/documentation/SwiftUI",
		},
		{
			name:     "absolute path",
			base:     baseURL,
			relative: "/tutorials/data/test",
			expected: "https://developer.apple.com/tutorials/data/test",
		},
		{
			name:     "relative path",
			base:     baseURL,
			relative: "test/path",
			expected: "https://developer.apple.com/test/path",
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
		data     []byte
		expected []string
	}{
		{
			name:     "simple URL field",
			data:     []byte(`{"url": "test.json"}`),
			expected: []string{"test.json"},
		},
		{
			name:     "doc:// identifier",
			data:     []byte(`{"type": "reference", "identifier": "doc://com.apple.documentation/documentation/SwiftUI"}`),
			expected: []string{"documentation/SwiftUI.json"},
		},
		{
			name:     "path field",
			data:     []byte(`{"path": "/documentation/swiftui/view"}`),
			expected: []string{"/documentation/swiftui/view"},
		},
		{
			name: "multiple URLs",
			data: []byte(`{
				"url": "first.json",
				"nested": {
					"secondURL": "second.json"
				},
				"array": [
					{"type": "ref", "identifier": "doc://com.apple.documentation/documentation/UIKit"}
				]
			}`),
			expected: []string{"first.json", "second.json", "documentation/UIKit.json"},
		},
		{
			name:     "no URLs",
			data:     []byte(`{"title": "Test", "value": 123}`),
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractJSONURLs(tt.data)
			if len(result) != len(tt.expected) {
				t.Errorf("extractJSONURLs() returned %d URLs, want %d", len(result), len(tt.expected))
			}
			for i, url := range tt.expected {
				found := false
				for _, r := range result {
					if r == url {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("extractJSONURLs() missing expected URL at index %d: %v", i, url)
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
			name:     "symbol with number",
			url:      "/documentation/uikit/uiview/1622418-alpha",
			expected: true,
		},
		{
			name:     "method path",
			url:      "/documentation/foundation/nsstring/method/name",
			expected: true,
		},
		{
			name:     "deep symbol path",
			url:      "/documentation/swift/array/element/index/value",
			expected: true,
		},
		{
			name:     "framework root",
			url:      "/documentation/swiftui",
			expected: false,
		},
		{
			name:     "class level",
			url:      "/documentation/uikit/uiview",
			expected: false,
		},
		{
			name:     "shallow path",
			url:      "/documentation/foundation/nsstring",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSymbolURL(tt.url)
			if result != tt.expected {
				t.Errorf("isSymbolURL() = %v, want %v", result, tt.expected)
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
			name:     "tutorials data URL",
			url:      "https://developer.apple.com/tutorials/data/documentation/SwiftUI.json",
			expected: true,
		},
		{
			name:     "regular documentation URL",
			url:      "https://developer.apple.com/documentation/SwiftUI",
			expected: false,
		},
		{
			name:     "other path",
			url:      "https://developer.apple.com/design/human-interface-guidelines",
			expected: false,
		},
		{
			name:     "relative data URL",
			url:      "/tutorials/data/index/swiftui",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isDataURL(tt.url)
			if result != tt.expected {
				t.Errorf("isDataURL() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBuildFrameworkURLs(t *testing.T) {
	tests := []struct {
		name         string
		framework    string
		expectedURLs []string
	}{
		{
			name:      "SwiftUI framework",
			framework: "SwiftUI",
			expectedURLs: []string{
				"tutorials/data/documentation/SwiftUI.json",
				"tutorials/data/index/swiftui",
			},
		},
		{
			name:      "framework with .json suffix",
			framework: "Foundation.json",
			expectedURLs: []string{
				"tutorials/data/documentation/Foundation.json",
				"tutorials/data/index/foundation",
			},
		},
		{
			name:      "UIKit framework",
			framework: "UIKit",
			expectedURLs: []string{
				"tutorials/data/documentation/UIKit.json",
				"tutorials/data/index/uikit",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildFrameworkURLs(tt.framework)
			if len(result) != len(tt.expectedURLs) {
				t.Errorf("buildFrameworkURLs() returned %d URLs, want %d", len(result), len(tt.expectedURLs))
			}
			for i, expected := range tt.expectedURLs {
				if i >= len(result) || result[i] != expected {
					t.Errorf("buildFrameworkURLs() URL at index %d = %v, want %v", i, result[i], expected)
				}
			}
		})
	}
}

func TestShouldExcludePath(t *testing.T) {
	// Save original value
	originalExcludePaths := *excludePaths
	defer func() {
		*excludePaths = originalExcludePaths
	}()

	tests := []struct {
		name         string
		excludePaths string
		pathToCheck  string
		expected     bool
	}{
		{
			name:         "no exclude paths",
			excludePaths: "",
			pathToCheck:  "/any/path",
			expected:     false,
		},
		{
			name:         "path matches single exclude pattern",
			excludePaths: "Mozilla",
			pathToCheck:  "/en-US/docs/Mozilla/guide",
			expected:     true,
		},
		{
			name:         "path matches one of multiple exclude patterns",
			excludePaths: "Mozilla,private,internal",
			pathToCheck:  "/docs/private/api",
			expected:     true,
		},
		{
			name:         "path doesn't match any exclude pattern",
			excludePaths: "Mozilla,private",
			pathToCheck:  "/documentation/SwiftUI",
			expected:     false,
		},
		{
			name:         "comma-separated with spaces",
			excludePaths: "Mozilla, private, internal",
			pathToCheck:  "/docs/internal/api",
			expected:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			*excludePaths = tt.excludePaths
			result := shouldExcludePath(tt.pathToCheck)
			if result != tt.expected {
				t.Errorf("shouldExcludePath() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAppledocsStatsTracking(t *testing.T) {
	app := &appledocs{
		visitedURLs: make(map[string]bool),
		badURLs:     make(map[string]bool),
		urlDepths:   make(map[string]int),
	}

	// Test initial stats
	processed, cacheHits, cacheMisses, errors, skipped := app.getStats()
	if processed != 0 || cacheHits != 0 || cacheMisses != 0 || errors != 0 || skipped != 0 {
		t.Errorf("Initial stats should be zero, got: processed=%d, cacheHits=%d, cacheMisses=%d, errors=%d, skipped=%d",
			processed, cacheHits, cacheMisses, errors, skipped)
	}

	// Test incrementing each stat
	app.incrementCacheHits()
	app.incrementCacheMisses()
	app.incrementErrors()
	app.incrementSkippedURLs()
	app.incrementSkippedSymbols()

	processed, cacheHits, cacheMisses, errors, skipped = app.getStats()
	if cacheHits != 1 || cacheMisses != 1 || errors != 1 || skipped != 2 {
		t.Errorf("Stats after incrementing: cacheHits=%d (want 1), cacheMisses=%d (want 1), errors=%d (want 1), skipped=%d (want 2)",
			cacheHits, cacheMisses, errors, skipped)
	}
}

func TestFetchWithCacheMockServer(t *testing.T) {
	// Create a temporary directory for cache
	tempDir := t.TempDir()
	originalCacheDir := *cacheDir
	*cacheDir = tempDir
	defer func() {
		*cacheDir = originalCacheDir
	}()

	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that browser-like headers are set
		if r.Header.Get("User-Agent") == "" {
			t.Errorf("User-Agent header not set")
		}
		if r.Header.Get("Accept") == "" {
			t.Errorf("Accept header not set")
		}

		switch r.URL.Path {
		case "/test.json":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"test": "data"}`))
		case "/notfound.json":
			w.WriteHeader(http.StatusNotFound)
		case "/error.json":
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	app := &appledocs{
		visitedURLs: make(map[string]bool),
		badURLs:     make(map[string]bool),
		urlDepths:   make(map[string]int),
	}

	client := &http.Client{Timeout: 5 * time.Second}

	// Test successful fetch
	testURL := server.URL + "/test.json"
	data, err := fetchWithCache(client, testURL, app)
	if err != nil {
		t.Errorf("fetchWithCache() error = %v", err)
	}
	if !isJSON(data) {
		t.Errorf("fetchWithCache() returned invalid JSON")
	}

	// Test cache hit (second fetch should use cache)
	originalCacheHits := app.cacheHits
	data2, err := fetchWithCache(client, testURL, app)
	if err != nil {
		t.Errorf("fetchWithCache() cache hit error = %v", err)
	}
	if !bytes.Equal(data, data2) {
		t.Errorf("fetchWithCache() cache hit returned different data")
	}
	if app.cacheHits <= originalCacheHits {
		t.Errorf("fetchWithCache() should have incremented cache hits")
	}

	// Test 404 error
	notFoundURL := server.URL + "/notfound.json"
	_, err = fetchWithCache(client, notFoundURL, app)
	if err == nil {
		t.Errorf("fetchWithCache() should return error for 404")
	}
	if !app.badURLs[notFoundURL] {
		t.Errorf("fetchWithCache() should add 404 URL to bad URLs")
	}

	// Test server error
	errorURL := server.URL + "/error.json"
	_, err = fetchWithCache(client, errorURL, app)
	if err == nil {
		t.Errorf("fetchWithCache() should return error for server error")
	}
}

func TestSaveToOutputDir(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()
	originalOutputDir := *outputDir
	*outputDir = tempDir
	defer func() {
		*outputDir = originalOutputDir
	}()

	tests := []struct {
		name        string
		path        string
		content     []byte
		prettyJSON  bool
		expectError bool
	}{
		{
			name:        "save regular file",
			path:        "test.txt",
			content:     []byte("hello world"),
			prettyJSON:  false,
			expectError: false,
		},
		{
			name:        "save JSON with pretty printing",
			path:        "test.json",
			content:     []byte(`{"key":"value"}`),
			prettyJSON:  true,
			expectError: false,
		},
		{
			name:        "save nested path",
			path:        "nested/path/file.json",
			content:     []byte(`{"nested":"data"}`),
			prettyJSON:  true,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalPrettyJSON := *prettyJSON
			*prettyJSON = tt.prettyJSON
			defer func() {
				*prettyJSON = originalPrettyJSON
			}()

			err := saveToOutputDir(tt.path, tt.content)
			if tt.expectError {
				if err == nil {
					t.Errorf("saveToOutputDir() expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("saveToOutputDir() error = %v", err)
				}

				// Verify file was created
				fullPath := filepath.Join(tempDir, tt.path)
				savedContent, err := os.ReadFile(fullPath)
				if err != nil {
					t.Errorf("saveToOutputDir() failed to create file: %v", err)
				}

				// For JSON files with pretty printing, verify it's still valid JSON
				if tt.prettyJSON && strings.HasSuffix(tt.path, ".json") {
					if !isJSON(savedContent) {
						t.Errorf("saveToOutputDir() saved invalid JSON")
					}
				}
			}
		})
	}
}

func TestIsTechnologyFile(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "top-level technology file",
			path:     "/tutorials/data/documentation/SwiftUI.json",
			expected: true,
		},
		{
			name:     "nested documentation file",
			path:     "/tutorials/data/documentation/SwiftUI/View.json",
			expected: false,
		},
		{
			name:     "non-documentation file",
			path:     "/tutorials/other/SwiftUI.json",
			expected: false,
		},
		{
			name:     "wrong extension",
			path:     "/tutorials/data/documentation/SwiftUI.html",
			expected: false,
		},
		{
			name:     "too short path",
			path:     "/tutorials/data/documentation",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isTechnologyFile(tt.path)
			if result != tt.expected {
				t.Errorf("isTechnologyFile() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestContextCancellation(t *testing.T) {
	// Create a context that will be cancelled quickly
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	app := &appledocs{
		client:      &http.Client{Timeout: 50 * time.Millisecond}, // Add client
		visitedURLs: make(map[string]bool),
		badURLs:     make(map[string]bool),
		urlDepths:   make(map[string]int),
	}

	// Create a channel that would normally block
	urlQueue := make(chan string, 1)

	// This should return quickly due to context cancellation or timeout
	err := app.processURL(ctx, "https://invalid-url-that-should-timeout", urlQueue)

	// We expect either a context cancellation error or a network error
	if err == nil {
		t.Errorf("processURL() should return error for cancelled context or invalid URL")
	}
}

func TestAddFileToTree(t *testing.T) {
	root := &TreeNode{
		Name:     "root",
		Path:     "",
		IsDir:    true,
		Children: []*TreeNode{},
	}

	// Test adding files to build a tree
	addFileToTree(root, "folder1/file1.json", "folder1/file1.json")
	addFileToTree(root, "folder1/file2.json", "folder1/file2.json")
	addFileToTree(root, "folder2/subfolder/file3.json", "folder2/subfolder/file3.json")

	// Verify tree structure
	if len(root.Children) != 2 {
		t.Errorf("Root should have 2 children, got %d", len(root.Children))
	}

	// Find folder1
	var folder1 *TreeNode
	for _, child := range root.Children {
		if child.Name == "folder1" && child.IsDir {
			folder1 = child
			break
		}
	}
	if folder1 == nil {
		t.Errorf("folder1 not found in tree")
	} else if len(folder1.Children) != 2 {
		t.Errorf("folder1 should have 2 children, got %d", len(folder1.Children))
	}

	// Find folder2
	var folder2 *TreeNode
	for _, child := range root.Children {
		if child.Name == "folder2" && child.IsDir {
			folder2 = child
			break
		}
	}
	if folder2 == nil {
		t.Errorf("folder2 not found in tree")
	} else if len(folder2.Children) != 1 {
		t.Errorf("folder2 should have 1 child, got %d", len(folder2.Children))
	}
}

func TestSortTree(t *testing.T) {
	root := &TreeNode{
		Name:  "root",
		IsDir: true,
		Children: []*TreeNode{
			{Name: "file2.json", IsDir: false},
			{Name: "dirB", IsDir: true, Children: []*TreeNode{}},
			{Name: "file1.json", IsDir: false},
			{Name: "dirA", IsDir: true, Children: []*TreeNode{}},
		},
	}

	sortTree(root)

	// Verify that directories come before files and names are sorted
	if len(root.Children) != 4 {
		t.Errorf("Expected 4 children, got %d", len(root.Children))
	}

	// First two should be directories in alphabetical order
	if !root.Children[0].IsDir || root.Children[0].Name != "dirA" {
		t.Errorf("First child should be dirA, got %s (isDir: %v)", root.Children[0].Name, root.Children[0].IsDir)
	}
	if !root.Children[1].IsDir || root.Children[1].Name != "dirB" {
		t.Errorf("Second child should be dirB, got %s (isDir: %v)", root.Children[1].Name, root.Children[1].IsDir)
	}

	// Last two should be files in alphabetical order
	if root.Children[2].IsDir || root.Children[2].Name != "file1.json" {
		t.Errorf("Third child should be file1.json, got %s (isDir: %v)", root.Children[2].Name, root.Children[2].IsDir)
	}
	if root.Children[3].IsDir || root.Children[3].Name != "file2.json" {
		t.Errorf("Fourth child should be file2.json, got %s (isDir: %v)", root.Children[3].Name, root.Children[3].IsDir)
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
					{Name: "subdir1", IsDir: true, Children: []*TreeNode{}},
					{Name: "file1.json", IsDir: false},
				},
			},
			{Name: "dir2", IsDir: true, Children: []*TreeNode{}},
			{Name: "file2.json", IsDir: false},
		},
	}

	count := countDirectories(root)
	// Should count dir1, subdir1, and dir2 (but not root)
	if count != 3 {
		t.Errorf("countDirectories() = %d, want 3", count)
	}
}

// TestBadURLsFileOperations tests loading and writing bad URLs files
func TestBadURLsFileOperations(t *testing.T) {
	tempDir := t.TempDir()
	badURLsPath := filepath.Join(tempDir, "bad-urls.txt")

	// Save original value
	originalBadURLsFile := *badURLsFile
	*badURLsFile = badURLsPath
	defer func() {
		*badURLsFile = originalBadURLsFile
	}()

	app := &appledocs{
		visitedURLs: make(map[string]bool),
		badURLs:     make(map[string]bool),
		urlDepths:   make(map[string]int),
	}

	// Test loading non-existent file (should not error)
	err := loadBadURLs(app)
	if err != nil {
		t.Errorf("loadBadURLs() should not error for non-existent file: %v", err)
	}

	// Add some bad URLs
	app.badURLs["https://example.com/bad1"] = true
	app.badURLs["https://example.com/bad2"] = true

	// Test writing bad URLs file
	err = writeBadURLsFile(app)
	if err != nil {
		t.Errorf("writeBadURLsFile() error = %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(badURLsPath); os.IsNotExist(err) {
		t.Errorf("writeBadURLsFile() did not create file")
	}

	// Test loading the written file
	newApp := &appledocs{
		visitedURLs: make(map[string]bool),
		badURLs:     make(map[string]bool),
		urlDepths:   make(map[string]int),
	}
	err = loadBadURLs(newApp)
	if err != nil {
		t.Errorf("loadBadURLs() error = %v", err)
	}

	// Verify URLs were loaded
	if !newApp.badURLs["https://example.com/bad1"] {
		t.Errorf("loadBadURLs() did not load bad1 URL")
	}
	if !newApp.badURLs["https://example.com/bad2"] {
		t.Errorf("loadBadURLs() did not load bad2 URL")
	}
}

// Benchmark for performance testing
func BenchmarkExtractJSONURLs(b *testing.B) {
	testData := []byte(`{
		"url": "test.json",
		"nested": {
			"secondURL": "second.json",
			"type": "reference",
			"identifier": "doc://com.apple.documentation/documentation/SwiftUI"
		},
		"array": [
			{"type": "ref", "identifier": "doc://com.apple.documentation/documentation/UIKit"},
			{"path": "/documentation/foundation/nsstring"}
		]
	}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = extractJSONURLs(testData)
	}
}

func BenchmarkIsJSON(b *testing.B) {
	testData := []byte(`{"complex": {"nested": {"structure": [1, 2, 3, {"deep": "value"}]}}}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isJSON(testData)
	}
}
