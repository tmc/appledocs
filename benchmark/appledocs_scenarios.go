// Package benchmark provides appledocs-specific benchmark scenarios
package benchmark

import (
	"path/filepath"
	"time"
)

// AppledocsScenarios returns benchmark scenarios specific to appledocs
func AppledocsScenarios() []Scenario {
	return []Scenario{
		{
			Name:        "json_parsing_small",
			Description: "Parse small JSON files (<10KB)",
			FilePatterns: []string{
				"testdata/json/small/*.json",
				".cache/**/small_*.json",
			},
			FileSizes:   []string{"small"},
			Concurrency: []int{1, 5, 10},
			Operations: []Operation{
				{
					Name:     "parse_json",
					Type:     "parse",
					Function: "json.Unmarshal",
				},
				{
					Name:     "extract_urls",
					Type:     "extract",
					Function: "extractJSONURLs",
				},
			},
			Tags:   []string{"json", "parsing", "small"},
			Weight: 0.8,
		},
		{
			Name:        "json_parsing_large",
			Description: "Parse large JSON files (>1MB)",
			FilePatterns: []string{
				"testdata/json/large/*.json",
				".cache/**/technologies.json",
				".cache/**/UIKit.json",
			},
			FileSizes:   []string{"large", "xlarge"},
			Concurrency: []int{1, 5, 10, 20},
			Operations: []Operation{
				{
					Name:     "parse_json",
					Type:     "parse",
					Function: "json.Unmarshal",
				},
				{
					Name:     "extract_urls",
					Type:     "extract",
					Function: "extractJSONURLs",
				},
				{
					Name:     "validate_structure",
					Type:     "validate",
					Function: "validateJSONStructure",
				},
			},
			Tags:   []string{"json", "parsing", "large"},
			Weight: 1.0,
		},
		{
			Name:        "concurrent_crawling",
			Description: "Concurrent URL processing with different concurrency levels",
			FilePatterns: []string{
				"testdata/urls.txt",
			},
			Concurrency: []int{1, 5, 10, 20, 50},
			Operations: []Operation{
				{
					Name:     "process_url_batch",
					Type:     "crawl",
					Function: "processURLBatch",
					Parameters: map[string]interface{}{
						"batch_size": 100,
						"timeout":    30 * time.Second,
					},
				},
			},
			Tags:   []string{"crawling", "concurrent", "network"},
			Weight: 1.2,
		},
		{
			Name:        "cache_operations",
			Description: "Cache read/write performance",
			FilePatterns: []string{
				".cache/**/*.json",
			},
			FileSizes:   []string{"small", "medium", "large"},
			Concurrency: []int{1, 10, 20},
			Operations: []Operation{
				{
					Name:     "cache_read",
					Type:     "cache",
					Function: "readFromCache",
				},
				{
					Name:     "cache_write",
					Type:     "cache",
					Function: "writeToCache",
				},
				{
					Name:     "cache_validate",
					Type:     "cache",
					Function: "validateCache",
				},
			},
			Tags:   []string{"cache", "io"},
			Weight: 0.7,
		},
		{
			Name:        "markdown_generation",
			Description: "Markdown generation from JSON documentation",
			FilePatterns: []string{
				".cache/**/documentation/*.json",
			},
			FileSizes:   []string{"medium", "large"},
			Concurrency: []int{1, 5, 10},
			Operations: []Operation{
				{
					Name:     "generate_markdown",
					Type:     "transform",
					Function: "generateMarkdown",
				},
				{
					Name:     "format_code_blocks",
					Type:     "transform",
					Function: "formatCodeBlocks",
				},
			},
			Tags:   []string{"markdown", "transform"},
			Weight: 0.9,
		},
		{
			Name:        "memory_intensive",
			Description: "Memory-intensive operations with large data sets",
			FilePatterns: []string{
				"testdata/json/xlarge/*.json",
			},
			FileSizes:   []string{"xlarge"},
			Concurrency: []int{1, 2, 5},
			Operations: []Operation{
				{
					Name:     "load_all_frameworks",
					Type:     "memory",
					Function: "loadAllFrameworks",
				},
				{
					Name:     "build_reference_index",
					Type:     "memory",
					Function: "buildReferenceIndex",
				},
			},
			Tags:   []string{"memory", "stress"},
			Weight: 1.5,
		},
		{
			Name:        "real_world_simulation",
			Description: "Simulate real-world appledocs usage patterns",
			FilePatterns: []string{
				".cache/**/technologies.json",
			},
			Concurrency: []int{10},
			Operations: []Operation{
				{
					Name:     "full_crawl_simulation",
					Type:     "integration",
					Function: "simulateFullCrawl",
					Parameters: map[string]interface{}{
						"max_depth":     3,
						"max_urls":      1000,
						"rate_limit":    10.0,
						"cache_enabled": true,
					},
				},
			},
			Tags:   []string{"integration", "real-world"},
			Weight: 2.0,
		},
	}
}

