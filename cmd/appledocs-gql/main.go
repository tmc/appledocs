package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

//go:embed templates/*.html
var templateFS embed.FS

// Document represents an Apple documentation file
type Document struct {
	ID       string                 `json:"id"`
	Path     string                 `json:"path"`
	Title    string                 `json:"title,omitempty"`
	Abstract string                 `json:"abstract,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Content  map[string]interface{} `json:"content,omitempty"`
}

// SearchResponse wraps search results
type SearchResponse struct {
	Results []*Document `json:"results"`
	Count   int         `json:"count"`
	Query   string      `json:"query"`
	Limit   int         `json:"limit"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// DocumentService handles document retrieval and search operations
type DocumentService struct {
	cacheDir string
}

// NewDocumentService creates a new document service
func NewDocumentService(cacheDir string) *DocumentService {
	return &DocumentService{
		cacheDir: cacheDir,
	}
}

// GetDocumentByPath retrieves a document by its path
func (s *DocumentService) GetDocumentByPath(path string) (*Document, error) {
	fullPath := filepath.Join(s.cacheDir, path)

	// If the path doesn't end with .json, append it
	if !strings.HasSuffix(fullPath, ".json") {
		fullPath += ".json"
	}

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	var content map[string]interface{}
	if err := json.Unmarshal(data, &content); err != nil {
		return nil, err
	}

	// Extract title and abstract from metadata if available
	var title, abstract string
	if metadata, ok := content["metadata"].(map[string]interface{}); ok {
		if t, ok := metadata["title"].(string); ok {
			title = t
		}
	}

	if sections, ok := content["abstract"].(map[string]interface{}); ok {
		if contentArr, ok := sections["content"].([]interface{}); ok && len(contentArr) > 0 {
			if text, ok := contentArr[0].(map[string]interface{}); ok {
				if t, ok := text["text"].(string); ok {
					abstract = t
				}
			}
		}
	}

	// Create document object
	doc := &Document{
		ID:       path,
		Path:     path,
		Title:    title,
		Abstract: abstract,
	}

	// Add metadata if available
	if metadata, ok := content["metadata"].(map[string]interface{}); ok {
		doc.Metadata = metadata
	}

	// Add full content
	doc.Content = content

	return doc, nil
}

// SearchDocuments searches for documents matching the query
func (s *DocumentService) SearchDocuments(query string, limit int) ([]*Document, error) {
	results := []*Document{}
	query = strings.ToLower(query)

	// Simple file search for now, can be improved later
	err := filepath.Walk(s.cacheDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}

		// Read the file and check if it contains the query
		data, err := os.ReadFile(path)
		if err != nil {
			return nil // Skip files we can't read
		}

		if strings.Contains(strings.ToLower(string(data)), query) {
			relativePath, err := filepath.Rel(s.cacheDir, path)
			if err != nil {
				return nil
			}

			doc, err := s.GetDocumentByPath(relativePath)
			if err != nil {
				return nil
			}

			results = append(results, doc)

			if limit > 0 && len(results) >= limit {
				return filepath.SkipAll
			}
		}

		return nil
	})

	return results, err
}

// GetFrameworks retrieves top-level frameworks
func (s *DocumentService) GetFrameworks() ([]*Document, error) {
	techPath := filepath.Join(s.cacheDir, "tutorials", "data", "documentation", "technologies.json")

	data, err := os.ReadFile(techPath)
	if err != nil {
		return nil, err
	}

	var content map[string]interface{}
	if err := json.Unmarshal(data, &content); err != nil {
		return nil, err
	}

	frameworks := []*Document{}

	if technologies, ok := content["technologies"].([]interface{}); ok {
		for _, tech := range technologies {
			if t, ok := tech.(map[string]interface{}); ok {
				id, _ := t["identifier"].(string)
				path, _ := t["path"].(string)
				title, _ := t["title"].(string)

				framework := &Document{
					ID:    id,
					Path:  path,
					Title: title,
				}

				if abstract, ok := t["abstract"].(map[string]interface{}); ok {
					if content, ok := abstract["content"].([]interface{}); ok && len(content) > 0 {
						if text, ok := content[0].(map[string]interface{}); ok {
							if t, ok := text["text"].(string); ok {
								framework.Abstract = t
							}
						}
					}
				}

				frameworks = append(frameworks, framework)
			}
		}
	}

	return frameworks, nil
}

