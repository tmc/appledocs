// Package benchmark provides reporting functionality for benchmark results
package benchmark

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// Reporter generates reports from benchmark results
type Reporter struct {
	config     Config
	outputDir  string
	templates  map[string]*template.Template
}

// NewReporter creates a new reporter
func NewReporter(config Config) *Reporter {
	return &Reporter{
		config:    config,
		outputDir: filepath.Join(config.OutputDir, "reports"),
		templates: make(map[string]*template.Template),
	}
}

// GenerateReport generates a comprehensive report
func (r *Reporter) GenerateReport(data *ReportData) error {
	// Create output directory
	if err := os.MkdirAll(r.outputDir, 0755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}
	
	// Generate different report formats
	if err := r.generateJSONReport(data); err != nil {
		return fmt.Errorf("generate JSON report: %w", err)
	}
	
	if r.config.ReportFormat == "html" || r.config.ReportFormat == "all" {
		if err := r.generateHTMLReport(data); err != nil {
			return fmt.Errorf("generate HTML report: %w", err)
		}
	}
	
	if r.config.ReportFormat == "markdown" || r.config.ReportFormat == "all" {
		if err := r.generateMarkdownReport(data); err != nil {
			return fmt.Errorf("generate Markdown report: %w", err)
		}
	}
	
	// Generate visualization if enabled
	if r.config.Visualization {
		if err := r.generateVisualization(data); err != nil {
			return fmt.Errorf("generate visualization: %w", err)
		}
	}
	
	return nil
}

// generateJSONReport generates a JSON report
func (r *Reporter) generateJSONReport(data *ReportData) error {
	filename := filepath.Join(r.outputDir, "benchmark_report.json")
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create JSON file: %w", err)
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	
	return encoder.Encode(data)
}

// generateHTMLReport generates an HTML report
func (r *Reporter) generateHTMLReport(data *ReportData) error {
	// Load HTML template
	tmpl, err := r.getHTMLTemplate()
	if err != nil {
		return fmt.Errorf("load HTML template: %w", err)
	}
	
	filename := filepath.Join(r.outputDir, "benchmark_report.html")
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create HTML file: %w", err)
	}
	defer file.Close()
	
	// Execute template
	return tmpl.Execute(file, data)
}

// generateMarkdownReport generates a Markdown report
func (r *Reporter) generateMarkdownReport(data *ReportData) error {
	filename := filepath.Join(r.outputDir, "benchmark_report.md")
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create Markdown file: %w", err)
	}
	defer file.Close()
	
	return r.writeMarkdownReport(file, data)
}

