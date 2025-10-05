// Package benchmark provides validation for performance benchmarks
package benchmark

import (
	"fmt"
	"math"
	"sort"
	"time"
)

// Validator validates benchmark results against thresholds and baselines
type Validator struct {
	config      Config
	thresholds  map[string]Threshold
}

// NewValidator creates a new validator
func NewValidator(config Config) *Validator {
	return &Validator{
		config:     config,
		thresholds: config.Thresholds,
	}
}

// ValidatePhase validates results for a specific phase
func (v *Validator) ValidatePhase(phase Phase, results []Result, baselines map[string]Baseline) *ValidationReport {
	report := &ValidationReport{
		Phase:       phase,
		Timestamp:   time.Now(),
		Results:     make([]ValidationResult, 0),
		Summary:     ValidationSummary{},
	}
	
	for _, result := range results {
		validation := v.validateResult(result, baselines)
		report.Results = append(report.Results, validation)
		
		// Update summary
		report.Summary.Total++
		if validation.HasRegressions() {
			report.Summary.Regressions++
		}
		if validation.HasImprovements() {
			report.Summary.Improvements++
		}
		if validation.HasViolations() {
			report.Summary.Violations++
		}
	}
	
	// Calculate overall scores
	report.Summary.OverallScore = v.calculateOverallScore(report.Results)
	report.Summary.PerformanceScore = v.calculatePerformanceScore(report.Results)
	report.Summary.StabilityScore = v.calculateStabilityScore(report.Results)
	
	return report
}

// validateResult validates a single result
func (v *Validator) validateResult(result Result, baselines map[string]Baseline) ValidationResult {
	validation := ValidationResult{
		Scenario:        result.Scenario,
		Operation:       result.Operation,
		Violations:      make([]Violation, 0),
		Regressions:     make([]Regression, 0),
		Improvements:    make([]Improvement, 0),
		Warnings:        make([]Warning, 0),
	}
	
	// Check threshold violations
	v.checkThresholds(result, &validation)
	
	// Check for regressions against baseline
	v.checkBaseline(result, baselines, &validation)
	
	// Check for stability issues
	v.checkStability(result, &validation)
	
	// Check for potential issues
	v.checkPotentialIssues(result, &validation)
	
	return validation
}

// checkThresholds checks if results violate configured thresholds
func (v *Validator) checkThresholds(result Result, validation *ValidationResult) {
	key := fmt.Sprintf("%s_%s", result.Scenario, result.Operation)
	threshold, exists := v.thresholds[key]
	if !exists {
		// Use default threshold
		threshold = v.getDefaultThreshold()
	}
	
	// Check duration threshold
	if threshold.MaxDuration > 0 && result.Duration > threshold.MaxDuration {
		validation.Violations = append(validation.Violations, Violation{
			Type:        "duration",
			Expected:    threshold.MaxDuration.String(),
			Actual:      result.Duration.String(),
			Severity:    "error",
			Description: "Duration exceeds maximum threshold",
		})
	}
	
	// Check memory threshold
	if threshold.MaxMemory > 0 && result.Memory.TotalAlloc > threshold.MaxMemory {
		validation.Violations = append(validation.Violations, Violation{
			Type:        "memory",
			Expected:    fmt.Sprintf("%d bytes", threshold.MaxMemory),
			Actual:      fmt.Sprintf("%d bytes", result.Memory.TotalAlloc),
			Severity:    "error",
			Description: "Memory usage exceeds maximum threshold",
		})
	}
	
	// Check allocations threshold
	if threshold.MaxAllocations > 0 && result.Memory.Allocations > threshold.MaxAllocations {
		validation.Violations = append(validation.Violations, Violation{
			Type:        "allocations",
			Expected:    fmt.Sprintf("%d", threshold.MaxAllocations),
			Actual:      fmt.Sprintf("%d", result.Memory.Allocations),
			Severity:    "warning",
			Description: "Allocations exceed maximum threshold",
		})
	}
	
	// Check GC pause threshold
	if threshold.MaxGCPause > 0 && result.Memory.GCPauseMax > threshold.MaxGCPause {
		validation.Violations = append(validation.Violations, Violation{
			Type:        "gc_pause",
			Expected:    threshold.MaxGCPause.String(),
			Actual:      result.Memory.GCPauseMax.String(),
			Severity:    "warning",
			Description: "GC pause exceeds maximum threshold",
		})
	}
}

