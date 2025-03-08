// Command appledocs mirrors Apple documentation JSON files to disk.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	outputDir    = flag.String("output", "output", "directory to store mirrored content")
	cacheDir     = flag.String("cache", ".cache", "directory to store HTTP cache")
	concurrency  = flag.Int("concurrency", 10, "number of concurrent downloads")
	forceRefresh = flag.Bool("force", false, "force refresh all content")
	baseURL      = flag.String("base", "https://developer.apple.com", "base URL for Apple docs")
	timeout      = flag.Duration("timeout", 30*time.Second, "HTTP request timeout")
	maxTime      = flag.Duration("max-time", time.Hour, "maximum time to run the program")
	verbose      = flag.Bool("verbose", false, "enable verbose logging")
	prettyJSON   = flag.Bool("pretty", true, "pretty-print JSON files")
	badURLsFile  = flag.String("bad-urls-file", ".cache/known-bad-urls.txt", "file containing URLs to skip")
)

// JSONFileEntry represents a found JSON file
type JSONFileEntry struct {
	Path string
	URL  string
}

// appledocs holds all the application settings
type appledocs struct {
	client         *http.Client
	visitedURLs    map[string]bool
	visitedMutex   sync.RWMutex
	jsonEntries    []JSONFileEntry
	entriesMutex   sync.Mutex
	processedCount int
	badURLs        map[string]bool // URLs known to be 404s or invalid

	// Status tracking metrics
	cacheHits   int
	cacheMisses int
	errors      int
	skippedURLs int
	statsMutex  sync.Mutex
}

func main() {
	flag.Parse()

	// Create output and cache directories
	for _, dir := range []string{*outputDir, *cacheDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Failed to create directory %q: %v", dir, err)
		}
	}

	// Setup cancellable context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), *maxTime)
	defer cancel()

	// Setup signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("Received interrupt signal, shutting down...")
		cancel()
	}()

	// Run the crawler
	if err := run(ctx); err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Mirroring complete.")
}

// run is the main entry point for the application logic
func run(ctx context.Context) error {
	client := &http.Client{Timeout: *timeout}
	startURL := resolveURL(*baseURL, "/tutorials/data/documentation/technologies.json")

	app := &appledocs{
		client:      client,
		visitedURLs: make(map[string]bool),
		badURLs:     make(map[string]bool),
	}
	
	// Load known bad URLs if file exists
	if err := loadBadURLs(app); err != nil && *verbose {
		log.Printf("Warning: failed to load bad URLs file: %v", err)
	}

	// Mark the start URL as visited
	app.visitedURLs[startURL] = true

	// Setup progress reporting
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				processed, cacheHits, cacheMisses, errors, skipped := app.getStats()
				log.Printf("In progress... Processed: %d files | Cache: %d hits, %d misses | Errors: %d | Skipped: %d",
					processed, cacheHits, cacheMisses, errors, skipped)
			}
		}
	}()

	// Create a buffered channel with sufficient capacity
	// Use a much larger buffer to handle the thousands of URLs discovered in large docs
	urlQueue := make(chan string, *concurrency*5000)

	// Create a context with cancellation for worker control
	workerCtx, cancelWorkers := context.WithCancel(ctx)
	defer cancelWorkers()

	// Create a wait group for the worker pool
	var wg sync.WaitGroup

	// Launch a fixed pool of workers
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			if *verbose {
				log.Printf("Worker %d starting", workerID)
			}

			for {
				select {
				case <-workerCtx.Done():
					if *verbose {
						log.Printf("Worker %d shutting down (context done)", workerID)
					}
					return
				case url, ok := <-urlQueue:
					if !ok {
						if *verbose {
							log.Printf("Worker %d shutting down (channel closed)", workerID)
						}
						return
					}

					// Process the URL directly - no need to spawn another goroutine
					if err := app.processURL(workerCtx, url, urlQueue); err != nil {
						if err != context.Canceled && err != context.DeadlineExceeded {
							log.Printf("Error processing %q: %v", url, err)
						}
					}
				}
			}
		}(i)
	}

	// Add the initial URL to the queue
	urlQueue <- startURL

	// Wait for context cancellation or completion
	<-ctx.Done()

	// Cancel worker context
	cancelWorkers()

	// Signal workers to finish but don't close channel yet
	log.Printf("Waiting for workers to complete...")

	// Use a timeout for waiting so we don't hang forever
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	// Wait for workers with timeout
	select {
	case <-done:
		log.Printf("All workers finished gracefully")
	case <-time.After(5 * time.Second):
		log.Printf("Some workers still running after timeout")
	}

	// Now it's safe to close the channel
	close(urlQueue)

	// Create index.html with the list of JSON files
	if err := createJSONIndexHTML(*outputDir, app.jsonEntries); err != nil {
		log.Printf("Failed to create index.html: %v", err)
	}

	// Report final statistics
	processed, cacheHits, cacheMisses, errors, skipped := app.getStats()
	log.Printf("Final statistics:")
	log.Printf("  - Processed: %d JSON files", processed)
	log.Printf("  - Cache: %d hits, %d misses", cacheHits, cacheMisses)
	log.Printf("  - Errors: %d", errors)
	log.Printf("  - Skipped URLs: %d", skipped)

	return nil
}

