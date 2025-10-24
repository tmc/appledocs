package main

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/tmc/appledocs"
	"github.com/tmc/appledocs/occ2go"
)

// extractSymbolsFromAPICollections finds C function declarations in -api.json files.
// These files contain references to symbols that may not have individual JSON files.
func extractSymbolsFromAPICollections(fsys *appledocs.FS, framework string, verbose bool) map[string]*appledocs.Document {
	// Find all -api.json files in the framework directory
	apiFiles := []string{}
	err := fs.WalkDir(fsys, framework, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, "-api.json") {
			apiFiles = append(apiFiles, path)
		}
		return nil
	})
	if err != nil && verbose {
		fmt.Fprintf(os.Stderr, "Warning: failed to walk framework directory for API collections: %v\n", err)
		return nil
	}

	if verbose && len(apiFiles) > 0 {
		fmt.Fprintf(os.Stderr, "Found %d API collection files\n", len(apiFiles))
	}

	// Extract references from each API collection file and create synthetic documents
	syntheticDocs := make(map[string]*appledocs.Document)
	for _, apiFile := range apiFiles {
		doc, err := fsys.ReadDocument(apiFile)
		if err != nil {
			if verbose {
				fmt.Fprintf(os.Stderr, "Warning: failed to read API collection %s: %v\n", apiFile, err)
			}
			continue
		}

		// Process references to create synthetic documents for symbols
		if doc.References != nil {
			// Sort reference keys for stable output
			var refKeys []string
			for identifier := range doc.References {
				refKeys = append(refKeys, identifier)
			}
			sort.Strings(refKeys)

			for _, identifier := range refKeys {
				ref := doc.References[identifier]
				// Only process function symbols (role: "symbol", kind: "symbol")
				if ref.Role != "symbol" || ref.Kind != "symbol" {
					continue
				}

				// Skip if we already have this symbol
				if _, exists := syntheticDocs[identifier]; exists {
					continue
				}

				// Skip if no fragments (can't generate code without them)
				if len(ref.Fragments) == 0 {
					continue
				}

				// Only process C functions - skip Swift/ObjC class/struct declarations
				// C functions start with "func" keyword followed by identifier
				if len(ref.Fragments) > 0 && ref.Fragments[0].Kind == "keyword" {
					keyword := ref.Fragments[0].Text
					// Skip Swift-specific keywords
					if keyword == "class" || keyword == "struct" || keyword == "enum" || keyword == "protocol" || keyword == "typealias" {
						continue
					}
					// Only process "func" and "let" (for constants)
					if keyword != "func" && keyword != "let" {
						continue
					}
				}

				// Determine if this is a C function by checking preciseIdentifier in fragments
				// C functions must have:
				// 1. Either a function identifier with c:@F@ OR all types are C types (c:@T@, c:@...)
				// 2. NO Swift types (s:) in return type or parameters
				isCFunction := false
				hasSwiftTypes := false

				// Check for Swift types (if any typeIdentifier has s: prefix, it's Swift)
				for _, frag := range ref.Fragments {
					if frag.Kind == "typeIdentifier" && frag.PreciseIdentifier != "" {
						if strings.HasPrefix(frag.PreciseIdentifier, "s:") {
							hasSwiftTypes = true
							break
						}
					}
				}
				if hasSwiftTypes {
					continue
				}

				// Now check if it's a C function
				for _, frag := range ref.Fragments {
					if frag.Kind == "identifier" && frag.PreciseIdentifier != "" {
						if strings.HasPrefix(frag.PreciseIdentifier, "c:@F@") {
							isCFunction = true
							break
						}
					}
				}
				// If no identifier with c:@F@ found, check if all types are C types
				if !isCFunction {
					allTypesAreC := true
					foundAnyType := false
					for _, frag := range ref.Fragments {
						if frag.Kind == "typeIdentifier" && frag.PreciseIdentifier != "" {
							foundAnyType = true
							if !strings.HasPrefix(frag.PreciseIdentifier, "c:@") {
								allTypesAreC = false
								break
							}
						}
					}
					if foundAnyType && allTypesAreC {
						isCFunction = true
					}
				}
				if !isCFunction {
					continue
				}

				// Convert Fragments to Tokens for the parser
				// Filter out Swift parameter labels and convert Swift syntax to C-like syntax
				var tokens []appledocs.Token
				skipNextColonText := false
				for i, frag := range ref.Fragments {
					// Skip Swift parameter labels
					if frag.Kind == "externalParam" || frag.Kind == "internalParam" {
						// Skip the next ": " text token that follows parameter labels
						skipNextColonText = true
						continue
					}

					// Skip ": " text tokens that follow parameter labels
					if skipNextColonText && frag.Kind == "text" && strings.HasPrefix(strings.TrimSpace(frag.Text), ":") {
						skipNextColonText = false
						// Keep the comma and space if present (e.g., "?, " becomes ", ")
						if strings.Contains(frag.Text, ",") {
							tokens = append(tokens, appledocs.Token{
								Kind: "text",
								Text: ", ",
							})
						}
						continue
					}
					skipNextColonText = false

					// Remove Swift optional markers (?) from type identifiers
					// They appear in text tokens like "?, " or "?) -> "
					text := frag.Text
					if frag.Kind == "text" && strings.Contains(text, "?") {
						// Replace "?, " with ", " and "?) " with ") "
						text = strings.ReplaceAll(text, "?", "")
					}

					// Skip empty text after processing
					if frag.Kind == "text" && strings.TrimSpace(text) == "" && i > 0 {
						continue
					}

					tokens = append(tokens, appledocs.Token{
						Kind: frag.Kind,
						Text: text,
					})
				}

				// Create synthetic document from reference
				// We need to provide tokens in PrimaryContentSections for the parser
				syntheticDoc := &appledocs.Document{
					Identifier: appledocs.Identifier{
						InterfaceLanguage: "occ",
						URL:               identifier,
					},
					Kind: ref.Kind,
					Metadata: appledocs.Metadata{
						Title:      ref.Title,
						Role:       ref.Role,
						SymbolKind: ref.SymbolKind,
						Fragments:  ref.Fragments,
						ExternalID: extractExternalIDFromFragments(ref.Fragments, ref.Title),
					},
					Abstract: ref.Abstract,
					PrimaryContentSections: []appledocs.ContentSection{
						{
							Kind: "declarations",
							Declarations: []appledocs.Declaration{
								{
									Tokens: tokens,
								},
							},
						},
					},
				}

				syntheticDocs[identifier] = syntheticDoc
			}
		}
	}

	if verbose && len(syntheticDocs) > 0 {
		fmt.Fprintf(os.Stderr, "Created %d synthetic documents from API collection references\n", len(syntheticDocs))
	}

	return syntheticDocs
}

