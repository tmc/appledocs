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
	"regexp"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	// Directories and URLs
	outputDir    = flag.String("output", "output", "directory to store mirrored content")
	cacheDir     = flag.String("cache", ".cache", "directory to store HTTP cache")
	baseURL      = flag.String("base", "https://developer.apple.com", "base URL for Apple docs")
	entryPoint   = flag.String("entry-point", "/tutorials/data/documentation/technologies.json", "path to start crawling from")
	badURLsFile  = flag.String("bad-urls-file", ".cache/known-bad-urls.txt", "file containing URLs to skip")
	excludePaths = flag.String("exclude-paths", "en-US/docs/Mozilla", "comma-separated list of paths to exclude from crawling")

	// Crawling options
	concurrency  = flag.Int("concurrency", 1, "number of concurrent downloads")
	delay        = flag.Duration("delay", 0, "delay between urls")
	forceRefresh = flag.Bool("force", false, "force refresh all content")
	timeout      = flag.Duration("timeout", 30*time.Second, "HTTP request timeout")
	maxTime      = flag.Duration("max-time", time.Hour, "maximum time to run the program")
	skipSymbols  = flag.Bool("skip-symbols", false, "skip individual symbol level documentation")
	printURLs    = flag.Bool("print-urls", false, "only print discovered URLs from entry point and exit")

	// Output options
	prettyJSON = flag.Bool("pretty", true, "pretty-print JSON files")
	verbose    = flag.Bool("verbose", false, "enable verbose logging")

	// Mode selection
	mode = flag.String("mode", "crawl", "operation mode: crawl, html, markdown, or all")

	// Markdown-specific options
	mdOutputDir = flag.String("md-output", "markdown", "directory to store Markdown documentation")

	// Legacy flags for backward compatibility
	generateMD = flag.Bool("markdown", false, "generate Markdown documentation from the JSON files")
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
	urlDepths      map[string]int  // Track semantic depth of each URL
	depthMutex     sync.RWMutex    // Mutex for urlDepths

	// Status tracking metrics
	cacheHits      int
	cacheMisses    int
	errors         int
	skippedURLs    int
	skippedSymbols int // Count of URLs skipped because they're symbol-level docs
	statsMutex     sync.Mutex
}

// buildFrameworkURLs constructs URLs for a specific framework
func buildFrameworkURLs(frameworkName string) []string {
	// Normalize framework name
	frameworkName = strings.TrimSuffix(frameworkName, ".json")

	// Create lowercase version for index URLs
	lowerFramework := strings.ToLower(frameworkName)

	// Use both documentation and index URLs
	urls := []string{
		// Main documentation URL
		fmt.Sprintf("tutorials/data/documentation/%s.json", frameworkName),
		// Index URL (uses lowercase without .json extension)
		fmt.Sprintf("tutorials/data/index/%s", lowerFramework),
	}

	return urls
}

// fetchAndExtractURLs fetches a URL and extracts URLs from its content
func fetchAndExtractURLs(client *http.Client, app *appledocs, fetchURL string) ([]string, error) {
	log.Printf("Fetching URL: %s", fetchURL)
	data, err := fetchWithCache(client, fetchURL, app)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch %s: %v", fetchURL, err)
	}

	// Extract URLs
	urls := extractJSONURLs(data)
	log.Printf("Found %d URLs in %s", len(urls), fetchURL)

	// Return the extracted URLs
	return urls, nil
}

// isDataURL checks if the URL is a tutorial data URL
func isDataURL(url string) bool {
	// Only include URLs with tutorials/data/ path component which indicates Apple data files
	return strings.Contains(url, "/tutorials/data/")
}

