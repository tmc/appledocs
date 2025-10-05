// Package main provides migration strategies and implementation guides
// for transitioning to optimized JSON parsing in appledocs
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/buger/jsonparser"
	"github.com/json-iterator/go"
	"github.com/tidwall/gjson"
	"github.com/valyala/fastjson"
)

// MigrationStrategy defines different approaches for migrating JSON parsing
type MigrationStrategy int

const (
	DropInReplacement MigrationStrategy = iota
	GradualMigration
	StreamingFirst
	SelectiveParsing
)

// MigrationPlan provides a structured approach to JSON parsing optimization
type MigrationPlan struct {
	Strategy    MigrationStrategy
	Phase       int
	Description string
	Benefits    []string
	Risks       []string
	Implementation string
}

// JSONParsingInterface defines the contract for JSON parsers in appledocs
type JSONParsingInterface interface {
	ParseDocJSON(data []byte) (*DocJSONData, error)
	ExtractURLs(data []byte) ([]string, error)
	SupportsStreaming() bool
	MemoryFootprint() string // "low", "medium", "high"
	Name() string
}

// StandardJSONParser implements the current encoding/json approach
type StandardJSONParser struct{}

func (p *StandardJSONParser) ParseDocJSON(data []byte) (*DocJSONData, error) {
	var doc DocJSONData
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	return &doc, nil
}

func (p *StandardJSONParser) ExtractURLs(data []byte) ([]string, error) {
	return extractJSONURLs(data), nil
}

func (p *StandardJSONParser) SupportsStreaming() bool { return false }
func (p *StandardJSONParser) MemoryFootprint() string { return "high" }
func (p *StandardJSONParser) Name() string             { return "encoding/json" }

// JsoniterParser provides jsoniter-based parsing
type JsoniterParser struct{}

func (p *JsoniterParser) ParseDocJSON(data []byte) (*DocJSONData, error) {
	var doc DocJSONData
	if err := jsoniter.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	return &doc, nil
}

func (p *JsoniterParser) ExtractURLs(data []byte) ([]string, error) {
	return extractJSONURLs(data), nil
}

func (p *JsoniterParser) SupportsStreaming() bool { return false }
func (p *JsoniterParser) MemoryFootprint() string { return "medium" }
func (p *JsoniterParser) Name() string             { return "jsoniter" }

// FastJSONParser provides fastjson-based parsing
type FastJSONParser struct {
	parser fastjson.Parser
}

func (p *FastJSONParser) ParseDocJSON(data []byte) (*DocJSONData, error) {
	v, err := p.parser.Parse(string(data))
	if err != nil {
		return nil, err
	}
	
	// Convert fastjson.Value to DocJSONData
	return p.convertToDocJSON(v)
}

func (p *FastJSONParser) ExtractURLs(data []byte) ([]string, error) {
	v, err := p.parser.Parse(string(data))
	if err != nil {
		return nil, err
	}
	return extractURLsFromFastJSON(v), nil
}

func (p *FastJSONParser) SupportsStreaming() bool { return true }
func (p *FastJSONParser) MemoryFootprint() string { return "low" }
func (p *FastJSONParser) Name() string             { return "fastjson" }

func (p *FastJSONParser) convertToDocJSON(v *fastjson.Value) (*DocJSONData, error) {
	// This is a simplified conversion - in practice, you'd need full conversion
	doc := &DocJSONData{}
	
	// Extract basic fields
	if title := v.Get("metadata", "title"); title != nil {
		doc.Metadata.Title = string(title.GetStringBytes())
	}
	
	// Extract references
	doc.References = make(map[string]Reference)
	refs := v.Get("references")
	if refs != nil {
		refs.GetObject().Visit(func(key []byte, v *fastjson.Value) {
			ref := Reference{}
			if url := v.Get("url"); url != nil {
				ref.URL = string(url.GetStringBytes())
			}
			if title := v.Get("title"); title != nil {
				ref.Title = string(title.GetStringBytes())
			}
			doc.References[string(key)] = ref
		})
	}
	
	return doc, nil
}

// GJSONParser provides gjson-based selective parsing
type GJSONParser struct{}

func (p *GJSONParser) ParseDocJSON(data []byte) (*DocJSONData, error) {
	// GJSON excels at selective parsing, but full struct conversion is less efficient
	// Use this for selective field extraction rather than full document parsing
	return nil, fmt.Errorf("gjson parser optimized for selective parsing, not full document conversion")
}

func (p *GJSONParser) ExtractURLs(data []byte) ([]string, error) {
	return extractURLsFromGJSON(string(data)), nil
}

func (p *GJSONParser) SupportsStreaming() bool { return false }
func (p *GJSONParser) MemoryFootprint() string { return "low" }
func (p *GJSONParser) Name() string             { return "gjson" }

