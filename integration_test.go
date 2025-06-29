package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestEndToEndWorkflow tests the complete workflow from fetching to markdown generation
func TestEndToEndWorkflow(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Store original flag values
	originals := struct {
		outputDir   string
		cacheDir    string
		mdOutputDir string
		baseURL     string
		concurrency int
		prettyJSON  bool
		verbose     bool
	}{
		*outputDir,
		*cacheDir,
		*mdOutputDir,
		*baseURL,
		*concurrency,
		*prettyJSON,
		*verbose,
	}

	defer func() {
		*outputDir = originals.outputDir
		*cacheDir = originals.cacheDir
		*mdOutputDir = originals.mdOutputDir
		*baseURL = originals.baseURL
		*concurrency = originals.concurrency
		*prettyJSON = originals.prettyJSON
		*verbose = originals.verbose
	}()

	// Set up temporary directories
	tempDir := t.TempDir()
	*outputDir = filepath.Join(tempDir, "output")
	*cacheDir = filepath.Join(tempDir, "cache")
	*mdOutputDir = filepath.Join(tempDir, "markdown")
	*concurrency = 2
	*prettyJSON = true
	*verbose = false

	// Create test data
	testFramework := DocJSONData{
		Metadata: Metadata{
			Title:       "TestFramework",
			Role:        "framework",
			RoleHeading: "Framework",
			Modules:     []Module{{Name: "TestFramework"}},
			Platforms: []Platform{
				{Name: "iOS", IntroducedAt: "13.0"},
				{Name: "macOS", IntroducedAt: "10.15"},
			},
		},
		Abstract: []TextContent{
			{Type: "text", Text: "A comprehensive test framework for integration testing."},
		},
		PrimaryContentSections: []ContentSection{
			{
				Kind: "declarations",
				Declarations: []Declaration{
					{
						Languages: []string{"swift"},
						Tokens: []Fragment{
							{Text: "import TestFramework"},
						},
					},
				},
			},
			{
				Kind: "content",
				Content: []ContentBlock{
					{
						Type:  "heading",
						Level: 2,
						Text:  "Overview",
					},
					{
						Type: "paragraph",
						InlineContent: []InlineContent{
							{Type: "text", Text: "This framework provides testing capabilities."},
						},
					},
				},
			},
		},
		TopicSections: []TopicSection{
			{
				Title:       "Classes",
				Identifiers: []string{"doc://com.apple.documentation/documentation/TestFramework/TestClass"},
			},
		},
		References: map[string]Reference{
			"doc://com.apple.documentation/documentation/TestFramework/TestClass": {
				Title: "TestClass",
				URL:   "doc://com.apple.documentation/documentation/TestFramework/TestClass",
				Role:  "class",
				Abstract: []TextContent{
					{Type: "text", Text: "A test class for demonstration."},
				},
			},
		},
	}

	testClass := DocJSONData{
		Metadata: Metadata{
			Title:       "TestClass",
			Role:        "symbol",
			RoleHeading: "Class",
		},
		Abstract: []TextContent{
			{Type: "text", Text: "A test class for demonstration purposes."},
		},
		PrimaryContentSections: []ContentSection{
			{
				Kind: "declarations",
				Declarations: []Declaration{
					{
						Languages: []string{"swift"},
						Tokens: []Fragment{
							{Text: "class TestClass"},
						},
					},
				},
			},
		},
	}

	// Set up test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		switch r.URL.Path {
		case "/tutorials/data/documentation/technologies.json":
			// Return a simple technologies index
			response := map[string]interface{}{
				"technologies": []interface{}{
					map[string]interface{}{
						"title": "TestFramework",
						"url":   "documentation/TestFramework.json",
					},
				},
			}
			json.NewEncoder(w).Encode(response)
			
		case "/tutorials/data/documentation/TestFramework.json":
			json.NewEncoder(w).Encode(testFramework)
			
		case "/tutorials/data/documentation/TestFramework/TestClass.json":
			json.NewEncoder(w).Encode(testClass)
			
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, `{"error": "Not found"}`)
		}
	}))
	defer server.Close()

	*baseURL = server.URL

	// Test Phase 1: Crawling
	t.Run("crawl_phase", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := run(ctx)
		if err != nil {
			t.Fatalf("Crawling phase failed: %v", err)
		}

		// Verify files were created
		expectedFiles := []string{
			"tutorials/data/documentation/technologies.json",
			"tutorials/data/documentation/TestFramework.json",
			"tutorials/data/documentation/TestFramework/TestClass.json",
		}

		for _, file := range expectedFiles {
			fullPath := filepath.Join(*outputDir, file)
			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				t.Errorf("Expected file %s was not created", file)
			}
		}

		// Verify cache was created
		cacheFiles, err := filepath.Glob(filepath.Join(*cacheDir, "*", "*"))
		if err != nil {
			t.Errorf("Error checking cache files: %v", err)
		}
		if len(cacheFiles) == 0 {
			t.Errorf("No cache files were created")
		}
	})

	// Test Phase 2: HTML Generation
	t.Run("html_generation", func(t *testing.T) {
		jsonFiles, err := scanOutputDirectory(*outputDir)
		if err != nil {
			t.Fatalf("Error scanning output directory: %v", err)
		}

		if len(jsonFiles) == 0 {
			t.Fatalf("No JSON files found for HTML generation")
		}

		err = createJSONIndexHTML(*outputDir, jsonFiles)
		if err != nil {
			t.Fatalf("HTML generation failed: %v", err)
		}

		// Note: HTML generation is currently disabled in createJSONIndexHTML
		// but we test that it doesn't error
	})

	// Test Phase 3: Markdown Generation
	t.Run("markdown_generation", func(t *testing.T) {
		err := generateMarkdown(*outputDir, *mdOutputDir)
		if err != nil {
			t.Fatalf("Markdown generation failed: %v", err)
		}

		// Verify markdown files were created
		expectedMdFiles := []string{
			"tutorials/data/documentation/TestFramework.md",
			"tutorials/data/documentation/TestFramework/TestClass.md",
			"index.md",
		}

		for _, file := range expectedMdFiles {
			fullPath := filepath.Join(*mdOutputDir, file)
			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				t.Errorf("Expected markdown file %s was not created", file)
			} else {
				// Check that the file has content
				content, err := os.ReadFile(fullPath)
				if err != nil {
					t.Errorf("Error reading markdown file %s: %v", file, err)
				} else if len(content) == 0 {
					t.Errorf("Markdown file %s is empty", file)
				}
			}
		}

		// Verify index.md has expected structure
		indexPath := filepath.Join(*mdOutputDir, "index.md")
		indexContent, err := os.ReadFile(indexPath)
		if err != nil {
			t.Errorf("Error reading index.md: %v", err)
		} else {
			indexStr := string(indexContent)
			expectedElements := []string{
				"# Apple Documentation",
				"TestFramework",
				"Generated on",
			}
			for _, element := range expectedElements {
				if !strings.Contains(indexStr, element) {
					t.Errorf("Index.md should contain %q", element)
				}
			}
		}
	})
}

