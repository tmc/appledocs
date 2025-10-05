// Package benchmark provides examples and usage patterns for the benchmarking framework
package benchmark

import (
	"context"
	"fmt"
	"log"
	"time"
)

// ExampleBasicUsage demonstrates basic benchmark usage
func ExampleBasicUsage() {
	// Create integration manager
	manager, err := NewIntegrationManager(".")
	if err != nil {
		log.Fatalf("Failed to create integration manager: %v", err)
	}
	defer manager.Cleanup()
	
	// Establish baseline (run once)
	fmt.Println("Establishing performance baseline...")
	ctx := context.Background()
	if err := manager.EstablishBaseline(ctx); err != nil {
		log.Fatalf("Failed to establish baseline: %v", err)
	}
	
	// Run CI benchmarks (for continuous integration)
	fmt.Println("Running CI benchmarks...")
	ciResult, err := manager.RunCIBenchmarks(ctx)
	if err != nil {
		log.Fatalf("CI benchmarks failed: %v", err)
	}
	
	if !ciResult.Success {
		fmt.Printf("❌ CI benchmarks failed: %s\n", ciResult.Error)
	} else {
		fmt.Printf("✅ CI benchmarks passed in %v\n", ciResult.Duration)
	}
}

// ExampleFullSuite demonstrates running the complete benchmark suite
func ExampleFullSuite() {
	manager, err := NewIntegrationManager(".")
	if err != nil {
		log.Fatalf("Failed to create integration manager: %v", err)
	}
	defer manager.Cleanup()
	
	fmt.Println("Running full benchmark suite...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()
	
	result, err := manager.RunFullBenchmarkSuite(ctx)
	if err != nil {
		log.Fatalf("Full benchmark suite failed: %v", err)
	}
	
	// Print summary
	fmt.Printf("Benchmark suite completed in %v\n", result.Duration)
	fmt.Printf("Phases executed: %d\n", len(result.Phases))
	fmt.Printf("Comparisons generated: %d\n", len(result.Comparisons))
	fmt.Printf("Validation reports: %d\n", len(result.Validations))
	
	// Print phase results
	for phase, phaseResult := range result.Phases {
		fmt.Printf("Phase %s: %d results, duration: %v\n", 
			phase, len(phaseResult.Results), phaseResult.Duration)
	}
}

// ExampleCustomConfiguration demonstrates custom benchmark configuration
func ExampleCustomConfiguration() {
	// Create custom config
	config := Config{
		Iterations:       10,
		WarmupIterations: 3,
		Timeout:          10 * time.Minute,
		CPUProfile:       true,
		MemProfile:       true,
		ReportFormat:     "html",
		Visualization:    true,
		OutputDir:        "custom-benchmark-results",
		
		// Custom scenarios
		Scenarios: []Scenario{
			{
				Name:         "custom_json_test",
				Description:  "Custom JSON parsing test",
				FilePatterns: []string{"test-data/*.json"},
				FileSizes:    []string{"small", "medium"},
				Concurrency:  []int{1, 5},
				Operations: []Operation{
					{Name: "parse", Type: "parse"},
					{Name: "extract", Type: "extract"},
				},
				Weight: 1.0,
			},
		},
		
		// Custom thresholds
		Thresholds: map[string]Threshold{
			"custom_json_test_parse": {
				MaxDuration:      200 * time.Millisecond,
				MaxMemory:        20 * 1024 * 1024,
				RegressionMargin: 0.10,
			},
		},
	}
	
	// Create framework with custom config
	framework, err := NewFramework(config)
	if err != nil {
		log.Fatalf("Failed to create framework: %v", err)
	}
	
	// Run custom benchmark
	ctx := context.Background()
	runner := &BaselineRunner{}
	
	if err := framework.RunPhase(ctx, PhaseBaseline, runner); err != nil {
		log.Fatalf("Failed to run phase: %v", err)
	}
	
	fmt.Println("Custom benchmark completed successfully")
}

// ExampleRegressionDetection demonstrates regression detection
func ExampleRegressionDetection() {
	manager, err := NewIntegrationManager(".")
	if err != nil {
		log.Fatalf("Failed to create integration manager: %v", err)
	}
	defer manager.Cleanup()
	
	ctx := context.Background()
	
	// Simulate before and after refactoring
	fmt.Println("Running baseline (before refactoring)...")
	baselineRunner := &BaselineRunner{}
	if err := manager.framework.RunPhase(ctx, PhaseBaseline, baselineRunner); err != nil {
		log.Fatalf("Baseline failed: %v", err)
	}
	
	fmt.Println("Running optimized version (after refactoring)...")
	optimizedRunner := &OptimizedJSONRunner{}
	if err := manager.framework.RunPhase(ctx, PhaseJSONOptimize, optimizedRunner); err != nil {
		log.Fatalf("Optimized phase failed: %v", err)
	}
	
	// Load results
	baselineResults, err := manager.framework.storage.LoadResults(PhaseBaseline)
	if err != nil {
		log.Fatalf("Failed to load baseline results: %v", err)
	}
	
	optimizedResults, err := manager.framework.storage.LoadResults(PhaseJSONOptimize)
	if err != nil {
		log.Fatalf("Failed to load optimized results: %v", err)
	}
	
	// Compare results
	comparison := manager.framework.comparator.ComparePhases(
		PhaseBaseline, PhaseJSONOptimize, baselineResults, optimizedResults)
	
	// Check for regressions
	regressions := comparison.GetSignificantRegressions()
	improvements := comparison.GetSignificantImprovements()
	
	fmt.Printf("Comparison results:\n")
	fmt.Printf("- Significant regressions: %d\n", len(regressions))
	fmt.Printf("- Significant improvements: %d\n", len(improvements))
	fmt.Printf("- Average duration change: %.2f%%\n", comparison.Summary.AvgDurationChange*100)
	fmt.Printf("- Average memory change: %.2f%%\n", comparison.Summary.AvgMemoryChange*100)
	
	if len(regressions) > 0 {
		fmt.Println("\n❌ Regressions detected:")
		for _, reg := range regressions {
			fmt.Printf("  - %s.%s: %.2f%% slower (p=%.4f)\n", 
				reg.Scenario, reg.Operation, 
				reg.Changes.Duration.Relative*100, reg.Significance.PValue)
		}
	}
	
	if len(improvements) > 0 {
		fmt.Println("\n✅ Improvements detected:")
		for _, imp := range improvements {
			fmt.Printf("  - %s.%s: %.2f%% faster (p=%.4f)\n", 
				imp.Scenario, imp.Operation, 
				-imp.Changes.Duration.Relative*100, imp.Significance.PValue)
		}
	}
}

// ExampleMemoryProfiling demonstrates memory profiling and leak detection
func ExampleMemoryProfiling() {
	config := DefaultConfig()
	config.MemProfile = true
	config.Iterations = 5
	
	framework, err := NewFramework(config)
	if err != nil {
		log.Fatalf("Failed to create framework: %v", err)
	}
	
	ctx := context.Background()
	runner := &BaselineRunner{}
	
	fmt.Println("Running memory profiling benchmark...")
	if err := framework.RunPhase(ctx, PhaseBaseline, runner); err != nil {
		log.Fatalf("Memory profiling failed: %v", err)
	}
	
	// Load results and check for memory issues
	results, err := framework.storage.LoadResults(PhaseBaseline)
	if err != nil {
		log.Fatalf("Failed to load results: %v", err)
	}
	
	// Analyze memory patterns
	for _, result := range results {
		memStats := result.Memory
		
		// Check for potential memory leaks
		heapRetention := float64(memStats.HeapInuse) / float64(memStats.HeapInuse+memStats.HeapReleased)
		if heapRetention > 0.8 {
			fmt.Printf("⚠️  Potential memory leak in %s.%s: %.2f%% heap retention\n",
				result.Scenario, result.Operation, heapRetention*100)
		}
		
		// Check for excessive allocations
		if memStats.Allocations > 100000 {
			fmt.Printf("⚠️  High allocation count in %s.%s: %d allocations\n",
				result.Scenario, result.Operation, memStats.Allocations)
		}
		
		// Check for long GC pauses
		if memStats.GCPauseMax > 50*time.Millisecond {
			fmt.Printf("⚠️  Long GC pause in %s.%s: %v\n",
				result.Scenario, result.Operation, memStats.GCPauseMax)
		}
	}
	
	fmt.Println("Memory profiling completed. Check profiles in:", framework.profiler.GetProfileDir())
}

// ExamplePerformanceTrends demonstrates performance trend analysis
func ExamplePerformanceTrends() {
	manager, err := NewIntegrationManager(".")
	if err != nil {
		log.Fatalf("Failed to create integration manager: %v", err)
	}
	defer manager.Cleanup()
	
	// Simulate multiple benchmark runs over time
	phases := []Phase{PhaseBaseline, PhaseJSONOptimize, PhaseStreaming}
	ctx := context.Background()
	
	fmt.Println("Running multiple phases for trend analysis...")
	
	phaseResults := make(map[Phase][]Result)
	for _, phase := range phases {
		runner, _ := manager.createRunner(phase)
		runner.Setup()
		
		if err := manager.framework.RunPhase(ctx, phase, runner); err != nil {
			log.Printf("Phase %s failed: %v", phase, err)
			continue
		}
		
		results, _ := manager.framework.storage.LoadResults(phase)
		phaseResults[phase] = results
		
		runner.Teardown()
	}
	
	// Analyze trends
	comparison := manager.framework.comparator.CompareMultiplePhases(phases, phaseResults)
	
	fmt.Println("\nPerformance trends:")
	for key, trend := range comparison.Trends {
		fmt.Printf("- %s: %s trend (slope: %.2f, confidence: %s)\n",
			key, trend.Direction, trend.Slope, trend.Confidence)
	}
}

// ExampleCIIntegration demonstrates CI/CD integration
func ExampleCIIntegration() {
	// This example shows how to integrate with CI/CD systems
	
	manager, err := NewIntegrationManager(".")
	if err != nil {
		log.Fatalf("Failed to create integration manager: %v", err)
	}
	defer manager.Cleanup()
	
	// Configure for CI environment
	manager.config.CIMode = true
	manager.config.FailOnRegression = true
	manager.config.Iterations = 3 // Faster for CI
	manager.config.ReportFormat = "json"
	manager.config.Visualization = false
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	
	fmt.Println("Running CI benchmarks...")
	result, err := manager.RunCIBenchmarks(ctx)
	if err != nil {
		fmt.Printf("CI benchmarks failed: %v\n", err)
		// In CI, you would exit with non-zero code here
		return
	}
	
	if !result.Success {
		fmt.Printf("❌ Performance regression detected: %s\n", result.Error)
		// In CI, you would exit with non-zero code here
		return
	}
	
	fmt.Printf("✅ CI benchmarks passed in %v\n", result.Duration)
	
	// Generate CI-friendly output
	if result.Comparison != nil {
		summary := result.Comparison.Summary
		fmt.Printf("Performance changes: %.2f%% duration, %.2f%% memory\n",
			summary.AvgDurationChange*100, summary.AvgMemoryChange*100)
	}
}

// ExampleReportGeneration demonstrates report generation
func ExampleReportGeneration() {
	manager, err := NewIntegrationManager(".")
	if err != nil {
		log.Fatalf("Failed to create integration manager: %v", err)
	}
	defer manager.Cleanup()
	
	// Run a quick benchmark
	ctx := context.Background()
	runner := &BaselineRunner{}
	if err := manager.framework.RunPhase(ctx, PhaseBaseline, runner); err != nil {
		log.Fatalf("Benchmark failed: %v", err)
	}
	
	// Load results
	results, err := manager.framework.storage.LoadResults(PhaseBaseline)
	if err != nil {
		log.Fatalf("Failed to load results: %v", err)
	}
	
	// Generate validation report
	validation := manager.framework.validator.ValidatePhase(PhaseBaseline, results, manager.framework.baselines)
	
	// Generate phase report
	phaseReport, err := manager.framework.reporter.GeneratePhaseReport(PhaseBaseline, results, validation)
	if err != nil {
		log.Fatalf("Failed to generate phase report: %v", err)
	}
	
	// Create full report data
	reportData := &ReportData{
		GeneratedAt:  time.Now(),
		Platform:     getCurrentPlatform(),
		GitCommit:    manager.getCurrentGitCommit(),
		Config:       manager.config,
		PhaseResults: []*PhaseResult{phaseReport},
		Validations:  []*ValidationReport{validation},
	}
	
	// Generate summary
	summary := manager.framework.reporter.GenerateSummaryReport(
		reportData.PhaseResults, nil, reportData.Validations)
	reportData.Summary = *summary
	
	// Generate recommendations
	reportData.Recommendations = manager.framework.reporter.GenerateRecommendations(reportData)
	
	// Generate report
	if err := manager.framework.reporter.GenerateReport(reportData); err != nil {
		log.Fatalf("Failed to generate report: %v", err)
	}
	
	fmt.Printf("Report generated successfully in: %s\n", manager.config.OutputDir)
	fmt.Printf("Overall score: %.1f/100\n", reportData.Summary.OverallScore)
	fmt.Printf("Performance score: %.1f/100\n", reportData.Summary.PerformanceScore)
	fmt.Printf("Stability score: %.1f/100\n", reportData.Summary.StabilityScore)
}

// ExampleCustomRunner demonstrates creating a custom benchmark runner
func ExampleCustomRunner() {
	// Define a custom runner for testing specific optimizations
	type CustomOptimizationRunner struct {
		optimizationEnabled bool
	}
	
	runner := &CustomOptimizationRunner{optimizationEnabled: true}
	
	// Implement the BenchmarkRunner interface
	runner.Name = func() string { return "custom-optimization" }
	runner.Setup = func() error {
		fmt.Println("Setting up custom optimization...")
		return nil
	}
	runner.Teardown = func() error {
		fmt.Println("Cleaning up custom optimization...")
		return nil
	}
	runner.Run = func(ctx context.Context, op Operation, data []byte) (interface{}, error) {
		// Implement custom optimization logic
		fmt.Printf("Running optimized %s operation...\n", op.Type)
		
		switch op.Type {
		case "parse":
			// Custom parsing logic
			return map[string]interface{}{"optimized": true}, nil
		case "extract":
			// Custom extraction logic
			return []string{"optimized-url"}, nil
		default:
			return nil, fmt.Errorf("unsupported operation: %s", op.Type)
		}
	}
	
	// Use the custom runner
	config := DefaultConfig()
	config.Iterations = 3
	
	framework, err := NewFramework(config)
	if err != nil {
		log.Fatalf("Failed to create framework: %v", err)
	}
	
	ctx := context.Background()
	if err := framework.RunPhase(ctx, Phase("custom"), runner); err != nil {
		log.Fatalf("Custom benchmark failed: %v", err)
	}
	
	fmt.Println("Custom benchmark completed successfully")
}

// Helper function to create a BenchmarkRunner with function implementations
func createFunctionalRunner(name string, 
	setupFunc func() error,
	teardownFunc func() error,
	runFunc func(context.Context, Operation, []byte) (interface{}, error)) BenchmarkRunner {
	
	return &functionalRunner{
		name:         name,
		setupFunc:    setupFunc,
		teardownFunc: teardownFunc,
		runFunc:      runFunc,
	}
}

// functionalRunner implements BenchmarkRunner using function fields
type functionalRunner struct {
	name         string
	setupFunc    func() error
	teardownFunc func() error
	runFunc      func(context.Context, Operation, []byte) (interface{}, error)
}

func (r *functionalRunner) Name() string { return r.name }
func (r *functionalRunner) Setup() error { return r.setupFunc() }
func (r *functionalRunner) Teardown() error { return r.teardownFunc() }
func (r *functionalRunner) Run(ctx context.Context, op Operation, data []byte) (interface{}, error) {
	return r.runFunc(ctx, op, data)
}

// ExampleAdvancedConfiguration demonstrates advanced configuration options
func ExampleAdvancedConfiguration() {
	config := Config{
		// Performance settings
		Iterations:       10,
		WarmupIterations: 3,
		Timeout:          15 * time.Minute,
		
		// Profiling settings
		CPUProfile:   true,
		MemProfile:   true,
		BlockProfile: true,
		TraceEnabled: false,
		
		// Output settings
		OutputDir:     "advanced-benchmark-results",
		ReportFormat:  "all", // Generate JSON, HTML, and Markdown
		Visualization: true,
		
		// CI/CD settings
		CIMode:           false,
		FailOnRegression: true,
		
		// Advanced scenarios with different file sizes and concurrency levels
		Scenarios: []Scenario{
			{
				Name:         "comprehensive_json_parsing",
				Description:  "Comprehensive JSON parsing across all file sizes",
				FilePatterns: []string{"output/**/*.json", "test-data/**/*.json"},
				FileSizes:    []string{"small", "medium", "large", "xlarge"},
				Concurrency:  []int{1, 2, 5, 10, 20},
				Operations: []Operation{
					{Name: "parse", Type: "parse"},
					{Name: "extract_urls", Type: "extract"},
					{Name: "validate", Type: "validate"},
					{Name: "transform", Type: "transform"},
				},
				Tags:   []string{"json", "performance", "comprehensive"},
				Weight: 3.0,
			},
			{
				Name:         "memory_stress_test",
				Description:  "Memory stress testing with large files",
				FilePatterns: []string{"output/**/large*.json"},
				FileSizes:    []string{"large", "xlarge"},
				Concurrency:  []int{1, 5},
				Operations: []Operation{
					{Name: "memory_intensive_parse", Type: "parse"},
					{Name: "streaming_parse", Type: "stream"},
				},
				Tags:   []string{"memory", "stress", "large-files"},
				Weight: 2.0,
			},
		},
		
		// Strict performance thresholds
		Thresholds: map[string]Threshold{
			"comprehensive_json_parsing_parse": {
				MaxDuration:      1 * time.Second,
				MaxMemory:        100 * 1024 * 1024, // 100MB
				MaxAllocations:   50000,
				MaxGCPause:       20 * time.Millisecond,
				RegressionMargin: 0.05, // Only 5% regression allowed
			},
			"memory_stress_test_parse": {
				MaxDuration:      5 * time.Second,
				MaxMemory:        500 * 1024 * 1024, // 500MB
				MaxAllocations:   100000,
				MaxGCPause:       100 * time.Millisecond,
				RegressionMargin: 0.10, // 10% regression allowed
			},
		},
		
		// Multi-platform support
		Platforms: []Platform{
			{
				OS:        "darwin",
				Arch:      "amd64",
				CPUCores:  8,
				Memory:    16 * 1024 * 1024 * 1024, // 16GB
				GoVersion: "go1.21.0",
			},
			{
				OS:        "linux",
				Arch:      "amd64",
				CPUCores:  4,
				Memory:    8 * 1024 * 1024 * 1024, // 8GB
				GoVersion: "go1.21.0",
			},
		},
	}
	
	framework, err := NewFramework(config)
	if err != nil {
		log.Fatalf("Failed to create advanced framework: %v", err)
	}
	
	fmt.Println("Advanced benchmark framework created successfully")
	fmt.Printf("Scenarios: %d\n", len(config.Scenarios))
	fmt.Printf("Platforms: %d\n", len(config.Platforms))
	fmt.Printf("Output directory: %s\n", config.OutputDir)
}

// Main function to run all examples
func RunAllExamples() {
	fmt.Println("=== AppLeDocs Benchmark Framework Examples ===\n")
	
	examples := []struct {
		name string
		fn   func()
	}{
		{"Basic Usage", ExampleBasicUsage},
		{"Custom Configuration", ExampleCustomConfiguration},
		{"Regression Detection", ExampleRegressionDetection},
		{"Memory Profiling", ExampleMemoryProfiling},
		{"Performance Trends", ExamplePerformanceTrends},
		{"CI Integration", ExampleCIIntegration},
		{"Report Generation", ExampleReportGeneration},
		{"Advanced Configuration", ExampleAdvancedConfiguration},
	}
	
	for _, example := range examples {
		fmt.Printf("Running example: %s\n", example.name)
		fmt.Println(strings.Repeat("-", 50))
		
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Example %s panicked: %v\n", example.name, r)
				}
			}()
			
			example.fn()
		}()
		
		fmt.Printf("\nCompleted: %s\n", example.name)
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println()
	}
}