// extractPropertiesFromClassReferences scans class documents for property references
// that don't have separate JSON files. Many properties like NSButton.title are only
// documented in the class's references section, not as standalone files.
// Returns a map of className -> properties.
// propertyFiles is a set of "ClassName.propertyName" strings for properties that have separate files.
func extractPropertiesFromClassReferences(fsys *appledocs.FS, framework string, propertyFiles map[string]bool, verbose bool) map[string][]*occ2go.ParsedProperty {
	classProperties := make(map[string][]*occ2go.ParsedProperty)

	// Process all class documents in the framework
	for _, doc := range appledocs.Symbols(fsys, framework) {
		// Only process class documents (not property files, methods, etc.)
		if !strings.HasPrefix(doc.Metadata.ExternalID, "c:objc(cs)") {
			continue
		}
		// Skip property files: c:objc(cs)NSButton(py)attributedTitle
		if strings.Contains(doc.Metadata.ExternalID, "(py)") || strings.Contains(doc.Metadata.ExternalID, "(cpy)") {
			continue
		}

		// Extract class name from external ID: c:objc(cs)NSButton -> NSButton
		className := strings.TrimPrefix(doc.Metadata.ExternalID, "c:objc(cs)")

		// Scan references for properties
		if doc.References == nil {
			continue
		}

		// Sort reference keys for stable output
		var refKeys []string
		for refKey := range doc.References {
			refKeys = append(refKeys, refKey)
		}
		sort.Strings(refKeys)

		for _, refKey := range refKeys {
			ref := doc.References[refKey]
			// Only process symbol references with fragments
			if ref.Role != "symbol" || ref.Kind != "symbol" || len(ref.Fragments) == 0 {
				continue
			}

			// Check if this is a property (starts with 'var' or 'let' keyword)
			if ref.Fragments[0].Kind != "keyword" {
				continue
			}
			if ref.Fragments[0].Text != "var" && ref.Fragments[0].Text != "let" {
				continue
			}

			// Extract property info from fragments
			// Format: var <name>: <Type>
			// Convert relative URLs to full Apple developer URLs
			docURL := ref.URL
			if strings.HasPrefix(docURL, "/documentation/") {
				docURL = "https://developer.apple.com" + docURL
			}
			property := &occ2go.ParsedProperty{
				Attributes: []string{},
				DocURL:     docURL,
			}

			// 'let' properties are readonly
			if ref.Fragments[0].Text == "let" {
				property.Attributes = append(property.Attributes, "readonly")
			}

			// Extract abstract if available
			if len(ref.Abstract) > 0 {
				for _, abstractNode := range ref.Abstract {
					if abstractNode.Type == "text" && abstractNode.Text != "" {
						property.Abstract = abstractNode.Text
						break
					}
				}
			}

			// Parse property name and type from fragments
			// Expected pattern: keyword(" "), identifier(name), text(": "), typeIdentifier(type)
			// For nested types like NSImage.SymbolConfiguration, we need the last typeIdentifier
			var name, propType, preciseID string
			for i := 0; i < len(ref.Fragments); i++ {
				frag := ref.Fragments[i]
				switch frag.Kind {
				case "identifier":
					// First identifier after 'var'/'let' is the property name
					if name == "" {
						// Strip backticks from property names (Apple's docs use backticks for keywords)
						name = strings.Trim(frag.Text, "`")
					}
				case "typeIdentifier":
					// For nested types (e.g., NSImage.SymbolConfiguration), keep updating
					// until we find the last typeIdentifier with a preciseIdentifier
					if frag.PreciseIdentifier != "" {
						propType = frag.Text
						preciseID = frag.PreciseIdentifier
					} else if propType == "" {
						// Fallback: use first typeIdentifier if no preciseIdentifier found
						propType = frag.Text
					}
				}
			}

			if name == "" {
				continue // Skip if we couldn't extract the name
			}

			// Skip this property if it has a separate JSON file
			// Use lowercase for comparison to handle case variations (e.g., x vs X)
			propertyKey := className + "." + strings.ToLower(name)
			if propertyFiles[propertyKey] {
				// Skip properties that have separate JSON files
				continue
			}

			property.Name = name
			property.Type = propType

			// Map Swift types to ObjC types for proper code generation
			objcType := propType

			// Check if this is an ObjC class type using preciseIdentifier
			if preciseID != "" && strings.HasPrefix(preciseID, "c:objc(cs)") {
				// Extract ObjC class name from c:objc(cs)NSAttributedString -> NSAttributedString *
				className := strings.TrimPrefix(preciseID, "c:objc(cs)")
				objcType = className + " *"
			} else {
				// Map Swift primitive types to ObjC types
				switch propType {
				case "Bool":
					objcType = "BOOL"
				case "String":
					objcType = "NSString *"
				case "Int":
					objcType = "NSInteger"
				case "UInt":
					objcType = "NSUInteger"
				case "Double":
					objcType = "double"
				case "Float":
					objcType = "float"
				case "CGFloat":
					objcType = "CGFloat"
					// For other types, keep the Swift type - the type mapper will handle it
				}
			}
			property.ObjCType = objcType

			classProperties[className] = append(classProperties[className], property)
		}
	}

	if verbose && len(classProperties) > 0 {
		totalProps := 0
		for _, props := range classProperties {
			totalProps += len(props)
		}
		fmt.Fprintf(os.Stderr, "Extracted %d properties from %d class references sections\n", totalProps, len(classProperties))
	}

	return classProperties
}

