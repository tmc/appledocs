package crawler

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// fetchWithCache fetches a URL with caching and ETag support
func (c *Crawler) fetchWithCache(ctx context.Context, u string, cfg *Config) ([]byte, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		c.incrementErrors()
		return nil, fmt.Errorf("parse URL %q: %v", u, err)
	}

	// Create cache path
	basePath := filepath.Join(cfg.CacheDir, parsed.Host, parsed.Path)

	// For documentation URLs, always use directory structure with index.json
	// This prevents file/directory collisions when a path has both content and children
	// e.g., /documentation/foundation/nsstring (has content)
	//   and /documentation/foundation/nsstring/init (child path)
	var cachePath string
	var cacheDir string

	if strings.Contains(parsed.Path, "/documentation/") || strings.HasSuffix(parsed.Path, ".json") {
		// Use directory structure: path/index.json
		cacheDir = basePath
		// For .json URLs, strip the extension for the directory name
		if strings.HasSuffix(cacheDir, ".json") {
			cacheDir = strings.TrimSuffix(cacheDir, ".json")
		}
		cachePath = filepath.Join(cacheDir, "index.json")
	} else {
		// For non-documentation URLs, use the path as filename
		cacheDir = filepath.Dir(basePath)
		cachePath = basePath
	}

	// Add query params to filename if they're NOT language params
	if parsed.RawQuery != "" {
		query := parsed.Query()
		query.Del("language")
		if len(query) > 0 {
			queryStr := url.QueryEscape(query.Encode())
			if strings.HasSuffix(cachePath, "index.json") {
				// For index.json, insert query before extension
				cachePath = filepath.Join(cacheDir, "index."+queryStr+".json")
			} else {
				cachePath = cachePath + "." + queryStr
			}
		}
	}

	// Ensure directory exists
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		c.incrementErrors()
		return nil, fmt.Errorf("create cache directory: %v", err)
	}

	etagPath := cachePath + ".etag"

	// Check cache
	if !cfg.ForceRefresh {
		if fi, err := os.Stat(cachePath); err == nil && fi.Size() > 0 {
			data, err := os.ReadFile(cachePath)
			if err == nil {
				// For JSON files, verify that the content is valid JSON
				if strings.HasSuffix(parsed.Path, ".json") && !IsJSON(data) {
					log.Printf("Cache contains invalid JSON for %q, refetching", u)
					c.incrementCacheMisses()
				} else {
					// Check if we have an ETag to validate freshness
					if etagData, err := os.ReadFile(etagPath); err == nil && len(etagData) > 0 {
						// We'll validate with If-None-Match below
					} else {
						// No ETag, use cache as-is
						c.incrementCacheHits()
						c.recordBytesFromCache(int64(len(data)))
						if cfg.Verbose {
							log.Printf("Cache hit for %q (no ETag)", u)
						}
						return data, nil
					}
				}
			} else {
				log.Printf("Cache read error for %q: %v", u, err)
				c.incrementErrors()
				c.incrementCacheMisses()
			}
		} else {
			c.incrementCacheMisses()
		}
	} else {
		c.incrementCacheMisses()
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, "GET", u, nil)
	if err != nil {
		c.incrementErrors()
		return nil, fmt.Errorf("create request for %q: %v", u, err)
	}

	// Add browser-like headers
	AddBrowserLikeHeaders(req)

	// Add If-None-Match header if we have an ETag
	if etagData, err := os.ReadFile(etagPath); err == nil && len(etagData) > 0 {
		etag := strings.TrimSpace(string(etagData))
		req.Header.Set("If-None-Match", etag)
		if cfg.Verbose {
			log.Printf("Sending If-None-Match: %s for %q", etag, u)
		}
	}

	// Execute the request with retry logic
	var resp *http.Response
	maxRetries := 3
	for attempt := 0; attempt <= maxRetries; attempt++ {
		// Apply rate limiting before making request
		if err := c.rateLimiter.Wait(ctx); err != nil {
			return nil, fmt.Errorf("rate limit wait failed: %v", err)
		}

		startTime := time.Now()
		resp, err = c.client.Do(req)
		requestDuration := time.Since(startTime)

		// Record response time for successful requests
		if err == nil {
			c.recordResponseTime(requestDuration)
		}

		// Check if we should retry
		shouldRetry, retryReason := c.shouldRetryRequest(err, resp, attempt, maxRetries, u)
		if !shouldRetry {
			if err != nil {
				return nil, err
			}
			break
		}

		if resp != nil {
			resp.Body.Close()
		}

		// Handle retry backoff
		backoffDuration := time.Duration(1<<uint(attempt)) * time.Second
		if cfg.Verbose {
			log.Printf("Attempt %d failed for %q (%s), retrying in %v", attempt+1, u, retryReason, backoffDuration)
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(backoffDuration):
			// Continue to next attempt
		}
	}

	// Handle 304 Not Modified
	if resp.StatusCode == http.StatusNotModified {
		resp.Body.Close()
		if cfg.Verbose {
			log.Printf("304 Not Modified for %q, using cached content", u)
		}

		cachedData, err := os.ReadFile(cachePath)
		if err != nil {
			c.incrementErrors()
			return nil, fmt.Errorf("read cached content after 304 for %q: %v", u, err)
		}

		c.incrementCacheHits()
		c.recordBytesFromCache(int64(len(cachedData)))
		return cachedData, nil
	}

	defer resp.Body.Close()

	// Read response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		c.incrementErrors()
		return nil, fmt.Errorf("read %q: %v", u, err)
	}

	// Record bytes downloaded
	c.recordBytesDownloaded(int64(len(data)))

	// Write to cache with atomic operation
	if err := WriteFileAtomic(cachePath, data); err != nil {
		c.incrementErrors()
		return nil, fmt.Errorf("write cache %q: %v", cachePath, err)
	}

	// Store ETag if present
	if etag := resp.Header.Get("Etag"); etag != "" {
		if err := os.WriteFile(etagPath, []byte(etag), 0644); err != nil {
			if cfg.Verbose {
				log.Printf("Warning: Failed to write ETag for %q: %v", u, err)
			}
		} else if cfg.Verbose {
			log.Printf("Stored ETag %s for %q", etag, u)
		}
	}

	return data, nil
}