// checkBaseline checks for regressions against baseline
func (v *Validator) checkBaseline(result Result, baselines map[string]Baseline, validation *ValidationResult) {
	key := fmt.Sprintf("%s_%s", result.Scenario, result.Operation)
	baseline, exists := baselines[key]
	if !exists {
		validation.Warnings = append(validation.Warnings, Warning{
			Type:        "no_baseline",
			Description: "No baseline found for comparison",
		})
		return
	}
	
	// Check for regressions
	v.checkDurationRegression(result, baseline, validation)
	v.checkMemoryRegression(result, baseline, validation)
	v.checkAllocationRegression(result, baseline, validation)
	v.checkThroughputRegression(result, baseline, validation)
}

// checkDurationRegression checks for duration regression
func (v *Validator) checkDurationRegression(result Result, baseline Baseline, validation *ValidationResult) {
	currentDuration := float64(result.Duration)
	baselineDuration := float64(baseline.Metrics.Duration)
	
	if baselineDuration <= 0 {
		return
	}
	
	change := (currentDuration - baselineDuration) / baselineDuration
	threshold := v.getThresholdForScenario(result.Scenario, result.Operation)
	
	if change > threshold.RegressionMargin {
		validation.Regressions = append(validation.Regressions, Regression{
			Type:           "duration",
			BaselineValue:  baseline.Metrics.Duration.String(),
			CurrentValue:   result.Duration.String(),
			ChangePercent:  change * 100,
			Severity:       v.getRegressinSeverity(change),
			Description:    fmt.Sprintf("Duration regression: %.2f%% slower", change*100),
		})
	} else if change < -0.05 { // 5% improvement threshold
		validation.Improvements = append(validation.Improvements, Improvement{
			Type:           "duration",
			BaselineValue:  baseline.Metrics.Duration.String(),
			CurrentValue:   result.Duration.String(),
			ChangePercent:  -change * 100,
			Description:    fmt.Sprintf("Duration improvement: %.2f%% faster", -change*100),
		})
	}
}

// checkMemoryRegression checks for memory regression
func (v *Validator) checkMemoryRegression(result Result, baseline Baseline, validation *ValidationResult) {
	currentMemory := float64(result.Memory.TotalAlloc)
	baselineMemory := float64(baseline.Metrics.Memory)
	
	if baselineMemory <= 0 {
		return
	}
	
	change := (currentMemory - baselineMemory) / baselineMemory
	threshold := v.getThresholdForScenario(result.Scenario, result.Operation)
	
	if change > threshold.RegressionMargin {
		validation.Regressions = append(validation.Regressions, Regression{
			Type:           "memory",
			BaselineValue:  fmt.Sprintf("%d bytes", baseline.Metrics.Memory),
			CurrentValue:   fmt.Sprintf("%d bytes", result.Memory.TotalAlloc),
			ChangePercent:  change * 100,
			Severity:       v.getRegressinSeverity(change),
			Description:    fmt.Sprintf("Memory regression: %.2f%% more memory", change*100),
		})
	} else if change < -0.05 { // 5% improvement threshold
		validation.Improvements = append(validation.Improvements, Improvement{
			Type:           "memory",
			BaselineValue:  fmt.Sprintf("%d bytes", baseline.Metrics.Memory),
			CurrentValue:   fmt.Sprintf("%d bytes", result.Memory.TotalAlloc),
			ChangePercent:  -change * 100,
			Description:    fmt.Sprintf("Memory improvement: %.2f%% less memory", -change*100),
		})
	}
}

// checkAllocationRegression checks for allocation regression
func (v *Validator) checkAllocationRegression(result Result, baseline Baseline, validation *ValidationResult) {
	currentAllocs := float64(result.Memory.Allocations)
	baselineAllocs := float64(baseline.Metrics.Allocations)
	
	if baselineAllocs <= 0 {
		return
	}
	
	change := (currentAllocs - baselineAllocs) / baselineAllocs
	threshold := v.getThresholdForScenario(result.Scenario, result.Operation)
	
	if change > threshold.RegressionMargin {
		validation.Regressions = append(validation.Regressions, Regression{
			Type:           "allocations",
			BaselineValue:  fmt.Sprintf("%d", baseline.Metrics.Allocations),
			CurrentValue:   fmt.Sprintf("%d", result.Memory.Allocations),
			ChangePercent:  change * 100,
			Severity:       v.getRegressinSeverity(change),
			Description:    fmt.Sprintf("Allocation regression: %.2f%% more allocations", change*100),
		})
	} else if change < -0.05 { // 5% improvement threshold
		validation.Improvements = append(validation.Improvements, Improvement{
			Type:           "allocations",
			BaselineValue:  fmt.Sprintf("%d", baseline.Metrics.Allocations),
			CurrentValue:   fmt.Sprintf("%d", result.Memory.Allocations),
			ChangePercent:  -change * 100,
			Description:    fmt.Sprintf("Allocation improvement: %.2f%% fewer allocations", -change*100),
		})
	}
}

