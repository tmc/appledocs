//go:build ignore

// Package main provides enhanced benchmarking for JSON parsing libraries
// with comprehensive analysis for Apple documentation processing
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/buger/jsonparser"
	"github.com/json-iterator/go"
	"github.com/tidwall/gjson"
	"github.com/valyala/fastjson"
)

// EnhancedBenchmarkConfig provides comprehensive benchmark configuration
type EnhancedBenchmarkConfig struct {
	TestCases      []TestCase
	Iterations     int
	WarmupRounds   int
	MemoryProfile  bool
	CPUProfile     bool
	StreamingTests bool
}

// TestCase represents a specific test scenario
type TestCase struct {
	Name        string
	FilePath    string
	Data        []byte
	Size        int64
	Category    string // small, medium, large, xlarge
	Description string
}

// BenchmarkMetrics provides detailed performance metrics
type BenchmarkMetrics struct {
	Library          string
	TestCase         string
	Category         string
	Duration         time.Duration
	BytesPerSec      float64
	URLsPerSec       float64
	MemoryAllocated  uint64
	AllocsPerOp      uint64
	URLsExtracted    int
	Success          bool
	Error            string
	CPUPercent       float64
	GCPauses         time.Duration
	StreamingCapable bool
}

// PerformanceComparison aggregates performance across libraries
type PerformanceComparison struct {
	TestCase  string
	Category  string
	Results   map[string]BenchmarkMetrics
	Winner    string
	Speedup   float64
	MemWinner string
	MemSaving float64
}

// initEnhancedTestCases creates comprehensive test cases
func initEnhancedTestCases() (EnhancedBenchmarkConfig, error) {
	config := EnhancedBenchmarkConfig{
		Iterations:     5,
		WarmupRounds:   2,
		MemoryProfile:  true,
		StreamingTests: true,
	}

	// Define test case patterns with categories
	testPatterns := []struct {
		pattern     string
		category    string
		description string
		maxSize     int64
		minSize     int64
	}{
		{"**/technologies.json", "small", "Main framework listing", 50000, 1000},
		{"**/Foundation.json", "medium", "Foundation framework", 500000, 50000},
		{"**/SwiftUI.json", "medium", "SwiftUI framework", 500000, 50000},
		{"**/UIKit.json", "large", "UIKit framework", 2000000, 500000},
		{"**/View-Implementations.json", "xlarge", "Complex view implementations", 10000000, 2000000},
	}

	outputDir := "output"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		// Create synthetic test cases if no output directory
		config.TestCases = createSyntheticTestCases()
		return config, nil
	}

	// Find real test files
	for _, pattern := range testPatterns {
		matches, err := filepath.Glob(filepath.Join(outputDir, "**", pattern.pattern))
		if err != nil {
			continue
		}

		for _, match := range matches {
			info, err := os.Stat(match)
			if err != nil || info.IsDir() {
				continue
			}

			// Check size constraints
			if info.Size() < pattern.minSize || info.Size() > pattern.maxSize {
				continue
			}

			data, err := os.ReadFile(match)
			if err != nil {
				continue
			}

			// Validate JSON
			if !json.Valid(data) {
				continue
			}

			config.TestCases = append(config.TestCases, TestCase{
				Name:        fmt.Sprintf("%s_%s", pattern.category, filepath.Base(match)),
				FilePath:    match,
				Data:        data,
				Size:        info.Size(),
				Category:    pattern.category,
				Description: pattern.description,
			})
			break // Only take first match for each pattern
		}
	}

	// Add synthetic cases if we don't have enough real ones
	if len(config.TestCases) < 3 {
		config.TestCases = append(config.TestCases, createSyntheticTestCases()...)
	}

	return config, nil
}

// createSyntheticTestCases generates realistic synthetic test data
func createSyntheticTestCases() []TestCase {
	return []TestCase{
		{
			Name:        "small_synthetic_technologies",
			Data:        createSyntheticAppleJSON(50, "framework"),
			Size:        int64(len(createSyntheticAppleJSON(50, "framework"))),
			Category:    "small",
			Description: "Synthetic technologies listing",
		},
		{
			Name:        "medium_synthetic_framework", 
			Data:        createSyntheticAppleJSON(500, "class"),
			Size:        int64(len(createSyntheticAppleJSON(500, "class"))),
			Category:    "medium",
			Description: "Synthetic framework with classes",
		},
		{
			Name:        "large_synthetic_framework",
			Data:        createSyntheticAppleJSON(2000, "method"),
			Size:        int64(len(createSyntheticAppleJSON(2000, "method"))),
			Category:    "large", 
			Description: "Synthetic framework with methods",
		},
		{
			Name:        "xlarge_synthetic_docs",
			Data:        createSyntheticAppleJSON(10000, "symbol"),
			Size:        int64(len(createSyntheticAppleJSON(10000, "symbol"))),
			Category:    "xlarge",
			Description: "Synthetic comprehensive documentation",
		},
	}
}

