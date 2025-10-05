// Package benchmark provides benchmark runners for different operations
package benchmark

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

// BenchmarkRunner interface for running benchmarks
type BenchmarkRunner interface {
	// Run executes the benchmark operation
	Run(ctx context.Context, op Operation, data []byte) (interface{}, error)
	
	// Name returns the runner name
	Name() string
	
	// Setup prepares the runner (called once before benchmarks)
	Setup() error
	
	// Teardown cleans up (called once after benchmarks)
	Teardown() error
}

// BaselineRunner implements the baseline (current) implementation
type BaselineRunner struct {
	// Simulate current appledocs implementation
}

func (r *BaselineRunner) Name() string {
	return "baseline"
}

func (r *BaselineRunner) Setup() error {
	return nil
}

func (r *BaselineRunner) Teardown() error {
	return nil
}

func (r *BaselineRunner) Run(ctx context.Context, op Operation, data []byte) (interface{}, error) {
	switch op.Type {
	case "parse":
		return r.parseJSON(data)
	case "extract":
		return r.extractURLs(data)
	case "transform":
		return r.transformData(data)
	default:
		return nil, fmt.Errorf("unknown operation: %s", op.Type)
	}
}

func (r *BaselineRunner) parseJSON(data []byte) (interface{}, error) {
	var result map[string]interface{}
	return result, json.Unmarshal(data, &result)
}

func (r *BaselineRunner) extractURLs(data []byte) (interface{}, error) {
	// Simulate current URL extraction logic
	var doc map[string]interface{}
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	
	urls := make([]string, 0)
	extractURLsRecursive(doc, &urls)
	return urls, nil
}

func (r *BaselineRunner) transformData(data []byte) (interface{}, error) {
	// Simulate markdown transformation
	var doc map[string]interface{}
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	
	// Simple transformation simulation
	return fmt.Sprintf("# %v\n\nContent", doc["title"]), nil
}

// OptimizedJSONRunner implements JSON parsing optimizations
type OptimizedJSONRunner struct {
	parser JSONParser
}

type JSONParser interface {
	Parse(data []byte) (interface{}, error)
	ExtractURLs(data []byte) ([]string, error)
}

func (r *OptimizedJSONRunner) Name() string {
	return "json-optimized"
}

func (r *OptimizedJSONRunner) Setup() error {
	// Initialize optimized parser (e.g., jsoniter)
	return nil
}

func (r *OptimizedJSONRunner) Teardown() error {
	return nil
}

func (r *OptimizedJSONRunner) Run(ctx context.Context, op Operation, data []byte) (interface{}, error) {
	switch op.Type {
	case "parse":
		return r.parser.Parse(data)
	case "extract":
		return r.parser.ExtractURLs(data)
	default:
		return nil, fmt.Errorf("unsupported operation: %s", op.Type)
	}
}

// StreamingRunner implements streaming JSON processing
type StreamingRunner struct {
	bufferSize int
}

func (r *StreamingRunner) Name() string {
	return "streaming"
}

func (r *StreamingRunner) Setup() error {
	r.bufferSize = 4096 // 4KB buffer
	return nil
}

func (r *StreamingRunner) Teardown() error {
	return nil
}

func (r *StreamingRunner) Run(ctx context.Context, op Operation, data []byte) (interface{}, error) {
	// Implement streaming logic
	// This would use json.Decoder for streaming parsing
	return nil, fmt.Errorf("streaming not implemented")
}

// ConcurrentRunner implements concurrent processing
type ConcurrentRunner struct {
	workers int
}

func (r *ConcurrentRunner) Name() string {
	return "concurrent"
}

func (r *ConcurrentRunner) Setup() error {
	r.workers = runtime.NumCPU()
	return nil
}

func (r *ConcurrentRunner) Teardown() error {
	return nil
}

func (r *ConcurrentRunner) Run(ctx context.Context, op Operation, data []byte) (interface{}, error) {
	// Implement concurrent processing
	return nil, fmt.Errorf("concurrent processing not implemented")
}

// runOperations executes operations for a scenario
func (bf *BenchmarkFramework) runOperations(ctx context.Context, phase Phase, scenario Scenario, 
	runner BenchmarkRunner, testData []TestData, warmup bool) ([]Result, error) {
	
	results := make([]Result, 0)
	
	for _, data := range testData {
		for _, op := range scenario.Operations {
			select {
			case <-ctx.Done():
				return results, ctx.Err()
			default:
			}
			
			// Measure operation
			result := bf.measureOperation(ctx, phase, scenario, op, runner, data)
			
			if !warmup {
				results = append(results, result)
			}
		}
	}
	
	return results, nil
}

