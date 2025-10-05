// Package benchmark provides integration with the main appledocs application
package benchmark

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// IntegrationManager manages benchmark integration with appledocs
type IntegrationManager struct {
	framework  *BenchmarkFramework
	config     Config
	workingDir string
}

// NewIntegrationManager creates a new integration manager
func NewIntegrationManager(workingDir string) (*IntegrationManager, error) {
	config := DefaultConfig()
	config.OutputDir = filepath.Join(workingDir, "benchmark-results")
	
	framework, err := NewFramework(config)
	if err != nil {
		return nil, fmt.Errorf("create framework: %w", err)
	}
	
	return &IntegrationManager{
		framework:  framework,
		config:     config,
		workingDir: workingDir,
	}, nil
}

// DefaultConfig returns a default benchmark configuration for appledocs
func DefaultConfig() Config {
	return Config{
		Iterations:       5,
		WarmupIterations: 2,
		Timeout:          5 * time.Minute,
		CPUProfile:       true,
		MemProfile:       true,
		BlockProfile:     false,
		TraceEnabled:     false,
		OutputDir:        "benchmark-results",
		ReportFormat:     "html",
		Visualization:    true,
		CIMode:           false,
		FailOnRegression: false,
		
		Scenarios: []Scenario{
			{
				Name:         "json_parsing_small",
				Description:  "JSON parsing performance for small files",
				FilePatterns: []string{"output/**/*technologies*.json", "output/**/small_*.json"},
				FileSizes:    []string{"small"},
				Concurrency:  []int{1, 5, 10},
				Operations: []Operation{
					{Name: "parse", Type: "parse", Function: "parseDocument"},
					{Name: "extract_urls", Type: "extract", Function: "extractURLs"},
				},
				Weight: 1.0,
			},
			{
				Name:         "json_parsing_medium",
				Description:  "JSON parsing performance for medium files",
				FilePatterns: []string{"output/**/*View*.json", "output/**/medium_*.json"},
				FileSizes:    []string{"medium"},
				Concurrency:  []int{1, 5, 10},
				Operations: []Operation{
					{Name: "parse", Type: "parse", Function: "parseDocument"},
					{Name: "extract_urls", Type: "extract", Function: "extractURLs"},
					{Name: "transform_markdown", Type: "transform", Function: "transformToMarkdown"},
				},
				Weight: 2.0,
			},
			{
				Name:         "json_parsing_large",
				Description:  "JSON parsing performance for large files",
				FilePatterns: []string{"output/**/*Implementations*.json", "output/**/large_*.json"},
				FileSizes:    []string{"large"},
				Concurrency:  []int{1, 2, 5},
				Operations: []Operation{
					{Name: "parse", Type: "parse", Function: "parseDocument"},
					{Name: "extract_urls", Type: "extract", Function: "extractURLs"},
					{Name: "transform_markdown", Type: "transform", Function: "transformToMarkdown"},
				},
				Weight: 3.0,
			},
			{
				Name:         "concurrent_processing",
				Description:  "Concurrent processing performance",
				FilePatterns: []string{"output/**/*.json"},
				FileSizes:    []string{"small", "medium"},
				Concurrency:  []int{1, 5, 10, 20},
				Operations: []Operation{
					{Name: "concurrent_parse", Type: "parse", Function: "concurrentParse"},
					{Name: "concurrent_extract", Type: "extract", Function: "concurrentExtract"},
				},
				Weight: 2.0,
			},
			{
				Name:         "cache_performance",
				Description:  "HTTP cache performance",
				FilePatterns: []string{".cache/**/*.json"},
				FileSizes:    []string{"small", "medium", "large"},
				Concurrency:  []int{1, 5},
				Operations: []Operation{
					{Name: "cache_read", Type: "read", Function: "readFromCache"},
					{Name: "cache_write", Type: "write", Function: "writeToCache"},
				},
				Weight: 1.5,
			},
		},
		
		Thresholds: map[string]Threshold{
			"json_parsing_small_parse": {
				MaxDuration:      100 * time.Millisecond,
				MaxMemory:        10 * 1024 * 1024, // 10MB
				MaxAllocations:   10000,
				MaxGCPause:       10 * time.Millisecond,
				RegressionMargin: 0.15, // 15%
			},
			"json_parsing_medium_parse": {
				MaxDuration:      500 * time.Millisecond,
				MaxMemory:        50 * 1024 * 1024, // 50MB
				MaxAllocations:   50000,
				MaxGCPause:       25 * time.Millisecond,
				RegressionMargin: 0.20, // 20%
			},
			"json_parsing_large_parse": {
				MaxDuration:      2 * time.Second,
				MaxMemory:        200 * 1024 * 1024, // 200MB
				MaxAllocations:   200000,
				MaxGCPause:       50 * time.Millisecond,
				RegressionMargin: 0.25, // 25%
			},
		},
		
		Platforms: []Platform{
			{
				OS:        runtime.GOOS,
				Arch:      runtime.GOARCH,
				CPUCores:  runtime.NumCPU(),
				GoVersion: runtime.Version(),
			},
		},
	}
}