// Server represents our API server
type Server struct {
	documentService *DocumentService
}

// NewServer creates a new server
func NewServer(cacheDir string) *Server {
	return &Server{
		documentService: NewDocumentService(cacheDir),
	}
}

// handleDocument handles requests to get a document by path
func (s *Server) handleDocument(w http.ResponseWriter, r *http.Request) {
	// Only support GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the path parameter
	path := r.URL.Query().Get("path")
	if path == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Missing path parameter"})
		return
	}

	// Get the document
	doc, err := s.documentService.GetDocumentByPath(path)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{Error: fmt.Sprintf("Document not found: %v", err)})
		return
	}

	// Return the document as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doc)
}

// handleSearch handles search requests
func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request) {
	// Only support GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the query parameter
	query := r.URL.Query().Get("q")
	if query == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Missing query parameter"})
		return
	}

	// Get the limit parameter
	limitStr := r.URL.Query().Get("limit")
	limit := 10 // Default limit
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	// Perform the search
	results, err := s.documentService.SearchDocuments(query, limit)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: fmt.Sprintf("Search failed: %v", err)})
		return
	}

	// Return the results as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SearchResponse{
		Results: results,
		Count:   len(results),
		Query:   query,
		Limit:   limit,
	})
}

// handleFrameworks handles requests to get all frameworks
func (s *Server) handleFrameworks(w http.ResponseWriter, r *http.Request) {
	// Only support GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get all frameworks
	frameworks, err := s.documentService.GetFrameworks()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: fmt.Sprintf("Failed to get frameworks: %v", err)})
		return
	}

	// Return the frameworks as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(frameworks)
}

// handleGraphQL serves the GraphQL playground and processes GraphQL queries
func (s *Server) handleGraphQL(w http.ResponseWriter, r *http.Request) {
	// If this is a POST request, handle it as a GraphQL query
	if r.Method == http.MethodPost {
		s.processGraphQLQuery(w, r)
		return
	}

	// For GET requests, serve the GraphiQL interface
	htmlContent, err := templateFS.ReadFile("templates/graphql.html")
	if err != nil {
		// If we still couldn't find the HTML file, provide a basic interface
		html := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>GraphiQL - Apple Docs API</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/graphiql@1.0.0/graphiql.min.css" />
    <script src="https://cdn.jsdelivr.net/npm/react@16.13.1/umd/react.production.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/react-dom@16.13.1/umd/react-dom.production.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/graphiql@1.0.0/graphiql.min.js"></script>
    <style>
        body, html { margin: 0; padding: 0; height: 100%; overflow: hidden; }
        nav {
            background-color: #f5f5f5;
            border-bottom: 1px solid #e0e0e0;
            padding: 10px;
            display: flex;
            gap: 20px;
        }
        nav a {
            text-decoration: none;
            color: #333;
            font-weight: 500;
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
        }
        #graphiql { 
            height: calc(100vh - 41px);
        }
    </style>
</head>
<body>
    <nav>
        <a href="/">Home</a>
        <a href="/graphql">GraphiQL</a>
        <a href="/sandbox">Apollo Sandbox</a>
        <a href="/schema">Schema</a>
    </nav>
    <div id="graphiql">Loading GraphiQL...</div>
    <script>
        function graphQLFetcher(params) {
            return fetch('/graphql', {
                method: 'post',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(params)
            }).then(response => response.json());
        }
        ReactDOM.render(
            React.createElement(GraphiQL, { fetcher: graphQLFetcher }),
            document.getElementById('graphiql')
        );
    </script>
</body>
</html>`
		htmlContent = []byte(html)
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlContent)
}

// handleHome serves the home page with API documentation
func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Read template from embedded filesystem
	htmlContent, err := templateFS.ReadFile("templates/home.html")
	if err != nil {
		// If we still couldn't find the HTML file, provide a basic interface
		html := `<!DOCTYPE html>
<html>
<head>
    <title>Apple Docs API</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
        }
        h1 {
            border-bottom: 1px solid #eee;
            padding-bottom: 10px;
        }
        pre {
            background-color: #f5f5f5;
            padding: 15px;
            border-radius: 5px;
            overflow-x: auto;
        }
        .btn {
            display: inline-block;
            padding: 8px 16px;
            background-color: #0366d6;
            color: white;
            border-radius: 4px;
            text-decoration: none;
            margin: 20px 0;
        }
    </style>
