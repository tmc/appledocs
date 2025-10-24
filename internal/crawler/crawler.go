package crawler

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/time/rate"
)

// Pool for reusing slice allocations
var urlSlicePool = sync.Pool{
	New: func() interface{} {
		return make([]string, 0, 100) // Pre-allocate capacity
	},
}

// JSONFileEntry represents a found JSON file
type JSONFileEntry struct {
	Path string
	URL  string
}

// Crawler holds all the application settings
type Crawler struct {
	client           *http.Client
	visitedURLs      sync.Map // map[string]bool - tracks visited URLs, safe for concurrent access
	jsonEntries      []JSONFileEntry
	entriesMutex     sync.Mutex
	processedCount   int
	badURLs          map[string]bool // URLs known to be 404s or invalid
	urlDepths        map[string]int  // Track semantic depth of each URL
	depthMutex       sync.RWMutex    // Mutex for urlDepths
	rateLimiter      *rate.Limiter   // Rate limiter for HTTP requests
	entryPointPrefix string          // Path prefix to restrict crawling scope

	// Work tracking for graceful shutdown
	activeWork   sync.WaitGroup // Tracks active URL processing
	queueClosed  atomic.Bool    // Flag indicating queue is closed

	// Status tracking metrics
	cacheHits      int
	cacheMisses    int
	errors         int
	skippedURLs    int
	skippedSymbols int // Count of URLs skipped because they're symbol-level docs
	statsMutex     sync.Mutex

	// Enhanced metrics
	startTime            time.Time
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

// Config holds configuration for the crawler
type Config struct {
	BaseURL            string
	EntryPoint         string
	OutputDir          string
	CacheDir           string
	BadURLsFile        string
	ExcludePaths       string
	Concurrency        int
	RateLimit          float64
	ForceRefresh       bool
	Timeout            time.Duration
	SkipSymbols        bool
	PrettyJSON         bool
	FetchBothLanguages bool
	ChecksumValidation bool
	MaxDepth           int  // Maximum link depth to follow (0 = unlimited)
	ObjCOnly           bool // Only crawl Objective-C types (skip Swift-only docs)
	Verbose            bool
	Logger             *slog.Logger
}

// New creates a new crawler instance
func New(cfg *Config) *Crawler {
	client := &http.Client{Timeout: cfg.Timeout}

	var rateLimiter *rate.Limiter
	if cfg.RateLimit > 0 {
		rateLimiter = rate.NewLimiter(rate.Limit(cfg.RateLimit), int(cfg.RateLimit))
		if cfg.Logger != nil {
			cfg.Logger.Info("Rate limiting enabled", "requests_per_second", cfg.RateLimit)
		}
	} else {
		rateLimiter = rate.NewLimiter(rate.Inf, 0) // No limit
	}

	return &Crawler{
		client:      client,
		badURLs:     make(map[string]bool),
		urlDepths:   make(map[string]int),
		rateLimiter: rateLimiter,
		startTime:   time.Now(),
		httpErrors:  make(map[int]int),
	}
}

// Run executes the crawler with the given configuration
func (c *Crawler) Run(ctx context.Context, cfg *Config) error {
	// Create bad URLs directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(cfg.BadURLsFile), 0755); err != nil {
		log.Printf("Warning: failed to create directory for bad URLs file: %v", err)
	}

	// Load known bad URLs if file exists
	if err := c.loadBadURLs(cfg.BadURLsFile, cfg.Verbose); err != nil && cfg.Verbose {
		log.Printf("Warning: failed to load bad URLs file: %v", err)
	} else if cfg.Verbose {
		log.Printf("Loaded %d known bad URLs", len(c.badURLs))
	}

	// Set entry point prefix to restrict crawling scope
	// Extract the base path from entry point (e.g., /tutorials/data/documentation/foundation/NSOutputStream.json
	// becomes /tutorials/data/documentation/foundation/nsoutputstream)
	if parsed, err := url.Parse(cfg.EntryPoint); err == nil {
		basePath := parsed.Path
		// If EntryPoint is just a framework name (e.g., "CoreVideo"), expand it to full path
		if basePath == "" || (!strings.HasPrefix(basePath, "/") && !strings.Contains(basePath, "/")) {
			// Entry point is just framework name, construct path
			basePath = "/tutorials/data/documentation/" + cfg.EntryPoint
		}
		basePath = strings.TrimSuffix(basePath, ".json")
		basePath = strings.TrimSuffix(basePath, "/index")
		c.entryPointPrefix = strings.ToLower(basePath)
		if cfg.Verbose {
			log.Printf("Restricting crawl to paths under: %s", c.entryPointPrefix)
		}
	}

	// Determine start URLs
	startURLs := c.buildStartURLs(cfg)

	// Mark the start URLs as visited and set their depth to 0
	for _, startURL := range startURLs {
		c.visitedURLs.Store(startURL, true)
		c.depthMutex.Lock()
		c.urlDepths[startURL] = 0
		c.depthMutex.Unlock()
	}

	// Setup progress reporting
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	go c.reportProgress(ctx, ticker, cfg)

	// Create a buffered channel with sufficient capacity
	urlQueue := make(chan string, cfg.Concurrency*5000)

	// Create a context with cancellation for worker control
	workerCtx, cancelWorkers := context.WithCancel(ctx)
	defer cancelWorkers()

	// Create a wait group for the worker pool
	var wg sync.WaitGroup

	// Launch worker pool
	for i := 0; i < cfg.Concurrency; i++ {
		wg.Add(1)
		go c.worker(workerCtx, &wg, i, urlQueue, cfg)
	}

	// Add the initial URLs to the queue and track them
	for _, startURL := range startURLs {
		c.activeWork.Add(1)
		urlQueue <- startURL
	}

	// Launch goroutine to close queue when all work is done
	go func() {
		c.activeWork.Wait()
		c.queueClosed.Store(true)
		close(urlQueue)
		if cfg.Verbose {
			log.Printf("All work completed, closed URL queue")
		}
	}()

	// Wait for context cancellation or queue closure
	select {
	case <-ctx.Done():
		// Context cancelled (timeout or interrupt)
	case <-func() chan struct{} {
		// Wait for queue to close (all work done)
		ch := make(chan struct{})
		go func() {
			// Monitor queueClosed flag
			for !c.queueClosed.Load() {
				time.Sleep(100 * time.Millisecond)
			}
			close(ch)
		}()
		return ch
	}():
		// All work completed naturally
		if cfg.Verbose {
			log.Printf("Crawl completed - all URLs processed")
		}
	}

	// Cancel worker context
	cancelWorkers()

	// Signal workers to finish
	log.Printf("Waiting for workers to complete...")

	// Wait for workers with timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		log.Printf("All workers finished gracefully")
	case <-time.After(5 * time.Second):
		log.Printf("Some workers still running after timeout")
	}

	// Note: urlQueue is already closed by the work tracker goroutine

	// Write bad URLs file
	if badURLCount := len(c.badURLs); badURLCount > 0 {
		log.Printf("Writing %d bad URLs to %s", badURLCount, cfg.BadURLsFile)
		if err := c.writeBadURLsFile(cfg.BadURLsFile); err != nil {
			log.Printf("Warning: Failed to write bad URLs file: %v", err)
		}
	}

	// Report final statistics
	c.reportFinalStats(cfg)

	return nil
}

