// detect-cycles scans generated Go packages and detects circular import dependencies.
// It outputs detailed information about cycles including which types are causing them.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// CycleInfo contains information about a detected import cycle
type CycleInfo struct {
	Cycle     []string      `json:"cycle"`      // List of packages in the cycle
	Edges     []Edge        `json:"edges"`      // Detailed edge information
	BreakSuggestions []BreakSuggestion `json:"suggestions,omitempty"`
}

// Edge represents an import relationship between two packages
type Edge struct {
	From      string   `json:"from"`       // Package doing the importing
	To        string   `json:"to"`         // Package being imported
	Files     []string `json:"files"`      // Files containing the import
	Types     []TypeUsage `json:"types,omitempty"` // Types being used from imported package
}

// TypeUsage tracks how a type from another package is used
type TypeUsage struct {
	TypeName string `json:"type_name"` // e.g., "BarcodeDescriptor"
	Package  string `json:"package"`   // e.g., "coreimage"
	File     string `json:"file"`      // File where it's used
	Line     int    `json:"line"`      // Line number
	Context  string `json:"context"`   // e.g., "property", "parameter", "return"
}

// BreakSuggestion suggests where to break a cycle
type BreakSuggestion struct {
	Edge     Edge   `json:"edge"`
	Reason   string `json:"reason"`
	Action   string `json:"action"`
	Priority int    `json:"priority"` // Higher = more important
}

var (
	generatedDir = flag.String("dir", "../../generated", "Directory containing generated packages")
	outputJSON   = flag.String("output", "", "Output file for JSON report (default: stdout)")
	verbose      = flag.Bool("v", false, "Verbose output")
)

func main() {
	flag.Parse()

	if *verbose {
		fmt.Fprintf(os.Stderr, "Scanning directory: %s\n", *generatedDir)
	}

	// Build dependency graph
	graph, err := buildDependencyGraph(*generatedDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building dependency graph: %v\n", err)
		os.Exit(1)
	}

	if *verbose {
		fmt.Fprintf(os.Stderr, "Found %d packages\n", len(graph))
	}

	// Detect cycles
	cycles := detectCycles(graph)

	if *verbose {
		fmt.Fprintf(os.Stderr, "Found %d cycles\n", len(cycles))
	}

	// Output results
	output := map[string]interface{}{
		"cycles_count": len(cycles),
		"cycles":       cycles,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	if *outputJSON != "" {
		f, err := os.Create(*outputJSON)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		encoder = json.NewEncoder(f)
		encoder.SetIndent("", "  ")
	}

	if err := encoder.Encode(output); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}

	// Exit with non-zero if cycles found
	if len(cycles) > 0 {
		os.Exit(1)
	}
}

// DependencyGraph maps package names to their dependencies
type DependencyGraph map[string]*PackageNode

type PackageNode struct {
	Name    string
	Imports map[string]*Edge // Map of imported package -> edge info
}

func buildDependencyGraph(dir string) (DependencyGraph, error) {
	graph := make(DependencyGraph)
	baseModule := "github.com/tmc/appledocs/generated"

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip non-go files
		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		// Parse the file
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if err != nil {
			return fmt.Errorf("parsing %s: %w", path, err)
		}

		// Determine package name
		pkgName := file.Name.Name
		if pkgName == "" {
			return nil
		}

		// Ensure package node exists
		if _, ok := graph[pkgName]; !ok {
			graph[pkgName] = &PackageNode{
				Name:    pkgName,
				Imports: make(map[string]*Edge),
			}
		}

		// Track imports
		for _, imp := range file.Imports {
			importPath := strings.Trim(imp.Path.Value, `"`)

			// Only track imports within our generated code
			if !strings.HasPrefix(importPath, baseModule) {
				continue
			}

			// Extract package name from import path
			importedPkg := filepath.Base(importPath)

			// Skip self-imports
			if importedPkg == pkgName {
				continue
			}

			// Create or update edge
			edge := graph[pkgName].Imports[importedPkg]
			if edge == nil {
				edge = &Edge{
					From:  pkgName,
					To:    importedPkg,
					Files: []string{},
				}
				graph[pkgName].Imports[importedPkg] = edge
			}

			// Add file to edge
			relFile, _ := filepath.Rel(dir, path)
			edge.Files = append(edge.Files, relFile)
		}

		return nil
	})

	return graph, err
}

func detectCycles(graph DependencyGraph) []CycleInfo {
	var cycles []CycleInfo
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	var dfs func(pkg string, path []string)
	dfs = func(pkg string, path []string) {
		visited[pkg] = true
		recStack[pkg] = true
		path = append(path, pkg)

		if node, ok := graph[pkg]; ok {
			for imported := range node.Imports {
				if !visited[imported] {
					dfs(imported, path)
				} else if recStack[imported] {
					// Found a cycle
					cycleStart := -1
					for i, p := range path {
						if p == imported {
							cycleStart = i
							break
						}
					}
					if cycleStart >= 0 {
						cyclePath := append([]string{}, path[cycleStart:]...)
						cyclePath = append(cyclePath, imported) // Close the cycle

						// Build edges for this cycle
						var edges []Edge
						for i := 0; i < len(cyclePath)-1; i++ {
							from := cyclePath[i]
							to := cyclePath[i+1]
							if node, ok := graph[from]; ok {
								if edge, ok := node.Imports[to]; ok {
									edges = append(edges, *edge)
								}
							}
						}

						cycles = append(cycles, CycleInfo{
							Cycle: cyclePath,
							Edges: edges,
						})
					}
				}
			}
		}

		recStack[pkg] = false
	}

	for pkg := range graph {
		if !visited[pkg] {
			dfs(pkg, []string{})
		}
	}

	// Deduplicate cycles
	return deduplicateCycles(cycles)
}

func deduplicateCycles(cycles []CycleInfo) []CycleInfo {
	seen := make(map[string]bool)
	var unique []CycleInfo

	for _, cycle := range cycles {
		// Create a normalized key for the cycle
		key := normalizeCycle(cycle.Cycle)
		if !seen[key] {
			seen[key] = true
			unique = append(unique, cycle)
		}
	}

	return unique
}

func normalizeCycle(cycle []string) string {
	if len(cycle) == 0 {
		return ""
	}

	// Find the lexicographically smallest rotation
	minRotation := cycle
	for i := 1; i < len(cycle); i++ {
		rotation := append(cycle[i:], cycle[:i]...)
		if strings.Join(rotation, ",") < strings.Join(minRotation, ",") {
			minRotation = rotation
		}
	}

	return strings.Join(minRotation, ",")
}