// createSyntheticAppleJSON creates realistic Apple documentation JSON
func createSyntheticAppleJSON(refCount int, docType string) []byte {
	var buf bytes.Buffer
	
	// Create realistic Apple documentation structure
	buf.WriteString(`{
	"kind": "symbol",
	"identifier": {
		"url": "doc://com.apple.documentation/documentation/` + docType + `",
		"interfaceLanguage": "swift"
	},
	"metadata": {
		"title": "` + strings.Title(docType) + ` Documentation",
		"role": "` + docType + `",
		"modules": [{"name": "Foundation"}],
		"platforms": [
			{"name": "iOS", "introducedAt": "13.0"},
			{"name": "macOS", "introducedAt": "10.15"},
			{"name": "tvOS", "introducedAt": "13.0"},
			{"name": "watchOS", "introducedAt": "6.0"}
		]
	},
	"abstract": [
		{"type": "text", "text": "A comprehensive ` + docType + ` providing core functionality."}
	],
	"primaryContentSections": [
		{
			"kind": "declarations",
			"declarations": [
				{
					"languages": ["swift"],
					"tokens": [
						{"kind": "keyword", "text": "class"},
						{"kind": "text", "text": " "},
						{"kind": "identifier", "text": "` + strings.Title(docType) + `"}
					]
				}
			]
		}
	],
	"references": {`)

	// Generate references
	for i := 0; i < refCount; i++ {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf(`
		"doc://com.apple.documentation/documentation/%s/ref%d": {
			"title": "Reference %d",
			"abstract": [{"type": "text", "text": "Description for reference %d"}],
			"kind": "symbol",
			"role": "symbol",
			"url": "doc://com.apple.documentation/documentation/%s/ref%d.json",
			"type": "topic",
			"identifier": "doc://com.apple.documentation/documentation/%s/ref%d",
			"fragments": [
				{"kind": "identifier", "text": "ref%d"}
			]
		}`, docType, i, i, i, docType, i, docType, i, i))
	}

	buf.WriteString(`
	},
	"topicSections": [
		{
			"title": "` + strings.Title(docType) + ` References",
			"identifiers": [`)

	// Add identifiers to topic sections
	for i := 0; i < refCount; i++ {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf(`"doc://com.apple.documentation/documentation/%s/ref%d"`, docType, i))
	}

	buf.WriteString(`
			]
		}
	]
}`)

	return buf.Bytes()
}

// runEnhancedBenchmark executes comprehensive benchmark for a single library
func runEnhancedBenchmark(library string, testCase TestCase, config EnhancedBenchmarkConfig) BenchmarkMetrics {
	metrics := BenchmarkMetrics{
		Library:  library,
		TestCase: testCase.Name,
		Category: testCase.Category,
	}

	// Warmup rounds
	for i := 0; i < config.WarmupRounds; i++ {
		_, _ = parseWithLibrary(library, testCase.Data)
	}

	// Benchmark rounds
	var totalDuration time.Duration
	var totalMemory uint64
	var totalAllocs uint64
	var totalURLs int
	var successCount int

	for i := 0; i < config.Iterations; i++ {
		// Force garbage collection
		runtime.GC()
		
		// Get memory stats before
		var before, after runtime.MemStats
		runtime.ReadMemStats(&before)

		// Execute parsing
		start := time.Now()
		urls, err := parseWithLibrary(library, testCase.Data)
		duration := time.Since(start)

		// Get memory stats after
		runtime.ReadMemStats(&after)

		if err != nil {
			metrics.Error = err.Error()
			continue
		}

		totalDuration += duration
		totalMemory += after.TotalAlloc - before.TotalAlloc
		totalAllocs += after.Mallocs - before.Mallocs
		totalURLs += len(urls)
		successCount++
	}

	if successCount > 0 {
		metrics.Success = true
		metrics.Duration = totalDuration / time.Duration(successCount)
		metrics.MemoryAllocated = totalMemory / uint64(successCount)
		metrics.AllocsPerOp = totalAllocs / uint64(successCount)
		metrics.URLsExtracted = totalURLs / successCount
		
		// Calculate throughput metrics
		if metrics.Duration > 0 {
			metrics.BytesPerSec = float64(testCase.Size) / metrics.Duration.Seconds()
			metrics.URLsPerSec = float64(metrics.URLsExtracted) / metrics.Duration.Seconds()
		}

		// Check streaming capability
		metrics.StreamingCapable = supportsStreaming(library)
	}

	return metrics
}

