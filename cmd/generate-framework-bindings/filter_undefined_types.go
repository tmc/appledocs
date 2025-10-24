package main

import (
	"bufio"
	"regexp"
	"strings"

	"golang.org/x/tools/txtar"
)

// stripGoComments removes single-line (//) and multi-line (/* */) comments from Go code.
// This prevents false positives when scanning for type usage (e.g., "// Set the value" won't match "Set" type).
func stripGoComments(code string) string {
	var result strings.Builder
	inBlockComment := false

	for i := 0; i < len(code); i++ {
		// Check for block comment start
		if !inBlockComment && i+1 < len(code) && code[i] == '/' && code[i+1] == '*' {
			inBlockComment = true
			result.WriteString("  ") // Preserve spacing
			i++ // Skip the '*'
			continue
		}

		// Check for block comment end
		if inBlockComment && i+1 < len(code) && code[i] == '*' && code[i+1] == '/' {
			inBlockComment = false
			result.WriteString("  ") // Preserve spacing
			i++ // Skip the '/'
			continue
		}

		// Inside block comment - replace with space to preserve positions
		if inBlockComment {
			if code[i] == '\n' {
				result.WriteRune('\n')
			} else {
				result.WriteRune(' ')
			}
			continue
		}

		// Check for line comment start
		if i+1 < len(code) && code[i] == '/' && code[i+1] == '/' {
			// Skip until end of line
			for i < len(code) && code[i] != '\n' {
				result.WriteRune(' ')
				i++
			}
			if i < len(code) {
				result.WriteRune('\n')
			}
			continue
		}

		// Normal character
		result.WriteRune(rune(code[i]))
	}

	return result.String()
}

// filterUnusedUndefinedTypes scans the generated txtar archive and filters
// undefined_types.gen.go to only include types that are actually referenced
// in other generated files. This eliminates unused type aliases that were
// collected but never used (e.g., due to framework hierarchy violations).
// It also removes types that are already defined in typedefs.gen.go to prevent duplicates.
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

	// Extract all typedef names to exclude from undefined types
	typedefNames := extractTypedefNames(archive)
	if Debug != nil {
		Debug.Log(DebugUndefined, "Extracted typedef names", nil, "count", len(typedefNames), "undefined_count", len(undefinedTypes))
		Debug.Log(DebugUndefined, "Undefined types list", nil)
		for name := range undefinedTypes {
			Debug.Log(DebugUndefined, "  Undefined type", nil, "name", name)
		}
	}
	// Remove types that are already defined as typedefs
	removedCount := 0
	for typeName := range typedefNames {
		if undefinedTypes[typeName] {
			if Debug != nil {
				Debug.Log(DebugUndefined, "Removing typedef from undefined types", nil, "typename", typeName)
			}
			removedCount++
		}
		delete(undefinedTypes, typeName)
	}
	if Debug != nil {
		Debug.Log(DebugUndefined, "Removed typedefs", nil, "count", removedCount, "remaining", len(undefinedTypes))
	}
	if len(undefinedTypes) == 0 {
		// All undefined types are actually typedefs, remove the file
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

		// Strip comments to avoid false positives (e.g., "// Set the value" matching "Set" type)
		content = stripGoComments(content)

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

// extractTypedefNames parses typedefs.gen.go and extracts all typedef names.
func extractTypedefNames(archive *txtar.Archive) map[string]bool {
	types := make(map[string]bool)

	// Find typedefs.gen.go
	for _, file := range archive.Files {
		if file.Name != "typedefs.gen.go" {
			continue
		}

		// Pattern: "type TypeName" at start of line
		pattern := regexp.MustCompile(`^type\s+([A-Z][a-zA-Z0-9]*)\s+`)

		scanner := bufio.NewScanner(strings.NewReader(string(file.Data)))
		for scanner.Scan() {
			line := scanner.Text()
			if matches := pattern.FindStringSubmatch(line); matches != nil {
				types[matches[1]] = true
			}
		}
		break
	}

	return types
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