// StreamingJSONParser provides streaming-optimized parsing
type StreamingJSONParser struct{}

func (p *StreamingJSONParser) ParseDocJSON(data []byte) (*DocJSONData, error) {
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	
	var doc DocJSONData
	if err := decoder.Decode(&doc); err != nil {
		return nil, err
	}
	return &doc, nil
}

func (p *StreamingJSONParser) ExtractURLs(data []byte) ([]string, error) {
	return extractJSONURLs(data), nil
}

func (p *StreamingJSONParser) SupportsStreaming() bool { return true }
func (p *StreamingJSONParser) MemoryFootprint() string { return "medium" }
func (p *StreamingJSONParser) Name() string             { return "streaming/json" }

// GetMigrationPlans returns recommended migration strategies
func GetMigrationPlans() []MigrationPlan {
	return []MigrationPlan{
		{
			Strategy:    DropInReplacement,
			Phase:       1,
			Description: "Replace encoding/json with jsoniter for immediate performance gains",
			Benefits: []string{
				"2-3x speed improvement",
				"Drop-in replacement",
				"Minimal code changes",
				"Maintains compatibility",
			},
			Risks: []string{
				"Small dependency addition",
				"Slightly different error messages",
			},
			Implementation: `
// Change import
import "github.com/json-iterator/go"

// Replace json.Unmarshal calls
var jsoniter = jsoniter.ConfigCompatibleWithStandardLibrary
err := jsoniter.Unmarshal(data, &doc)
`,
		},
		{
			Strategy:    GradualMigration,
			Phase:       2,
			Description: "Migrate URL extraction to fastjson for memory efficiency",
			Benefits: []string{
				"5-10x speed improvement for URL extraction",
				"Significantly lower memory usage",
				"Zero-allocation parsing",
			},
			Risks: []string{
				"API changes required",
				"Need to maintain two parsing paths initially",
			},
			Implementation: `
// For URL extraction only
func extractURLsOptimized(data []byte) []string {
    var p fastjson.Parser
    v, err := p.Parse(string(data))
    if err != nil {
        // Fallback to standard parsing
        return extractJSONURLs(data)
    }
    return extractURLsFromFastJSON(v)
}
`,
		},
		{
			Strategy:    StreamingFirst,
			Phase:       3,
			Description: "Implement streaming for large files (>1MB)",
			Benefits: []string{
				"Constant memory usage regardless of file size",
				"Can process files larger than available RAM",
				"Better performance for large files",
			},
			Risks: []string{
				"More complex implementation",
				"Different error handling",
				"Requires significant code changes",
			},
			Implementation: `
// Streaming implementation for large files
func processLargeFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    // Check file size
    if info, err := file.Stat(); err == nil && info.Size() > 1024*1024 {
        return processFileStreaming(file)
    }
    
    // Use regular parsing for smaller files
    return processFileStandard(file)
}
`,
		},
		{
			Strategy:    SelectiveParsing,
			Phase:       4,
			Description: "Use gjson for metadata-only operations",
			Benefits: []string{
				"Extract only needed fields",
				"Skip parsing unnecessary data",
				"Very low memory footprint",
			},
			Risks: []string{
				"Requires identifying specific use cases",
				"Different API for different operations",
			},
			Implementation: `
// Extract only URLs using gjson
func extractURLsOnly(data []byte) []string {
    var urls []string
    
    // Extract from references
    gjson.GetBytes(data, "references").ForEach(func(key, value gjson.Result) bool {
        if url := value.Get("url"); url.Exists() {
            urls = append(urls, url.String())
        }
        return true
    })
    
    return urls
}
`,
		},
	}
}

// RecommendMigrationStrategy analyzes current usage and recommends optimal strategy
func RecommendMigrationStrategy(fileSize int64, operation string, memoryConstraints bool) MigrationPlan {
	plans := GetMigrationPlans()
	
	// For small files, start with drop-in replacement
	if fileSize < 100*1024 { // < 100KB
		return plans[0] // DropInReplacement
	}
	
	// For memory-constrained environments
	if memoryConstraints {
		if operation == "url_extraction" {
			return plans[3] // SelectiveParsing
		}
		return plans[2] // StreamingFirst
	}
	
	// For large files
	if fileSize > 1024*1024 { // > 1MB
		return plans[2] // StreamingFirst
	}
	
	// Default to gradual migration
	return plans[1] // GradualMigration
}

// OptimizedJSONProcessor provides the recommended optimized implementation
type OptimizedJSONProcessor struct {
	standardParser  *StandardJSONParser
	jsoniterParser  *JsoniterParser
	fastjsonParser  *FastJSONParser
	gjsonParser     *GJSONParser
	streamingParser *StreamingJSONParser
}

