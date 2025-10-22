package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// FrameworkSummary contains high-level statistics about a framework
type FrameworkSummary struct {
	Name                 string            `json:"name"`
	Version              string            `json:"version,omitempty"`
	Platforms            []string          `json:"platforms,omitempty"`
	ClassCount           int               `json:"class_count"`
	ProtocolCount        int               `json:"protocol_count"`
	EnumCount            int               `json:"enum_count"`
	StructCount          int               `json:"struct_count"`
	TypeAliasCount       int               `json:"type_alias_count"`
	FunctionCount        int               `json:"function_count"`
	ConstantCount        int               `json:"constant_count"`
	TotalSymbols         int               `json:"total_symbols"`
	TotalMethods         int               `json:"total_methods"`
	TotalProperties      int               `json:"total_properties"`
	TopLevelCategories   []string          `json:"top_level_categories,omitempty"`
	Dependencies         []string          `json:"dependencies,omitempty"`
	APICoverage          float64           `json:"api_coverage_percent,omitempty"`
	LastCacheUpdate      time.Time         `json:"last_cache_update,omitempty"`
	CachedFiles          int               `json:"cached_files"`
	SymbolsByKind        map[string]int    `json:"symbols_by_kind,omitempty"`
}

// generateFrameworkSummary analyzes cached framework data and generates a summary
func generateFrameworkSummary(cacheDir, framework string, jsonOutput bool) error {
	// Normalize framework name to lowercase for directory lookups
	frameworkLower := strings.ToLower(framework)
	frameworkDir := filepath.Join(cacheDir, "developer.apple.com", "tutorials", "data", "documentation", frameworkLower)

	// Check if framework directory exists
	if _, err := os.Stat(frameworkDir); os.IsNotExist(err) {
		// Try with original case
		frameworkDir = filepath.Join(cacheDir, "developer.apple.com", "tutorials", "data", "documentation", framework)
		if _, err := os.Stat(frameworkDir); os.IsNotExist(err) {
			return fmt.Errorf("framework %q not found in cache. Run with -mode=crawl first", framework)
		}
	}

	summary := FrameworkSummary{
		Name:          framework,
		SymbolsByKind: make(map[string]int),
	}

	// Read the main framework JSON file
	frameworkJSON := filepath.Join(cacheDir, "developer.apple.com", "tutorials", "data", "documentation", framework+".json")
	if _, err := os.Stat(frameworkJSON); err == nil {
		data, err := os.ReadFile(frameworkJSON)
		if err == nil {
			var doc map[string]interface{}
			if json.Unmarshal(data, &doc) == nil {
				// Extract metadata
				if metadata, ok := doc["metadata"].(map[string]interface{}); ok {
					// Extract platforms
					if platforms, ok := metadata["platforms"].([]interface{}); ok {
						for _, p := range platforms {
							if platform, ok := p.(map[string]interface{}); ok {
								if name, ok := platform["name"].(string); ok {
									summary.Platforms = append(summary.Platforms, name)
								}
							}
						}
					}

					// Extract version if available
					if modules, ok := metadata["modules"].([]interface{}); ok && len(modules) > 0 {
						if module, ok := modules[0].(map[string]interface{}); ok {
							if relatedModules, ok := module["relatedModules"].([]interface{}); ok {
								summary.Dependencies = make([]string, 0, len(relatedModules))
								for _, rm := range relatedModules {
									if rmStr, ok := rm.(string); ok {
										summary.Dependencies = append(summary.Dependencies, rmStr)
									}
								}
							}
						}
					}
				}

				// Extract top-level categories
				if topicSections, ok := doc["topicSections"].([]interface{}); ok {
					for _, section := range topicSections {
						if sectionMap, ok := section.(map[string]interface{}); ok {
							if title, ok := sectionMap["title"].(string); ok {
								summary.TopLevelCategories = append(summary.TopLevelCategories, title)
							}
						}
					}
				}

				// Get file modification time
				if fileInfo, err := os.Stat(frameworkJSON); err == nil {
					summary.LastCacheUpdate = fileInfo.ModTime()
				}
			}
		}
	}

	// Count files in framework directory
	symbolFiles := []string{}
	err := filepath.Walk(frameworkDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			symbolFiles = append(symbolFiles, path)
			summary.CachedFiles++

			// Update last cache time if this file is newer
			if info.ModTime().After(summary.LastCacheUpdate) {
				summary.LastCacheUpdate = info.ModTime()
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to walk framework directory: %w", err)
	}

	// Analyze symbol files to count types
	logger.Info("Analyzing symbol files", "count", len(symbolFiles))
	for _, symbolPath := range symbolFiles {
		data, err := os.ReadFile(symbolPath)
		if err != nil {
			continue
		}

		var doc map[string]interface{}
		if err := json.Unmarshal(data, &doc); err != nil {
			continue
		}

		// Extract symbol kind from metadata
		if metadata, ok := doc["metadata"].(map[string]interface{}); ok {
			if symbolKind, ok := metadata["symbolKind"].(string); ok {
				summary.SymbolsByKind[symbolKind]++

				// Categorize symbols
				switch symbolKind {
				case "class", "cl":
					summary.ClassCount++
				case "protocol", "intf":
					summary.ProtocolCount++
				case "enum", "tag":
					summary.EnumCount++
				case "struct":
					summary.StructCount++
				case "typealias", "tdef":
					summary.TypeAliasCount++
				case "func", "method", "clm", "instm", "intfm", "intfcm":
					summary.FunctionCount++
				case "var", "data", "instp", "intfp":
					summary.ConstantCount++
				}
			}

			// Check for methods/properties in the document
			if role, ok := metadata["role"].(string); ok {
				if role == "symbol" || role == "collectionGroup" {
					summary.TotalSymbols++
				}
			}
		}

		// Count methods in primaryContentSections
		if sections, ok := doc["primaryContentSections"].([]interface{}); ok {
			for _, section := range sections {
				if sectionMap, ok := section.(map[string]interface{}); ok {
					if kind, ok := sectionMap["kind"].(string); ok {
						if kind == "declarations" {
							if declarations, ok := sectionMap["declarations"].([]interface{}); ok {
								summary.TotalMethods += len(declarations)
							}
						}
					}
				}
			}
		}

		// Count properties in topicSections
		if topicSections, ok := doc["topicSections"].([]interface{}); ok {
			for _, section := range topicSections {
				if sectionMap, ok := section.(map[string]interface{}); ok {
					if identifiers, ok := sectionMap["identifiers"].([]interface{}); ok {
						// Check if this is a properties or methods section
						if title, ok := sectionMap["title"].(string); ok {
							titleLower := strings.ToLower(title)
							if strings.Contains(titleLower, "propert") {
								summary.TotalProperties += len(identifiers)
							} else if strings.Contains(titleLower, "method") || strings.Contains(titleLower, "function") {
								summary.TotalMethods += len(identifiers)
							}
						}
					}
				}
			}
		}
	}

	// Calculate API coverage (percentage of documented symbols with content)
	if summary.CachedFiles > 0 {
		summary.APICoverage = float64(summary.TotalSymbols) / float64(summary.CachedFiles) * 100
		if summary.APICoverage > 100 {
			summary.APICoverage = 100
		}
	}

	// Output the summary
	if jsonOutput {
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		return encoder.Encode(summary)
	}

	// Human-readable text output
	printFrameworkSummary(summary)
	return nil
}

// printFrameworkSummary prints a human-readable summary
func printFrameworkSummary(s FrameworkSummary) {
	fmt.Printf("\n")
	fmt.Printf("Framework Summary: %s\n", s.Name)
	fmt.Printf("%s\n", strings.Repeat("=", len(s.Name)+19))
	fmt.Printf("\n")

	if len(s.Platforms) > 0 {
		fmt.Printf("Platforms:          %s\n", strings.Join(s.Platforms, ", "))
	}
	if s.Version != "" {
		fmt.Printf("Version:            %s\n", s.Version)
	}
	if !s.LastCacheUpdate.IsZero() {
		fmt.Printf("Last Cache Update:  %s\n", s.LastCacheUpdate.Format("2006-01-02 15:04:05"))
	}
	fmt.Printf("\n")

	fmt.Printf("Symbol Statistics:\n")
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("  Classes:          %d\n", s.ClassCount)
	fmt.Printf("  Protocols:        %d\n", s.ProtocolCount)
	fmt.Printf("  Enums:            %d\n", s.EnumCount)
	fmt.Printf("  Structs:          %d\n", s.StructCount)
	fmt.Printf("  Type Aliases:     %d\n", s.TypeAliasCount)
	fmt.Printf("  Functions:        %d\n", s.FunctionCount)
	fmt.Printf("  Constants:        %d\n", s.ConstantCount)
	fmt.Printf("  %s\n", strings.Repeat("-", 38))
	fmt.Printf("  Total Symbols:    %d\n", s.TotalSymbols)
	fmt.Printf("\n")

	fmt.Printf("Content Statistics:\n")
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("  Total Methods:    %d\n", s.TotalMethods)
	fmt.Printf("  Total Properties: %d\n", s.TotalProperties)
	fmt.Printf("  Cached Files:     %d\n", s.CachedFiles)
	if s.APICoverage > 0 {
		fmt.Printf("  API Coverage:     %.1f%%\n", s.APICoverage)
	}
	fmt.Printf("\n")

	if len(s.SymbolsByKind) > 0 {
		fmt.Printf("Symbols by Kind:\n")
		fmt.Printf("%s\n", strings.Repeat("-", 40))

		// Sort by count descending
		type kindCount struct {
			kind  string
			count int
		}
		kinds := make([]kindCount, 0, len(s.SymbolsByKind))
		for k, v := range s.SymbolsByKind {
			kinds = append(kinds, kindCount{k, v})
		}
		sort.Slice(kinds, func(i, j int) bool {
			return kinds[i].count > kinds[j].count
		})

		for _, kc := range kinds {
			fmt.Printf("  %-20s %d\n", kc.kind+":", kc.count)
		}
		fmt.Printf("\n")
	}

	if len(s.TopLevelCategories) > 0 {
		fmt.Printf("Top-Level Categories:\n")
		fmt.Printf("%s\n", strings.Repeat("-", 40))
		for _, cat := range s.TopLevelCategories {
			fmt.Printf("  • %s\n", cat)
		}
		fmt.Printf("\n")
	}

	if len(s.Dependencies) > 0 {
		fmt.Printf("Dependencies:\n")
		fmt.Printf("%s\n", strings.Repeat("-", 40))
		for _, dep := range s.Dependencies {
			fmt.Printf("  • %s\n", dep)
		}
		fmt.Printf("\n")
	}
}

// listCachedFrameworks lists all frameworks that have been cached
func listCachedFrameworks(cacheDir string, jsonOutput bool) error {
	docsPath := filepath.Join(cacheDir, "developer.apple.com", "tutorials", "data", "documentation")

	// Check if documentation directory exists
	if _, err := os.Stat(docsPath); os.IsNotExist(err) {
		return fmt.Errorf("no cached documentation found in %s", cacheDir)
	}

	// Read directory
	entries, err := os.ReadDir(docsPath)
	if err != nil {
		return fmt.Errorf("failed to read documentation directory: %w", err)
	}

	type frameworkInfo struct {
		Name      string    `json:"name"`
		FileCount int       `json:"file_count"`
		LastMod   time.Time `json:"last_modified"`
	}

	frameworks := []frameworkInfo{}

	for _, entry := range entries {
		if entry.IsDir() {
			// Count files in this framework directory
			frameworkPath := filepath.Join(docsPath, entry.Name())
			fileCount := 0
			var lastMod time.Time

			filepath.Walk(frameworkPath, func(path string, info os.FileInfo, err error) error {
				if err == nil && !info.IsDir() && filepath.Ext(path) == ".json" {
					fileCount++
					if info.ModTime().After(lastMod) {
						lastMod = info.ModTime()
					}
				}
				return nil
			})

			if fileCount > 0 {
				frameworks = append(frameworks, frameworkInfo{
					Name:      entry.Name(),
					FileCount: fileCount,
					LastMod:   lastMod,
				})
			}
		}
	}

	// Sort by name
	sort.Slice(frameworks, func(i, j int) bool {
		return frameworks[i].Name < frameworks[j].Name
	})

	if jsonOutput {
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		return encoder.Encode(map[string]interface{}{
			"total_frameworks": len(frameworks),
			"frameworks":       frameworks,
		})
	}

	// Human-readable output
	fmt.Printf("\nCached Frameworks (%d total):\n", len(frameworks))
	fmt.Printf("%s\n\n", strings.Repeat("=", 70))

	for _, fw := range frameworks {
		fmt.Printf("%-30s  %6d files  (updated: %s)\n",
			fw.Name,
			fw.FileCount,
			fw.LastMod.Format("2006-01-02"),
		)
	}
	fmt.Printf("\n")

	return nil
}
