// generate-framework-bindings generates Go bindings for macOS/iOS frameworks
// from Apple's documentation JSON files.
//
// This generator supports framework-by-framework generation for:
// - ObjectiveC runtime
// - CoreGraphics
// - CoreFoundation
// - AppKit
// - Foundation
// - And other Apple frameworks
package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/tmc/appledocs"
	"github.com/tmc/appledocs/occ2go"
	"golang.org/x/tools/txtar"
)

//go:embed funcs.go templates.txtar templates_*.txtar config.yaml
var embeddedFS embed.FS

var (
	docTemplate               *template.Template
	coreGraphicsTypesTemplate *template.Template
	classesTemplate           *template.Template
	protocolsTemplate         *template.Template
	functionsGenTemplate      *template.Template
	methodsTemplate           *template.Template
)

var templateArchive *txtar.Archive
var variantArchives map[string]*txtar.Archive

// Generator encapsulates the state and methods for generating bindings
type Generator struct {
	Framework        string
	PackageName      string
	InputDir         string
	OutputModule     string
	Variant          string
	WithRefMethods   bool
	GenerateTests    bool
	GenerateExamples bool

	Functions []*occ2go.ParsedFunction
	Classes   []*occ2go.ParsedClass
	Protocols []*occ2go.ParsedProtocol

	// Computed/cached data
	frameworkAbstract string
	frameworkURL      string
	refTypes          []string
	typeMethods       map[string][]*occ2go.ParsedFunction
	typeToRef         map[string]string

	// Error collection
	Errors []error
}

// NewGenerator creates a new Generator instance
func NewGenerator(framework, packageName, inputDir, outputModule, variant string, withRefMethods, generateTests, generateExamples bool) *Generator {
	return &Generator{
		Framework:        framework,
		PackageName:      packageName,
		InputDir:         inputDir,
		OutputModule:     outputModule,
		Variant:          variant,
		WithRefMethods:   withRefMethods,
		GenerateTests:    generateTests,
		GenerateExamples: generateExamples,
		Errors:           make([]error, 0),
	}
}

// AddError adds an error to the error collection
func (g *Generator) AddError(err error) {
	if err != nil {
		g.Errors = append(g.Errors, err)
	}
}

// prepare computes cached data needed for generation
func (g *Generator) prepare() {
	g.frameworkAbstract, g.frameworkURL, _ = loadFrameworkMetadata(g.InputDir, g.Framework)
	g.refTypes = extractRefTypes(g.Functions, getFrameworkPrefix(g.Framework))
	g.typeMethods = groupFunctionsByType(g.Functions, g.Framework)
	g.typeToRef = make(map[string]string)
	for _, refType := range g.refTypes {
		prefix := getFrameworkPrefix(g.Framework)
		if strings.HasPrefix(refType, prefix) && strings.HasSuffix(refType, "Ref") {
			typeName := strings.TrimSuffix(strings.TrimPrefix(refType, prefix), "Ref")
			g.typeToRef[typeName] = refType
		}
	}
}

// SortClassesByDependency sorts classes topologically so parent classes come before children.
// This ensures that when a class extends another class in the same framework, the parent
// is generated first.
func (g *Generator) SortClassesByDependency() {
	if len(g.Classes) == 0 {
		return
	}

	// Build a map of class names for quick lookup
	classMap := make(map[string]*occ2go.ParsedClass)
	for _, cls := range g.Classes {
		classMap[cls.Name] = cls
	}

	// Track visit state: 0 = unvisited, 1 = visiting, 2 = visited
	visited := make(map[string]int)
	sorted := make([]*occ2go.ParsedClass, 0, len(g.Classes))

	// Depth-first search for topological sort
	var visit func(className string) bool
	visit = func(className string) bool {
		if visited[className] == 2 {
			return true // Already processed
		}
		if visited[className] == 1 {
			// Cycle detected - shouldn't happen with proper inheritance, but handle gracefully
			return false
		}

		cls, exists := classMap[className]
		if !exists {
			// Class not in this framework (e.g., NSObject, or from another framework)
			return true
		}

		visited[className] = 1 // Mark as visiting

		// Visit parent first if it exists in this framework
		if cls.SuperClass != "" && cls.SuperClass != "NSObject" {
			if !visit(cls.SuperClass) {
				return false // Cycle detected
			}
		}

		visited[className] = 2 // Mark as visited
		sorted = append(sorted, cls)
		return true
	}

	// Visit all classes
	for _, cls := range g.Classes {
		if visited[cls.Name] == 0 {
			visit(cls.Name)
		}
	}

	// Update the classes with the sorted order
	g.Classes = sorted
}

// GenerateMissingParentStubs creates stub class definitions for parent classes
// that are referenced but not defined in the current framework.
// This handles cases where documentation doesn't include abstract base classes.
func (g *Generator) GenerateMissingParentStubs() []*occ2go.ParsedClass {
	if len(g.Classes) == 0 {
		return nil
	}

	// Build a map of existing classes
	existing := make(map[string]bool)
	for _, cls := range g.Classes {
		existing[cls.Name] = true
	}

	// Find all referenced parent classes that don't exist
	missing := make(map[string]bool)
	for _, cls := range g.Classes {
		if cls.SuperClass != "" && cls.SuperClass != "NSObject" {
			if !existing[cls.SuperClass] {
				// Check if it's a cross-framework type
				resolvedType := resolveType(g.Framework, classToStructName(cls.SuperClass))
				// Only create stub if it's not from another framework
				if !strings.Contains(resolvedType, ".") {
					missing[cls.SuperClass] = true
				}
			}
		}
	}

	// Generate stub classes
	stubs := make([]*occ2go.ParsedClass, 0, len(missing))
	for className := range missing {
		stub := &occ2go.ParsedClass{
			Name:       className,
			SuperClass: "NSObject",
			Methods:    []*occ2go.ParsedMethod{},
			Properties: []*occ2go.ParsedProperty{},
			Comment:    fmt.Sprintf("Auto-generated stub for missing parent class %s", className),
			Abstract:   fmt.Sprintf("A parent class referenced by other %s classes.", g.Framework),
		}
		stubs = append(stubs, stub)
	}

	return stubs
}

// Helper methods for templates

// FunctionCount returns the number of functions
func (g *Generator) FunctionCount() int {
	return len(g.Functions)
}

