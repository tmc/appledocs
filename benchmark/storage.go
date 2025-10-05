// Package benchmark provides storage functionality for benchmark results
package benchmark

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Storage handles persistent storage of benchmark results
type Storage interface {
	// SaveResults saves benchmark results for a phase
	SaveResults(phase Phase, results []Result) error
	
	// LoadResults loads benchmark results for a phase
	LoadResults(phase Phase) ([]Result, error)
	
	// SaveBaseline saves a baseline
	SaveBaseline(baseline Baseline) error
	
	// LoadBaselines loads all baselines
	LoadBaselines() (map[string]Baseline, error)
	
	// GetHistory gets historical results for trending
	GetHistory(scenario, operation string, limit int) ([]HistoricalResult, error)
	
	// Close closes the storage
	Close() error
}

// SQLiteStorage implements Storage using SQLite
type SQLiteStorage struct {
	db *sql.DB
}

// NewStorage creates a new SQLite storage
func NewStorage(dbPath string) (Storage, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}
	
	storage := &SQLiteStorage{db: db}
	
	if err := storage.createTables(); err != nil {
		return nil, fmt.Errorf("create tables: %w", err)
	}
	
	return storage, nil
}

// createTables creates the database tables
func (s *SQLiteStorage) createTables() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS results (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			phase TEXT NOT NULL,
			scenario TEXT NOT NULL,
			operation TEXT NOT NULL,
			timestamp DATETIME NOT NULL,
			duration_ns INTEGER NOT NULL,
			memory_total INTEGER NOT NULL,
			memory_heap INTEGER NOT NULL,
			memory_allocations INTEGER NOT NULL,
			cpu_goroutines INTEGER NOT NULL,
			metadata TEXT,
			platform TEXT,
			git_commit TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		
		`CREATE TABLE IF NOT EXISTS baselines (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			scenario TEXT NOT NULL,
			operation TEXT NOT NULL,
			duration_ns INTEGER NOT NULL,
			memory_bytes INTEGER NOT NULL,
			allocations INTEGER NOT NULL,
			gc_pause_ns INTEGER NOT NULL,
			throughput REAL NOT NULL,
			statistics TEXT NOT NULL,
			timestamp DATETIME NOT NULL,
			git_commit TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(scenario, operation)
		)`,
		
		`CREATE TABLE IF NOT EXISTS validations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			phase TEXT NOT NULL,
			scenario TEXT NOT NULL,
			operation TEXT NOT NULL,
			violations INTEGER NOT NULL,
			regressions INTEGER NOT NULL,
			improvements INTEGER NOT NULL,
			warnings INTEGER NOT NULL,
			overall_score REAL NOT NULL,
			performance_score REAL NOT NULL,
			stability_score REAL NOT NULL,
			timestamp DATETIME NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		
		`CREATE TABLE IF NOT EXISTS comparisons (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			phase1 TEXT NOT NULL,
			phase2 TEXT NOT NULL,
			scenario TEXT NOT NULL,
			operation TEXT NOT NULL,
			duration_change REAL NOT NULL,
			memory_change REAL NOT NULL,
			allocation_change REAL NOT NULL,
			significant BOOLEAN NOT NULL,
			p_value REAL,
			timestamp DATETIME NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		
		`CREATE INDEX IF NOT EXISTS idx_results_phase ON results(phase)`,
		`CREATE INDEX IF NOT EXISTS idx_results_scenario ON results(scenario)`,
		`CREATE INDEX IF NOT EXISTS idx_results_timestamp ON results(timestamp)`,
		`CREATE INDEX IF NOT EXISTS idx_baselines_scenario ON baselines(scenario, operation)`,
		`CREATE INDEX IF NOT EXISTS idx_validations_phase ON validations(phase)`,
		`CREATE INDEX IF NOT EXISTS idx_comparisons_phases ON comparisons(phase1, phase2)`,
	}
	
	for _, query := range queries {
		if _, err := s.db.Exec(query); err != nil {
			return fmt.Errorf("execute query: %w", err)
		}
	}
	
	return nil
}