// parseWithLibrary parses JSON using the specified library
func parseWithLibrary(library string, data []byte) ([]string, error) {
	switch library {
	case "encoding/json":
		return parseWithStandardJSON(data)
	case "jsoniter":
		return parseWithJsoniter(data)
	case "fastjson":
		return parseWithFastJSON(data)
	case "gjson":
		return parseWithGJSON(data)
	case "jsonparser":
		return parseWithJSONParser(data)
	case "streaming/json":
		return parseWithStreamingJSON(data)
	default:
		return nil, fmt.Errorf("unknown library: %s", library)
	}
}

// supportsStreaming returns true if library supports streaming
func supportsStreaming(library string) bool {
	switch library {
	case "streaming/json", "jsonparser", "fastjson":
		return true
	default:
		return false
	}
}

// runComprehensiveBenchmarks executes all benchmarks and generates analysis
func runComprehensiveBenchmarks() error {
	fmt.Println("Initializing enhanced JSON parsing benchmarks...")
	
	config, err := initEnhancedTestCases()
	if err != nil {
		return fmt.Errorf("initialize test cases: %v", err)
	}

	if len(config.TestCases) == 0 {
		return fmt.Errorf("no test cases available")
	}

	libraries := []string{
		"encoding/json",
		"jsoniter",
		"fastjson", 
		"gjson",
		"jsonparser",
		"streaming/json",
	}

	fmt.Printf("Running benchmarks with %d test cases and %d libraries...\n", 
		len(config.TestCases), len(libraries))

	var allMetrics []BenchmarkMetrics
	var comparisons []PerformanceComparison

	// Run benchmarks for each test case
	for _, testCase := range config.TestCases {
		fmt.Printf("Benchmarking %s (%s, %d KB)...\n", 
			testCase.Name, testCase.Category, testCase.Size/1024)

		comparison := PerformanceComparison{
			TestCase: testCase.Name,
			Category: testCase.Category,
			Results:  make(map[string]BenchmarkMetrics),
		}

		for _, library := range libraries {
			metrics := runEnhancedBenchmark(library, testCase, config)
			allMetrics = append(allMetrics, metrics)
			comparison.Results[library] = metrics
		}

		// Determine winners for this test case
		determineWinners(&comparison)
		comparisons = append(comparisons, comparison)
	}

	// Generate comprehensive report
	generateEnhancedReport(allMetrics, comparisons)
	
	// Generate streaming analysis if enabled
	if config.StreamingTests {
		generateStreamingAnalysis(allMetrics)
	}

	return nil
}

// determineWinners finds the best performing libraries for speed and memory
func determineWinners(comparison *PerformanceComparison) {
	var fastestLib string
	var fastestSpeed float64
	var memEfficientLib string
	var lowestMem uint64 = ^uint64(0) // Max uint64

	for lib, metrics := range comparison.Results {
		if !metrics.Success {
			continue
		}

		// Find fastest
		if metrics.BytesPerSec > fastestSpeed {
			fastestSpeed = metrics.BytesPerSec
			fastestLib = lib
		}

		// Find most memory efficient  
		if metrics.MemoryAllocated < lowestMem {
			lowestMem = metrics.MemoryAllocated
			memEfficientLib = lib
		}
	}

	comparison.Winner = fastestLib
	comparison.MemWinner = memEfficientLib

	// Calculate speedup vs standard library
	if stdMetrics, ok := comparison.Results["encoding/json"]; ok && stdMetrics.Success {
		if winnerMetrics, ok := comparison.Results[fastestLib]; ok {
			comparison.Speedup = winnerMetrics.BytesPerSec / stdMetrics.BytesPerSec
		}
		if memMetrics, ok := comparison.Results[memEfficientLib]; ok {
			if stdMetrics.MemoryAllocated > 0 {
				comparison.MemSaving = float64(stdMetrics.MemoryAllocated) / float64(memMetrics.MemoryAllocated)
			}
		}
	}
}

