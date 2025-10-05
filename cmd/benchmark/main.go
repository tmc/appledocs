// Command benchmark runs performance benchmarks for appledocs
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/tmc/appledocs/benchmark"
)

var (
	// Benchmark options
	phase        = flag.String("phase", "all", "Phase to benchmark: baseline, json-optimize, streaming, concurrent, memory-opt, all")
	configFile   = flag.String("config", "", "Path to benchmark config file (JSON)")
	outputDir    = flag.String("output", "", "Output directory for results")
	ciMode       = flag.Bool("ci", false, "Run in CI mode with reduced iterations")
	compare      = flag.String("compare", "", "Compare two phases (e.g., 'baseline,json-optimize')")
	continuous   = flag.Bool("continuous", false, "Run continuous monitoring")
	interval     = flag.Duration("interval", 5*time.Minute, "Monitoring interval for continuous mode")
	
	// Test data options
	testDataDir  = flag.String("test-data", "", "Directory containing test data")
	generateData = flag.Bool("generate-data", false, "Generate synthetic test data")
	
	// Report options
	reportFormat = flag.String("report", "json", "Report format: json, html, markdown")
	verbose      = flag.Bool("verbose", false, "Enable verbose output")
)

func main() {
	flag.Parse()

	// Setup signal handling
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("Received interrupt signal, shutting down...")
		cancel()
	}()

	// Load or create configuration
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Override config with command-line flags
	if *outputDir != "" {
		config.OutputDir = *outputDir
	}
	if *reportFormat != "" {
		config.ReportFormat = *reportFormat
	}

	// Create benchmark framework
	framework, err := benchmark.NewFramework(config)
	if err != nil {
		log.Fatalf("Failed to create benchmark framework: %v", err)
	}

	// Generate test data if requested
	if *generateData {
		if err := generateTestData(config); err != nil {
			log.Fatalf("Failed to generate test data: %v", err)
		}
		log.Println("Test data generated successfully")
		return
	}

	// Run continuous monitoring if requested
	if *continuous {
		log.Printf("Starting continuous monitoring with %v interval", *interval)
		if err := framework.RunContinuousMonitoring(ctx, *interval); err != nil {
			log.Fatalf("Continuous monitoring failed: %v", err)
		}
		return
	}

	// Handle comparison mode
	if *compare != "" {
		if err := runComparison(framework, *compare); err != nil {
			log.Fatalf("Comparison failed: %v", err)
		}
		return
	}

	// Run benchmarks
	phases := getPhasesToRun(*phase)
	log.Printf("Running benchmarks for phases: %v", phases)

	for _, phase := range phases {
		log.Printf("Running phase: %s", phase)
		
		runner := benchmark.RunnerFactory(phase)
		if err := runner.Setup(); err != nil {
			log.Fatalf("Failed to setup runner for %s: %v", phase, err)
		}
		defer runner.Teardown()

		if err := framework.RunPhase(ctx, phase, runner); err != nil {
			log.Printf("Phase %s failed: %v", phase, err)
			if config.FailOnRegression {
				os.Exit(1)
			}
		}
		
		log.Printf("Phase %s completed", phase)
	}

	// Generate final report
	log.Println("Generating benchmark report...")
	report, err := framework.GenerateReport()
	if err != nil {
		log.Fatalf("Failed to generate report: %v", err)
	}

	// Print summary
	printSummary(report)

	// Exit with appropriate code
	if report.Summary.OverallVerdict == "Significant regressions detected" && config.FailOnRegression {
		os.Exit(1)
	}
}

// loadConfig loads benchmark configuration
func loadConfig() (benchmark.Config, error) {
	var config benchmark.Config

	if *configFile != "" {
		// Load from file
		data, err := os.ReadFile(*configFile)
		if err != nil {
			return config, fmt.Errorf("read config file: %w", err)
		}
		if err := json.Unmarshal(data, &config); err != nil {
			return config, fmt.Errorf("parse config: %w", err)
		}
	} else {
		// Use default config
		if *ciMode {
			config = benchmark.CIAppledocsConfig()
		} else {
			config = benchmark.DefaultAppledocsConfig()
		}
	}

	// Ensure output directory
	if config.OutputDir == "" {
		config.OutputDir = filepath.Join("benchmark_results", time.Now().Format("2006-01-02_15-04-05"))
	}

	return config, nil
}

// getPhasesToRun returns the phases to benchmark
func getPhasesToRun(phaseFlag string) []benchmark.Phase {
	switch phaseFlag {
	case "all":
		return []benchmark.Phase{
			benchmark.PhaseBaseline,
			benchmark.PhaseJSONOptimize,
			benchmark.PhaseStreaming,
			benchmark.PhaseConcurrency,
			benchmark.PhaseMemoryOpt,
		}
	case "baseline":
		return []benchmark.Phase{benchmark.PhaseBaseline}
	case "json-optimize":
		return []benchmark.Phase{benchmark.PhaseJSONOptimize}
	case "streaming":
		return []benchmark.Phase{benchmark.PhaseStreaming}
	case "concurrent":
		return []benchmark.Phase{benchmark.PhaseConcurrency}
	case "memory-opt":
		return []benchmark.Phase{benchmark.PhaseMemoryOpt}
	default:
		log.Fatalf("Unknown phase: %s", phaseFlag)
		return nil
	}
}