// Helper methods for incrementing metrics
func (app *appledocs) incrementCacheHits() {
	app.statsMutex.Lock()
	app.cacheHits++
	app.statsMutex.Unlock()
}

func (app *appledocs) incrementCacheMisses() {
	app.statsMutex.Lock()
	app.cacheMisses++
	app.statsMutex.Unlock()
}

func (app *appledocs) incrementErrors() {
	app.statsMutex.Lock()
	app.errors++
	app.statsMutex.Unlock()
}

func (app *appledocs) incrementSkippedURLs() {
	app.statsMutex.Lock()
	app.skippedURLs++
	app.statsMutex.Unlock()
}

// getStats returns current statistics in a thread-safe way
func (app *appledocs) getStats() (int, int, int, int, int) {
	app.statsMutex.Lock()
	defer app.statsMutex.Unlock()
	app.entriesMutex.Lock()
	processed := len(app.jsonEntries)
	app.entriesMutex.Unlock()
	return processed, app.cacheHits, app.cacheMisses, app.errors, app.skippedURLs
}

// processURL handles a single URL, fetching and processing it
func (app *appledocs) processURL(ctx context.Context, u string, urlQueue chan<- string) error {
	if *verbose {
		log.Printf("Processing %s", u)
	}

	// The concurrency is now managed by the worker pool
	// No need for explicit semaphore acquisition

	// Parse URL to get relative path
	parsed, err := url.Parse(u)
	if err != nil {
		app.incrementErrors()
		return fmt.Errorf("parse URL %q: %v", u, err)
	}

	// Check if this is a valid URL to process
	// Accept both .json files and index URLs
	isJsonFile := strings.HasSuffix(parsed.Path, ".json")
	isIndexFile := strings.Contains(parsed.Path, "/tutorials/data/index/")
	
	if !isJsonFile && !isIndexFile {
		app.incrementSkippedURLs()
		return fmt.Errorf("not a supported URL: %s", u)
	}

	// Get content
	data, err := fetchWithCache(app.client, u, app)
	if err != nil {
		return fmt.Errorf("fetch %q: %v", u, err)
	}

	// Save to output
	outputPath := parsed.Path
	if strings.HasPrefix(outputPath, "/") {
		outputPath = outputPath[1:]
	}
	
	// For index files that don't have a .json extension, add one for consistency
	if isIndexFile && !strings.HasSuffix(outputPath, ".json") {
		outputPath += ".json"
	}

	if err := saveToOutputDir(outputPath, data); err != nil {
		return fmt.Errorf("save %q: %v", outputPath, err)
	}

	if *verbose {
		log.Printf("Saved %s", outputPath)
	}

	// Create entry for the index
	app.entriesMutex.Lock()
	app.jsonEntries = append(app.jsonEntries, JSONFileEntry{
		Path: outputPath,
		URL:  u,
	})
	app.entriesMutex.Unlock()

	// Extract and enqueue new JSON URLs
	newURLs := extractJSONURLs(data)
	if *verbose {
		log.Printf("Found %d URLs in %s", len(newURLs), outputPath)
	}

	// Check if this is a technology file and try to fetch its related URLs
	if isTechnologyFile(parsed.Path) {
		// Get the technology name
		parts := strings.Split(parsed.Path, "/")
		if len(parts) >= 5 {
			technologyFile := parts[4]
			technologyName := strings.TrimSuffix(technologyFile, ".json")
			
			// 1. Add the index URL for the technology
			indexURL := "/tutorials/data/index/" + strings.ToLower(technologyName)
			newURLs = append(newURLs, indexURL)
			
			// 2. Try both capitalization variants for the documentation file
			lowerURL := "/tutorials/data/documentation/" + strings.ToLower(technologyName) + ".json"
				
			upperName := strings.Title(technologyName)
			upperURL := "/tutorials/data/documentation/" + upperName + ".json"
			
			// Add URLs to the queue if they're different from the current one
			if lowerURL != parsed.Path {
				newURLs = append(newURLs, lowerURL)
			}
			if upperURL != parsed.Path {
				newURLs = append(newURLs, upperURL)
			}
			
			if *verbose {
				log.Printf("Added related URLs for technology %s: index and case variants", technologyName)
			}
		}
	}

	// Queue new URLs for processing
	newURLsAdded := app.queueNewURLs(newURLs, urlQueue)

	if *verbose {
		log.Printf("Queued %d new URLs from %s", newURLsAdded, outputPath)
	}

	return nil
}