// buildStartURLs constructs the initial URLs to crawl
func (c *Crawler) buildStartURLs(cfg *Config) []string {
	var startURLs []string

	// Check if entry point is a framework name (not a full path)
	if cfg.EntryPoint != "/tutorials/data/documentation/technologies.json" &&
		!strings.Contains(cfg.EntryPoint, "/") && !strings.HasSuffix(cfg.EntryPoint, ".json") {
		// This looks like a framework name (e.g., "imageio", "coremedia")
		frameworkURLs := BuildFrameworkURLs(cfg.EntryPoint)
		for _, urlPath := range frameworkURLs {
			startURLs = append(startURLs, ResolveURL(cfg.BaseURL, urlPath))
		}
	} else {
		// Use the entry point as-is (full path or default)
		startURLs = []string{ResolveURL(cfg.BaseURL, cfg.EntryPoint)}
	}

	return startURLs
}

// worker processes URLs from the queue
func (c *Crawler) worker(ctx context.Context, wg *sync.WaitGroup, workerID int, urlQueue chan string, cfg *Config) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Worker %d panic recovered: %v", workerID, r)
		}
	}()

	if cfg.Verbose {
		log.Printf("Worker %d starting", workerID)
	}

	for {
		select {
		case <-ctx.Done():
			if cfg.Verbose {
				log.Printf("Worker %d shutting down (context done)", workerID)
			}
			return
		case u, ok := <-urlQueue:
			if !ok {
				if cfg.Verbose {
					log.Printf("Worker %d shutting down (channel closed)", workerID)
				}
				return
			}

			// Process the URL
			if err := c.processURL(ctx, u, urlQueue, cfg); err != nil {
				if err != context.Canceled && err != context.DeadlineExceeded {
					log.Printf("Error processing %q: %v", u, err)
				}
			}

			// Mark this URL as processed
			c.activeWork.Done()
		}
	}
}