// ClassCount returns the number of classes
func (g *Generator) ClassCount() int {
	return len(g.Classes)
}

// ProtocolCount returns the number of protocols
func (g *Generator) ProtocolCount() int {
	return len(g.Protocols)
}

// MinVersion returns the minimum macOS version
func (g *Generator) MinVersion() string {
	return findMinimumMacOSVersion(g.Functions)
}

// Abstract returns the framework abstract
func (g *Generator) Abstract() string {
	return g.frameworkAbstract
}

// DocURL returns the framework documentation URL
func (g *Generator) DocURL() string {
	return g.frameworkURL
}

// RefTypes returns the ref types
func (g *Generator) RefTypes() []string {
	return g.refTypes
}

// TypeMethods returns the type methods map
func (g *Generator) TypeMethods() map[string][]*occ2go.ParsedFunction {
	return g.typeMethods
}

// TypeToRef returns the type to ref map
func (g *Generator) TypeToRef() map[string]string {
	return g.typeToRef
}

// Count returns function count (for backward compatibility)
func (g *Generator) Count() int {
	return len(g.Functions)
}

var verbose bool

func init() {
	// Load configuration
	if err := loadConfig(); err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	// Load base templates
	templatesData, err := embeddedFS.ReadFile("templates.txtar")
	if err != nil {
		panic(fmt.Errorf("failed to read templates.txtar: %w", err))
	}
	templateArchive = txtar.Parse(templatesData)

	// Load variant templates
	variantArchives = make(map[string]*txtar.Archive)
	entries, err := fs.ReadDir(embeddedFS, ".")
	if err == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			name := entry.Name()
			// Match templates_<variant>.txtar pattern
			if strings.HasPrefix(name, "templates_") && strings.HasSuffix(name, ".txtar") {
				variantName := strings.TrimSuffix(strings.TrimPrefix(name, "templates_"), ".txtar")
				data, err := embeddedFS.ReadFile(name)
				if err == nil {
					variantArchives[variantName] = txtar.Parse(data)
					if verbose {
						fmt.Fprintf(os.Stderr, "Loaded template variant: %s\n", variantName)
					}
				}
			}
		}
	}

	templates := make(map[string]string)
	for _, file := range templateArchive.Files {
		templates[file.Name] = string(file.Data)
	}

	var parseErr error
	docTemplate, parseErr = template.New("doc.gen.go").Funcs(templateFuncs).Parse(templates["doc.gen.go"])
	if parseErr != nil {
		panic(fmt.Errorf("failed to parse doc.gen.go: %w", parseErr))
	}

	coreGraphicsTypesTemplate, parseErr = template.New("types.gen.go").Funcs(templateFuncs).Parse(templates["types.gen.go"])
	if parseErr != nil {
		panic(fmt.Errorf("failed to parse types.gen.go: %w", parseErr))
	}

	classesTemplate, parseErr = template.New("classes.gen.go").Funcs(templateFuncs).Parse(templates["classes.gen.go"])
	if parseErr != nil {
		panic(fmt.Errorf("failed to parse classes.gen.go: %w", parseErr))
	}

	protocolsTemplate, parseErr = template.New("protocols.gen.go").Funcs(templateFuncs).Parse(templates["protocols.gen.go"])
	if parseErr != nil {
		panic(fmt.Errorf("failed to parse protocols.gen.go: %w", parseErr))
	}

	functionsGenTemplate, parseErr = template.New("functions.gen.go").Funcs(templateFuncs).Parse(templates["functions.gen.go"])
	if parseErr != nil {
		panic(fmt.Errorf("failed to parse functions.gen.go: %w", parseErr))
	}

	methodsTemplate, parseErr = template.New("methods.gen.go").Funcs(templateFuncs).Parse(templates["methods.gen.go"])
	if parseErr != nil {
		panic(fmt.Errorf("failed to parse methods.gen.go: %w", parseErr))
	}
}

// getTemplateVariant returns the template content for a given file, checking variant archives.
// Supports comma-separated variant layering from separate files. For example:
//
//	variant="darwinkit,ref-methods" and filename="types.gen.go" will look for (in order):
//	1. types.gen.go in templates_ref-methods.txtar
//	2. types.gen.go in templates_darwinkit.txtar
//	3. types.gen.go in templates.txtar (fallback)
//
// The first match wins, allowing later variants to override earlier ones.
func getTemplateVariant(filename, variant string) (string, error) {
	// Try variants in reverse order (rightmost first) so later variants override earlier ones
	if variant != "" {
		variants := strings.Split(variant, ",")
		for i := len(variants) - 1; i >= 0; i-- {
			v := strings.TrimSpace(variants[i])
			if v == "" {
				continue
			}
			// Look in variant archive
			if archive, ok := variantArchives[v]; ok {
				for _, file := range archive.Files {
					if file.Name == filename {
						return string(file.Data), nil
					}
				}
			}
		}
	}

	// Fallback to base template
	for _, file := range templateArchive.Files {
		if file.Name == filename {
			return string(file.Data), nil
		}
	}

	return "", fmt.Errorf("template not found: %s", filename)
}

// extractSymbolsFromAPICollections finds all -api.json collection files for a framework
// and creates synthetic documents from their references. This allows processing symbols
// that only exist in API collections without individual documentation files.
// Returns a map of identifier -> synthetic document.
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
			for identifier, ref := range doc.References {
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

	// List all .json files in the input directory
	entries, err := os.ReadDir(inputDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read input directory: %w", err)
	}

	var frameworks []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		// Framework JSON files are named like "Foundation.json" or "AppKit.json"
		if !strings.HasSuffix(name, ".json") {
			continue
		}
		// Extract framework name (remove .json extension)
		frameworkName := strings.TrimSuffix(name, ".json")

		// Check if it matches the regexp
		if re.MatchString(frameworkName) {
			frameworks = append(frameworks, frameworkName)
		}
	}

	sort.Strings(frameworks)
	return frameworks, nil
}

// classbelongsToFramework checks if a class belongs to the specified framework
// by examining the metadata.modules field in the document.
func classbelongsToFramework(doc *appledocs.Document, framework string) bool {
	if doc == nil || doc.Metadata.Modules == nil || len(doc.Metadata.Modules) == 0 {
		// If no module info, default to including it (backward compatibility)
		return true
	}

	// Normalize framework name for comparison (handle variations)
	normalizedFramework := strings.ToLower(strings.ReplaceAll(framework, " ", ""))

	for _, module := range doc.Metadata.Modules {
		normalizedModule := strings.ToLower(strings.ReplaceAll(module.Name, " ", ""))
		if normalizedModule == normalizedFramework {
			return true
		}
	}

	return false
}