// queueNewURLs adds new URLs to the processing queue if they haven't been visited
func (app *appledocs) queueNewURLs(newURLs []string, urlQueue chan<- string) int {
	var added int

	for _, newURL := range newURLs {
		resolvedURL := resolveURL(*baseURL, newURL)

		// Check if URL is known to be bad
		if app.badURLs[resolvedURL] {
			if *verbose {
				log.Printf("Skipping known bad URL: %s", resolvedURL)
			}
			continue
		}

		// Check if URL has been visited (using read lock)
		app.visitedMutex.RLock()
		visited := app.visitedURLs[resolvedURL]
		app.visitedMutex.RUnlock()

		if visited {
			continue
		}

		// Mark as visited (using write lock)
		app.visitedMutex.Lock()
		// Double-check after acquiring write lock
		if app.visitedURLs[resolvedURL] {
			app.visitedMutex.Unlock()
			continue
		}
		app.visitedURLs[resolvedURL] = true
		app.visitedMutex.Unlock()

		// Try to send to channel, but don't block or panic if it's closed
		select {
		case urlQueue <- resolvedURL:
			added++
		default:
			// Channel might be full or closed, skip this URL
			if *verbose {
				log.Printf("Skipping URL %s (channel full or closed)", resolvedURL)
			}
		}
	}

	return added
}

// isJSON checks if the data is valid JSON
func isJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}

