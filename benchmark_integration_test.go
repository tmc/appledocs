// Package main provides integration tests for the benchmarking framework
package main

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/tmc/appledocs/benchmark"
)

// TestBenchmarkFrameworkIntegration tests the complete benchmark framework
func TestBenchmarkFrameworkIntegration(t *testing.T) {
	// Skip in short mode
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Create test configuration
	config := benchmark.Config{
		Iterations:       3,
		WarmupIterations: 1,
		Timeout:          30 * time.Second,
		CPUProfile:       false,
		MemProfile:       true,
		OutputDir:        filepath.Join(os.TempDir(), "appledocs_benchmark_test"),
		ReportFormat:     "json",
		CIMode:           true,
		FailOnRegression: false,
		Scenarios: []benchmark.Scenario{
			{
				Name:        "test_json_parsing",
				Description: "Test JSON parsing performance",
				FileSizes:   []string{"small"},
				Concurrency: []int{1},
				Operations: []benchmark.Operation{
					{
						Name:     "parse_json",
						Type:     "parse",
						Function: "json.Unmarshal",
					},
				},
				Weight: 1.0,
			},
		},
		Thresholds: map[string]benchmark.Threshold{
			"test_json_parsing": {
				MaxDuration:      100 * time.Millisecond,
				MaxMemory:        10 * 1024 * 1024,
				MaxAllocations:   1000,
				RegressionMargin: 0.2,
			},
		},
	}

	// Create framework
	framework, err := benchmark.NewFramework(config)
	if err != nil {
		t.Fatalf("Failed to create framework: %v", err)
	}

	// Create test context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// Run baseline phase
	t.Run("BaselinePhase", func(t *testing.T) {
		runner := benchmark.NewAppledocsBaselineRunner()
		if err := runner.Setup(); err != nil {
			t.Fatalf("Failed to setup runner: %v", err)
		}
		defer runner.Teardown()

		if err := framework.RunPhase(ctx, benchmark.PhaseBaseline, runner); err != nil {
			t.Errorf("Baseline phase failed: %v", err)
		}
	})

	// Run optimized phase
	t.Run("OptimizedPhase", func(t *testing.T) {
		runner := benchmark.NewAppledocsOptimizedRunner()
		if err := runner.Setup(); err != nil {
			t.Fatalf("Failed to setup runner: %v", err)
		}
		defer runner.Teardown()

		if err := framework.RunPhase(ctx, benchmark.PhaseJSONOptimize, runner); err != nil {
			t.Errorf("Optimized phase failed: %v", err)
		}
	})

	// Compare results
	t.Run("CompareResults", func(t *testing.T) {
		comparison := framework.ComparePhases(benchmark.PhaseBaseline, benchmark.PhaseJSONOptimize)
		if comparison == nil {
			t.Fatal("Failed to get comparison")
		}

		// Check for improvements
		if len(comparison.Improvements) == 0 && len(comparison.Regressions) == 0 {
			t.Log("No significant performance changes detected")
		}

		// Log results
		for _, improvement := range comparison.Improvements {
			t.Logf("Improvement: %s", improvement)
		}
		for _, regression := range comparison.Regressions {
			t.Logf("Regression: %s", regression)
		}
	})

	// Generate report
	t.Run("GenerateReport", func(t *testing.T) {
		report, err := framework.GenerateReport()
		if err != nil {
			t.Errorf("Failed to generate report: %v", err)
		}

		// Verify report structure
		if report.Summary == nil {
			t.Error("Report missing summary")
		}
		if len(report.PhaseResults) == 0 {
			t.Error("Report missing phase results")
		}
	})

	// Cleanup
	os.RemoveAll(config.OutputDir)
}

// TestAppledocsScenarios tests appledocs-specific scenarios
func TestAppledocsScenarios(t *testing.T) {
	scenarios := benchmark.AppledocsScenarios()
	
	// Verify we have all expected scenarios
	expectedScenarios := map[string]bool{
		"json_parsing_small":    false,
		"json_parsing_large":    false,
		"concurrent_crawling":   false,
		"cache_operations":      false,
		"markdown_generation":   false,
		"memory_intensive":      false,
		"real_world_simulation": false,
	}

	for _, scenario := range scenarios {
		if _, expected := expectedScenarios[scenario.Name]; expected {
			expectedScenarios[scenario.Name] = true
		} else {
			t.Errorf("Unexpected scenario: %s", scenario.Name)
		}

		// Validate scenario
		if scenario.Description == "" {
			t.Errorf("Scenario %s missing description", scenario.Name)
		}
		if len(scenario.Operations) == 0 {
			t.Errorf("Scenario %s has no operations", scenario.Name)
		}
		if scenario.Weight <= 0 {
			t.Errorf("Scenario %s has invalid weight: %f", scenario.Name, scenario.Weight)
		}
	}

	// Check all expected scenarios were found
	for name, found := range expectedScenarios {
		if !found {
			t.Errorf("Expected scenario not found: %s", name)
		}
	}
}

