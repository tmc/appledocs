// errors.go provides error handling infrastructure for the generator

package main

import (
	"fmt"
	"strings"
)

// ErrorCollector accumulates errors during generation with context
type ErrorCollector struct {
	errors []ErrorWithContext
}

// ErrorWithContext wraps an error with file and symbol context
type ErrorWithContext struct {
	File   string
	Symbol string
	Stage  string // "parse", "generate", "validate"
	Err    error
}

// Add adds an error with context to the collector
func (ec *ErrorCollector) Add(file, symbol, stage string, err error) {
	if err != nil {
		ec.errors = append(ec.errors, ErrorWithContext{
			File:   file,
			Symbol: symbol,
			Stage:  stage,
			Err:    err,
		})
	}
}

// Addf adds a formatted error with context
func (ec *ErrorCollector) Addf(file, symbol, stage, format string, args ...interface{}) {
	ec.Add(file, symbol, stage, fmt.Errorf(format, args...))
}

// HasErrors returns true if any errors were collected
func (ec *ErrorCollector) HasErrors() bool {
	return len(ec.errors) > 0
}

// Count returns the number of errors collected
func (ec *ErrorCollector) Count() int {
	return len(ec.errors)
}

// Errors returns all collected errors
func (ec *ErrorCollector) Errors() []ErrorWithContext {
	return ec.errors
}

// Summary returns a formatted summary of all errors
func (ec *ErrorCollector) Summary() string {
	if !ec.HasErrors() {
		return "No errors"
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Collected %d errors:\n\n", len(ec.errors)))

	// Group errors by stage
	byStage := make(map[string][]ErrorWithContext)
	for _, e := range ec.errors {
		byStage[e.Stage] = append(byStage[e.Stage], e)
	}

	for stage, errs := range byStage {
		sb.WriteString(fmt.Sprintf("=== %s errors (%d) ===\n", stage, len(errs)))
		for i, e := range errs {
			sb.WriteString(fmt.Sprintf("%d. %s", i+1, e.Error()))
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

// Error implements error interface
func (e ErrorWithContext) Error() string {
	parts := []string{}

	if e.File != "" {
		parts = append(parts, fmt.Sprintf("file=%s", e.File))
	}
	if e.Symbol != "" {
		parts = append(parts, fmt.Sprintf("symbol=%s", e.Symbol))
	}
	if e.Stage != "" {
		parts = append(parts, fmt.Sprintf("stage=%s", e.Stage))
	}

	context := strings.Join(parts, " ")
	if context != "" {
		return fmt.Sprintf("[%s] %v", context, e.Err)
	}
	return e.Err.Error()
}

// ErrorStats provides statistics about collected errors
type ErrorStats struct {
	Total      int
	ByStage    map[string]int
	ByFramework map[string]int
}

// Stats computes statistics from collected errors
func (ec *ErrorCollector) Stats() ErrorStats {
	stats := ErrorStats{
		Total:       len(ec.errors),
		ByStage:     make(map[string]int),
		ByFramework: make(map[string]int),
	}

	for _, e := range ec.errors {
		stats.ByStage[e.Stage]++
		// Extract framework from file path if possible
		// This is a simple heuristic
		if e.File != "" {
			stats.ByFramework[e.File]++
		}
	}

	return stats
}
