package main

import (
	"bufio"
	"regexp"
	"strings"

	"golang.org/x/tools/txtar"
)

// filterUnusedUndefinedTypes scans the generated txtar archive and filters
// undefined_types.gen.go to only include types that are actually referenced
// in other generated files. This eliminates unused type aliases that were
// collected but never used (e.g., due to framework hierarchy violations).
func filterUnusedUndefinedTypes(archive *txtar.Archive) *txtar.Archive {
	// Find the undefined_types.gen.go file
	var undefinedTypesIdx = -1
	for i, file := range archive.Files {
		if file.Name == "undefined_types.gen.go" {
			undefinedTypesIdx = i
			break
		}
	}

	// If no undefined_types file, nothing to filter
	if undefinedTypesIdx < 0 {
		return archive
	}

	// Extract all type names defined in undefined_types.gen.go
	undefinedTypes := extractUndefinedTypeNames(string(archive.Files[undefinedTypesIdx].Data))
	if len(undefinedTypes) == 0 {
		return archive
	}

	// Scan all OTHER files for type usage
	usedTypes := make(map[string]bool)
	// Match unqualified type names only (not pkg.Type)
	// Negative lookbehind for '.' is not supported in Go regex, so we'll filter manually
	typePattern := regexp.MustCompile(`\b([A-Z][a-zA-Z0-9]*)\b`)

	for i, file := range archive.Files {
		// Skip undefined_types.gen.go itself
		if i == undefinedTypesIdx {
			continue
		}

		content := string(file.Data)
		// Find all potential type references
		matches := typePattern.FindAllStringIndex(content, -1)
		for _, match := range matches {
			typeName := content[match[0]:match[1]]
			// Skip if this type is preceded by a dot (qualified type like pkg.Type)
			if match[0] > 0 && content[match[0]-1] == '.' {
				continue
			}
			if undefinedTypes[typeName] {
				usedTypes[typeName] = true
			}
		}
	}

	// If no undefined types are used, we can skip the file entirely
	if len(usedTypes) == 0 {
		// Remove the undefined_types.gen.go file from the archive
		newFiles := make([]txtar.File, 0, len(archive.Files)-1)
		for i, file := range archive.Files {
			if i != undefinedTypesIdx {
				newFiles = append(newFiles, file)
			}
		}
		return &txtar.Archive{
			Comment: archive.Comment,
			Files:   newFiles,
		}
	}

	// Regenerate undefined_types.gen.go with only used types
	newContent := regenerateUndefinedTypesFile(
		string(archive.Files[undefinedTypesIdx].Data),
		usedTypes,
	)

	// Update the archive
	archive.Files[undefinedTypesIdx].Data = []byte(newContent)
	return archive
}

// extractUndefinedTypeNames parses undefined_types.gen.go and extracts
// all type alias names defined in it.
func extractUndefinedTypeNames(content string) map[string]bool {
	types := make(map[string]bool)

	// Pattern: "TypeName = _undefined // ..."
	// Matches lines like: "Array = _undefined // referenced in ObjectiveC"
	pattern := regexp.MustCompile(`^\s*([A-Z][a-zA-Z0-9]*)\s*=\s*_undefined`)

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		if matches := pattern.FindStringSubmatch(line); matches != nil {
			types[matches[1]] = true
		}
	}

	return types
}

// regenerateUndefinedTypesFile rebuilds the undefined_types.gen.go content
// with only the types that are actually used.
func regenerateUndefinedTypesFile(originalContent string, usedTypes map[string]bool) string {
	var result strings.Builder

	// Keep header and opening type block
	scanner := bufio.NewScanner(strings.NewReader(originalContent))
	inTypeBlock := false
	typePattern := regexp.MustCompile(`^\s*([A-Z][a-zA-Z0-9]*)\s*=\s*_undefined`)

	for scanner.Scan() {
		line := scanner.Text()

		// Check if we're entering the type block
		if strings.Contains(line, "type (") {
			inTypeBlock = true
			result.WriteString(line)
			result.WriteString("\n")
			continue
		}

		// Check if we're exiting the type block
		if inTypeBlock && strings.TrimSpace(line) == ")" {
			result.WriteString(line)
			result.WriteString("\n")
			inTypeBlock = false
			continue
		}

		// If we're in the type block, filter type definitions
		if inTypeBlock {
			if matches := typePattern.FindStringSubmatch(line); matches != nil {
				typeName := matches[1]
				if usedTypes[typeName] {
					result.WriteString(line)
					result.WriteString("\n")
				}
				// Skip unused types
				continue
			}
		}

		// Keep all other lines (header, comments, etc.)
		result.WriteString(line)
		result.WriteString("\n")
	}

	return result.String()
}
