package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// ValidationError represents a validation error with context
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
	Path    string
}

func (e ValidationError) Error() string {
	if e.Path != "" {
		return fmt.Sprintf("validation error in %s at %s: %s (value: %v)", e.Path, e.Field, e.Message, e.Value)
	}
	return fmt.Sprintf("validation error at %s: %s (value: %v)", e.Field, e.Message, e.Value)
}

// ValidationResult contains the results of validation
type ValidationResult struct {
	Valid   bool
	Errors  []ValidationError
	Warnings []ValidationError
}

// AddError adds a validation error
func (vr *ValidationResult) AddError(field, message string, value interface{}, path ...string) {
	pathStr := ""
	if len(path) > 0 {
		pathStr = path[0]
	}
	vr.Errors = append(vr.Errors, ValidationError{
		Field:   field,
		Value:   value,
		Message: message,
		Path:    pathStr,
	})
	vr.Valid = false
}

// AddWarning adds a validation warning
func (vr *ValidationResult) AddWarning(field, message string, value interface{}, path ...string) {
	pathStr := ""
	if len(path) > 0 {
		pathStr = path[0]
	}
	vr.Warnings = append(vr.Warnings, ValidationError{
		Field:   field,
		Value:   value,
		Message: message,
		Path:    pathStr,
	})
}

// ValidateCommandLineFlags validates all command-line flags for correctness
func ValidateCommandLineFlags() ValidationResult {
	result := ValidationResult{Valid: true}

	// Validate output directory
	if *outputDir == "" {
		result.AddError("output", "output directory cannot be empty", *outputDir)
	} else if !isValidPath(*outputDir) {
		result.AddError("output", "output directory path is invalid", *outputDir)
	}

	// Validate cache directory
	if *cacheDir == "" {
		result.AddError("cache", "cache directory cannot be empty", *cacheDir)
	} else if !isValidPath(*cacheDir) {
		result.AddError("cache", "cache directory path is invalid", *cacheDir)
	}

	// Validate base URL
	if *baseURL == "" {
		result.AddError("base", "base URL cannot be empty", *baseURL)
	} else if _, err := url.Parse(*baseURL); err != nil {
		result.AddError("base", "base URL is invalid", *baseURL)
	} else if !strings.HasPrefix(*baseURL, "http") {
		result.AddError("base", "base URL must start with http or https", *baseURL)
	}

	// Validate entry point
	if *entryPoint == "" {
		result.AddError("entry-point", "entry point cannot be empty", *entryPoint)
	}

	// Validate concurrency
	if *concurrency <= 0 {
		result.AddError("concurrency", "concurrency must be positive", *concurrency)
	} else if *concurrency > 100 {
		result.AddWarning("concurrency", "concurrency is very high, may overwhelm server", *concurrency)
	}

	// Validate rate limit
	if *rateLimit < 0 {
		result.AddError("rate-limit", "rate limit cannot be negative", *rateLimit)
	} else if *rateLimit > 1000 {
		result.AddWarning("rate-limit", "rate limit is very high", *rateLimit)
	}

	// Validate timeout
	if *timeout <= 0 {
		result.AddError("timeout", "timeout must be positive", *timeout)
	} else if *timeout < 1*time.Second {
		result.AddWarning("timeout", "timeout is very short", *timeout)
	} else if *timeout > 5*time.Minute {
		result.AddWarning("timeout", "timeout is very long", *timeout)
	}

	// Validate max time
	if *maxTime <= 0 {
		result.AddError("max-time", "max time must be positive", *maxTime)
	}

	// Validate mode
	validModes := map[string]bool{"crawl": true, "html": true, "markdown": true, "all": true}
	if !validModes[*mode] {
		result.AddError("mode", "invalid mode, must be one of: crawl, html, markdown, all", *mode)
	}

	// Validate log level
	validLogLevels := map[string]bool{"debug": true, "info": true, "warn": true, "error": true}
	if !validLogLevels[*logLevel] {
		result.AddError("log-level", "invalid log level, must be one of: debug, info, warn, error", *logLevel)
	}

	// Validate export metrics path if provided
	if *exportMetrics != "" && *exportMetrics != "true" && *exportMetrics != "1" {
		if !isValidPath(*exportMetrics) {
			result.AddError("export-metrics", "export metrics path is invalid", *exportMetrics)
		}
	}

	return result
}

