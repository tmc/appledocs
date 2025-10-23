package main

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strings"
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
)

// AllCategories lists all available debug categories
var AllCategories = []string{
	DebugTypeMap, DebugUndefined, DebugHierarchy, DebugEnumDedup,
	DebugImports, DebugTimeInterval, DebugObject, DebugEnumType,
	DebugEnumCreate, DebugEnumCases, DebugEnumAttach, DebugParser,
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

// InitDebug initializes the global debug logger with the specified categories and filter.
// Categories can be comma-separated or "all" to enable all categories.
// FilterSpec can be:
//   - A single regex pattern (applies to all enabled categories)
//   - Per-category patterns: "category:pattern,category2:pattern2"
//
// Examples:
//   InitDebug("typemap", "Coder|Error")                    // typemap with filter
//   InitDebug("typemap,hierarchy", "typemap:Coder,hierarchy:Broadcast")  // per-category
//   InitDebug("all", "")                                   // all categories, no filter
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

	// Create slog handler for stderr
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	Debug = &DebugLogger{
		enabled: enabled,
		filters: filters,
		logger:  slog.New(handler),
	}
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
	fmt.Println()
	fmt.Println("Filter patterns are Go regular expressions. Filters are matched against:")
	fmt.Println("  - typemap: objcType and goType")
	fmt.Println("  - undefined: typeName")
	fmt.Println("  - hierarchy: className and methodName")
	fmt.Println("  - imports: importPath and className")
	fmt.Println("  - enum-*: enumName")
	fmt.Println("  - parser: docPath and symbolName")
}
