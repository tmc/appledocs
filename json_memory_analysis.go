//go:build ignore

// Package main provides comprehensive memory analysis for JSON parsing
// in the appledocs project
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"time"

	"github.com/buger/jsonparser"
	"github.com/json-iterator/go"
	"github.com/tidwall/gjson"
	"github.com/valyala/fastjson"
)

// MemoryProfile represents memory usage for a specific operation
type MemoryProfile struct {
	Library        string
	FileName       string
	FileSize       int64
	ParseDuration  time.Duration
	HeapAllocBefore uint64
	HeapAllocAfter  uint64
	HeapAllocDelta  uint64
	SysMemBefore   uint64
	SysMemAfter    uint64
	SysMemDelta    uint64
	NumGCBefore    uint32
	NumGCAfter     uint32
	GCPauseBefore  uint64
	GCPauseAfter   uint64
	URLsExtracted  int
	Success        bool
	Error          string
}

// ProfileResult aggregates results from multiple parsing attempts
type ProfileResult struct {
	Library       string
	TestFiles     []string
	AvgDuration   time.Duration
	AvgMemoryUsed uint64
	MemoryEfficiency float64 // MB/s
	TotalURLs     int
	SuccessRate   float64
}

// getMemoryStats captures current memory statistics
func getMemoryStats() (uint64, uint64, uint32, uint64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.HeapAlloc, m.Sys, m.NumGC, m.PauseTotalNs
}

// profileJSONLibrary profiles a specific JSON library with given test files
func profileJSONLibrary(library string, testFiles []string) ([]MemoryProfile, error) {
	var profiles []MemoryProfile

	for _, filename := range testFiles {
		profile := MemoryProfile{
			Library:  library,
			FileName: filename,
		}

		// Get file size
		if info, err := os.Stat(filename); err == nil {
			profile.FileSize = info.Size()
		}

		// Read file data
		data, err := os.ReadFile(filename)
		if err != nil {
			profile.Error = fmt.Sprintf("read file: %v", err)
			profiles = append(profiles, profile)
			continue
		}

		// Force garbage collection before measurement
		runtime.GC()
		runtime.GC() // Call twice to ensure cleanup

		// Capture before state
		profile.HeapAllocBefore, profile.SysMemBefore, profile.NumGCBefore, profile.GCPauseBefore = getMemoryStats()

		// Perform parsing based on library
		start := time.Now()
		var urls []string

		switch library {
		case "encoding/json":
			urls, err = parseWithStandardJSON(data)
		case "jsoniter":
			urls, err = parseWithJsoniter(data)
		case "fastjson":
			urls, err = parseWithFastJSON(data)
		case "gjson":
			urls, err = parseWithGJSON(data)
		case "jsonparser":
			urls, err = parseWithJSONParser(data)
		case "streaming/json":
			urls, err = parseWithStreamingJSON(data)
		default:
			err = fmt.Errorf("unknown library: %s", library)
		}

		profile.ParseDuration = time.Since(start)

		// Capture after state
		profile.HeapAllocAfter, profile.SysMemAfter, profile.NumGCAfter, profile.GCPauseAfter = getMemoryStats()

		// Calculate deltas
		profile.HeapAllocDelta = profile.HeapAllocAfter - profile.HeapAllocBefore
		profile.SysMemDelta = profile.SysMemAfter - profile.SysMemBefore

		if err != nil {
			profile.Error = err.Error()
		} else {
			profile.Success = true
			profile.URLsExtracted = len(urls)
		}

		profiles = append(profiles, profile)
	}

	return profiles, nil
}

// parseWithStandardJSON parses JSON using encoding/json
func parseWithStandardJSON(data []byte) ([]string, error) {
	var doc DocJSONData
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	return extractURLsFromDoc(&doc), nil
}

// parseWithJsoniter parses JSON using jsoniter
func parseWithJsoniter(data []byte) ([]string, error) {
	var doc DocJSONData
	if err := jsoniter.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	return extractURLsFromDoc(&doc), nil
}

