package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// AnalyzeFrameworkDependencies scans Apple documentation to build a framework dependency graph
func AnalyzeFrameworkDependencies(docsRoot string) (map[string][]string, error) {
	// Map of framework -> list of frameworks it depends on
	deps := make(map[string][]string)
	frameworkSet := make(map[string]map[string]bool) // framework -> set of dependencies

	// Walk through all documentation
	err := filepath.Walk(docsRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process index.json files
		if !strings.HasSuffix(path, "index.json") {
			return nil
		}

		// Determine which framework this file belongs to
		rel, _ := filepath.Rel(docsRoot, path)
		parts := strings.Split(rel, string(filepath.Separator))
		if len(parts) < 1 {
			return nil
		}
		currentFramework := parts[0]

		// Read and parse the JSON
		data, err := os.ReadFile(path)
		if err != nil {
			return nil // Skip files we can't read
		}

		var doc map[string]interface{}
		if err := json.Unmarshal(data, &doc); err != nil {
			return nil // Skip invalid JSON
		}

		// Initialize set for this framework if needed
		if frameworkSet[currentFramework] == nil {
			frameworkSet[currentFramework] = make(map[string]bool)
		}

		// Look for cross-framework type references in various places
		extractCrossFrameworkRefs(doc, currentFramework, frameworkSet)

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Convert sets to slices
	for fw, depSet := range frameworkSet {
		for dep := range depSet {
			deps[fw] = append(deps[fw], dep)
		}
	}

	return deps, nil
}

// extractCrossFrameworkRefs extracts cross-framework type references from a documentation JSON
func extractCrossFrameworkRefs(doc map[string]interface{}, currentFramework string, frameworkSet map[string]map[string]bool) {
	// Check primaryContentSections for declarations
	if sections, ok := doc["primaryContentSections"].([]interface{}); ok {
		for _, sec := range sections {
			if section, ok := sec.(map[string]interface{}); ok {
				if decls, ok := section["declarations"].([]interface{}); ok {
					for _, d := range decls {
						if decl, ok := d.(map[string]interface{}); ok {
							extractRefsFromDeclaration(decl, currentFramework, frameworkSet)
						}
					}
				}
			}
		}
	}

	// Check relationshipsSections for inheritance
	if sections, ok := doc["relationshipsSections"].([]interface{}); ok {
		for _, sec := range sections {
			if section, ok := sec.(map[string]interface{}); ok {
				if identifiers, ok := section["identifiers"].([]interface{}); ok {
					for _, id := range identifiers {
						if idStr, ok := id.(string); ok {
							// Extract framework from identifier like "doc://CoreImage/documentation/CoreImage/CIBarcodeDescriptor"
							if fw := extractFrameworkFromIdentifier(idStr); fw != "" && fw != currentFramework {
								frameworkSet[currentFramework][fw] = true
							}
						}
					}
				}
			}
		}
	}

	// Check references for type information
	if refs, ok := doc["references"].(map[string]interface{}); ok {
		for _, ref := range refs {
			if refMap, ok := ref.(map[string]interface{}); ok {
				if fragments, ok := refMap["fragments"].([]interface{}); ok {
					for _, frag := range fragments {
						if fragMap, ok := frag.(map[string]interface{}); ok {
							if preciseID, ok := fragMap["preciseIdentifier"].(string); ok {
								if fw := extractFrameworkFromPreciseIdentifier(preciseID); fw != "" && fw != currentFramework {
									frameworkSet[currentFramework][fw] = true
								}
							}
						}
					}
				}
			}
		}
	}
}

func extractRefsFromDeclaration(decl map[string]interface{}, currentFramework string, frameworkSet map[string]map[string]bool) {
	if tokens, ok := decl["tokens"].([]interface{}); ok {
		for _, tok := range tokens {
			if token, ok := tok.(map[string]interface{}); ok {
				if preciseID, ok := token["preciseIdentifier"].(string); ok {
					if fw := extractFrameworkFromPreciseIdentifier(preciseID); fw != "" && fw != currentFramework {
						frameworkSet[currentFramework][fw] = true
					}
				}
			}
		}
	}
}

// extractFrameworkFromIdentifier extracts framework name from doc:// identifiers
// Example: "doc://CoreImage/documentation/CoreImage/CIBarcodeDescriptor" -> "CoreImage"
func extractFrameworkFromIdentifier(identifier string) string {
	if !strings.HasPrefix(identifier, "doc://") {
		return ""
	}
	identifier = strings.TrimPrefix(identifier, "doc://")
	parts := strings.Split(identifier, "/")
	if len(parts) > 0 && parts[0] != "com.apple.documentation" && parts[0] != "com.externally.resolved.symbol" {
		return parts[0]
	}
	return ""
}

// extractFrameworkFromPreciseIdentifier extracts framework from type prefixes in preciseIdentifiers
// Example: "c:objc(cs)CIBarcodeDescriptor" -> "CoreImage" (by analyzing the CI prefix)
// This is a heuristic approach - we look for common prefixes
func extractFrameworkFromPreciseIdentifier(preciseID string) string {
	// Extract the type name from the preciseIdentifier
	// Format: c:objc(cs)ClassName or c:objc(pl)ProtocolName
	var typeName string

	if idx := strings.LastIndex(preciseID, ")"); idx != -1 && idx < len(preciseID)-1 {
		typeName = preciseID[idx+1:]
	} else {
		return ""
	}

	// Map common prefixes to frameworks
	prefixMap := map[string]string{
		"CI":  "CoreImage",
		"CG":  "CoreGraphics",
		"CA":  "QuartzCore", // CoreAnimation
		"CF":  "CoreFoundation",
		"NS":  "Foundation", // or AppKit - need more context
		"UI":  "UIKit",
		"AV":  "AVFoundation",
		"VN":  "Vision",
		"CK":  "CloudKit",
		"SK":  "SpriteKit",
		"SCN": "SceneKit",
		"AR":  "ARKit",
		"ML":  "CoreML",
		"MT":  "Metal", // or MetalKit
	}

	// Try to match prefix
	for prefix, framework := range prefixMap {
		if strings.HasPrefix(typeName, prefix) {
			return framework
		}
	}

	return ""
}

// BuildFrameworkLayers uses the dependency graph to determine framework layers
// Returns a map of framework -> layer number (0 = base, higher = depends on lower)
func BuildFrameworkLayers(deps map[string][]string) map[string]int {
	layers := make(map[string]int)

	// Start with frameworks that have no dependencies
	for fw, fwDeps := range deps {
		if len(fwDeps) == 0 {
			layers[fw] = 0
		}
	}

	// Iteratively assign layers based on maximum dependency layer + 1
	changed := true
	maxIterations := 100 // Prevent infinite loops
	iterations := 0

	for changed && iterations < maxIterations {
		changed = false
		iterations++

		for fw, fwDeps := range deps {
			if _, assigned := layers[fw]; assigned {
				continue // Already assigned
			}

			// Check if all dependencies have been assigned
			maxDepLayer := -1
			allAssigned := true
			for _, dep := range fwDeps {
				if depLayer, ok := layers[dep]; ok {
					if depLayer > maxDepLayer {
						maxDepLayer = depLayer
					}
				} else {
					allAssigned = false
					break
				}
			}

			if allAssigned {
				layers[fw] = maxDepLayer + 1
				changed = true
			}
		}
	}

	// Assign remaining frameworks (those in circular deps) a high layer number
	for fw := range deps {
		if _, assigned := layers[fw]; !assigned {
			layers[fw] = 99 // High number for circular dependencies
		}
	}

	return layers
}

func main() {
	docsRoot := filepath.Join(os.Getenv("HOME"), ".appledocs/cache/developer.apple.com/tutorials/data/documentation")

	fmt.Println("Analyzing framework dependencies from Apple documentation...")
	deps, err := AnalyzeFrameworkDependencies(docsRoot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nFramework Dependencies:")
	for fw, fwDeps := range deps {
		if len(fwDeps) > 0 {
			fmt.Printf("%s -> %v\n", fw, fwDeps)
		}
	}

	fmt.Println("\nBuilding framework layers...")
	layers := BuildFrameworkLayers(deps)

	// Group by layer
	layerGroups := make(map[int][]string)
	for fw, layer := range layers {
		layerGroups[layer] = append(layerGroups[layer], fw)
	}

	fmt.Println("\nFramework Layers:")
	for i := 0; i <= 10; i++ {
		if fws, ok := layerGroups[i]; ok {
			fmt.Printf("Layer %d: %v\n", i, fws)
		}
	}
	if fws, ok := layerGroups[99]; ok {
		fmt.Printf("Layer 99 (circular deps): %v\n", fws)
	}
}