// prettyPrintJSON formats JSON data with proper indentation
func prettyPrintJSON(data []byte) ([]byte, error) {
	var out bytes.Buffer
	var jsonData interface{}

	// Parse JSON
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	// Format with indentation
	encoder := json.NewEncoder(&out)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(jsonData); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

// fetchWithCache fetches a URL with caching.
func fetchWithCache(client *http.Client, u string, app *appledocs) ([]byte, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		app.incrementErrors()
		return nil, fmt.Errorf("parse URL %q: %v", u, err)
	}

	// Create cache path
	cachePath := filepath.Join(*cacheDir, parsed.Host, parsed.Path)
	if parsed.RawQuery != "" {
		cachePath = filepath.Join(cachePath + "." + url.QueryEscape(parsed.RawQuery))
	}

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(cachePath), 0755); err != nil {
		app.incrementErrors()
		return nil, fmt.Errorf("create cache directory: %v", err)
	}

	// Check cache
	if !*forceRefresh {
		// check exists, and is not empty:
		if fi, err := os.Stat(cachePath); err == nil && fi.Size() > 0 {
			// Read from cache
			data, err := os.ReadFile(cachePath)
			// Check for empty/malformed cache files
			if err == nil {
				// For JSON files, verify that the content is valid JSON
				if strings.HasSuffix(parsed.Path, ".json") && !isJSON(data) {
					log.Printf("Cache contains invalid JSON for %q, refetching", u)
					app.incrementCacheMisses()
				} else {
					app.incrementCacheHits()
					if *verbose {
						log.Printf("Cache hit for %q", u)
					}
					return data, nil
				}
			} else {
				log.Printf("Cache read error for %q: %v", u, err)
				app.incrementErrors()
				app.incrementCacheMisses()
			}
		} else {
			app.incrementCacheMisses()
		}
	} else {
		app.incrementCacheMisses()
	}

	// Fetch URL
	resp, err := client.Get(u)
	if err != nil {
		app.incrementErrors()
		// Add to bad URLs list if it's a network error
		if strings.Contains(err.Error(), "no such host") || 
		   strings.Contains(err.Error(), "connection refused") ||
		   strings.Contains(err.Error(), "timeout") {
			app.badURLs[u] = true
			// Save to known-bad-urls file
			appendToBadURLsFile(u)
		}
		return nil, fmt.Errorf("fetch %q: %v", u, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		app.incrementErrors()
		// Add to bad URLs list if it's a 404
		if resp.StatusCode == http.StatusNotFound {
			app.badURLs[u] = true
			// Save to known-bad-urls file
			appendToBadURLsFile(u)
		}
		return nil, fmt.Errorf("fetch %q: %s", u, resp.Status)
	}

	// Read response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		app.incrementErrors()
		return nil, fmt.Errorf("read %q: %v", u, err)
	}

	// Write to cache
	if err := os.WriteFile(cachePath, data, 0644); err != nil {
		app.incrementErrors()
		return nil, fmt.Errorf("write cache %q: %v", cachePath, err)
	}

	return data, nil
}

// saveToOutputDir saves content to the output directory.
func saveToOutputDir(path string, content []byte) error {
	outputPath := filepath.Join(*outputDir, path)
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("create output directory: %v", err)
	}

	// Pretty-print JSON files if enabled
	if *prettyJSON && strings.HasSuffix(path, ".json") {
		// Try to pretty-print the JSON, but if it fails, use the original content
		if prettyContent, err := prettyPrintJSON(content); err == nil {
			content = prettyContent
		} else {
			log.Printf("Failed to pretty-print JSON file %s: %v", path, err)
		}
	}

	return os.WriteFile(outputPath, content, 0644)
}

// resolveURL resolves a potentially relative URL against the base URL.
func resolveURL(base, relative string) string {
	if strings.HasPrefix(relative, "http://") || strings.HasPrefix(relative, "https://") {
		return relative
	}

	// Handle documentation paths
	if strings.HasPrefix(relative, "documentation/") {
		return *baseURL + "/tutorials/data/" + relative
	}

	if strings.HasPrefix(relative, "/") {
		return *baseURL + relative
	}
	return *baseURL + "/" + relative
}

// extractJSONURLs extracts URLs to other JSON files from a JSON response.
func extractJSONURLs(data []byte) []string {
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil
	}

	urls := []string{}
	extractJSONURLsFromValue(result, &urls)
	return urls
}

