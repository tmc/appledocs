package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// SchemaAnalyzer analyzes JSON schema patterns
type SchemaAnalyzer struct {
	fieldFrequency    map[string]int            // field path -> occurrence count
	fieldTypes        map[string]map[string]int // field path -> type -> count
	fieldExamples     map[string][]interface{}  // field path -> examples
	arrayElementTypes map[string]map[string]int // array field -> element type -> count
}

// NewSchemaAnalyzer creates a new schema analyzer
func NewSchemaAnalyzer() *SchemaAnalyzer {
	return &SchemaAnalyzer{
		fieldFrequency:    make(map[string]int),
		fieldTypes:        make(map[string]map[string]int),
		fieldExamples:     make(map[string][]interface{}),
		arrayElementTypes: make(map[string]map[string]int),
	}
}

// AnalyzeDirectory analyzes all JSON files in a directory
func (a *SchemaAnalyzer) AnalyzeDirectory(root string, maxFiles int) error {
	count := 0
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}
		if maxFiles > 0 && count >= maxFiles {
			return filepath.SkipDir
		}

		if err := a.analyzeFile(path); err != nil {
			log.Printf("Warning: failed to analyze %s: %v", path, err)
		}

		count++
		if count%100 == 0 {
			log.Printf("Analyzed %d files...", count)
		}
		return nil
	})

	log.Printf("Analyzed %d JSON files", count)
	return err
}

// analyzeFile analyzes a single JSON file
func (a *SchemaAnalyzer) analyzeFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var doc map[string]interface{}
	if err := json.Unmarshal(data, &doc); err != nil {
		return err
	}

	a.analyzeValue("", doc)
	return nil
}

// analyzeValue recursively analyzes a JSON value
func (a *SchemaAnalyzer) analyzeValue(path string, value interface{}) {
	if value == nil {
		return
	}

	// Record field frequency
	if path != "" {
		a.fieldFrequency[path]++

		// Record type
		typeName := getTypeName(value)
		if a.fieldTypes[path] == nil {
			a.fieldTypes[path] = make(map[string]int)
		}
		a.fieldTypes[path][typeName]++

		// Save example (limit to 3 examples per field)
		if len(a.fieldExamples[path]) < 3 {
			a.fieldExamples[path] = append(a.fieldExamples[path], value)
		}
	}

	// Recursively analyze nested structures
	switch v := value.(type) {
	case map[string]interface{}:
		for key, val := range v {
			childPath := path
			if childPath != "" {
				childPath += "."
			}
			childPath += key
			a.analyzeValue(childPath, val)
		}

	case []interface{}:
		if len(v) > 0 {
			// Analyze array element types
			if a.arrayElementTypes[path] == nil {
				a.arrayElementTypes[path] = make(map[string]int)
			}
			for _, elem := range v {
				elemType := getTypeName(elem)
				a.arrayElementTypes[path][elemType]++
			}

			// Analyze first element in detail
			a.analyzeValue(path+"[]", v[0])
		}
	}
}

// getTypeName returns a string representation of the type
func getTypeName(value interface{}) string {
	if value == nil {
		return "null"
	}

	switch v := value.(type) {
	case bool:
		return "bool"
	case float64:
		if v == float64(int64(v)) {
			return "int"
		}
		return "float64"
	case string:
		return "string"
	case []interface{}:
		if len(v) == 0 {
			return "[]unknown"
		}
		return "[]" + getTypeName(v[0])
	case map[string]interface{}:
		return "object"
	default:
		return "unknown"
	}
}

// PrintReport prints an analysis report
func (a *SchemaAnalyzer) PrintReport() {
	fmt.Println("=== Schema Analysis Report ===\n")

	// Get sorted field paths
	var paths []string
	for path := range a.fieldFrequency {
		paths = append(paths, path)
	}
	sort.Slice(paths, func(i, j int) bool {
		// Sort by frequency descending, then alphabetically
		if a.fieldFrequency[paths[i]] != a.fieldFrequency[paths[j]] {
			return a.fieldFrequency[paths[i]] > a.fieldFrequency[paths[j]]
		}
		return paths[i] < paths[j]
	})

	// Print top-level fields
	fmt.Println("## Top-Level Fields (by frequency):")
	for _, path := range paths {
		if !strings.Contains(path, ".") && !strings.Contains(path, "[]") {
			freq := a.fieldFrequency[path]
			types := a.fieldTypes[path]
			fmt.Printf("  %-30s freq=%4d  types=%v\n", path, freq, types)
		}
	}

	// Print common nested structures
	fmt.Println("\n## Common Nested Structures (freq > 50%):")
	threshold := a.fieldFrequency["metadata"] / 2 // 50% of documents
	for _, path := range paths {
		if strings.Contains(path, ".") && !strings.Contains(path, "[]") {
			freq := a.fieldFrequency[path]
			if freq >= threshold {
				types := a.fieldTypes[path]
				fmt.Printf("  %-50s freq=%4d  types=%v\n", path, freq, types)
			}
		}
	}

	// Print array patterns
	fmt.Println("\n## Array Element Types:")
	var arrayPaths []string
	for path := range a.arrayElementTypes {
		arrayPaths = append(arrayPaths, path)
	}
	sort.Strings(arrayPaths)
	for _, path := range arrayPaths {
		if a.fieldFrequency[path] >= threshold {
			elemTypes := a.arrayElementTypes[path]
			fmt.Printf("  %-50s %v\n", path, elemTypes)
		}
	}

	// Field statistics
	fmt.Printf("\n## Statistics:\n")
	fmt.Printf("  Total unique fields: %d\n", len(a.fieldFrequency))
	fmt.Printf("  Fields in >90%% of docs: %d\n", countFieldsAboveThreshold(a.fieldFrequency, a.fieldFrequency["metadata"]*9/10))
	fmt.Printf("  Fields in >50%% of docs: %d\n", countFieldsAboveThreshold(a.fieldFrequency, threshold))
}

func countFieldsAboveThreshold(freq map[string]int, threshold int) int {
	count := 0
	for _, f := range freq {
		if f >= threshold {
			count++
		}
	}
	return count
}

// analyzeSchema is the main entry point
func analyzeSchema(docsPath string, maxFiles int) error {
	analyzer := NewSchemaAnalyzer()

	if err := analyzer.AnalyzeDirectory(docsPath, maxFiles); err != nil {
		return fmt.Errorf("analyze directory: %w", err)
	}

	analyzer.PrintReport()
	return nil
}
