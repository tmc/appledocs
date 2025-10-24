package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// Debug categories - use these constants throughout the codebase
const (
	DebugTypeMap      = "typemap"
	DebugUndefined    = "undefined"
	DebugHierarchy    = "hierarchy"
	DebugEnumDedup    = "enum-dedup"
	DebugImports      = "imports"
	DebugTimeInterval = "timeinterval"
	DebugObject       = "object"
	DebugEnumType     = "enum-type"
	DebugEnumCreate   = "enum-create"
	DebugEnumCases    = "enum-cases"
	DebugEnumAttach   = "enum-attach"
	DebugParser       = "parser"
	DebugTemplates    = "templates"
)

// AllCategories lists all available debug categories
var AllCategories = []string{
	DebugTypeMap, DebugUndefined, DebugHierarchy, DebugEnumDedup,
	DebugImports, DebugTimeInterval, DebugObject, DebugEnumType,
	DebugEnumCreate, DebugEnumCases, DebugEnumAttach, DebugParser,
	DebugTemplates,
}

// CategoryDescriptions provides help text for each category
var CategoryDescriptions = map[string]string{
	DebugTypeMap:      "Type mapping resolution (ObjC → Go type conversion)",
	DebugUndefined:    "Undefined type handling and resolution",
	DebugHierarchy:    "Class hierarchy violation detection",
	DebugEnumDedup:    "Enum deduplication and merging",
	DebugImports:      "Import path resolution",
	DebugTimeInterval: "TimeInterval typedef handling",
	DebugObject:       "Object/IObject type conversions",
	DebugEnumType:     "Enum type resolution",
	DebugEnumCreate:   "Enum creation from documentation",
	DebugEnumCases:    "Enum case parsing",
	DebugEnumAttach:   "Enum attachment to classes",
	DebugParser:       "occ2go parser operations",
	DebugTemplates:    "Template execution debug comments in generated code",
}

// DebugLogger provides structured debug logging with category-based filtering
// and regex pattern matching for focused debugging.
type DebugLogger struct {
	enabled map[string]bool
	filters map[string]*regexp.Regexp
	logger  *slog.Logger
}

// Debug is the global debug logger instance
var Debug *DebugLogger

// Logger is the global logger for general output (info, warnings, errors)
var Logger *slog.Logger

// VerboseLogger is used for verbose output (enabled with -v flag)
var VerboseLogger *slog.Logger

// SimpleHandler is a custom slog handler that outputs logs in a simplified format:
// "LEVEL: message key=value key2=value2"
type SimpleHandler struct {
	w     *os.File
	level slog.Level
}

// Enabled reports whether the handler handles records at the given level.
func (h *SimpleHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

// Handle formats and writes a log record.
func (h *SimpleHandler) Handle(_ context.Context, r slog.Record) error {
	// Format: "LEVEL: message key=value key2=value2"
	buf := make([]byte, 0, 256)

	// Add level
	buf = append(buf, r.Level.String()...)
	buf = append(buf, ": "...)

	// Add message
	buf = append(buf, r.Message...)

	// Add attributes
	r.Attrs(func(a slog.Attr) bool {
		buf = append(buf, ' ')
		buf = append(buf, a.Key...)
		buf = append(buf, '=')
		buf = append(buf, a.Value.String()...)
		return true
	})

	buf = append(buf, '\n')
	_, err := h.w.Write(buf)
	return err
}

// WithAttrs returns a new handler with the given attributes.
func (h *SimpleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// For simplicity, we don't support persistent attributes
	return h
}

// WithGroup returns a new handler with the given group.
func (h *SimpleHandler) WithGroup(name string) slog.Handler {
	// For simplicity, we don't support groups
	return h
}

// InitDebug initializes the global debug logger with the specified categories and filter.
// Categories can be comma-separated or "all" to enable all categories.
// FilterSpec can be:
//   - A single regex pattern (applies to all enabled categories)
//   - Per-category patterns: "category:pattern,category2:pattern2"
//
// Examples:
//
//	InitDebug("typemap", "Coder|Error")                    // typemap with filter
//	InitDebug("typemap,hierarchy", "typemap:Coder,hierarchy:Broadcast")  // per-category
//	InitDebug("all", "")                                   // all categories, no filter
func InitDebug(categories, filterSpec string) {
	enabled := make(map[string]bool)
	filters := make(map[string]*regexp.Regexp)

	// Parse comma-separated categories
	if categories != "" {
		for _, cat := range strings.Split(categories, ",") {
			cat = strings.TrimSpace(cat)
			if cat == "all" {
				for _, c := range AllCategories {
					enabled[c] = true
				}
			} else {
				enabled[cat] = true
			}
		}
	}

	// Parse filter specification
	if filterSpec != "" {
		if !strings.Contains(filterSpec, ":") {
			// Global filter - applies to all enabled categories
			re, err := regexp.Compile(filterSpec)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Warning: Invalid debug filter pattern %q: %v\n", filterSpec, err)
			} else {
				for cat := range enabled {
					filters[cat] = re
				}
			}
		} else {
			// Per-category filters: "typemap:Coder,hierarchy:Broadcast"
			for _, spec := range strings.Split(filterSpec, ",") {
				parts := strings.SplitN(spec, ":", 2)
				if len(parts) == 2 {
					cat, pattern := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
					re, err := regexp.Compile(pattern)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Warning: Invalid regex for category %s: %q: %v\n",
							cat, pattern, err)
						continue
					}
					filters[cat] = re
				}
			}
		}
	}

	// Backwards compatibility: Check DEBUG_* environment variables
	for _, cat := range AllCategories {
		envKey := "DEBUG_" + strings.ToUpper(strings.ReplaceAll(cat, "-", "_"))
		if os.Getenv(envKey) == "1" {
			enabled[cat] = true
		}
	}

	// Check for global DEBUG_FILTER environment variable
	if envFilter := os.Getenv("DEBUG_FILTER"); envFilter != "" && filterSpec == "" {
		re, err := regexp.Compile(envFilter)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Invalid DEBUG_FILTER pattern %q: %v\n", envFilter, err)
		} else {
			for cat := range enabled {
				filters[cat] = re
			}
		}
	}

	// Create simplified handler for debug logs
	handler := &SimpleHandler{
		w:     os.Stderr,
		level: slog.LevelDebug,
	}

	Debug = &DebugLogger{
		enabled: enabled,
		filters: filters,
		logger:  slog.New(handler),
	}

	// Initialize general-purpose logger with simplified format
	infoHandler := &SimpleHandler{
		w:     os.Stderr,
		level: slog.LevelInfo,
	}
	Logger = slog.New(infoHandler)

	// Initialize verbose logger (starts disabled, call SetVerbose(true) to enable)
	SetVerbose(false)
}