// writeMarkdownReport writes the Markdown report
func (r *Reporter) writeMarkdownReport(w io.Writer, data *ReportData) error {
	// Title and metadata
	fmt.Fprintf(w, "# AppLeDocs Benchmark Report\n\n")
	fmt.Fprintf(w, "**Generated:** %s\n", data.GeneratedAt.Format(time.RFC3339))
	fmt.Fprintf(w, "**Platform:** %s/%s\n", data.Platform.OS, data.Platform.Arch)
	fmt.Fprintf(w, "**Go Version:** %s\n", data.Platform.GoVersion)
	fmt.Fprintf(w, "**Git Commit:** %s\n\n", data.GitCommit)
	
	// Executive Summary
	fmt.Fprintf(w, "## Executive Summary\n\n")
	fmt.Fprintf(w, "- **Total Phases:** %d\n", len(data.PhaseResults))
	fmt.Fprintf(w, "- **Total Scenarios:** %d\n", data.Summary.TotalScenarios)
	fmt.Fprintf(w, "- **Total Operations:** %d\n", data.Summary.TotalOperations)
	fmt.Fprintf(w, "- **Overall Score:** %.1f/100\n", data.Summary.OverallScore)
	fmt.Fprintf(w, "- **Performance Score:** %.1f/100\n", data.Summary.PerformanceScore)
	fmt.Fprintf(w, "- **Stability Score:** %.1f/100\n\n", data.Summary.StabilityScore)
	
	// Key Findings
	fmt.Fprintf(w, "## Key Findings\n\n")
	for _, finding := range data.Summary.KeyFindings {
		fmt.Fprintf(w, "- %s\n", finding)
	}
	fmt.Fprintf(w, "\n")
	
	// Phase Results
	fmt.Fprintf(w, "## Phase Results\n\n")
	for _, phase := range data.PhaseResults {
		r.writePhaseSection(w, phase)
	}
	
	// Comparisons
	if len(data.Comparisons) > 0 {
		fmt.Fprintf(w, "## Phase Comparisons\n\n")
		for _, comparison := range data.Comparisons {
			r.writeComparisonSection(w, comparison)
		}
	}
	
	// Validation Results
	if len(data.Validations) > 0 {
		fmt.Fprintf(w, "## Validation Results\n\n")
		for _, validation := range data.Validations {
			r.writeValidationSection(w, validation)
		}
	}
	
	// Recommendations
	if len(data.Recommendations) > 0 {
		fmt.Fprintf(w, "## Recommendations\n\n")
		for i, rec := range data.Recommendations {
			fmt.Fprintf(w, "%d. **%s** (%s priority)\n", i+1, rec.Title, rec.Priority)
			fmt.Fprintf(w, "   %s\n\n", rec.Description)
		}
	}
	
	// Appendix
	fmt.Fprintf(w, "## Appendix\n\n")
	fmt.Fprintf(w, "### Configuration\n\n")
	fmt.Fprintf(w, "```json\n")
	configJSON, _ := json.MarshalIndent(data.Config, "", "  ")
	fmt.Fprintf(w, "%s\n", configJSON)
	fmt.Fprintf(w, "```\n\n")
	
	return nil
}

// writePhaseSection writes a phase section
func (r *Reporter) writePhaseSection(w io.Writer, phase *PhaseResult) {
	fmt.Fprintf(w, "### Phase: %s\n\n", phase.Phase)
	fmt.Fprintf(w, "**Duration:** %s\n", phase.Duration)
	fmt.Fprintf(w, "**Results:** %d\n", len(phase.Results))
	fmt.Fprintf(w, "**Score:** %.1f/100\n\n", phase.Score)
	
	// Top performers
	if len(phase.TopPerformers) > 0 {
		fmt.Fprintf(w, "#### Top Performers\n\n")
		fmt.Fprintf(w, "| Scenario | Operation | Duration | Memory | Score |\n")
		fmt.Fprintf(w, "|----------|-----------|----------|--------|---------|\n")
		for _, perf := range phase.TopPerformers {
			fmt.Fprintf(w, "| %s | %s | %s | %s | %.1f |\n",
				perf.Scenario, perf.Operation, perf.Duration, 
				formatBytes(perf.Memory), perf.Score)
		}
		fmt.Fprintf(w, "\n")
	}
	
	// Issues
	if len(phase.Issues) > 0 {
		fmt.Fprintf(w, "#### Issues\n\n")
		for _, issue := range phase.Issues {
			fmt.Fprintf(w, "- **%s:** %s\n", strings.Title(issue.Severity), issue.Description)
		}
		fmt.Fprintf(w, "\n")
	}
}

