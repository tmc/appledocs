// Package main provides comprehensive benchmarks for JSON parsing libraries
// tailored for Apple documentation processing
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/buger/jsonparser"
	"github.com/json-iterator/go"
	"github.com/tidwall/gjson"
	"github.com/valyala/fastjson"
)

// Test data sizes for benchmarking
type TestDataSet struct {
	Name     string
	FilePath string
	Data     []byte
	Size     int64
}

// Benchmark configuration
type BenchmarkConfig struct {
	TestDataSets []TestDataSet
	Iterations   int
}

// Memory usage tracking
type MemoryStats struct {
	AllocsBefore   uint64
	AllocsAfter    uint64
	AllocsDelta    uint64
	BytesBefore    uint64
	BytesAfter     uint64
	BytesDelta     uint64
	SysBefore      uint64
	SysAfter       uint64
	SysDelta       uint64
	GCCycles       uint32
	GCPauseTotalNs uint64
}

// Benchmark result
type BenchmarkResult struct {
	Library       string
	TestCase      string
	Duration      time.Duration
	MemoryStats   MemoryStats
	ParsedURLs    int
	Success       bool
	Error         string
	BytesPerSec   float64
	AllocsPerOp   float64
}

// Test data initialization
func initTestData() (BenchmarkConfig, error) {
	config := BenchmarkConfig{
		Iterations: 10,
	}

	// Find test files of different sizes
	testFiles := []struct {
		pattern string
		name    string
	}{
		{"**/technologies.json", "small"},
		{"**/MPSGraph.json", "medium"},
		{"**/View-Implementations.json", "large"},
	}

	outputDir := "output"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		outputDir = "."
	}

	for _, tf := range testFiles {
		matches, err := filepath.Glob(filepath.Join(outputDir, "**", tf.pattern))
		if err != nil {
			continue
		}

		for _, match := range matches {
			if info, err := os.Stat(match); err == nil && !info.IsDir() {
				data, err := os.ReadFile(match)
				if err != nil {
					continue
				}

				config.TestDataSets = append(config.TestDataSets, TestDataSet{
					Name:     fmt.Sprintf("%s_%s", tf.name, filepath.Base(match)),
					FilePath: match,
					Data:     data,
					Size:     info.Size(),
				})
				break // Only take the first match for each pattern
			}
		}
	}

	// If no files found, create synthetic test data
	if len(config.TestDataSets) == 0 {
		config.TestDataSets = append(config.TestDataSets, createSyntheticTestData()...)
	}

	return config, nil
}

// Create synthetic test data for benchmarking
func createSyntheticTestData() []TestDataSet {
	return []TestDataSet{
		{
			Name: "small_synthetic",
			Data: []byte(`{
				"metadata": {"title": "Test", "platforms": [{"name": "iOS", "introducedAt": "14.0"}]},
				"abstract": [{"type": "text", "text": "Small test data"}],
				"references": {
					"ref1": {"url": "doc://test/path1.json", "title": "Ref 1"},
					"ref2": {"url": "doc://test/path2.json", "title": "Ref 2"}
				},
				"topicSections": [{"identifiers": ["ref1", "ref2"], "title": "Topics"}]
			}`),
			Size: 400,
		},
		{
			Name: "medium_synthetic",
			Data: createLargeTestJSON(1000), // 1000 references
			Size: 150000,                    // ~150KB
		},
		{
			Name: "large_synthetic",
			Data: createLargeTestJSON(10000), // 10000 references
			Size: 1500000,                    // ~1.5MB
		},
	}
}

// Create large JSON for testing
func createLargeTestJSON(refCount int) []byte {
	var buf bytes.Buffer
	buf.WriteString(`{
		"metadata": {"title": "Large Test", "platforms": [{"name": "iOS", "introducedAt": "14.0"}]},
		"abstract": [{"type": "text", "text": "Large test data with many references"}],
		"references": {`)

	for i := 0; i < refCount; i++ {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf(`
			"ref%d": {
				"url": "doc://test/path%d.json",
				"title": "Reference %d",
				"abstract": [{"type": "text", "text": "Description for reference %d"}],
				"kind": "symbol",
				"role": "symbol"
			}`, i, i, i, i))
	}

	buf.WriteString(`},
		"topicSections": [{"identifiers": [`)

	for i := 0; i < refCount; i++ {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf(`"ref%d"`, i))
	}

	buf.WriteString(`], "title": "All Topics"}]
	}`)

	return buf.Bytes()
}

// Memory tracking utilities
func getMemoryStats() MemoryStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return MemoryStats{
		AllocsBefore:   m.Mallocs,
		BytesBefore:    m.TotalAlloc,
		SysBefore:      m.Sys,
		GCCycles:       m.NumGC,
		GCPauseTotalNs: m.PauseTotalNs,
	}
}