// checkThroughputRegression checks for throughput regression
func (v *Validator) checkThroughputRegression(result Result, baseline Baseline, validation *ValidationResult) {
	// Calculate throughput from result metadata
	fileSize, ok := result.Metadata["file_size"].(int64)
	if !ok {
		return
	}
	
	currentThroughput := float64(fileSize) / result.Duration.Seconds()
	baselineThroughput := baseline.Metrics.Throughput
	
	if baselineThroughput <= 0 {
		return
	}
	
	change := (currentThroughput - baselineThroughput) / baselineThroughput
	threshold := v.getThresholdForScenario(result.Scenario, result.Operation)
	
	if change < -threshold.RegressionMargin {
		validation.Regressions = append(validation.Regressions, Regression{
			Type:           "throughput",
			BaselineValue:  fmt.Sprintf("%.2f MB/s", baselineThroughput/1024/1024),
			CurrentValue:   fmt.Sprintf("%.2f MB/s", currentThroughput/1024/1024),
			ChangePercent:  -change * 100,
			Severity:       v.getRegressinSeverity(-change),
			Description:    fmt.Sprintf("Throughput regression: %.2f%% slower processing", -change*100),
		})
	} else if change > 0.05 { // 5% improvement threshold
		validation.Improvements = append(validation.Improvements, Improvement{
			Type:           "throughput",
			BaselineValue:  fmt.Sprintf("%.2f MB/s", baselineThroughput/1024/1024),
			CurrentValue:   fmt.Sprintf("%.2f MB/s", currentThroughput/1024/1024),
			ChangePercent:  change * 100,
			Description:    fmt.Sprintf("Throughput improvement: %.2f%% faster processing", change*100),
		})
	}
}

// checkStability checks for stability issues
func (v *Validator) checkStability(result Result, validation *ValidationResult) {
	// Check for high variance in statistics
	if stats, ok := result.Metadata["duration_stats"].(Statistics); ok {
		if stats.CV > 0.3 { // Coefficient of variation > 30%
			validation.Warnings = append(validation.Warnings, Warning{
				Type:        "high_variance",
				Description: fmt.Sprintf("High duration variance (CV: %.2f%%)", stats.CV*100),
			})
		}
	}
	
	// Check for excessive GC activity
	if result.Memory.NumGC > 100 { // Arbitrary threshold
		validation.Warnings = append(validation.Warnings, Warning{
			Type:        "excessive_gc",
			Description: fmt.Sprintf("High GC activity: %d cycles", result.Memory.NumGC),
		})
	}
	
	// Check for long GC pauses
	if result.Memory.GCPauseMax > 100*time.Millisecond {
		validation.Warnings = append(validation.Warnings, Warning{
			Type:        "long_gc_pause",
			Description: fmt.Sprintf("Long GC pause: %v", result.Memory.GCPauseMax),
		})
	}
}

// checkPotentialIssues checks for potential performance issues
func (v *Validator) checkPotentialIssues(result Result, validation *ValidationResult) {
	// Check for memory leaks (high heap retention)
	if result.Memory.HeapInuse > 0 && result.Memory.HeapReleased > 0 {
		retention := float64(result.Memory.HeapInuse) / float64(result.Memory.HeapInuse+result.Memory.HeapReleased)
		if retention > 0.8 { // 80% retention might indicate a leak
			validation.Warnings = append(validation.Warnings, Warning{
				Type:        "potential_leak",
				Description: fmt.Sprintf("High heap retention: %.2f%%", retention*100),
			})
		}
	}
	
	// Check for excessive allocations
	if result.Memory.Allocations > 0 {
		fileSize, ok := result.Metadata["file_size"].(int64)
		if ok && fileSize > 0 {
			allocsPerByte := float64(result.Memory.Allocations) / float64(fileSize)
			if allocsPerByte > 10.0 { // More than 10 allocations per byte
				validation.Warnings = append(validation.Warnings, Warning{
					Type:        "excessive_allocations",
					Description: fmt.Sprintf("High allocation rate: %.2f allocs/byte", allocsPerByte),
				})
			}
		}
	}
}