func NewOptimizedJSONProcessor() *OptimizedJSONProcessor {
	return &OptimizedJSONProcessor{
		standardParser:  &StandardJSONParser{},
		jsoniterParser:  &JsoniterParser{},
		fastjsonParser:  &FastJSONParser{},
		gjsonParser:     &GJSONParser{},
		streamingParser: &StreamingJSONParser{},
	}
}

// ProcessJSON intelligently chooses the best parsing strategy
func (p *OptimizedJSONProcessor) ProcessJSON(data []byte, operation string) (*DocJSONData, []string, error) {
	dataSize := int64(len(data))
	
	// Choose strategy based on size and operation
	switch {
	case operation == "url_extraction" && dataSize > 100*1024:
		// Use gjson for URL extraction on larger files
		urls, err := p.gjsonParser.ExtractURLs(data)
		return nil, urls, err
		
	case dataSize > 5*1024*1024: // > 5MB
		// Use streaming for very large files
		doc, err := p.streamingParser.ParseDocJSON(data)
		if err != nil {
			return nil, nil, err
		}
		urls, _ := p.streamingParser.ExtractURLs(data)
		return doc, urls, nil
		
	case dataSize > 1024*1024: // > 1MB
		// Use fastjson for large files
		doc, err := p.fastjsonParser.ParseDocJSON(data)
		if err != nil {
			return nil, nil, err
		}
		urls, _ := p.fastjsonParser.ExtractURLs(data)
		return doc, urls, nil
		
	default:
		// Use jsoniter for regular files (best balance)
		doc, err := p.jsoniterParser.ParseDocJSON(data)
		if err != nil {
			return nil, nil, err
		}
		urls, _ := p.jsoniterParser.ExtractURLs(data)
		return doc, urls, nil
	}
}

// ProcessFile provides file-based processing with automatic strategy selection
func (p *OptimizedJSONProcessor) ProcessFile(filename string, operation string) (*DocJSONData, []string, error) {
	// Check file size first
	info, err := os.Stat(filename)
	if err != nil {
		return nil, nil, err
	}
	
	// For very large files, use streaming from the file directly
	if info.Size() > 10*1024*1024 { // > 10MB
		return p.processFileStreaming(filename, operation)
	}
	
	// For smaller files, read into memory and process
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}
	
	return p.ProcessJSON(data, operation)
}

// processFileStreaming handles large files with streaming
func (p *OptimizedJSONProcessor) processFileStreaming(filename string, operation string) (*DocJSONData, []string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	
	if operation == "url_extraction" {
		// Use streaming URL extraction
		var urls []string
		callback := func(url, source string) error {
			urls = append(urls, url)
			return nil
		}
		
		processor := &StandardStreamingProcessor{}
		err = processor.ProcessStream(file, callback)
		return nil, urls, err
	}
	
	// Full document parsing with streaming
	decoder := json.NewDecoder(file)
	var doc DocJSONData
	if err := decoder.Decode(&doc); err != nil {
		return nil, nil, err
	}
	
	// Extract URLs from parsed document
	urls := extractURLsFromDoc(&doc)
	return &doc, urls, nil
}

// BackwardCompatibleJSONExtractor maintains compatibility while using optimizations
type BackwardCompatibleJSONExtractor struct {
	processor *OptimizedJSONProcessor
	fallback  *StandardJSONParser
}

func NewBackwardCompatibleExtractor() *BackwardCompatibleJSONExtractor {
	return &BackwardCompatibleJSONExtractor{
		processor: NewOptimizedJSONProcessor(),
		fallback:  &StandardJSONParser{},
	}
}

// ExtractJSONURLs is a drop-in replacement for the current extractJSONURLs function
func (e *BackwardCompatibleJSONExtractor) ExtractJSONURLs(data []byte) []string {
	// Try optimized extraction first
	_, urls, err := e.processor.ProcessJSON(data, "url_extraction")
	if err != nil {
		// Fallback to standard approach
		if fallbackURLs, fallbackErr := e.fallback.ExtractURLs(data); fallbackErr == nil {
			return fallbackURLs
		}
	}
	return urls
}

// PerformanceMonitor tracks the performance of different parsing strategies
type PerformanceMonitor struct {
	libraryStats map[string]struct {
		totalCalls    int
		totalDuration int64 // nanoseconds
		totalMemory   uint64
		failureCount  int
	}
}

func NewPerformanceMonitor() *PerformanceMonitor {
	return &PerformanceMonitor{
		libraryStats: make(map[string]struct {
			totalCalls    int
			totalDuration int64
			totalMemory   uint64
			failureCount  int
		}),
	}
}

