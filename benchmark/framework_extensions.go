// Package benchmark provides extensions to the benchmark framework
package benchmark

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// ComparePhases compares results between two phases
func (bf *BenchmarkFramework) ComparePhases(baseline, current Phase) *PhaseComparison {
	bf.mu.RLock()
	baselineResults := bf.results[baseline]
	currentResults := bf.results[current]
	bf.mu.RUnlock()

	if len(baselineResults) == 0 || len(currentResults) == 0 {
		return nil
	}

	return bf.comparator.ComparePhases(baselineResults, currentResults)
}

// GenerateReport generates a comprehensive benchmark report
func (bf *BenchmarkFramework) GenerateReport() (*BenchmarkReport, error) {
	bf.mu.RLock()
	defer bf.mu.RUnlock()

	report := &BenchmarkReport{
		Timestamp:    time.Now(),
		Config:       bf.config,
		PhaseResults: make(map[Phase]PhaseResult),
		Comparisons:  make(map[string]*PhaseComparison),
	}

	// Collect results for each phase
	for phase, results := range bf.results {
		phaseResult := bf.analyzePhaseResults(phase, results)
		report.PhaseResults[phase] = phaseResult
	}

	// Generate comparisons if we have baseline
	if baselineResults, hasBaseline := bf.results[PhaseBaseline]; hasBaseline {
		for phase, results := range bf.results {
			if phase != PhaseBaseline && len(results) > 0 {
				comparison := bf.comparator.ComparePhases(baselineResults, results)
				report.Comparisons[string(phase)] = comparison
			}
		}
	}

	// Generate summary
	report.Summary = bf.generateSummary(report)

	// Save report
	if err := bf.saveReport(report); err != nil {
		return report, fmt.Errorf("save report: %w", err)
	}

	return report, nil
}

// BenchmarkReport holds the complete benchmark report
type BenchmarkReport struct {
	Timestamp    time.Time                    `json:"timestamp"`
	Config       Config                       `json:"config"`
	PhaseResults map[Phase]PhaseResult        `json:"phase_results"`
	Comparisons  map[string]*PhaseComparison  `json:"comparisons"`
	Summary      *ReportSummary               `json:"summary"`
}

// PhaseResult holds results for a single phase
type PhaseResult struct {
	Phase          Phase             `json:"phase"`
	ScenarioStats  map[string]ScenarioStats `json:"scenario_stats"`
	TotalDuration  time.Duration     `json:"total_duration"`
	SuccessRate    float64           `json:"success_rate"`
	AverageMemory  int64             `json:"average_memory_bytes"`
}

// ScenarioStats holds statistics for a scenario
type ScenarioStats struct {
	OperationCount int               `json:"operation_count"`
	SuccessCount   int               `json:"success_count"`
	ErrorCount     int               `json:"error_count"`
	DurationStats  Statistics        `json:"duration_stats"`
	MemoryStats    Statistics        `json:"memory_stats"`
}

// ReportSummary holds the report summary
type ReportSummary struct {
	TotalPhases      int                      `json:"total_phases"`
	BestPhase        Phase                    `json:"best_phase"`
	WorstPhase       Phase                    `json:"worst_phase"`
	OverallVerdict   string                   `json:"overall_verdict"`
	Recommendations  []string                 `json:"recommendations"`
	CriticalFindings []string                 `json:"critical_findings"`
}

// analyzePhaseResults analyzes results for a phase
func (bf *BenchmarkFramework) analyzePhaseResults(phase Phase, results []Result) PhaseResult {
	phaseResult := PhaseResult{
		Phase:         phase,
		ScenarioStats: make(map[string]ScenarioStats),
	}

	// Group by scenario
	scenarioResults := make(map[string][]Result)
	for _, result := range results {
		scenarioResults[result.Scenario] = append(scenarioResults[result.Scenario], result)
	}

	// Calculate stats for each scenario
	totalDuration := time.Duration(0)
	totalMemory := int64(0)
	successCount := 0
	totalCount := 0

	for scenario, results := range scenarioResults {
		stats := ScenarioStats{
			OperationCount: len(results),
		}

		durations := make([]float64, 0)
		memories := make([]float64, 0)

		for _, result := range results {
			totalCount++
			if len(result.Errors) == 0 {
				stats.SuccessCount++
				successCount++
			} else {
				stats.ErrorCount++
			}

			durations = append(durations, float64(result.Duration))
			memories = append(memories, float64(result.Memory.TotalAlloc))
			
			totalDuration += result.Duration
			totalMemory += result.Memory.TotalAlloc
		}

		stats.DurationStats = calculateStatistics(durations)
		stats.MemoryStats = calculateStatistics(memories)

		phaseResult.ScenarioStats[scenario] = stats
	}

	phaseResult.TotalDuration = totalDuration
	if totalCount > 0 {
		phaseResult.SuccessRate = float64(successCount) / float64(totalCount)
		phaseResult.AverageMemory = totalMemory / int64(totalCount)
	}

	return phaseResult
}

