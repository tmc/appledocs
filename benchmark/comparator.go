// Package benchmark provides comparison functionality for benchmark results
package benchmark

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/montanaflynn/stats"
)

// Comparator compares benchmark results between phases
type Comparator struct {
	config Config
}

// NewComparator creates a new comparator
func NewComparator(config Config) *Comparator {
	return &Comparator{config: config}
}

// ComparePhases compares results between two phases
func (c *Comparator) ComparePhases(baseline, current []Result) *PhaseComparison {
	comparison := &PhaseComparison{
		BaselinePhase: PhaseBaseline,
		CurrentPhase:  current[0].Phase,
		Timestamp:     time.Now(),
		Comparisons:   make([]OperationComparison, 0),
	}

	// Group results by scenario and operation
	baselineMap := c.groupResults(baseline)
	currentMap := c.groupResults(current)

	// Compare each operation
	for key, baselineResults := range baselineMap {
		currentResults, exists := currentMap[key]
		if !exists {
			comparison.MissingOperations = append(comparison.MissingOperations, key)
			continue
		}

		opComparison := c.compareOperation(key, baselineResults, currentResults)
		comparison.Comparisons = append(comparison.Comparisons, opComparison)
	}

	// Find new operations
	for key := range currentMap {
		if _, exists := baselineMap[key]; !exists {
			comparison.NewOperations = append(comparison.NewOperations, key)
		}
	}

	// Calculate summary
	comparison.calculateSummary()

	return comparison
}

// PhaseComparison holds comparison results between phases
type PhaseComparison struct {
	BaselinePhase     Phase                  `json:"baseline_phase"`
	CurrentPhase      Phase                  `json:"current_phase"`
	Timestamp         time.Time              `json:"timestamp"`
	Comparisons       []OperationComparison  `json:"comparisons"`
	Summary           ComparisonSummary      `json:"summary"`
	MissingOperations []string               `json:"missing_operations,omitempty"`
	NewOperations     []string               `json:"new_operations,omitempty"`
}

// OperationComparison holds comparison for a single operation
type OperationComparison struct {
	Scenario           string              `json:"scenario"`
	Operation          string              `json:"operation"`
	BaselineMetrics    AggregatedMetrics   `json:"baseline_metrics"`
	CurrentMetrics     AggregatedMetrics   `json:"current_metrics"`
	Changes            MetricChanges       `json:"changes"`
	StatisticalAnalysis StatAnalysis       `json:"statistical_analysis"`
	Verdict            string              `json:"verdict"` // improved, regressed, unchanged
}

// AggregatedMetrics holds aggregated metrics for comparison
type AggregatedMetrics struct {
	Duration        DurationMetrics    `json:"duration"`
	Memory          MemoryMetrics      `json:"memory"`
	Throughput      float64            `json:"throughput"`
	ErrorRate       float64            `json:"error_rate"`
	SampleSize      int                `json:"sample_size"`
}

// DurationMetrics holds duration-related metrics
type DurationMetrics struct {
	Mean   time.Duration `json:"mean"`
	Median time.Duration `json:"median"`
	P95    time.Duration `json:"p95"`
	P99    time.Duration `json:"p99"`
	StdDev time.Duration `json:"std_dev"`
}

// MemoryMetrics holds memory-related metrics
type MemoryMetrics struct {
	MeanAlloc       int64 `json:"mean_alloc"`
	MeanAllocations int64 `json:"mean_allocations"`
	MaxHeap         int64 `json:"max_heap"`
	GCPauseMean     time.Duration `json:"gc_pause_mean"`
	GCPauseP99      time.Duration `json:"gc_pause_p99"`
}

// MetricChanges holds percentage changes between metrics
type MetricChanges struct {
	DurationChange      float64 `json:"duration_change_percent"`
	MemoryChange        float64 `json:"memory_change_percent"`
	AllocationChange    float64 `json:"allocation_change_percent"`
	ThroughputChange    float64 `json:"throughput_change_percent"`
	GCPauseChange       float64 `json:"gc_pause_change_percent"`
}

