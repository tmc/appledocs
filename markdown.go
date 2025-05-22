// Package main implements Markdown rendering functionality for Apple docs
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// DocJSONData represents the structure of the Apple documentation JSON
type DocJSONData struct {
	Abstract               []TextContent         `json:"abstract,omitempty"`
	Hierarchy              Hierarchy             `json:"hierarchy,omitempty"`
	Identifier             Identifier            `json:"identifier,omitempty"`
	Kind                   string                `json:"kind,omitempty"`
	Metadata               Metadata              `json:"metadata,omitempty"`
	PrimaryContentSections []ContentSection      `json:"primaryContentSections,omitempty"`
	References             map[string]Reference  `json:"references,omitempty"`
	TopicSections          []TopicSection        `json:"topicSections,omitempty"`
	RelationshipsSections  []RelationshipSection `json:"relationshipsSections,omitempty"`
	SeeAlsoSections        []SeeAlsoSection      `json:"seeAlsoSections,omitempty"`
}

// Hierarchy represents the document hierarchy
type Hierarchy struct {
	Paths [][]string `json:"paths,omitempty"`
}

// Identifier contains document identifier information
type Identifier struct {
	InterfaceLanguage string `json:"interfaceLanguage,omitempty"`
	URL               string `json:"url,omitempty"`
}

// Metadata contains document metadata
type Metadata struct {
	ExternalID     string     `json:"externalID,omitempty"`
	Fragments      []Fragment `json:"fragments,omitempty"`
	Modules        []Module   `json:"modules,omitempty"`
	Role           string     `json:"role,omitempty"`
	RoleHeading    string     `json:"roleHeading,omitempty"`
	SymbolKind     string     `json:"symbolKind,omitempty"`
	Title          string     `json:"title,omitempty"`
	Platforms      []Platform `json:"platforms,omitempty"`
	NavigatorTitle []Fragment `json:"navigatorTitle,omitempty"`
}

// Platform represents a platform compatibility information
type Platform struct {
	Beta         bool   `json:"beta,omitempty"`
	Deprecated   bool   `json:"deprecated,omitempty"`
	IntroducedAt string `json:"introducedAt,omitempty"`
	Name         string `json:"name,omitempty"`
	Unavailable  bool   `json:"unavailable,omitempty"`
}

// Module represents a module reference
type Module struct {
	Name string `json:"name,omitempty"`
}

// Fragment represents a text fragment
type Fragment struct {
	Kind string `json:"kind,omitempty"`
	Text string `json:"text,omitempty"`
}

// TextContent represents a text content block
type TextContent struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

// ContentSection represents a section of content
type ContentSection struct {
	Kind         string         `json:"kind,omitempty"`
	Content      []ContentBlock `json:"content,omitempty"`
	Declarations []Declaration  `json:"declarations,omitempty"`
	Mentions     []string       `json:"mentions,omitempty"`
}

// Declaration represents a code declaration
type Declaration struct {
	Languages []string   `json:"languages,omitempty"`
	Platforms []string   `json:"platforms,omitempty"`
	Tokens    []Fragment `json:"tokens,omitempty"`
}

// ContentBlock represents a block of content
type ContentBlock struct {
	Type          string          `json:"type,omitempty"`
	Anchor        string          `json:"anchor,omitempty"`
	Level         int             `json:"level,omitempty"`
	Text          string          `json:"text,omitempty"`
	InlineContent []InlineContent `json:"inlineContent,omitempty"`
	Items         []Item          `json:"items,omitempty"`
	Content       []ContentBlock  `json:"content,omitempty"`
	Name          string          `json:"name,omitempty"`
	Style         string          `json:"style,omitempty"`
}

// InlineContent represents inline content elements
type InlineContent struct {
	Type          string          `json:"type,omitempty"`
	Text          string          `json:"text,omitempty"`
	Code          string          `json:"code,omitempty"`
	InlineContent []InlineContent `json:"inlineContent,omitempty"`
	Identifier    string          `json:"identifier,omitempty"`
	IsActive      bool            `json:"isActive,omitempty"`
}

// Item represents an item in a list
type Item struct {
	Content []ContentBlock `json:"content,omitempty"`
}

// TopicSection represents a section of related topics
type TopicSection struct {
	Anchor      string   `json:"anchor,omitempty"`
	Identifiers []string `json:"identifiers,omitempty"`
	Title       string   `json:"title,omitempty"`
	Generated   bool     `json:"generated,omitempty"`
}

// RelationshipSection represents a section of related items
type RelationshipSection struct {
	Kind        string   `json:"kind,omitempty"`
	Title       string   `json:"title,omitempty"`
	Type        string   `json:"type,omitempty"`
	Identifiers []string `json:"identifiers,omitempty"`
}

// SeeAlsoSection represents a section of "see also" references
type SeeAlsoSection struct {
	Anchor      string   `json:"anchor,omitempty"`
	Title       string   `json:"title,omitempty"`
	Generated   bool     `json:"generated,omitempty"`
	Identifiers []string `json:"identifiers,omitempty"`
}