// RunFullBenchmarkSuite runs the complete benchmark suite for refactoring validation
func (im *IntegrationManager) RunFullBenchmarkSuite(ctx context.Context) (*FullSuiteResult, error) {
	result := &FullSuiteResult{
		StartTime: time.Now(),
		Phases:    make(map[Phase]*PhaseExecutionResult),
	}
	
	// Define refactoring phases to test
	phases := []Phase{
		PhaseBaseline,
		PhaseJSONOptimize,
		PhaseStreaming,
		PhaseConcurrency,
		PhaseMemoryOpt,
		PhaseCacheOpt,
	}
	
	// Run each phase
	for _, phase := range phases {
		phaseResult, err := im.runPhase(ctx, phase)
		if err != nil {
			return nil, fmt.Errorf("run phase %s: %w", phase, err)
		}
		result.Phases[phase] = phaseResult
	}
	
	// Generate comparisons
	result.Comparisons = im.generateAllComparisons(result.Phases)
	
	// Generate validations
	result.Validations = im.generateAllValidations(result.Phases)
	
	// Generate comprehensive report
	reportData := im.buildReportData(result)
	if err := im.framework.reporter.GenerateReport(reportData); err != nil {
		return nil, fmt.Errorf("generate report: %w", err)
	}
	
	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	
	return result, nil
}

// runPhase executes a single benchmark phase
func (im *IntegrationManager) runPhase(ctx context.Context, phase Phase) (*PhaseExecutionResult, error) {
	result := &PhaseExecutionResult{
		Phase:     phase,
		StartTime: time.Now(),
	}
	
	// Create runner for this phase
	runner, err := im.createRunner(phase)
	if err != nil {
		return nil, fmt.Errorf("create runner: %w", err)
	}
	
	// Setup runner
	if err := runner.Setup(); err != nil {
		return nil, fmt.Errorf("setup runner: %w", err)
	}
	defer runner.Teardown()
	
	// Run benchmarks for this phase
	if err := im.framework.RunPhase(ctx, phase, runner); err != nil {
		return nil, fmt.Errorf("run phase benchmarks: %w", err)
	}
	
	// Load results
	phaseResults, err := im.framework.storage.LoadResults(phase)
	if err != nil {
		return nil, fmt.Errorf("load phase results: %w", err)
	}
	
	result.Results = phaseResults
	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	
	return result, nil
}

// createRunner creates a benchmark runner for a specific phase
func (im *IntegrationManager) createRunner(phase Phase) (BenchmarkRunner, error) {
	switch phase {
	case PhaseBaseline:
		return &BaselineRunner{}, nil
	case PhaseJSONOptimize:
		return &OptimizedJSONRunner{}, nil
	case PhaseStreaming:
		return &StreamingRunner{}, nil
	case PhaseConcurrency:
		return &ConcurrentRunner{}, nil
	case PhaseMemoryOpt:
		return &MemoryOptimizedRunner{}, nil
	case PhaseCacheOpt:
		return &CacheOptimizedRunner{}, nil
	default:
		return nil, fmt.Errorf("unknown phase: %s", phase)
	}
}