</head>
<body>
    <h1>Apple Docs API</h1>
    <p>This API provides access to the cached Apple documentation.</p>
    
    <a href="/graphql" class="btn">GraphQL Playground</a>
    
    <h2>API Endpoints</h2>
    <ul>
        <li><code>GET /api/document?path=...</code> - Get a document by path</li>
        <li><code>GET /api/search?q=...&limit=...</code> - Search for documents</li>
        <li><code>GET /api/frameworks</code> - List all frameworks</li>
        <li><code>GET /graphql</code> - GraphQL interface</li>
    </ul>
</body>
</html>`
		htmlContent = []byte(html)
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlContent)
}

// handleSandbox serves the Apollo Sandbox playground
func (s *Server) handleSandbox(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Apollo Sandbox - Apple Docs API</title>
    <style>
        body, html { margin: 0; padding: 0; height: 100%; overflow: hidden; }
        nav {
            background-color: #f5f5f5;
            border-bottom: 1px solid #e0e0e0;
            padding: 10px;
            display: flex;
            gap: 20px;
        }
        nav a {
            text-decoration: none;
            color: #333;
            font-weight: 500;
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
        }
        .sandbox-container { 
            width: 100%; 
            height: calc(100vh - 41px);
        }
    </style>
</head>
<body>
    <nav>
        <a href="/">Home</a>
        <a href="/graphql">GraphiQL</a>
        <a href="/sandbox">Apollo Sandbox</a>
        <a href="/schema">Schema</a>
    </nav>
    <div id="embedded-sandbox" class="sandbox-container"></div>
    
    <script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
    <script>
      new window.EmbeddedSandbox({
        target: '#embedded-sandbox',
        initialEndpoint: window.location.protocol + '//' + window.location.host + '/graphql',
      });
    </script>
</body>
</html>`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

// handleSchema serves the schema documentation
func (s *Server) handleSchema(w http.ResponseWriter, r *http.Request) {
	// Read template from embedded filesystem
	htmlContent, err := templateFS.ReadFile("templates/schema.html")
	if err != nil {
		// If that fails, serve the raw schema GraphQL file
		schemaContent, schemaErr := os.ReadFile("enhanced-schema.graphql")

		if schemaErr == nil {
			w.Header().Set("Content-Type", "text/plain")
			w.Write(schemaContent)
			return
		}

		// If all else fails, return a 404
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlContent)
}