// Reference represents a document reference
type Reference struct {
	Abstract       []TextContent `json:"abstract,omitempty"`
	Identifier     string        `json:"identifier,omitempty"`
	Kind           string        `json:"kind,omitempty"`
	Role           string        `json:"role,omitempty"`
	Title          string        `json:"title,omitempty"`
	Type           string        `json:"type,omitempty"`
	URL            string        `json:"url,omitempty"`
	Fragments      []Fragment    `json:"fragments,omitempty"`
	NavigatorTitle []Fragment    `json:"navigatorTitle,omitempty"`
	Beta           bool          `json:"beta,omitempty"`
}

// generateMarkdown generates Markdown documentation from the JSON documentation
func generateMarkdown(inputDir, outputDir string) error {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("create output directory: %v", err)
	}

	// Process JSON files in a set structure to ensure we process them in a logical order
	fmt.Println("Starting to process JSON files...")

	// First, scan the input directory for all JSON files
	jsonFiles, err := scanDirectoryForJSON(inputDir)
	if err != nil {
		return fmt.Errorf("scan directory: %v", err)
	}

	fmt.Printf("Found %d JSON files to process\n", len(jsonFiles))

	// Process files concurrently
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10) // Limit to 10 concurrent operations

	for _, jsonFile := range jsonFiles {
		wg.Add(1)
		semaphore <- struct{}{} // Acquire semaphore

		go func(file string) {
			defer wg.Done()
			defer func() { <-semaphore }() // Release semaphore

			relativePath, err := filepath.Rel(inputDir, file)
			if err != nil {
				fmt.Printf("Error getting relative path for %s: %v\n", file, err)
				return
			}

			// Create corresponding markdown file path
			mdFile := filepath.Join(outputDir, changeExtension(relativePath, ".md"))

			// Ensure output directory exists
			mdDir := filepath.Dir(mdFile)
			if err := os.MkdirAll(mdDir, 0755); err != nil {
				fmt.Printf("Error creating directory for %s: %v\n", mdFile, err)
				return
			}

			// Convert JSON to Markdown
			if err := convertJSONToMarkdown(file, mdFile); err != nil {
				fmt.Printf("Error converting %s to markdown: %v\n", file, err)
			}
		}(jsonFile)
	}

	wg.Wait()

	// Create index file with links to all frameworks
	createFrameworkIndex(outputDir, jsonFiles)

	return nil
}

// scanDirectoryForJSON recursively scans a directory for JSON files
func scanDirectoryForJSON(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".json") {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

// changeExtension changes the extension of a file path
func changeExtension(path, newExt string) string {
	ext := filepath.Ext(path)
	return path[:len(path)-len(ext)] + newExt
}

// convertJSONToMarkdown converts a JSON file to Markdown
func convertJSONToMarkdown(jsonPath, mdPath string) error {
	// Read JSON file
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return fmt.Errorf("open JSON file: %v", err)
	}
	defer jsonFile.Close()

	// Parse JSON
	var doc DocJSONData
	decoder := json.NewDecoder(jsonFile)
	if err := decoder.Decode(&doc); err != nil {
		return fmt.Errorf("decode JSON: %v", err)
	}

	// Create Markdown file
	mdFile, err := os.Create(mdPath)
	if err != nil {
		return fmt.Errorf("create markdown file: %v", err)
	}
	defer mdFile.Close()

	// Generate Markdown content
	if err := writeMarkdownContent(mdFile, &doc); err != nil {
		return fmt.Errorf("write markdown: %v", err)
	}

	return nil
}