func (ms *MemoryStats) Update() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	ms.AllocsAfter = m.Mallocs
	ms.BytesAfter = m.TotalAlloc
	ms.SysAfter = m.Sys
	ms.AllocsDelta = ms.AllocsAfter - ms.AllocsBefore
	ms.BytesDelta = ms.BytesAfter - ms.BytesBefore
	ms.SysDelta = ms.SysAfter - ms.SysBefore
}

// Benchmark standard library encoding/json
func benchmarkStdJSON(data []byte) BenchmarkResult {
	runtime.GC()
	memStats := getMemoryStats()
	start := time.Now()

	var doc DocJSONData
	if err := json.Unmarshal(data, &doc); err != nil {
		return BenchmarkResult{
			Library:  "encoding/json",
			Duration: time.Since(start),
			Success:  false,
			Error:    err.Error(),
		}
	}

	// Extract URLs to simulate real usage
	urls := extractURLsFromDoc(&doc)
	duration := time.Since(start)
	memStats.Update()

	return BenchmarkResult{
		Library:     "encoding/json",
		Duration:    duration,
		MemoryStats: memStats,
		ParsedURLs:  len(urls),
		Success:     true,
		BytesPerSec: float64(len(data)) / duration.Seconds(),
		AllocsPerOp: float64(memStats.AllocsDelta),
	}
}

// Benchmark jsoniter
func benchmarkJsoniter(data []byte) BenchmarkResult {
	runtime.GC()
	memStats := getMemoryStats()
	start := time.Now()

	var doc DocJSONData
	if err := jsoniter.Unmarshal(data, &doc); err != nil {
		return BenchmarkResult{
			Library:  "jsoniter",
			Duration: time.Since(start),
			Success:  false,
			Error:    err.Error(),
		}
	}

	// Extract URLs to simulate real usage
	urls := extractURLsFromDoc(&doc)
	duration := time.Since(start)
	memStats.Update()

	return BenchmarkResult{
		Library:     "jsoniter",
		Duration:    duration,
		MemoryStats: memStats,
		ParsedURLs:  len(urls),
		Success:     true,
		BytesPerSec: float64(len(data)) / duration.Seconds(),
		AllocsPerOp: float64(memStats.AllocsDelta),
	}
}

// Benchmark fastjson
func benchmarkFastJSON(data []byte) BenchmarkResult {
	runtime.GC()
	memStats := getMemoryStats()
	start := time.Now()

	var p fastjson.Parser
	v, err := p.Parse(string(data))
	if err != nil {
		return BenchmarkResult{
			Library:  "fastjson",
			Duration: time.Since(start),
			Success:  false,
			Error:    err.Error(),
		}
	}

	// Extract URLs to simulate real usage
	urls := extractURLsFromFastJSON(v)
	duration := time.Since(start)
	memStats.Update()

	return BenchmarkResult{
		Library:     "fastjson",
		Duration:    duration,
		MemoryStats: memStats,
		ParsedURLs:  len(urls),
		Success:     true,
		BytesPerSec: float64(len(data)) / duration.Seconds(),
		AllocsPerOp: float64(memStats.AllocsDelta),
	}
}

// Benchmark gjson
func benchmarkGJSON(data []byte) BenchmarkResult {
	runtime.GC()
	memStats := getMemoryStats()
	start := time.Now()

	// Extract URLs directly using gjson
	urls := extractURLsFromGJSON(string(data))
	duration := time.Since(start)
	memStats.Update()

	return BenchmarkResult{
		Library:     "gjson",
		Duration:    duration,
		MemoryStats: memStats,
		ParsedURLs:  len(urls),
		Success:     true,
		BytesPerSec: float64(len(data)) / duration.Seconds(),
		AllocsPerOp: float64(memStats.AllocsDelta),
	}
}

// Benchmark jsonparser
func benchmarkJSONParser(data []byte) BenchmarkResult {
	runtime.GC()
	memStats := getMemoryStats()
	start := time.Now()

	// Extract URLs using jsonparser
	urls := extractURLsFromJSONParser(data)
	duration := time.Since(start)
	memStats.Update()

	return BenchmarkResult{
		Library:     "jsonparser",
		Duration:    duration,
		MemoryStats: memStats,
		ParsedURLs:  len(urls),
		Success:     true,
		BytesPerSec: float64(len(data)) / duration.Seconds(),
		AllocsPerOp: float64(memStats.AllocsDelta),
	}
}