// AppledocsThresholds returns performance thresholds for appledocs
func AppledocsThresholds() map[string]Threshold {
	return map[string]Threshold{
		"json_parsing_small": {
			MaxDuration:      100 * time.Millisecond,
			MaxMemory:        10 * 1024 * 1024, // 10MB
			MaxAllocations:   1000,
			MaxGCPause:       10 * time.Millisecond,
			RegressionMargin: 0.1, // 10% margin
		},
		"json_parsing_large": {
			MaxDuration:      1 * time.Second,
			MaxMemory:        100 * 1024 * 1024, // 100MB
			MaxAllocations:   10000,
			MaxGCPause:       50 * time.Millisecond,
			RegressionMargin: 0.15, // 15% margin
		},
		"concurrent_crawling": {
			MaxDuration:      5 * time.Second,
			MaxMemory:        500 * 1024 * 1024, // 500MB
			MaxAllocations:   50000,
			MaxGCPause:       100 * time.Millisecond,
			RegressionMargin: 0.2, // 20% margin
		},
		"cache_operations": {
			MaxDuration:      50 * time.Millisecond,
			MaxMemory:        50 * 1024 * 1024, // 50MB
			MaxAllocations:   5000,
			MaxGCPause:       20 * time.Millisecond,
			RegressionMargin: 0.1, // 10% margin
		},
		"markdown_generation": {
			MaxDuration:      2 * time.Second,
			MaxMemory:        200 * 1024 * 1024, // 200MB
			MaxAllocations:   20000,
			MaxGCPause:       50 * time.Millisecond,
			RegressionMargin: 0.15, // 15% margin
		},
		"memory_intensive": {
			MaxDuration:      10 * time.Second,
			MaxMemory:        1024 * 1024 * 1024, // 1GB
			MaxAllocations:   100000,
			MaxGCPause:       200 * time.Millisecond,
			RegressionMargin: 0.25, // 25% margin
		},
		"real_world_simulation": {
			MaxDuration:      30 * time.Second,
			MaxMemory:        2048 * 1024 * 1024, // 2GB
			MaxAllocations:   500000,
			MaxGCPause:       500 * time.Millisecond,
			RegressionMargin: 0.3, // 30% margin
		},
	}
}

// DefaultAppledocsConfig returns default benchmark configuration for appledocs
func DefaultAppledocsConfig() Config {
	return Config{
		Iterations:       5,
		WarmupIterations: 2,
		Timeout:          5 * time.Minute,
		CPUProfile:       true,
		MemProfile:       true,
		BlockProfile:     false,
		TraceEnabled:     false,
		Scenarios:        AppledocsScenarios(),
		Thresholds:       AppledocsThresholds(),
		OutputDir:        filepath.Join("benchmark_results", time.Now().Format("2006-01-02_15-04-05")),
		ReportFormat:     "json",
		Visualization:    true,
		CIMode:           false,
		FailOnRegression: true,
		Platforms: []Platform{
			getCurrentPlatform(),
		},
	}
}

// CIAppledocsConfig returns CI-optimized benchmark configuration
func CIAppledocsConfig() Config {
	config := DefaultAppledocsConfig()
	config.CIMode = true
	config.Iterations = 3
	config.WarmupIterations = 1
	config.CPUProfile = false
	config.MemProfile = true
	config.TraceEnabled = false
	config.Visualization = false
	
	// Filter scenarios for CI - only critical ones
	ciScenarios := make([]Scenario, 0)
	for _, scenario := range config.Scenarios {
		if scenario.Weight >= 1.0 {
			ciScenarios = append(ciScenarios, scenario)
		}
	}
	config.Scenarios = ciScenarios
	
	return config
}

// LoadTestData loads test data for appledocs benchmarks
func LoadTestData(config Config) error {
	// Ensure test data directories exist
	testDirs := []string{
		"testdata/json/small",
		"testdata/json/medium",
		"testdata/json/large",
		"testdata/json/xlarge",
	}
	
	for _, dir := range testDirs {
		if err := ensureDir(dir); err != nil {
			return err
		}
	}
	
	// Generate sample test data if needed
	if err := generateTestDataIfNeeded(); err != nil {
		return err
	}
	
	return nil
}

func ensureDir(path string) error {
	return nil // Implementation would create directory
}

func generateTestDataIfNeeded() error {
	return nil // Implementation would generate test data
}