// writeMarkdownContent writes the Markdown content to a file
func writeMarkdownContent(w io.Writer, doc *DocJSONData) error {
	// Write title
	title := doc.Metadata.Title
	if title == "" {
		// Try to find title in abstract
		for _, abstract := range doc.Abstract {
			if abstract.Text != "" {
				title = abstract.Text
				break
			}
		}
	}

	// If we still don't have a title, use the file name from the identifier
	if title == "" && doc.Identifier.URL != "" {
		parts := strings.Split(doc.Identifier.URL, "/")
		if len(parts) > 0 {
			title = parts[len(parts)-1]
		}
	}

	fmt.Fprintf(w, "# %s\n\n", title)

	// Write breadcrumb navigation if hierarchy exists
	if len(doc.Hierarchy.Paths) > 0 && len(doc.Hierarchy.Paths[0]) > 1 {
		fmt.Fprintf(w, "**Navigation:** ")
		path := doc.Hierarchy.Paths[0]
		for i, component := range path {
			if i > 0 {
				fmt.Fprintf(w, " › ")
			}
			// Convert component to a more readable format
			readable := strings.ReplaceAll(component, "_", " ")
			readable = strings.Title(readable)
			fmt.Fprintf(w, "%s", readable)
		}
		fmt.Fprintf(w, "\n\n")
	}

	// Write metadata with enhanced styling
	// Only include metadata section if we have actual metadata to show
	hasMetadata := doc.Metadata.RoleHeading != "" || len(doc.Metadata.Modules) > 0 || len(doc.Metadata.Platforms) > 0

	if hasMetadata {
		fmt.Fprintf(w, "---\n\n")

		if doc.Metadata.RoleHeading != "" {
			fmt.Fprintf(w, "**Type:** `%s`\n\n", doc.Metadata.RoleHeading)
		}

		if len(doc.Metadata.Modules) > 0 {
			modules := make([]string, 0, len(doc.Metadata.Modules))
			for _, module := range doc.Metadata.Modules {
				modules = append(modules, fmt.Sprintf("`%s`", module.Name))
			}
			fmt.Fprintf(w, "**Framework:** %s\n\n", strings.Join(modules, ", "))
		}

		if len(doc.Metadata.Platforms) > 0 {
			fmt.Fprintf(w, "**Platform Availability:**\n\n")
			for _, platform := range doc.Metadata.Platforms {
				status := ""
				badge := ""
				if platform.Deprecated {
					status = " (Deprecated)"
					badge = " ⚠️"
				} else if platform.Beta {
					status = " (Beta)"
					badge = " 🧪"
				} else {
					badge = " ✅"
				}
				fmt.Fprintf(w, "- %s%s **%s**%s\n", badge, platform.Name, platform.IntroducedAt, status)
			}
			fmt.Fprintf(w, "\n")
		}

		fmt.Fprintf(w, "---\n\n")
	}

	// Write abstract
	if len(doc.Abstract) > 0 {
		fmt.Fprintf(w, "## Overview\n\n")
		for _, abstract := range doc.Abstract {
			if abstract.Text != "" {
				fmt.Fprintf(w, "%s\n\n", abstract.Text)
			}
		}
	}

	// Write declarations if present
	for _, section := range doc.PrimaryContentSections {
		if section.Kind == "declarations" && len(section.Declarations) > 0 {
			fmt.Fprintf(w, "## Declaration\n\n")

			for i, decl := range section.Declarations {
				// Determine language for code block
				language := "swift"
				languageDisplay := "Swift"
				if len(decl.Languages) > 0 {
					// Use first language in list
					originalLang := decl.Languages[0]
					language = strings.ToLower(originalLang)
					languageDisplay = originalLang
					
					// Map common language names to markdown code block identifiers and display names
					switch language {
					case "swift":
						languageDisplay = "Swift"
					case "objective-c":
						language = "objectivec"
						languageDisplay = "Objective-C"
					case "objective-c++":
						language = "objectivec"
						languageDisplay = "Objective-C++"
					case "c++":
						language = "cpp"
						languageDisplay = "C++"
					default:
						// Capitalize first letter for display
						if len(originalLang) > 0 {
							languageDisplay = strings.Title(strings.ToLower(originalLang))
						}
					}
				}

				// Add language heading for multi-language declarations
				if len(section.Declarations) > 1 {
					fmt.Fprintf(w, "### %s\n\n", languageDisplay)
				}

				fmt.Fprintf(w, "```%s\n", language)

				// Concatenate all tokens to form the declaration
				for _, token := range decl.Tokens {
					fmt.Fprintf(w, "%s", token.Text)
				}
				fmt.Fprintf(w, "\n```\n\n")

				// Add platform information if available
				if len(decl.Platforms) > 0 {
					fmt.Fprintf(w, "*Available on:* %s\n\n", strings.Join(decl.Platforms, ", "))
				}

				// Add separator between declarations if there are multiple
				if i < len(section.Declarations)-1 {
					fmt.Fprintf(w, "---\n\n")
				}
			}
		}
	}

	// Write content sections
	for _, section := range doc.PrimaryContentSections {
		if section.Kind == "content" && len(section.Content) > 0 {
			// Process content blocks
			for _, block := range section.Content {
				writeContentBlock(w, block, doc.References, 0)
			}
		}
	}

	// Write topic sections
	if len(doc.TopicSections) > 0 {
		fmt.Fprintf(w, "## Topics\n\n")

		for _, section := range doc.TopicSections {
			fmt.Fprintf(w, "### %s\n\n", section.Title)

			// Group items by type for better organization
			itemsByType := make(map[string][]Reference)
			ungroupedItems := []Reference{}

			for _, id := range section.Identifiers {
				if ref, ok := doc.References[id]; ok {
					if ref.Role != "" {
						itemsByType[ref.Role] = append(itemsByType[ref.Role], ref)
					} else {
						ungroupedItems = append(ungroupedItems, ref)
					}
				}
			}

			// Write grouped items
			typeOrder := []string{"class", "protocol", "struct", "enum", "function", "method", "property", "type", "symbol"}
			for _, itemType := range typeOrder {
				if items, exists := itemsByType[itemType]; exists {
					if len(itemsByType) > 1 {
						fmt.Fprintf(w, "#### %ss\n\n", strings.Title(itemType))
					}
					for _, ref := range items {
						writeTopicReference(w, ref)
					}
					fmt.Fprintf(w, "\n")
				}
			}

			// Write any remaining types not in our predefined order
			for itemType, items := range itemsByType {
				found := false
				for _, knownType := range typeOrder {
					if itemType == knownType {
						found = true
						break
					}
				}
				if !found && len(items) > 0 {
					if len(itemsByType) > 1 {
						fmt.Fprintf(w, "#### %ss\n\n", strings.Title(itemType))
					}
					for _, ref := range items {
						writeTopicReference(w, ref)
					}
					fmt.Fprintf(w, "\n")
				}
			}

			// Write ungrouped items
			for _, ref := range ungroupedItems {
				writeTopicReference(w, ref)
			}

			if len(ungroupedItems) > 0 {
				fmt.Fprintf(w, "\n")
			}
		}
	}

	// Write relationship sections
	if len(doc.RelationshipsSections) > 0 {
		fmt.Fprintf(w, "## Relationships\n\n")

		for _, section := range doc.RelationshipsSections {
			fmt.Fprintf(w, "### %s\n\n", section.Title)

			for _, id := range section.Identifiers {
				if ref, ok := doc.References[id]; ok {
					// Format URL for markdown compatibility
					url := formatURL(ref.URL)
					fmt.Fprintf(w, "- [%s](%s)\n", ref.Title, url)
				}
			}

			fmt.Fprintf(w, "\n")
		}
	}

	// Write see also sections
	if len(doc.SeeAlsoSections) > 0 {
		fmt.Fprintf(w, "## See Also\n\n")

		for _, section := range doc.SeeAlsoSections {
			fmt.Fprintf(w, "### %s\n\n", section.Title)

			for _, id := range section.Identifiers {
				if ref, ok := doc.References[id]; ok {
					// Get abstract if available
					abstract := ""
					if len(ref.Abstract) > 0 && len(ref.Abstract[0].Text) > 0 {
						// Truncate long abstracts for better readability
						if len(ref.Abstract[0].Text) > 100 {
							abstract = ref.Abstract[0].Text[:97] + "..."
						} else {
							abstract = ref.Abstract[0].Text
						}
					}

					// Format URL for markdown compatibility
					url := formatURL(ref.URL)

					if abstract != "" {
						fmt.Fprintf(w, "- [%s](%s) - %s\n", ref.Title, url, abstract)
					} else {
						fmt.Fprintf(w, "- [%s](%s)\n", ref.Title, url)
					}
				}
			}

			fmt.Fprintf(w, "\n")
		}
	}

	return nil
}

