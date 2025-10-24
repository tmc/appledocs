// detect-cycles using go list for more accurate dependency information
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GoListPackage represents the JSON output from go list
type GoListPackage struct {
	ImportPath string   `json:"ImportPath"`
	Name       string   `json:"Name"`
	Imports    []string `json:"Imports"`
	Dir        string   `json:"Dir"`
	Error      *struct {
		Err string `json:"Err"`
	} `json:"Error,omitempty"`
}

var (
	dir        = flag.String("dir", "../../generated", "Directory to scan")
	outputJSON = flag.String("output", "", "Output JSON file")
	verbose    = flag.Bool("v", false, "Verbose output")
	baseModule = flag.String("module", "github.com/tmc/appledocs/generated", "Base module path")
)

func main() {
	flag.Parse()

	if *verbose {
		fmt.Fprintf(os.Stderr, "Using go list to analyze packages in: %s\n", *dir)
	}

	// Get package information using go list
	packages, err := getPackagesWithGoList(*dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running go list: %v\n", err)
		os.Exit(1)
	}

	if *verbose {
		fmt.Fprintf(os.Stderr, "Found %d packages\n", len(packages))
	}

	// Build dependency graph
	graph := buildGraphFromGoList(packages, *baseModule)

	// Detect cycles
	cycles := detectCycles(graph)

	if *verbose {
		fmt.Fprintf(os.Stderr, "Found %d cycles\n", len(cycles))
	}

	// Output results
	output := map[string]interface{}{
		"cycles_count": len(cycles),
		"cycles":       cycles,
		"packages":     len(packages),
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

	// Print human-readable summary to stderr if verbose
	if *verbose && len(cycles) > 0 {
		fmt.Fprintf(os.Stderr, "\nCycle Details:\n")
		for i, cycle := range cycles {
			fmt.Fprintf(os.Stderr, "\nCycle %d: %s\n", i+1, strings.Join(cycle.Cycle, " → "))
			for _, edge := range cycle.Edges {
				fmt.Fprintf(os.Stderr, "  %s imports %s\n", edge.From, edge.To)
			}
		}
	}

	// Exit with error if cycles found
	if len(cycles) > 0 {
		os.Exit(1)
	}
}

func getPackagesWithGoList(dir string) ([]GoListPackage, error) {
	cmd := exec.Command("go", "list", "-json", "./...")
	cmd.Dir = dir

	output, err := cmd.Output()
	if err != nil {
		// go list can return errors but still produce output
		// Try to parse what we got
		if len(output) == 0 {
			return nil, fmt.Errorf("go list failed: %w", err)
		}
	}

	// go list -json outputs one JSON object per line
	// We need to parse them individually
	var packages []GoListPackage
	decoder := json.NewDecoder(strings.NewReader(string(output)))

	for decoder.More() {
		var pkg GoListPackage
		if err := decoder.Decode(&pkg); err != nil {
			// Skip packages that couldn't be decoded
			continue
		}
		packages = append(packages, pkg)
	}

	return packages, nil
}

func buildGraphFromGoList(packages []GoListPackage, baseModule string) DependencyGraph {
	graph := make(DependencyGraph)

	// First pass: create nodes for all packages
	for _, pkg := range packages {
		if !strings.HasPrefix(pkg.ImportPath, baseModule) {
			continue
		}

		// Extract package name from import path
		pkgName := pkg.Name
		if pkgName == "" {
			// Fallback to last component of import path
			parts := strings.Split(pkg.ImportPath, "/")
			pkgName = parts[len(parts)-1]
		}

		if _, exists := graph[pkgName]; !exists {
			graph[pkgName] = &PackageNode{
				Name:    pkgName,
				Imports: make(map[string]*Edge),
			}
		}
	}

	// Second pass: add edges for imports
	for _, pkg := range packages {
		if !strings.HasPrefix(pkg.ImportPath, baseModule) {
			continue
		}

		pkgName := pkg.Name
		if pkgName == "" {
			parts := strings.Split(pkg.ImportPath, "/")
			pkgName = parts[len(parts)-1]
		}

		node := graph[pkgName]
		for _, imp := range pkg.Imports {
			// Only track imports within our base module
			if !strings.HasPrefix(imp, baseModule) {
				continue
			}

			// Extract imported package name
			impParts := strings.Split(imp, "/")
			impPkgName := impParts[len(impParts)-1]

			// Skip self-imports
			if impPkgName == pkgName {
				continue
			}

			// Check if imported package exists in our graph
			if _, exists := graph[impPkgName]; !exists {
				// Package not in our graph, skip it
				continue
			}

			// Add or update edge
			if _, exists := node.Imports[impPkgName]; !exists {
				node.Imports[impPkgName] = &Edge{
					From:  pkgName,
					To:    impPkgName,
					Files: []string{pkg.Dir},
				}
			}
		}
	}

	return graph
}

// Rest of the code from main.go (DependencyGraph, detectCycles, etc.)
// Copy the type definitions and cycle detection logic here

type CycleInfo struct {
	Cycle []string `json:"cycle"`
	Edges []Edge   `json:"edges"`
}

type Edge struct {
	From  string   `json:"from"`
	To    string   `json:"to"`
	Files []string `json:"files"`
}

type DependencyGraph map[string]*PackageNode

type PackageNode struct {
	Name    string
	Imports map[string]*Edge
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
						cyclePath = append(cyclePath, imported)

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

	return deduplicateCycles(cycles)
}

func deduplicateCycles(cycles []CycleInfo) []CycleInfo {
	seen := make(map[string]bool)
	var unique []CycleInfo

	for _, cycle := range cycles {
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

	minRotation := cycle
	for i := 1; i < len(cycle); i++ {
		rotation := append(cycle[i:], cycle[:i]...)
		if strings.Join(rotation, ",") < strings.Join(minRotation, ",") {
			minRotation = rotation
		}
	}

	return strings.Join(minRotation, ",")
}
