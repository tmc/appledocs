// Package benchmark provides a comprehensive benchmarking framework for appledocs refactoring
package benchmark

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/montanaflynn/stats"
)

// Phase represents a refactoring phase
type Phase string

const (
	PhaseBaseline     Phase = "baseline"
	PhaseJSONOptimize Phase = "json-optimize"
	PhaseStreaming    Phase = "streaming"
	PhaseConcurrency  Phase = "concurrency"
	PhaseMemoryOpt    Phase = "memory-opt"
	PhaseCacheOpt     Phase = "cache-opt"
)

// BenchmarkFramework manages performance benchmarking across refactoring phases
type BenchmarkFramework struct {
	config      Config
	results     map[Phase][]Result
	baselines   map[string]Baseline
	mu          sync.RWMutex
	profiler    *Profiler
	reporter    *Reporter
	validator   *Validator
	comparator  *Comparator
	storage     Storage
}

// Config holds benchmark configuration
type Config struct {
	// Benchmark settings
	Iterations       int                    `json:"iterations"`
	WarmupIterations int                    `json:"warmup_iterations"`
	Timeout          time.Duration          `json:"timeout"`
	CPUProfile       bool                   `json:"cpu_profile"`
	MemProfile       bool                   `json:"mem_profile"`
	BlockProfile     bool                   `json:"block_profile"`
	TraceEnabled     bool                   `json:"trace_enabled"`
	
	// Test scenarios
	Scenarios        []Scenario             `json:"scenarios"`
	
	// Performance thresholds
	Thresholds       map[string]Threshold   `json:"thresholds"`
	
	// Output settings
	OutputDir        string                 `json:"output_dir"`
	ReportFormat     string                 `json:"report_format"` // json, html, markdown
	Visualization    bool                   `json:"visualization"`
	
	// Integration settings
	CIMode           bool                   `json:"ci_mode"`
	FailOnRegression bool                   `json:"fail_on_regression"`
	
	// Platform-specific settings
	Platforms        []Platform             `json:"platforms"`
}

// Scenario represents a benchmark scenario
type Scenario struct {
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	FilePatterns    []string               `json:"file_patterns"`
	FileSizes       []string               `json:"file_sizes"` // small, medium, large, xlarge
	Concurrency     []int                  `json:"concurrency_levels"`
	Operations      []Operation            `json:"operations"`
	Tags            []string               `json:"tags"`
	Weight          float64                `json:"weight"` // Importance weight for scoring
}

// Operation represents a specific operation to benchmark
type Operation struct {
	Name        string                 `json:"name"`
	Type        string                 `json:"type"` // parse, extract, transform, write
	Function    string                 `json:"function"`
	Parameters  map[string]interface{} `json:"parameters"`
}

// Threshold defines performance thresholds
type Threshold struct {
	MaxDuration      time.Duration `json:"max_duration"`
	MaxMemory        int64         `json:"max_memory_bytes"`
	MaxAllocations   int64         `json:"max_allocations"`
	MaxGCPause       time.Duration `json:"max_gc_pause"`
	RegressionMargin float64       `json:"regression_margin"` // Allowed regression percentage
}

// Platform represents a test platform
type Platform struct {
	OS           string `json:"os"`
	Arch         string `json:"arch"`
	CPUCores     int    `json:"cpu_cores"`
	Memory       int64  `json:"memory_gb"`
	GoVersion    string `json:"go_version"`
}

// Result holds benchmark results
type Result struct {
	Phase           Phase                  `json:"phase"`
	Scenario        string                 `json:"scenario"`
	Operation       string                 `json:"operation"`
	Timestamp       time.Time              `json:"timestamp"`
	Duration        time.Duration          `json:"duration"`
	Memory          MemoryStats            `json:"memory"`
	CPU             CPUStats               `json:"cpu"`
	Network         NetworkStats           `json:"network"`
	Cache           CacheStats             `json:"cache"`
	Errors          []Error                `json:"errors,omitempty"`
	Metadata        map[string]interface{} `json:"metadata"`
	Platform        Platform               `json:"platform"`
}

// MemoryStats holds memory-related statistics
type MemoryStats struct {
	HeapAlloc      int64         `json:"heap_alloc"`
	HeapInuse      int64         `json:"heap_inuse"`
	HeapReleased   int64         `json:"heap_released"`
	StackInuse     int64         `json:"stack_inuse"`
	Allocations    int64         `json:"allocations"`
	TotalAlloc     int64         `json:"total_alloc"`
	Sys            int64         `json:"sys"`
	NumGC          uint32        `json:"num_gc"`
	GCPauseTotal   time.Duration `json:"gc_pause_total"`
	GCPauseAvg     time.Duration `json:"gc_pause_avg"`
	GCPauseMax     time.Duration `json:"gc_pause_max"`
	GCPauseP99     time.Duration `json:"gc_pause_p99"`
}