// writeTopicReference writes a single topic reference with enhanced formatting
func writeTopicReference(w io.Writer, ref Reference) {
	// Get abstract if available
	abstract := ""
	if len(ref.Abstract) > 0 && len(ref.Abstract[0].Text) > 0 {
		// Truncate long abstracts for better readability
		if len(ref.Abstract[0].Text) > 120 {
			abstract = ref.Abstract[0].Text[:117] + "..."
		} else {
			abstract = ref.Abstract[0].Text
		}
	}

	// Format URL for markdown compatibility
	url := formatURL(ref.URL)

	// Add type indicator
	typeIndicator := ""
	if ref.Role != "" {
		switch ref.Role {
		case "class":
			typeIndicator = " 🏛️"
		case "protocol":
			typeIndicator = " 📋"
		case "struct":
			typeIndicator = " 🧱"
		case "enum":
			typeIndicator = " 📝"
		case "function":
			typeIndicator = " ⚡"
		case "method":
			typeIndicator = " 🔧"
		case "property":
			typeIndicator = " 📊"
		case "type":
			typeIndicator = " 🏷️"
		default:
			typeIndicator = " 🔗"
		}
	}

	// Add beta indicator
	betaIndicator := ""
	if ref.Beta {
		betaIndicator = " 🧪"
	}

	if abstract != "" {
		fmt.Fprintf(w, "- **[%s](%s)**%s%s  \n  %s\n", ref.Title, url, typeIndicator, betaIndicator, abstract)
	} else {
		fmt.Fprintf(w, "- **[%s](%s)**%s%s\n", ref.Title, url, typeIndicator, betaIndicator)
	}
}

// formatURL ensures URLs are properly formatted for markdown links
func formatURL(urlStr string) string {
	// Handle doc:// scheme URLs
	if strings.HasPrefix(urlStr, "doc://") {
		// Convert doc:// URLs to relative markdown links
		docPath := strings.TrimPrefix(urlStr, "doc://")
		parts := strings.SplitN(docPath, "/", 2)
		if len(parts) > 1 {
			return "/documentation/" + parts[1] + ".md"
		}
		return "#"
	}

	// Handle absolute URLs (keep them as is)
	if strings.HasPrefix(urlStr, "http://") || strings.HasPrefix(urlStr, "https://") {
		return urlStr
	}

	// Handle relative URLs
	if !strings.HasSuffix(urlStr, ".md") && !strings.HasSuffix(urlStr, ".html") {
		// Add .md extension for markdown files
		if strings.HasSuffix(urlStr, ".json") {
			return strings.TrimSuffix(urlStr, ".json") + ".md"
		}
		return urlStr + ".md"
	}

	return urlStr
}

