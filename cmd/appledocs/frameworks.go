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
	"text/tabwriter"
	"time"
)

// PlatformInfo contains platform availability information
type PlatformInfo struct {
	Name         string `json:"name"`
	IntroducedAt string `json:"introduced_at,omitempty"`
	Beta         bool   `json:"beta,omitempty"`
	Deprecated   bool   `json:"deprecated,omitempty"`
	Unavailable  bool   `json:"unavailable,omitempty"`
}

// FrameworkInfo contains information about an available framework
type FrameworkInfo struct {
	Name        string          `json:"name"`
	Title       string          `json:"title"`
	URL         string          `json:"url"`
	Description string          `json:"description,omitempty"`
	Platforms   []PlatformInfo  `json:"platforms,omitempty"`
}

// FrameworkList contains the list of available frameworks
type FrameworkList struct {
	Frameworks  []FrameworkInfo `json:"frameworks"`
	LastUpdated time.Time       `json:"last_updated"`
	Source      string          `json:"source"`
}

// ListFilters contains filtering options for framework listing
type ListFilters struct {
	Platform   string
	MinVersion string
	Pattern    string
	ShowAll    bool
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

			fw := FrameworkInfo{
				Name:        frameworkName,
				Title:       title,
				URL:         url,
				Description: description,
			}

			// Try to get platform information from cached framework JSON
			fw.Platforms = getFrameworkPlatforms(cacheDir, frameworkName)

			frameworks = append(frameworks, fw)
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
// getFrameworkPlatforms extracts platform information from a framework's cached JSON
func getFrameworkPlatforms(cacheDir, frameworkName string) []PlatformInfo {
	// Try both lowercase and original case
	possiblePaths := []string{
		filepath.Join(cacheDir, "developer.apple.com/tutorials/data/documentation", strings.ToLower(frameworkName), "index.json"),
		filepath.Join(cacheDir, "developer.apple.com/tutorials/data/documentation", frameworkName, "index.json"),
	}

	var data []byte
	var err error
	for _, jsonPath := range possiblePaths {
		data, err = os.ReadFile(jsonPath)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil
	}

	// Parse JSON
	var doc map[string]interface{}
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil
	}

	// Extract platform metadata
	metadata, ok := doc["metadata"].(map[string]interface{})
	if !ok {
		return nil
	}

	platformsData, ok := metadata["platforms"].([]interface{})
	if !ok {
		return nil
	}

	platforms := make([]PlatformInfo, 0, len(platformsData))
	for _, p := range platformsData {
		pMap, ok := p.(map[string]interface{})
		if !ok {
			continue
		}

		platform := PlatformInfo{
			Name: getString(pMap, "name"),
			IntroducedAt: getString(pMap, "introducedAt"),
			Beta: getBool(pMap, "beta"),
			Deprecated: getBool(pMap, "deprecated"),
			Unavailable: getBool(pMap, "unavailable"),
		}

		platforms = append(platforms, platform)
	}

	return platforms
}

func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

func getBool(m map[string]interface{}, key string) bool {
	if v, ok := m[key].(bool); ok {
		return v
	}
	return false
}

// filterFrameworks applies filters to the framework list
func filterFrameworks(frameworks []FrameworkInfo, filters ListFilters) []FrameworkInfo {
	if filters.Platform == "" && filters.MinVersion == "" && filters.Pattern == "" {
		return frameworks
	}

	filtered := make([]FrameworkInfo, 0, len(frameworks))
	for _, fw := range frameworks {
		// Pattern filter
		if filters.Pattern != "" {
			matched, _ := filepath.Match(filters.Pattern, fw.Name)
			if !matched {
				continue
			}
		}

		// Platform filter
		if filters.Platform != "" {
			hasPlatform := false
			for _, p := range fw.Platforms {
				if strings.EqualFold(p.Name, filters.Platform) {
					hasPlatform = true

					// MinVersion filter
					if filters.MinVersion != "" && p.IntroducedAt != "" {
						// Simple version comparison (works for X.Y format)
						if compareVersion(p.IntroducedAt, filters.MinVersion) >= 0 {
							break
						}
						hasPlatform = false
					}
					break
				}
			}
			if !hasPlatform {
				continue
			}
		}

		filtered = append(filtered, fw)
	}

	return filtered
}

// compareVersion compares two version strings (e.g., "15.0" vs "14.0")
// Returns: -1 if v1 < v2, 0 if v1 == v2, 1 if v1 > v2
func compareVersion(v1, v2 string) int {
	parts1 := strings.Split(v1, ".")
	parts2 := strings.Split(v2, ".")

	maxLen := len(parts1)
	if len(parts2) > maxLen {
		maxLen = len(parts2)
	}

	for i := 0; i < maxLen; i++ {
		var n1, n2 int
		if i < len(parts1) {
			fmt.Sscanf(parts1[i], "%d", &n1)
		}
		if i < len(parts2) {
			fmt.Sscanf(parts2[i], "%d", &n2)
		}

		if n1 < n2 {
			return -1
		}
		if n1 > n2 {
			return 1
		}
	}

	return 0
}

func listAvailableFrameworks(ctx context.Context, cacheDir, baseURL string, jsonOutput, refresh bool, filters ListFilters) error {
	list, err := getFrameworkList(ctx, cacheDir, baseURL, refresh)
	if err != nil {
		return fmt.Errorf("failed to get framework list: %w", err)
	}

	// Apply filters
	filteredFrameworks := filterFrameworks(list.Frameworks, filters)
	list.Frameworks = filteredFrameworks

	if jsonOutput {
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		return encoder.Encode(list)
	}

	// Human-readable output (table format)
	fmt.Printf("\nAvailable Frameworks (%d total)\n", len(list.Frameworks))
	fmt.Printf("Last updated: %s\n\n", list.LastUpdated.Format("2006-01-02 15:04:05"))

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, "FRAMEWORK\tPLATFORMS\n")
	fmt.Fprintf(w, "%s\t%s\n", strings.Repeat("-", 30), strings.Repeat("-", 80))

	for _, fw := range list.Frameworks {
		// Format platform information
		platformStr := ""
		if len(fw.Platforms) > 0 {
			platformStrs := make([]string, 0, len(fw.Platforms))
			for _, p := range fw.Platforms {
				pStr := p.Name
				if p.IntroducedAt != "" {
					pStr += " " + p.IntroducedAt + "+"
				}
				platformStrs = append(platformStrs, pStr)
			}
			platformStr = strings.Join(platformStrs, ", ")
		} else {
			platformStr = "(not cached)"
		}

		fmt.Fprintf(w, "%s\t%s\n", fw.Name, platformStr)
	}

	w.Flush()

	fmt.Printf("\nUsage: appledocs crawl <framework>\n")
	fmt.Printf("       appledocs list --platform macOS\n")
	fmt.Printf("       appledocs list --pattern '^Core'\n")
	fmt.Printf("Example: appledocs crawl Foundation\n\n")

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