// TestPrintURLsOnlyWithMockServer tests the -print-urls functionality
func TestPrintURLsOnlyWithMockServer(t *testing.T) {
	// Store original flag values
	originals := struct {
		baseURL     string
		entryPoint  string
		cacheDir    string
		timeout     time.Duration
		verbose     bool
	}{
		*baseURL,
		*entryPoint,
		*cacheDir,
		*timeout,
		*verbose,
	}

	defer func() {
		*baseURL = originals.baseURL
		*entryPoint = originals.entryPoint
		*cacheDir = originals.cacheDir
		*timeout = originals.timeout
		*verbose = originals.verbose
	}()

	tempDir := t.TempDir()
	*cacheDir = tempDir
	*timeout = 5 * time.Second
	*verbose = false

	// Set up test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		switch r.URL.Path {
		case "/tutorials/data/documentation/technologies.json":
			response := map[string]interface{}{
				"technologies": []interface{}{
					map[string]interface{}{
						"title": "SwiftUI",
						"url":   "documentation/SwiftUI.json",
					},
					map[string]interface{}{
						"title": "UIKit",
						"url":   "documentation/UIKit.json",
					},
				},
			}
			json.NewEncoder(w).Encode(response)
			
		case "/tutorials/data/documentation/SwiftUI.json":
			response := map[string]interface{}{
				"metadata": map[string]interface{}{
					"title": "SwiftUI",
				},
				"references": map[string]interface{}{
					"ref1": map[string]interface{}{
						"type":       "reference",
						"identifier": "doc://com.apple.documentation/documentation/SwiftUI/View",
					},
				},
			}
			json.NewEncoder(w).Encode(response)
			
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	*baseURL = server.URL

	tests := []struct {
		name       string
		entryPoint string
		expectURLs bool
	}{
		{
			name:       "default technologies.json",
			entryPoint: "/tutorials/data/documentation/technologies.json",
			expectURLs: true,
		},
		{
			name:       "specific framework",
			entryPoint: "SwiftUI",
			expectURLs: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			*entryPoint = tt.entryPoint

			ctx := context.Background()
			err := printURLsOnly(ctx)
			if err != nil {
				t.Errorf("printURLsOnly() error = %v", err)
			}
			
			// In a real test, we would capture stdout and verify the URLs
			// For now, we just ensure no error occurs
		})
	}
}