// processGraphQLQuery handles GraphQL query execution including introspection
func (s *Server) processGraphQLQuery(w http.ResponseWriter, r *http.Request) {
	// Set JSON content type for all responses
	w.Header().Set("Content-Type", "application/json")

	// Parse the request body
	var params map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": []map[string]interface{}{
				{"message": "Invalid request format: " + err.Error()},
			},
		})
		return
	}

	// Extract the query
	query, ok := params["query"].(string)
	if !ok {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": []map[string]interface{}{
				{"message": "Query parameter is required"},
			},
		})
		return
	}

	// Handle introspection queries
	if strings.Contains(query, "__schema") || strings.Contains(query, "__type") {
		s.handleIntrospection(w, query, params)
		return
	}

	// Handle regular queries
	if strings.Contains(query, "document(path:") {
		s.handleDocumentQuery(w, query, params)
		return
	}

	if strings.Contains(query, "search(query:") {
		s.handleSearchQuery(w, query, params)
		return
	}

	if strings.Contains(query, "frameworks") {
		s.handleFrameworksQuery(w, query, params)
		return
	}

	// Unsupported query
	json.NewEncoder(w).Encode(map[string]interface{}{
		"errors": []map[string]interface{}{
			{"message": "Unsupported query type"},
		},
	})
}

