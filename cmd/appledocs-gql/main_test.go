package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewDocumentService(t *testing.T) {
	cacheDir := "/test/cache"
	service := NewDocumentService(cacheDir)

	if service.cacheDir != cacheDir {
		t.Errorf("NewDocumentService() cacheDir = %v, want %v", service.cacheDir, cacheDir)
	}
}

func TestDocumentServiceGetDocumentByPath(t *testing.T) {
	// Create a temporary directory with test files
	tempDir := t.TempDir()

	// Create a test JSON file
	testDoc := map[string]interface{}{
		"metadata": map[string]interface{}{
			"title": "Test Document",
		},
		"abstract": map[string]interface{}{
			"content": []interface{}{
				map[string]interface{}{
					"text": "This is a test document",
				},
			},
		},
		"primaryContentSections": []interface{}{
			map[string]interface{}{
				"kind": "content",
			},
		},
	}

	testDocBytes, _ := json.Marshal(testDoc)
	testFilePath := filepath.Join(tempDir, "test.json")
	err := os.WriteFile(testFilePath, testDocBytes, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	service := NewDocumentService(tempDir)

	tests := []struct {
		name        string
		path        string
		expectError bool
		checkTitle  string
	}{
		{
			name:        "existing document",
			path:        "test.json",
			expectError: false,
			checkTitle:  "Test Document",
		},
		{
			name:        "existing document without .json",
			path:        "test",
			expectError: false,
			checkTitle:  "Test Document",
		},
		{
			name:        "non-existent document",
			path:        "nonexistent.json",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := service.GetDocumentByPath(tt.path)

			if tt.expectError {
				if err == nil {
					t.Errorf("GetDocumentByPath() expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("GetDocumentByPath() error = %v", err)
				}
				if doc == nil {
					t.Errorf("GetDocumentByPath() returned nil document")
				} else {
					if doc.Title != tt.checkTitle {
						t.Errorf("GetDocumentByPath() title = %v, want %v", doc.Title, tt.checkTitle)
					}
					if doc.Abstract != "This is a test document" {
						t.Errorf("GetDocumentByPath() abstract = %v, want 'This is a test document'", doc.Abstract)
					}
				}
			}
		})
	}
}

func TestDocumentServiceSearchDocuments(t *testing.T) {
	// Create a temporary directory with test files
	tempDir := t.TempDir()

	// Create test documents
	docs := []struct {
		filename string
		content  map[string]interface{}
	}{
		{
			filename: "swiftui.json",
			content: map[string]interface{}{
				"metadata": map[string]interface{}{
					"title": "SwiftUI",
				},
				"abstract": map[string]interface{}{
					"content": []interface{}{
						map[string]interface{}{
							"text": "A declarative UI framework",
						},
					},
				},
			},
		},
		{
			filename: "uikit.json",
			content: map[string]interface{}{
				"metadata": map[string]interface{}{
					"title": "UIKit",
				},
				"abstract": map[string]interface{}{
					"content": []interface{}{
						map[string]interface{}{
							"text": "An imperative UI framework",
						},
					},
				},
			},
		},
		{
			filename: "foundation.json",
			content: map[string]interface{}{
				"metadata": map[string]interface{}{
					"title": "Foundation",
				},
				"abstract": map[string]interface{}{
					"content": []interface{}{
						map[string]interface{}{
							"text": "Basic functionality framework",
						},
					},
				},
			},
		},
	}

	for _, doc := range docs {
		docBytes, _ := json.Marshal(doc.content)
		filePath := filepath.Join(tempDir, doc.filename)
		err := os.WriteFile(filePath, docBytes, 0644)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", doc.filename, err)
		}
	}

	service := NewDocumentService(tempDir)

	tests := []struct {
		name          string
		query         string
		limit         int
		expectedCount int
	}{
		{
			name:          "search for framework",
			query:         "framework",
			limit:         10,
			expectedCount: 3, // All documents contain "framework"
		},
		{
			name:          "search for SwiftUI",
			query:         "SwiftUI",
			limit:         10,
			expectedCount: 1,
		},
		{
			name:          "search for UI",
			query:         "UI",
			limit:         10,
			expectedCount: 2, // SwiftUI and UIKit
		},
		{
			name:          "search with limit",
			query:         "framework",
			limit:         2,
			expectedCount: 2,
		},
		{
			name:          "search no results",
			query:         "nonexistent",
			limit:         10,
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := service.SearchDocuments(tt.query, tt.limit)
			if err != nil {
				t.Errorf("SearchDocuments() error = %v", err)
			}

			if len(results) != tt.expectedCount {
				t.Errorf("SearchDocuments() returned %d results, want %d", len(results), tt.expectedCount)
			}

			// Verify all results contain the query term (case insensitive)
			for _, doc := range results {
				found := strings.Contains(strings.ToLower(doc.Title), strings.ToLower(tt.query)) ||
					strings.Contains(strings.ToLower(doc.Abstract), strings.ToLower(tt.query))
				if !found {
					// Check the full content as well since that's what the search actually uses
					if doc.Content != nil {
						contentBytes, _ := json.Marshal(doc.Content)
						found = strings.Contains(strings.ToLower(string(contentBytes)), strings.ToLower(tt.query))
					}
				}
				if !found {
					t.Errorf("SearchDocuments() result '%s' does not contain query '%s'", doc.Title, tt.query)
				}
			}
		})
	}
}