// generateEnhancedReport creates a comprehensive performance report
func generateEnhancedReport(metrics []BenchmarkMetrics, comparisons []PerformanceComparison) {
	fmt.Println("\n=== Enhanced JSON Parsing Benchmark Report ===")
	fmt.Println()

	// Group metrics by category
	categoryMetrics := make(map[string][]BenchmarkMetrics)
	for _, metric := range metrics {
		categoryMetrics[metric.Category] = append(categoryMetrics[metric.Category], metric)
	}

	// Print detailed results by category
	categories := []string{"small", "medium", "large", "xlarge"}
	for _, category := range categories {
		if catMetrics, exists := categoryMetrics[category]; exists {
			fmt.Printf("=== %s Files ===\n", strings.Title(category))
			printCategoryResults(catMetrics)
			fmt.Println()
		}
	}

	// Print comparison summary
	fmt.Println("=== Performance Summary ===")
	fmt.Printf("%-25s %-15s %-12s %-15s %-12s\n", 
		"Test Case", "Speed Winner", "Speedup", "Memory Winner", "Mem Saving")
	fmt.Printf("%-25s %-15s %-12s %-15s %-12s\n", 
		"---------", "------------", "-------", "-------------", "-----------")

	for _, comp := range comparisons {
		fmt.Printf("%-25s %-15s %-12.2fx %-15s %-12.2fx\n",
			truncateString(comp.TestCase, 24),
			comp.Winner,
			comp.Speedup,
			comp.MemWinner,
			comp.MemSaving)
	}

	// Print overall recommendations
	generateRecommendations(metrics)
}

// printCategoryResults prints detailed results for a category
func printCategoryResults(metrics []BenchmarkMetrics) {
	// Group by test case
	testCases := make(map[string][]BenchmarkMetrics)
	for _, metric := range metrics {
		testCases[metric.TestCase] = append(testCases[metric.TestCase], metric)
	}

	for testCase, caseMetrics := range testCases {
		fmt.Printf("Test: %s\n", testCase)
		fmt.Printf("%-15s %-10s %-12s %-10s %-10s %-10s\n", 
			"Library", "Duration", "MB/s", "URLs/s", "Mem(KB)", "Success")
		fmt.Printf("%-15s %-10s %-12s %-10s %-10s %-10s\n", 
			"-------", "--------", "----", "------", "-------", "-------")

		// Sort by performance
		sort.Slice(caseMetrics, func(i, j int) bool {
			return caseMetrics[i].BytesPerSec > caseMetrics[j].BytesPerSec
		})

		for _, metric := range caseMetrics {
			status := "✓"
			if !metric.Success {
				status = "✗"
			}

			fmt.Printf("%-15s %-10s %-12.1f %-10.1f %-10d %-10s\n",
				metric.Library,
				metric.Duration.Round(time.Microsecond),
				metric.BytesPerSec/1024/1024,
				metric.URLsPerSec,
				metric.MemoryAllocated/1024,
				status)
		}
		fmt.Println()
	}
}