// handleIntrospection processes GraphQL introspection queries
func (s *Server) handleIntrospection(w http.ResponseWriter, query string, params map[string]interface{}) {
	// Verify schema file exists
	_, err := os.ReadFile("enhanced-schema.graphql")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": []map[string]interface{}{
				{"message": "Schema file not found: " + err.Error()},
			},
		})
		return
	}

	// Handle __schema query
	if strings.Contains(query, "__schema") {
		// Build a basic schema response
		queryType := map[string]interface{}{
			"name": "Query",
			"kind": "OBJECT",
		}

		types := []map[string]interface{}{
			{
				"name":        "Query",
				"kind":        "OBJECT",
				"description": "Root query type",
				"fields": []map[string]interface{}{
					{
						"name":        "document",
						"description": "Get a document by its path",
						"args": []map[string]interface{}{
							{
								"name":        "path",
								"description": "Path to the document",
								"type": map[string]interface{}{
									"kind": "NON_NULL",
									"ofType": map[string]interface{}{
										"kind": "SCALAR",
										"name": "String",
									},
								},
							},
						},
						"type": map[string]interface{}{
							"kind": "OBJECT",
							"name": "Document",
						},
					},
					{
						"name":        "search",
						"description": "Search for documents",
						"args": []map[string]interface{}{
							{
								"name":        "query",
								"description": "Search query",
								"type": map[string]interface{}{
									"kind": "NON_NULL",
									"ofType": map[string]interface{}{
										"kind": "SCALAR",
										"name": "String",
									},
								},
							},
							{
								"name":        "limit",
								"description": "Maximum number of results",
								"type": map[string]interface{}{
									"kind": "SCALAR",
									"name": "Int",
								},
							},
						},
						"type": map[string]interface{}{
							"kind": "OBJECT",
							"name": "SearchResult",
						},
					},
					{
						"name":        "frameworks",
						"description": "List all frameworks",
						"args":        []map[string]interface{}{},
						"type": map[string]interface{}{
							"kind": "NON_NULL",
							"ofType": map[string]interface{}{
								"kind": "LIST",
								"ofType": map[string]interface{}{
									"kind": "NON_NULL",
									"ofType": map[string]interface{}{
										"kind": "OBJECT",
										"name": "Framework",
									},
								},
							},
						},
					},
				},
			},
			{
				"name":        "Document",
				"kind":        "OBJECT",
				"description": "Apple documentation item",
				"fields": []map[string]interface{}{
					{
						"name": "id",
						"type": map[string]interface{}{
							"kind": "NON_NULL",
							"ofType": map[string]interface{}{
								"kind": "SCALAR",
								"name": "String",
							},
						},
					},
					{
						"name": "path",
						"type": map[string]interface{}{
							"kind": "NON_NULL",
							"ofType": map[string]interface{}{
								"kind": "SCALAR",
								"name": "String",
							},
						},
					},
					{
						"name": "title",
						"type": map[string]interface{}{
							"kind": "SCALAR",
							"name": "String",
						},
					},
					{
						"name": "abstract",
						"type": map[string]interface{}{
							"kind": "SCALAR",
							"name": "String",
						},
					},
				},
			},
			{
				"name":        "SearchResult",
				"kind":        "OBJECT",
				"description": "Search result",
				"fields": []map[string]interface{}{
					{
						"name": "documents",
						"type": map[string]interface{}{
							"kind": "NON_NULL",
							"ofType": map[string]interface{}{
								"kind": "LIST",
								"ofType": map[string]interface{}{
									"kind": "NON_NULL",
									"ofType": map[string]interface{}{
										"kind": "OBJECT",
										"name": "Document",
									},
								},
							},
						},
					},
					{
						"name": "count",
						"type": map[string]interface{}{
							"kind": "NON_NULL",
							"ofType": map[string]interface{}{
								"kind": "SCALAR",
								"name": "Int",
							},
						},
					},
					{
						"name": "query",
						"type": map[string]interface{}{
							"kind": "NON_NULL",
							"ofType": map[string]interface{}{
								"kind": "SCALAR",
								"name": "String",
							},
						},
					},
				},
			},
			{
				"name":        "Framework",
				"kind":        "OBJECT",
				"description": "Framework information",
				"fields": []map[string]interface{}{
					{
						"name": "id",
						"type": map[string]interface{}{
							"kind": "NON_NULL",
							"ofType": map[string]interface{}{
								"kind": "SCALAR",
								"name": "String",
							},
						},
					},
					{
						"name": "path",
						"type": map[string]interface{}{
							"kind": "NON_NULL",
							"ofType": map[string]interface{}{
								"kind": "SCALAR",
								"name": "String",
							},
						},
					},
					{
						"name": "title",
						"type": map[string]interface{}{
							"kind": "NON_NULL",
							"ofType": map[string]interface{}{
								"kind": "SCALAR",
								"name": "String",
							},
						},
					},
					{
						"name": "abstract",
						"type": map[string]interface{}{
							"kind": "SCALAR",
							"name": "String",
						},
					},
				},
			},
		}

		// Add scalar types
		for _, name := range []string{"String", "Int", "Float", "Boolean", "ID", "JSON"} {
			types = append(types, map[string]interface{}{
				"name": name,
				"kind": "SCALAR",
			})
		}

		schemaResponse := map[string]interface{}{
			"queryType": queryType,
			"types":     types,
			"directives": []map[string]interface{}{
				{
					"name":        "include",
					"description": "Directs the executor to include this field or fragment only when the argument is true.",
					"locations":   []string{"FIELD", "FRAGMENT_SPREAD", "INLINE_FRAGMENT"},
					"args": []map[string]interface{}{
						{
							"name":        "if",
							"description": "Included when true.",
							"type": map[string]interface{}{
								"kind": "NON_NULL",
								"ofType": map[string]interface{}{
									"kind": "SCALAR",
									"name": "Boolean",
								},
							},
						},
					},
				},
				{
					"name":        "skip",
					"description": "Directs the executor to skip this field or fragment when the argument is true.",
					"locations":   []string{"FIELD", "FRAGMENT_SPREAD", "INLINE_FRAGMENT"},
					"args": []map[string]interface{}{
						{
							"name":        "if",
							"description": "Skipped when true.",
							"type": map[string]interface{}{
								"kind": "NON_NULL",
								"ofType": map[string]interface{}{
									"kind": "SCALAR",
									"name": "Boolean",
								},
							},
						},
					},
				},
			},
			"mutationType": nil,
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": map[string]interface{}{
				"__schema": schemaResponse,
			},
		})
		return
	}

	// Handle __type query
	if strings.Contains(query, "__type") {
		// Extract type name
		typeNameMatch := regexp.MustCompile(`__type\(name:\s*"([^"]+)"`).FindStringSubmatch(query)
		if len(typeNameMatch) < 2 {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"errors": []map[string]interface{}{
					{"message": "Invalid __type query format"},
				},
			})
			return
		}

		typeName := typeNameMatch[1]
		typeInfo := map[string]interface{}{}

		// Basic type info for common types
		switch typeName {
		case "Query":
			typeInfo = map[string]interface{}{
				"name":        "Query",
				"kind":        "OBJECT",
				"description": "Root query type",
				"fields": []map[string]interface{}{
					{
						"name":        "document",
						"description": "Get a document by its path",
						"args": []map[string]interface{}{
							{
								"name":        "path",
								"description": "Path to the document",
								"type": map[string]interface{}{
									"kind": "NON_NULL",
									"ofType": map[string]interface{}{
										"kind": "SCALAR",
										"name": "String",
									},
								},
							},
						},
						"type": map[string]interface{}{
							"kind": "OBJECT",
							"name": "Document",
						},
					},
					{
						"name":        "search",
						"description": "Search for documents",
						"args": []map[string]interface{}{
							{
								"name":        "query",
								"description": "Search query",
								"type": map[string]interface{}{
									"kind": "NON_NULL",
									"ofType": map[string]interface{}{
										"kind": "SCALAR",
										"name": "String",
									},
								},
							},
							{
								"name":        "limit",
								"description": "Maximum number of results",
								"type": map[string]interface{}{
									"kind": "SCALAR",
									"name": "Int",
								},
							},
						},
						"type": map[string]interface{}{
							"kind": "OBJECT",
							"name": "SearchResult",
						},
					},
					{
						"name":        "frameworks",
						"description": "List all frameworks",
						"args":        []map[string]interface{}{},
						"type": map[string]interface{}{
							"kind": "NON_NULL",
							"ofType": map[string]interface{}{
								"kind": "LIST",
								"ofType": map[string]interface{}{
									"kind": "NON_NULL",
									"ofType": map[string]interface{}{
										"kind": "OBJECT",
										"name": "Framework",
									},
								},
							},
						},
					},
				},
			}
		case "Document":
			typeInfo = map[string]interface{}{
				"name":        "Document",
				"kind":        "OBJECT",
				"description": "Apple documentation item",
				"fields": []map[string]interface{}{
					{
						"name": "id",
						"type": map[string]interface{}{
							"kind": "NON_NULL",
							"ofType": map[string]interface{}{
								"kind": "SCALAR",
								"name": "String",
							},
						},
					},
					{
						"name": "path",
						"type": map[string]interface{}{
							"kind": "NON_NULL",
							"ofType": map[string]interface{}{
								"kind": "SCALAR",
								"name": "String",
							},
						},
					},
					{
						"name": "title",
						"type": map[string]interface{}{
							"kind": "SCALAR",
							"name": "String",
						},
					},
					{
						"name": "abstract",
						"type": map[string]interface{}{
							"kind": "SCALAR",
							"name": "String",
						},
					},
				},
			}
		case "SearchResult":
			typeInfo = map[string]interface{}{
				"name":        "SearchResult",
				"kind":        "OBJECT",
				"description": "Search result",
				"fields": []map[string]interface{}{
					{
						"name": "documents",
						"type": map[string]interface{}{
							"kind": "NON_NULL",
							"ofType": map[string]interface{}{
								"kind": "LIST",
								"ofType": map[string]interface{}{
									"kind": "NON_NULL",
									"ofType": map[string]interface{}{
										"kind": "OBJECT",
										"name": "Document",
									},
								},
							},
						},
					},
					{
						"name": "count",
						"type": map[string]interface{}{
							"kind": "NON_NULL",
							"ofType": map[string]interface{}{
								"kind": "SCALAR",
								"name": "Int",
							},
						},
					},
					{
						"name": "query",
						"type": map[string]interface{}{
							"kind": "NON_NULL",
							"ofType": map[string]interface{}{
								"kind": "SCALAR",
								"name": "String",
							},
						},
					},
				},
			}
		case "String", "Int", "Float", "Boolean", "ID", "JSON":
			typeInfo = map[string]interface{}{
				"name": typeName,
				"kind": "SCALAR",
			}
		default:
			// Return null for unknown types
			typeInfo = nil
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": map[string]interface{}{
				"__type": typeInfo,
			},
		})
		return
	}
}