// RunCIBenchmarks runs benchmarks suitable for CI/CD
func (im *IntegrationManager) RunCIBenchmarks(ctx context.Context) (*CIResult, error) {
	// Configure for CI mode
	im.config.CIMode = true
	im.config.FailOnRegression = true
	im.config.Iterations = 3
	im.config.WarmupIterations = 1
	im.config.CPUProfile = false
	im.config.MemProfile = false
	im.config.Visualization = false
	
	result := &CIResult{
		StartTime: time.Now(),
		Success:   true,
	}
	
	// Run baseline and current implementation
	phases := []Phase{PhaseBaseline, PhaseJSONOptimize}
	phaseResults := make(map[Phase][]Result)
	
	for _, phase := range phases {
		runner, err := im.createRunner(phase)
		if err != nil {
			result.Success = false
			result.Error = err.Error()
			return result, err
		}
		
		if err := runner.Setup(); err != nil {
			result.Success = false
			result.Error = err.Error()
			return result, err
		}
		
		if err := im.framework.RunPhase(ctx, phase, runner); err != nil {
			result.Success = false
			result.Error = err.Error()
			runner.Teardown()
			return result, err
		}
		
		runner.Teardown()
		
		// Load results
		results, err := im.framework.storage.LoadResults(phase)
		if err != nil {
			result.Success = false
			result.Error = err.Error()
			return result, err
		}
		
		phaseResults[phase] = results
	}
	
	// Compare phases
	if len(phaseResults) >= 2 {
		comparison := im.framework.comparator.ComparePhases(
			PhaseBaseline, PhaseJSONOptimize,
			phaseResults[PhaseBaseline], phaseResults[PhaseJSONOptimize])
		
		result.Comparison = comparison
		
		// Check for regressions
		if comparison.Summary.Regressions > 0 {
			result.Success = false
			result.Error = fmt.Sprintf("Performance regressions detected: %d", comparison.Summary.Regressions)
		}
	}
	
	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	
	return result, nil
}

// EstablishBaseline establishes performance baselines for future comparisons
func (im *IntegrationManager) EstablishBaseline(ctx context.Context) error {
	// Run baseline phase
	runner := &BaselineRunner{}
	if err := runner.Setup(); err != nil {
		return fmt.Errorf("setup baseline runner: %w", err)
	}
	defer runner.Teardown()
	
	if err := im.framework.RunPhase(ctx, PhaseBaseline, runner); err != nil {
		return fmt.Errorf("run baseline phase: %w", err)
	}
	
	// Load results
	results, err := im.framework.storage.LoadResults(PhaseBaseline)
	if err != nil {
		return fmt.Errorf("load baseline results: %w", err)
	}
	
	// Create baselines from results
	baselines := im.createBaselinesFromResults(results)
	
	// Save baselines
	for _, baseline := range baselines {
		if err := im.framework.storage.SaveBaseline(baseline); err != nil {
			return fmt.Errorf("save baseline: %w", err)
		}
	}
	
	return nil
}

// createBaselinesFromResults creates baselines from benchmark results
func (im *IntegrationManager) createBaselinesFromResults(results []Result) []Baseline {
	// Group results by scenario and operation
	grouped := make(map[string][]Result)
	for _, result := range results {
		key := fmt.Sprintf("%s_%s", result.Scenario, result.Operation)
		grouped[key] = append(grouped[key], result)
	}
	
	baselines := make([]Baseline, 0)
	
	for key, results := range grouped {
		parts := strings.Split(key, "_")
		if len(parts) < 2 {
			continue
		}
		
		scenario := strings.Join(parts[:len(parts)-1], "_")
		operation := parts[len(parts)-1]
		
		// Calculate baseline metrics
		metrics := im.calculateBaselineMetrics(results)
		statistics := im.calculateBaselineStatistics(results)
		
		baseline := Baseline{
			Scenario:   scenario,
			Operation:  operation,
			Metrics:    metrics,
			Statistics: statistics,
			Timestamp:  time.Now(),
			GitCommit:  im.getCurrentGitCommit(),
		}
		
		baselines = append(baselines, baseline)
	}
	
	return baselines
}

// calculateBaselineMetrics calculates baseline metrics from results
func (im *IntegrationManager) calculateBaselineMetrics(results []Result) BaselineMetrics {
	if len(results) == 0 {
		return BaselineMetrics{}
	}
	
	var totalDuration time.Duration
	var totalMemory, totalAllocations int64
	var totalGCPause time.Duration
	var totalBytes int64
	
	for _, result := range results {
		totalDuration += result.Duration
		totalMemory += result.Memory.TotalAlloc
		totalAllocations += result.Memory.Allocations
		totalGCPause += result.Memory.GCPauseTotal
		
		if fileSize, ok := result.Metadata["file_size"].(int64); ok {
			totalBytes += fileSize
		}
	}
	
	avgDuration := totalDuration / time.Duration(len(results))
	avgMemory := totalMemory / int64(len(results))
	avgAllocations := totalAllocations / int64(len(results))
	avgGCPause := totalGCPause / time.Duration(len(results))
	
	// Calculate throughput
	throughput := 0.0
	if avgDuration > 0 && totalBytes > 0 {
		throughput = float64(totalBytes/int64(len(results))) / avgDuration.Seconds()
	}
	
	return BaselineMetrics{
		Duration:    avgDuration,
		Memory:      avgMemory,
		Allocations: avgAllocations,
		GCPause:     avgGCPause,
		Throughput:  throughput,
	}
}