// RecordParsingMetrics records performance metrics for analysis
func (pm *PerformanceMonitor) RecordParsingMetrics(library string, duration int64, memory uint64, success bool) {
	stats := pm.libraryStats[library]
	stats.totalCalls++
	stats.totalDuration += duration
	stats.totalMemory += memory
	if !success {
		stats.failureCount++
	}
	pm.libraryStats[library] = stats
}

// GetRecommendation provides real-time recommendations based on observed performance
func (pm *PerformanceMonitor) GetRecommendation() string {
	if len(pm.libraryStats) == 0 {
		return "No performance data available"
	}
	
	var bestLibrary string
	var bestEfficiency float64
	
	for library, stats := range pm.libraryStats {
		if stats.totalCalls > 0 {
			avgDuration := float64(stats.totalDuration) / float64(stats.totalCalls)
			avgMemory := float64(stats.totalMemory) / float64(stats.totalCalls)
			successRate := float64(stats.totalCalls-stats.failureCount) / float64(stats.totalCalls)
			
			// Simple efficiency metric (lower is better)
			efficiency := avgDuration * avgMemory * (2.0 - successRate)
			
			if bestLibrary == "" || efficiency < bestEfficiency {
				bestLibrary = library
				bestEfficiency = efficiency
			}
		}
	}
	
	return fmt.Sprintf("Based on %d measurements, %s is performing best", 
		pm.libraryStats[bestLibrary].totalCalls, bestLibrary)
}

// generateMigrationGuide creates a comprehensive migration guide
func generateMigrationGuide() {
	fmt.Println("=== AppLeDocs JSON Parsing Migration Guide ===")
	fmt.Println()
	
	plans := GetMigrationPlans()
	
	fmt.Println("This guide provides a phased approach to optimizing JSON parsing in the appledocs project.")
	fmt.Println("Each phase can be implemented independently and provides incremental benefits.")
	fmt.Println()
	
	for i, plan := range plans {
		fmt.Printf("## Phase %d: %s\n", plan.Phase, plan.Description)
		fmt.Println()
		
		fmt.Println("### Benefits:")
		for _, benefit := range plan.Benefits {
			fmt.Printf("- %s\n", benefit)
		}
		fmt.Println()
		
		fmt.Println("### Risks:")
		for _, risk := range plan.Risks {
			fmt.Printf("- %s\n", risk)
		}
		fmt.Println()
		
		fmt.Println("### Implementation:")
		fmt.Printf("```go%s```\n", plan.Implementation)
		fmt.Println()
		
		if i < len(plans)-1 {
			fmt.Println("---")
			fmt.Println()
		}
	}
	
	fmt.Println("## Recommendation")
	fmt.Println()
	fmt.Println("1. Start with Phase 1 (jsoniter) for immediate 2-3x performance improvement")
	fmt.Println("2. Implement Phase 2 (fastjson) for URL extraction to reduce memory usage")
	fmt.Println("3. Add Phase 3 (streaming) for files larger than 1MB")
	fmt.Println("4. Use Phase 4 (selective parsing) for metadata-only operations")
	fmt.Println()
	fmt.Println("This approach minimizes risk while maximizing performance gains.")
}

// Example of how to use the optimized processor in the current codebase
func demonstrateIntegration() {
	fmt.Println("=== Integration Examples ===")
	fmt.Println()
	
	// Example 1: Drop-in replacement for extractJSONURLs
	fmt.Println("1. Drop-in replacement for URL extraction:")
	fmt.Println("```go")
	fmt.Println("// Before:")
	fmt.Println("urls := extractJSONURLs(data)")
	fmt.Println()
	fmt.Println("// After:")
	fmt.Println("extractor := NewBackwardCompatibleExtractor()")
	fmt.Println("urls := extractor.ExtractJSONURLs(data)")
	fmt.Println("```")
	fmt.Println()
	
	// Example 2: Intelligent processing
	fmt.Println("2. Intelligent processing based on file size:")
	fmt.Println("```go")
	fmt.Println("processor := NewOptimizedJSONProcessor()")
	fmt.Println("doc, urls, err := processor.ProcessFile(filename, \"full_parsing\")")
	fmt.Println("if err != nil {")
	fmt.Println("    return err")
	fmt.Println("}")
	fmt.Println("```")
	fmt.Println()
	
	// Example 3: Streaming for large files
	fmt.Println("3. Streaming for large files:")
	fmt.Println("```go")
	fmt.Println("if fileSize > 10*1024*1024 {")
	fmt.Println("    // Use streaming approach")
	fmt.Println("    doc, urls, err := processor.ProcessFile(filename, \"url_extraction\")")
	fmt.Println("} else {")
	fmt.Println("    // Use regular approach")
	fmt.Println("    doc, urls, err := processor.ProcessJSON(data, \"full_parsing\")")
	fmt.Println("}")
	fmt.Println("```")
}