// SaveResults saves benchmark results
func (s *SQLiteStorage) SaveResults(phase Phase, results []Result) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()
	
	stmt, err := tx.Prepare(`
		INSERT INTO results (
			phase, scenario, operation, timestamp, duration_ns, 
			memory_total, memory_heap, memory_allocations, cpu_goroutines,
			metadata, platform, git_commit
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()
	
	for _, result := range results {
		metadataJSON, _ := json.Marshal(result.Metadata)
		platformJSON, _ := json.Marshal(result.Platform)
		
		// Get git commit if available
		gitCommit := ""
		if commit, ok := result.Metadata["git_commit"].(string); ok {
			gitCommit = commit
		}
		
		_, err := stmt.Exec(
			string(phase),
			result.Scenario,
			result.Operation,
			result.Timestamp,
			result.Duration.Nanoseconds(),
			result.Memory.TotalAlloc,
			result.Memory.HeapAlloc,
			result.Memory.Allocations,
			result.CPU.Goroutines,
			string(metadataJSON),
			string(platformJSON),
			gitCommit,
		)
		if err != nil {
			return fmt.Errorf("insert result: %w", err)
		}
	}
	
	return tx.Commit()
}

// LoadResults loads benchmark results
func (s *SQLiteStorage) LoadResults(phase Phase) ([]Result, error) {
	rows, err := s.db.Query(`
		SELECT scenario, operation, timestamp, duration_ns, 
			   memory_total, memory_heap, memory_allocations, cpu_goroutines,
			   metadata, platform, git_commit
		FROM results 
		WHERE phase = ? 
		ORDER BY timestamp DESC
	`, string(phase))
	if err != nil {
		return nil, fmt.Errorf("query results: %w", err)
	}
	defer rows.Close()
	
	var results []Result
	for rows.Next() {
		var result Result
		var durationNs int64
		var metadataJSON, platformJSON, gitCommit string
		
		err := rows.Scan(
			&result.Scenario,
			&result.Operation,
			&result.Timestamp,
			&durationNs,
			&result.Memory.TotalAlloc,
			&result.Memory.HeapAlloc,
			&result.Memory.Allocations,
			&result.CPU.Goroutines,
			&metadataJSON,
			&platformJSON,
			&gitCommit,
		)
		if err != nil {
			return nil, fmt.Errorf("scan result: %w", err)
		}
		
		result.Phase = phase
		result.Duration = time.Duration(durationNs)
		
		// Unmarshal metadata
		if metadataJSON != "" {
			json.Unmarshal([]byte(metadataJSON), &result.Metadata)
		}
		
		// Unmarshal platform
		if platformJSON != "" {
			json.Unmarshal([]byte(platformJSON), &result.Platform)
		}
		
		// Add git commit to metadata
		if gitCommit != "" {
			if result.Metadata == nil {
				result.Metadata = make(map[string]interface{})
			}
			result.Metadata["git_commit"] = gitCommit
		}
		
		results = append(results, result)
	}
	
	return results, nil
}

// SaveBaseline saves a baseline
func (s *SQLiteStorage) SaveBaseline(baseline Baseline) error {
	statisticsJSON, _ := json.Marshal(baseline.Statistics)
	
	_, err := s.db.Exec(`
		INSERT OR REPLACE INTO baselines (
			scenario, operation, duration_ns, memory_bytes, allocations,
			gc_pause_ns, throughput, statistics, timestamp, git_commit
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		baseline.Scenario,
		baseline.Operation,
		baseline.Metrics.Duration.Nanoseconds(),
		baseline.Metrics.Memory,
		baseline.Metrics.Allocations,
		baseline.Metrics.GCPause.Nanoseconds(),
		baseline.Metrics.Throughput,
		string(statisticsJSON),
		baseline.Timestamp,
		baseline.GitCommit,
	)
	
	if err != nil {
		return fmt.Errorf("insert baseline: %w", err)
	}
	
	return nil
}

