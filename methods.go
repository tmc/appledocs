package appledocs

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"
)

// MethodInfo contains information about a discovered method or property.
type MethodInfo struct {
	// ClassName is the name of the class this method belongs to
	ClassName string

	// Name is the method or property name
	Name string

	// Kind is "method", "property", or "initializer"
	Kind string

	// IsClassMethod indicates if this is a class method (vs instance method)
	IsClassMethod bool

	// ExternalID is the c:objc(...) identifier
	ExternalID string

	// Signature is the parsed method signature from fragments
	Signature string

	// Fragments contains the raw declaration tokens
	Fragments []Fragment

	// Title is the display title from the documentation
	Title string

	// Abstract is the brief description
	Abstract []InlineContent

	// FilePath is the relative path to the JSON file
	FilePath string
}

// ClassMethods groups methods, properties, and initializers by class.
type ClassMethods struct {
	// ClassName is the name of the class
	ClassName string

	// ClassExternalID is the c:objc(cs)ClassName identifier
	ClassExternalID string

	// InstanceMethods contains instance methods
	InstanceMethods []MethodInfo

	// ClassMethods contains class methods
	ClassMethods []MethodInfo

	// Properties contains properties
	Properties []MethodInfo

	// Initializers contains init methods
	Initializers []MethodInfo
}

// externalIDPattern matches the c:objc(...) pattern in externalIDs.
var externalIDPattern = regexp.MustCompile(`^c:objc\((cs|pl)\)([^(]+)(?:\((im|cm|py)\)(.+))?$`)

// parseExternalID parses a c:objc(...) external ID and extracts components.
//
// Patterns:
//   - c:objc(cs)ClassName - class symbol
//   - c:objc(pl)ProtocolName - protocol
//   - c:objc(cs)ClassName(im)methodName - instance method
//   - c:objc(cs)ClassName(cm)methodName - class method
//   - c:objc(cs)ClassName(py)propertyName - property
//
// Returns: (symbolType, className, memberType, memberName, ok)
func parseExternalID(externalID string) (symbolType, className, memberType, memberName string, ok bool) {
	matches := externalIDPattern.FindStringSubmatch(externalID)
	if matches == nil {
		return "", "", "", "", false
	}

	symbolType = matches[1] // "cs" or "pl"
	className = matches[2]

	if len(matches) > 3 && matches[3] != "" {
		memberType = matches[3] // "im", "cm", or "py"
		memberName = matches[4]
	}

	return symbolType, className, memberType, memberName, true
}

// DiscoverMethods scans a filesystem for Apple documentation and discovers all methods.
//
// The fsys parameter should point to the root of cached Apple documentation,
// typically ~/.appledocs/cache/developer.apple.com/tutorials/data/documentation
//
// Files that cannot be parsed are silently skipped. Only symbols with c:objc external IDs
// are included in the results.
//
// Returns a map of className -> ClassMethods.
func DiscoverMethods(fsys *FS) (map[string]*ClassMethods, error) {
	classes := make(map[string]*ClassMethods)

	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Only process JSON files
		if filepath.Ext(path) != ".json" {
			return nil
		}

		// Read and parse the JSON file
		doc, err := fsys.ReadDocument(path)
		if err != nil {
			// Skip files that can't be parsed
			return nil
		}

		// Check if this has a c:objc externalID
		externalID := doc.Metadata.ExternalID
		if externalID == "" || !strings.HasPrefix(externalID, "c:objc") {
			return nil
		}

		// Parse the external ID
		symbolType, className, memberType, memberName, ok := parseExternalID(externalID)
		if !ok {
			return nil
		}

		// Handle class definitions
		if memberType == "" && symbolType == "cs" {
			// This is a class definition
			if classes[className] == nil {
				classes[className] = &ClassMethods{
					ClassName:       className,
					ClassExternalID: externalID,
				}
			}
			return nil
		}

		// Handle methods and properties
		if memberType != "" && symbolType == "cs" {
			// Ensure class entry exists
			if classes[className] == nil {
				classes[className] = &ClassMethods{
					ClassName:       className,
					ClassExternalID: fmt.Sprintf("c:objc(cs)%s", className),
				}
			}

			// Build method signature from fragments
			signature := buildSignature(doc.Metadata.Fragments)

			methodInfo := MethodInfo{
				ClassName:     className,
				Name:          memberName,
				Kind:          doc.Metadata.SymbolKind,
				IsClassMethod: memberType == "cm",
				ExternalID:    externalID,
				Signature:     signature,
				Fragments:     doc.Metadata.Fragments,
				Title:         doc.Metadata.Title,
				Abstract:      doc.Abstract,
				FilePath:      path,
			}

			// Categorize by type
			switch {
			case memberType == "py":
				classes[className].Properties = append(classes[className].Properties, methodInfo)
			case memberType == "cm":
				classes[className].ClassMethods = append(classes[className].ClassMethods, methodInfo)
			case memberType == "im" && strings.HasPrefix(memberName, "init"):
				classes[className].Initializers = append(classes[className].Initializers, methodInfo)
			case memberType == "im":
				classes[className].InstanceMethods = append(classes[className].InstanceMethods, methodInfo)
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("appledocs.DiscoverMethods: walk directory: %w", err)
	}

	return classes, nil
}

// buildSignature constructs a method signature from declaration fragments.
func buildSignature(fragments []Fragment) string {
	var parts []string
	for _, frag := range fragments {
		parts = append(parts, frag.Text)
	}
	return strings.Join(parts, "")
}