// printURLsOnly fetches the entry point URL and prints all discovered URLs
func printURLsOnly() error {
	client := &http.Client{Timeout: *timeout}

	// Create simple app instance for cache tracking
	app := &appledocs{
		client:      client,
		visitedURLs: make(map[string]bool),
		badURLs:     make(map[string]bool),
	}

	// Ensure we have a cache directory
	if err := os.MkdirAll(*cacheDir, 0755); err != nil {
		return fmt.Errorf("failed to create cache directory: %v", err)
	}

	// Load bad URLs to avoid known problem URLs
	if err := loadBadURLs(app); err != nil && *verbose {
		log.Printf("Warning: failed to load bad URLs file: %v", err)
	}

	var allURLs []string

	// Check if a specific entry point was specified, not the default technologies.json
	if *entryPoint != "/tutorials/data/documentation/technologies.json" {
		// Remove any leading slash and .json suffix
		cleanEntry := strings.TrimPrefix(*entryPoint, "/")
		cleanEntry = strings.TrimSuffix(cleanEntry, ".json")

		// Split the entry into components
		components := strings.Split(cleanEntry, "/")

		// Determine if we have a framework or framework/class pattern
		var frameworkURLs []string

		if len(components) >= 2 {
			// We have a class-specific path
			framework := components[0]
			class := components[1]
			log.Printf("Fetching URLs for class: %s in framework: %s", class, framework)

			// For a class, we just use the direct class URL
			frameworkURLs = []string{
				fmt.Sprintf("tutorials/data/documentation/%s/%s.json", framework, class),
			}
		} else if len(components) == 1 {
			// We just have a framework
			framework := components[0]
			log.Printf("Fetching URLs for framework: %s", framework)

			// For a framework, use both doc and index URLs
			frameworkURLs = buildFrameworkURLs(framework)
		}

		// Use a map to track unique URLs
		uniqueURLs := make(map[string]bool)

		// Try each URL
		for _, urlPath := range frameworkURLs {
			fullURL := resolveURL(*baseURL, urlPath)
			extractedURLs, err := fetchAndExtractURLs(client, app, fullURL)
			if err != nil {
				// Log but don't fail - some paths might not exist
				log.Printf("Warning: %v", err)
				continue
			}

			// Add extracted URLs to our collection, ensuring uniqueness
			for _, u := range extractedURLs {
				resolvedURL := resolveURL(*baseURL, u)
				// Only include tutorial data URLs
				if isDataURL(resolvedURL) && !uniqueURLs[resolvedURL] {
					uniqueURLs[resolvedURL] = true
					allURLs = append(allURLs, resolvedURL)
				}
			}

			// Also add the source URL itself if it's not already included
			if !uniqueURLs[fullURL] {
				uniqueURLs[fullURL] = true
				allURLs = append(allURLs, fullURL)
			}
		}
	} else {
		// Default to technologies.json (removed leading slash)
		startURL := resolveURL(*baseURL, "tutorials/data/documentation/technologies.json")

		// Fetch the technologies index
		log.Printf("Fetching technologies index: %s", startURL)
		extractedURLs, err := fetchAndExtractURLs(client, app, startURL)
		if err != nil {
			return fmt.Errorf("failed to fetch technologies index: %v", err)
		}

		// Use a map to track unique URLs
		uniqueURLs := make(map[string]bool)

		// Process the extracted URLs
		for _, u := range extractedURLs {
			resolvedURL := resolveURL(*baseURL, u)
			if !uniqueURLs[resolvedURL] {
				uniqueURLs[resolvedURL] = true
				allURLs = append(allURLs, resolvedURL)
			}
		}
	}

	// Filter to only include tutorial data URLs
	var filteredURLs []string

	for _, url := range allURLs {
		// Only include URLs with the proper pattern
		if isDataURL(url) {
			filteredURLs = append(filteredURLs, url)
		}
	}

	log.Printf("Found a total of %d unique data URLs", len(filteredURLs))

	// Ensure log messages are flushed before printing URLs
	time.Sleep(100 * time.Millisecond)

	// Sort URLs for consistent output
	sort.Strings(filteredURLs)

	// Print URLs to stdout
	for _, resolvedURL := range filteredURLs {
		fmt.Println(resolvedURL)
	}

	return nil
}