// ValidateAppleDocJSON validates Apple documentation JSON structure and content
func ValidateAppleDocJSON(data []byte, sourcePath string) ValidationResult {
	result := ValidationResult{Valid: true}

	// Basic JSON validity check
	if !isJSON(data) {
		result.AddError("json", "invalid JSON format", string(data[:min(100, len(data))]), sourcePath)
		return result
	}

	// Parse into DocJSONData structure
	var doc DocJSONData
	if err := json.Unmarshal(data, &doc); err != nil {
		result.AddError("schema", "failed to parse as Apple documentation JSON", err.Error(), sourcePath)
		return result
	}

	// Validate document structure
	validateDocumentStructure(&doc, &result, sourcePath)

	return result
}

// ValidateDataIntegrity performs comprehensive data integrity checks
func ValidateDataIntegrity(filePath string, data []byte) ValidationResult {
	result := ValidationResult{Valid: true}

	// Check file size
	if len(data) == 0 {
		result.AddError("size", "file is empty", len(data), filePath)
		return result
	}

	// Check for extremely large files that might indicate corruption
	if len(data) > 50*1024*1024 { // 50MB
		result.AddWarning("size", "file is very large, might indicate corruption", len(data), filePath)
	}

	// Check for binary data in what should be text
	if containsBinaryData(data) {
		result.AddError("encoding", "file contains binary data", "binary data detected", filePath)
	}

	// Validate JSON structure if it's a JSON file
	if strings.HasSuffix(filePath, ".json") {
		jsonResult := ValidateAppleDocJSON(data, filePath)
		result.Errors = append(result.Errors, jsonResult.Errors...)
		result.Warnings = append(result.Warnings, jsonResult.Warnings...)
		if !jsonResult.Valid {
			result.Valid = false
		}
	}

	return result
}