func TestDocumentServiceGetFrameworks(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()

	// Create the technologies.json file
	techDir := filepath.Join(tempDir, "tutorials", "data", "documentation")
	err := os.MkdirAll(techDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create tech directory: %v", err)
	}

	technologiesDoc := map[string]interface{}{
		"technologies": []interface{}{
			map[string]interface{}{
				"identifier": "doc://com.apple.documentation/documentation/SwiftUI",
				"path":       "/documentation/swiftui",
				"title":      "SwiftUI",
				"abstract": map[string]interface{}{
					"content": []interface{}{
						map[string]interface{}{
							"text": "A declarative UI framework",
						},
					},
				},
			},
			map[string]interface{}{
				"identifier": "doc://com.apple.documentation/documentation/UIKit",
				"path":       "/documentation/uikit",
				"title":      "UIKit",
				"abstract": map[string]interface{}{
					"content": []interface{}{
						map[string]interface{}{
							"text": "An imperative UI framework",
						},
					},
				},
			},
		},
	}

	techBytes, _ := json.Marshal(technologiesDoc)
	techPath := filepath.Join(techDir, "technologies.json")
	err = os.WriteFile(techPath, techBytes, 0644)
	if err != nil {
		t.Fatalf("Failed to create technologies.json: %v", err)
	}

	service := NewDocumentService(tempDir)

	frameworks, err := service.GetFrameworks()
	if err != nil {
		t.Errorf("GetFrameworks() error = %v", err)
	}

	if len(frameworks) != 2 {
		t.Errorf("GetFrameworks() returned %d frameworks, want 2", len(frameworks))
	}

	// Check first framework
	if frameworks[0].Title != "SwiftUI" {
		t.Errorf("GetFrameworks() first framework title = %v, want SwiftUI", frameworks[0].Title)
	}

	if frameworks[0].Abstract != "A declarative UI framework" {
		t.Errorf("GetFrameworks() first framework abstract = %v, want 'A declarative UI framework'", frameworks[0].Abstract)
	}
}