// TestErrorHandlingInRealScenarios tests error handling with various failure modes
func TestErrorHandlingInRealScenarios(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Store original values
	originals := struct {
		outputDir string
		cacheDir  string
		baseURL   string
		timeout   time.Duration
	}{
		*outputDir,
		*cacheDir,
		*baseURL,
		*timeout,
	}

	defer func() {
		*outputDir = originals.outputDir
		*cacheDir = originals.cacheDir
		*baseURL = originals.baseURL
		*timeout = originals.timeout
	}()

	tempDir := t.TempDir()
	*outputDir = filepath.Join(tempDir, "output")
	*cacheDir = filepath.Join(tempDir, "cache")
	*timeout = 2 * time.Second

	tests := []struct {
		name           string
		serverBehavior func(w http.ResponseWriter, r *http.Request)
		expectError    bool
	}{
		{
			name: "server returns 404",
			serverBehavior: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			},
			expectError: true,
		},
		{
			name: "server returns invalid JSON",
			serverBehavior: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, "{invalid json}")
			},
			expectError: false, // Should handle gracefully
		},
		{
			name: "server returns empty response",
			serverBehavior: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, "{}")
			},
			expectError: false,
		},
		{
			name: "server returns 500 error",
			serverBehavior: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(tt.serverBehavior))
			defer server.Close()

			*baseURL = server.URL

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			err := run(ctx)
			
			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			} else if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

// TestConcurrentOperations tests the system under concurrent load
func TestConcurrentOperations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Store original values
	originals := struct {
		outputDir   string
		cacheDir    string
		baseURL     string
		concurrency int
		timeout     time.Duration
	}{
		*outputDir,
		*cacheDir,
		*baseURL,
		*concurrency,
		*timeout,
	}

	defer func() {
		*outputDir = originals.outputDir
		*cacheDir = originals.cacheDir
		*baseURL = originals.baseURL
		*concurrency = originals.concurrency
		*timeout = originals.timeout
	}()

	tempDir := t.TempDir()
	*outputDir = filepath.Join(tempDir, "output")
	*cacheDir = filepath.Join(tempDir, "cache")
	*concurrency = 5 // Higher concurrency for testing
	*timeout = 5 * time.Second

	// Create a server that simulates multiple frameworks
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		switch r.URL.Path {
		case "/tutorials/data/documentation/technologies.json":
			// Return multiple frameworks to test concurrent processing
			technologies := make([]map[string]interface{}, 10)
			for i := 0; i < 10; i++ {
				technologies[i] = map[string]interface{}{
					"title": fmt.Sprintf("Framework%d", i),
					"url":   fmt.Sprintf("documentation/Framework%d.json", i),
				}
			}
			response := map[string]interface{}{
				"technologies": technologies,
			}
			json.NewEncoder(w).Encode(response)
			
		default:
			// Handle framework requests
			if strings.HasPrefix(r.URL.Path, "/tutorials/data/documentation/Framework") {
				response := map[string]interface{}{
					"metadata": map[string]interface{}{
						"title": strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/tutorials/data/documentation/"), ".json"),
					},
					"abstract": []map[string]interface{}{
						{"type": "text", "text": "A test framework"},
					},
				}
				json.NewEncoder(w).Encode(response)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}
	}))
	defer server.Close()

	*baseURL = server.URL

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := run(ctx)
	if err != nil {
		t.Fatalf("Concurrent operations failed: %v", err)
	}

	// Verify that multiple files were processed
	jsonFiles, err := scanOutputDirectory(*outputDir)
	if err != nil {
		t.Fatalf("Error scanning output directory: %v", err)
	}

	if len(jsonFiles) < 5 {
		t.Errorf("Expected at least 5 JSON files, got %d", len(jsonFiles))
	}
}