// measureOperation measures a single operation
func (bf *BenchmarkFramework) measureOperation(ctx context.Context, phase Phase, scenario Scenario, 
	op Operation, runner BenchmarkRunner, data TestData) Result {
	
	// Capture initial state
	runtime.GC()
	startMem := captureMemStats()
	startTime := time.Now()
	startGoroutines := runtime.NumGoroutine()
	
	// Run operation
	output, err := runner.Run(ctx, op, data.Data)
	
	// Capture final state
	duration := time.Since(startTime)
	endMem := captureMemStats()
	
	result := Result{
		Phase:     phase,
		Scenario:  scenario.Name,
		Operation: op.Name,
		Timestamp: startTime,
		Duration:  duration,
		Memory:    calculateMemoryDelta(startMem, endMem),
		CPU: CPUStats{
			Goroutines: runtime.NumGoroutine() - startGoroutines,
		},
		Platform: getCurrentPlatform(),
		Metadata: map[string]interface{}{
			"file_size":     data.Size,
			"file_category": data.Category,
			"runner":        runner.Name(),
			"output_size":   getOutputSize(output),
		},
	}
	
	if err != nil {
		result.Errors = append(result.Errors, Error{
			Type:      "operation",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
	}
	
	return result
}

// captureMemStats captures current memory statistics
func captureMemStats() runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m
}

// calculateMemoryDelta calculates memory changes
func calculateMemoryDelta(start, end runtime.MemStats) MemoryStats {
	// Calculate GC pause statistics
	gcPauses := make([]time.Duration, 0)
	for i := 0; i < int(end.NumGC-start.NumGC) && i < 256; i++ {
		pauseNs := end.PauseNs[(int(end.NumGC)-1-i)%256]
		gcPauses = append(gcPauses, time.Duration(pauseNs))
	}
	
	var totalPause, maxPause time.Duration
	for _, pause := range gcPauses {
		totalPause += pause
		if pause > maxPause {
			maxPause = pause
		}
	}
	
	avgPause := time.Duration(0)
	if len(gcPauses) > 0 {
		avgPause = totalPause / time.Duration(len(gcPauses))
	}
	
	return MemoryStats{
		HeapAlloc:    int64(end.HeapAlloc),
		HeapInuse:    int64(end.HeapInuse),
		HeapReleased: int64(end.HeapReleased),
		StackInuse:   int64(end.StackInuse),
		Allocations:  int64(end.Mallocs - start.Mallocs),
		TotalAlloc:   int64(end.TotalAlloc - start.TotalAlloc),
		Sys:          int64(end.Sys),
		NumGC:        end.NumGC - start.NumGC,
		GCPauseTotal: totalPause,
		GCPauseAvg:   avgPause,
		GCPauseMax:   maxPause,
	}
}

// getCurrentPlatform returns current platform information
func getCurrentPlatform() Platform {
	return Platform{
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		CPUCores:  runtime.NumCPU(),
		GoVersion: runtime.Version(),
	}
}

// getOutputSize calculates the size of output
func getOutputSize(output interface{}) int64 {
	switch v := output.(type) {
	case []byte:
		return int64(len(v))
	case string:
		return int64(len(v))
	case []string:
		total := 0
		for _, s := range v {
			total += len(s)
		}
		return int64(total)
	default:
		// Estimate via JSON encoding
		if data, err := json.Marshal(output); err == nil {
			return int64(len(data))
		}
		return 0
	}
}

// extractURLsRecursive recursively extracts URLs from a map
func extractURLsRecursive(data interface{}, urls *[]string) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			if key == "url" || strings.HasSuffix(key, "URL") || strings.HasSuffix(key, "Uri") {
				if str, ok := value.(string); ok && strings.HasSuffix(str, ".json") {
					*urls = append(*urls, str)
				}
			}
			extractURLsRecursive(value, urls)
		}
	case []interface{}:
		for _, item := range v {
			extractURLsRecursive(item, urls)
		}
	}
}

// aggregateResults aggregates results from multiple iterations
func (bf *BenchmarkFramework) aggregateResults(scenario Scenario, iterations [][]Result) []Result {
	// Group results by operation
	operationResults := make(map[string][][]Result)
	
	for _, iteration := range iterations {
		for _, result := range iteration {
			key := result.Operation
			operationResults[key] = append(operationResults[key], iteration)
		}
	}
	
	aggregated := make([]Result, 0)
	
	for operation, iterResults := range operationResults {
		if len(iterResults) == 0 {
			continue
		}
		
		// Calculate statistics for each operation
		durations := make([]float64, 0)
		memories := make([]float64, 0)
		allocations := make([]float64, 0)
		
		var lastResult Result
		for _, iteration := range iterResults {
			for _, result := range iteration {
				if result.Operation == operation {
					durations = append(durations, float64(result.Duration))
					memories = append(memories, float64(result.Memory.TotalAlloc))
					allocations = append(allocations, float64(result.Memory.Allocations))
					lastResult = result
				}
			}
		}
		
		// Calculate statistics
		durationStats := calculateStatistics(durations)
		memoryStats := calculateStatistics(memories)
		allocationStats := calculateStatistics(allocations)
		
		// Create aggregated result
		aggResult := lastResult
		aggResult.Duration = time.Duration(durationStats.Mean)
		aggResult.Memory.TotalAlloc = int64(memoryStats.Mean)
		aggResult.Memory.Allocations = int64(allocationStats.Mean)
		
		// Add statistics to metadata
		aggResult.Metadata["duration_stats"] = durationStats
		aggResult.Metadata["memory_stats"] = memoryStats
		aggResult.Metadata["allocation_stats"] = allocationStats
		aggResult.Metadata["iterations"] = len(durations)
		
		aggregated = append(aggregated, aggResult)
	}
	
	return aggregated
}