// parseWithFastJSON parses JSON using fastjson
func parseWithFastJSON(data []byte) ([]string, error) {
	var p fastjson.Parser
	v, err := p.Parse(string(data))
	if err != nil {
		return nil, err
	}
	return extractURLsFromFastJSON(v), nil
}

// parseWithGJSON parses JSON using gjson
func parseWithGJSON(data []byte) ([]string, error) {
	return extractURLsFromGJSON(string(data)), nil
}

// parseWithJSONParser parses JSON using jsonparser
func parseWithJSONParser(data []byte) ([]string, error) {
	return extractURLsFromJSONParser(data), nil
}

// parseWithStreamingJSON parses JSON using streaming approach
func parseWithStreamingJSON(data []byte) ([]string, error) {
	decoder := json.NewDecoder(bytesReader{data, 0})
	var doc DocJSONData
	if err := decoder.Decode(&doc); err != nil {
		return nil, err
	}
	return extractURLsFromDoc(&doc), nil
}

// bytesReader implements io.Reader for byte slice
type bytesReader struct {
	data []byte
	pos  int
}

func (r *bytesReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.data) {
		return 0, fmt.Errorf("EOF")
	}
	n = copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}

// generateMemoryReport creates a comprehensive memory usage report
func generateMemoryReport(profiles []MemoryProfile) {
	fmt.Println("=== JSON Parsing Memory Analysis Report ===")
	fmt.Println()

	// Group profiles by library
	libraryProfiles := make(map[string][]MemoryProfile)
	for _, profile := range profiles {
		libraryProfiles[profile.Library] = append(libraryProfiles[profile.Library], profile)
	}

	// Print detailed analysis for each library
	for library, libProfiles := range libraryProfiles {
		fmt.Printf("Library: %s\n", library)
		fmt.Printf("%-40s %-10s %-12s %-12s %-10s %-8s\n", 
			"File", "Size(KB)", "Duration", "Memory(KB)", "URLs", "Success")
		fmt.Printf("%-40s %-10s %-12s %-12s %-10s %-8s\n", 
			"----", "-------", "--------", "---------", "----", "-------")

		var totalDuration time.Duration
		var totalMemory uint64
		var totalURLs int
		var successCount int

		for _, profile := range libProfiles {
			filename := filepath.Base(profile.FileName)
			if len(filename) > 38 {
				filename = filename[:35] + "..."
			}
			
			status := "✓"
			if !profile.Success {
				status = "✗"
			}

			fmt.Printf("%-40s %-10d %-12s %-12d %-10d %-8s\n",
				filename,
				profile.FileSize/1024,
				profile.ParseDuration.Round(time.Microsecond),
				profile.HeapAllocDelta/1024,
				profile.URLsExtracted,
				status)

			if profile.Success {
				totalDuration += profile.ParseDuration
				totalMemory += profile.HeapAllocDelta
				totalURLs += profile.URLsExtracted
				successCount++
			}
		}

		// Calculate averages
		if successCount > 0 {
			avgDuration := totalDuration / time.Duration(successCount)
			avgMemory := totalMemory / uint64(successCount)
			
			fmt.Printf("\nSummary for %s:\n", library)
			fmt.Printf("  Average Duration: %v\n", avgDuration.Round(time.Microsecond))
			fmt.Printf("  Average Memory:   %d KB\n", avgMemory/1024)
			fmt.Printf("  Total URLs:       %d\n", totalURLs)
			fmt.Printf("  Success Rate:     %.1f%%\n", float64(successCount)/float64(len(libProfiles))*100)
		}
		fmt.Println()
	}

	// Create comparative analysis
	fmt.Println("=== Comparative Performance Analysis ===")
	
	var results []ProfileResult
	for library, libProfiles := range libraryProfiles {
		result := ProfileResult{Library: library}
		
		var totalDuration time.Duration
		var totalMemory uint64
		var successCount int

		for _, profile := range libProfiles {
			result.TestFiles = append(result.TestFiles, profile.FileName)
			if profile.Success {
				totalDuration += profile.ParseDuration
				totalMemory += profile.HeapAllocDelta
				result.TotalURLs += profile.URLsExtracted
				successCount++
			}
		}

		if successCount > 0 {
			result.AvgDuration = totalDuration / time.Duration(successCount)
			result.AvgMemoryUsed = totalMemory / uint64(successCount)
			result.SuccessRate = float64(successCount) / float64(len(libProfiles)) * 100
			
			// Calculate memory efficiency (URLs processed per MB per second)
			if result.AvgDuration > 0 && result.AvgMemoryUsed > 0 {
				urlsPerSec := float64(result.TotalURLs) / result.AvgDuration.Seconds()
				memoryMB := float64(result.AvgMemoryUsed) / (1024 * 1024)
				result.MemoryEfficiency = urlsPerSec / memoryMB
			}
		}
		
		results = append(results, result)
	}

	// Sort by memory efficiency
	sort.Slice(results, func(i, j int) bool {
		return results[i].MemoryEfficiency > results[j].MemoryEfficiency
	})

	fmt.Printf("%-15s %-12s %-12s %-15s %-10s %-10s\n", 
		"Library", "Avg Duration", "Avg Mem(KB)", "Efficiency", "URLs", "Success%")
	fmt.Printf("%-15s %-12s %-12s %-15s %-10s %-10s\n", 
		"-------", "------------", "-----------", "----------", "----", "--------")
	
	for _, result := range results {
		fmt.Printf("%-15s %-12s %-12d %-15.2f %-10d %-10.1f\n",
			result.Library,
			result.AvgDuration.Round(time.Microsecond),
			result.AvgMemoryUsed/1024,
			result.MemoryEfficiency,
			result.TotalURLs,
			result.SuccessRate)
	}
}