// SetVerbose enables or disables verbose logging
func SetVerbose(enabled bool) {
	var level slog.Level
	if enabled {
		level = slog.LevelDebug
	} else {
		level = slog.LevelWarn // Only show warnings and errors when not verbose
	}
	verboseHandler := &SimpleHandler{
		w:     os.Stderr,
		level: level,
	}
	VerboseLogger = slog.New(verboseHandler)
}

// Enabled returns true if the given category is enabled for debugging
func (d *DebugLogger) Enabled(category string) bool {
	if d == nil {
		return false
	}
	return d.enabled[category]
}

// Match returns true if the category is enabled AND at least one context string
// matches the category's regex filter (if any filter is set).
// If no filter is set for the category, returns true if category is enabled.
func (d *DebugLogger) Match(category string, contexts ...string) bool {
	if !d.Enabled(category) {
		return false
	}

	filter, hasFilter := d.filters[category]
	if !hasFilter {
		return true // No filter = match everything in this category
	}

	// Check if any context string matches the filter
	for _, ctx := range contexts {
		if filter.MatchString(ctx) {
			return true
		}
	}
	return false
}

// Log outputs a debug message if the category is enabled and contexts match the filter.
// Uses structured logging with slog, automatically including the category in output.
func (d *DebugLogger) Log(category, msg string, contexts []string, args ...any) {
	if d.Match(category, contexts...) {
		allArgs := append([]any{"category", category}, args...)
		d.logger.Debug(msg, allArgs...)
	}
}

// Convenience methods for each category with automatic context extraction

// TypeMap logs type mapping operations, filtering by objcType or goType
func (d *DebugLogger) TypeMap(msg string, objcType, goType string, args ...any) {
	d.Log(DebugTypeMap, msg, []string{objcType, goType}, args...)
}

// Undefined logs undefined type handling, filtering by type name
func (d *DebugLogger) Undefined(msg string, typeName string, args ...any) {
	d.Log(DebugUndefined, msg, []string{typeName}, args...)
}

// Hierarchy logs class hierarchy operations, filtering by class or method name
func (d *DebugLogger) Hierarchy(msg string, className, methodName string, args ...any) {
	d.Log(DebugHierarchy, msg, []string{className, methodName}, args...)
}

// EnumDedup logs enum deduplication operations, filtering by enum name
func (d *DebugLogger) EnumDedup(msg string, enumName string, args ...any) {
	d.Log(DebugEnumDedup, msg, []string{enumName}, args...)
}

// Imports logs import resolution, filtering by import path or class name
func (d *DebugLogger) Imports(msg string, importPath, className string, args ...any) {
	d.Log(DebugImports, msg, []string{importPath, className}, args...)
}

// TimeInterval logs TimeInterval typedef handling, filtering by type name
func (d *DebugLogger) TimeInterval(msg string, typeName string, args ...any) {
	d.Log(DebugTimeInterval, msg, []string{typeName}, args...)
}