// getClassFramework returns the framework name from the document's metadata.
func getClassFramework(doc *appledocs.Document) string {
	if doc == nil || doc.Metadata.Modules == nil || len(doc.Metadata.Modules) == 0 {
		return "unknown"
	}
	return doc.Metadata.Modules[0].Name
}

// generateFramework generates bindings for a single framework.
func generateFramework(framework, inputDir, outputDir, filterRegexp string, txtarOutput bool, variant string, withRefMethods, generateTests, generateExamples bool) error {
	// Check if this is an iOS-only framework and skip if so
	_, _, iOSOnly := loadFrameworkMetadata(inputDir, framework)
	if iOSOnly {
		// Already logged in loadFrameworkMetadata with verbose flag
		return nil // Skip generation silently
	}

	// Open the appledocs filesystem
	fsys, err := appledocs.Open(inputDir)
	if err != nil {
		return fmt.Errorf("failed to open appledocs filesystem: %w", err)
	}

	// Parse all symbols in the framework using the appledocs iterator
	var functions []*occ2go.ParsedFunction
	var classes []*occ2go.ParsedClass
	var protocols []*occ2go.ParsedProtocol

	processedFiles := 0
	parseErrors := 0

	// First, extract synthetic documents from API collection pages
	syntheticDocs := extractSymbolsFromAPICollections(fsys, framework, verbose)

	// Keep track of properties separately so we can attach them to classes later
	classPropertiesMap := make(map[string][]*occ2go.ParsedProperty)

	// Process regular symbols
	for path, doc := range appledocs.Symbols(fsys, framework) {
		processedFiles++
		fn, cls, proto, err := occ2go.ParseDocument(doc)
		if err == nil {
			if fn != nil {
				functions = append(functions, fn)
			}
			if cls != nil {
				// Only include classes that actually belong to this framework
				if classbelongsToFramework(doc, framework) {
					classes = append(classes, cls)
				} else if verbose {
					fmt.Fprintf(os.Stderr, "Skipping class %s (belongs to different framework: %s)\n", cls.Name, getClassFramework(doc))
				}
			}
			if proto != nil {
				protocols = append(protocols, proto)
			}
		} else if strings.Contains(err.Error(), "property (use ParseProperty)") {
			// This is a property file, parse it as a property
			property, propErr := occ2go.ParseProperty(doc)
			if propErr != nil {
				parseErrors++
				if verbose {
					fmt.Fprintf(os.Stderr, "Warning: failed to parse property %s: %v\n", path, propErr)
				}
			} else if property != nil {
				// Extract class name from external ID
				externalID := doc.Metadata.ExternalID
				if strings.Contains(externalID, "(py)") || strings.Contains(externalID, "(cpy)") {
					parts := strings.Split(externalID, "(")
					if len(parts) >= 2 {
						className := strings.TrimPrefix(parts[1], "cs)")
						classPropertiesMap[className] = append(classPropertiesMap[className], property)
					}
				}
			}
		} else {
			parseErrors++
			if verbose {
				fmt.Fprintf(os.Stderr, "Warning: failed to parse %s: %v\n", path, err)
			}
		}
	}

	// Also process synthetic documents from API collection references
	for identifier, doc := range syntheticDocs {
		processedFiles++
		fn, cls, proto, err := occ2go.ParseDocument(doc)
		if err == nil {
			if fn != nil {
				functions = append(functions, fn)
			}
			if cls != nil {
				// Only include classes that actually belong to this framework
				if classbelongsToFramework(doc, framework) {
					classes = append(classes, cls)
				} else if verbose {
					fmt.Fprintf(os.Stderr, "Skipping class %s (belongs to different framework: %s)\n", cls.Name, getClassFramework(doc))
				}
			}
			if proto != nil {
				protocols = append(protocols, proto)
			}
		} else {
			parseErrors++
			if verbose {
				fmt.Fprintf(os.Stderr, "Warning: failed to parse synthetic document %s: %v\n", identifier, err)
			}
		}
	}

	if verbose {
		fmt.Fprintf(os.Stderr, "Processed %d files (%d parse errors)\n", processedFiles, parseErrors)
		fmt.Fprintf(os.Stderr, "Found: %d functions, %d classes, %d protocols\n", len(functions), len(classes), len(protocols))
	}

	// Populate currentFrameworkClasses map for cross-framework type detection
	// This helps resolve whether a type like "Number" or "AccessibilityCustomAction"
	// exists in the current framework or is from another framework
	currentFrameworkClasses = make(map[string]bool)
	for _, cls := range classes {
		// Store both the original name and the stripped name
		strippedName := stripObjCPrefix(cls.Name)
		currentFrameworkClasses[strippedName] = true
	}

	// Second pass: collect methods for each class
	// (Properties are already collected in the first pass)
	if len(classes) > 0 {
		classMethodsMap := make(map[string][]*occ2go.ParsedMethod)
		// Track seen selectors per class to prevent duplicates from documentation
		classSeenSelectors := make(map[string]map[string]bool)

		methodCount := 0
		for _, doc := range appledocs.Symbols(fsys, framework) {
			externalID := doc.Metadata.ExternalID

			// Check for methods: c:objc(cs)ClassName(im)methodName or c:objc(cs)ClassName(cm)methodName
			if strings.Contains(externalID, "(im)") || strings.Contains(externalID, "(cm)") {
				method, err := occ2go.ParseMethod(doc)
				if err == nil && method != nil {
					// Extract class name from external ID
					// Format: c:objc(cs)NSButton(im)initWithFrame:
					parts := strings.Split(externalID, "(")
					if len(parts) >= 2 {
						className := strings.TrimPrefix(parts[1], "cs)")

						// Initialize selector tracking for this class if needed
						if classSeenSelectors[className] == nil {
							classSeenSelectors[className] = make(map[string]bool)
						}

						// Create a unique key combining selector and method type (class vs instance)
						methodKey := method.Selector
						if method.IsClassMethod {
							methodKey = "class:" + methodKey
						} else {
							methodKey = "instance:" + methodKey
						}

						// Debug: Print when we detect duplicates
						if className == "CKRecord" && method.Selector == "creationDate" {
							if classSeenSelectors[className][methodKey] {
								fmt.Fprintf(os.Stderr, "DEBUG: Skipping duplicate method %s.%s from externalID: %s\n", className, method.Selector, externalID)
							} else {
								fmt.Fprintf(os.Stderr, "DEBUG: Adding method %s.%s from externalID: %s\n", className, method.Selector, externalID)
							}
						}

						// Only add if we haven't seen this exact method before
						if !classSeenSelectors[className][methodKey] {
							classMethodsMap[className] = append(classMethodsMap[className], method)
							classSeenSelectors[className][methodKey] = true
							methodCount++
						}
					}
				}
			}
		}

		// Attach methods and properties to classes
		propertyCount := 0
		for i := range classes {
			if methods, ok := classMethodsMap[classes[i].Name]; ok {
				classes[i].Methods = methods
			}
			if properties, ok := classPropertiesMap[classes[i].Name]; ok {
				classes[i].Properties = properties
				propertyCount += len(properties)
			}
		}

		if verbose {
			fmt.Fprintf(os.Stderr, "Found %d methods and %d properties for %d classes\n", methodCount, propertyCount, len(classes))
		}
	}

	// Fail if no symbols were found and no filter was applied
	if len(functions) == 0 && len(classes) == 0 && len(protocols) == 0 && filterRegexp == "" && processedFiles == 0 {
		return fmt.Errorf("no symbols found for framework %s", framework)
	}

	// Deduplicate functions by name (keep first occurrence)
	seenFunctions := make(map[string]bool)
	uniqueFunctions := make([]*occ2go.ParsedFunction, 0, len(functions))
	for _, fn := range functions {
		if !seenFunctions[fn.Name] {
			seenFunctions[fn.Name] = true
			uniqueFunctions = append(uniqueFunctions, fn)
		}
	}
	if len(uniqueFunctions) < len(functions) {
		if verbose {
			fmt.Fprintf(os.Stderr, "Deduplicated functions: %d -> %d\n", len(functions), len(uniqueFunctions))
		}
		functions = uniqueFunctions
	}

	// Apply regexp filter if specified
	if filterRegexp != "" {
		re, err := regexp.Compile(filterRegexp)
		if err != nil {
			return fmt.Errorf("invalid filter regexp: %w", err)
		}

		filteredFunctions := make([]*occ2go.ParsedFunction, 0)
		for _, fn := range functions {
			if re.MatchString(fn.Name) {
				filteredFunctions = append(filteredFunctions, fn)
			}
		}

		filteredClasses := make([]*occ2go.ParsedClass, 0)
		for _, cls := range classes {
			if re.MatchString(cls.Name) {
				filteredClasses = append(filteredClasses, cls)
			}
		}

		filteredProtocols := make([]*occ2go.ParsedProtocol, 0)
		for _, proto := range protocols {
			if re.MatchString(proto.Name) {
				filteredProtocols = append(filteredProtocols, proto)
			}
		}

		if verbose {
			fmt.Fprintf(os.Stderr, "Filter '%s' matched: %d/%d functions, %d/%d classes, %d/%d protocols\n",
				filterRegexp,
				len(filteredFunctions), len(functions),
				len(filteredClasses), len(classes),
				len(filteredProtocols), len(protocols))
		}

		functions = filteredFunctions
		classes = filteredClasses
		protocols = filteredProtocols
	}

	// Create output directory
	packageName := strings.ToLower(framework)

	// Check if outputDir already ends with the package name to avoid double-nesting
	// (e.g., when running from generated/appkit/ with -output .)
	outDir := outputDir
	if filepath.Base(outputDir) != packageName {
		outDir = filepath.Join(outputDir, packageName)
	}

	if !txtarOutput {
		if err := os.MkdirAll(outDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	// Generate bindings
	if txtarOutput {
		if err := generateTxtar(os.Stdout, framework, packageName, inputDir, functions, classes, protocols, withRefMethods, generateTests, generateExamples, variant); err != nil {
			return fmt.Errorf("failed to generate bindings: %w", err)
		}
	} else {
		if err := generateFiles(outDir, framework, packageName, inputDir, functions, classes, protocols, withRefMethods, generateTests, generateExamples, variant); err != nil {
			return fmt.Errorf("failed to generate bindings: %w", err)
		}
		fmt.Printf("Generated %s bindings in %s\n", framework, outDir)
	}

	return nil
}

func main() {
	framework := flag.String("framework", "CoreGraphics", "Framework to generate bindings for (supports regexp patterns like 'Core.*' or '^(AppKit|Foundation)$')")
	inputDir := flag.String("input", "", "Input directory with JSON files (defaults to ~/.appledocs/cache/developer.apple.com/tutorials/data/documentation)")
	outputDir := flag.String("output", "generated", "Output directory for generated bindings")
	filterRegexp := flag.String("filter", "", "Only generate symbols matching this regexp (e.g., '^CGRect' or '^NS(Window|View)')")
	txtarOutput := flag.Bool("txtar", false, "Output as txtar format to stdout instead of files")
	variant := flag.String("variant", "", "Comma-separated template variants (e.g., 'darwinkit' or 'base,ref-methods'). Later variants override earlier ones.")
	withRefMethods := flag.Bool("with-ref-methods", false, "Generate struct-wrapped Ref types with methods (enables method-style API)")
	generateTests := flag.Bool("generate-tests", false, "Generate test files for the bindings")
	generateExamples := flag.Bool("generate-examples", false, "Generate example code demonstrating API usage")
	generateObjcRuntime := flag.Bool("generate-objc-runtime", false, "Generate only the objc runtime package (framework-independent)")
	verboseFlag := flag.Bool("v", false, "Enable verbose output")
	flag.Parse()

	verbose = *verboseFlag

	// Handle objc runtime generation
	if *generateObjcRuntime {
		objcDir := filepath.Join(*outputDir, "objc")
		if err := generateObjcRuntimePackage(objcDir); err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to generate objc runtime: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Generated objc runtime package in %s\n", objcDir)
		return
	}

	// Default to cache directory if not specified
	if *inputDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to get home directory: %v\n", err)
			os.Exit(1)
		}
		*inputDir = filepath.Join(homeDir, ".appledocs/cache/developer.apple.com/tutorials/data/documentation")
	}

	// Discover available frameworks by looking for .json files in the input directory
	frameworks, err := discoverFrameworks(*inputDir, *framework)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to discover frameworks: %v\n", err)
		os.Exit(1)
	}

	if len(frameworks) == 0 {
		fmt.Fprintf(os.Stderr, "Error: no frameworks found matching pattern '%s'\n", *framework)
		os.Exit(1)
	}

	if verbose {
		if len(frameworks) == 1 {
			fmt.Fprintf(os.Stderr, "Generating bindings for %s\n", frameworks[0])
		} else {
			fmt.Fprintf(os.Stderr, "Generating bindings for %d frameworks matching '%s': %v\n", len(frameworks), *framework, frameworks)
		}
		fmt.Fprintf(os.Stderr, "Input: %s\n", *inputDir)
		fmt.Fprintf(os.Stderr, "Output: %s\n", *outputDir)
	}

	// Build cross-framework type registry for proper type resolution
	if err := buildCrossFrameworkTypeRegistry(*outputDir); err != nil {
		// Non-fatal: warn but continue (registry just won't be populated)
		if verbose {
			fmt.Fprintf(os.Stderr, "Warning: failed to build cross-framework type registry: %v\n", err)
		}
	} else if verbose && len(crossFrameworkTypeRegistry) > 0 {
		fmt.Fprintf(os.Stderr, "Built cross-framework type registry with %d types\n", len(crossFrameworkTypeRegistry))
	}

	// Generate bindings for each matching framework
	for _, fw := range frameworks {
		if err := generateFramework(fw, *inputDir, *outputDir, *filterRegexp, *txtarOutput, *variant, *withRefMethods, *generateTests, *generateExamples); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating %s: %v\n", fw, err)
			os.Exit(1)
		}
	}
}