// getThresholdForScenario gets threshold for a specific scenario
func (v *Validator) getThresholdForScenario(scenario, operation string) Threshold {
	key := fmt.Sprintf("%s_%s", scenario, operation)
	if threshold, exists := v.thresholds[key]; exists {
		return threshold
	}
	return v.getDefaultThreshold()
}

// getDefaultThreshold returns default threshold values
func (v *Validator) getDefaultThreshold() Threshold {
	return Threshold{
		MaxDuration:      10 * time.Second,
		MaxMemory:        100 * 1024 * 1024, // 100MB
		MaxAllocations:   1000000,            // 1M allocations
		MaxGCPause:       50 * time.Millisecond,
		RegressionMargin: 0.2, // 20% regression margin
	}
}

// getRegressinSeverity determines severity of regression
func (v *Validator) getRegressinSeverity(change float64) string {
	absChange := math.Abs(change)
	switch {
	case absChange > 0.5: // 50%+ change
		return "critical"
	case absChange > 0.2: // 20%+ change
		return "major"
	case absChange > 0.1: // 10%+ change
		return "minor"
	default:
		return "negligible"
	}
}

// calculateOverallScore calculates overall benchmark score
func (v *Validator) calculateOverallScore(results []ValidationResult) float64 {
	if len(results) == 0 {
		return 0.0
	}
	
	totalScore := 0.0
	for _, result := range results {
		score := 100.0 // Start with perfect score
		
		// Deduct points for violations
		for _, violation := range result.Violations {
			switch violation.Severity {
			case "critical":
				score -= 50
			case "error":
				score -= 25
			case "warning":
				score -= 10
			}
		}
		
		// Deduct points for regressions
		for _, regression := range result.Regressions {
			switch regression.Severity {
			case "critical":
				score -= 40
			case "major":
				score -= 20
			case "minor":
				score -= 10
			}
		}
		
		// Add points for improvements
		for _, improvement := range result.Improvements {
			score += 5 // Small bonus for improvements
		}
		
		// Deduct points for warnings
		score -= float64(len(result.Warnings)) * 2
		
		// Ensure score doesn't go below 0
		if score < 0 {
			score = 0
		}
		
		totalScore += score
	}
	
	return totalScore / float64(len(results))
}

// calculatePerformanceScore calculates performance-specific score
func (v *Validator) calculatePerformanceScore(results []ValidationResult) float64 {
	if len(results) == 0 {
		return 0.0
	}
	
	totalScore := 0.0
	for _, result := range results {
		score := 100.0
		
		// Focus on performance-related violations and regressions
		for _, violation := range result.Violations {
			if violation.Type == "duration" || violation.Type == "memory" {
				score -= 30
			}
		}
		
		for _, regression := range result.Regressions {
			if regression.Type == "duration" || regression.Type == "memory" || regression.Type == "throughput" {
				score -= math.Min(regression.ChangePercent, 50) // Cap at 50 points
			}
		}
		
		for _, improvement := range result.Improvements {
			if improvement.Type == "duration" || improvement.Type == "memory" || improvement.Type == "throughput" {
				score += math.Min(improvement.ChangePercent/2, 25) // Cap at 25 points
			}
		}
		
		if score < 0 {
			score = 0
		}
		
		totalScore += score
	}
	
	return totalScore / float64(len(results))
}

// calculateStabilityScore calculates stability score
func (v *Validator) calculateStabilityScore(results []ValidationResult) float64 {
	if len(results) == 0 {
		return 0.0
	}
	
	totalScore := 0.0
	for _, result := range results {
		score := 100.0
		
		// Deduct points for stability issues
		for _, warning := range result.Warnings {
			switch warning.Type {
			case "high_variance":
				score -= 20
			case "excessive_gc":
				score -= 15
			case "long_gc_pause":
				score -= 10
			case "potential_leak":
				score -= 25
			case "excessive_allocations":
				score -= 15
			}
		}
		
		if score < 0 {
			score = 0
		}
		
		totalScore += score
	}
	
	return totalScore / float64(len(results))
}