// writeComparisonSection writes a comparison section
func (r *Reporter) writeComparisonSection(w io.Writer, comparison *PhaseComparison) {
	fmt.Fprintf(w, "### %s vs %s\n\n", comparison.Phase1, comparison.Phase2)
	
	summary := comparison.Summary
	fmt.Fprintf(w, "**Total Comparisons:** %d\n", summary.Total)
	fmt.Fprintf(w, "**Improvements:** %d\n", summary.Improvements)
	fmt.Fprintf(w, "**Regressions:** %d\n", summary.Regressions)
	fmt.Fprintf(w, "**Significant Changes:** %d\n\n", summary.Significant)
	
	// Significant regressions
	regressions := comparison.GetSignificantRegressions()
	if len(regressions) > 0 {
		fmt.Fprintf(w, "#### Significant Regressions\n\n")
		fmt.Fprintf(w, "| Scenario | Operation | Change | P-Value |\n")
		fmt.Fprintf(w, "|----------|-----------|--------|---------|\n")
		for _, reg := range regressions {
			fmt.Fprintf(w, "| %s | %s | %.1f%% | %.4f |\n",
				reg.Scenario, reg.Operation, 
				reg.Changes.Duration.Relative*100, reg.Significance.PValue)
		}
		fmt.Fprintf(w, "\n")
	}
	
	// Significant improvements
	improvements := comparison.GetSignificantImprovements()
	if len(improvements) > 0 {
		fmt.Fprintf(w, "#### Significant Improvements\n\n")
		fmt.Fprintf(w, "| Scenario | Operation | Change | P-Value |\n")
		fmt.Fprintf(w, "|----------|-----------|--------|---------|\n")
		for _, imp := range improvements {
			fmt.Fprintf(w, "| %s | %s | %.1f%% | %.4f |\n",
				imp.Scenario, imp.Operation, 
				-imp.Changes.Duration.Relative*100, imp.Significance.PValue)
		}
		fmt.Fprintf(w, "\n")
	}
}

// writeValidationSection writes a validation section
func (r *Reporter) writeValidationSection(w io.Writer, validation *ValidationReport) {
	fmt.Fprintf(w, "### Validation: %s\n\n", validation.Phase)
	
	summary := validation.Summary
	fmt.Fprintf(w, "**Overall Score:** %.1f/100\n", summary.OverallScore)
	fmt.Fprintf(w, "**Performance Score:** %.1f/100\n", summary.PerformanceScore)
	fmt.Fprintf(w, "**Stability Score:** %.1f/100\n", summary.StabilityScore)
	fmt.Fprintf(w, "**Total Issues:** %d\n\n", summary.Violations+summary.Regressions)
	
	// Critical issues
	criticalIssues := validation.GetMostCriticalIssues(5)
	if len(criticalIssues) > 0 {
		fmt.Fprintf(w, "#### Critical Issues\n\n")
		for _, issue := range criticalIssues {
			fmt.Fprintf(w, "- %s\n", issue)
		}
		fmt.Fprintf(w, "\n")
	}
}

// generateVisualization generates visualization files
func (r *Reporter) generateVisualization(data *ReportData) error {
	// Generate Chart.js HTML file
	filename := filepath.Join(r.outputDir, "charts.html")
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create charts file: %w", err)
	}
	defer file.Close()
	
	return r.writeChartsHTML(file, data)
}

// writeChartsHTML writes the charts HTML
func (r *Reporter) writeChartsHTML(w io.Writer, data *ReportData) error {
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
    <title>AppLeDocs Benchmark Charts</title>
    <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; }
        .chart-container { width: 800px; height: 400px; margin: 20px 0; }
        h1, h2 { color: #333; }
    </style>
</head>
<body>
    <h1>AppLeDocs Benchmark Visualizations</h1>
`)
	
	// Performance over time chart
	fmt.Fprintf(w, `
    <h2>Performance Over Time</h2>
    <div class="chart-container">
        <canvas id="performanceChart"></canvas>
    </div>
    
    <script>
    const performanceCtx = document.getElementById('performanceChart').getContext('2d');
    const performanceChart = new Chart(performanceCtx, {
        type: 'line',
        data: {
            labels: %s,
            datasets: [{
                label: 'Average Duration (ms)',
                data: %s,
                borderColor: 'rgb(75, 192, 192)',
                tension: 0.1
            }]
        },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            scales: {
                y: {
                    beginAtZero: true
                }
            }
        }
    });
    </script>
`, r.getPhaseLabels(data), r.getPerformanceData(data))
	
	// Memory usage chart
	fmt.Fprintf(w, `
    <h2>Memory Usage</h2>
    <div class="chart-container">
        <canvas id="memoryChart"></canvas>
    </div>
    
    <script>
    const memoryCtx = document.getElementById('memoryChart').getContext('2d');
    const memoryChart = new Chart(memoryCtx, {
        type: 'bar',
        data: {
            labels: %s,
            datasets: [{
                label: 'Average Memory (MB)',
                data: %s,
                backgroundColor: 'rgba(255, 99, 132, 0.2)',
                borderColor: 'rgba(255, 99, 132, 1)',
                borderWidth: 1
            }]
        },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            scales: {
                y: {
                    beginAtZero: true
                }
            }
        }
    });
    </script>