// TestResourceCleanup tests that resources are properly cleaned up
func TestResourceCleanup(t *testing.T) {
	tempDir := t.TempDir()
	
	// Store original values
	originals := struct {
		outputDir string
		cacheDir  string
		baseURL   string
	}{
		*outputDir,
		*cacheDir,
		*baseURL,
	}

	defer func() {
		*outputDir = originals.outputDir
		*cacheDir = originals.cacheDir
		*baseURL = originals.baseURL
	}()

	*outputDir = filepath.Join(tempDir, "output")
	*cacheDir = filepath.Join(tempDir, "cache")

	// Simple server for testing
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"test": "data"}`)
	}))
	defer server.Close()

	*baseURL = server.URL

	// Create context that times out quickly to test cleanup
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// This should timeout and test cleanup paths
	err := run(ctx)
	
	// Should handle timeout gracefully
	if err != nil && err != context.DeadlineExceeded {
		t.Errorf("Unexpected error during timeout test: %v", err)
	}

	// Verify directories were created (basic cleanup verification)
	if _, err := os.Stat(*outputDir); os.IsNotExist(err) {
		t.Errorf("Output directory should be created even with timeout")
	}
	
	if _, err := os.Stat(*cacheDir); os.IsNotExist(err) {
		t.Errorf("Cache directory should be created even with timeout")
	}
}

// BenchmarkEndToEndWorkflow benchmarks the complete workflow
func BenchmarkEndToEndWorkflow(b *testing.B) {
	if testing.Short() {
		b.Skip("Skipping benchmark in short mode")
	}

	// Store original values
	originals := struct {
		outputDir   string
		cacheDir    string
		mdOutputDir string
		baseURL     string
		concurrency int
		verbose     bool
	}{
		*outputDir,
		*cacheDir,
		*mdOutputDir,
		*baseURL,
		*concurrency,
		*verbose,
	}

	defer func() {
		*outputDir = originals.outputDir
		*cacheDir = originals.cacheDir
		*mdOutputDir = originals.mdOutputDir
		*baseURL = originals.baseURL
		*concurrency = originals.concurrency
		*verbose = originals.verbose
	}()

	*concurrency = 2
	*verbose = false

	// Create a simple test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		switch r.URL.Path {
		case "/tutorials/data/documentation/technologies.json":
			response := map[string]interface{}{
				"technologies": []interface{}{
					map[string]interface{}{
						"title": "TestFramework",
						"url":   "documentation/TestFramework.json",
					},
				},
			}
			json.NewEncoder(w).Encode(response)
			
		case "/tutorials/data/documentation/TestFramework.json":
			response := map[string]interface{}{
				"metadata": map[string]interface{}{
					"title": "TestFramework",
				},
				"abstract": []map[string]interface{}{
					{"type": "text", "text": "A test framework"},
				},
			}
			json.NewEncoder(w).Encode(response)
			
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	*baseURL = server.URL

	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Create fresh temp directories for each iteration
		tempDir := b.TempDir()
		*outputDir = filepath.Join(tempDir, "output")
		*cacheDir = filepath.Join(tempDir, "cache")
		*mdOutputDir = filepath.Join(tempDir, "markdown")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		
		// Run crawl phase
		err := run(ctx)
		if err != nil {
			b.Fatalf("Crawl phase failed: %v", err)
		}

		// Run markdown generation
		err = generateMarkdown(*outputDir, *mdOutputDir)
		if err != nil {
			b.Fatalf("Markdown generation failed: %v", err)
		}

		cancel()
	}
}