func main() {
	flag.Parse()

	// Handle legacy flag conversion for backward compatibility
	if *generateMD && *mode == "crawl" {
		*mode = "markdown"
	}

	// Special case for print-urls mode
	if *printURLs {
		if err := printURLsOnly(); err != nil {
			log.Fatalf("Error: %v", err)
		}
		return
	}

	// Validate mode
	validModes := map[string]bool{"crawl": true, "html": true, "markdown": true, "all": true}
	if !validModes[*mode] {
		log.Fatalf("Invalid mode: %s. Must be one of: crawl, html, markdown, or all", *mode)
	}

	// Create necessary directories
	dirsToCreate := []string{}
	if *mode == "crawl" || *mode == "all" {
		dirsToCreate = append(dirsToCreate, *outputDir, *cacheDir)
	}
	if *mode == "markdown" || *mode == "all" {
		dirsToCreate = append(dirsToCreate, *mdOutputDir)
	}

	for _, dir := range dirsToCreate {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Failed to create directory %q: %v", dir, err)
		}
	}

	// Crawl mode - fetch JSON files
	if *mode == "crawl" || *mode == "all" {
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
			log.Fatalf("Error during crawling: %v", err)
		}

		log.Printf("Mirroring complete.")
	}

	// HTML mode - generate HTML index
	if *mode == "html" || *mode == "all" {
		log.Printf("Generating HTML index...")
		jsonFiles, err := scanOutputDirectory(*outputDir)
		if err != nil {
			log.Fatalf("Error scanning output directory: %v", err)
		}

		if err := createJSONIndexHTML(*outputDir, jsonFiles); err != nil {
			log.Fatalf("Error generating HTML: %v", err)
		}

		log.Printf("HTML generation complete. Open %s/index.html to view.", *outputDir)
	}

	// Markdown mode - generate Markdown files
	if *mode == "markdown" || *mode == "all" {
		log.Printf("Generating Markdown documentation...")
		if err := generateMarkdown(*outputDir, *mdOutputDir); err != nil {
			log.Fatalf("Error generating Markdown: %v", err)
		}
		log.Printf("Markdown generation complete. Output in: %s", *mdOutputDir)
	}
}

// run is the main entry point for the application logic
func run(ctx context.Context) error {
	client := &http.Client{Timeout: *timeout}
	startURL := resolveURL(*baseURL, *entryPoint)

	app := &appledocs{
		client:      client,
		visitedURLs: make(map[string]bool),
		badURLs:     make(map[string]bool),
		urlDepths:   make(map[string]int),
	}

	// Create bad URLs directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(*badURLsFile), 0755); err != nil {
		log.Printf("Warning: failed to create directory for bad URLs file: %v", err)
	}

	// Load known bad URLs if file exists
	if err := loadBadURLs(app); err != nil && *verbose {
		log.Printf("Warning: failed to load bad URLs file: %v", err)
	} else if *verbose {
		log.Printf("Loaded %d known bad URLs", len(app.badURLs))
	}

	// Mark the start URL as visited and set its depth to 0
	app.visitedURLs[startURL] = true
	app.depthMutex.Lock()
	app.urlDepths[startURL] = 0
	app.depthMutex.Unlock()

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

	// Make sure badURLs file is written
	badURLCount := len(app.badURLs)
	if badURLCount > 0 {
		log.Printf("Writing %d bad URLs to %s", badURLCount, *badURLsFile)
		if err := writeBadURLsFile(app); err != nil {
			log.Printf("Warning: Failed to write bad URLs file: %v", err)
		}
	}

	// Report final statistics
	processed, cacheHits, cacheMisses, errors, skipped := app.getStats()
	log.Printf("Final statistics:")
	log.Printf("  - Processed: %d JSON files", processed)
	log.Printf("  - Cache: %d hits, %d misses", cacheHits, cacheMisses)
	log.Printf("  - Errors: %d", errors)
	log.Printf("  - Skipped URLs: %d", skipped)
	if *skipSymbols {
		log.Printf("  - Skipped symbol URLs: %d", app.skippedSymbols)
	}

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

func (app *appledocs) incrementSkippedSymbols() {
	app.statsMutex.Lock()
	app.skippedSymbols++
	app.skippedURLs++ // Also count in the general skipped URLs
	app.statsMutex.Unlock()
}

// getStats returns current statistics in a thread-safe way
func (app *appledocs) getStats() (int, int, int, int, int) {
	app.statsMutex.Lock()
	defer app.statsMutex.Unlock()
	app.entriesMutex.Lock()
	processed := len(app.jsonEntries)
	app.entriesMutex.Unlock()

	// Note: depthLimits is counted as part of skippedURLs for backward compatibility
	// with the existing reporting, so we don't need to return it separately
	return processed, app.cacheHits, app.cacheMisses, app.errors, app.skippedURLs
}

