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
	"log/slog"
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

	"golang.org/x/time/rate"
)

// Pool for reusing slice allocations
var urlSlicePool = sync.Pool{
	New: func() interface{} {
		return make([]string, 0, 100) // Pre-allocate capacity
	},
}

// Global structured logger
var logger *slog.Logger

// Build-time variables (set via -ldflags)
var (
	version   = "dev"
	commit    = "none"
	buildTime = "unknown"
)

var (
	// Version flag
	showVersion = flag.Bool("version", false, "show version information and exit")

	// Directories and URLs
	outputDir    = flag.String("output", "output", "directory to store mirrored content")
	cacheDir     = flag.String("cache", filepath.Join(os.Getenv("HOME"), ".appledocs", "cache"), "directory to store HTTP cache")
	baseURL      = flag.String("base", "https://developer.apple.com", "base URL for Apple docs")
	entryPoint   = flag.String("entry-point", "/tutorials/data/documentation/technologies.json", "path to start crawling from")
	badURLsFile  = flag.String("bad-urls-file", filepath.Join(os.Getenv("HOME"), ".appledocs", "cache", "known-bad-urls.txt"), "file containing URLs to skip")
	excludePaths = flag.String("exclude-paths", "en-US/docs/Mozilla", "comma-separated list of paths to exclude from crawling")

	// Crawling options
	concurrency   = flag.Int("concurrency", 1, "number of concurrent downloads")
	delay         = flag.Duration("delay", 0, "delay between urls")
	rateLimit     = flag.Float64("rate-limit", 10.0, "requests per second rate limit (0 = no limit)")
	forceRefresh  = flag.Bool("force", false, "force refresh all content")
	timeout       = flag.Duration("timeout", 30*time.Second, "HTTP request timeout")
	maxTime       = flag.Duration("max-time", time.Hour, "maximum time to run the program")
	skipSymbols   = flag.Bool("skip-symbols", false, "skip individual symbol level documentation")
	printURLs     = flag.Bool("print-urls", false, "only print discovered URLs from entry point and exit")

	// Output options
	prettyJSON    = flag.Bool("pretty", true, "pretty-print JSON files")
	verbose       = flag.Bool("verbose", false, "enable verbose logging")
	logLevel      = flag.String("log-level", "info", "log level: debug, info, warn, error")
	exportMetrics = flag.String("export-metrics", "", "export detailed metrics to JSON file (optional path)")
	validateCache = flag.Bool("validate-cache", false, "validate cache integrity on startup")
	checksumValidation = flag.Bool("checksum-validation", false, "enable enhanced checksum-based cache validation")
	fetchBothLanguages = flag.Bool("fetch-both-languages", true, "fetch both Swift and Objective-C variants")

	// Mode selection
	mode = flag.String("mode", "crawl", "operation mode: crawl, html, markdown, gentypes, analyze, list-demos, or all")

	// list-demos options
	listDemosFramework = flag.String("framework", "", "filter demo code by framework name (optional)")
	downloadDemos = flag.Bool("download", false, "download demo code (requires -mode list-demos)")
	demosOutputDir = flag.String("demos-output", filepath.Join(os.Getenv("HOME"), "go", "src", "github.com", "tmc", "appledocs-examples"), "directory to store downloaded demo code")

	// Markdown-specific options
	mdOutputDir = flag.String("md-output", "markdown", "directory to store Markdown documentation")

	// Type generation options
	genTypesOutput = flag.String("gentypes-output", "types/types_generated.go", "output file for generated types")
	genTypesMaxFiles = flag.Int("gentypes-max-files", 1000, "maximum number of files to scan for type generation")

	// Legacy flags for backward compatibility
	generateMD = flag.Bool("markdown", false, "generate Markdown documentation from the JSON files")
)

// JSONFileEntry represents a found JSON file
type JSONFileEntry struct {
	Path string
	URL  string
}