// StatAnalysis holds statistical analysis results
type StatAnalysis struct {
	TTest            TTestResult `json:"t_test"`
	MannWhitneyU     float64     `json:"mann_whitney_u"`
	EffectSize       float64     `json:"effect_size"`
	PowerAnalysis    float64     `json:"power_analysis"`
	SignificantChange bool       `json:"significant_change"`
}

// TTestResult holds t-test results
type TTestResult struct {
	TStatistic float64 `json:"t_statistic"`
	PValue     float64 `json:"p_value"`
	DegreesOfFreedom int `json:"degrees_of_freedom"`
}

// ComparisonSummary holds summary statistics
type ComparisonSummary struct {
	TotalOperations   int     `json:"total_operations"`
	ImprovedCount     int     `json:"improved_count"`
	RegressedCount    int     `json:"regressed_count"`
	UnchangedCount    int     `json:"unchanged_count"`
	OverallImprovement float64 `json:"overall_improvement_percent"`
	SignificantChanges int     `json:"significant_changes"`
}

// groupResults groups results by scenario and operation
func (c *Comparator) groupResults(results []Result) map[string][]Result {
	grouped := make(map[string][]Result)
	for _, result := range results {
		key := fmt.Sprintf("%s/%s", result.Scenario, result.Operation)
		grouped[key] = append(grouped[key], result)
	}
	return grouped
}

// compareOperation compares results for a single operation
func (c *Comparator) compareOperation(key string, baseline, current []Result) OperationComparison {
	comparison := OperationComparison{
		Scenario:        baseline[0].Scenario,
		Operation:       baseline[0].Operation,
		BaselineMetrics: c.aggregateMetrics(baseline),
		CurrentMetrics:  c.aggregateMetrics(current),
	}

	// Calculate changes
	comparison.Changes = c.calculateChanges(comparison.BaselineMetrics, comparison.CurrentMetrics)

	// Perform statistical analysis
	comparison.StatisticalAnalysis = c.performStatisticalAnalysis(baseline, current)

	// Determine verdict
	comparison.Verdict = c.determineVerdict(comparison.Changes, comparison.StatisticalAnalysis)

	return comparison
}

// aggregateMetrics aggregates metrics from multiple results
func (c *Comparator) aggregateMetrics(results []Result) AggregatedMetrics {
	if len(results) == 0 {
		return AggregatedMetrics{}
	}

	// Extract durations and memory values
	durations := make([]float64, len(results))
	allocBytes := make([]float64, len(results))
	allocCounts := make([]float64, len(results))
	gcPauses := make([]float64, len(results))
	heapSizes := make([]int64, len(results))

	for i, r := range results {
		durations[i] = float64(r.Duration)
		allocBytes[i] = float64(r.Memory.TotalAlloc)
		allocCounts[i] = float64(r.Memory.Allocations)
		gcPauses[i] = float64(r.Memory.GCPauseAvg)
		heapSizes[i] = r.Memory.HeapAlloc
	}

	// Calculate statistics
	durationStats := calculateStatistics(durations)
	allocStats := calculateStatistics(allocBytes)
	allocCountStats := calculateStatistics(allocCounts)
	gcStats := calculateStatistics(gcPauses)

	// Find max heap
	maxHeap := int64(0)
	for _, h := range heapSizes {
		if h > maxHeap {
			maxHeap = h
		}
	}

	// Calculate throughput (operations per second)
	throughput := 0.0
	if durationStats.Mean > 0 {
		throughput = 1e9 / durationStats.Mean // Convert nanoseconds to ops/sec
	}

	return AggregatedMetrics{
		Duration: DurationMetrics{
			Mean:   time.Duration(durationStats.Mean),
			Median: time.Duration(durationStats.Median),
			P95:    time.Duration(durationStats.P95),
			P99:    time.Duration(durationStats.P99),
			StdDev: time.Duration(durationStats.StdDev),
		},
		Memory: MemoryMetrics{
			MeanAlloc:       int64(allocStats.Mean),
			MeanAllocations: int64(allocCountStats.Mean),
			MaxHeap:         maxHeap,
			GCPauseMean:     time.Duration(gcStats.Mean),
			GCPauseP99:      time.Duration(gcStats.P99),
		},
		Throughput: throughput,
		SampleSize: len(results),
	}
}

