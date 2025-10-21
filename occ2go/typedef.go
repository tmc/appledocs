package occ2go

import (
	"fmt"
	"strings"

	"github.com/tmc/appledocs"
)

// ParseTypedef parses a typedef declaration from Apple documentation.
// Handles both simple typedefs (typedef int CIFormat) and block typedefs (typedef CGRect (^Callback)(...)).
func ParseTypedef(doc *appledocs.Document) (*ParsedTypedef, error) {
	if doc == nil {
		return nil, fmt.Errorf("document is nil")
	}

	// Get ObjectiveC variant tokens
	tokens := GetObjectiveCVariant(doc)
	if tokens == nil || len(tokens) == 0 {
		return nil, fmt.Errorf("no declaration found for typedef")
	}

	typedef := &ParsedTypedef{
		Availability: ExtractAvailability(doc.Metadata.Platforms),
		DocURL:       ConvertDocURLToWeb(doc.Identifier.URL),
		Abstract:     ExtractAbstract(doc.Abstract),
	}

	// Parse tokens to extract typedef information
	// Expected format: "typedef" <type> <name> ";"
	var tokenText []string
	for _, tok := range tokens {
		if tok.Text != "" {
			tokenText = append(tokenText, tok.Text)
		}
	}

	decl := strings.Join(tokenText, "")

	// Extract the name from the typedef
	// For "typedef int CIFormat;", we want "CIFormat"
	// For "typedef CGRect (^CIKernelROICallback)(int index, CGRect destRect);", we want "CIKernelROICallback"

	if !strings.HasPrefix(decl, "typedef") {
		return nil, fmt.Errorf("not a typedef declaration: %s", decl)
	}

	// Remove typedef keyword and trailing semicolon
	decl = strings.TrimPrefix(decl, "typedef")
	decl = strings.TrimSuffix(decl, ";")
	decl = strings.TrimSpace(decl)

	// Check for block typedef (has ^)
	if strings.Contains(decl, "^") {
		// Block typedef: "CGRect (^CIKernelROICallback)(int index, CGRect destRect)"
		// Find the name between ^ and the parameter list
		startIdx := strings.Index(decl, "^")
		if startIdx == -1 {
			return nil, fmt.Errorf("malformed block typedef")
		}
		endIdx := strings.Index(decl[startIdx:], ")")
		if endIdx == -1 {
			return nil, fmt.Errorf("malformed block typedef")
		}
		typedef.Name = strings.TrimSpace(decl[startIdx+1 : startIdx+endIdx])
		typedef.BaseType = decl // Keep the full declaration
	} else {
		// Simple typedef: "int CIFormat"
		// The last token is the name
		parts := strings.Fields(decl)
		if len(parts) < 2 {
			return nil, fmt.Errorf("malformed typedef: %s", decl)
		}
		typedef.Name = parts[len(parts)-1]
		typedef.BaseType = strings.Join(parts[:len(parts)-1], " ")
	}

	// Check if this is a typed enum by looking for constants in topic sections
	if doc.TopicSections != nil {
		for _, section := range doc.TopicSections {
			// Look for sections with format constants
			if len(section.Identifiers) > 0 {
				typedef.IsTypedEnum = true
				// Extract constant names from identifiers
				// Identifiers are like: "doc://com.apple.coreimage/documentation/CoreImage/CIFormat/ARGB8"
				for _, id := range section.Identifiers {
					// Extract the last component as the constant name
					parts := strings.Split(id, "/")
					if len(parts) > 0 {
						constName := parts[len(parts)-1]
						typedef.Constants = append(typedef.Constants, constName)
					}
				}
			}
		}
	}

	return typedef, nil
}

// ParseConstant parses a constant declaration from Apple documentation.
// Handles extern const declarations (e.g., CORE_IMAGE_EXPORT const CIFormat kCIFormatARGB8).
func ParseConstant(doc *appledocs.Document) (*ParsedConstant, error) {
	if doc == nil {
		return nil, fmt.Errorf("document is nil")
	}

	// Get ObjectiveC variant tokens
	tokens := GetObjectiveCVariant(doc)
	if tokens == nil || len(tokens) == 0 {
		return nil, fmt.Errorf("no declaration found for constant")
	}

	constant := &ParsedConstant{
		Availability: ExtractAvailability(doc.Metadata.Platforms),
		DocURL:       ConvertDocURLToWeb(doc.Identifier.URL),
		Abstract:     ExtractAbstract(doc.Abstract),
	}

	// Extract name and type from tokens
	// Look for identifier tokens that are not keywords
	var identifiers []string
	for _, tok := range tokens {
		if tok.Kind == "identifier" || tok.Kind == "typeIdentifier" {
			identifiers = append(identifiers, tok.Text)
		}
	}

	// For "const CIFormat kCIFormatARGB8", we get ["CIFormat", "kCIFormatARGB8"]
	// Type is first identifier, name is last
	if len(identifiers) >= 2 {
		constant.Type = identifiers[0]
		constant.Name = identifiers[len(identifiers)-1]
	} else if len(identifiers) == 1 {
		// Fall back to using the name from metadata
		constant.Name = doc.Metadata.Title
		constant.Type = identifiers[0]
	} else {
		// Use metadata title as name
		constant.Name = doc.Metadata.Title
		// Try to infer type from context - for now, leave empty
	}

	return constant, nil
}