// TestRunnerImplementations tests different runner implementations
func TestRunnerImplementations(t *testing.T) {
	testData := []byte(`{
		"title": "Test Document",
		"abstract": [{"type": "text", "text": "Test abstract"}],
		"metadata": {"platforms": [{"name": "iOS"}]},
		"references": {
			"ref1": {"url": "test1.json"},
			"ref2": {"url": "test2.json"}
		}
	}`)

	runners := []benchmark.BenchmarkRunner{
		benchmark.NewAppledocsBaselineRunner(),
		benchmark.NewAppledocsOptimizedRunner(),
		benchmark.NewAppledocsStreamingRunner(),
		benchmark.NewAppledocsConcurrentRunner(),
		benchmark.NewAppledocsMemoryOptimizedRunner(),
	}

	ctx := context.Background()

	for _, runner := range runners {
		t.Run(runner.Name(), func(t *testing.T) {
			// Setup
			if err := runner.Setup(); err != nil {
				t.Fatalf("Setup failed: %v", err)
			}
			defer runner.Teardown()

			// Test JSON parsing
			t.Run("ParseJSON", func(t *testing.T) {
				op := benchmark.Operation{
					Name:     "parse_json",
					Type:     "parse",
					Function: "json.Unmarshal",
				}

				result, err := runner.Run(ctx, op, testData)
				if err != nil {
					t.Errorf("Parse failed: %v", err)
				}

				// Verify result is a map
				if _, ok := result.(map[string]interface{}); !ok {
					t.Errorf("Expected map result, got %T", result)
				}
			})

			// Test URL extraction
			t.Run("ExtractURLs", func(t *testing.T) {
				op := benchmark.Operation{
					Name:     "extract_urls",
					Type:     "extract",
					Function: "extractJSONURLs",
				}

				result, err := runner.Run(ctx, op, testData)
				if err != nil {
					t.Errorf("Extract URLs failed: %v", err)
				}

				// Verify result is a slice of strings
				urls, ok := result.([]string)
				if !ok {
					t.Errorf("Expected []string result, got %T", result)
				}

				// Should find 2 URLs
				if len(urls) != 2 {
					t.Errorf("Expected 2 URLs, found %d", len(urls))
				}
			})

			// Test markdown generation
			t.Run("GenerateMarkdown", func(t *testing.T) {
				op := benchmark.Operation{
					Name:     "generate_markdown",
					Type:     "transform",
					Function: "generateMarkdown",
				}

				result, err := runner.Run(ctx, op, testData)
				if err != nil {
					t.Errorf("Generate markdown failed: %v", err)
				}

				// Verify result is a string
				markdown, ok := result.(string)
				if !ok {
					t.Errorf("Expected string result, got %T", result)
				}

				// Should contain the title
				if len(markdown) == 0 {
					t.Error("Generated markdown is empty")
				}
			})
		})
	}
}

// TestBenchmarkValidation tests benchmark validation
func TestBenchmarkValidation(t *testing.T) {
	config := benchmark.DefaultAppledocsConfig()
	validator := benchmark.NewValidator(config)

	// Create mock results
	results := []benchmark.Result{
		{
			Phase:     benchmark.PhaseBaseline,
			Scenario:  "json_parsing_small",
			Operation: "parse_json",
			Duration:  50 * time.Millisecond,
			Memory: benchmark.MemoryStats{
				TotalAlloc:  5 * 1024 * 1024,
				Allocations: 500,
			},
		},
		{
			Phase:     benchmark.PhaseJSONOptimize,
			Scenario:  "json_parsing_small",
			Operation: "parse_json",
			Duration:  40 * time.Millisecond,
			Memory: benchmark.MemoryStats{
				TotalAlloc:  4 * 1024 * 1024,
				Allocations: 400,
			},
		},
	}

	// Validate results
	report := validator.ValidateResults(results)
	
	// Should pass validation
	if report.HasFailures() {
		t.Errorf("Validation failed: %v", report.Summary())
	}

	// Should detect improvement
	if !report.HasImprovements() {
		t.Error("Expected improvements to be detected")
	}
}