// calculateBaselineStatistics calculates baseline statistics
func (im *IntegrationManager) calculateBaselineStatistics(results []Result) Statistics {
	durations := make([]float64, len(results))
	for i, result := range results {
		durations[i] = float64(result.Duration)
	}
	
	return calculateStatistics(durations)
}

// getCurrentGitCommit gets the current git commit hash
func (im *IntegrationManager) getCurrentGitCommit() string {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = im.workingDir
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

// Additional runner implementations for different phases

// MemoryOptimizedRunner implements memory optimizations
type MemoryOptimizedRunner struct {
	poolEnabled bool
}

func (r *MemoryOptimizedRunner) Name() string {
	return "memory-optimized"
}

func (r *MemoryOptimizedRunner) Setup() error {
	r.poolEnabled = true
	return nil
}

func (r *MemoryOptimizedRunner) Teardown() error {
	return nil
}

func (r *MemoryOptimizedRunner) Run(ctx context.Context, op Operation, data []byte) (interface{}, error) {
	// Implement memory optimizations like object pooling
	switch op.Type {
	case "parse":
		return r.parseWithPooling(data)
	case "extract":
		return r.extractWithPooling(data)
	default:
		return nil, fmt.Errorf("unsupported operation: %s", op.Type)
	}
}

func (r *MemoryOptimizedRunner) parseWithPooling(data []byte) (interface{}, error) {
	// Implement pooled parsing
	return nil, fmt.Errorf("memory optimized parsing not implemented")
}

func (r *MemoryOptimizedRunner) extractWithPooling(data []byte) (interface{}, error) {
	// Implement pooled URL extraction
	return nil, fmt.Errorf("memory optimized extraction not implemented")
}

// CacheOptimizedRunner implements cache optimizations
type CacheOptimizedRunner struct {
	cache map[string]interface{}
}

func (r *CacheOptimizedRunner) Name() string {
	return "cache-optimized"
}

func (r *CacheOptimizedRunner) Setup() error {
	r.cache = make(map[string]interface{})
	return nil
}

func (r *CacheOptimizedRunner) Teardown() error {
	r.cache = nil
	return nil
}

func (r *CacheOptimizedRunner) Run(ctx context.Context, op Operation, data []byte) (interface{}, error) {
	// Implement cached operations
	cacheKey := fmt.Sprintf("%s_%x", op.Name, data[:min(len(data), 32)])
	
	if cached, exists := r.cache[cacheKey]; exists {
		return cached, nil
	}
	
	// Process and cache result
	var result interface{}
	var err error
	
	switch op.Type {
	case "parse":
		result, err = r.parseWithCaching(data)
	case "extract":
		result, err = r.extractWithCaching(data)
	default:
		return nil, fmt.Errorf("unsupported operation: %s", op.Type)
	}
	
	if err == nil {
		r.cache[cacheKey] = result
	}
	
	return result, err
}

func (r *CacheOptimizedRunner) parseWithCaching(data []byte) (interface{}, error) {
	// Implement cached parsing
	return nil, fmt.Errorf("cache optimized parsing not implemented")
}

func (r *CacheOptimizedRunner) extractWithCaching(data []byte) (interface{}, error) {
	// Implement cached URL extraction
	return nil, fmt.Errorf("cache optimized extraction not implemented")
}

// Helper function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Data structures for integration results

// FullSuiteResult contains results from a full benchmark suite
type FullSuiteResult struct {
	StartTime   time.Time                           `json:"start_time"`
	EndTime     time.Time                           `json:"end_time"`
	Duration    time.Duration                       `json:"duration"`
	Phases      map[Phase]*PhaseExecutionResult     `json:"phases"`
	Comparisons []*PhaseComparison                  `json:"comparisons"`
	Validations []*ValidationReport                 `json:"validations"`
}

// PhaseExecutionResult contains results from executing a single phase
type PhaseExecutionResult struct {
	Phase     Phase         `json:"phase"`
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
	Duration  time.Duration `json:"duration"`
	Results   []Result      `json:"results"`
}

// CIResult contains results from CI benchmark run
type CIResult struct {
	StartTime  time.Time        `json:"start_time"`
	EndTime    time.Time        `json:"end_time"`
	Duration   time.Duration    `json:"duration"`
	Success    bool             `json:"success"`
	Error      string           `json:"error,omitempty"`
	Comparison *PhaseComparison `json:"comparison,omitempty"`
}

// Helper methods for IntegrationManager

func (im *IntegrationManager) generateAllComparisons(phases map[Phase]*PhaseExecutionResult) []*PhaseComparison {
	comparisons := make([]*PhaseComparison, 0)
	
	phaseList := []Phase{PhaseBaseline, PhaseJSONOptimize, PhaseStreaming, PhaseConcurrency, PhaseMemoryOpt, PhaseCacheOpt}
	
	// Compare adjacent phases
	for i := 0; i < len(phaseList)-1; i++ {
		phase1 := phaseList[i]
		phase2 := phaseList[i+1]
		
		if result1, ok := phases[phase1]; ok {
			if result2, ok := phases[phase2]; ok {
				comparison := im.framework.comparator.ComparePhases(
					phase1, phase2, result1.Results, result2.Results)
				comparisons = append(comparisons, comparison)
			}
		}
	}
	
	// Compare all phases with baseline
	if baseline, ok := phases[PhaseBaseline]; ok {
		for _, phase := range phaseList[1:] {
			if result, ok := phases[phase]; ok {
				comparison := im.framework.comparator.ComparePhases(
					PhaseBaseline, phase, baseline.Results, result.Results)
				comparisons = append(comparisons, comparison)
			}
		}
	}
	
	return comparisons
}

func (im *IntegrationManager) generateAllValidations(phases map[Phase]*PhaseExecutionResult) []*ValidationReport {
	validations := make([]*ValidationReport, 0)
	
	for phase, result := range phases {
		validation := im.framework.validator.ValidatePhase(phase, result.Results, im.framework.baselines)
		validations = append(validations, validation)
	}
	
	return validations
}

func (im *IntegrationManager) buildReportData(result *FullSuiteResult) *ReportData {
	// Build phase results
	phaseResults := make([]*PhaseResult, 0)
	for _, phase := range []Phase{PhaseBaseline, PhaseJSONOptimize, PhaseStreaming, PhaseConcurrency, PhaseMemoryOpt, PhaseCacheOpt} {
		if phaseResult, ok := result.Phases[phase]; ok {
			// Find corresponding validation
			var validation *ValidationReport
			for _, v := range result.Validations {
				if v.Phase == phase {
					validation = v
					break
				}
			}
			
			phaseReport, _ := im.framework.reporter.GeneratePhaseReport(phase, phaseResult.Results, validation)
			phaseResults = append(phaseResults, phaseReport)
		}
	}
	
	// Generate summary
	summary := im.framework.reporter.GenerateSummaryReport(phaseResults, result.Comparisons, result.Validations)
	
	// Generate recommendations
	data := &ReportData{
		GeneratedAt:  time.Now(),
		Platform:     getCurrentPlatform(),
		GitCommit:    im.getCurrentGitCommit(),
		Config:       im.config,
		Summary:      *summary,
		PhaseResults: phaseResults,
		Comparisons:  result.Comparisons,
		Validations:  result.Validations,
	}
	
	data.Recommendations = im.framework.reporter.GenerateRecommendations(data)
	
	return data
}

// SaveConfiguration saves the benchmark configuration
func (im *IntegrationManager) SaveConfiguration() error {
	configPath := filepath.Join(im.config.OutputDir, "benchmark_config.json")
	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("create config file: %w", err)
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(im.config)
}

// LoadConfiguration loads benchmark configuration
func (im *IntegrationManager) LoadConfiguration(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("open config file: %w", err)
	}
	defer file.Close()
	
	decoder := json.NewDecoder(file)
	return decoder.Decode(&im.config)
}

// GetStatus returns the current status of the benchmark framework
func (im *IntegrationManager) GetStatus() (*Status, error) {
	status := &Status{
		Timestamp: time.Now(),
		Framework: "ready",
		Storage:   "connected",
	}
	
	// Check if storage is accessible
	if _, err := im.framework.storage.LoadBaselines(); err != nil {
		status.Storage = "error: " + err.Error()
	}
	
	// Get recent results count
	phases := []Phase{PhaseBaseline, PhaseJSONOptimize, PhaseStreaming, PhaseConcurrency, PhaseMemoryOpt, PhaseCacheOpt}
	for _, phase := range phases {
		results, err := im.framework.storage.LoadResults(phase)
		if err == nil {
			status.RecentResults += len(results)
		}
	}
	
	return status, nil
}

// Status represents the current framework status
type Status struct {
	Timestamp     time.Time `json:"timestamp"`
	Framework     string    `json:"framework"`
	Storage       string    `json:"storage"`
	RecentResults int       `json:"recent_results"`
}

// Cleanup performs cleanup of resources
func (im *IntegrationManager) Cleanup() error {
	if im.framework != nil && im.framework.storage != nil {
		return im.framework.storage.Close()
	}
	return nil
}