// shouldRetryRequest determines if a request should be retried
func (c *Crawler) shouldRetryRequest(err error, resp *http.Response, attempt, maxRetries int, u string) (bool, string) {
	if err != nil {
		// Check if it's a retryable network error
		isRetryable := strings.Contains(err.Error(), "timeout") ||
			strings.Contains(err.Error(), "connection reset") ||
			strings.Contains(err.Error(), "temporary failure")

		if isRetryable && attempt < maxRetries {
			c.incrementRetryCount()
			return true, fmt.Sprintf("network error: %v", err)
		}

		c.incrementErrors()
		// Add to bad URLs list if it's a permanent network error
		if strings.Contains(err.Error(), "no such host") ||
			strings.Contains(err.Error(), "connection refused") ||
			(!isRetryable && strings.Contains(err.Error(), "timeout")) {
			c.badURLs[u] = true
			AppendToBadURLsFile(u, "")
		}
		return false, ""
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotModified {
		// Record HTTP error
		c.recordHTTPError(resp.StatusCode)

		// Check if it's a retryable status code
		isRetryableStatus := resp.StatusCode == http.StatusTooManyRequests ||
			resp.StatusCode == http.StatusInternalServerError ||
			resp.StatusCode == http.StatusBadGateway ||
			resp.StatusCode == http.StatusServiceUnavailable ||
			resp.StatusCode == http.StatusGatewayTimeout

		if isRetryableStatus && attempt < maxRetries {
			c.incrementRetryCount()
			return true, fmt.Sprintf("HTTP %d: %s", resp.StatusCode, resp.Status)
		}

		c.incrementErrors()

		// Add to bad URLs list for client errors
		if resp.StatusCode == http.StatusForbidden ||
			resp.StatusCode == http.StatusNotFound ||
			resp.StatusCode == http.StatusMethodNotAllowed ||
			resp.StatusCode == http.StatusGone {
			c.badURLs[u] = true
			AppendToBadURLsFile(u, "")
		}
		return false, ""
	}

	return false, ""
}

// saveToOutputDir saves content to the output directory
func (c *Crawler) saveToOutputDir(path string, content []byte, cfg *Config) error {
	outputPath := filepath.Join(cfg.OutputDir, path)
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("create output directory: %v", err)
	}

	// Pretty-print JSON files if enabled
	if cfg.PrettyJSON && strings.HasSuffix(path, ".json") {
		if prettyContent, err := PrettyPrintJSON(content); err == nil {
			content = prettyContent
		} else {
			log.Printf("Failed to pretty-print JSON file %s: %v", path, err)
		}
	}

	return WriteFileAtomic(outputPath, content)
}