// Object logs Object/IObject conversions, filtering by type name
func (d *DebugLogger) Object(msg string, typeName string, args ...any) {
	d.Log(DebugObject, msg, []string{typeName}, args...)
}

// EnumType logs enum type resolution, filtering by enum name
func (d *DebugLogger) EnumType(msg string, enumName string, args ...any) {
	d.Log(DebugEnumType, msg, []string{enumName}, args...)
}

// EnumCreate logs enum creation, filtering by enum name
func (d *DebugLogger) EnumCreate(msg string, enumName string, args ...any) {
	d.Log(DebugEnumCreate, msg, []string{enumName}, args...)
}

// EnumCases logs enum case parsing, filtering by enum name
func (d *DebugLogger) EnumCases(msg string, enumName string, args ...any) {
	d.Log(DebugEnumCases, msg, []string{enumName}, args...)
}

// EnumAttach logs enum attachment to classes, filtering by enum or class name
func (d *DebugLogger) EnumAttach(msg string, enumName, className string, args ...any) {
	d.Log(DebugEnumAttach, msg, []string{enumName, className}, args...)
}

// Parser logs occ2go parser operations, filtering by document path or symbol name
func (d *DebugLogger) Parser(msg string, docPath, symbolName string, args ...any) {
	d.Log(DebugParser, msg, []string{docPath, symbolName}, args...)
}

// PrintDebugHelp prints comprehensive help about available debug categories and usage
func PrintDebugHelp() {
	fmt.Println("Available debug categories:")
	fmt.Println()

	// Calculate max width for alignment
	maxLen := 0
	for _, cat := range AllCategories {
		if len(cat) > maxLen {
			maxLen = len(cat)
		}
	}

	// Print categories with descriptions
	for _, cat := range AllCategories {
		desc := CategoryDescriptions[cat]
		fmt.Printf("  %-*s  %s\n", maxLen, cat, desc)
	}

	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  --debug=CATEGORIES              Enable debug categories (comma-separated)")
	fmt.Println("  --debug-filter=PATTERN          Filter debug output with regex pattern")
	fmt.Println("  --debug-filter=CAT:PATTERN,...  Per-category regex filters")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  # Debug all type mapping")
	fmt.Println("  --debug=typemap")
	fmt.Println()
	fmt.Println("  # Add debug comments in generated code (shows template execution)")
	fmt.Println("  --debug=templates")
	fmt.Println()
	fmt.Println("  # Debug only Coder and Error types")
	fmt.Println("  --debug=typemap --debug-filter='Coder|Error'")
	fmt.Println()
	fmt.Println("  # Debug NSEnergy and NSLength types")
	fmt.Println("  --debug=typemap --debug-filter='NSEnergy.*|NSLength.*'")
	fmt.Println()
	fmt.Println("  # Different filters for different categories")
	fmt.Println("  --debug=typemap,hierarchy --debug-filter='typemap:Coder,hierarchy:Broadcast'")
	fmt.Println()
	fmt.Println("  # Debug everything")
	fmt.Println("  --debug=all")
	fmt.Println()
	fmt.Println("Environment variables:")
	fmt.Println("  DEBUG=typemap                   Enable categories")
	fmt.Println("  DEBUG_FILTER='Coder|Error'      Global filter pattern")
	fmt.Println("  DEBUG_TYPEMAP=1                 Backwards compatible (no filter)")
	fmt.Println("  DEBUG_TEMPLATES=1               Enable template debug comments")
	fmt.Println()
	fmt.Println("Filter patterns are Go regular expressions. Filters are matched against:")
	fmt.Println("  - typemap: objcType and goType")
	fmt.Println("  - undefined: typeName")
	fmt.Println("  - hierarchy: className and methodName")
	fmt.Println("  - imports: importPath and className")
	fmt.Println("  - enum-*: enumName")
	fmt.Println("  - parser: docPath and symbolName")
}

// LogWithFunc logs a message with a custom function name for better traceability in template functions.
func LogWithFunc(logger *slog.Logger, level slog.Level, funcName, msg string, args ...any) {
	if logger == nil {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.AddAttrs(slog.String("func", funcName))
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			if key, ok := args[i].(string); ok {
				r.AddAttrs(slog.Any(key, args[i+1]))
			}
		}
	}
	_ = logger.Handler().Handle(nil, r)
}

// DebugWithFunc logs a debug message with a custom function name for template functions.
func DebugWithFunc(funcName, msg string, args ...any) {
	if VerboseLogger != nil {
		LogWithFunc(VerboseLogger, slog.LevelDebug, funcName, msg, args...)
	}
}

// InfoWithFunc logs an info message with a custom function name for template functions.
func InfoWithFunc(funcName, msg string, args ...any) {
	if Logger != nil {
		LogWithFunc(Logger, slog.LevelInfo, funcName, msg, args...)
	}
}