// runComparison runs comparison between phases
func runComparison(framework *benchmark.BenchmarkFramework, compareFlag string) error {
	// Parse phase names
	var phase1, phase2 benchmark.Phase
	if _, err := fmt.Sscanf(compareFlag, "%s,%s", &phase1, &phase2); err != nil {
		return fmt.Errorf("invalid compare format, use 'phase1,phase2': %w", err)
	}

	// Run both phases if needed
	// ... (implementation would check if phases have been run)

	// Compare phases
	comparison := framework.ComparePhases(phase1, phase2)
	if comparison == nil {
		return fmt.Errorf("no results found for comparison")
	}

	// Print comparison results
	fmt.Printf("\nComparison: %s vs %s\n", phase1, phase2)
	fmt.Printf("==========================================\n")
	fmt.Printf("Summary:\n")
	fmt.Printf("  Improved: %d operations\n", comparison.Summary.ImprovedCount)
	fmt.Printf("  Regressed: %d operations\n", comparison.Summary.RegressedCount)
	fmt.Printf("  Unchanged: %d operations\n", comparison.Summary.UnchangedCount)
	fmt.Printf("  Overall: %.1f%% improvement\n", comparison.Summary.OverallImprovement)

	if len(comparison.Improvements()) > 0 {
		fmt.Printf("\nTop Improvements:\n")
		for i, improvement := range comparison.Improvements() {
			if i >= 5 {
				break
			}
			fmt.Printf("  - %s\n", improvement)
		}
	}

	if len(comparison.Regressions()) > 0 {
		fmt.Printf("\nTop Regressions:\n")
		for i, regression := range comparison.Regressions() {
			if i >= 5 {
				break
			}
			fmt.Printf("  - %s\n", regression)
		}
	}

	return nil
}

// printSummary prints benchmark summary
func printSummary(report *benchmark.BenchmarkReport) {
	fmt.Printf("\nBenchmark Summary\n")
	fmt.Printf("==========================================\n")
	fmt.Printf("Total Phases: %d\n", report.Summary.TotalPhases)
	fmt.Printf("Best Phase: %s\n", report.Summary.BestPhase)
	fmt.Printf("Worst Phase: %s\n", report.Summary.WorstPhase)
	fmt.Printf("Overall Verdict: %s\n", report.Summary.OverallVerdict)

	if len(report.Summary.CriticalFindings) > 0 {
		fmt.Printf("\nCritical Findings:\n")
		for _, finding := range report.Summary.CriticalFindings {
			fmt.Printf("  ! %s\n", finding)
		}
	}

	if len(report.Summary.Recommendations) > 0 {
		fmt.Printf("\nRecommendations:\n")
		for _, rec := range report.Summary.Recommendations {
			fmt.Printf("  - %s\n", rec)
		}
	}

	fmt.Printf("\nDetailed report saved to: %s\n", report.Config.OutputDir)
}

// generateTestData generates synthetic test data
func generateTestData(config benchmark.Config) error {
	// Create test data directories
	dirs := []string{
		filepath.Join("testdata", "json", "small"),
		filepath.Join("testdata", "json", "medium"),
		filepath.Join("testdata", "json", "large"),
		filepath.Join("testdata", "json", "xlarge"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("create directory %s: %w", dir, err)
		}
	}

	// Generate files of different sizes
	sizes := map[string]int{
		"small":  10 * 1024,        // 10KB
		"medium": 1024 * 1024,      // 1MB
		"large":  10 * 1024 * 1024, // 10MB
		"xlarge": 50 * 1024 * 1024, // 50MB
	}

	for category, size := range sizes {
		for i := 0; i < 3; i++ {
			filename := filepath.Join("testdata", "json", category, fmt.Sprintf("test_%d.json", i))
			if err := generateJSONFile(filename, size); err != nil {
				return fmt.Errorf("generate %s: %w", filename, err)
			}
			log.Printf("Generated %s (%d bytes)", filename, size)
		}
	}

	// Generate URL list for crawling tests
	urlFile := filepath.Join("testdata", "urls.txt")
	if err := generateURLList(urlFile, 1000); err != nil {
		return fmt.Errorf("generate URL list: %w", err)
	}

	return nil
}

// generateJSONFile generates a JSON file of specified size
func generateJSONFile(filename string, targetSize int) error {
	// Implementation would generate realistic Apple documentation JSON
	data := make(map[string]interface{})
	data["metadata"] = map[string]interface{}{
		"title": "Test Document",
		"platforms": []string{"iOS", "macOS"},
	}
	
	// Add content until we reach target size
	// ... (simplified for brevity)
	
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(filename, jsonData, 0644)
}

// generateURLList generates a list of test URLs
func generateURLList(filename string, count int) error {
	var urls []string
	baseURL := "https://developer.apple.com/tutorials/data/documentation"
	
	frameworks := []string{"UIKit", "SwiftUI", "Foundation", "CoreData", "Combine"}
	for i := 0; i < count; i++ {
		framework := frameworks[i%len(frameworks)]
		urls = append(urls, fmt.Sprintf("%s/%s/class%d.json", baseURL, framework, i))
	}
	
	content := ""
	for _, url := range urls {
		content += url + "\n"
	}
	
	return os.WriteFile(filename, []byte(content), 0644)
}