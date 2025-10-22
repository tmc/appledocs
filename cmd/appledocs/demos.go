package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// DemoCodeInfo holds information about a sample code download
type DemoCodeInfo struct {
	Title       string
	Framework   string
	DownloadURL string
	DocURL      string
}

// listDemoCode scans cached documentation and lists all available sample code downloads
func listDemoCode(cacheDir, frameworkFilter string, download bool, outputDir string) error {
	demos := []DemoCodeInfo{}

	// Build the search path
	searchPath := filepath.Join(cacheDir, "developer.apple.com", "tutorials", "data", "documentation")

	// Walk through cached JSON files
	err := filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip files with errors
		}

		if !info.IsDir() && filepath.Ext(path) == ".json" {
			// Extract framework name from path
			relPath, _ := filepath.Rel(searchPath, path)
			framework := filepath.Base(filepath.Dir(relPath))

			// Apply framework filter if specified
			if frameworkFilter != "" && !strings.EqualFold(framework, frameworkFilter) {
				return nil
			}

			// Read and parse the JSON file
			data, err := os.ReadFile(path)
			if err != nil {
				return nil // Skip unreadable files
			}

			var doc map[string]interface{}
			if err := json.Unmarshal(data, &doc); err != nil {
				return nil // Skip malformed JSON
			}

			// Check for sampleCodeDownload field
			if sampleCodeDownload, ok := doc["sampleCodeDownload"].(map[string]interface{}); ok {
				if action, ok := sampleCodeDownload["action"].(map[string]interface{}); ok {
					if downloadPath, ok := action["identifier"].(string); ok {
						// Build full download URL
						var downloadURL string
						if strings.HasPrefix(downloadPath, "http://") || strings.HasPrefix(downloadPath, "https://") {
							downloadURL = downloadPath
						} else {
							downloadURL = "https://docs-assets.developer.apple.com/published/" + downloadPath
						}

						// Get title from metadata or identifier
						title := ""
						if metadata, ok := doc["metadata"].(map[string]interface{}); ok {
							if t, ok := metadata["title"].(string); ok {
								title = t
							}
						}
						if title == "" {
							if identifier, ok := doc["identifier"].(map[string]interface{}); ok {
								if url, ok := identifier["url"].(string); ok {
									title = url
								}
							}
						}

						// Build doc URL
						docURL := ""
						if identifier, ok := doc["identifier"].(map[string]interface{}); ok {
							if url, ok := identifier["url"].(string); ok {
								docURL = "https://developer.apple.com" + url
							}
						}

						demos = append(demos, DemoCodeInfo{
							Title:       title,
							Framework:   framework,
							DownloadURL: downloadURL,
							DocURL:      docURL,
						})
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to walk cache directory: %w", err)
	}

	// Sort by framework, then title
	sort.Slice(demos, func(i, j int) bool {
		if demos[i].Framework != demos[j].Framework {
			return demos[i].Framework < demos[j].Framework
		}
		return demos[i].Title < demos[j].Title
	})

	// Download demos if requested
	if download {
		if err := downloadDemoCode(demos, outputDir); err != nil {
			return fmt.Errorf("failed to download demo code: %w", err)
		}
		return nil
	}

	// Print results
	fmt.Printf("\nFound %d demo code examples:\n\n", len(demos))

	currentFramework := ""
	for _, demo := range demos {
		if demo.Framework != currentFramework {
			currentFramework = demo.Framework
			fmt.Printf("\n%s:\n", currentFramework)
			fmt.Printf("%s\n", strings.Repeat("-", len(currentFramework)+1))
		}
		fmt.Printf("  • %s\n", demo.Title)
		fmt.Printf("    Download: %s\n", demo.DownloadURL)
		if demo.DocURL != "" {
			fmt.Printf("    Documentation: %s\n", demo.DocURL)
		}
		fmt.Println()
	}

	return nil
}

// downloadDemoCode downloads all demo code examples to the output directory
func downloadDemoCode(demos []DemoCodeInfo, outputDir string) error {
	// Create output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	fmt.Printf("\nDownloading %d demo code examples to %s\n\n", len(demos), outputDir)

	// Track statistics
	successful := 0
	failed := 0

	// Download each demo
	for i, demo := range demos {
		// Create framework subdirectory
		frameworkDir := filepath.Join(outputDir, demo.Framework)
		if err := os.MkdirAll(frameworkDir, 0755); err != nil {
			if logger != nil {
				logger.Error("Failed to create framework directory", "framework", demo.Framework, "error", err)
			}
			failed++
			continue
		}

		// Extract filename from URL
		filename := filepath.Base(demo.DownloadURL)
		outputPath := filepath.Join(frameworkDir, filename)

		// Skip if already downloaded
		if _, err := os.Stat(outputPath); err == nil {
			fmt.Printf("[%d/%d] ✓ %s (already exists)\n", i+1, len(demos), demo.Title)
			successful++
			continue
		}

		// Download the file
		fmt.Printf("[%d/%d] Downloading %s...\n", i+1, len(demos), demo.Title)
		resp, err := http.Get(demo.DownloadURL)
		if err != nil {
			if logger != nil {
				logger.Error("Failed to download demo", "title", demo.Title, "url", demo.DownloadURL, "error", err)
			}
			failed++
			continue
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			if logger != nil {
				logger.Error("Failed to download demo", "title", demo.Title, "status", resp.Status)
			}
			failed++
			continue
		}

		// Write to file
		outFile, err := os.Create(outputPath)
		if err != nil {
			resp.Body.Close()
			if logger != nil {
				logger.Error("Failed to create output file", "path", outputPath, "error", err)
			}
			failed++
			continue
		}

		_, err = io.Copy(outFile, resp.Body)
		outFile.Close()
		resp.Body.Close()

		if err != nil {
			os.Remove(outputPath) // Clean up partial download
			if logger != nil {
				logger.Error("Failed to write demo file", "path", outputPath, "error", err)
			}
			failed++
			continue
		}

		fmt.Printf("[%d/%d] ✓ %s\n", i+1, len(demos), demo.Title)
		successful++
	}

	// Print summary
	fmt.Printf("\n%s\n", strings.Repeat("=", 60))
	fmt.Printf("Download complete: %d successful, %d failed\n", successful, failed)
	fmt.Printf("Output directory: %s\n", outputDir)
	fmt.Printf("%s\n\n", strings.Repeat("=", 60))

	return nil
}
