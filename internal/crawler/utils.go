package crawler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

// BuildFrameworkURLs constructs URLs for a specific framework
func BuildFrameworkURLs(frameworkName string) []string {
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

// ResolveURL resolves a potentially relative URL against the base URL
func ResolveURL(base, relative string) string {
	// Handle absolute URLs - convert old documentation URLs to new API format
	if strings.HasPrefix(relative, "http://") || strings.HasPrefix(relative, "https://") {
		// Convert old-style documentation URLs to new JSON API format
		// https://developer.apple.com/documentation/foundation/nsdata
		// -> https://developer.apple.com/tutorials/data/documentation/foundation/nsdata.json
		if strings.Contains(relative, "developer.apple.com/documentation/") &&
			!strings.Contains(relative, "/tutorials/data/") {
			// Extract the path after /documentation/
			parts := strings.SplitN(relative, "/documentation/", 2)
			if len(parts) == 2 {
				docPath := parts[1]
				// Add .json if not present
				if !strings.HasSuffix(docPath, ".json") {
					docPath = docPath + ".json"
				}
				return base + "/tutorials/data/documentation/" + docPath
			}
		}
		return relative
	}

	// Handle documentation paths
	if strings.HasPrefix(relative, "documentation/") {
		return base + "/tutorials/data/" + relative
	}

	if strings.HasPrefix(relative, "/") {
		return base + relative
	}
	return base + "/" + relative
}

// ResolveURLsWithLanguageVariants returns URLs for both language variants if enabled
func ResolveURLsWithLanguageVariants(base, relative string, fetchBoth bool) []string {
	baseURL := ResolveURL(base, relative)

	// Only add language variants for .json documentation URLs
	if !fetchBoth || !strings.HasSuffix(baseURL, ".json") || !strings.Contains(baseURL, "/tutorials/data/documentation/") {
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

// ExtractJSONURLs extracts URLs to other JSON files from a JSON response
func ExtractJSONURLs(data []byte) []string {
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

// extractJSONURLsFromValue recursively extracts all URLs from a JSON value
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

// IsSymbolURL determines if a URL likely points to individual symbol documentation
func IsSymbolURL(urlPath string) bool {
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

// ShouldExcludePath checks if a URL path should be excluded based on user-defined exclude patterns
func ShouldExcludePath(pathToCheck, excludePaths string, verbose bool) bool {
	if excludePaths == "" {
		return false
	}

	patterns := strings.Split(excludePaths, ",")
	for _, pattern := range patterns {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" {
			continue
		}

		// Check if the path contains this pattern
		if strings.Contains(pathToCheck, pattern) {
			if verbose {
				log.Printf("Excluding path %q because it matches pattern %q", pathToCheck, pattern)
			}
			return true
		}
	}

	return false
}

// IsJSON checks if the data is valid JSON
func IsJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}

// PrettyPrintJSON formats JSON data with proper indentation
func PrettyPrintJSON(data []byte) ([]byte, error) {
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

// WriteFileAtomic writes data to a file atomically by writing to a temp file first
func WriteFileAtomic(filename string, data []byte) error {
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

// AppendToBadURLsFile adds a URL to the bad URLs file
func AppendToBadURLsFile(url, badURLsFile string) {
	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(badURLsFile), 0755); err != nil {
		log.Printf("Error creating bad URLs directory: %v", err)
		return
	}

	// Create file if it doesn't exist
	file, err := os.OpenFile(badURLsFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

// AddBrowserLikeHeaders adds headers to a request to make it look like a browser request
func AddBrowserLikeHeaders(req *http.Request) {
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

// writeBadURLsFile writes all bad URLs to the bad URLs file (unexported method)
func writeBadURLsFileInternal(badURLs map[string]bool, badURLsFile string) error {
	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(badURLsFile), 0755); err != nil {
		return err
	}

	// Get a sorted list of bad URLs for more consistent file content
	urls := make([]string, 0, len(badURLs))
	for url := range badURLs {
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
	return os.WriteFile(badURLsFile, []byte(content.String()), 0644)
}