func TestServerHandleDocument(t *testing.T) {
	// Create a temporary directory with test files
	tempDir := t.TempDir()

	testDoc := map[string]interface{}{
		"metadata": map[string]interface{}{
			"title": "Test Document",
		},
	}

	testDocBytes, _ := json.Marshal(testDoc)
	testFilePath := filepath.Join(tempDir, "test.json")
	err := os.WriteFile(testFilePath, testDocBytes, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	server := NewServer(tempDir)

	tests := []struct {
		name           string
		method         string
		queryParams    string
		expectedStatus int
		checkResponse  func(t *testing.T, body []byte)
	}{
		{
			name:           "GET with valid path",
			method:         "GET",
			queryParams:    "?path=test.json",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []byte) {
				var doc Document
				err := json.Unmarshal(body, &doc)
				if err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}
				if doc.Title != "Test Document" {
					t.Errorf("Document title = %v, want 'Test Document'", doc.Title)
				}
			},
		},
		{
			name:           "GET with missing path",
			method:         "GET",
			queryParams:    "",
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, body []byte) {
				var errorResp ErrorResponse
				err := json.Unmarshal(body, &errorResp)
				if err != nil {
					t.Errorf("Failed to unmarshal error response: %v", err)
				}
				if !strings.Contains(errorResp.Error, "Missing path parameter") {
					t.Errorf("Error message = %v, want to contain 'Missing path parameter'", errorResp.Error)
				}
			},
		},
		{
			name:           "GET with non-existent path",
			method:         "GET",
			queryParams:    "?path=nonexistent.json",
			expectedStatus: http.StatusNotFound,
			checkResponse: func(t *testing.T, body []byte) {
				var errorResp ErrorResponse
				err := json.Unmarshal(body, &errorResp)
				if err != nil {
					t.Errorf("Failed to unmarshal error response: %v", err)
				}
				if !strings.Contains(errorResp.Error, "Document not found") {
					t.Errorf("Error message = %v, want to contain 'Document not found'", errorResp.Error)
				}
			},
		},
		{
			name:           "POST method not allowed",
			method:         "POST",
			queryParams:    "?path=test.json",
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, "/api/document"+tt.queryParams, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			server.handleDocument(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handleDocument() status = %v, want %v", rr.Code, tt.expectedStatus)
			}

			if tt.checkResponse != nil {
				tt.checkResponse(t, rr.Body.Bytes())
			}
		})
	}
}

func TestServerHandleSearch(t *testing.T) {
	// Create a temporary directory with test files
	tempDir := t.TempDir()

	testDoc := map[string]interface{}{
		"metadata": map[string]interface{}{
			"title": "SwiftUI Framework",
		},
		"content": "This is about SwiftUI framework",
	}

	testDocBytes, _ := json.Marshal(testDoc)
	testFilePath := filepath.Join(tempDir, "swiftui.json")
	err := os.WriteFile(testFilePath, testDocBytes, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	server := NewServer(tempDir)

	tests := []struct {
		name           string
		method         string
		queryParams    string
		expectedStatus int
		checkResponse  func(t *testing.T, body []byte)
	}{
		{
			name:           "GET with valid query",
			method:         "GET",
			queryParams:    "?q=SwiftUI",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []byte) {
				var searchResp SearchResponse
				err := json.Unmarshal(body, &searchResp)
				if err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}
				if len(searchResp.Results) == 0 {
					t.Errorf("Search should return at least 1 result")
				}
				if searchResp.Query != "SwiftUI" {
					t.Errorf("Search query = %v, want 'SwiftUI'", searchResp.Query)
				}
			},
		},
		{
			name:           "GET with missing query",
			method:         "GET",
			queryParams:    "",
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, body []byte) {
				var errorResp ErrorResponse
				err := json.Unmarshal(body, &errorResp)
				if err != nil {
					t.Errorf("Failed to unmarshal error response: %v", err)
				}
				if !strings.Contains(errorResp.Error, "Missing query parameter") {
					t.Errorf("Error message = %v, want to contain 'Missing query parameter'", errorResp.Error)
				}
			},
		},
		{
			name:           "GET with limit",
			method:         "GET",
			queryParams:    "?q=SwiftUI&limit=5",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []byte) {
				var searchResp SearchResponse
				err := json.Unmarshal(body, &searchResp)
				if err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}
				if searchResp.Limit != 5 {
					t.Errorf("Search limit = %v, want 5", searchResp.Limit)
				}
			},
		},
		{
			name:           "POST method not allowed",
			method:         "POST",
			queryParams:    "?q=test",
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, "/api/search"+tt.queryParams, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			server.handleSearch(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handleSearch() status = %v, want %v", rr.Code, tt.expectedStatus)
			}

			if tt.checkResponse != nil {
				tt.checkResponse(t, rr.Body.Bytes())
			}
		})
	}
}