// CPUStats holds CPU-related statistics
type CPUStats struct {
	UserTime       time.Duration `json:"user_time"`
	SystemTime     time.Duration `json:"system_time"`
	CPUPercent     float64       `json:"cpu_percent"`
	Goroutines     int           `json:"goroutines"`
	CGoCalls       int64         `json:"cgo_calls"`
}

// NetworkStats holds network-related statistics
type NetworkStats struct {
	RequestsTotal  int64         `json:"requests_total"`
	BytesSent      int64         `json:"bytes_sent"`
	BytesReceived  int64         `json:"bytes_received"`
	AvgLatency     time.Duration `json:"avg_latency"`
	P99Latency     time.Duration `json:"p99_latency"`
	ErrorRate      float64       `json:"error_rate"`
}

// CacheStats holds cache-related statistics
type CacheStats struct {
	Hits           int64   `json:"hits"`
	Misses         int64   `json:"misses"`
	HitRate        float64 `json:"hit_rate"`
	Size           int64   `json:"size_bytes"`
	Evictions      int64   `json:"evictions"`
}

// Error represents a benchmark error
type Error struct {
	Type        string    `json:"type"`
	Message     string    `json:"message"`
	Timestamp   time.Time `json:"timestamp"`
	StackTrace  string    `json:"stack_trace,omitempty"`
}

// Baseline represents performance baseline for comparison
type Baseline struct {
	Scenario       string                 `json:"scenario"`
	Operation      string                 `json:"operation"`
	Metrics        BaselineMetrics        `json:"metrics"`
	Statistics     Statistics             `json:"statistics"`
	Timestamp      time.Time              `json:"timestamp"`
	GitCommit      string                 `json:"git_commit"`
}

// BaselineMetrics holds baseline performance metrics
type BaselineMetrics struct {
	Duration       time.Duration `json:"duration"`
	Memory         int64         `json:"memory_bytes"`
	Allocations    int64         `json:"allocations"`
	GCPause        time.Duration `json:"gc_pause"`
	Throughput     float64       `json:"throughput"`
}

// Statistics holds statistical analysis of results
type Statistics struct {
	Mean           float64 `json:"mean"`
	Median         float64 `json:"median"`
	StdDev         float64 `json:"std_dev"`
	Min            float64 `json:"min"`
	Max            float64 `json:"max"`
	P50            float64 `json:"p50"`
	P90            float64 `json:"p90"`
	P95            float64 `json:"p95"`
	P99            float64 `json:"p99"`
	CV             float64 `json:"coefficient_variation"` // StdDev/Mean
	CI95Lower      float64 `json:"ci95_lower"`
	CI95Upper      float64 `json:"ci95_upper"`
}

// NewFramework creates a new benchmark framework
func NewFramework(config Config) (*BenchmarkFramework, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	// Create output directory
	if err := os.MkdirAll(config.OutputDir, 0755); err != nil {
		return nil, fmt.Errorf("create output dir: %w", err)
	}

	bf := &BenchmarkFramework{
		config:     config,
		results:    make(map[Phase][]Result),
		baselines:  make(map[string]Baseline),
		profiler:   NewProfiler(config),
		reporter:   NewReporter(config),
		validator:  NewValidator(config),
		comparator: NewComparator(config),
	}

	// Initialize storage
	storage, err := NewStorage(filepath.Join(config.OutputDir, "benchmark.db"))
	if err != nil {
		return nil, fmt.Errorf("initialize storage: %w", err)
	}
	bf.storage = storage

	// Load existing baselines
	if err := bf.loadBaselines(); err != nil {
		return nil, fmt.Errorf("load baselines: %w", err)
	}

	return bf, nil
}

// RunPhase executes benchmarks for a specific refactoring phase
func (bf *BenchmarkFramework) RunPhase(ctx context.Context, phase Phase, runner BenchmarkRunner) error {
	bf.mu.Lock()
	defer bf.mu.Unlock()

	log := bf.createPhaseLogger(phase)
	log.Info("Starting benchmark phase", "scenarios", len(bf.config.Scenarios))

	phaseResults := make([]Result, 0)
	
	for _, scenario := range bf.config.Scenarios {
		if err := bf.runScenario(ctx, phase, scenario, runner, &phaseResults); err != nil {
			log.Error("Scenario failed", "scenario", scenario.Name, "error", err)
			if bf.config.CIMode && bf.config.FailOnRegression {
				return err
			}
		}
	}

	bf.results[phase] = phaseResults
	
	// Save results to storage
	if err := bf.storage.SaveResults(phase, phaseResults); err != nil {
		log.Error("Failed to save results", "error", err)
	}

	// Validate results
	validationReport := bf.validator.ValidatePhase(phase, phaseResults, bf.baselines)
	if validationReport.HasRegressions() && bf.config.FailOnRegression {
		return fmt.Errorf("performance regressions detected: %v", validationReport.Summary())
	}

	return nil
}