// validateDocumentStructure validates the structure of a DocJSONData document
func validateDocumentStructure(doc *DocJSONData, result *ValidationResult, sourcePath string) {
	// Validate metadata
	if doc.Metadata.Title == "" {
		result.AddWarning("metadata.title", "document has no title", "", sourcePath)
	}

	if doc.Metadata.Role == "" {
		result.AddWarning("metadata.role", "document has no role specified", "", sourcePath)
	}

	// Validate known roles
	validRoles := map[string]bool{
		"framework": true, "class": true, "protocol": true, "structure": true,
		"enumeration": true, "symbol": true, "method": true, "property": true,
		"function": true, "variable": true, "typeAlias": true, "module": true,
	}
	if doc.Metadata.Role != "" && !validRoles[doc.Metadata.Role] {
		result.AddWarning("metadata.role", "unknown role type", doc.Metadata.Role, sourcePath)
	}

	// Validate platforms
	for i, platform := range doc.Metadata.Platforms {
		if platform.Name == "" {
			result.AddError(fmt.Sprintf("metadata.platforms[%d].name", i), "platform name is required", "", sourcePath)
		}

		// Validate platform names
		validPlatforms := map[string]bool{
			"iOS": true, "macOS": true, "watchOS": true, "tvOS": true, "visionOS": true,
			"Mac Catalyst": true, "iPadOS": true, "DriverKit": true,
		}
		if platform.Name != "" && !validPlatforms[platform.Name] {
			result.AddWarning(fmt.Sprintf("metadata.platforms[%d].name", i), "unknown platform", platform.Name, sourcePath)
		}

		// Validate version format
		if platform.IntroducedAt != "" && !isValidVersionString(platform.IntroducedAt) {
			result.AddWarning(fmt.Sprintf("metadata.platforms[%d].introducedAt", i), "invalid version format", platform.IntroducedAt, sourcePath)
		}
	}

	// Validate references
	for refID, ref := range doc.References {
		if ref.Title == "" {
			result.AddWarning(fmt.Sprintf("references[%s].title", refID), "reference has no title", "", sourcePath)
		}

		if ref.URL == "" {
			result.AddWarning(fmt.Sprintf("references[%s].url", refID), "reference has no URL", "", sourcePath)
		}

		// Validate reference URL format
		if ref.URL != "" && !isValidReferenceURL(ref.URL) {
			result.AddWarning(fmt.Sprintf("references[%s].url", refID), "reference URL format is unusual", ref.URL, sourcePath)
		}
	}

	// Validate content sections
	for i, section := range doc.PrimaryContentSections {
		if section.Kind == "" {
			result.AddWarning(fmt.Sprintf("primaryContentSections[%d].kind", i), "content section has no kind", "", sourcePath)
		}

		// Validate known section kinds
		validKinds := map[string]bool{
			"content": true, "declarations": true, "parameters": true, "returns": true,
			"discussion": true, "overview": true, "availability": true,
		}
		if section.Kind != "" && !validKinds[section.Kind] {
			result.AddWarning(fmt.Sprintf("primaryContentSections[%d].kind", i), "unknown section kind", section.Kind, sourcePath)
		}
	}

	// Validate topic sections
	for i, topic := range doc.TopicSections {
		if topic.Title == "" {
			result.AddWarning(fmt.Sprintf("topicSections[%d].title", i), "topic section has no title", "", sourcePath)
		}

		if len(topic.Identifiers) == 0 {
			result.AddWarning(fmt.Sprintf("topicSections[%d].identifiers", i), "topic section has no identifiers", "", sourcePath)
		}
	}
}

// Helper functions

func isValidPath(path string) bool {
	if path == "" {
		return false
	}
	// Check for obviously invalid characters
	invalidChars := []string{"\x00", "<", ">", "|", "\""}
	for _, char := range invalidChars {
		if strings.Contains(path, char) {
			return false
		}
	}
	return true
}

func containsBinaryData(data []byte) bool {
	// Check first 1KB for binary data
	checkLen := min(1024, len(data))
	nullCount := 0
	for i := 0; i < checkLen; i++ {
		if data[i] == 0 {
			nullCount++
		}
		// If we find too many null bytes, it's likely binary
		if nullCount > 5 {
			return true
		}
	}
	return false
}

func isValidVersionString(version string) bool {
	// Match patterns like "13.0", "10.15.4", "1.0", etc.
	matched, _ := regexp.MatchString(`^\d+\.\d+(\.\d+)?$`, version)
	return matched
}

func isValidReferenceURL(url string) bool {
	// Apple doc URLs should match certain patterns
	return strings.HasPrefix(url, "doc://") || 
		   strings.HasPrefix(url, "https://") ||
		   strings.HasPrefix(url, "/documentation/")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ValidateCache validates the integrity of the cache directory
func ValidateCache(cacheDir string) ValidationResult {
	result := ValidationResult{Valid: true}

	// Check if cache directory exists
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		result.AddWarning("cache", "cache directory does not exist", cacheDir)
		return result
	}

	// Walk cache directory and validate files
	err := filepath.Walk(cacheDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			result.AddError("cache", "error accessing cache file", err.Error(), path)
			return nil // Continue walking
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check for empty files
		if info.Size() == 0 {
			result.AddWarning("cache", "empty cache file", "", path)
			return nil
		}

		// Check for very old cache files (older than 30 days)
		if time.Since(info.ModTime()) > 30*24*time.Hour {
			result.AddWarning("cache", "very old cache file", info.ModTime(), path)
		}

		return nil
	})

	if err != nil {
		result.AddError("cache", "failed to walk cache directory", err.Error(), cacheDir)
	}

	return result
}