// writeContentBlock writes a content block to the Markdown file
func writeContentBlock(w io.Writer, block ContentBlock, refs map[string]Reference, level int) {
	switch block.Type {
	case "heading":
		headingLevel := block.Level
		// Ensure heading level is between 1-6
		if headingLevel < 1 {
			headingLevel = 1
		} else if headingLevel > 6 {
			headingLevel = 6
		}
		fmt.Fprintf(w, "%s %s\n\n", strings.Repeat("#", headingLevel), block.Text)

	case "paragraph":
		if len(block.InlineContent) > 0 {
			// Process inline content
			for _, inline := range block.InlineContent {
				writeInlineContent(w, inline, refs)
			}
			fmt.Fprintf(w, "\n\n")
		}

	case "unorderedList":
		if len(block.Items) > 0 {
			// Add a newline before the list for proper markdown rendering
			if level == 0 {
				fmt.Fprintf(w, "\n")
			}

			for _, item := range block.Items {
				// Indentation for nested lists
				fmt.Fprintf(w, "%s- ", strings.Repeat("  ", level))

				if len(item.Content) > 0 {
					// For simple paragraph content, write it on the same line
					if len(item.Content) == 1 && item.Content[0].Type == "paragraph" && len(item.Content[0].InlineContent) > 0 {
						for _, inline := range item.Content[0].InlineContent {
							writeInlineContent(w, inline, refs)
						}
						fmt.Fprintf(w, "\n")
					} else {
						// For complex content, write each content block
						fmt.Fprintf(w, "\n")
						for _, content := range item.Content {
							writeContentBlock(w, content, refs, level+1)
						}
					}
				} else {
					fmt.Fprintf(w, "\n")
				}
			}

			// Add an extra newline after the list
			if level == 0 {
				fmt.Fprintf(w, "\n")
			}
		}

	case "orderedList":
		if len(block.Items) > 0 {
			// Add a newline before the list for proper markdown rendering
			if level == 0 {
				fmt.Fprintf(w, "\n")
			}

			for i, item := range block.Items {
				// Indentation for nested lists
				fmt.Fprintf(w, "%s%d. ", strings.Repeat("  ", level), i+1)

				if len(item.Content) > 0 {
					// For simple paragraph content, write it on the same line
					if len(item.Content) == 1 && item.Content[0].Type == "paragraph" && len(item.Content[0].InlineContent) > 0 {
						for _, inline := range item.Content[0].InlineContent {
							writeInlineContent(w, inline, refs)
						}
						fmt.Fprintf(w, "\n")
					} else {
						// For complex content, write each content block
						fmt.Fprintf(w, "\n")
						for _, content := range item.Content {
							writeContentBlock(w, content, refs, level+1)
						}
					}
				} else {
					fmt.Fprintf(w, "\n")
				}
			}

			// Add an extra newline after the list
			if level == 0 {
				fmt.Fprintf(w, "\n")
			}
		}

	case "aside":
		fmt.Fprintf(w, "> **%s**\n>\n", block.Name)
		if len(block.Content) > 0 {
			for _, content := range block.Content {
				// Write each line of content with the blockquote prefix
				var buf bytes.Buffer
				writeContentBlock(&buf, content, refs, level+1)
				lines := strings.Split(buf.String(), "\n")

				for _, line := range lines {
					if line != "" {
						fmt.Fprintf(w, "> %s\n", line)
					} else if len(line) == 0 {
						fmt.Fprintf(w, ">\n")
					}
				}
			}
		}
		fmt.Fprintf(w, "\n")

	case "codeListing":
		// Try to determine the language for syntax highlighting
		language := "swift" // Default to Swift for Apple docs

		// Check if the style field contains language information
		if block.Style != "" {
			// Map common style names to markdown language identifiers
			switch strings.ToLower(block.Style) {
			case "objective-c":
				language = "objectivec"
			case "c++", "cpp":
				language = "cpp"
			case "c":
				language = "c"
			case "java":
				language = "java"
			case "javascript", "js":
				language = "javascript"
			case "shell", "bash", "terminal":
				language = "bash"
			case "python":
				language = "python"
			}
		}

		fmt.Fprintf(w, "```%s\n", language)

		if len(block.Content) > 0 {
			for _, content := range block.Content {
				// For code blocks, we need to handle differently to preserve formatting
				if content.Type == "paragraph" && len(content.InlineContent) > 0 {
					for _, inline := range content.InlineContent {
						if inline.Type == "text" {
							fmt.Fprintf(w, "%s", inline.Text)
						} else if inline.Type == "codeVoice" {
							fmt.Fprintf(w, "%s", inline.Code)
						}
					}
					fmt.Fprintf(w, "\n")
				} else {
					writeContentBlock(w, content, refs, level+1)
				}
			}
		}

		fmt.Fprintf(w, "```\n\n")

	case "table":
		// Handle tables if they appear in the documentation
		if len(block.Content) > 0 {
			fmt.Fprintf(w, "| ")

			// Determine if first row is header
			headerRow := len(block.Content) > 0 && block.Content[0].Type == "tableRow"

			// Write the header row
			if headerRow {
				for _, cell := range block.Content[0].Content {
					if cell.Type == "tableCell" && len(cell.InlineContent) > 0 {
						fmt.Fprintf(w, " ")
						for _, inline := range cell.InlineContent {
							writeInlineContent(w, inline, refs)
						}
						fmt.Fprintf(w, " |")
					}
				}
				fmt.Fprintf(w, "\n|")

				// Write the separator row
				for range block.Content[0].Content {
					fmt.Fprintf(w, " --- |")
				}
				fmt.Fprintf(w, "\n")

				// Write the data rows
				for i := 1; i < len(block.Content); i++ {
					if block.Content[i].Type == "tableRow" {
						fmt.Fprintf(w, "| ")
						for _, cell := range block.Content[i].Content {
							if cell.Type == "tableCell" && len(cell.InlineContent) > 0 {
								fmt.Fprintf(w, " ")
								for _, inline := range cell.InlineContent {
									writeInlineContent(w, inline, refs)
								}
								fmt.Fprintf(w, " |")
							}
						}
						fmt.Fprintf(w, "\n")
					}
				}
			}

			fmt.Fprintf(w, "\n")
		}
	}
}