// calculateChanges calculates percentage changes between metrics
func (c *Comparator) calculateChanges(baseline, current AggregatedMetrics) MetricChanges {
	return MetricChanges{
		DurationChange:   percentageChange(float64(baseline.Duration.Mean), float64(current.Duration.Mean)),
		MemoryChange:     percentageChange(float64(baseline.Memory.MeanAlloc), float64(current.Memory.MeanAlloc)),
		AllocationChange: percentageChange(float64(baseline.Memory.MeanAllocations), float64(current.Memory.MeanAllocations)),
		ThroughputChange: percentageChange(baseline.Throughput, current.Throughput),
		GCPauseChange:    percentageChange(float64(baseline.Memory.GCPauseMean), float64(current.Memory.GCPauseMean)),
	}
}

// percentageChange calculates percentage change
func percentageChange(baseline, current float64) float64 {
	if baseline == 0 {
		if current == 0 {
			return 0
		}
		return 100 // 100% increase from 0
	}
	return ((current - baseline) / baseline) * 100
}

// performStatisticalAnalysis performs statistical tests
func (c *Comparator) performStatisticalAnalysis(baseline, current []Result) StatAnalysis {
	// Extract duration values for analysis
	baselineValues := make([]float64, len(baseline))
	currentValues := make([]float64, len(current))

	for i, r := range baseline {
		baselineValues[i] = float64(r.Duration)
	}
	for i, r := range current {
		currentValues[i] = float64(r.Duration)
	}

	// Perform t-test
	tTest := performTTest(baselineValues, currentValues)

	// Calculate effect size (Cohen's d)
	effectSize := calculateEffectSize(baselineValues, currentValues)

	// Determine if change is statistically significant
	significant := tTest.PValue < 0.05 && math.Abs(effectSize) > 0.5

	return StatAnalysis{
		TTest:             tTest,
		EffectSize:        effectSize,
		SignificantChange: significant,
		PowerAnalysis:     calculatePower(len(baselineValues), effectSize),
	}
}

// performTTest performs Welch's t-test (for unequal variances)
func performTTest(sample1, sample2 []float64) TTestResult {
	n1 := float64(len(sample1))
	n2 := float64(len(sample2))

	if n1 < 2 || n2 < 2 {
		return TTestResult{}
	}

	// Calculate means
	mean1, _ := stats.Mean(sample1)
	mean2, _ := stats.Mean(sample2)

	// Calculate variances
	var1, _ := stats.Variance(sample1)
	var2, _ := stats.Variance(sample2)

	// Welch's t-statistic
	se := math.Sqrt(var1/n1 + var2/n2)
	if se == 0 {
		return TTestResult{}
	}

	t := (mean1 - mean2) / se

	// Degrees of freedom (Welch-Satterthwaite equation)
	df := math.Pow(var1/n1+var2/n2, 2) / 
		(math.Pow(var1/n1, 2)/(n1-1) + math.Pow(var2/n2, 2)/(n2-1))

	// Approximate p-value (would need proper t-distribution CDF)
	// This is a simplified approximation
	pValue := 2 * (1 - normalCDF(math.Abs(t)))

	return TTestResult{
		TStatistic:       t,
		PValue:           pValue,
		DegreesOfFreedom: int(df),
	}
}

// calculateEffectSize calculates Cohen's d effect size
func calculateEffectSize(sample1, sample2 []float64) float64 {
	mean1, _ := stats.Mean(sample1)
	mean2, _ := stats.Mean(sample2)
	
	var1, _ := stats.Variance(sample1)
	var2, _ := stats.Variance(sample2)
	
	// Pooled standard deviation
	n1 := float64(len(sample1))
	n2 := float64(len(sample2))
	pooledSD := math.Sqrt(((n1-1)*var1 + (n2-1)*var2) / (n1 + n2 - 2))
	
	if pooledSD == 0 {
		return 0
	}
	
	return (mean1 - mean2) / pooledSD
}