// processURL handles a single URL, fetching and processing it
func (c *Crawler) processURL(ctx context.Context, u string, urlQueue chan<- string, cfg *Config) error {
	if cfg.Verbose {
		log.Printf("Processing %s", u)
	}

	// First, check if this URL is already known to be bad
	if c.badURLs[u] {
		if cfg.Verbose {
			log.Printf("Skipping known bad URL: %s", u)
		}
		c.incrementSkippedURLs()
		return fmt.Errorf("known bad URL: %s", u)
	}

	// Parse URL to get relative path
	parsed, err := url.Parse(u)
	if err != nil {
		c.incrementErrors()
		return fmt.Errorf("parse URL %q: %v", u, err)
	}

	// Check if this URL matches any exclude patterns
	if ShouldExcludePath(parsed.Path, cfg.ExcludePaths, cfg.Verbose) {
		c.badURLs[u] = true
		AppendToBadURLsFile(u, cfg.BadURLsFile)
		c.incrementSkippedURLs()
		if cfg.Verbose {
			log.Printf("Skipping excluded path, added to bad URLs: %s", parsed.Path)
		}
		return fmt.Errorf("excluded path: %s", parsed.Path)
	}

	// Check for invalid URL patterns
	if c.shouldSkipURL(parsed) {
		c.badURLs[u] = true
		AppendToBadURLsFile(u, cfg.BadURLsFile)
		c.incrementSkippedURLs()
		if cfg.Verbose {
			log.Printf("Added unsupported URL to bad URLs: %s", u)
		}
		return fmt.Errorf("not a supported URL type: %s", u)
	}

	// Get content (this also caches it)
	data, err := c.fetchWithCache(ctx, u, cfg)
	if err != nil {
		return fmt.Errorf("fetch %q: %v", u, err)
	}

	// Check if this is an Objective-C document when ObjCOnly is enabled
	if cfg.ObjCOnly && !IsObjectiveCDocument(data) {
		if cfg.Verbose {
			log.Printf("Skipping non-Objective-C document: %s", u)
		}
		// Still mark as visited to avoid re-checking
		c.visitedURLs.Store(u, true)
		return nil
	}

	// Record content type for metrics
	outputPath := strings.TrimPrefix(parsed.Path, "/")
	c.recordContentType(outputPath)

	// Create entry for the index
	c.entriesMutex.Lock()
	c.jsonEntries = append(c.jsonEntries, JSONFileEntry{
		Path: outputPath,
		URL:  u,
	})
	c.entriesMutex.Unlock()

	// Extract and enqueue new JSON URLs
	newURLs := ExtractJSONURLs(data)

	// If skipSymbols is enabled, filter out symbol-level documentation URLs
	if cfg.SkipSymbols {
		filteredURLs := make([]string, 0, len(newURLs))
		for _, url := range newURLs {
			if !IsSymbolURL(url) {
				filteredURLs = append(filteredURLs, url)
			} else {
				c.incrementSkippedSymbols()
				if cfg.Verbose {
					log.Printf("Skipping symbol URL: %s", url)
				}
			}
		}
		if cfg.Verbose && len(newURLs) != len(filteredURLs) {
			log.Printf("Filtered out %d symbol URLs", len(newURLs)-len(filteredURLs))
		}
		newURLs = filteredURLs
	}

	if cfg.Verbose {
		log.Printf("Found %d URLs in %s", len(newURLs), outputPath)
	}

	// Queue new URLs for processing
	newURLsAdded := c.queueNewURLs(newURLs, u, urlQueue, cfg)

	if cfg.Verbose {
		log.Printf("Queued %d new URLs from %s", newURLsAdded, outputPath)
	}

	return nil
}