// writeInlineContent writes inline content to the Markdown file
func writeInlineContent(w io.Writer, inline InlineContent, refs map[string]Reference) {
	switch inline.Type {
	case "text":
		// Handle text with potential escaping for Markdown special characters
		// Escape chars: [, ], (, ), <, >, `, *, _, #, +, -, ., !, |
		text := inline.Text
		// Only escape if not already in code block
		if strings.ContainsAny(text, "[]()<>`*_#+-.!|") {
			// But don't over-escape if already part of Markdown syntax
			// This is a simplified approach - might need tweaking for complex docs
			if !strings.HasPrefix(text, "```") && !strings.HasPrefix(text, "`") {
				// Simple replacements for common characters
				text = strings.ReplaceAll(text, "<", "&lt;")
				text = strings.ReplaceAll(text, ">", "&gt;")
			}
		}
		fmt.Fprintf(w, "%s", text)

	case "codeVoice":
		// Handle inline code
		code := inline.Code
		// Escape backticks in code by using multiple backticks if needed
		if strings.Contains(code, "`") {
			fmt.Fprintf(w, "`` %s ``", code)
		} else {
			fmt.Fprintf(w, "`%s`", code)
		}

	case "reference":
		if ref, ok := refs[inline.Identifier]; ok {
			// Format URL for markdown compatibility
			url := formatURL(ref.URL)

			// Determine if the reference is active
			title := ref.Title
			if title == "" {
				// Try to construct a meaningful title
				if ref.Kind != "" && ref.Role != "" {
					title = fmt.Sprintf("%s %s", ref.Kind, ref.Role)
				} else if len(ref.Fragments) > 0 {
					var titleBuilder strings.Builder
					for _, fragment := range ref.Fragments {
						titleBuilder.WriteString(fragment.Text)
					}
					title = titleBuilder.String()
				} else if len(ref.NavigatorTitle) > 0 {
					var titleBuilder strings.Builder
					for _, navigator := range ref.NavigatorTitle {
						titleBuilder.WriteString(navigator.Text)
					}
					title = titleBuilder.String()
				} else {
					// Last resort - use the identifier
					title = inline.Identifier
				}
			}

			// Beta indicator
			betaIndicator := ""
			if ref.Beta {
				betaIndicator = " (Beta)"
			}

			fmt.Fprintf(w, "[%s%s](%s)", title, betaIndicator, url)
		} else {
			// If reference not found, try to create a meaningful link
			parts := strings.Split(inline.Identifier, "/")
			if len(parts) > 0 {
				lastPart := parts[len(parts)-1]
				// Check if it's likely a doc identifier
				if strings.HasPrefix(inline.Identifier, "doc://") {
					docPath := strings.TrimPrefix(inline.Identifier, "doc://")
					parts := strings.SplitN(docPath, "/", 2)
					if len(parts) > 1 {
						fmt.Fprintf(w, "[%s](/documentation/%s.md)", lastPart, parts[1])
					} else {
						fmt.Fprintf(w, "[%s](#)", lastPart)
					}
				} else {
					fmt.Fprintf(w, "[%s](%s)", lastPart, inline.Identifier)
				}
			} else {
				// If no meaningful parts, just use the identifier
				fmt.Fprintf(w, "[%s](#)", inline.Identifier)
			}
		}

	case "emphasis":
		fmt.Fprintf(w, "*")
		for _, content := range inline.InlineContent {
			writeInlineContent(w, content, refs)
		}
		fmt.Fprintf(w, "*")

	case "strong":
		fmt.Fprintf(w, "**")
		for _, content := range inline.InlineContent {
			writeInlineContent(w, content, refs)
		}
		fmt.Fprintf(w, "**")

	case "image":
		if ref, ok := refs[inline.Identifier]; ok {
			// Format image URL to ensure it works in markdown
			imageURL := ref.URL
			if strings.HasPrefix(imageURL, "/") {
				// For local images, use a relative path or an absolute URL
				// Adjust this based on where the actual images will be hosted
				imageURL = "https://developer.apple.com" + imageURL
			}

			// Add alt text
			altText := ref.Title
			if altText == "" {
				altText = "Image"
				// Try to find a more descriptive alt text
				if len(ref.Abstract) > 0 && ref.Abstract[0].Text != "" {
					altText = ref.Abstract[0].Text
					// Truncate long alt text
					if len(altText) > 50 {
						altText = altText[:47] + "..."
					}
				}
			}

			fmt.Fprintf(w, "![%s](%s)", altText, imageURL)
		} else {
			// If reference not found, create a placeholder
			fmt.Fprintf(w, "![Image](#)")
		}

	case "link":
		// Handle explicit links
		destination := inline.Identifier
		var linkText string

		// Extract link text from inlineContent
		if len(inline.InlineContent) > 0 {
			var textBuilder strings.Builder
			for _, content := range inline.InlineContent {
				if content.Type == "text" {
					textBuilder.WriteString(content.Text)
				}
			}
			linkText = textBuilder.String()
		}

		if linkText == "" {
			linkText = destination
		}

		fmt.Fprintf(w, "[%s](%s)", linkText, destination)
	}
}