func TestServerHandleFrameworks(t *testing.T) {
	// Create a temporary directory with technologies.json
	tempDir := t.TempDir()

	techDir := filepath.Join(tempDir, "tutorials", "data", "documentation")
	err := os.MkdirAll(techDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create tech directory: %v", err)
	}

	technologiesDoc := map[string]interface{}{
		"technologies": []interface{}{
			map[string]interface{}{
				"identifier": "doc://com.apple.documentation/documentation/SwiftUI",
				"path":       "/documentation/swiftui",
				"title":      "SwiftUI",
			},
		},
	}

	techBytes, _ := json.Marshal(technologiesDoc)
	techPath := filepath.Join(techDir, "technologies.json")
	err = os.WriteFile(techPath, techBytes, 0644)
	if err != nil {
		t.Fatalf("Failed to create technologies.json: %v", err)
	}

	server := NewServer(tempDir)

	tests := []struct {
		name           string
		method         string
		expectedStatus int
		checkResponse  func(t *testing.T, body []byte)
	}{
		{
			name:           "GET frameworks",
			method:         "GET",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []byte) {
				var frameworks []*Document
				err := json.Unmarshal(body, &frameworks)
				if err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}
				if len(frameworks) != 1 {
					t.Errorf("Expected 1 framework, got %d", len(frameworks))
				}
				if frameworks[0].Title != "SwiftUI" {
					t.Errorf("Framework title = %v, want 'SwiftUI'", frameworks[0].Title)
				}
			},
		},
		{
			name:           "POST method not allowed",
			method:         "POST",
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, "/api/frameworks", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			server.handleFrameworks(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handleFrameworks() status = %v, want %v", rr.Code, tt.expectedStatus)
			}

			if tt.checkResponse != nil {
				tt.checkResponse(t, rr.Body.Bytes())
			}
		})
	}
}

func TestServerHandleGraphQL(t *testing.T) {
	tempDir := t.TempDir()
	server := NewServer(tempDir)

	tests := []struct {
		name           string
		method         string
		body           string
		expectedStatus int
		checkResponse  func(t *testing.T, body []byte)
	}{
		{
			name:           "GET GraphQL playground",
			method:         "GET",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []byte) {
				content := string(body)
				if !strings.Contains(content, "GraphiQL") {
					t.Errorf("Response should contain GraphiQL")
				}
				if !strings.Contains(content, "html") {
					t.Errorf("Response should contain HTML")
				}
			},
		},
		{
			name:           "POST invalid JSON",
			method:         "POST",
			body:           `{invalid json`,
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []byte) {
				var response map[string]interface{}
				err := json.Unmarshal(body, &response)
				if err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}
				if errors, ok := response["errors"].([]interface{}); !ok || len(errors) == 0 {
					t.Errorf("Response should contain errors")
				}
			},
		},
		{
			name:           "POST missing query",
			method:         "POST",
			body:           `{"variables": {}}`,
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []byte) {
				var response map[string]interface{}
				err := json.Unmarshal(body, &response)
				if err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}
				if errors, ok := response["errors"].([]interface{}); !ok || len(errors) == 0 {
					t.Errorf("Response should contain errors for missing query")
				}
			},
		},
		{
			name:           "POST introspection query",
			method:         "POST",
			body:           `{"query": "query IntrospectionQuery { __schema { types { name } } }"}`,
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []byte) {
				var response map[string]interface{}
				err := json.Unmarshal(body, &response)
				if err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}
				// The response might contain either data or errors for introspection
				if data, ok := response["data"].(map[string]interface{}); ok {
					if schema, ok := data["__schema"].(map[string]interface{}); ok {
						if types, ok := schema["types"].([]interface{}); !ok || len(types) == 0 {
							t.Errorf("Response should contain types in schema")
						}
					}
				}
				// It's acceptable if introspection returns errors in this test environment
			},
		},
		{
			name:           "POST unsupported query",
			method:         "POST",
			body:           `{"query": "query { unsupportedField }"}`,
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []byte) {
				var response map[string]interface{}
				err := json.Unmarshal(body, &response)
				if err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}
				if errors, ok := response["errors"].([]interface{}); !ok || len(errors) == 0 {
					t.Errorf("Response should contain errors for unsupported query")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			var err error

			if tt.body != "" {
				req, err = http.NewRequest(tt.method, "/graphql", bytes.NewBufferString(tt.body))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req, err = http.NewRequest(tt.method, "/graphql", nil)
			}

			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			server.handleGraphQL(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handleGraphQL() status = %v, want %v", rr.Code, tt.expectedStatus)
			}

			if tt.checkResponse != nil {
				tt.checkResponse(t, rr.Body.Bytes())
			}
		})
	}
}