// calculateStatistics calculates statistical measures
func calculateStatistics(values []float64) Statistics {
	if len(values) == 0 {
		return Statistics{}
	}
	
	mean, _ := stats.Mean(values)
	median, _ := stats.Median(values)
	stdDev, _ := stats.StandardDeviation(values)
	min, _ := stats.Min(values)
	max, _ := stats.Max(values)
	
	// Calculate percentiles
	p50, _ := stats.Percentile(values, 50)
	p90, _ := stats.Percentile(values, 90)
	p95, _ := stats.Percentile(values, 95)
	p99, _ := stats.Percentile(values, 99)
	
	// Coefficient of variation
	cv := 0.0
	if mean > 0 {
		cv = stdDev / mean
	}
	
	// 95% confidence interval
	ci95Lower, ci95Upper := confidenceInterval(values, 0.95)
	
	return Statistics{
		Mean:      mean,
		Median:    median,
		StdDev:    stdDev,
		Min:       min,
		Max:       max,
		P50:       p50,
		P90:       p90,
		P95:       p95,
		P99:       p99,
		CV:        cv,
		CI95Lower: ci95Lower,
		CI95Upper: ci95Upper,
	}
}

// confidenceInterval calculates confidence interval
func confidenceInterval(values []float64, confidence float64) (lower, upper float64) {
	if len(values) == 0 {
		return 0, 0
	}
	
	mean, _ := stats.Mean(values)
	stdErr, _ := stats.StandardError(values)
	
	// For 95% confidence, z = 1.96
	z := 1.96
	if confidence == 0.99 {
		z = 2.576
	}
	
	margin := z * stdErr
	return mean - margin, mean + margin
}

// generateSyntheticData generates synthetic test data
func (bf *BenchmarkFramework) generateSyntheticData(scenario Scenario) []TestData {
	sizes := map[string]int64{
		"small":  10 * 1024,        // 10KB
		"medium": 1024 * 1024,      // 1MB
		"large":  10 * 1024 * 1024, // 10MB
		"xlarge": 50 * 1024 * 1024, // 50MB
	}
	
	testData := make([]TestData, 0)
	
	for _, sizeCategory := range scenario.FileSizes {
		size, ok := sizes[sizeCategory]
		if !ok {
			continue
		}
		
		data := generateJSONData(size)
		testData = append(testData, TestData{
			Name:     fmt.Sprintf("synthetic_%s.json", sizeCategory),
			Data:     data,
			Size:     int64(len(data)),
			Category: sizeCategory,
		})
	}
	
	return testData
}

// generateJSONData generates synthetic JSON data of specified size
func generateJSONData(targetSize int64) []byte {
	// Create a structure similar to Apple documentation
	doc := map[string]interface{}{
		"metadata": map[string]interface{}{
			"title": "Synthetic Test Document",
			"platforms": []map[string]interface{}{
				{"name": "iOS", "introducedAt": "14.0"},
				{"name": "macOS", "introducedAt": "11.0"},
			},
		},
		"abstract": []map[string]interface{}{
			{"type": "text", "text": "This is a synthetic test document for benchmarking"},
		},
		"references": make(map[string]interface{}),
		"topicSections": []map[string]interface{}{},
	}
	
	// Add references until we reach target size
	refCount := 0
	for {
		refID := fmt.Sprintf("ref%d", refCount)
		doc["references"].(map[string]interface{})[refID] = map[string]interface{}{
			"url":   fmt.Sprintf("doc://test/reference%d.json", refCount),
			"title": fmt.Sprintf("Reference %d", refCount),
			"abstract": []map[string]interface{}{
				{"type": "text", "text": fmt.Sprintf("Description for reference %d", refCount)},
			},
		}
		refCount++
		
		// Check size periodically
		if refCount%100 == 0 {
			data, _ := json.Marshal(doc)
			if int64(len(data)) >= targetSize {
				return data
			}
		}
	}
}