// generateSummary generates report summary
func (bf *BenchmarkFramework) generateSummary(report *BenchmarkReport) *ReportSummary {
	summary := &ReportSummary{
		TotalPhases:     len(report.PhaseResults),
		Recommendations: make([]string, 0),
		CriticalFindings: make([]string, 0),
	}

	// Find best and worst phases
	var bestPhase Phase
	var worstPhase Phase
	bestDuration := time.Duration(0)
	worstDuration := time.Duration(0)

	for phase, result := range report.PhaseResults {
		avgDuration := result.TotalDuration / time.Duration(len(bf.config.Scenarios))
		
		if bestDuration == 0 || avgDuration < bestDuration {
			bestDuration = avgDuration
			bestPhase = phase
		}
		if avgDuration > worstDuration {
			worstDuration = avgDuration
			worstPhase = phase
		}
	}

	summary.BestPhase = bestPhase
	summary.WorstPhase = worstPhase

	// Analyze comparisons
	improvementCount := 0
	regressionCount := 0
	
	for _, comparison := range report.Comparisons {
		improvementCount += comparison.Summary.ImprovedCount
		regressionCount += comparison.Summary.RegressedCount
		
		// Add critical findings
		if comparison.Summary.RegressedCount > comparison.Summary.ImprovedCount {
			summary.CriticalFindings = append(summary.CriticalFindings, 
				fmt.Sprintf("Phase %s shows more regressions than improvements", comparison.CurrentPhase))
		}
	}

	// Determine overall verdict
	if improvementCount > regressionCount*2 {
		summary.OverallVerdict = "Significant improvements detected"
	} else if regressionCount > improvementCount*2 {
		summary.OverallVerdict = "Significant regressions detected"
	} else {
		summary.OverallVerdict = "Mixed results - further analysis needed"
	}

	// Generate recommendations
	summary.Recommendations = bf.generateRecommendations(report)

	return summary
}

// generateRecommendations generates actionable recommendations
func (bf *BenchmarkFramework) generateRecommendations(report *BenchmarkReport) []string {
	recommendations := make([]string, 0)

	// Check for memory issues
	for phase, result := range report.PhaseResults {
		if result.AverageMemory > 100*1024*1024 { // 100MB average
			recommendations = append(recommendations, 
				fmt.Sprintf("High memory usage in %s phase - consider memory optimization", phase))
		}
	}

	// Check for performance regressions
	for phaseName, comparison := range report.Comparisons {
		if comparison.Summary.RegressedCount > 0 {
			for _, comp := range comparison.Comparisons {
				if comp.Verdict == "regressed" && comp.Changes.DurationChange > 50 {
					recommendations = append(recommendations,
						fmt.Sprintf("Investigate %s/%s in %s - 50%% performance regression", 
							comp.Scenario, comp.Operation, phaseName))
				}
			}
		}
	}

	// Check success rates
	for phase, result := range report.PhaseResults {
		if result.SuccessRate < 0.95 {
			recommendations = append(recommendations,
				fmt.Sprintf("Low success rate (%.1f%%) in %s phase - investigate failures", 
					result.SuccessRate*100, phase))
		}
	}

	return recommendations
}

// saveReport saves the report to disk
func (bf *BenchmarkFramework) saveReport(report *BenchmarkReport) error {
	// Save JSON report
	jsonPath := filepath.Join(bf.config.OutputDir, "benchmark_report.json")
	jsonData, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal report: %w", err)
	}

	if err := os.WriteFile(jsonPath, jsonData, 0644); err != nil {
		return fmt.Errorf("write JSON report: %w", err)
	}

	// Generate HTML report if requested
	if bf.config.ReportFormat == "html" || bf.config.Visualization {
		htmlPath := filepath.Join(bf.config.OutputDir, "benchmark_report.html")
		if err := bf.reporter.GenerateHTMLReport(report, htmlPath); err != nil {
			return fmt.Errorf("generate HTML report: %w", err)
		}
	}

	// Generate Markdown report if requested
	if bf.config.ReportFormat == "markdown" {
		mdPath := filepath.Join(bf.config.OutputDir, "benchmark_report.md")
		if err := bf.reporter.GenerateMarkdownReport(report, mdPath); err != nil {
			return fmt.Errorf("generate Markdown report: %w", err)
		}
	}

	return nil
}

// loadBaselines loads existing baseline results
func (bf *BenchmarkFramework) loadBaselines() error {
	baselines, err := bf.storage.LoadBaselines()
	if err != nil {
		return fmt.Errorf("load baselines: %w", err)
	}

	bf.baselines = baselines
	return nil
}