// generateFiles generates all files to disk
func generateFiles(outDir, framework, packageName, inputDir string, functions []*occ2go.ParsedFunction, classes []*occ2go.ParsedClass, protocols []*occ2go.ParsedProtocol, withRefMethods, generateTests, generateExamples bool, variant string) error {
	// Determine output module - default to github.com/tmc/appledocs/generated for now
	outputModule := "github.com/tmc/appledocs/generated"

	// Create generator to get gen.go data
	gen := NewGenerator(framework, packageName, inputDir, outputModule, variant, withRefMethods, generateTests, generateExamples)
	gen.Functions = functions
	gen.Classes = classes
	gen.Protocols = protocols
	gen.prepare()

	// Generate stubs for missing parent classes
	stubs := gen.GenerateMissingParentStubs()
	if len(stubs) > 0 {
		gen.Classes = append(stubs, gen.Classes...)
	}

	// Try to use the module template if it exists
	if _, err := getTemplateVariant("module", variant); err == nil {
		// Generate using module template and parse txtar output
		var buf bytes.Buffer
		if err := gen.GenerateTxtarFromModule(&buf); err != nil {
			return err
		}

		// Parse txtar output
		archive := txtar.Parse(buf.Bytes())

		// Write each file from the archive
		for _, file := range archive.Files {
			if file.Name == "" {
				continue
			}
			filePath := filepath.Join(outDir, file.Name)
			// Create parent directory if needed
			if dir := filepath.Dir(filePath); dir != "." {
				if err := os.MkdirAll(dir, 0755); err != nil {
					return fmt.Errorf("failed to create directory %s: %w", dir, err)
				}
			}
			if err := os.WriteFile(filePath, file.Data, 0644); err != nil {
				return fmt.Errorf("failed to write %s: %w", file.Name, err)
			}
		}

		// Log any collected errors/warnings
		if verbose {
			for _, e := range gen.Errors {
				fmt.Fprintf(os.Stderr, "Warning: %v\n", e)
			}
		}
		return nil
	}

	// Fallback to individual file generation
	generators := []struct {
		filename string
		generate func(io.Writer) error
	}{
		{"gen.go", func(w io.Writer) error {
			templateContent, err := getTemplateVariant("gen.go", variant)
			if err != nil {
				return err
			}
			tmpl, err := template.New("gen.go").Funcs(templateFuncs).Parse(templateContent)
			if err != nil {
				return err
			}
			return tmpl.Execute(w, gen)
		}},
		{"doc.gen.go", func(w io.Writer) error { return generateDoc(w, framework, packageName, inputDir, functions, variant) }},
		{"types.gen.go", func(w io.Writer) error {
			return generateTypes(w, framework, packageName, functions, withRefMethods, variant)
		}},
		{"functions.gen.go", func(w io.Writer) error {
			return generateFunctions(w, framework, packageName, functions, withRefMethods, variant)
		}},
	}

	if withRefMethods {
		generators = append(generators, struct {
			filename string
			generate func(io.Writer) error
		}{"methods.gen.go", func(w io.Writer) error { return generateMethods(w, framework, packageName, functions, variant) }})
	}

	if len(classes) > 0 {
		generators = append(generators, struct {
			filename string
			generate func(io.Writer) error
		}{"classes.gen.go", func(w io.Writer) error { return generateClasses(w, framework, packageName, classes, variant) }})
	}

	if len(protocols) > 0 {
		generators = append(generators, struct {
			filename string
			generate func(io.Writer) error
		}{"protocols.gen.go", func(w io.Writer) error { return generateProtocols(w, framework, packageName, protocols, variant) }})
	}

	if generateTests {
		generators = append(generators, struct {
			filename string
			generate func(io.Writer) error
		}{"functions_test.gen.go", func(w io.Writer) error { return generateTestsFile(w, framework, packageName, functions, variant) }})
	}

	if generateExamples {
		generators = append(generators, struct {
			filename string
			generate func(io.Writer) error
		}{"examples_test.gen.go", func(w io.Writer) error { return generateExamplesFile(w, framework, packageName, functions, variant) }})
	}

	for _, gen := range generators {
		f, err := os.Create(filepath.Join(outDir, gen.filename))
		if err != nil {
			return err
		}
		err = gen.generate(f)
		f.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

// GenerateTxtarFromModule generates the entire txtar output using the module template
func (g *Generator) GenerateTxtarFromModule(w io.Writer) error {
	g.prepare()

	// Load all templates as named templates
	moduleContent, err := getTemplateVariant("module", g.Variant)
	if err != nil {
		return err
	}

	// Create master template and parse all sub-templates
	tmpl := template.New("module").Funcs(templateFuncs)

	// Dynamically discover all templates from the archive (except "module")
	templateFiles := make(map[string]bool)

	// Collect from base archive
	for _, file := range templateArchive.Files {
		if file.Name != "module" && file.Name != "" {
			templateFiles[file.Name] = true
		}
	}

	// Collect from variant archives
	if g.Variant != "" {
		variants := strings.Split(g.Variant, ",")
		for _, v := range variants {
			v = strings.TrimSpace(v)
			if archive, ok := variantArchives[v]; ok {
				for _, file := range archive.Files {
					if file.Name != "module" && file.Name != "" {
						templateFiles[file.Name] = true
					}
				}
			}
		}
	}

	// Parse all discovered templates as associated templates
	for filename := range templateFiles {
		content, err := getTemplateVariant(filename, g.Variant)
		if err != nil {
			g.AddError(fmt.Errorf("warning: skipping template %s: %w", filename, err))
			continue // Skip if template doesn't exist
		}
		_, err = tmpl.New(filename).Parse(content)
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", filename, err)
		}
	}

	// Parse the module template last
	tmpl, err = tmpl.Parse(moduleContent)
	if err != nil {
		return fmt.Errorf("failed to parse module template: %w", err)
	}

	// Execute template - pass Generator directly as context
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, g); err != nil {
		return err
	}

	// Replace #-- with -- to convert to proper txtar format
	output := strings.ReplaceAll(buf.String(), "#-- ", "-- ")
	_, err = w.Write([]byte(output))
	return err
}