// URL extraction functions for different libraries
func extractURLsFromDoc(doc *DocJSONData) []string {
	var urls []string
	
	// Extract from references
	for _, ref := range doc.References {
		if ref.URL != "" && strings.HasSuffix(ref.URL, ".json") {
			urls = append(urls, ref.URL)
		}
	}
	
	// Extract from topic sections
	for _, section := range doc.TopicSections {
		urls = append(urls, section.Identifiers...)
	}
	
	return urls
}

func extractURLsFromFastJSON(v *fastjson.Value) []string {
	var urls []string
	
	// Extract from references
	refs := v.Get("references")
	if refs != nil {
		refs.GetObject().Visit(func(key []byte, v *fastjson.Value) {
			if url := v.Get("url"); url != nil {
				if urlStr := string(url.GetStringBytes()); strings.HasSuffix(urlStr, ".json") {
					urls = append(urls, urlStr)
				}
			}
		})
	}
	
	// Extract from topic sections
	topics := v.Get("topicSections")
	if topics != nil {
		for _, topic := range topics.GetArray() {
			identifiers := topic.Get("identifiers")
			if identifiers != nil {
				for _, id := range identifiers.GetArray() {
					urls = append(urls, string(id.GetStringBytes()))
				}
			}
		}
	}
	
	return urls
}

func extractURLsFromGJSON(data string) []string {
	var urls []string
	
	// Extract from references
	gjson.Get(data, "references").ForEach(func(key, value gjson.Result) bool {
		if url := value.Get("url"); url.Exists() && strings.HasSuffix(url.String(), ".json") {
			urls = append(urls, url.String())
		}
		return true
	})
	
	// Extract from topic sections
	gjson.Get(data, "topicSections").ForEach(func(key, value gjson.Result) bool {
		value.Get("identifiers").ForEach(func(key, value gjson.Result) bool {
			urls = append(urls, value.String())
			return true
		})
		return true
	})
	
	return urls
}

func extractURLsFromJSONParser(data []byte) []string {
	var urls []string
	
	// Extract from references
	jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		if url, err := jsonparser.GetString(value, "url"); err == nil && strings.HasSuffix(url, ".json") {
			urls = append(urls, url)
		}
		return nil
	}, "references")
	
	// Extract from topic sections
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			urls = append(urls, string(value))
		}, "identifiers")
	}, "topicSections")
	
	return urls
}

// Benchmark streaming JSON parsing
func benchmarkStreamingJSON(data []byte) BenchmarkResult {
	runtime.GC()
	memStats := getMemoryStats()
	start := time.Now()

	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	
	var urls []string
	var doc DocJSONData
	
	if err := decoder.Decode(&doc); err != nil {
		return BenchmarkResult{
			Library:  "streaming/json",
			Duration: time.Since(start),
			Success:  false,
			Error:    err.Error(),
		}
	}
	
	urls = extractURLsFromDoc(&doc)
	duration := time.Since(start)
	memStats.Update()

	return BenchmarkResult{
		Library:     "streaming/json",
		Duration:    duration,
		MemoryStats: memStats,
		ParsedURLs:  len(urls),
		Success:     true,
		BytesPerSec: float64(len(data)) / duration.Seconds(),
		AllocsPerOp: float64(memStats.AllocsDelta),
	}
}

// Run all benchmarks
func runBenchmarks(config BenchmarkConfig) []BenchmarkResult {
	var results []BenchmarkResult
	
	benchmarks := []struct {
		name string
		fn   func([]byte) BenchmarkResult
	}{
		{"encoding/json", benchmarkStdJSON},
		{"jsoniter", benchmarkJsoniter},
		{"fastjson", benchmarkFastJSON},
		{"gjson", benchmarkGJSON},
		{"jsonparser", benchmarkJSONParser},
		{"streaming/json", benchmarkStreamingJSON},
	}
	
	for _, testData := range config.TestDataSets {
		for _, benchmark := range benchmarks {
			// Run multiple iterations and take the average
			var totalDuration time.Duration
			var totalBytes, totalAllocs uint64
			var successCount int
			var lastError string
			
			for i := 0; i < config.Iterations; i++ {
				result := benchmark.fn(testData.Data)
				if result.Success {
					totalDuration += result.Duration
					totalBytes += result.MemoryStats.BytesDelta
					totalAllocs += result.MemoryStats.AllocsDelta
					successCount++
				} else {
					lastError = result.Error
				}
			}
			
			if successCount > 0 {
				avgResult := BenchmarkResult{
					Library:     benchmark.name,
					TestCase:    testData.Name,
					Duration:    totalDuration / time.Duration(successCount),
					Success:     true,
					BytesPerSec: float64(testData.Size*int64(successCount)) / totalDuration.Seconds(),
					AllocsPerOp: float64(totalAllocs) / float64(successCount),
				}
				avgResult.MemoryStats.BytesDelta = totalBytes / uint64(successCount)
				avgResult.MemoryStats.AllocsDelta = totalAllocs / uint64(successCount)
				results = append(results, avgResult)
			} else {
				results = append(results, BenchmarkResult{
					Library:  benchmark.name,
					TestCase: testData.Name,
					Success:  false,
					Error:    lastError,
				})
			}
		}
	}
	
	return results
}