// crawler holds all the application settings
type crawler struct {
	client       *http.Client
	visitedURLs  sync.Map // map[string]bool - tracks visited URLs, safe for concurrent access
	jsonEntries  []JSONFileEntry
	entriesMutex sync.Mutex
	processedCount int
	badURLs        map[string]bool // URLs known to be 404s or invalid
	urlDepths      map[string]int  // Track semantic depth of each URL
	depthMutex     sync.RWMutex    // Mutex for urlDepths
	rateLimiter    *rate.Limiter   // Rate limiter for HTTP requests

	// Status tracking metrics
	cacheHits      int
	cacheMisses    int
	errors         int
	skippedURLs    int
	skippedSymbols int // Count of URLs skipped because they're symbol-level docs
	statsMutex     sync.Mutex

	// Enhanced metrics
	startTime         time.Time
	totalBytesDownloaded int64
	totalBytesFromCache  int64
	avgResponseTime      time.Duration
	totalResponseTime    time.Duration
	requestCount         int
	httpErrors           map[int]int // HTTP status code -> count
	retryCount           int         // Total number of retries
	frameworkCount       int         // Number of frameworks processed
	classCount           int         // Number of classes processed
	methodCount          int         // Number of methods processed
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
func fetchAndExtractURLs(ctx context.Context, client *http.Client, app *crawler, fetchURL string) ([]string, error) {
	log.Printf("Fetching URL: %s", fetchURL)
	data, err := fetchWithCache(ctx, client, fetchURL, app)
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
func printURLsOnly(ctx context.Context) error {
	client := &http.Client{Timeout: *timeout}

	// Initialize rate limiter
	var rateLimiter *rate.Limiter
	if *rateLimit > 0 {
		rateLimiter = rate.NewLimiter(rate.Limit(*rateLimit), int(*rateLimit))
		if logger != nil {
			logger.Info("Rate limiting enabled", "requests_per_second", *rateLimit)
		}
	} else {
		rateLimiter = rate.NewLimiter(rate.Inf, 0) // No limit
	}

	// Create simple app instance for cache tracking
	app := &crawler{
		client:      client,
		badURLs:     make(map[string]bool),
		urlDepths:   make(map[string]int),
		rateLimiter: rateLimiter,
		startTime:   time.Now(),
		httpErrors:  make(map[int]int),
	}
	// visitedURLs is a sync.Map, no initialization needed

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
			extractedURLs, err := fetchAndExtractURLs(ctx, client, app, fullURL)
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
		extractedURLs, err := fetchAndExtractURLs(ctx, client, app, startURL)
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

// initLogger initializes the structured logger
func initLogger() error {
	var level slog.Level
	switch strings.ToLower(*logLevel) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn", "warning":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		return fmt.Errorf("invalid log level: %s", *logLevel)
	}

	opts := &slog.HandlerOptions{
		Level: level,
		AddSource: level == slog.LevelDebug,
	}

	// Use text handler for human-readable logs
	handler := slog.NewTextHandler(os.Stderr, opts)
	logger = slog.New(handler)
	
	// Set as default logger
	slog.SetDefault(logger)
	
	return nil
}

func main() {
	flag.Parse()

	// Handle version flag
	if *showVersion {
		fmt.Printf("appledocs version %s\n", version)
		fmt.Printf("  commit: %s\n", commit)
		fmt.Printf("  built:  %s\n", buildTime)
		os.Exit(0)
	}

	// Initialize structured logger
	if err := initLogger(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Enhanced command-line flag validation
	flagValidation := ValidateCommandLineFlags()
	if !flagValidation.Valid {
		log.Printf("Command-line validation errors:")
		for _, err := range flagValidation.Errors {
			log.Printf("  - %s", err.Error())
		}
		os.Exit(1)
	}
	
	// Print warnings if any
	if len(flagValidation.Warnings) > 0 {
		log.Printf("Command-line validation warnings:")
		for _, warning := range flagValidation.Warnings {
			log.Printf("  - %s", warning.Error())
		}
	}

	// Handle legacy flag conversion for backward compatibility
	if *generateMD && *mode == "crawl" {
		*mode = "markdown"
	}

	// Special case for print-urls mode
	if *printURLs {
		ctx := context.Background()
		if err := printURLsOnly(ctx); err != nil {
			log.Fatalf("Error: %v", err)
		}
		return
	}

	// Validate mode
	validModes := map[string]bool{"crawl": true, "html": true, "markdown": true, "gentypes": true, "analyze": true, "list-demos": true, "all": true}
	if !validModes[*mode] {
		log.Fatalf("Invalid mode: %s. Must be one of: crawl, html, markdown, gentypes, analyze, list-demos, or all", *mode)
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
		logger.Info("Starting HTML index generation", "output_dir", *outputDir)
		jsonFiles, err := scanOutputDirectory(*outputDir)
		if err != nil {
			logger.Error("Failed to scan output directory", "error", err, "dir", *outputDir)
			os.Exit(1)
		}

		if err := createJSONIndexHTML(*outputDir, jsonFiles); err != nil {
			logger.Error("HTML index generation failed", "error", err)
			os.Exit(1)
		}

		logger.Info("HTML index generation completed", "file", filepath.Join(*outputDir, "index.html"))
	}

	// Markdown mode - generate Markdown files
	if *mode == "markdown" || *mode == "all" {
		logger.Info("Starting Markdown generation", "output_dir", *mdOutputDir)
		if err := generateMarkdown(*outputDir, *mdOutputDir); err != nil {
			logger.Error("Markdown generation failed", "error", err)
			os.Exit(1)
		}
		logger.Info("Markdown generation completed", "output_dir", *mdOutputDir)
	}

	// Type generation mode - generate Go types from JSON schema
	if *mode == "gentypes" || *mode == "all" {
		logger.Info("Starting type generation", "output", *genTypesOutput, "max_files", *genTypesMaxFiles)
		docsPath := filepath.Join(*outputDir, "tutorials", "data", "documentation")
		if err := generateTypes(docsPath, *genTypesOutput, *genTypesMaxFiles); err != nil {
			logger.Error("Type generation failed", "error", err)
			os.Exit(1)
		}
		logger.Info("Type generation completed", "output", *genTypesOutput)
	}

	// Schema analysis mode - analyze JSON schema patterns
	if *mode == "analyze" {
		logger.Info("Starting schema analysis", "max_files", *genTypesMaxFiles)
		docsPath := filepath.Join(*outputDir, "tutorials", "data", "documentation")
		if err := analyzeSchema(docsPath, *genTypesMaxFiles); err != nil {
			logger.Error("Schema analysis failed", "error", err)
			os.Exit(1)
		}
	}

	// List demos mode - list available sample code downloads
	if *mode == "list-demos" {
		if *downloadDemos {
			logger.Info("Downloading demo code", "output", *demosOutputDir)
		} else {
			logger.Info("Listing demo code from cached documentation")
		}
		if err := listDemoCode(*cacheDir, *listDemosFramework, *downloadDemos, *demosOutputDir); err != nil {
			logger.Error("Demo listing failed", "error", err)
			os.Exit(1)
		}
	}
}

// run is the main entry point for the application logic
func run(ctx context.Context) error {
	client := &http.Client{Timeout: *timeout}
	startURL := resolveURL(*baseURL, *entryPoint)

	// Initialize rate limiter
	var rateLimiter *rate.Limiter
	if *rateLimit > 0 {
		rateLimiter = rate.NewLimiter(rate.Limit(*rateLimit), int(*rateLimit))
		if logger != nil {
			logger.Info("Rate limiting enabled", "requests_per_second", *rateLimit)
		}
	} else {
		rateLimiter = rate.NewLimiter(rate.Inf, 0) // No limit
	}

	app := &crawler{
		client:      client,
		badURLs:     make(map[string]bool),
		urlDepths:   make(map[string]int),
		rateLimiter: rateLimiter,
		startTime:   time.Now(),
		httpErrors:  make(map[int]int),
	}
	// visitedURLs is a sync.Map, no initialization needed

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

	// Validate cache integrity if requested
	if *validateCache {
		log.Printf("Validating cache integrity...")
		var cacheValidation ValidationResult
		
		if *checksumValidation {
			log.Printf("Using enhanced checksum-based validation...")
			cacheValidation = ValidateCacheIntegrityWithChecksums(*cacheDir)
		} else {
			cacheValidation = ValidateCache(*cacheDir)
		}
		
		if len(cacheValidation.Errors) > 0 {
			log.Printf("Cache validation errors found:")
			for _, err := range cacheValidation.Errors {
				log.Printf("  - %s", err.Error())
			}
		}
		if len(cacheValidation.Warnings) > 0 {
			log.Printf("Cache validation warnings:")
			for _, warning := range cacheValidation.Warnings {
				log.Printf("  - %s", warning.Error())
			}
		}
		if len(cacheValidation.Errors) == 0 && len(cacheValidation.Warnings) == 0 {
			log.Printf("Cache validation completed successfully")
		}
	}

	// Mark the start URL as visited and set its depth to 0
	app.visitedURLs.Store(startURL, true)
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
				metrics := app.getEnhancedMetrics()
				
				// Enhanced progress reporting with more detail
				log.Printf("Progress: %d files processed | Cache: %d hits (%.1f%%), %d misses | %.2f MB/s | Avg: %v/req | Errors: %d | Retries: %d", 
					metrics.Processed, 
					metrics.CacheHits, 
					metrics.CacheHitRate,
					metrics.CacheMisses,
					metrics.DownloadRate,
					metrics.AvgResponseTime,
					metrics.Errors,
					metrics.RetryCount)
				
				// Additional metrics when verbose
				if *verbose && metrics.EstimatedTimeRemaining > 0 {
					log.Printf("Content: %d frameworks, %d classes, %d methods | ETA: %v", 
						metrics.FrameworkCount, 
						metrics.ClassCount, 
						metrics.MethodCount,
						metrics.EstimatedTimeRemaining.Round(time.Second))
				}
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
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Worker %d panic recovered: %v", workerID, r)
				}
			}()
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

	// Report enhanced final statistics
	metrics := app.getEnhancedMetrics()
	log.Printf("Final statistics:")
	log.Printf("  - Runtime: %v", metrics.RuntimeDuration.Round(time.Second))
	log.Printf("  - Processed: %d JSON files (%.2f files/sec)", metrics.Processed, metrics.ProcessingRate)
	log.Printf("  - Cache: %d hits (%.1f%%), %d misses", metrics.CacheHits, metrics.CacheHitRate, metrics.CacheMisses)
	log.Printf("  - Data: %.2f MB downloaded, %.2f MB from cache (%.2f MB/s)", 
		float64(metrics.TotalBytesDownloaded)/1024/1024,
		float64(metrics.TotalBytesFromCache)/1024/1024,
		metrics.DownloadRate)
	log.Printf("  - Network: %d requests, avg %v/req, %d retries", 
		metrics.RequestCount, metrics.AvgResponseTime, metrics.RetryCount)
	log.Printf("  - Content: %d frameworks, %d classes, %d methods", 
		metrics.FrameworkCount, metrics.ClassCount, metrics.MethodCount)
	log.Printf("  - Errors: %d", metrics.Errors)
	log.Printf("  - Skipped URLs: %d", metrics.SkippedURLs)
	if *skipSymbols {
		log.Printf("  - Skipped symbol URLs: %d", metrics.SkippedSymbols)
	}
	
	// Report HTTP errors if any
	if len(metrics.HTTPErrors) > 0 {
		log.Printf("  - HTTP Errors by status code:")
		for statusCode, count := range metrics.HTTPErrors {
			log.Printf("    - %d: %d errors", statusCode, count)
		}
	}

	// Export metrics to JSON file if requested
	if *exportMetrics != "" {
		metricsPath := *exportMetrics
		if metricsPath == "true" || metricsPath == "1" {
			metricsPath = "metrics.json"
		}
		
		metricsJSON, err := json.MarshalIndent(metrics, "", "  ")
		if err != nil {
			log.Printf("Warning: Failed to marshal metrics: %v", err)
		} else {
			if err := os.WriteFile(metricsPath, metricsJSON, 0644); err != nil {
				log.Printf("Warning: Failed to write metrics file %s: %v", metricsPath, err)
			} else {
				log.Printf("Detailed metrics exported to: %s", metricsPath)
			}
		}
	}

	return nil
}

// Helper methods for incrementing metrics
func (app *crawler) incrementCacheHits() {
	app.statsMutex.Lock()
	app.cacheHits++
	app.statsMutex.Unlock()
}

func (app *crawler) incrementCacheMisses() {
	app.statsMutex.Lock()
	app.cacheMisses++
	app.statsMutex.Unlock()
}

func (app *crawler) incrementErrors() {
	app.statsMutex.Lock()
	app.errors++
	app.statsMutex.Unlock()
}

func (app *crawler) incrementSkippedURLs() {
	app.statsMutex.Lock()
	app.skippedURLs++
	app.statsMutex.Unlock()
}

func (app *crawler) incrementSkippedSymbols() {
	app.statsMutex.Lock()
	app.skippedSymbols++
	app.skippedURLs++ // Also count in the general skipped URLs
	app.statsMutex.Unlock()
}

// Enhanced metrics methods
func (app *crawler) recordResponseTime(duration time.Duration) {
	app.statsMutex.Lock()
	app.totalResponseTime += duration
	app.requestCount++
	if app.requestCount > 0 {
		app.avgResponseTime = app.totalResponseTime / time.Duration(app.requestCount)
	}
	app.statsMutex.Unlock()
}

func (app *crawler) recordBytesDownloaded(bytes int64) {
	app.statsMutex.Lock()
	app.totalBytesDownloaded += bytes
	app.statsMutex.Unlock()
}

func (app *crawler) recordBytesFromCache(bytes int64) {
	app.statsMutex.Lock()
	app.totalBytesFromCache += bytes
	app.statsMutex.Unlock()
}

func (app *crawler) recordHTTPError(statusCode int) {
	app.statsMutex.Lock()
	if app.httpErrors == nil {
		app.httpErrors = make(map[int]int)
	}
	app.httpErrors[statusCode]++
	app.statsMutex.Unlock()
}

func (app *crawler) incrementRetryCount() {
	app.statsMutex.Lock()
	app.retryCount++
	app.statsMutex.Unlock()
}

func (app *crawler) recordContentType(path string) {
	app.statsMutex.Lock()
	defer app.statsMutex.Unlock()
	
	// Classify content based on URL path patterns
	if strings.Contains(path, "/documentation/") {
		pathParts := strings.Split(path, "/")
		// Count frameworks (depth 3: /tutorials/data/documentation/FrameworkName.json)
		if len(pathParts) >= 4 && strings.HasSuffix(pathParts[3], ".json") && !strings.Contains(pathParts[3], "/") {
			app.frameworkCount++
		}
		// Count classes (depth 4: /tutorials/data/documentation/Framework/Class.json)
		if len(pathParts) >= 5 && strings.HasSuffix(pathParts[4], ".json") {
			app.classCount++
		}
		// Count methods (depth 5+: deeper nesting indicates symbols/methods)
		if len(pathParts) >= 6 {
			app.methodCount++
		}
	}
}

// MetricsSnapshot represents a comprehensive snapshot of all metrics
type MetricsSnapshot struct {
	// Basic metrics
	Processed    int    `json:"processed"`
	CacheHits    int    `json:"cache_hits"`
	CacheMisses  int    `json:"cache_misses"`
	Errors       int    `json:"errors"`
	SkippedURLs  int    `json:"skipped_urls"`
	SkippedSymbols int  `json:"skipped_symbols"`
	
	// Enhanced metrics
	StartTime            time.Time         `json:"start_time"`
	RuntimeDuration      time.Duration     `json:"runtime_duration"`
	TotalBytesDownloaded int64             `json:"total_bytes_downloaded"`
	TotalBytesFromCache  int64             `json:"total_bytes_from_cache"`
	AvgResponseTime      time.Duration     `json:"avg_response_time"`
	RequestCount         int               `json:"request_count"`
	HTTPErrors           map[int]int       `json:"http_errors"`
	RetryCount           int               `json:"retry_count"`
	FrameworkCount       int               `json:"framework_count"`
	ClassCount           int               `json:"class_count"`
	MethodCount          int               `json:"method_count"`
	
	// Calculated metrics
	CacheHitRate         float64           `json:"cache_hit_rate"`
	DownloadRate         float64           `json:"download_rate_mbps"`
	ProcessingRate       float64           `json:"processing_rate_per_sec"`
	EstimatedTimeRemaining time.Duration   `json:"estimated_time_remaining"`
}

// getStats returns current statistics in a thread-safe way (legacy method)
func (app *crawler) getStats() (int, int, int, int, int) {
	app.statsMutex.Lock()
	defer app.statsMutex.Unlock()
	app.entriesMutex.Lock()
	processed := len(app.jsonEntries)
	app.entriesMutex.Unlock()

	// Note: depthLimits is counted as part of skippedURLs for backward compatibility
	// with the existing reporting, so we don't need to return it separately
	return processed, app.cacheHits, app.cacheMisses, app.errors, app.skippedURLs
}

// getEnhancedMetrics returns comprehensive metrics snapshot
func (app *crawler) getEnhancedMetrics() MetricsSnapshot {
	app.statsMutex.Lock()
	defer app.statsMutex.Unlock()
	app.entriesMutex.Lock()
	processed := len(app.jsonEntries)
	app.entriesMutex.Unlock()
	
	now := time.Now()
	runtime := now.Sub(app.startTime)
	
	// Calculate derived metrics
	var cacheHitRate float64
	totalRequests := app.cacheHits + app.cacheMisses
	if totalRequests > 0 {
		cacheHitRate = float64(app.cacheHits) / float64(totalRequests) * 100
	}
	
	// Download rate in MB/s
	var downloadRate float64
	if runtime.Seconds() > 0 {
		totalMB := float64(app.totalBytesDownloaded) / 1024 / 1024
		downloadRate = totalMB / runtime.Seconds()
	}
	
	// Processing rate (files per second)
	var processingRate float64
	if runtime.Seconds() > 0 {
		processingRate = float64(processed) / runtime.Seconds()
	}
	
	// Estimate time remaining (very rough estimate)
	var estimatedTimeRemaining time.Duration
	// Count visited URLs (sync.Map doesn't have a Len method)
	totalURLs := 0
	app.visitedURLs.Range(func(_, _ interface{}) bool {
		totalURLs++
		return true
	})

	if processed > 0 && totalURLs > processed && processingRate > 0 {
		remaining := totalURLs - processed
		estimatedTimeRemaining = time.Duration(float64(remaining)/processingRate) * time.Second
	}
	
	return MetricsSnapshot{
		Processed:              processed,
		CacheHits:              app.cacheHits,
		CacheMisses:            app.cacheMisses,
		Errors:                 app.errors,
		SkippedURLs:            app.skippedURLs,
		SkippedSymbols:         app.skippedSymbols,
		StartTime:              app.startTime,
		RuntimeDuration:        runtime,
		TotalBytesDownloaded:   app.totalBytesDownloaded,
		TotalBytesFromCache:    app.totalBytesFromCache,
		AvgResponseTime:        app.avgResponseTime,
		RequestCount:           app.requestCount,
		HTTPErrors:             app.httpErrors,
		RetryCount:             app.retryCount,
		FrameworkCount:         app.frameworkCount,
		ClassCount:             app.classCount,
		MethodCount:            app.methodCount,
		CacheHitRate:           cacheHitRate,
		DownloadRate:           downloadRate,
		ProcessingRate:         processingRate,
		EstimatedTimeRemaining: estimatedTimeRemaining,
	}
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
func (app *crawler) processURL(ctx context.Context, u string, urlQueue chan<- string) error {
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
	data, err := fetchWithCache(ctx, app.client, u, app)
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

	// Record content type for metrics
	app.recordContentType(outputPath)

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
func (app *crawler) queueNewURLs(newURLs []string, urlQueue chan<- string) int {
	var added int

	for _, newURL := range newURLs {
		// Get all language variants for this URL
		resolvedURLs := resolveURLsWithLanguageVariants(*baseURL, newURL)

		for _, resolvedURL := range resolvedURLs {
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

			// Check if URL has been visited using atomic LoadOrStore
			// This eliminates the race condition that existed with double-check locking
			_, alreadyVisited := app.visitedURLs.LoadOrStore(resolvedURL, true)
			if alreadyVisited {
				continue
			}

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
func fetchWithCache(ctx context.Context, client *http.Client, u string, app *crawler) ([]byte, error) {
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

	// ETag cache path
	etagPath := cachePath + ".etag"

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
					// Check if we have an ETag to validate freshness
					if etagData, err := os.ReadFile(etagPath); err == nil && len(etagData) > 0 {
						etag := strings.TrimSpace(string(etagData))
						if *verbose {
							log.Printf("Validating cache for %q with ETag: %s", u, etag)
						}
						// We'll validate with If-None-Match below
						// For now, fall through to make the request
					} else {
						// No ETag, use cache as-is (legacy behavior)
						app.incrementCacheHits()
						app.recordBytesFromCache(int64(len(data)))
						if *verbose {
							log.Printf("Cache hit for %q (no ETag)", u)
						}
						return data, nil
					}
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
	req, err := http.NewRequestWithContext(ctx, "GET", u, nil)
	if err != nil {
		app.incrementErrors()
		return nil, fmt.Errorf("create request for %q: %v", u, err)
	}

	// Add browser-like headers
	addBrowserLikeHeaders(req)

	// Add If-None-Match header if we have an ETag
	if etagData, err := os.ReadFile(etagPath); err == nil && len(etagData) > 0 {
		etag := strings.TrimSpace(string(etagData))
		req.Header.Set("If-None-Match", etag)
		if *verbose {
			log.Printf("Sending If-None-Match: %s for %q", etag, u)
		}
	}

	// Execute the request with retry logic
	var resp *http.Response
	maxRetries := 3
	for attempt := 0; attempt <= maxRetries; attempt++ {
		// Apply rate limiting before making request
		if err := app.rateLimiter.Wait(ctx); err != nil {
			return nil, fmt.Errorf("rate limit wait failed: %v", err)
		}

		var err error
		startTime := time.Now()
		resp, err = client.Do(req)
		requestDuration := time.Since(startTime)
		
		// Record response time for successful requests
		if err == nil {
			app.recordResponseTime(requestDuration)
		}
		
		// Check if we should retry this attempt
		shouldRetry := false
		var retryReason string
		
		if err != nil {
			// Check if it's a retryable network error
			isRetryable := strings.Contains(err.Error(), "timeout") ||
				strings.Contains(err.Error(), "connection reset") ||
				strings.Contains(err.Error(), "temporary failure")
			
			if isRetryable && attempt < maxRetries {
				shouldRetry = true
				retryReason = fmt.Sprintf("network error: %v", err)
				app.incrementRetryCount()
			} else {
				app.incrementErrors()
				// Add to bad URLs list if it's a permanent network error
				if strings.Contains(err.Error(), "no such host") ||
					strings.Contains(err.Error(), "connection refused") ||
					(!isRetryable && strings.Contains(err.Error(), "timeout")) {
					app.badURLs[u] = true
					appendToBadURLsFile(u)
				}
				return nil, fmt.Errorf("fetch %q (after %d attempts): %v", u, attempt+1, err)
			}
		} else if resp.StatusCode == http.StatusNotModified {
			// 304 Not Modified - content hasn't changed, use cached version
			resp.Body.Close()
			if *verbose {
				log.Printf("304 Not Modified for %q, using cached content", u)
			}

			// Read cached content
			cachedData, err := os.ReadFile(cachePath)
			if err != nil {
				app.incrementErrors()
				return nil, fmt.Errorf("read cached content after 304 for %q: %v", u, err)
			}

			app.incrementCacheHits()
			app.recordBytesFromCache(int64(len(cachedData)))
			return cachedData, nil
		} else if resp.StatusCode != http.StatusOK {
			// Record HTTP error
			app.recordHTTPError(resp.StatusCode)

			// Check if it's a retryable status code
			isRetryableStatus := resp.StatusCode == http.StatusTooManyRequests ||
				resp.StatusCode == http.StatusInternalServerError ||
				resp.StatusCode == http.StatusBadGateway ||
				resp.StatusCode == http.StatusServiceUnavailable ||
				resp.StatusCode == http.StatusGatewayTimeout

			if isRetryableStatus && attempt < maxRetries {
				shouldRetry = true
				retryReason = fmt.Sprintf("HTTP %d: %s", resp.StatusCode, resp.Status)
				app.incrementRetryCount()
				resp.Body.Close() // Close the response body before retrying
			} else {
				// Handle non-retryable status codes or exhausted retries
				defer resp.Body.Close()
				app.incrementErrors()

				// Add to bad URLs list for client errors that won't be resolved by retrying
				if resp.StatusCode == http.StatusForbidden ||
					resp.StatusCode == http.StatusNotFound ||
					resp.StatusCode == http.StatusMethodNotAllowed ||
					resp.StatusCode == http.StatusGone {
					app.badURLs[u] = true
					appendToBadURLsFile(u)
					if *verbose {
						log.Printf("Added bad URL due to %d status: %s", resp.StatusCode, u)
					}
				}
				return nil, fmt.Errorf("fetch %q: %s", u, resp.Status)
			}
		} else {
			// Success, break out of retry loop
			break
		}
		
		// Handle retry backoff
		if shouldRetry {
			backoffDuration := time.Duration(1<<uint(attempt)) * time.Second
			if *verbose {
				log.Printf("Attempt %d failed for %q (%s), retrying in %v", attempt+1, u, retryReason, backoffDuration)
			}
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(backoffDuration):
				// Continue to next attempt
			}
		}
	}
	defer resp.Body.Close()

	// Read response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		app.incrementErrors()
		return nil, fmt.Errorf("read %q: %v", u, err)
	}

	// Record bytes downloaded
	app.recordBytesDownloaded(int64(len(data)))

	// Validate data integrity
	if validation := ValidateDataIntegrity(u, data); !validation.Valid {
		app.incrementErrors()
		if *verbose {
			log.Printf("Data validation failed for %q:", u)
			for _, err := range validation.Errors {
				log.Printf("  - %s", err.Error())
			}
		}
		// Still cache the data but warn about issues
		for _, warning := range validation.Warnings {
			if *verbose {
				log.Printf("  Warning: %s", warning.Error())
			}
		}
	}

	// Write to cache with atomic operation
	if err := writeFileAtomic(cachePath, data); err != nil {
		app.incrementErrors()
		return nil, fmt.Errorf("write cache %q: %v", cachePath, err)
	}

	// Store ETag if present in response headers
	if etag := resp.Header.Get("Etag"); etag != "" {
		if err := os.WriteFile(etagPath, []byte(etag), 0644); err != nil {
			// Log but don't fail - ETag storage is optional
			if *verbose {
				log.Printf("Warning: Failed to write ETag for %q: %v", u, err)
			}
		} else if *verbose {
			log.Printf("Stored ETag %s for %q", etag, u)
		}
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

	return writeFileAtomic(outputPath, content)
}

// writeFileAtomic writes data to a file atomically by writing to a temp file first
func writeFileAtomic(filename string, data []byte) error {
	// Create a temporary file in the same directory
	dir := filepath.Dir(filename)
	tmpFile, err := os.CreateTemp(dir, ".tmp-appledocs-*")
	if err != nil {
		return fmt.Errorf("create temporary file: %v", err)
	}
	tmpPath := tmpFile.Name()
	
	// Ensure cleanup of temp file on error
	defer func() {
		if tmpFile != nil {
			tmpFile.Close()
			os.Remove(tmpPath)
		}
	}()
	
	// Write data to temp file
	if _, err := tmpFile.Write(data); err != nil {
		return fmt.Errorf("write to temporary file: %v", err)
	}
	
	// Sync to ensure data is written to disk
	if err := tmpFile.Sync(); err != nil {
		return fmt.Errorf("sync temporary file: %v", err)
	}
	
	// Close temp file
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("close temporary file: %v", err)
	}
	tmpFile = nil // Mark as closed to avoid double-close in defer
	
	// Atomically move temp file to final location
	if err := os.Rename(tmpPath, filename); err != nil {
		return fmt.Errorf("rename temporary file: %v", err)
	}
	
	return nil
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

// resolveURLsWithLanguageVariants returns URLs for both language variants if enabled
func resolveURLsWithLanguageVariants(base, relative string) []string {
	baseURL := resolveURL(base, relative)

	// Only add language variants for .json documentation URLs
	if !*fetchBothLanguages || !strings.HasSuffix(baseURL, ".json") || !strings.Contains(baseURL, "/tutorials/data/documentation/") {
		return []string{baseURL}
	}

	// Check if URL already has a language parameter
	if strings.Contains(baseURL, "?language=") || strings.Contains(baseURL, "&language=") {
		return []string{baseURL}
	}

	// Generate both Swift and Objective-C variants
	urls := make([]string, 0, 2)

	// Check if URL already has query parameters
	if strings.Contains(baseURL, "?") {
		urls = append(urls, baseURL+"&language=swift")
		urls = append(urls, baseURL+"&language=objc")
	} else {
		urls = append(urls, baseURL+"?language=swift")
		urls = append(urls, baseURL+"?language=objc")
	}

	return urls
}

// extractJSONURLs extracts URLs to other JSON files from a JSON response.
func extractJSONURLs(data []byte) []string {
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil
	}

	// Get a slice from the pool
	urls := urlSlicePool.Get().([]string)
	urls = urls[:0] // Reset length but keep capacity
	
	extractJSONURLsFromValue(result, &urls)
	
	// Create a copy to return and put the slice back in the pool
	result_urls := make([]string, len(urls))
	copy(result_urls, urls)
	urlSlicePool.Put(urls)
	
	return result_urls
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

	// Generate HTML using the function from html.go
	indexPath := filepath.Join(outputDir, "index.html")
	return generateHTMLFile(indexPath, data)
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
func loadBadURLs(app *crawler) error {
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
func writeBadURLsFile(app *crawler) error {
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

// DemoCodeInfo holds information about a sample code download
type DemoCodeInfo struct {
	Title      string
	Framework  string
	DownloadURL string
	DocURL     string
}

// listDemoCode scans cached documentation and lists all available sample code downloads
func listDemoCode(cacheDir, frameworkFilter string, download bool, outputDir string) error {
	demos := []DemoCodeInfo{}

	// Build the search path
	searchPath := filepath.Join(cacheDir, "developer.apple.com", "tutorials", "data", "documentation")

	// Walk through cached JSON files
	err := filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip files with errors
		}

		if !info.IsDir() && filepath.Ext(path) == ".json" {
			// Extract framework name from path
			relPath, _ := filepath.Rel(searchPath, path)
			framework := filepath.Base(filepath.Dir(relPath))

			// Apply framework filter if specified
			if frameworkFilter != "" && !strings.EqualFold(framework, frameworkFilter) {
				return nil
			}

			// Read and parse the JSON file
			data, err := os.ReadFile(path)
			if err != nil {
				return nil // Skip unreadable files
			}

			var doc map[string]interface{}
			if err := json.Unmarshal(data, &doc); err != nil {
				return nil // Skip malformed JSON
			}

			// Check for sampleCodeDownload field
			if sampleCodeDownload, ok := doc["sampleCodeDownload"].(map[string]interface{}); ok {
				if action, ok := sampleCodeDownload["action"].(map[string]interface{}); ok {
					if downloadPath, ok := action["identifier"].(string); ok {
						// Build full download URL
						var downloadURL string
						if strings.HasPrefix(downloadPath, "http://") || strings.HasPrefix(downloadPath, "https://") {
							downloadURL = downloadPath
						} else {
							downloadURL = "https://docs-assets.developer.apple.com/published/" + downloadPath
						}

						// Get title from metadata or identifier
						title := ""
						if metadata, ok := doc["metadata"].(map[string]interface{}); ok {
							if t, ok := metadata["title"].(string); ok {
								title = t
							}
						}
						if title == "" {
							if identifier, ok := doc["identifier"].(map[string]interface{}); ok {
								if url, ok := identifier["url"].(string); ok {
									title = url
								}
							}
						}

						// Build doc URL
						docURL := ""
						if identifier, ok := doc["identifier"].(map[string]interface{}); ok {
							if url, ok := identifier["url"].(string); ok {
								docURL = "https://developer.apple.com" + url
							}
						}

						demos = append(demos, DemoCodeInfo{
							Title:       title,
							Framework:   framework,
							DownloadURL: downloadURL,
							DocURL:      docURL,
						})
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to walk cache directory: %w", err)
	}

	// Sort by framework, then title
	sort.Slice(demos, func(i, j int) bool {
		if demos[i].Framework != demos[j].Framework {
			return demos[i].Framework < demos[j].Framework
		}
		return demos[i].Title < demos[j].Title
	})

	// Download demos if requested
	if download {
		if err := downloadDemoCode(demos, outputDir); err != nil {
			return fmt.Errorf("failed to download demo code: %w", err)
		}
		return nil
	}

	// Print results
	fmt.Printf("\nFound %d demo code examples:\n\n", len(demos))

	currentFramework := ""
	for _, demo := range demos {
		if demo.Framework != currentFramework {
			currentFramework = demo.Framework
			fmt.Printf("\n%s:\n", currentFramework)
			fmt.Printf("%s\n", strings.Repeat("-", len(currentFramework)+1))
		}
		fmt.Printf("  • %s\n", demo.Title)
		fmt.Printf("    Download: %s\n", demo.DownloadURL)
		if demo.DocURL != "" {
			fmt.Printf("    Documentation: %s\n", demo.DocURL)
		}
		fmt.Println()
	}

	return nil
}

// downloadDemoCode downloads all demo code examples to the output directory
func downloadDemoCode(demos []DemoCodeInfo, outputDir string) error {
	// Create output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	fmt.Printf("\nDownloading %d demo code examples to %s\n\n", len(demos), outputDir)

	// Track statistics
	successful := 0
	failed := 0

	// Download each demo
	for i, demo := range demos {
		// Create framework subdirectory
		frameworkDir := filepath.Join(outputDir, demo.Framework)
		if err := os.MkdirAll(frameworkDir, 0755); err != nil {
			logger.Error("Failed to create framework directory", "framework", demo.Framework, "error", err)
			failed++
			continue
		}

		// Extract filename from URL
		filename := filepath.Base(demo.DownloadURL)
		outputPath := filepath.Join(frameworkDir, filename)

		// Skip if already downloaded
		if _, err := os.Stat(outputPath); err == nil {
			fmt.Printf("[%d/%d] ✓ %s (already exists)\n", i+1, len(demos), demo.Title)
			successful++
			continue
		}

		// Download the file
		fmt.Printf("[%d/%d] Downloading %s...\n", i+1, len(demos), demo.Title)
		resp, err := http.Get(demo.DownloadURL)
		if err != nil {
			logger.Error("Failed to download demo", "title", demo.Title, "url", demo.DownloadURL, "error", err)
			failed++
			continue
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			logger.Error("Failed to download demo", "title", demo.Title, "status", resp.Status)
			failed++
			continue
		}

		// Write to file
		outFile, err := os.Create(outputPath)
		if err != nil {
			resp.Body.Close()
			logger.Error("Failed to create output file", "path", outputPath, "error", err)
			failed++
			continue
		}

		_, err = io.Copy(outFile, resp.Body)
		outFile.Close()
		resp.Body.Close()

		if err != nil {
			os.Remove(outputPath) // Clean up partial download
			logger.Error("Failed to write demo file", "path", outputPath, "error", err)
			failed++
			continue
		}

		fmt.Printf("[%d/%d] ✓ %s\n", i+1, len(demos), demo.Title)
		successful++
	}

	// Print summary
	fmt.Printf("\n%s\n", strings.Repeat("=", 60))
	fmt.Printf("Download complete: %d successful, %d failed\n", successful, failed)
	fmt.Printf("Output directory: %s\n", outputDir)
	fmt.Printf("%s\n\n", strings.Repeat("=", 60))

	return nil
}