// extractJSONURLsFromValue recursively extracts all URLs from a JSON value.
func extractJSONURLsFromValue(v interface{}, urls *[]string) {
	switch val := v.(type) {
	case map[string]interface{}:
		// Special handling for destination objects with identifiers
		if _, hasType := val["type"]; hasType {
			if identifier, ok := val["identifier"].(string); ok {
				// Handle doc:// URLs by converting them to HTTP URLs
				if strings.HasPrefix(identifier, "doc://") {
					docPath := strings.TrimPrefix(identifier, "doc://")
					parts := strings.SplitN(docPath, "/", 2)
					if len(parts) > 1 {
						// Convert doc:// URL to a JSON path
						jsonURL := parts[1] + ".json"
						*urls = append(*urls, jsonURL)
					}
				} else if !strings.HasPrefix(identifier, "http") && !strings.HasPrefix(identifier, "https") {
					// Handle relative URLs
					*urls = append(*urls, identifier)
				}
			}
		}

		for k, v := range val {
			// Check if this is a URL field that points to a JSON file
			if (k == "url" || strings.HasSuffix(k, "URL") || strings.HasSuffix(k, "Uri")) &&
				v != nil {
				if urlStr, ok := v.(string); ok && strings.HasSuffix(urlStr, ".json") {
					*urls = append(*urls, urlStr)
				}
			}
			extractJSONURLsFromValue(v, urls)
		}
	case []interface{}:
		for _, item := range val {
			extractJSONURLsFromValue(item, urls)
		}
	}
}

// TreeNode represents a node in the file tree
type TreeNode struct {
	Name     string
	Path     string
	IsDir    bool
	Children []*TreeNode
}

// createJSONIndexHTML creates an index.html file in the output directory
// that shows a tree view of all the JSON files that were mirrored.
func createJSONIndexHTML(outputDir string, jsonFiles []JSONFileEntry) error {
	// Scan the output directory to make sure we include all JSON files
	scannedFiles, err := scanOutputDirectory(outputDir)
	if err != nil {
		log.Printf("Warning: Error scanning output directory: %v", err)
		// Continue with the files we know about
	} else if len(scannedFiles) > len(jsonFiles) {
		// Use the scanned files if we found more
		log.Printf("Found %d additional files in output directory", len(scannedFiles)-len(jsonFiles))
		jsonFiles = scannedFiles
	}
	
	// Build tree from files
	root := buildFileTree(jsonFiles)
	
	// Calculate stats
	dirCount := countDirectories(root)

	data := struct {
		Root      *TreeNode
		FileCount int
		DirCount  int
		Timestamp string
	}{
		Root:      root,
		FileCount: len(jsonFiles),
		DirCount:  dirCount,
		Timestamp: time.Now().Format(time.RFC1123),
	}

	// Create the HTML file with embedded JavaScript
	indexPath := filepath.Join(outputDir, "index.html")
	if err := generateHTMLFile(indexPath, data); err != nil {
		return err
	}

	log.Printf("Created tree-view index.html with %d files across %d directories at %s", 
		len(jsonFiles), dirCount, outputDir)
	return nil
}

// scanOutputDirectory walks the output directory and finds all JSON files
func scanOutputDirectory(outputDir string) ([]JSONFileEntry, error) {
	var files []JSONFileEntry

	err := filepath.Walk(outputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip directories
		if info.IsDir() {
			return nil
		}
		
		// Only include JSON files
		if strings.HasSuffix(strings.ToLower(path), ".json") {
			// Convert absolute path to relative path from output directory
			relPath, err := filepath.Rel(outputDir, path)
			if err != nil {
				return err
			}
			
			files = append(files, JSONFileEntry{
				Path: relPath,
				URL:  relPath, // URL could be reconstructed if needed
			})
		}
		
		return nil
	})

	return files, err
}

// buildFileTree constructs a tree structure from the flat list of file paths
func buildFileTree(files []JSONFileEntry) *TreeNode {
	root := &TreeNode{
		Name:     "root",
		Path:     "",
		IsDir:    true,
		Children: []*TreeNode{},
	}

	for _, file := range files {
		addFileToTree(root, file.Path, file.Path)
	}

	// Sort the tree nodes
	sortTree(root)

	return root
}