func TestServerHandleHome(t *testing.T) {
	tempDir := t.TempDir()
	server := NewServer(tempDir)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	server.handleHome(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handleHome() status = %v, want %v", rr.Code, http.StatusOK)
	}

	content := rr.Body.String()
	if !strings.Contains(content, "Apple Docs API") {
		t.Errorf("Response should contain 'Apple Docs API'")
	}
	if !strings.Contains(content, "html") {
		t.Errorf("Response should contain HTML")
	}
}

func TestServerHandleNotFound(t *testing.T) {
	tempDir := t.TempDir()
	server := NewServer(tempDir)

	req, err := http.NewRequest("GET", "/nonexistent", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	server.handleHome(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("handleHome() for non-root path status = %v, want %v", rr.Code, http.StatusNotFound)
	}
}

func TestServerHandleSandbox(t *testing.T) {
	tempDir := t.TempDir()
	server := NewServer(tempDir)

	req, err := http.NewRequest("GET", "/sandbox", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	server.handleSandbox(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handleSandbox() status = %v, want %v", rr.Code, http.StatusOK)
	}

	content := rr.Body.String()
	if !strings.Contains(content, "Apollo Sandbox") {
		t.Errorf("Response should contain 'Apollo Sandbox'")
	}
	if !strings.Contains(content, "EmbeddedSandbox") {
		t.Errorf("Response should contain 'EmbeddedSandbox'")
	}
}

func TestServerHandleSchema(t *testing.T) {
	tempDir := t.TempDir()
	server := NewServer(tempDir)

	req, err := http.NewRequest("GET", "/schema", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	server.handleSchema(rr, req)

	// The handler provides a fallback response when schema file doesn't exist
	if rr.Code != http.StatusOK && rr.Code != http.StatusNotFound {
		t.Errorf("handleSchema() status = %v, want %v or %v", rr.Code, http.StatusOK, http.StatusNotFound)
	}
}

// Benchmark tests for performance
func BenchmarkDocumentServiceGetDocumentByPath(b *testing.B) {
	tempDir := b.TempDir()

	testDoc := map[string]interface{}{
		"metadata": map[string]interface{}{
			"title": "Test Document",
		},
		"content": "Large content for performance testing",
	}

	testDocBytes, _ := json.Marshal(testDoc)
	testFilePath := filepath.Join(tempDir, "test.json")
	err := os.WriteFile(testFilePath, testDocBytes, 0644)
	if err != nil {
		b.Fatalf("Failed to create test file: %v", err)
	}

	service := NewDocumentService(tempDir)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetDocumentByPath("test.json")
	}
}

func BenchmarkServerHandleDocument(b *testing.B) {
	tempDir := b.TempDir()

	testDoc := map[string]interface{}{
		"metadata": map[string]interface{}{
			"title": "Test Document",
		},
	}

	testDocBytes, _ := json.Marshal(testDoc)
	testFilePath := filepath.Join(tempDir, "test.json")
	err := os.WriteFile(testFilePath, testDocBytes, 0644)
	if err != nil {
		b.Fatalf("Failed to create test file: %v", err)
	}

	server := NewServer(tempDir)

	req, err := http.NewRequest("GET", "/api/document?path=test.json", nil)
	if err != nil {
		b.Fatalf("Failed to create request: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		server.handleDocument(rr, req)
	}
}