// TestCIIntegration tests CI mode configuration
func TestCIIntegration(t *testing.T) {
	config := benchmark.CIAppledocsConfig()

	// Verify CI optimizations
	if !config.CIMode {
		t.Error("CI mode not enabled")
	}
	if config.Iterations > 3 {
		t.Errorf("CI mode should have fewer iterations, got %d", config.Iterations)
	}
	if config.CPUProfile {
		t.Error("CPU profiling should be disabled in CI mode")
	}

	// Verify only critical scenarios are included
	for _, scenario := range config.Scenarios {
		if scenario.Weight < 1.0 {
			t.Errorf("CI mode should only include critical scenarios, found %s with weight %f", 
				scenario.Name, scenario.Weight)
		}
	}
}

// TestMemoryLeakDetection tests memory leak detection
func TestMemoryLeakDetection(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping memory leak test in short mode")
	}

	config := benchmark.Config{
		Iterations:       5,
		WarmupIterations: 2,
		MemProfile:       true,
		OutputDir:        filepath.Join(os.TempDir(), "appledocs_memleak_test"),
		Scenarios: []benchmark.Scenario{
			{
				Name:        "memory_leak_test",
				Description: "Test for memory leaks",
				Operations: []benchmark.Operation{
					{
						Name:     "allocate_memory",
						Type:     "memory",
						Function: "allocateMemory",
					},
				},
			},
		},
	}

	framework, err := benchmark.NewFramework(config)
	if err != nil {
		t.Fatalf("Failed to create framework: %v", err)
	}

	// Run test with memory leak simulation
	runner := &memoryLeakRunner{
		AppledocsBaselineRunner: benchmark.NewAppledocsBaselineRunner(),
	}

	ctx := context.Background()
	if err := framework.RunPhase(ctx, benchmark.PhaseBaseline, runner); err != nil {
		t.Errorf("Memory leak test failed: %v", err)
	}

	// Analyze memory profile
	analysis := framework.AnalyzeMemoryGrowth(benchmark.PhaseBaseline)
	if analysis == nil {
		t.Fatal("Failed to analyze memory growth")
	}

	// Check if memory leak was detected
	if !analysis.HasLeak {
		t.Error("Expected memory leak to be detected")
	}

	// Cleanup
	os.RemoveAll(config.OutputDir)
}

// memoryLeakRunner simulates a memory leak
type memoryLeakRunner struct {
	*benchmark.AppledocsBaselineRunner
	leaked [][]byte
}

func (r *memoryLeakRunner) Run(ctx context.Context, op benchmark.Operation, data []byte) (interface{}, error) {
	if op.Function == "allocateMemory" {
		// Simulate memory leak by holding references
		leak := make([]byte, 1024*1024) // 1MB
		r.leaked = append(r.leaked, leak)
		return len(r.leaked), nil
	}
	return r.AppledocsBaselineRunner.Run(ctx, op, data)
}

// BenchmarkJSONParsing provides a real Go benchmark
func BenchmarkJSONParsing(b *testing.B) {
	// Create test data
	testData := generateTestJSON(1024 * 1024) // 1MB

	runners := map[string]benchmark.BenchmarkRunner{
		"Baseline":  benchmark.NewAppledocsBaselineRunner(),
		"Optimized": benchmark.NewAppledocsOptimizedRunner(),
		"Streaming": benchmark.NewAppledocsStreamingRunner(),
	}

	op := benchmark.Operation{
		Name:     "parse_json",
		Type:     "parse",
		Function: "json.Unmarshal",
	}

	ctx := context.Background()

	for name, runner := range runners {
		b.Run(name, func(b *testing.B) {
			runner.Setup()
			defer runner.Teardown()

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				_, err := runner.Run(ctx, op, testData)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// generateTestJSON generates test JSON data
func generateTestJSON(size int) []byte {
	doc := map[string]interface{}{
		"metadata": map[string]interface{}{
			"title": "Test Document",
			"platforms": []map[string]interface{}{
				{"name": "iOS", "introducedAt": "14.0"},
			},
		},
		"references": make(map[string]interface{}),
	}

	// Add references until we reach target size
	refCount := 0
	for {
		refID := "ref" + string(rune(refCount))
		doc["references"].(map[string]interface{})[refID] = map[string]interface{}{
			"url":   "test" + string(rune(refCount)) + ".json",
			"title": "Reference " + string(rune(refCount)),
		}
		refCount++

		data, _ := json.Marshal(doc)
		if len(data) >= size {
			return data
		}
	}
}