// LoadBaselines loads all baselines
func (s *SQLiteStorage) LoadBaselines() (map[string]Baseline, error) {
	rows, err := s.db.Query(`
		SELECT scenario, operation, duration_ns, memory_bytes, allocations,
			   gc_pause_ns, throughput, statistics, timestamp, git_commit
		FROM baselines
	`)
	if err != nil {
		return nil, fmt.Errorf("query baselines: %w", err)
	}
	defer rows.Close()
	
	baselines := make(map[string]Baseline)
	for rows.Next() {
		var baseline Baseline
		var durationNs, gcPauseNs int64
		var statisticsJSON, gitCommit string
		
		err := rows.Scan(
			&baseline.Scenario,
			&baseline.Operation,
			&durationNs,
			&baseline.Metrics.Memory,
			&baseline.Metrics.Allocations,
			&gcPauseNs,
			&baseline.Metrics.Throughput,
			&statisticsJSON,
			&baseline.Timestamp,
			&gitCommit,
		)
		if err != nil {
			return nil, fmt.Errorf("scan baseline: %w", err)
		}
		
		baseline.Metrics.Duration = time.Duration(durationNs)
		baseline.Metrics.GCPause = time.Duration(gcPauseNs)
		baseline.GitCommit = gitCommit
		
		// Unmarshal statistics
		if statisticsJSON != "" {
			json.Unmarshal([]byte(statisticsJSON), &baseline.Statistics)
		}
		
		key := fmt.Sprintf("%s_%s", baseline.Scenario, baseline.Operation)
		baselines[key] = baseline
	}
	
	return baselines, nil
}

// GetHistory gets historical results for trending
func (s *SQLiteStorage) GetHistory(scenario, operation string, limit int) ([]HistoricalResult, error) {
	rows, err := s.db.Query(`
		SELECT phase, timestamp, duration_ns, memory_total, memory_allocations, git_commit
		FROM results 
		WHERE scenario = ? AND operation = ?
		ORDER BY timestamp DESC
		LIMIT ?
	`, scenario, operation, limit)
	if err != nil {
		return nil, fmt.Errorf("query history: %w", err)
	}
	defer rows.Close()
	
	var history []HistoricalResult
	for rows.Next() {
		var result HistoricalResult
		var phaseStr string
		var durationNs int64
		var gitCommit sql.NullString
		
		err := rows.Scan(
			&phaseStr,
			&result.Timestamp,
			&durationNs,
			&result.Memory,
			&result.Allocations,
			&gitCommit,
		)
		if err != nil {
			return nil, fmt.Errorf("scan history: %w", err)
		}
		
		result.Phase = Phase(phaseStr)
		result.Duration = time.Duration(durationNs)
		result.Scenario = scenario
		result.Operation = operation
		if gitCommit.Valid {
			result.GitCommit = gitCommit.String
		}
		
		history = append(history, result)
	}
	
	return history, nil
}