// calculatePower calculates statistical power
func calculatePower(sampleSize int, effectSize float64) float64 {
	// Simplified power calculation
	// Real implementation would use non-central t-distribution
	n := float64(sampleSize)
	d := math.Abs(effectSize)
	
	// Approximate power for alpha = 0.05
	z := d * math.Sqrt(n/2)
	power := normalCDF(z - 1.96)
	
	return power
}

// normalCDF approximates the normal cumulative distribution function
func normalCDF(x float64) float64 {
	// Approximation of the error function
	a1 := 0.254829592
	a2 := -0.284496736
	a3 := 1.421413741
	a4 := -1.453152027
	a5 := 1.061405429
	p := 0.3275911

	sign := 1.0
	if x < 0 {
		sign = -1.0
	}
	x = math.Abs(x) / math.Sqrt(2.0)

	t := 1.0 / (1.0 + p*x)
	y := 1.0 - (((((a5*t+a4)*t)+a3)*t+a2)*t+a1)*t*math.Exp(-x*x)

	return 0.5 * (1.0 + sign*y)
}

// determineVerdict determines the overall verdict
func (c *Comparator) determineVerdict(changes MetricChanges, stats StatAnalysis) string {
	// Get threshold for this scenario (would need scenario context)
	// For now, use a default threshold
	margin := 10.0 // 10% default regression margin

	// Check for significant regression
	if changes.DurationChange > margin && stats.SignificantChange {
		return "regressed"
	}

	// Check for significant improvement
	if changes.DurationChange < -margin && stats.SignificantChange {
		return "improved"
	}

	// Check memory regression
	if changes.MemoryChange > margin*2 { // Higher threshold for memory
		return "regressed"
	}

	return "unchanged"
}

// calculateSummary calculates comparison summary
func (pc *PhaseComparison) calculateSummary() {
	pc.Summary.TotalOperations = len(pc.Comparisons)

	totalImprovement := 0.0
	significantCount := 0

	for _, comp := range pc.Comparisons {
		switch comp.Verdict {
		case "improved":
			pc.Summary.ImprovedCount++
			totalImprovement += -comp.Changes.DurationChange
		case "regressed":
			pc.Summary.RegressedCount++
			totalImprovement += -comp.Changes.DurationChange
		case "unchanged":
			pc.Summary.UnchangedCount++
		}

		if comp.StatisticalAnalysis.SignificantChange {
			significantCount++
		}
	}

	pc.Summary.SignificantChanges = significantCount
	if pc.Summary.TotalOperations > 0 {
		pc.Summary.OverallImprovement = totalImprovement / float64(pc.Summary.TotalOperations)
	}
}

// Improvements returns list of improvements
func (pc *PhaseComparison) Improvements() []string {
	improvements := make([]string, 0)
	for _, comp := range pc.Comparisons {
		if comp.Verdict == "improved" {
			improvements = append(improvements, fmt.Sprintf(
				"%s/%s: %.1f%% faster (%.2fms → %.2fms)",
				comp.Scenario, comp.Operation,
				-comp.Changes.DurationChange,
				comp.BaselineMetrics.Duration.Mean.Seconds()*1000,
				comp.CurrentMetrics.Duration.Mean.Seconds()*1000,
			))
		}
	}
	return improvements
}

// Regressions returns list of regressions
func (pc *PhaseComparison) Regressions() []string {
	regressions := make([]string, 0)
	for _, comp := range pc.Comparisons {
		if comp.Verdict == "regressed" {
			regressions = append(regressions, fmt.Sprintf(
				"%s/%s: %.1f%% slower (%.2fms → %.2fms)",
				comp.Scenario, comp.Operation,
				comp.Changes.DurationChange,
				comp.BaselineMetrics.Duration.Mean.Seconds()*1000,
				comp.CurrentMetrics.Duration.Mean.Seconds()*1000,
			))
		}
	}
	return regressions
}

// GenerateComparisonReport generates a detailed comparison report
func (c *Comparator) GenerateComparisonReport(comparison *PhaseComparison) *ComparisonReport {
	report := &ComparisonReport{
		Comparison: comparison,
		Charts:     c.generateCharts(comparison),
		Analysis:   c.generateAnalysis(comparison),
	}

	return report
}