// AnalyzeMemoryGrowth analyzes memory growth patterns
func (bf *BenchmarkFramework) AnalyzeMemoryGrowth(phase Phase) *MemoryAnalysis {
	bf.mu.RLock()
	results := bf.results[phase]
	bf.mu.RUnlock()

	if len(results) == 0 {
		return nil
	}

	analysis := &MemoryAnalysis{
		Phase:     phase,
		Timestamp: time.Now(),
	}

	// Group by operation to track memory growth
	operationMemory := make(map[string][]int64)
	for _, result := range results {
		key := fmt.Sprintf("%s/%s", result.Scenario, result.Operation)
		operationMemory[key] = append(operationMemory[key], result.Memory.HeapAlloc)
	}

	// Analyze growth patterns
	for operation, memories := range operationMemory {
		if len(memories) < 3 {
			continue
		}

		// Calculate growth rate
		growthRate := float64(memories[len(memories)-1]-memories[0]) / float64(memories[0])
		
		// Check for potential leak (consistent growth)
		isIncreasing := true
		for i := 1; i < len(memories); i++ {
			if memories[i] < memories[i-1] {
				isIncreasing = false
				break
			}
		}

		if isIncreasing && growthRate > 0.5 { // 50% growth
			analysis.HasLeak = true
			analysis.LeakOperations = append(analysis.LeakOperations, LeakInfo{
				Operation:  operation,
				GrowthRate: growthRate,
				Samples:    len(memories),
			})
		}
	}

	return analysis
}

// MemoryAnalysis holds memory growth analysis
type MemoryAnalysis struct {
	Phase          Phase      `json:"phase"`
	Timestamp      time.Time  `json:"timestamp"`
	HasLeak        bool       `json:"has_leak"`
	LeakOperations []LeakInfo `json:"leak_operations,omitempty"`
}

// LeakInfo holds information about a potential memory leak
type LeakInfo struct {
	Operation  string  `json:"operation"`
	GrowthRate float64 `json:"growth_rate"`
	Samples    int     `json:"samples"`
}

// RunContinuousMonitoring runs continuous performance monitoring
func (bf *BenchmarkFramework) RunContinuousMonitoring(ctx context.Context, interval time.Duration) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			// Run a lightweight benchmark
			if err := bf.runMonitoringBenchmark(ctx); err != nil {
				logger := bf.createPhaseLogger("monitoring")
				logger.Error("Monitoring benchmark failed", "error", err)
			}
		}
	}
}

// runMonitoringBenchmark runs a lightweight monitoring benchmark
func (bf *BenchmarkFramework) runMonitoringBenchmark(ctx context.Context) error {
	// Use a subset of scenarios for monitoring
	monitoringScenarios := make([]Scenario, 0)
	for _, scenario := range bf.config.Scenarios {
		if scenario.Weight >= 1.0 { // Only high-priority scenarios
			monitoringScenarios = append(monitoringScenarios, scenario)
		}
	}

	// Create a monitoring config
	monitoringConfig := bf.config
	monitoringConfig.Scenarios = monitoringScenarios
	monitoringConfig.Iterations = 1
	monitoringConfig.WarmupIterations = 0

	// Run benchmark
	runner := RunnerFactory(PhaseBaseline)
	if err := bf.RunPhase(ctx, PhaseBaseline, runner); err != nil {
		return err
	}

	// Check for alerts
	bf.checkPerformanceAlerts()

	return nil
}

// checkPerformanceAlerts checks for performance alerts
func (bf *BenchmarkFramework) checkPerformanceAlerts() {
	bf.mu.RLock()
	defer bf.mu.RUnlock()

	// Check latest results against thresholds
	for phase, results := range bf.results {
		for _, result := range results {
			threshold, exists := bf.config.Thresholds[result.Scenario]
			if !exists {
				continue
			}

			// Check duration threshold
			if result.Duration > threshold.MaxDuration {
				bf.createAlert(AlertInfo{
					Type:      "duration",
					Phase:     phase,
					Scenario:  result.Scenario,
					Operation: result.Operation,
					Value:     float64(result.Duration),
					Threshold: float64(threshold.MaxDuration),
					Message:   fmt.Sprintf("Duration exceeded threshold: %v > %v", result.Duration, threshold.MaxDuration),
				})
			}

			// Check memory threshold
			if result.Memory.HeapAlloc > threshold.MaxMemory {
				bf.createAlert(AlertInfo{
					Type:      "memory",
					Phase:     phase,
					Scenario:  result.Scenario,
					Operation: result.Operation,
					Value:     float64(result.Memory.HeapAlloc),
					Threshold: float64(threshold.MaxMemory),
					Message:   fmt.Sprintf("Memory exceeded threshold: %d > %d", result.Memory.HeapAlloc, threshold.MaxMemory),
				})
			}
		}
	}
}

// AlertInfo holds alert information
type AlertInfo struct {
	Type      string    `json:"type"`
	Phase     Phase     `json:"phase"`
	Scenario  string    `json:"scenario"`
	Operation string    `json:"operation"`
	Value     float64   `json:"value"`
	Threshold float64   `json:"threshold"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// createAlert creates a performance alert
func (bf *BenchmarkFramework) createAlert(alert AlertInfo) {
	alert.Timestamp = time.Now()
	
	// Log alert
	logger := bf.createPhaseLogger(alert.Phase)
	logger.Warn("Performance alert", 
		"type", alert.Type,
		"scenario", alert.Scenario,
		"operation", alert.Operation,
		"message", alert.Message)
	
	// Save alert to storage
	bf.storage.SaveAlert(alert)
}