// handleDocumentQuery processes document queries
func (s *Server) handleDocumentQuery(w http.ResponseWriter, query string, params map[string]interface{}) {
	// Extract the path parameter from the query
	pathMatch := regexp.MustCompile(`document\(path:\s*"([^"]+)"`).FindStringSubmatch(query)
	if len(pathMatch) < 2 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": []map[string]interface{}{
				{"message": "Invalid document query format"},
			},
		})
		return
	}

	path := pathMatch[1]

	// Get the document
	doc, err := s.documentService.GetDocumentByPath(path)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": []map[string]interface{}{
				{"message": fmt.Sprintf("Document not found: %v", err)},
			},
		})
		return
	}

	// Return the document
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": map[string]interface{}{
			"document": doc,
		},
	})
}

// handleSearchQuery processes search queries
func (s *Server) handleSearchQuery(w http.ResponseWriter, query string, params map[string]interface{}) {
	// Extract search parameters
	queryMatch := regexp.MustCompile(`search\(query:\s*"([^"]+)"`).FindStringSubmatch(query)
	if len(queryMatch) < 2 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": []map[string]interface{}{
				{"message": "Invalid search query format"},
			},
		})
		return
	}

	searchQuery := queryMatch[1]

	// Extract limit if provided
	limit := 10 // Default
	limitMatch := regexp.MustCompile(`limit:\s*(\d+)`).FindStringSubmatch(query)
	if len(limitMatch) >= 2 {
		parsedLimit, err := strconv.Atoi(limitMatch[1])
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	// Perform the search
	results, err := s.documentService.SearchDocuments(searchQuery, limit)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": []map[string]interface{}{
				{"message": fmt.Sprintf("Search failed: %v", err)},
			},
		})
		return
	}

	// Return the search results
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": map[string]interface{}{
			"search": map[string]interface{}{
				"documents": results,
				"count":     len(results),
				"query":     searchQuery,
			},
		},
	})
}

