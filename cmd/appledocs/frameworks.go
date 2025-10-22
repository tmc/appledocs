package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// FrameworkInfo contains information about an available framework
type FrameworkInfo struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
}

// FrameworkList contains the list of available frameworks
type FrameworkList struct {
	Frameworks  []FrameworkInfo `json:"frameworks"`
	LastUpdated time.Time       `json:"last_updated"`
	Source      string          `json:"source"`
}

// getFrameworkList fetches and caches the list of available frameworks from technologies.json
func getFrameworkList(ctx context.Context, cacheDir, baseURL string, refresh bool) (*FrameworkList, error) {
	// Cache path for framework list
	listPath := filepath.Join(cacheDir, "framework-list.json")

	// Try to load from cache first (unless refresh requested)
	if !refresh {
		if data, err := os.ReadFile(listPath); err == nil {
			var list FrameworkList
			if err := json.Unmarshal(data, &list); err == nil {
				// Check if cache is still fresh (less than 7 days old)
				if time.Since(list.LastUpdated) < 7*24*time.Hour {
					return &list, nil
				}
			}
		}
	}

	// Fetch technologies.json
	technologiesURL := baseURL + "/tutorials/data/documentation/technologies.json"
	logger.Info("Fetching framework list", "url", technologiesURL)

	req, err := http.NewRequestWithContext(ctx, "GET", technologiesURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch technologies.json: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse technologies.json
	var doc map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&doc); err != nil {
		return nil, fmt.Errorf("failed to parse technologies.json: %w", err)
	}

	// Extract frameworks from references
	frameworks := make([]FrameworkInfo, 0, 300)
	if references, ok := doc["references"].(map[string]interface{}); ok {
		for refID, refData := range references {
			// Only process doc:// URLs that point to frameworks
			if !strings.HasPrefix(refID, "doc://com.apple.documentation/documentation/") {
				continue
			}

			ref, ok := refData.(map[string]interface{})
			if !ok {
				continue
			}

			// Get framework name from identifier
			identifier := refID
			frameworkName := strings.TrimPrefix(identifier, "doc://com.apple.documentation/documentation/")

			// Skip nested paths (we only want top-level frameworks)
			if strings.Contains(frameworkName, "/") {
				continue
			}

			// Get title and URL
			title, _ := ref["title"].(string)
			url, _ := ref["url"].(string)

			// Skip if no title
			if title == "" {
				title = frameworkName
			}

			// Get description from abstract
			var description string
			if abstract, ok := ref["abstract"].([]interface{}); ok && len(abstract) > 0 {
				if abstractItem, ok := abstract[0].(map[string]interface{}); ok {
					if text, ok := abstractItem["text"].(string); ok {
						description = text
					}
				}
			}

			frameworks = append(frameworks, FrameworkInfo{
				Name:        frameworkName,
				Title:       title,
				URL:         url,
				Description: description,
			})
		}
	}

	// Sort by name
	sort.Slice(frameworks, func(i, j int) bool {
		return frameworks[i].Name < frameworks[j].Name
	})

	// Create framework list
	list := &FrameworkList{
		Frameworks:  frameworks,
		LastUpdated: time.Now(),
		Source:      technologiesURL,
	}

	// Cache the list
	if err := os.MkdirAll(filepath.Dir(listPath), 0755); err == nil {
		if data, err := json.MarshalIndent(list, "", "  "); err == nil {
			os.WriteFile(listPath, data, 0644)
		}
	}

	return list, nil
}

// listAvailableFrameworks lists all available frameworks
func listAvailableFrameworks(ctx context.Context, cacheDir, baseURL string, jsonOutput, refresh bool) error {
	list, err := getFrameworkList(ctx, cacheDir, baseURL, refresh)
	if err != nil {
		return fmt.Errorf("failed to get framework list: %w", err)
	}

	if jsonOutput {
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		return encoder.Encode(list)
	}

	// Human-readable output
	fmt.Printf("\nAvailable Frameworks (%d total):\n", len(list.Frameworks))
	fmt.Printf("%s\n", strings.Repeat("=", 80))
	fmt.Printf("Last updated: %s\n\n", list.LastUpdated.Format("2006-01-02 15:04:05"))

	for _, fw := range list.Frameworks {
		fmt.Printf("%-30s  %s\n", fw.Name, fw.Title)
		if fw.Description != "" && len(fw.Description) < 100 {
			fmt.Printf("    %s\n", fw.Description)
		}
	}
	fmt.Printf("\n")
	fmt.Printf("Usage: ./appledocs -framework=<name>\n")
	fmt.Printf("Example: ./appledocs -framework=Foundation\n\n")

	return nil
}

// findFramework searches for a framework by name (case-insensitive)
func findFramework(ctx context.Context, cacheDir, baseURL, frameworkName string) (*FrameworkInfo, error) {
	list, err := getFrameworkList(ctx, cacheDir, baseURL, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get framework list: %w", err)
	}

	// Try exact match first
	for _, fw := range list.Frameworks {
		if strings.EqualFold(fw.Name, frameworkName) {
			return &fw, nil
		}
	}

	// Try case-insensitive match
	nameLower := strings.ToLower(frameworkName)
	for _, fw := range list.Frameworks {
		if strings.ToLower(fw.Name) == nameLower {
			return &fw, nil
		}
	}

	// Try prefix match
	matches := []FrameworkInfo{}
	for _, fw := range list.Frameworks {
		if strings.HasPrefix(strings.ToLower(fw.Name), nameLower) {
			matches = append(matches, fw)
		}
	}

	if len(matches) == 1 {
		logger.Info("Found framework by prefix match", "input", frameworkName, "match", matches[0].Name)
		return &matches[0], nil
	}

	if len(matches) > 1 {
		names := make([]string, len(matches))
		for i, m := range matches {
			names[i] = m.Name
		}
		return nil, fmt.Errorf("ambiguous framework name %q, matches: %s", frameworkName, strings.Join(names, ", "))
	}

	return nil, fmt.Errorf("framework %q not found. Use -mode=list-frameworks to see all available frameworks", frameworkName)
}

// resolveFrameworkEntryPoint resolves a framework name to its entry point URL
func resolveFrameworkEntryPoint(ctx context.Context, cacheDir, baseURL, frameworkName string) (string, error) {
	fw, err := findFramework(ctx, cacheDir, baseURL, frameworkName)
	if err != nil {
		return "", err
	}

	// Use the framework name (which may have been normalized)
	return fw.Name, nil
}