// generateTxtar generates all files as txtar format
func generateTxtar(w io.Writer, framework, packageName, inputDir string, functions []*occ2go.ParsedFunction, classes []*occ2go.ParsedClass, protocols []*occ2go.ParsedProtocol, withRefMethods, generateTests, generateExamples bool, variant string) error {
	// Determine output module - default to github.com/tmc/appledocs/generated for now
	outputModule := "github.com/tmc/appledocs/generated"

	// Create generator instance
	gen := NewGenerator(framework, packageName, inputDir, outputModule, variant, withRefMethods, generateTests, generateExamples)
	gen.Functions = functions
	gen.Classes = classes
	gen.Protocols = protocols

	// Try to use the module template if it exists
	if _, err := getTemplateVariant("module", variant); err == nil {
		if err := gen.GenerateTxtarFromModule(w); err != nil {
			return err
		}
		// Log any collected errors/warnings
		if verbose {
			for _, e := range gen.Errors {
				fmt.Fprintf(os.Stderr, "Warning: %v\n", e)
			}
		}
		return nil
	}

	// Fallback to individual file generation
	files := make(map[string][]byte)

	genFile := func(filename string, generator func(io.Writer) error) error {
		var buf bytes.Buffer
		if err := generator(&buf); err != nil {
			return err
		}
		files[filename] = buf.Bytes()
		return nil
	}

	if err := genFile("doc.gen.go", func(w io.Writer) error { return generateDoc(w, framework, packageName, inputDir, functions, variant) }); err != nil {
		return err
	}
	if err := genFile("types.gen.go", func(w io.Writer) error {
		return generateTypes(w, framework, packageName, functions, withRefMethods, variant)
	}); err != nil {
		return err
	}
	if err := genFile("functions.gen.go", func(w io.Writer) error {
		return generateFunctions(w, framework, packageName, functions, withRefMethods, variant)
	}); err != nil {
		return err
	}
	if withRefMethods {
		if err := genFile("methods.gen.go", func(w io.Writer) error { return generateMethods(w, framework, packageName, functions, variant) }); err != nil {
			return err
		}
	}
	if len(classes) > 0 {
		if err := genFile("classes.gen.go", func(w io.Writer) error { return generateClasses(w, framework, packageName, classes, variant) }); err != nil {
			return err
		}
	}
	if len(protocols) > 0 {
		if err := genFile("protocols.gen.go", func(w io.Writer) error { return generateProtocols(w, framework, packageName, protocols, variant) }); err != nil {
			return err
		}
	}

	// Write txtar format
	fmt.Fprintf(w, "# Generated bindings for %s framework\n", framework)
	fmt.Fprintf(w, "# Package: %s\n#\n", packageName)
	fmt.Fprintf(w, "# Functions: %d\n", len(functions))
	fmt.Fprintf(w, "# Classes: %d\n", len(classes))
	fmt.Fprintf(w, "# Protocols: %d\n\n", len(protocols))

	fileNames := make([]string, 0, len(files))
	for name := range files {
		fileNames = append(fileNames, name)
	}
	sort.Strings(fileNames)

	for _, name := range fileNames {
		content := files[name]
		fmt.Fprintf(w, "-- %s --\n", name)
		w.Write(content)
		if len(content) > 0 && content[len(content)-1] != '\n' {
			fmt.Fprintf(w, "\n")
		}
	}

	return nil
}