`, r.getPhaseLabels(data), r.getMemoryData(data))
	
	fmt.Fprintf(w, `
</body>
</html>`)
	
	return nil
}

// getHTMLTemplate returns the HTML template
func (r *Reporter) getHTMLTemplate() (*template.Template, error) {
	if tmpl, exists := r.templates["html"]; exists {
		return tmpl, nil
	}
	
	// Load or create HTML template
	tmplContent := r.getHTMLTemplateContent()
	tmpl, err := template.New("html").Parse(tmplContent)
	if err != nil {
		return nil, fmt.Errorf("parse HTML template: %w", err)
	}
	
	r.templates["html"] = tmpl
	return tmpl, nil
}

// getHTMLTemplateContent returns the HTML template content
func (r *Reporter) getHTMLTemplateContent() string {
	return `<!DOCTYPE html>
<html>
<head>
    <title>AppLeDocs Benchmark Report</title>
    <style>
        body { 
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif; 
            max-width: 1200px; 
            margin: 0 auto; 
            padding: 20px;
            line-height: 1.6;
        }
        .header { 
            border-bottom: 2px solid #007AFF; 
            padding-bottom: 20px; 
            margin-bottom: 30px; 
        }
        .summary { 
            background-color: #f8f9fa; 
            padding: 20px; 
            border-radius: 8px; 
            margin-bottom: 30px; 
        }
        .phase { 
            margin-bottom: 40px; 
            border: 1px solid #e9ecef; 
            border-radius: 8px; 
            padding: 20px; 
        }
        .phase h3 { 
            color: #007AFF; 
            margin-top: 0; 
        }
        table { 
            width: 100%; 
            border-collapse: collapse; 
            margin: 20px 0; 
        }
        th, td { 
            border: 1px solid #dee2e6; 
            padding: 12px; 
            text-align: left; 
        }
        th { 
            background-color: #f8f9fa; 
            font-weight: 600; 
        }
        .metric { 
            display: inline-block; 
            margin: 10px 20px 10px 0; 
        }
        .metric-label { 
            font-weight: 600; 
            color: #666; 
        }
        .metric-value { 
            font-size: 1.2em; 
            color: #333; 
        }
        .improvement { color: #28a745; }
        .regression { color: #dc3545; }
        .stable { color: #6c757d; }
    </style>
</head>
<body>
    <div class="header">
        <h1>AppLeDocs Benchmark Report</h1>
        <p><strong>Generated:</strong> {{.GeneratedAt.Format "2006-01-02 15:04:05"}}</p>
        <p><strong>Platform:</strong> {{.Platform.OS}}/{{.Platform.Arch}}</p>
        <p><strong>Go Version:</strong> {{.Platform.GoVersion}}</p>
        <p><strong>Git Commit:</strong> {{.GitCommit}}</p>
    </div>
    
    <div class="summary">
        <h2>Executive Summary</h2>
        <div class="metric">
            <div class="metric-label">Overall Score</div>
            <div class="metric-value">{{printf "%.1f" .Summary.OverallScore}}/100</div>
        </div>
        <div class="metric">
            <div class="metric-label">Performance Score</div>
            <div class="metric-value">{{printf "%.1f" .Summary.PerformanceScore}}/100</div>
        </div>
        <div class="metric">
            <div class="metric-label">Stability Score</div>
            <div class="metric-value">{{printf "%.1f" .Summary.StabilityScore}}/100</div>
        </div>
        <div class="metric">
            <div class="metric-label">Total Phases</div>
            <div class="metric-value">{{len .PhaseResults}}</div>
        </div>
    </div>
    
    {{range .PhaseResults}}
    <div class="phase">
        <h3>Phase: {{.Phase}}</h3>
        <p><strong>Duration:</strong> {{.Duration}}</p>
        <p><strong>Score:</strong> {{printf "%.1f" .Score}}/100</p>
        <p><strong>Results:</strong> {{len .Results}}</p>
        
        {{if .TopPerformers}}
        <h4>Top Performers</h4>
        <table>
            <tr>
                <th>Scenario</th>
                <th>Operation</th>
                <th>Duration</th>
                <th>Memory</th>
                <th>Score</th>
            </tr>
            {{range .TopPerformers}}
            <tr>
                <td>{{.Scenario}}</td>
                <td>{{.Operation}}</td>
                <td>{{.Duration}}</td>
                <td>{{.Memory}}</td>
                <td>{{printf "%.1f" .Score}}</td>
            </tr>
            {{end}}
        </table>
        {{end}}
    </div>
    {{end}}
    
    {{if .Recommendations}}
    <div class="recommendations">
        <h2>Recommendations</h2>
        <ol>
        {{range .Recommendations}}
            <li>
                <strong>{{.Title}}</strong> ({{.Priority}} priority)
                <p>{{.Description}}</p>
            </li>
        {{end}}
        </ol>
    </div>
    {{end}}
</body>
</html>`
}

// Helper functions for data extraction

func (r *Reporter) getPhaseLabels(data *ReportData) string {
	labels := make([]string, len(data.PhaseResults))
	for i, phase := range data.PhaseResults {
		labels[i] = string(phase.Phase)
	}
	labelsJSON, _ := json.Marshal(labels)
	return string(labelsJSON)
}

func (r *Reporter) getPerformanceData(data *ReportData) string {
	values := make([]float64, len(data.PhaseResults))
	for i, phase := range data.PhaseResults {
		// Calculate average duration for the phase
		if len(phase.Results) > 0 {
			var total time.Duration
			for _, result := range phase.Results {
				total += result.Duration
			}
			values[i] = float64(total/time.Duration(len(phase.Results))) / float64(time.Millisecond)
		}
	}
	valuesJSON, _ := json.Marshal(values)
	return string(valuesJSON)
}

func (r *Reporter) getMemoryData(data *ReportData) string {
	values := make([]float64, len(data.PhaseResults))
	for i, phase := range data.PhaseResults {
		// Calculate average memory for the phase
		if len(phase.Results) > 0 {
			var total int64
			for _, result := range phase.Results {
				total += result.Memory.TotalAlloc
			}
			values[i] = float64(total/int64(len(phase.Results))) / (1024 * 1024) // Convert to MB
		}
	}
	valuesJSON, _ := json.Marshal(values)
	return string(valuesJSON)
}

// formatBytes formats bytes in human readable format
func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// ReportData contains all data for report generation
type ReportData struct {
	GeneratedAt     time.Time            `json:"generated_at"`
	Platform        Platform             `json:"platform"`
	GitCommit       string               `json:"git_commit"`
	Config          Config               `json:"config"`
	Summary         ReportSummary        `json:"summary"`
	PhaseResults    []*PhaseResult       `json:"phase_results"`
	Comparisons     []*PhaseComparison   `json:"comparisons"`
	Validations     []*ValidationReport  `json:"validations"`
	Recommendations []Recommendation     `json:"recommendations"`
}

// ReportSummary contains summary information
type ReportSummary struct {
	TotalScenarios    int      `json:"total_scenarios"`
	TotalOperations   int      `json:"total_operations"`
	OverallScore      float64  `json:"overall_score"`
	PerformanceScore  float64  `json:"performance_score"`
	StabilityScore    float64  `json:"stability_score"`
	KeyFindings       []string `json:"key_findings"`
}

// PhaseResult contains results for a phase
type PhaseResult struct {
	Phase         Phase          `json:"phase"`
	Duration      time.Duration  `json:"duration"`
	Results       []Result       `json:"results"`
	Score         float64        `json:"score"`
	TopPerformers []Performer    `json:"top_performers"`
	Issues        []Issue        `json:"issues"`
}

// Performer represents a top performing operation
type Performer struct {
	Scenario  string        `json:"scenario"`
	Operation string        `json:"operation"`
	Duration  time.Duration `json:"duration"`
	Memory    int64         `json:"memory"`
	Score     float64       `json:"score"`
}

// Issue represents a performance issue
type Issue struct {
	Severity    string `json:"severity"`
	Description string `json:"description"`
	Impact      string `json:"impact"`
}

// Recommendation represents a performance recommendation
type Recommendation struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Impact      string `json:"impact"`
	Effort      string `json:"effort"`
}

// GeneratePhaseReport generates a report for a specific phase
func (r *Reporter) GeneratePhaseReport(phase Phase, results []Result, validation *ValidationReport) (*PhaseResult, error) {
	phaseResult := &PhaseResult{
		Phase:   phase,
		Results: results,
		Issues:  make([]Issue, 0),
	}
	
	// Calculate phase duration
	if len(results) > 0 {
		earliest := results[0].Timestamp
		latest := results[0].Timestamp
		for _, result := range results {
			if result.Timestamp.Before(earliest) {
				earliest = result.Timestamp
			}
			if result.Timestamp.After(latest) {
				latest = result.Timestamp
			}
		}
		phaseResult.Duration = latest.Sub(earliest)
	}
	
	// Calculate score
	if validation != nil {
		phaseResult.Score = validation.Summary.OverallScore
	}
	
	// Find top performers
	phaseResult.TopPerformers = r.findTopPerformers(results, 5)
	
	// Extract issues from validation
	if validation != nil {
		phaseResult.Issues = r.extractIssues(validation)
	}
	
	return phaseResult, nil
}

// findTopPerformers finds the top performing operations
func (r *Reporter) findTopPerformers(results []Result, limit int) []Performer {
	// Score each result (lower duration and memory = higher score)
	type scoredResult struct {
		result Result
		score  float64
	}
	
	scored := make([]scoredResult, 0)
	for _, result := range results {
		// Simple scoring: inverse of duration + memory
		score := 1000000.0 / (float64(result.Duration) + float64(result.Memory.TotalAlloc)/1000)
		scored = append(scored, scoredResult{result: result, score: score})
	}
	
	// Sort by score
	sort.Slice(scored, func(i, j int) bool {
		return scored[i].score > scored[j].score
	})
	
	// Convert to performers
	performers := make([]Performer, 0)
	for i, sr := range scored {
		if i >= limit {
			break
		}
		performers = append(performers, Performer{
			Scenario:  sr.result.Scenario,
			Operation: sr.result.Operation,
			Duration:  sr.result.Duration,
			Memory:    sr.result.Memory.TotalAlloc,
			Score:     sr.score,
		})
	}
	
	return performers
}

// extractIssues extracts issues from validation report
func (r *Reporter) extractIssues(validation *ValidationReport) []Issue {
	issues := make([]Issue, 0)
	
	for _, result := range validation.Results {
		// Add violations as issues
		for _, violation := range result.Violations {
			issues = append(issues, Issue{
				Severity:    violation.Severity,
				Description: fmt.Sprintf("%s.%s: %s", result.Scenario, result.Operation, violation.Description),
				Impact:      "performance",
			})
		}
		
		// Add regressions as issues
		for _, regression := range result.Regressions {
			issues = append(issues, Issue{
				Severity:    regression.Severity,
				Description: fmt.Sprintf("%s.%s: %s", result.Scenario, result.Operation, regression.Description),
				Impact:      "regression",
			})
		}
	}
	
	return issues
}

// GenerateRecommendations generates performance recommendations
func (r *Reporter) GenerateRecommendations(data *ReportData) []Recommendation {
	recommendations := make([]Recommendation, 0)
	
	// Analyze overall performance
	if data.Summary.PerformanceScore < 70 {
		recommendations = append(recommendations, Recommendation{
			Title:       "Improve Overall Performance",
			Description: "Performance score is below 70. Consider optimizing the slowest operations.",
			Priority:    "high",
			Impact:      "high",
			Effort:      "medium",
		})
	}
	
	// Analyze stability
	if data.Summary.StabilityScore < 80 {
		recommendations = append(recommendations, Recommendation{
			Title:       "Improve Stability",
			Description: "Stability score indicates high variance or GC issues. Review memory management.",
			Priority:    "medium",
			Impact:      "medium",
			Effort:      "medium",
		})
	}
	
	// Check for regressions
	for _, comparison := range data.Comparisons {
		if comparison.Summary.Regressions > 0 {
			recommendations = append(recommendations, Recommendation{
				Title:       fmt.Sprintf("Address Regressions in %s vs %s", comparison.Phase1, comparison.Phase2),
				Description: fmt.Sprintf("Found %d regressions between phases", comparison.Summary.Regressions),
				Priority:    "high",
				Impact:      "high",
				Effort:      "low",
			})
		}
	}
	
	// Check for validation issues
	for _, validation := range data.Validations {
		if validation.Summary.Violations > 0 {
			recommendations = append(recommendations, Recommendation{
				Title:       fmt.Sprintf("Fix Threshold Violations in %s", validation.Phase),
				Description: fmt.Sprintf("Found %d threshold violations", validation.Summary.Violations),
				Priority:    "medium",
				Impact:      "medium",
				Effort:      "low",
			})
		}
	}
	
	return recommendations
}

// GenerateSummaryReport generates a comprehensive summary
func (r *Reporter) GenerateSummaryReport(phaseResults []*PhaseResult, comparisons []*PhaseComparison, validations []*ValidationReport) *ReportSummary {
	summary := &ReportSummary{
		KeyFindings: make([]string, 0),
	}
	
	// Count scenarios and operations
	scenarios := make(map[string]bool)
	operations := make(map[string]bool)
	
	for _, phase := range phaseResults {
		for _, result := range phase.Results {
			scenarios[result.Scenario] = true
			operations[result.Operation] = true
		}
	}
	
	summary.TotalScenarios = len(scenarios)
	summary.TotalOperations = len(operations)
	
	// Calculate average scores
	if len(validations) > 0 {
		var totalOverall, totalPerformance, totalStability float64
		for _, validation := range validations {
			totalOverall += validation.Summary.OverallScore
			totalPerformance += validation.Summary.PerformanceScore
			totalStability += validation.Summary.StabilityScore
		}
		
		summary.OverallScore = totalOverall / float64(len(validations))
		summary.PerformanceScore = totalPerformance / float64(len(validations))
		summary.StabilityScore = totalStability / float64(len(validations))
	}
	
	// Generate key findings
	summary.KeyFindings = r.generateKeyFindings(phaseResults, comparisons, validations)
	
	return summary
}

// generateKeyFindings generates key findings from the data
func (r *Reporter) generateKeyFindings(phaseResults []*PhaseResult, comparisons []*PhaseComparison, validations []*ValidationReport) []string {
	findings := make([]string, 0)
	
	// Performance findings
	if len(phaseResults) > 0 {
		bestPhase := phaseResults[0]
		worstPhase := phaseResults[0]
		
		for _, phase := range phaseResults {
			if phase.Score > bestPhase.Score {
				bestPhase = phase
			}
			if phase.Score < worstPhase.Score {
				worstPhase = phase
			}
		}
		
		findings = append(findings, fmt.Sprintf("Best performing phase: %s (%.1f/100)", bestPhase.Phase, bestPhase.Score))
		findings = append(findings, fmt.Sprintf("Lowest performing phase: %s (%.1f/100)", worstPhase.Phase, worstPhase.Score))
	}
	
	// Comparison findings
	totalImprovements := 0
	totalRegressions := 0
	
	for _, comparison := range comparisons {
		totalImprovements += comparison.Summary.Improvements
		totalRegressions += comparison.Summary.Regressions
	}
	
	if totalImprovements > 0 {
		findings = append(findings, fmt.Sprintf("Total improvements across phases: %d", totalImprovements))
	}
	if totalRegressions > 0 {
		findings = append(findings, fmt.Sprintf("Total regressions across phases: %d", totalRegressions))
	}
	
	// Stability findings
	for _, validation := range validations {
		if validation.Summary.StabilityScore < 60 {
			findings = append(findings, fmt.Sprintf("Phase %s shows stability concerns (score: %.1f)", validation.Phase, validation.Summary.StabilityScore))
		}
	}
	
	return findings
}