// shouldSkipURL checks if a URL should be skipped based on its pattern
func (c *Crawler) shouldSkipURL(parsed *url.URL) bool {
	isMediaFile := strings.Contains(parsed.Path, "/media-")
	isMarketingPage := strings.HasSuffix(parsed.Path, "-hero") ||
		strings.HasPrefix(parsed.Path, "/devLink-") ||
		strings.HasPrefix(parsed.Path, "/link-") ||
		strings.Contains(parsed.Path, "-module") ||
		strings.Contains(parsed.Path, "-dynamic-") ||
		strings.Contains(parsed.Path, "#") ||
		strings.Contains(parsed.Path, "managedapp")

	// Skip standard library math functions that are re-exported (e.g., acos, sin, etc.)
	pathLower := strings.ToLower(parsed.Path)
	isStdLibMath := isStandardMathFunction(pathLower)

	// Skip C preprocessor macros and compiler directives
	isCMacro := isCPreprocessorMacro(pathLower)

	return isMediaFile ||
		isMarketingPage ||
		isStdLibMath ||
		isCMacro ||
		strings.Contains(parsed.Path, "/assets/") ||
		strings.HasSuffix(parsed.Path, ".css") ||
		strings.HasSuffix(parsed.Path, ".js") ||
		strings.HasSuffix(parsed.Path, ".png") ||
		strings.HasSuffix(parsed.Path, ".jpg") ||
		strings.HasSuffix(parsed.Path, ".svg") ||
		strings.HasSuffix(parsed.Path, ".pdf")
}

// isStandardMathFunction checks if the path represents a standard C/Swift math function
func isStandardMathFunction(pathLower string) bool {
	// List of standard math functions that frameworks sometimes re-export
	mathFuncs := []string{
		"/acos(", "/acosh(", "/asin(", "/asinh(", "/atan(", "/atan2(", "/atanh(",
		"/cbrt(", "/cos(", "/cosh(", "/erf(", "/erfc(", "/exp(", "/exp2(", "/expm1(",
		"/fdim(", "/fmax(", "/fmin(", "/hypot(", "/ilogb(", "/j0(", "/j1(", "/jn(",
		"/ldexp(", "/lgamma(", "/log(", "/log10(", "/log1p(", "/log2(", "/logb(",
		"/nan(", "/nearbyint(", "/nextafter(", "/pow(", "/remquo(", "/rint(",
		"/sin(", "/sinh(", "/tan(", "/tanh(", "/tgamma(", "/y0(", "/y1(", "/yn(",
		"/copysign(",
	}

	for _, fn := range mathFuncs {
		if strings.Contains(pathLower, fn) {
			return true
		}
	}
	return false
}

// isCPreprocessorMacro checks if the path represents a C preprocessor macro or compiler directive
func isCPreprocessorMacro(pathLower string) bool {
	// Common patterns for C macros and compiler directives
	macroPatterns := []string{
		"_extern", "_inline", "_local", "_deprecated", "_obsolete",
		"_soft_deprecated", "_boxable", "_bridge_", "_nonnull", "_nullable",
		"_version", "_hdr_", "_pure", "_extern_32", "_extern_64",
	}

	for _, pattern := range macroPatterns {
		if strings.Contains(pathLower, pattern) {
			return true
		}
	}

	// Also skip paths that look like macro definitions (all caps with underscores)
	parts := strings.Split(pathLower, "/")
	if len(parts) > 0 {
		lastPart := parts[len(parts)-1]
		// Check if it looks like a macro (contains underscores and mostly uppercase)
		if strings.Contains(lastPart, "_") {
			upperCount := 0
			for _, c := range lastPart {
				if c >= 'A' && c <= 'Z' {
					upperCount++
				}
			}
			// If more than 50% uppercase and has underscores, likely a macro
			if float64(upperCount)/float64(len(lastPart)) > 0.5 {
				return true
			}
		}
	}

	return false
}

// Utility functions for crawler logic follow...
// (Truncated for brevity - full implementation would include all helper methods)