// SaveValidation saves validation results
func (s *SQLiteStorage) SaveValidation(report *ValidationReport) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()
	
	stmt, err := tx.Prepare(`
		INSERT INTO validations (
			phase, scenario, operation, violations, regressions, improvements,
			warnings, overall_score, performance_score, stability_score, timestamp
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()
	
	for _, result := range report.Results {
		_, err := stmt.Exec(
			string(report.Phase),
			result.Scenario,
			result.Operation,
			len(result.Violations),
			len(result.Regressions),
			len(result.Improvements),
			len(result.Warnings),
			report.Summary.OverallScore,
			report.Summary.PerformanceScore,
			report.Summary.StabilityScore,
			report.Timestamp,
		)
		if err != nil {
			return fmt.Errorf("insert validation: %w", err)
		}
	}
	
	return tx.Commit()
}

// SaveComparison saves comparison results
func (s *SQLiteStorage) SaveComparison(comparison *PhaseComparison) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()
	
	stmt, err := tx.Prepare(`
		INSERT INTO comparisons (
			phase1, phase2, scenario, operation, duration_change, memory_change,
			allocation_change, significant, p_value, timestamp
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()
	
	for _, comp := range comparison.Comparisons {
		_, err := stmt.Exec(
			string(comparison.Phase1),
			string(comparison.Phase2),
			comp.Scenario,
			comp.Operation,
			comp.Changes.Duration.Relative,
			comp.Changes.Memory.Relative,
			comp.Changes.Allocations.Relative,
			comp.Significance.Significant,
			comp.Significance.PValue,
			comparison.Timestamp,
		)
		if err != nil {
			return fmt.Errorf("insert comparison: %w", err)
		}
	}
	
	return tx.Commit()
}

// GetTrendData gets trend data for visualization
func (s *SQLiteStorage) GetTrendData(scenario, operation string, days int) ([]TrendPoint, error) {
	rows, err := s.db.Query(`
		SELECT phase, timestamp, duration_ns, memory_total, memory_allocations
		FROM results 
		WHERE scenario = ? AND operation = ?
		  AND timestamp >= datetime('now', '-' || ? || ' days')
		ORDER BY timestamp ASC
	`, scenario, operation, days)
	if err != nil {
		return nil, fmt.Errorf("query trend data: %w", err)
	}
	defer rows.Close()
	
	var trends []TrendPoint
	for rows.Next() {
		var point TrendPoint
		var phaseStr string
		var durationNs int64
		
		err := rows.Scan(
			&phaseStr,
			&point.Timestamp,
			&durationNs,
			&point.Memory,
			&point.Allocations,
		)
		if err != nil {
			return nil, fmt.Errorf("scan trend point: %w", err)
		}
		
		point.Phase = Phase(phaseStr)
		point.Duration = time.Duration(durationNs)
		
		trends = append(trends, point)
	}
	
	return trends, nil
}

// GetPerformanceReport gets a performance report
func (s *SQLiteStorage) GetPerformanceReport(phase Phase) (*PerformanceReport, error) {
	// Get summary statistics
	var report PerformanceReport
	report.Phase = phase
	report.GeneratedAt = time.Now()
	
	// Get total results
	err := s.db.QueryRow(`
		SELECT COUNT(*) FROM results WHERE phase = ?
	`, string(phase)).Scan(&report.TotalResults)
	if err != nil {
		return nil, fmt.Errorf("query total results: %w", err)
	}
	
	// Get average duration
	var avgDurationNs sql.NullInt64
	err = s.db.QueryRow(`
		SELECT AVG(duration_ns) FROM results WHERE phase = ?
	`, string(phase)).Scan(&avgDurationNs)
	if err != nil {
		return nil, fmt.Errorf("query average duration: %w", err)
	}
	if avgDurationNs.Valid {
		report.AvgDuration = time.Duration(avgDurationNs.Int64)
	}
	
	// Get average memory
	var avgMemory sql.NullInt64
	err = s.db.QueryRow(`
		SELECT AVG(memory_total) FROM results WHERE phase = ?
	`, string(phase)).Scan(&avgMemory)
	if err != nil {
		return nil, fmt.Errorf("query average memory: %w", err)
	}
	if avgMemory.Valid {
		report.AvgMemory = avgMemory.Int64
	}
	
	// Get top scenarios by duration
	rows, err := s.db.Query(`
		SELECT scenario, operation, AVG(duration_ns) as avg_duration
		FROM results 
		WHERE phase = ?
		GROUP BY scenario, operation
		ORDER BY avg_duration DESC
		LIMIT 10
	`, string(phase))
	if err != nil {
		return nil, fmt.Errorf("query top scenarios: %w", err)
	}
	defer rows.Close()
	
	report.TopScenarios = make([]ScenarioSummary, 0)
	for rows.Next() {
		var summary ScenarioSummary
		var avgDurationNs int64
		
		err := rows.Scan(&summary.Scenario, &summary.Operation, &avgDurationNs)
		if err != nil {
			return nil, fmt.Errorf("scan scenario summary: %w", err)
		}
		
		summary.AvgDuration = time.Duration(avgDurationNs)
		report.TopScenarios = append(report.TopScenarios, summary)
	}
	
	return &report, nil
}

// Close closes the database connection
func (s *SQLiteStorage) Close() error {
	return s.db.Close()
}

// Helper function to load baselines from storage
func (bf *BenchmarkFramework) loadBaselines() error {
	baselines, err := bf.storage.LoadBaselines()
	if err != nil {
		return fmt.Errorf("load baselines: %w", err)
	}
	
	bf.baselines = baselines
	return nil
}

// Helper structures for storage

// HistoricalResult represents a historical benchmark result
type HistoricalResult struct {
	Phase       Phase     `json:"phase"`
	Scenario    string    `json:"scenario"`
	Operation   string    `json:"operation"`
	Timestamp   time.Time `json:"timestamp"`
	Duration    time.Duration `json:"duration"`
	Memory      int64     `json:"memory"`
	Allocations int64     `json:"allocations"`
	GitCommit   string    `json:"git_commit"`
}

// TrendPoint represents a point in a trend analysis
type TrendPoint struct {
	Phase       Phase     `json:"phase"`
	Timestamp   time.Time `json:"timestamp"`
	Duration    time.Duration `json:"duration"`
	Memory      int64     `json:"memory"`
	Allocations int64     `json:"allocations"`
}

// PerformanceReport represents a performance report
type PerformanceReport struct {
	Phase        Phase             `json:"phase"`
	GeneratedAt  time.Time         `json:"generated_at"`
	TotalResults int               `json:"total_results"`
	AvgDuration  time.Duration     `json:"avg_duration"`
	AvgMemory    int64             `json:"avg_memory"`
	TopScenarios []ScenarioSummary `json:"top_scenarios"`
}

// ScenarioSummary represents a scenario summary
type ScenarioSummary struct {
	Scenario    string        `json:"scenario"`
	Operation   string        `json:"operation"`
	AvgDuration time.Duration `json:"avg_duration"`
	AvgMemory   int64         `json:"avg_memory"`
}

// BatchSaveResults saves multiple results efficiently
func (s *SQLiteStorage) BatchSaveResults(phaseResults map[Phase][]Result) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()
	
	stmt, err := tx.Prepare(`
		INSERT INTO results (
			phase, scenario, operation, timestamp, duration_ns, 
			memory_total, memory_heap, memory_allocations, cpu_goroutines,
			metadata, platform, git_commit
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()
	
	for phase, results := range phaseResults {
		for _, result := range results {
			metadataJSON, _ := json.Marshal(result.Metadata)
			platformJSON, _ := json.Marshal(result.Platform)
			
			// Get git commit if available
			gitCommit := ""
			if commit, ok := result.Metadata["git_commit"].(string); ok {
				gitCommit = commit
			}
			
			_, err := stmt.Exec(
				string(phase),
				result.Scenario,
				result.Operation,
				result.Timestamp,
				result.Duration.Nanoseconds(),
				result.Memory.TotalAlloc,
				result.Memory.HeapAlloc,
				result.Memory.Allocations,
				result.CPU.Goroutines,
				string(metadataJSON),
				string(platformJSON),
				gitCommit,
			)
			if err != nil {
				return fmt.Errorf("insert result: %w", err)
			}
		}
	}
	
	return tx.Commit()
}

// GetPhaseStats gets statistics for a phase
func (s *SQLiteStorage) GetPhaseStats(phase Phase) (*PhaseStats, error) {
	var stats PhaseStats
	stats.Phase = phase
	
	// Get counts
	err := s.db.QueryRow(`
		SELECT COUNT(*), COUNT(DISTINCT scenario), COUNT(DISTINCT operation)
		FROM results WHERE phase = ?
	`, string(phase)).Scan(&stats.TotalResults, &stats.UniqueScenarios, &stats.UniqueOperations)
	if err != nil {
		return nil, fmt.Errorf("query phase stats: %w", err)
	}
	
	// Get duration stats
	err = s.db.QueryRow(`
		SELECT MIN(duration_ns), MAX(duration_ns), AVG(duration_ns)
		FROM results WHERE phase = ?
	`, string(phase)).Scan(&stats.MinDuration, &stats.MaxDuration, &stats.AvgDuration)
	if err != nil {
		return nil, fmt.Errorf("query duration stats: %w", err)
	}
	
	// Get memory stats
	err = s.db.QueryRow(`
		SELECT MIN(memory_total), MAX(memory_total), AVG(memory_total)
		FROM results WHERE phase = ?
	`, string(phase)).Scan(&stats.MinMemory, &stats.MaxMemory, &stats.AvgMemory)
	if err != nil {
		return nil, fmt.Errorf("query memory stats: %w", err)
	}
	
	return &stats, nil
}

// PhaseStats represents statistics for a phase
type PhaseStats struct {
	Phase            Phase `json:"phase"`
	TotalResults     int   `json:"total_results"`
	UniqueScenarios  int   `json:"unique_scenarios"`
	UniqueOperations int   `json:"unique_operations"`
	MinDuration      int64 `json:"min_duration_ns"`
	MaxDuration      int64 `json:"max_duration_ns"`
	AvgDuration      int64 `json:"avg_duration_ns"`
	MinMemory        int64 `json:"min_memory_bytes"`
	MaxMemory        int64 `json:"max_memory_bytes"`
	AvgMemory        int64 `json:"avg_memory_bytes"`
}