// handleFrameworksQuery processes frameworks queries
func (s *Server) handleFrameworksQuery(w http.ResponseWriter, query string, params map[string]interface{}) {
	// Get all frameworks
	frameworks, err := s.documentService.GetFrameworks()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": []map[string]interface{}{
				{"message": fmt.Sprintf("Failed to get frameworks: %v", err)},
			},
		})
		return
	}

	// Return the frameworks
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": map[string]interface{}{
			"frameworks": frameworks,
		},
	})
}

func main() {
	// Parse command line flags
	port := flag.String("port", "8080", "HTTP server port")
	cacheDir := flag.String("cache", ".cache", "Directory containing cached documentation")
	flag.Parse()

	// Create server
	server := NewServer(*cacheDir)

	// Set up routes
	http.HandleFunc("/", server.handleHome)
	http.HandleFunc("/api/document", server.handleDocument)
	http.HandleFunc("/api/search", server.handleSearch)
	http.HandleFunc("/api/frameworks", server.handleFrameworks)
	http.HandleFunc("/graphql", server.handleGraphQL)
	http.HandleFunc("/sandbox", server.handleSandbox)
	http.HandleFunc("/schema", server.handleSchema)

	// Start server
	log.Printf("Apple Docs API server started. Connect to http://localhost:%s/ for documentation", *port)
	log.Printf("GraphQL Playground: http://localhost:%s/graphql", *port)
	log.Printf("Apollo Sandbox: http://localhost:%s/sandbox", *port)
	log.Printf("GraphQL Schema Documentation: http://localhost:%s/schema", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