// queueNewURLs adds new URLs to the processing queue if they haven't been visited
func (c *Crawler) queueNewURLs(newURLs []string, parentURL string, urlQueue chan<- string, cfg *Config) int {
	var added int

	// Get parent depth
	c.depthMutex.RLock()
	parentDepth := c.urlDepths[parentURL]
	c.depthMutex.RUnlock()

	// Calculate child depth
	childDepth := parentDepth + 1

	// If max depth is set and we've exceeded it, don't queue any children
	if cfg.MaxDepth > 0 && childDepth > cfg.MaxDepth {
		if cfg.Verbose {
			log.Printf("Max depth %d reached, skipping %d URLs from %s", cfg.MaxDepth, len(newURLs), parentURL)
		}
		return 0
	}

	for _, newURL := range newURLs {
		// Get all language variants for this URL
		resolvedURLs := ResolveURLsWithLanguageVariants(cfg.BaseURL, newURL, cfg.FetchBothLanguages)

		for _, resolvedURL := range resolvedURLs {
			// Check if URL is known to be bad
			if c.badURLs[resolvedURL] {
				if cfg.Verbose {
					log.Printf("Skipping known bad URL: %s", resolvedURL)
				}
				continue
			}

			// Parse URL to check if it should be excluded
			parsedURL, err := url.Parse(resolvedURL)
			if err == nil && ShouldExcludePath(parsedURL.Path, cfg.ExcludePaths, cfg.Verbose) {
				c.badURLs[resolvedURL] = true
				AppendToBadURLsFile(resolvedURL, cfg.BadURLsFile)
				if cfg.Verbose {
					log.Printf("Skipping excluded path when queueing: %s", parsedURL.Path)
				}
				continue
			}

			// Check if URL is within entry point scope (only crawl children of entry point)
			if c.entryPointPrefix != "" && parsedURL != nil {
				urlPath := strings.ToLower(parsedURL.Path)
				if !strings.HasPrefix(urlPath, c.entryPointPrefix) {
					if cfg.Verbose {
						log.Printf("Skipping URL outside entry point scope: %s (prefix: %s)", parsedURL.Path, c.entryPointPrefix)
					}
					continue
				}
			}

			// Check if URL has been visited using atomic LoadOrStore
			_, alreadyVisited := c.visitedURLs.LoadOrStore(resolvedURL, true)
			if alreadyVisited {
				continue
			}

			// Set depth for this URL
			c.depthMutex.Lock()
			c.urlDepths[resolvedURL] = childDepth
			c.depthMutex.Unlock()

			// Try to send to channel
			select {
			case urlQueue <- resolvedURL:
				added++
			default:
				// Channel might be full or closed
				if cfg.Verbose {
					log.Printf("Skipping URL %s (channel full or closed)", resolvedURL)
				}
			}
		}
	}

	return added
}

// loadBadURLs loads the list of known bad URLs from a file
func (c *Crawler) loadBadURLs(badURLsFile string, verbose bool) error {
	if _, err := os.Stat(badURLsFile); os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	content, err := os.ReadFile(badURLsFile)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		c.badURLs[line] = true
		if verbose {
			log.Printf("Added bad URL: %s", line)
		}
	}

	if verbose {
		log.Printf("Loaded %d known bad URLs", len(c.badURLs))
	}
	return nil
}

// writeBadURLsFile writes all bad URLs to the bad URLs file
func (c *Crawler) writeBadURLsFile(badURLsFile string) error {
	return writeBadURLsFileInternal(c.badURLs, badURLsFile)
}

// Metrics methods
func (c *Crawler) incrementCacheHits() {
	c.statsMutex.Lock()
	c.cacheHits++
	c.statsMutex.Unlock()
}

func (c *Crawler) incrementCacheMisses() {
	c.statsMutex.Lock()
	c.cacheMisses++
	c.statsMutex.Unlock()
}

func (c *Crawler) incrementErrors() {
	c.statsMutex.Lock()
	c.errors++
	c.statsMutex.Unlock()
}

func (c *Crawler) incrementSkippedURLs() {
	c.statsMutex.Lock()
	c.skippedURLs++
	c.statsMutex.Unlock()
}

func (c *Crawler) incrementSkippedSymbols() {
	c.statsMutex.Lock()
	c.skippedSymbols++
	c.skippedURLs++
	c.statsMutex.Unlock()
}

func (c *Crawler) recordResponseTime(duration time.Duration) {
	c.statsMutex.Lock()
	c.totalResponseTime += duration
	c.requestCount++
	if c.requestCount > 0 {
		c.avgResponseTime = c.totalResponseTime / time.Duration(c.requestCount)
	}
	c.statsMutex.Unlock()
}

func (c *Crawler) recordBytesDownloaded(bytes int64) {
	c.statsMutex.Lock()
	c.totalBytesDownloaded += bytes
	c.statsMutex.Unlock()
}

func (c *Crawler) recordBytesFromCache(bytes int64) {
	c.statsMutex.Lock()
	c.totalBytesFromCache += bytes
	c.statsMutex.Unlock()
}

func (c *Crawler) recordHTTPError(statusCode int) {
	c.statsMutex.Lock()
	if c.httpErrors == nil {
		c.httpErrors = make(map[int]int)
	}
	c.httpErrors[statusCode]++
	c.statsMutex.Unlock()
}

func (c *Crawler) incrementRetryCount() {
	c.statsMutex.Lock()
	c.retryCount++
	c.statsMutex.Unlock()
}

func (c *Crawler) recordContentType(path string) {
	c.statsMutex.Lock()
	defer c.statsMutex.Unlock()

	if strings.Contains(path, "/documentation/") {
		pathParts := strings.Split(path, "/")
		// Count frameworks
		if len(pathParts) >= 4 && strings.HasSuffix(pathParts[3], ".json") && !strings.Contains(pathParts[3], "/") {
			c.frameworkCount++
		}
		// Count classes
		if len(pathParts) >= 5 && strings.HasSuffix(pathParts[4], ".json") {
			c.classCount++
		}
		// Count methods
		if len(pathParts) >= 6 {
			c.methodCount++
		}
	}
}