// shouldExcludePath checks if a URL path should be excluded based on user-defined exclude patterns
func shouldExcludePath(pathToCheck string) bool {
	if *excludePaths == "" {
		return false
	}

	patterns := strings.Split(*excludePaths, ",")
	for _, pattern := range patterns {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" {
			continue
		}

		// Check if the path contains this pattern
		if strings.Contains(pathToCheck, pattern) {
			if *verbose {
				log.Printf("Excluding path %q because it matches pattern %q", pathToCheck, pattern)
			}
			return true
		}
	}

	return false
}

// processURL handles a single URL, fetching and processing it
func (app *appledocs) processURL(ctx context.Context, u string, urlQueue chan<- string) error {
	if *verbose {
		log.Printf("Processing %s", u)
		if *delay > time.Duration(0) {
			log.Println("waiting", *delay)
			time.Sleep(*delay)
		}
	}

	// First, check if this URL is already known to be bad
	if app.badURLs[u] {
		if *verbose {
			log.Printf("Skipping known bad URL: %s", u)
		}
		app.incrementSkippedURLs()
		return fmt.Errorf("known bad URL: %s", u)
	}

	// The concurrency is now managed by the worker pool
	// No need for explicit semaphore acquisition

	// Parse URL to get relative path
	parsed, err := url.Parse(u)
	if err != nil {
		app.incrementErrors()
		return fmt.Errorf("parse URL %q: %v", u, err)
	}

	// Check if this URL matches any exclude patterns
	if shouldExcludePath(parsed.Path) {
		// Add to bad URLs to prevent future attempts to process this URL
		app.badURLs[u] = true
		appendToBadURLsFile(u)
		app.incrementSkippedURLs()
		if *verbose {
			log.Printf("Skipping excluded path, added to bad URLs: %s", parsed.Path)
		}
		return fmt.Errorf("excluded path: %s", parsed.Path)
	}

	// Check if this is a valid URL to process
	isMediaFile := strings.Contains(parsed.Path, "/media-")

	// Check for UUID-like paths or other patterns that aren't valid tutorial URLs
	isUUIDLike := false
	// Only mark URLs as marketing pages if they're not part of the tutorials path
	isMarketingPage := false
	isMarketingPage = strings.HasSuffix(parsed.Path, "-hero") ||
		strings.HasPrefix(parsed.Path, "/devLink-") ||
		strings.HasPrefix(parsed.Path, "/link-") ||
		strings.Contains(parsed.Path, "-module") ||
		strings.Contains(parsed.Path, "-dynamic-") ||
		strings.Contains(parsed.Path, "#") ||
		strings.Contains(parsed.Path, "managedapp")
	// For specific URLs that we know will always fail or aren't relevant, add them to bad URLs list
	if isMediaFile ||
		isUUIDLike ||
		isMarketingPage ||
		strings.Contains(parsed.Path, "/assets/") ||
		strings.HasSuffix(parsed.Path, ".css") ||
		strings.HasSuffix(parsed.Path, ".js") ||
		strings.HasSuffix(parsed.Path, ".png") ||
		strings.HasSuffix(parsed.Path, ".jpg") ||
		strings.HasSuffix(parsed.Path, ".svg") ||
		strings.HasSuffix(parsed.Path, ".pdf") {
		app.badURLs[u] = true
		appendToBadURLsFile(u)
		app.incrementSkippedURLs()
		if *verbose {
			log.Printf("Added unsupported URL to bad URLs: %s", u)
		}
		return fmt.Errorf("not a supported URL type: %s", u)
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

	// If skipSymbols is enabled, filter out symbol-level documentation URLs
	if *skipSymbols {
		filteredURLs := make([]string, 0, len(newURLs))
		for _, url := range newURLs {
			if !isSymbolURL(url) {
				filteredURLs = append(filteredURLs, url)
			} else {
				app.incrementSkippedSymbols()
				if *verbose {
					log.Printf("Skipping symbol URL: %s", url)
				}
			}
		}
		if *verbose && len(newURLs) != len(filteredURLs) {
			log.Printf("Filtered out %d symbol URLs", len(newURLs)-len(filteredURLs))
		}
		newURLs = filteredURLs
	}

	if *verbose {
		log.Printf("Found %d URLs in %s", len(newURLs), outputPath)
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

		// Parse URL to check if it should be excluded
		parsedURL, err := url.Parse(resolvedURL)
		if err == nil && shouldExcludePath(parsedURL.Path) {
			// Add to bad URLs so we don't attempt it again
			app.badURLs[resolvedURL] = true
			appendToBadURLsFile(resolvedURL)
			if *verbose {
				log.Printf("Skipping excluded path when queueing: %s", parsedURL.Path)
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

// addBrowserLikeHeaders adds headers to a request to make it look like a browser request
func addBrowserLikeHeaders(req *http.Request) {
	// Standard browser-like headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("sec-ch-ua", "\"Not(A:Brand\";v=\"99\", \"Google Chrome\";v=\"133\", \"Chromium\";v=\"133\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"macOS\"")

	// Set referer based on URL
	if strings.Contains(req.URL.Path, ".json") {
		docPath := strings.TrimSuffix(req.URL.Path, ".json")
		req.Header.Set("Referer", fmt.Sprintf("%s://%s/documentation%s", req.URL.Scheme, req.URL.Host, docPath))
	} else {
		req.Header.Set("Referer", fmt.Sprintf("%s://%s/", req.URL.Scheme, req.URL.Host))
	}
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

	// Create a new request so we can add headers
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		app.incrementErrors()
		return nil, fmt.Errorf("create request for %q: %v", u, err)
	}

	// Add browser-like headers
	addBrowserLikeHeaders(req)

	// Execute the request
	resp, err := client.Do(req)
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
		// Add to bad URLs list if it's any client error that won't be resolved by retrying
		// These include 403 (Forbidden), 404 (Not Found), 405 (Method Not Allowed), 410 (Gone)
		if resp.StatusCode == http.StatusForbidden ||
			resp.StatusCode == http.StatusNotFound ||
			resp.StatusCode == http.StatusMethodNotAllowed ||
			resp.StatusCode == http.StatusGone {
			app.badURLs[u] = true
			// Save to known-bad-urls file
			appendToBadURLsFile(u)
			if *verbose {
				log.Printf("Added bad URL due to %d status: %s", resp.StatusCode, u)
			}
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

			// Check if this is a path field (used by index files)
			if k == "path" && v != nil {
				if pathStr, ok := v.(string); ok {
					// Only add paths that look like documentation paths
					if strings.HasPrefix(pathStr, "/documentation/") {
						// For any documentation path with reasonable depth, add it as-is
						// We only care about paths deep enough to be meaningful
						if strings.Count(pathStr, "/") > 2 {
							*urls = append(*urls, pathStr)
						}
					}
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
	// Always scan the tutorials directory to get files directly from disk
	tutorialsDir := filepath.Join(outputDir, "tutorials")
	var filesFromDisk []JSONFileEntry

	if _, err := os.Stat(tutorialsDir); err == nil {
		log.Printf("Scanning the tutorials directory for JSON files...")

		// Only scan one directory level deep to keep it manageable
		_, err := os.ReadDir(tutorialsDir)
		if err != nil {
			log.Printf("Warning: Error reading tutorials directory: %v", err)
		} else {
			// Process only data directory which contains the main content
			dataDir := filepath.Join(tutorialsDir, "data")
			if _, err := os.Stat(dataDir); err == nil {
				// Process only first 1000 files to avoid timeout
				const maxFiles = 1000
				filesAdded := 0

				err := filepath.Walk(dataDir, func(path string, info os.FileInfo, err error) error {
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
							return nil
						}

						filesFromDisk = append(filesFromDisk, JSONFileEntry{
							Path: relPath,
							URL:  relPath, // URL could be reconstructed if needed
						})

						filesAdded++
						if filesAdded >= maxFiles {
							log.Printf("Reached limit of %d files, stopping scan", maxFiles)
							return filepath.SkipDir
						}
					}

					return nil
				})

				if err != nil {
					log.Printf("Warning: Error walking data directory: %v", err)
				}
			}
		}

		log.Printf("Found %d JSON files on disk", len(filesFromDisk))
		jsonFiles = filesFromDisk
	} else {
		log.Printf("Tutorials directory not found, using crawled files")
	}

	// Build tree from files
	// root := buildFileTree(jsonFiles)

	// // Calculate stats
	// dirCount := countDirectories(root)

	// data := struct {
	// 	Root      *TreeNode
	// 	FileCount int
	// 	DirCount  int
	// 	Timestamp string
	// }{
	// 	Root:      root,
	// 	FileCount: len(jsonFiles),
	// 	DirCount:  dirCount,
	// 	Timestamp: time.Now().Format(time.RFC1123),
	// }

	// // Generate HTML using the function from html.go
	// indexPath := filepath.Join(outputDir, "index.html")
	// return generateHTMLFile(indexPath, data)
	return nil
}

// scanOutputDirectory walks the output directory and finds all JSON files
func scanOutputDirectory(scanDir string) ([]JSONFileEntry, error) {
	var files []JSONFileEntry
	baseDir := *outputDir

	// Track the total size of all JSON files
	var totalSize int64

	// Count files by directory to potentially skip directories with too many files
	dirFileCounts := make(map[string]int)

	// First pass: count files by directory
	err := filepath.Walk(scanDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories in count
		if info.IsDir() {
			return nil
		}

		// Only count JSON files
		if strings.HasSuffix(strings.ToLower(path), ".json") {
			dir := filepath.Dir(path)
			dirFileCounts[dir]++
			totalSize += info.Size()
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Set a reasonable file limit per directory to avoid huge HTML files
	const maxFilesPerDir = 200

	// Second pass: collect files, limiting per directory
	err = filepath.Walk(scanDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Only include JSON files
		if strings.HasSuffix(strings.ToLower(path), ".json") {
			dir := filepath.Dir(path)

			// Skip if this directory has too many files
			if dirFileCounts[dir] > maxFilesPerDir {
				// Include only a marker file for directories with too many files
				if len(files) == 0 || !strings.Contains(files[len(files)-1].Path, dir) {
					// Calculate relative directory path
					dirRelPath, err := filepath.Rel(baseDir, dir)
					if err != nil {
						return err
					}

					files = append(files, JSONFileEntry{
						Path: filepath.Join(dirRelPath, "_TOO_MANY_FILES_.json"),
						URL:  "too_many_files", // Special marker
					})
				}
				return nil
			}

			// Create a relative path from output directory
			relPath, err := filepath.Rel(baseDir, path)
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

	log.Printf("Scanned %d MB of JSON files, limited to %d files in HTML tree", totalSize/(1024*1024), len(files))
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

// isSymbolURL determines if a URL likely points to individual symbol documentation
// This helps filter out the deepest level API documentation when using -skip-symbols
func isSymbolURL(urlPath string) bool {
	// Check for patterns that indicate symbol-level docs:
	// URLs with type paths like: /documentation/uikit/uiview/1622418-alpha
	if match, _ := regexp.MatchString(`/documentation/[^/]+/[^/]+/\d+\-`, urlPath); match {
		return true
	}

	// URLs with method or property paths
	if match, _ := regexp.MatchString(`/documentation/[^/]+/[^/]+/[^/]+/[^/]+$`, urlPath); match {
		return true
	}

	// Symbol reference paths with 3+ segments after the framework
	parts := strings.Split(urlPath, "/")
	if strings.Contains(urlPath, "/documentation/") && len(parts) >= 6 {
		return true
	}

	return false
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

// writeBadURLsFile writes all bad URLs to the bad URLs file
func writeBadURLsFile(app *appledocs) error {
	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(*badURLsFile), 0755); err != nil {
		return err
	}

	// Get a sorted list of bad URLs for more consistent file content
	urls := make([]string, 0, len(app.badURLs))
	for url := range app.badURLs {
		urls = append(urls, url)
	}
	sort.Strings(urls)

	// Create the file content with a header
	var content strings.Builder
	content.WriteString("# Known bad URLs for appledocs\n")
	content.WriteString("# Last updated: " + time.Now().Format(time.RFC3339) + "\n")
	content.WriteString("# Do not edit this file manually\n\n")

	for _, url := range urls {
		content.WriteString(url + "\n")
	}

	// Write to file
	return os.WriteFile(*badURLsFile, []byte(content.String()), 0644)
}