// generateStreamingAnalysis provides streaming-specific analysis
func generateStreamingAnalysis(metrics []BenchmarkMetrics) {
	fmt.Println("=== Streaming Capability Analysis ===")
	
	streamingLibs := []string{}
	nonStreamingLibs := []string{}
	
	libCapabilities := make(map[string]bool)
	for _, metric := range metrics {
		if _, exists := libCapabilities[metric.Library]; !exists {
			libCapabilities[metric.Library] = metric.StreamingCapable
			if metric.StreamingCapable {
				streamingLibs = append(streamingLibs, metric.Library)
			} else {
				nonStreamingLibs = append(nonStreamingLibs, metric.Library)
			}
		}
	}
	
	fmt.Printf("Streaming-capable libraries: %s\n", strings.Join(streamingLibs, ", "))
	fmt.Printf("Non-streaming libraries: %s\n", strings.Join(nonStreamingLibs, ", "))
	fmt.Println()
	
	// Analyze memory usage for large files
	largeFileMetrics := []BenchmarkMetrics{}
	for _, metric := range metrics {
		if metric.Category == "large" || metric.Category == "xlarge" {
			largeFileMetrics = append(largeFileMetrics, metric)
		}
	}
	
	if len(largeFileMetrics) > 0 {
		fmt.Println("Memory usage for large files:")
		fmt.Printf("%-15s %-10s %-15s %-15s\n", "Library", "Streaming", "Avg Memory(KB)", "Memory Efficiency")
		fmt.Printf("%-15s %-10s %-15s %-15s\n", "-------", "---------", "--------------", "-----------------")
		
		libStats := make(map[string]struct {
			totalMemory uint64
			count       int
			streaming   bool
		})
		
		for _, metric := range largeFileMetrics {
			if metric.Success {
				stats := libStats[metric.Library]
				stats.totalMemory += metric.MemoryAllocated
				stats.count++
				stats.streaming = metric.StreamingCapable
				libStats[metric.Library] = stats
			}
		}
		
		for lib, stats := range libStats {
			if stats.count > 0 {
				avgMemory := stats.totalMemory / uint64(stats.count)
				streamingStr := "No"
				if stats.streaming {
					streamingStr = "Yes"
				}
				
				efficiency := "Low"
				if avgMemory < 1024*1024 { // < 1MB
					efficiency = "High"
				} else if avgMemory < 5*1024*1024 { // < 5MB
					efficiency = "Medium"
				}
				
				fmt.Printf("%-15s %-10s %-15d %-15s\n",
					lib, streamingStr, avgMemory/1024, efficiency)
			}
		}
	}
}

// generateRecommendations provides usage recommendations
func generateRecommendations(metrics []BenchmarkMetrics) {
	fmt.Println("\n=== Recommendations ===")
	
	// Analyze patterns across categories
	libPerformance := make(map[string]struct {
		totalSpeed    float64
		totalMemory   uint64
		successCount  int
		categories    map[string]int
	})
	
	for _, metric := range metrics {
		if metric.Success {
			stats := libPerformance[metric.Library]
			stats.totalSpeed += metric.BytesPerSec
			stats.totalMemory += metric.MemoryAllocated
			stats.successCount++
			if stats.categories == nil {
				stats.categories = make(map[string]int)
			}
			stats.categories[metric.Category]++
			libPerformance[metric.Library] = stats
		}
	}
	
	// Find best overall performer
	var bestOverall string
	var bestSpeed float64
	
	for lib, stats := range libPerformance {
		if stats.successCount > 0 {
			avgSpeed := stats.totalSpeed / float64(stats.successCount)
			if avgSpeed > bestSpeed {
				bestSpeed = avgSpeed
				bestOverall = lib
			}
		}
	}
	
	fmt.Printf("🏆 Best Overall Performance: %s\n", bestOverall)
	
	// Category-specific recommendations
	recommendations := map[string]string{
		"small":  "For small files, use fastjson or gjson for best performance",
		"medium": "For medium files, jsoniter provides good balance of speed and memory",
		"large":  "For large files, consider streaming approaches or fastjson",
		"xlarge": "For very large files, streaming JSON or jsonparser recommended",
	}
	
	for category, rec := range recommendations {
		fmt.Printf("📂 %s: %s\n", strings.Title(category), rec)
	}
	
	fmt.Println("\n💡 General Guidelines:")
	fmt.Println("- Use fastjson for maximum speed with acceptable memory usage")
	fmt.Println("- Use gjson for selective parsing and minimal memory footprint")
	fmt.Println("- Use jsoniter as drop-in replacement for encoding/json with better performance")
	fmt.Println("- Use streaming approaches for files > 10MB to avoid memory issues")
	fmt.Println("- Consider jsonparser for very memory-constrained environments")
}

// truncateString truncates a string to the specified length
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// BenchmarkEnhancedJSON provides Go benchmark integration
func BenchmarkEnhancedJSON(b *testing.B) {
	config, err := initEnhancedTestCases()
	if err != nil {
		b.Fatalf("Failed to initialize test cases: %v", err)
	}
	
	if len(config.TestCases) == 0 {
		b.Skip("No test cases available")
	}
	
	// Use first test case for Go benchmarks
	testCase := config.TestCases[0]
	
	libraries := []string{
		"encoding/json",
		"jsoniter", 
		"fastjson",
		"gjson",
		"jsonparser",
	}
	
	for _, library := range libraries {
		b.Run(library, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := parseWithLibrary(library, testCase.Data)
				if err != nil {
					b.Fatalf("Parse error: %v", err)
				}
			}
		})
	}
}