// Test functions for Go testing framework
func BenchmarkJSONParsing(b *testing.B) {
	config, err := initTestData()
	if err != nil {
		b.Fatalf("Failed to initialize test data: %v", err)
	}
	
	if len(config.TestDataSets) == 0 {
		b.Skip("No test data available")
	}
	
	testData := config.TestDataSets[0] // Use first dataset for Go benchmarks
	
	b.Run("StdJSON", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var doc DocJSONData
			json.Unmarshal(testData.Data, &doc)
		}
	})
	
	b.Run("Jsoniter", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var doc DocJSONData
			jsoniter.Unmarshal(testData.Data, &doc)
		}
	})
	
	b.Run("FastJSON", func(b *testing.B) {
		var p fastjson.Parser
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			p.Parse(string(testData.Data))
		}
	})
	
	b.Run("GJSON", func(b *testing.B) {
		data := string(testData.Data)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			extractURLsFromGJSON(data)
		}
	})
	
	b.Run("JSONParser", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			extractURLsFromJSONParser(testData.Data)
		}
	})
	
	b.Run("StreamingJSON", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			reader := bytes.NewReader(testData.Data)
			decoder := json.NewDecoder(reader)
			var doc DocJSONData
			decoder.Decode(&doc)
		}
	})
}

// Memory benchmark
func BenchmarkJSONMemory(b *testing.B) {
	config, err := initTestData()
	if err != nil {
		b.Fatalf("Failed to initialize test data: %v", err)
	}
	
	if len(config.TestDataSets) == 0 {
		b.Skip("No test data available")
	}
	
	testData := config.TestDataSets[len(config.TestDataSets)-1] // Use largest dataset
	
	b.Run("StdJSON", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var doc DocJSONData
			json.Unmarshal(testData.Data, &doc)
		}
	})
	
	b.Run("FastJSON", func(b *testing.B) {
		var p fastjson.Parser
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			p.Parse(string(testData.Data))
		}
	})
}

// Print benchmark results
func printBenchmarkResults(results []BenchmarkResult) {
	fmt.Printf("JSON Parsing Library Benchmark Results\n")
	fmt.Printf("=====================================\n\n")
	
	// Group by test case
	testCases := make(map[string][]BenchmarkResult)
	for _, result := range results {
		testCases[result.TestCase] = append(testCases[result.TestCase], result)
	}
	
	for testCase, caseResults := range testCases {
		fmt.Printf("Test Case: %s\n", testCase)
		fmt.Printf("%-15s %-12s %-12s %-12s %-12s %-10s\n", 
			"Library", "Duration", "MB/s", "Allocs/Op", "Bytes/Op", "URLs")
		fmt.Printf("%-15s %-12s %-12s %-12s %-12s %-10s\n", 
			"-------", "--------", "----", "---------", "--------", "----")
		
		for _, result := range caseResults {
			if result.Success {
				fmt.Printf("%-15s %-12s %-12.2f %-12.0f %-12d %-10d\n",
					result.Library,
					result.Duration.Round(time.Microsecond).String(),
					result.BytesPerSec/1024/1024,
					result.AllocsPerOp,
					result.MemoryStats.BytesDelta,
					result.ParsedURLs)
			} else {
				fmt.Printf("%-15s %-12s %-12s %-12s %-12s %-10s\n",
					result.Library, "FAILED", "-", "-", "-", "-")
			}
		}
		fmt.Printf("\n")
	}
}

// Standalone benchmark runner
func runStandaloneBenchmarks() {
	fmt.Println("Initializing JSON parsing library benchmarks...")
	
	config, err := initTestData()
	if err != nil {
		fmt.Printf("Error initializing test data: %v\n", err)
		return
	}
	
	fmt.Printf("Running benchmarks with %d test datasets...\n", len(config.TestDataSets))
	
	results := runBenchmarks(config)
	printBenchmarkResults(results)
	
	// Print summary
	fmt.Printf("Summary:\n")
	fmt.Printf("- Total test datasets: %d\n", len(config.TestDataSets))
	fmt.Printf("- Libraries tested: 6\n")
	fmt.Printf("- Iterations per test: %d\n", config.Iterations)
	
	var totalSize int64
	for _, ds := range config.TestDataSets {
		totalSize += ds.Size
	}
	fmt.Printf("- Total data processed: %d KB\n", totalSize/1024)
}