// createFrameworkIndex creates an index file with links to all frameworks
func createFrameworkIndex(outputDir string, jsonFiles []string) error {
	// Extract framework names, categories, and classes
	frameworks := make(map[string]bool)
	frameworkClasses := make(map[string][]string) // Classes for each framework
	categories := make(map[string][]string)       // Categories for grouping

	// Known categories to organize frameworks better
	knownCategories := map[string][]string{
		"App Frameworks": {
			"AppKit", "UIKit", "SwiftUI", "Foundation", "Swift", "Objective-C",
			"Combine", "Observation", "GameKit", "MapKit", "WebKit", "CoreData",
		},
		"Graphics & Media": {
			"Metal", "CoreGraphics", "CoreImage", "AVFoundation", "AVKit",
			"QuartzCore", "SceneKit", "SpriteKit", "CoreAnimation", "RealityKit",
			"ARKit", "Vision", "VisionKit", "MediaPlayer", "PhotoKit", "CoreML",
		},
		"System": {
			"Network", "CoreBluetooth", "CoreLocation", "Security", "SystemConfiguration",
			"FileProvider", "BackgroundTasks", "LocalAuthentication", "Intents",
			"UserNotifications", "CloudKit", "CoreMotion", "CoreNFC", "EndpointSecurity",
		},
		"Developer Tools": {
			"XCTest", "Instruments", "Xcode", "XcodeKit", "Xcode-Release-Notes",
			"DeveloperToolsSupport", "Testing",
		},
		"Extended Reality": {
			"visionOS", "ARKit", "RealityKit", "Spatial",
		},
		"Web & Services": {
			"WebKit", "SafariServices", "CloudKit", "CloudKitJS", "MapKitJS",
			"AppleMusicAPI", "AppStoreServerAPI", "WeatherKit", "WeatherKitRESTAPI",
		},
	}

	// First pass: identify frameworks
	for _, file := range jsonFiles {
		// Simple heuristic: look for top-level framework files
		if strings.Contains(file, "/documentation/") && !strings.Contains(filepath.Base(file), "/") {
			frameworkName := filepath.Base(file)
			frameworkName = strings.TrimSuffix(frameworkName, filepath.Ext(frameworkName))
			frameworks[frameworkName] = true

			// Initialize classes array
			frameworkClasses[frameworkName] = []string{}

			// Categorize frameworks
			categorized := false
			for category, categoryItems := range knownCategories {
				for _, item := range categoryItems {
					if strings.EqualFold(item, frameworkName) {
						categories[category] = append(categories[category], frameworkName)
						categorized = true
						break
					}
				}
				if categorized {
					break
				}
			}

			// If not categorized, put in "Other"
			if !categorized {
				categories["Other"] = append(categories["Other"], frameworkName)
			}
		}
	}

	// Second pass: identify classes for each framework
	for _, file := range jsonFiles {
		// Extract path parts to identify classes within frameworks
		parts := strings.Split(file, "/")
		if len(parts) >= 3 {
			// Try to identify if this is a class within a framework
			for framework := range frameworks {
				// Check if this file is a class within this framework
				if strings.Contains(file, "/documentation/"+framework+"/") {
					// Extract class name from file path
					fileName := filepath.Base(file)
					className := strings.TrimSuffix(fileName, filepath.Ext(fileName))

					// Check if it's not just the framework file itself
					if className != framework {
						// Check if we don't already have this class
						hasClass := false
						for _, existingClass := range frameworkClasses[framework] {
							if existingClass == className {
								hasClass = true
								break
							}
						}

						if !hasClass {
							frameworkClasses[framework] = append(frameworkClasses[framework], className)
						}
					}
				}
			}
		}
	}

	// Sort classes for each framework
	for framework := range frameworkClasses {
		sort.Strings(frameworkClasses[framework])
	}

	// Sort category names and framework names within each category
	var categoryList []string
	for category := range categories {
		categoryList = append(categoryList, category)
	}
	sort.Strings(categoryList)

	for category := range categories {
		sort.Strings(categories[category])
	}

	// Create index file
	indexPath := filepath.Join(outputDir, "index.md")
	indexFile, err := os.Create(indexPath)
	if err != nil {
		return fmt.Errorf("create index file: %v", err)
	}
	defer indexFile.Close()

	// Write index content
	fmt.Fprintf(indexFile, "# Apple Documentation\n\n")
	fmt.Fprintf(indexFile, "This is a mirror of Apple's developer documentation converted to Markdown format.\n\n")

	// Table of contents
	fmt.Fprintf(indexFile, "## Contents\n\n")
	for _, category := range categoryList {
		fmt.Fprintf(indexFile, "- [%s](#%s)\n", category, strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(category, " & ", "-"), " ", "-")))
	}
	fmt.Fprintf(indexFile, "\n\n")

	// Write frameworks by category with collapsible sections
	for _, category := range categoryList {
		fmt.Fprintf(indexFile, "## %s\n\n", category)

		frameworksInCategory := categories[category]

		// Create a table with 2 columns for better navigation
		fmt.Fprintf(indexFile, "<table>\n<tr>\n")

		numCols := 2
		numRows := (len(frameworksInCategory) + numCols - 1) / numCols

		// Create columns
		for col := 0; col < numCols; col++ {
			fmt.Fprintf(indexFile, "<td width=\"50%%\">\n\n")

			// Write frameworks for this column with collapsible sections for classes
			for row := 0; row < numRows; row++ {
				idx := row + col*numRows
				if idx < len(frameworksInCategory) {
					framework := frameworksInCategory[idx]
					path := "tutorials/data/documentation/" + framework + ".md"

					// Framework name as a link
					fmt.Fprintf(indexFile, "### [%s](%s)\n\n", framework, path)

					// Add collapsible section with classes if we have any
					classes := frameworkClasses[framework]
					if len(classes) > 0 {
						// GitHub-flavored markdown for collapsible sections
						fmt.Fprintf(indexFile, "<details>\n<summary>Major Classes</summary>\n\n")

						// Only show up to 15 classes to avoid overwhelming lists
						classLimit := 15
						if len(classes) > classLimit {
							// Show a subset with a count of remaining classes
							for i := 0; i < classLimit; i++ {
								className := classes[i]
								classPath := "tutorials/data/documentation/" + framework + "/" + className + ".md"
								fmt.Fprintf(indexFile, "- [%s](%s)\n", className, classPath)
							}
							fmt.Fprintf(indexFile, "- ... and %d more classes\n", len(classes)-classLimit)
						} else {
							// Show all classes
							for _, className := range classes {
								classPath := "tutorials/data/documentation/" + framework + "/" + className + ".md"
								fmt.Fprintf(indexFile, "- [%s](%s)\n", className, classPath)
							}
						}

						fmt.Fprintf(indexFile, "\n</details>\n\n")
					}
				}
			}

			fmt.Fprintf(indexFile, "\n</td>\n")
		}

		fmt.Fprintf(indexFile, "</tr>\n</table>\n\n")
	}

	// Add section specifically for EndpointSecurity if available
	if frameworks["EndpointSecurity"] {
		fmt.Fprintf(indexFile, "## EndpointSecurity Framework Detail\n\n")
		fmt.Fprintf(indexFile, "The EndpointSecurity framework provides system event information to security tools. ")
		fmt.Fprintf(indexFile, "It enables the development of security products that monitor system events, ")
		fmt.Fprintf(indexFile, "ensure system integrity, and contain malicious behavior.\n\n")

		// Link to main documentation
		fmt.Fprintf(indexFile, "**[EndpointSecurity Documentation](tutorials/data/documentation/EndpointSecurity.md)**\n\n")

		// Create collapsible section for classes
		fmt.Fprintf(indexFile, "<details>\n<summary>EndpointSecurity Classes and Types</summary>\n\n")

		classes := frameworkClasses["EndpointSecurity"]
		if len(classes) > 0 {
			for _, className := range classes {
				classPath := "tutorials/data/documentation/EndpointSecurity/" + className + ".md"
				fmt.Fprintf(indexFile, "- [%s](%s)\n", className, classPath)
			}
		} else {
			fmt.Fprintf(indexFile, "- No classes found\n")
		}

		fmt.Fprintf(indexFile, "\n</details>\n\n")
	}

	// Add footer with generation timestamp
	fmt.Fprintf(indexFile, "---\n\n")
	fmt.Fprintf(indexFile, "*Generated on %s*\n", time.Now().Format("January 2, 2006 at 15:04:05 MST"))

	return nil
}
