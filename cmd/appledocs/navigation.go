package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// NavItem represents a navigation item with optional children
type NavItem struct {
	Title    string    `json:"title"`
	Path     string    `json:"path"`
	Children []NavItem `json:"children,omitempty"`
}

// NavSection represents a section in the navigation
type NavSection struct {
	Title string    `json:"title"`
	Items []NavItem `json:"items"`
}

// Navigation represents the complete navigation structure
type Navigation struct {
	Sections []NavSection `json:"sections"`
}

// generateNavigation creates the navigation JSON file by scanning the JSON documentation
func generateNavigation(inputDir, outputDir string) error {
	nav := Navigation{
		Sections: []NavSection{},
	}

	// Build a hierarchical structure from the JSON files
	// First, scan for all JSON files
	jsonFiles, err := scanDirectoryForJSON(inputDir)
	if err != nil {
		return fmt.Errorf("scan directory: %v", err)
	}

	// Group files by framework/technology
	frameworks := make(map[string]*frameworkInfo)

	for _, jsonFile := range jsonFiles {
		// Read JSON to get metadata
		data, err := os.ReadFile(jsonFile)
		if err != nil {
			continue
		}

		var doc DocJSONData
		if err := json.Unmarshal(data, &doc); err != nil {
			continue
		}

		// Skip if no hierarchy
		if len(doc.Hierarchy.Paths) == 0 || len(doc.Hierarchy.Paths[0]) == 0 {
			continue
		}

		// Get the relative path for the URL
		relativePath, err := filepath.Rel(inputDir, jsonFile)
		if err != nil {
			continue
		}

		// Convert to URL path (remove .json extension and add leading slash)
		urlPath := "/" + strings.TrimSuffix(relativePath, ".json")

		// Parse the hierarchy to understand the structure
		hierarchy := doc.Hierarchy.Paths[0]

		// Need at least 2 levels (technologies > framework) to be useful
		if len(hierarchy) < 2 {
			continue
		}

		// hierarchy[1] is the actual framework ID
		frameworkID := hierarchy[1]
		frameworkTitle := doc.Metadata.Title
		if ref, ok := doc.References[frameworkID]; ok {
			frameworkTitle = ref.Title
		}

		// Initialize framework info if needed
		if _, exists := frameworks[frameworkID]; !exists {
			frameworks[frameworkID] = &frameworkInfo{
				ID:      frameworkID,
				Title:   frameworkTitle,
				Classes: make(map[string]*classInfo),
			}
		}

		fw := frameworks[frameworkID]

		// Determine what type of document this is based on hierarchy
		// The hierarchy lists the PARENT path, not including the current document
		// hierarchy[0] = technologies
		// hierarchy[1] = framework
		// hierarchy[2] = class (if present)
		switch len(hierarchy) {
		case 2:
			// Hierarchy: [technologies, framework]
			// Could be either framework-level doc OR a class doc
			// Check if the path suggests it's a class
			pathParts := strings.Split(urlPath, "/")
			if len(pathParts) > 4 {
				// This is likely a class document
				// e.g., /tutorials/data/documentation/SecurityFoundation/SFAuthorization
				// Use the document's own identifier as classID if available
				classID := frameworkID + "/" + doc.Metadata.Title
				// Check if there's a self-reference
				for ref := range doc.References {
					if strings.Contains(ref, doc.Metadata.Title) {
						classID = ref
						break
					}
				}

				className := doc.Metadata.Title

				if _, exists := fw.Classes[classID]; !exists {
					fw.Classes[classID] = &classInfo{
						ID:      classID,
						Title:   className,
						Path:    urlPath,
						Methods: []methodInfo{},
					}
				} else {
					// Update the path if it was created earlier without a path
					fw.Classes[classID].Path = urlPath
				}
			} else {
				// This is a framework-level document
				fw.Path = urlPath
			}

		case 3:
			// Hierarchy: [technologies, framework, class]
			// hierarchy[2] is the parent class, this document is a method/property
			classID := hierarchy[2]
			className := ""
			if ref, ok := doc.References[classID]; ok {
				className = ref.Title
			}

			// The current document's title is the method name
			methodID := classID + "/" + doc.Metadata.Title
			methodName := doc.Metadata.Title

			// Ensure class exists
			if _, exists := fw.Classes[classID]; !exists {
				fw.Classes[classID] = &classInfo{
					ID:      classID,
					Title:   className,
					Methods: []methodInfo{},
				}
			}

			// Add method
			fw.Classes[classID].Methods = append(fw.Classes[classID].Methods, methodInfo{
				ID:    methodID,
				Title: methodName,
				Path:  urlPath,
			})
		}
	}

	// Build the navigation structure from the collected data
	// Add Overview section
	overviewItems := []NavItem{}

	// Add Technologies link if we have any frameworks
	if len(frameworks) > 0 {
		overviewItems = append(overviewItems, NavItem{
			Title: "Technologies",
			Path:  "/tutorials/data/documentation/technologies",
		})
	}

	// Add framework links to overview
	frameworkList := make([]*frameworkInfo, 0, len(frameworks))
	for _, fw := range frameworks {
		frameworkList = append(frameworkList, fw)
	}
	sort.Slice(frameworkList, func(i, j int) bool {
		return frameworkList[i].Title < frameworkList[j].Title
	})

	for _, fw := range frameworkList {
		if fw.Path != "" {
			overviewItems = append(overviewItems, NavItem{
				Title: fw.Title,
				Path:  fw.Path,
			})
		}
	}

	if len(overviewItems) > 0 {
		nav.Sections = append(nav.Sections, NavSection{
			Title: "Overview",
			Items: overviewItems,
		})
	}

	// Add framework sections
	for _, fw := range frameworkList {
		if len(fw.Classes) == 0 {
			continue
		}

		// Sort classes
		classList := make([]*classInfo, 0, len(fw.Classes))
		for _, cls := range fw.Classes {
			classList = append(classList, cls)
		}
		sort.Slice(classList, func(i, j int) bool {
			return classList[i].Title < classList[j].Title
		})

		// Build nav items for classes
		classItems := []NavItem{}
		for _, cls := range classList {
			item := NavItem{
				Title: cls.Title,
				Path:  cls.Path,
			}

			// Sort methods
			sort.Slice(cls.Methods, func(i, j int) bool {
				return cls.Methods[i].Title < cls.Methods[j].Title
			})

			// Add children (methods)
			if len(cls.Methods) > 0 {
				item.Children = make([]NavItem, 0, len(cls.Methods))
				for _, method := range cls.Methods {
					item.Children = append(item.Children, NavItem{
						Title: method.Title,
						Path:  method.Path,
					})
				}
			}

			classItems = append(classItems, item)
		}

		if len(classItems) > 0 {
			nav.Sections = append(nav.Sections, NavSection{
				Title: fw.Title,
				Items: classItems,
			})
		}
	}

	// Write navigation JSON file
	navFile := filepath.Join(outputDir, "docs-navigation.json")
	navData, err := json.MarshalIndent(nav, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal navigation: %v", err)
	}

	if err := os.WriteFile(navFile, navData, 0644); err != nil {
		return fmt.Errorf("write navigation file: %v", err)
	}

	fmt.Printf("Generated navigation: %s\n", navFile)
	return nil
}

// frameworkInfo holds information about a framework
type frameworkInfo struct {
	ID      string
	Title   string
	Path    string
	Classes map[string]*classInfo
}

// classInfo holds information about a class
type classInfo struct {
	ID      string
	Title   string
	Path    string
	Methods []methodInfo
}

// methodInfo holds information about a method
type methodInfo struct {
	ID    string
	Title string
	Path  string
}