// ComparisonReport holds the full comparison report
type ComparisonReport struct {
	Comparison *PhaseComparison         `json:"comparison"`
	Charts     []Chart                  `json:"charts"`
	Analysis   map[string]string        `json:"analysis"`
}

// Chart represents a visualization chart
type Chart struct {
	Type   string                 `json:"type"`
	Title  string                 `json:"title"`
	Data   map[string]interface{} `json:"data"`
}

// generateCharts generates visualization data
func (c *Comparator) generateCharts(comparison *PhaseComparison) []Chart {
	charts := make([]Chart, 0)

	// Performance comparison bar chart
	perfData := make(map[string]interface{})
	labels := make([]string, 0)
	baselineValues := make([]float64, 0)
	currentValues := make([]float64, 0)

	for _, comp := range comparison.Comparisons {
		labels = append(labels, comp.Operation)
		baselineValues = append(baselineValues, comp.BaselineMetrics.Duration.Mean.Seconds()*1000)
		currentValues = append(currentValues, comp.CurrentMetrics.Duration.Mean.Seconds()*1000)
	}

	perfData["labels"] = labels
	perfData["baseline"] = baselineValues
	perfData["current"] = currentValues

	charts = append(charts, Chart{
		Type:  "bar",
		Title: "Performance Comparison",
		Data:  perfData,
	})

	return charts
}

// generateAnalysis generates textual analysis
func (c *Comparator) generateAnalysis(comparison *PhaseComparison) map[string]string {
	analysis := make(map[string]string)

	// Overall assessment
	if comparison.Summary.OverallImprovement > 0 {
		analysis["overall"] = fmt.Sprintf(
			"Overall performance improved by %.1f%% across %d operations",
			comparison.Summary.OverallImprovement,
			comparison.Summary.TotalOperations,
		)
	} else if comparison.Summary.OverallImprovement < 0 {
		analysis["overall"] = fmt.Sprintf(
			"Overall performance regressed by %.1f%% across %d operations",
			-comparison.Summary.OverallImprovement,
			comparison.Summary.TotalOperations,
		)
	} else {
		analysis["overall"] = "No significant overall performance change detected"
	}

	// Statistical significance
	analysis["significance"] = fmt.Sprintf(
		"%d out of %d operations showed statistically significant changes",
		comparison.Summary.SignificantChanges,
		comparison.Summary.TotalOperations,
	)

	// Top improvements
	improvements := comparison.Improvements()
	if len(improvements) > 0 {
		analysis["top_improvements"] = fmt.Sprintf(
			"Top improvements: %s",
			improvements[0],
		)
	}

	// Top regressions
	regressions := comparison.Regressions()
	if len(regressions) > 0 {
		analysis["top_regressions"] = fmt.Sprintf(
			"Top regressions: %s",
			regressions[0],
		)
	}

	return analysis
}

// RankOperationsByImprovement ranks operations by improvement percentage
func (pc *PhaseComparison) RankOperationsByImprovement() []OperationRanking {
	rankings := make([]OperationRanking, 0)

	for _, comp := range pc.Comparisons {
		ranking := OperationRanking{
			Scenario:         comp.Scenario,
			Operation:        comp.Operation,
			ImprovementPercent: -comp.Changes.DurationChange,
			MemoryImprovement: -comp.Changes.MemoryChange,
			Significant:      comp.StatisticalAnalysis.SignificantChange,
		}
		rankings = append(rankings, ranking)
	}

	// Sort by improvement percentage
	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].ImprovementPercent > rankings[j].ImprovementPercent
	})

	return rankings
}

// OperationRanking holds ranking information for an operation
type OperationRanking struct {
	Scenario           string  `json:"scenario"`
	Operation          string  `json:"operation"`
	ImprovementPercent float64 `json:"improvement_percent"`
	MemoryImprovement  float64 `json:"memory_improvement_percent"`
	Significant        bool    `json:"statistically_significant"`
}