// extractExternalIDFromFragments attempts to construct an external ID from fragments
// For C functions, this typically looks like: c:@F@FunctionName
func extractExternalIDFromFragments(fragments []appledocs.Fragment, title string) string {
	// Look for the function name in fragments
	for _, frag := range fragments {
		if frag.Kind == "identifier" {
			// Construct C function external ID
			return "c:@F@" + frag.Text
		}
	}
	// Fallback: use title if available
	if title != "" {
		// Remove Swift parameter syntax like (_:) from title
		cleanTitle := strings.TrimSuffix(title, "(_:)")
		cleanTitle = strings.TrimSuffix(cleanTitle, "()")
		return "c:@F@" + cleanTitle
	}
	return ""
}

// discoverFrameworks discovers all frameworks in the input directory that match the given pattern.
// The pattern can be a literal framework name or a regexp pattern.
func discoverFrameworks(inputDir, pattern string) ([]string, error) {
	// Try to compile the pattern as a regexp
	re, err := regexp.Compile(pattern)
	if err != nil {
		// If it's not a valid regexp, treat it as a literal framework name
		return []string{pattern}, nil
	}

	// List all directories in the input directory (each is a framework)
	entries, err := os.ReadDir(inputDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read input directory: %w", err)
	}

	var frameworks []string
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		frameworkName := entry.Name()

		// Check if it matches the regexp
		if re.MatchString(frameworkName) {
			frameworks = append(frameworks, frameworkName)
		}
	}

	sort.Strings(frameworks)
	return frameworks, nil
}

// frameworkModuleNames maps framework directory names to their Apple documentation module names.
// Apple's documentation uses different names than the actual framework names on disk.
var frameworkModuleNames = map[string][]string{
	"QuartzCore": {"Core Animation", "QuartzCore"},
	"AppKit":     {"AppKit", "App Kit"},
	"Foundation": {"Foundation"},
	"ObjectiveC": {"Objective-C Runtime", "ObjectiveC"},
	// Add more mappings as needed
}

// classbelongsToFramework checks if a class belongs to the specified framework
// by examining the metadata.modules field in the document.
func classbelongsToFramework(doc *appledocs.Document, framework string) bool {
	if doc == nil || doc.Metadata.Modules == nil || len(doc.Metadata.Modules) == 0 {
		// If no module info, default to including it (backward compatibility)
		return true
	}

	// Get the possible module names for this framework
	possibleNames := frameworkModuleNames[framework]
	if len(possibleNames) == 0 {
		// If no mapping exists, use the framework name itself
		possibleNames = []string{framework}
	}

	// Normalize and check each possible name
	for _, module := range doc.Metadata.Modules {
		normalizedModule := strings.ToLower(strings.ReplaceAll(module.Name, " ", ""))
		for _, possible := range possibleNames {
			normalizedPossible := strings.ToLower(strings.ReplaceAll(possible, " ", ""))
			if normalizedModule == normalizedPossible {
				return true
			}
		}
	}
	return false
}