// runScenario executes a single benchmark scenario
func (bf *BenchmarkFramework) runScenario(ctx context.Context, phase Phase, scenario Scenario, 
	runner BenchmarkRunner, results *[]Result) error {
	
	log := bf.createScenarioLogger(scenario.Name)
	log.Info("Running scenario", "operations", len(scenario.Operations))

	// Prepare test data
	testData, err := bf.prepareTestData(scenario)
	if err != nil {
		return fmt.Errorf("prepare test data: %w", err)
	}

	// Run warmup iterations
	if bf.config.WarmupIterations > 0 {
		log.Debug("Running warmup iterations", "count", bf.config.WarmupIterations)
		for i := 0; i < bf.config.WarmupIterations; i++ {
			_ = bf.runOperations(ctx, phase, scenario, runner, testData, true)
		}
	}

	// Collect results for each iteration
	iterationResults := make([][]Result, bf.config.Iterations)
	
	for i := 0; i < bf.config.Iterations; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		
		log.Debug("Running iteration", "number", i+1)
		
		// Start profiling if enabled
		if bf.config.CPUProfile && i == 0 {
			if err := bf.profiler.StartCPUProfile(scenario.Name); err != nil {
				log.Warn("Failed to start CPU profile", "error", err)
			}
		}
		
		// Run operations
		opResults, err := bf.runOperations(ctx, phase, scenario, runner, testData, false)
		if err != nil {
			log.Error("Operation failed", "iteration", i, "error", err)
			continue
		}
		
		iterationResults[i] = opResults
		
		// Stop profiling after first iteration
		if bf.config.CPUProfile && i == 0 {
			bf.profiler.StopCPUProfile()
		}
		
		// Memory profile after each iteration
		if bf.config.MemProfile {
			if err := bf.profiler.CaptureMemProfile(scenario.Name, i); err != nil {
				log.Warn("Failed to capture memory profile", "error", err)
			}
		}
	}

	// Aggregate results
	aggregated := bf.aggregateResults(scenario, iterationResults)
	*results = append(*results, aggregated...)

	return nil
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.Iterations < 1 {
		return fmt.Errorf("iterations must be at least 1")
	}
	if c.OutputDir == "" {
		return fmt.Errorf("output directory required")
	}
	if len(c.Scenarios) == 0 {
		return fmt.Errorf("at least one scenario required")
	}
	return nil
}

// prepareTestData prepares test data for a scenario
func (bf *BenchmarkFramework) prepareTestData(scenario Scenario) ([]TestData, error) {
	var testData []TestData
	
	// Load files matching patterns
	for _, pattern := range scenario.FilePatterns {
		files, err := filepath.Glob(pattern)
		if err != nil {
			return nil, fmt.Errorf("glob pattern %s: %w", pattern, err)
		}
		
		for _, file := range files {
			data, err := os.ReadFile(file)
			if err != nil {
				return nil, fmt.Errorf("read file %s: %w", file, err)
			}
			
			info, _ := os.Stat(file)
			size := categorizeFileSize(info.Size())
			
			// Check if size matches scenario requirements
			if len(scenario.FileSizes) > 0 && !contains(scenario.FileSizes, size) {
				continue
			}
			
			testData = append(testData, TestData{
				Name: filepath.Base(file),
				Path: file,
				Data: data,
				Size: info.Size(),
				Category: size,
			})
		}
	}
	
	// Generate synthetic data if needed
	if len(testData) == 0 {
		testData = bf.generateSyntheticData(scenario)
	}
	
	return testData, nil
}

// TestData represents test data for benchmarking
type TestData struct {
	Name     string
	Path     string
	Data     []byte
	Size     int64
	Category string // small, medium, large, xlarge
}

// categorizeFileSize categorizes file size
func categorizeFileSize(size int64) string {
	switch {
	case size < 10*1024: // < 10KB
		return "small"
	case size < 1024*1024: // < 1MB
		return "medium"
	case size < 10*1024*1024: // < 10MB
		return "large"
	default:
		return "xlarge"
	}
}

// contains checks if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// createPhaseLogger creates a logger for a phase
func (bf *BenchmarkFramework) createPhaseLogger(phase Phase) Logger {
	return Logger{phase: string(phase)}
}

// createScenarioLogger creates a logger for a scenario
func (bf *BenchmarkFramework) createScenarioLogger(scenario string) Logger {
	return Logger{scenario: scenario}
}

// Logger is a simple logger for benchmarking
type Logger struct {
	phase    string
	scenario string
}

func (l Logger) Info(msg string, args ...interface{}) {
	fmt.Printf("[INFO] %s: %s %v\n", l.getPrefix(), msg, args)
}

func (l Logger) Debug(msg string, args ...interface{}) {
	fmt.Printf("[DEBUG] %s: %s %v\n", l.getPrefix(), msg, args)
}

func (l Logger) Warn(msg string, args ...interface{}) {
	fmt.Printf("[WARN] %s: %s %v\n", l.getPrefix(), msg, args)
}

func (l Logger) Error(msg string, args ...interface{}) {
	fmt.Printf("[ERROR] %s: %s %v\n", l.getPrefix(), msg, args)
}

func (l Logger) getPrefix() string {
	if l.phase != "" {
		return l.phase
	}
	if l.scenario != "" {
		return l.scenario
	}
	return "benchmark"
}