// ValidationReport holds validation results for a phase
type ValidationReport struct {
	Phase     Phase                `json:"phase"`
	Timestamp time.Time            `json:"timestamp"`
	Results   []ValidationResult   `json:"results"`
	Summary   ValidationSummary    `json:"summary"`
}

// ValidationSummary summarizes validation results
type ValidationSummary struct {
	Total            int     `json:"total"`
	Violations       int     `json:"violations"`
	Regressions      int     `json:"regressions"`
	Improvements     int     `json:"improvements"`
	Warnings         int     `json:"warnings"`
	OverallScore     float64 `json:"overall_score"`
	PerformanceScore float64 `json:"performance_score"`
	StabilityScore   float64 `json:"stability_score"`
}

// ValidationResult holds validation results for a single benchmark
type ValidationResult struct {
	Scenario     string        `json:"scenario"`
	Operation    string        `json:"operation"`
	Violations   []Violation   `json:"violations"`
	Regressions  []Regression  `json:"regressions"`
	Improvements []Improvement `json:"improvements"`
	Warnings     []Warning     `json:"warnings"`
}

// Violation represents a threshold violation
type Violation struct {
	Type        string `json:"type"`
	Expected    string `json:"expected"`
	Actual      string `json:"actual"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
}

// Regression represents a performance regression
type Regression struct {
	Type           string  `json:"type"`
	BaselineValue  string  `json:"baseline_value"`
	CurrentValue   string  `json:"current_value"`
	ChangePercent  float64 `json:"change_percent"`
	Severity       string  `json:"severity"`
	Description    string  `json:"description"`
}

// Improvement represents a performance improvement
type Improvement struct {
	Type           string  `json:"type"`
	BaselineValue  string  `json:"baseline_value"`
	CurrentValue   string  `json:"current_value"`
	ChangePercent  float64 `json:"change_percent"`
	Description    string  `json:"description"`
}

// Warning represents a potential issue
type Warning struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

// HasRegressions checks if there are any regressions
func (vr *ValidationResult) HasRegressions() bool {
	return len(vr.Regressions) > 0
}

// HasImprovements checks if there are any improvements
func (vr *ValidationResult) HasImprovements() bool {
	return len(vr.Improvements) > 0
}

// HasViolations checks if there are any violations
func (vr *ValidationResult) HasViolations() bool {
	return len(vr.Violations) > 0
}

// HasRegressions checks if the report has any regressions
func (vr *ValidationReport) HasRegressions() bool {
	return vr.Summary.Regressions > 0
}

// Summary returns a summary of the validation report
func (vr *ValidationReport) Summary() string {
	if vr.Summary.Regressions > 0 {
		return fmt.Sprintf("%d regressions detected (score: %.1f)", vr.Summary.Regressions, vr.Summary.OverallScore)
	}
	if vr.Summary.Improvements > 0 {
		return fmt.Sprintf("%d improvements detected (score: %.1f)", vr.Summary.Improvements, vr.Summary.OverallScore)
	}
	return fmt.Sprintf("No significant changes (score: %.1f)", vr.Summary.OverallScore)
}

// SortResults sorts validation results by severity
func (vr *ValidationReport) SortResults() {
	sort.Slice(vr.Results, func(i, j int) bool {
		scoreI := len(vr.Results[i].Violations)*3 + len(vr.Results[i].Regressions)*2 + len(vr.Results[i].Warnings)
		scoreJ := len(vr.Results[j].Violations)*3 + len(vr.Results[j].Regressions)*2 + len(vr.Results[j].Warnings)
		return scoreI > scoreJ
	})
}

// GetMostCriticalIssues returns the most critical issues
func (vr *ValidationReport) GetMostCriticalIssues(limit int) []string {
	issues := make([]string, 0)
	
	for _, result := range vr.Results {
		for _, violation := range result.Violations {
			if violation.Severity == "critical" || violation.Severity == "error" {
				issues = append(issues, fmt.Sprintf("%s.%s: %s", result.Scenario, result.Operation, violation.Description))
				if len(issues) >= limit {
					return issues
				}
			}
		}
		
		for _, regression := range result.Regressions {
			if regression.Severity == "critical" || regression.Severity == "major" {
				issues = append(issues, fmt.Sprintf("%s.%s: %s", result.Scenario, result.Operation, regression.Description))
				if len(issues) >= limit {
					return issues
				}
			}
		}
	}
	
	return issues
}