// addFileToTree recursively adds a file to the tree
func addFileToTree(node *TreeNode, path string, fullPath string) {
	parts := strings.Split(path, "/")

	if len(parts) == 1 {
		// This is a file at this level
		node.Children = append(node.Children, &TreeNode{
			Name:  parts[0],
			Path:  fullPath,
			IsDir: false,
		})
		return
	}

	// This is a directory
	dirName := parts[0]
	restPath := strings.Join(parts[1:], "/")

	// Look for existing directory node
	var dirNode *TreeNode
	for _, child := range node.Children {
		if child.IsDir && child.Name == dirName {
			dirNode = child
			break
		}
	}

	// Create new directory node if it doesn't exist
	if dirNode == nil {
		dirNode = &TreeNode{
			Name:     dirName,
			Path:     "",
			IsDir:    true,
			Children: []*TreeNode{},
		}
		node.Children = append(node.Children, dirNode)
	}

	// Add the rest of the path to the directory node
	addFileToTree(dirNode, restPath, fullPath)
}

// sortTree sorts the tree by name, directories first
func sortTree(node *TreeNode) {
	// Sort children
	sort.Slice(node.Children, func(i, j int) bool {
		// Directories come before files
		if node.Children[i].IsDir && !node.Children[j].IsDir {
			return true
		}
		if !node.Children[i].IsDir && node.Children[j].IsDir {
			return false
		}
		// Then sort by name
		return node.Children[i].Name < node.Children[j].Name
	})

	// Recursively sort children
	for _, child := range node.Children {
		if child.IsDir {
			sortTree(child)
		}
	}
}

// countDirectories returns the total number of directories in the tree
func countDirectories(node *TreeNode) int {
	count := 0
	if node.IsDir && node.Name != "root" {
		count = 1
	}

	for _, child := range node.Children {
		if child.IsDir {
			count += countDirectories(child)
		}
	}

	return count
}

// isTechnologyFile checks if the given path represents a technology file
func isTechnologyFile(path string) bool {
	// Check if the file is in the documentation directory and not in a subdirectory
	// For example: /tutorials/data/documentation/EndpointSecurity.json
	parts := strings.Split(path, "/")
	if len(parts) < 4 {
		return false
	}
	
	// Check if this is a top-level technology file
	// The pattern should be /tutorials/data/documentation/TechnologyName.json
	if parts[1] == "tutorials" && parts[2] == "data" && parts[3] == "documentation" {
		// If there are no further subdirectories and it ends with .json
		return len(parts) == 5 && strings.HasSuffix(parts[4], ".json")
	}
	
	return false
}

// createIndexURLForTechnology creates an index URL for a given technology path
func createIndexURLForTechnology(path string) string {
	// Extract the technology name from the path
	// Example: /tutorials/data/documentation/EndpointSecurity.json -> EndpointSecurity
	parts := strings.Split(path, "/")
	if len(parts) < 5 {
		return ""
	}
	
	technologyFile := parts[4]
	technologyName := strings.TrimSuffix(technologyFile, ".json")
	
	// Skip common non-technology files
	if technologyName == "technologies" {
		return ""
	}
	
	// Create the index URL: /tutorials/data/index/technologyname
	return "/tutorials/data/index/" + strings.ToLower(technologyName)
}

// loadBadURLs loads the list of known bad URLs from a file
func loadBadURLs(app *appledocs) error {
	// Check if the file exists
	if _, err := os.Stat(*badURLsFile); os.IsNotExist(err) {
		// File doesn't exist, which is fine
		return nil
	} else if err != nil {
		return err
	}

	// Read the file
	content, err := os.ReadFile(*badURLsFile)
	if err != nil {
		return err
	}

	// Parse each line as a URL
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			// Skip empty lines and comments
			continue
		}
		app.badURLs[line] = true
		if *verbose {
			log.Printf("Added bad URL: %s", line)
		}
	}

	if *verbose {
		log.Printf("Loaded %d known bad URLs", len(app.badURLs))
	}
	return nil
}

// appendToBadURLsFile adds a URL to the bad URLs file
func appendToBadURLsFile(url string) {
	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(*badURLsFile), 0755); err != nil {
		log.Printf("Error creating bad URLs directory: %v", err)
		return
	}

	// Create file if it doesn't exist
	file, err := os.OpenFile(*badURLsFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening bad URLs file: %v", err)
		return
	}
	defer file.Close()

	// Write URL to file
	if _, err := file.WriteString(url + "\n"); err != nil {
		log.Printf("Error writing to bad URLs file: %v", err)
	}
}