// generateDoc generates package documentation
func generateDoc(w io.Writer, framework, packageName, inputDir string, functions []*occ2go.ParsedFunction, variant string) error {
	frameworkAbstract, frameworkURL, _ := loadFrameworkMetadata(inputDir, framework)

	data := struct {
		Framework   string
		PackageName string
		MinVersion  string
		Abstract    string
		DocURL      string
	}{
		Framework:   framework,
		PackageName: packageName,
		MinVersion:  findMinimumMacOSVersion(functions),
		Abstract:    frameworkAbstract,
		DocURL:      frameworkURL,
	}

	// Load template with variant support
	templateContent, err := getTemplateVariant("doc.gen.go", variant)
	if err != nil {
		return err
	}
	tmpl, err := template.New("doc.gen.go").Funcs(templateFuncs).Parse(templateContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// generateTypes generates framework-specific type definitions
func generateTypes(w io.Writer, framework, packageName string, functions []*occ2go.ParsedFunction, withRefMethods bool, variant string) error {
	refTypes := extractRefTypes(functions, getFrameworkPrefix(framework))
	data := struct {
		Framework      string
		PackageName    string
		RefTypes       []string
		WithRefMethods bool
	}{framework, packageName, refTypes, withRefMethods}

	// Load template with variant support
	templateContent, err := getTemplateVariant("types.gen.go", variant)
	if err != nil {
		// Fallback: write empty types file if template not found
		fmt.Fprintf(w, "// Code generated from Apple documentation for %s. DO NOT EDIT.\n\n", framework)
		fmt.Fprintf(w, "package %s\n", packageName)
		return nil
	}

	tmpl, err := template.New("types.gen.go").Funcs(templateFuncs).Parse(templateContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// generateFunctions generates function bindings
func generateFunctions(w io.Writer, framework, packageName string, functions []*occ2go.ParsedFunction, withRefMethods bool, variant string) error {
	data := struct {
		Framework      string
		PackageName    string
		Count          int
		Functions      []*occ2go.ParsedFunction
		WithRefMethods bool
	}{framework, packageName, len(functions), functions, withRefMethods}

	// Load template with variant support
	templateContent, err := getTemplateVariant("functions.gen.go", variant)
	if err != nil {
		return err
	}
	tmpl, err := template.New("functions.gen.go").Funcs(templateFuncs).Parse(templateContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// generateClasses generates class declarations
func generateClasses(w io.Writer, framework, packageName string, classes []*occ2go.ParsedClass, variant string) error {
	data := struct {
		Framework   string
		PackageName string
		Count       int
		Classes     []*occ2go.ParsedClass
	}{framework, packageName, len(classes), classes}

	// Load template with variant support
	templateContent, err := getTemplateVariant("classes.gen.go", variant)
	if err != nil {
		return err
	}
	tmpl, err := template.New("classes.gen.go").Funcs(templateFuncs).Parse(templateContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// generateProtocols generates protocol declarations
func generateProtocols(w io.Writer, framework, packageName string, protocols []*occ2go.ParsedProtocol, variant string) error {
	data := struct {
		Framework   string
		PackageName string
		Count       int
		Protocols   []*occ2go.ParsedProtocol
	}{framework, packageName, len(protocols), protocols}

	// Load template with variant support
	templateContent, err := getTemplateVariant("protocols.gen.go", variant)
	if err != nil {
		return err
	}
	tmpl, err := template.New("protocols.gen.go").Funcs(templateFuncs).Parse(templateContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// generateMethods generates method-style wrappers
func generateMethods(w io.Writer, framework, packageName string, functions []*occ2go.ParsedFunction, variant string) error {
	// Group functions by type
	typeMethods := groupFunctionsByType(functions, framework)

	// Build a map from type names to their underlying ref type
	typeToRef := make(map[string]string)
	refTypes := extractRefTypes(functions, getFrameworkPrefix(framework))

	for _, refType := range refTypes {
		// Extract type name from ref type (e.g., CGContextRef -> Context)
		prefix := getFrameworkPrefix(framework)
		if strings.HasPrefix(refType, prefix) && strings.HasSuffix(refType, "Ref") {
			typeName := strings.TrimSuffix(strings.TrimPrefix(refType, prefix), "Ref")
			typeToRef[typeName] = refType
		}
	}

	data := struct {
		Framework   string
		PackageName string
		TypeMethods map[string][]*occ2go.ParsedFunction
		TypeToRef   map[string]string
	}{framework, packageName, typeMethods, typeToRef}

	// Load template with variant support
	templateContent, err := getTemplateVariant("methods.gen.go", variant)
	if err != nil {
		return err
	}
	tmpl, err := template.New("methods.gen.go").Funcs(templateFuncs).Parse(templateContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// getFrameworkPrefix returns the common type prefix for a framework
func getFrameworkPrefix(framework string) string {
	switch framework {
	case "CoreGraphics":
		return "CG"
	case "CoreFoundation":
		return "CF"
	case "CoreAudio":
		return "CA"
	default:
		return ""
	}
}

// extractRefTypes extracts all Ref types from function signatures
func extractRefTypes(functions []*occ2go.ParsedFunction, prefix string) []string {
	refTypesMap := make(map[string]bool)

	for _, fn := range functions {
		if strings.HasPrefix(fn.ReturnType, prefix) && strings.HasSuffix(fn.ReturnType, "Ref") {
			refTypesMap[fn.ReturnType] = true
		}
		for _, param := range fn.Parameters {
			typ := strings.TrimSpace(param.Type)
			if strings.HasPrefix(typ, prefix) && strings.HasSuffix(typ, "Ref") {
				refTypesMap[typ] = true
			}
		}
	}

	refTypes := make([]string, 0, len(refTypesMap))
	for typ := range refTypesMap {
		refTypes = append(refTypes, typ)
	}
	sort.Strings(refTypes)
	return refTypes
}

// loadFrameworkMetadata loads the framework-level JSON to extract abstract, URL, and platform info
func loadFrameworkMetadata(inputDir, framework string) (abstract string, docURL string, iOSOnly bool) {
	frameworkPath := filepath.Join(inputDir, framework+".json")
	data, err := os.ReadFile(frameworkPath)
	if err != nil {
		// Try lowercase
		frameworkPath = filepath.Join(inputDir, strings.ToLower(framework)+".json")
		data, err = os.ReadFile(frameworkPath)
		if err != nil {
			return "", "", false
		}
	}

	var doc struct {
		Abstract []struct {
			Text string `json:"text"`
			Type string `json:"type"`
		} `json:"abstract"`
		Identifier struct {
			URL string `json:"url"`
		} `json:"identifier"`
		Metadata struct {
			Platforms []struct {
				Name         string `json:"name"`
				IntroducedAt string `json:"introducedAt"`
			} `json:"platforms"`
		} `json:"metadata"`
		PrimaryContentSections []struct {
			Kind    string `json:"kind"`
			Content []struct {
				Type    string `json:"type"`
				Style   string `json:"style"`
				Name    string `json:"name"`
				Content []struct {
					Type          string `json:"type"`
					InlineContent []struct {
						Type string `json:"type"`
						Text string `json:"text"`
					} `json:"inlineContent"`
				} `json:"content"`
			} `json:"content"`
		} `json:"primaryContentSections"`
	}

	if err := json.Unmarshal(data, &doc); err != nil {
		return "", "", false
	}

	// Extract abstract text
	for _, item := range doc.Abstract {
		if item.Type == "text" && item.Text != "" {
			abstract = item.Text
			break
		}
	}

	// Check if this is an iOS-only framework (no macOS support)
	hasMacOS := false
	for _, platform := range doc.Metadata.Platforms {
		if platform.Name == "macOS" {
			hasMacOS = true
			break
		}
	}

	if !hasMacOS && len(doc.Metadata.Platforms) > 0 {
		// iOS-only framework - skip generation for macOS bindings
		platformNames := make([]string, 0, len(doc.Metadata.Platforms))
		for _, p := range doc.Metadata.Platforms {
			platformNames = append(platformNames, p.Name)
		}

		if verbose {
			fmt.Fprintf(os.Stderr, "Framework %s is iOS-only (platforms: %s). Skipping generation for macOS.\n",
				framework, strings.Join(platformNames, ", "))
		}

		docURL = occ2go.ConvertDocURLToWeb(doc.Identifier.URL)
		return abstract, docURL, true // iOS-only
	}

	// Extract deprecation information from Important asides
	deprecationText := extractDeprecationNotice(doc.PrimaryContentSections)
	if deprecationText != "" {
		if verbose {
			fmt.Fprintf(os.Stderr, "Framework %s is deprecated: %s\n", framework, deprecationText)
		}
		// Update config with deprecation information
		if config != nil {
			fwConfig := config.Frameworks[framework]
			fwConfig.Deprecated = true
			fwConfig.DeprecationReason = deprecationText
			config.Frameworks[framework] = fwConfig
		}
	}

	docURL = occ2go.ConvertDocURLToWeb(doc.Identifier.URL)
	return abstract, docURL, false // Not iOS-only
}

// extractDeprecationNotice extracts deprecation text from Important asides in content sections
func extractDeprecationNotice(sections []struct {
	Kind    string `json:"kind"`
	Content []struct {
		Type    string `json:"type"`
		Style   string `json:"style"`
		Name    string `json:"name"`
		Content []struct {
			Type          string `json:"type"`
			InlineContent []struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"inlineContent"`
		} `json:"content"`
	} `json:"content"`
}) string {
	for _, section := range sections {
		if section.Kind != "content" {
			continue
		}
		for _, content := range section.Content {
			if content.Type == "aside" && content.Style == "important" {
				// Extract text from inline content
				var texts []string
				for _, para := range content.Content {
					if para.Type == "paragraph" {
						for _, inline := range para.InlineContent {
							if inline.Type == "text" && inline.Text != "" {
								texts = append(texts, inline.Text)
							}
						}
					}
				}
				deprecationText := strings.Join(texts, "")
				// Check if it mentions "Do not use" or deprecation
				if strings.Contains(deprecationText, "Do not use") || strings.Contains(deprecationText, "deprecated") {
					// Clean up the text
					deprecationText = strings.TrimSpace(deprecationText)
					deprecationText = strings.ReplaceAll(deprecationText, "  ", " ")
					return deprecationText
				}
			}
		}
	}
	return ""
}

// findMinimumMacOSVersion finds the minimum macOS version across all functions
func findMinimumMacOSVersion(functions []*occ2go.ParsedFunction) string {
	var minVersion string

	for _, fn := range functions {
		if ver, ok := fn.Availability.IntroducedAt["macOS"]; ok && ver != "" {
			if minVersion == "" || compareVersionStrings(ver, minVersion) < 0 {
				minVersion = ver
			}
		}
	}

	return minVersion
}

// compareVersionStrings compares two version strings semantically
func compareVersionStrings(a, b string) int {
	if a == b {
		return 0
	}

	maj1, min1, ok1 := parseVersion(a)
	maj2, min2, ok2 := parseVersion(b)

	if !ok1 || !ok2 {
		return strings.Compare(a, b)
	}

	if maj1 < maj2 {
		return -1
	}
	if maj1 > maj2 {
		return 1
	}

	if min1 < min2 {
		return -1
	}
	if min1 > min2 {
		return 1
	}

	return 0
}

// parseVersion parses a version string like "10.14" into major and minor components
func parseVersion(s string) (major, minor int, ok bool) {
	parts := strings.SplitN(s, ".", 2)
	if len(parts) == 0 {
		return 0, 0, false
	}

	if _, err := fmt.Sscanf(parts[0], "%d", &major); err != nil {
		return 0, 0, false
	}

	minor = 0
	if len(parts) > 1 {
		fmt.Sscanf(parts[1], "%d", &minor)
	}

	return major, minor, true
}

// generateTestsFile generates test file
func generateTestsFile(w io.Writer, framework, packageName string, functions []*occ2go.ParsedFunction, variant string) error {
	data := struct {
		Framework   string
		PackageName string
		Functions   []*occ2go.ParsedFunction
	}{framework, packageName, functions}

	// Load template with variant support
	templateContent, err := getTemplateVariant("functions_test.gen.go", variant)
	if err != nil {
		return err
	}
	tmpl, err := template.New("functions_test.gen.go").Funcs(templateFuncs).Parse(templateContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// generateExamplesFile generates examples file
func generateExamplesFile(w io.Writer, framework, packageName string, functions []*occ2go.ParsedFunction, variant string) error {
	data := struct {
		Framework   string
		PackageName string
		Functions   []*occ2go.ParsedFunction
	}{framework, packageName, functions}

	// Load template with variant support
	templateContent, err := getTemplateVariant("examples_test.gen.go", variant)
	if err != nil {
		return err
	}
	tmpl, err := template.New("examples_test.gen.go").Funcs(templateFuncs).Parse(templateContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// generateObjcRuntimePackage generates the framework-independent objc runtime package
// by extracting all templates with the objc/ prefix from the template archive.
func generateObjcRuntimePackage(outputDir string) error {
	// Collect all templates that start with "objc/"
	objcFiles := make(map[string][]byte)

	for _, file := range templateArchive.Files {
		if strings.HasPrefix(file.Name, "objc/") {
			// Strip the "objc/" prefix to get the filename
			filename := strings.TrimPrefix(file.Name, "objc/")
			objcFiles[filename] = file.Data
		}
	}

	if len(objcFiles) == 0 {
		return fmt.Errorf("no objc/ templates found in template archive")
	}

	// Create output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Write each file
	for name, data := range objcFiles {
		filePath := filepath.Join(outputDir, name)
		if err := os.WriteFile(filePath, data, 0644); err != nil {
			return fmt.Errorf("failed to write %s: %w", name, err)
		}
		if verbose {
			fmt.Fprintf(os.Stderr, "Generated %s\n", filePath)
		}
	}

	return nil
}