// findTestFiles discovers JSON files for testing
func findTestFiles() ([]string, error) {
	var testFiles []string
	
	// Look for test files in output directory
	outputDir := "output"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("output directory not found")
	}

	err := filepath.Walk(outputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			// Skip very small files (< 1KB) and very large files (> 10MB)
			if info.Size() > 1024 && info.Size() < 10*1024*1024 {
				testFiles = append(testFiles, path)
			}
		}
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	// Limit to a manageable number of test files
	if len(testFiles) > 20 {
		// Sort by size and take a representative sample
		sort.Slice(testFiles, func(i, j int) bool {
			info1, _ := os.Stat(testFiles[i])
			info2, _ := os.Stat(testFiles[j])
			return info1.Size() < info2.Size()
		})
		
		// Take every nth file to get a good size distribution
		step := len(testFiles) / 20
		var sampledFiles []string
		for i := 0; i < len(testFiles); i += step {
			sampledFiles = append(sampledFiles, testFiles[i])
		}
		testFiles = sampledFiles
	}
	
	return testFiles, nil
}

// runMemoryAnalysis executes the complete memory analysis
func runMemoryAnalysis() error {
	fmt.Println("Starting comprehensive JSON parsing memory analysis...")
	
	testFiles, err := findTestFiles()
	if err != nil {
		return fmt.Errorf("find test files: %v", err)
	}
	
	if len(testFiles) == 0 {
		return fmt.Errorf("no suitable test files found")
	}
	
	fmt.Printf("Found %d test files for analysis\n", len(testFiles))
	
	libraries := []string{
		"encoding/json",
		"jsoniter", 
		"fastjson",
		"gjson",
		"jsonparser",
		"streaming/json",
	}
	
	var allProfiles []MemoryProfile
	
	for _, library := range libraries {
		fmt.Printf("Profiling %s...\n", library)
		profiles, err := profileJSONLibrary(library, testFiles)
		if err != nil {
			fmt.Printf("Warning: Error profiling %s: %v\n", library, err)
			continue
		}
		allProfiles = append(allProfiles, profiles...)
	}
	
	// Generate the comprehensive report
	generateMemoryReport(allProfiles)
	
	return nil
}