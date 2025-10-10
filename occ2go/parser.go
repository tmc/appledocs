package occ2go

import (
	"fmt"
	"strings"

	"github.com/tmc/appledocs"
)

// ParseDocument parses an appledocs.Document and extracts function/class/protocol definitions.
func ParseDocument(doc *appledocs.Document) (*ParsedFunction, *ParsedClass, *ParsedProtocol, error) {
	if doc == nil {
		return nil, nil, nil, fmt.Errorf("document is nil")
	}

	externalID := doc.Metadata.ExternalID
	availability := ExtractAvailability(doc.Metadata.Platforms)
	docURL := ConvertDocURLToWeb(doc.Identifier.URL)
	abstract := ExtractAbstract(doc.Abstract)

	// Try to get ObjectiveC variant first
	tokens := GetObjectiveCVariant(doc)
	if tokens == nil {
		// Fall back to primary declarations
		for _, section := range doc.PrimaryContentSections {
			if len(section.Declarations) > 0 && len(section.Declarations[0].Tokens) > 0 {
				tokens = section.Declarations[0].Tokens
				break
			}
		}
	}

	if tokens == nil || len(tokens) == 0 {
		return nil, nil, nil, fmt.Errorf("no declaration found for %s", externalID)
	}

	// Determine symbol type by external ID prefix
	switch {
	case strings.HasPrefix(externalID, "c:@F@"):
		// C function
		fn, err := ParseFunctionDeclaration(tokens)
		if err != nil {
			return nil, nil, nil, err
		}
		fn.Availability = availability
		fn.DocURL = docURL
		fn.Abstract = abstract
		return fn, nil, nil, nil

	case strings.HasPrefix(externalID, "c:objc(cs)"):
		// Objective-C class
		cls := ParseClassDeclaration(tokens)
		if cls == nil {
			return nil, nil, nil, fmt.Errorf("failed to parse class declaration")
		}
		cls.Availability = availability
		return nil, cls, nil, nil

	case strings.HasPrefix(externalID, "c:objc(pl)"):
		// Objective-C protocol
		proto := ParseProtocolDeclaration(tokens)
		if proto == nil {
			return nil, nil, nil, fmt.Errorf("failed to parse protocol declaration")
		}
		proto.Availability = availability
		return nil, nil, proto, nil

	case strings.HasPrefix(externalID, "c:@E@"),
		strings.HasPrefix(externalID, "c:@T@"),
		strings.HasPrefix(externalID, "c:@SA@"),
		strings.HasPrefix(externalID, "c:@UA@"),
		strings.HasPrefix(externalID, "c:objc(cy)"),
		strings.HasPrefix(externalID, "s:"),
		strings.HasPrefix(externalID, "doc:"):
		// Known but unsupported types:
		// c:@E@ = enums
		// c:@T@ = typedefs
		// c:@SA@ = struct
		// c:@UA@ = union
		// c:objc(cy) = ObjC category
		// s: = Swift symbols
		// doc: = documentation pages
		return nil, nil, nil, fmt.Errorf("skipping unsupported symbol type: %s", externalID)

	default:
		return nil, nil, nil, fmt.Errorf("unsupported symbol type: %s", externalID)
	}
}

// ExtractAbstract extracts the text description from the abstract field.
func ExtractAbstract(abstract []appledocs.InlineContent) string {
	if len(abstract) == 0 {
		return ""
	}
	var parts []string
	for _, content := range abstract {
		if content.Text != "" {
			parts = append(parts, content.Text)
		}
	}
	return strings.Join(parts, " ")
}

// ConvertDocURLToWeb converts an Apple documentation identifier URL to a web-accessible URL.
func ConvertDocURLToWeb(idURL string) string {
	// Convert: doc://com.apple.documentation/documentation/coregraphics/cgfloat
	// To: https://developer.apple.com/documentation/coregraphics/cgfloat
	if strings.HasPrefix(idURL, "doc://com.apple.documentation/documentation/") {
		path := strings.TrimPrefix(idURL, "doc://com.apple.documentation/documentation/")
		return "https://developer.apple.com/documentation/" + path
	}
	return idURL
}

// ExtractAvailability converts Platform metadata to Availability.
// Returns availability information for all platforms found in the metadata.
func ExtractAvailability(platforms []appledocs.Platform) Availability {
	avail := Availability{
		IntroducedAt: make(map[string]string),
		DeprecatedAt: make(map[string]string),
	}

	for _, p := range platforms {
		if p.Unavailable {
			continue
		}

		if p.IntroducedAt != "" {
			avail.IntroducedAt[p.Name] = p.IntroducedAt
		}

		if p.DeprecatedAt != "" {
			avail.DeprecatedAt[p.Name] = p.DeprecatedAt
		}

		if p.Beta {
			avail.Beta = true
		}
	}

	return avail
}

// GetObjectiveCVariant extracts Objective-C tokens from document variants.
func GetObjectiveCVariant(doc *appledocs.Document) []appledocs.Token {
	for _, variant := range doc.VariantOverrides {
		for _, trait := range variant.Traits {
			if trait.InterfaceLanguage == "occ" {
				for _, patch := range variant.Patch {
					if patch.Op != "replace" {
						continue
					}

					// Check for direct tokens replacement
					if strings.Contains(patch.Path, "declarations/0/tokens") {
						// Need to unmarshal from the interface{}
						if data, ok := patch.Value.([]interface{}); ok {
							tokens := make([]appledocs.Token, 0, len(data))
							for _, item := range data {
								if tokenMap, ok := item.(map[string]interface{}); ok {
									token := appledocs.Token{}
									if kind, ok := tokenMap["kind"].(string); ok {
										token.Kind = kind
									}
									if text, ok := tokenMap["text"].(string); ok {
										token.Text = text
									}
									tokens = append(tokens, token)
								}
							}
							return tokens
						}
					}

					// Check for primaryContentSections replacement (common for functions)
					if patch.Path == "/primaryContentSections/0" {
						if sectionMap, ok := patch.Value.(map[string]interface{}); ok {
							if decls, ok := sectionMap["declarations"].([]interface{}); ok {
								if len(decls) > 0 {
									if declMap, ok := decls[0].(map[string]interface{}); ok {
										if tokens, ok := declMap["tokens"].([]interface{}); ok {
											result := make([]appledocs.Token, 0, len(tokens))
											for _, item := range tokens {
												if tokenMap, ok := item.(map[string]interface{}); ok {
													token := appledocs.Token{}
													if kind, ok := tokenMap["kind"].(string); ok {
														token.Kind = kind
													}
													if text, ok := tokenMap["text"].(string); ok {
														token.Text = text
													}
													result = append(result, token)
												}
											}
											return result
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

// ParseFunctionDeclaration parses a C function declaration from tokens.
func ParseFunctionDeclaration(tokens []appledocs.Token) (*ParsedFunction, error) {
	fn := &ParsedFunction{
		Parameters: []Parameter{},
	}

	// Filter out Swift-only declarations
	for _, tok := range tokens {
		if tok.Kind == "keyword" && (tok.Text == "class" || tok.Text == "static" || tok.Text == "var") {
			return nil, fmt.Errorf("skipping Swift declaration with keyword: %s", tok.Text)
		}
		if tok.Text == "->" {
			return nil, fmt.Errorf("skipping Swift declaration with -> syntax")
		}
	}

	i := 0
	// Skip "extern" keyword
	if i < len(tokens) && tokens[i].Kind == "keyword" && tokens[i].Text == "extern" {
		i++
		for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
			i++
		}
	}

	// Skip additional keywords (const, static, inline, etc)
	for i < len(tokens) && tokens[i].Kind == "keyword" {
		if tokens[i].Text == "const" || tokens[i].Text == "static" ||
			tokens[i].Text == "inline" || tokens[i].Text == "__attribute__" {
			i++
			for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
				i++
			}
		} else {
			break
		}
	}

	// Collect return type
	returnTypeParts := []string{}
	foundFunctionName := false

	for i < len(tokens) && !foundFunctionName {
		if tokens[i].Kind == "identifier" || tokens[i].Kind == "typeIdentifier" {
			// Check if this is the function name
			j := i + 1
			for j < len(tokens) && tokens[j].Kind == "text" && strings.TrimSpace(tokens[j].Text) == "" {
				j++
			}
			if j < len(tokens) && (tokens[j].Text == "(" || strings.HasPrefix(tokens[j].Text, "(")) {
				fn.Name = tokens[i].Text
				// Handle "(" being part of the same token or separate
				if tokens[j].Text == "(" {
					i = j + 1
				} else {
					i = j
				}
				foundFunctionName = true
				break
			}
		}

		if tokens[i].Text != "" && strings.TrimSpace(tokens[i].Text) != "" {
			text := tokens[i].Text
			// Skip attributes and other noise
			if !strings.Contains(text, "__attribute__") &&
			   !strings.Contains(text, "__OSX_AVAILABLE") &&
			   !strings.Contains(text, "API_") {
				returnTypeParts = append(returnTypeParts, text)
			}
		}
		i++
	}

	if !foundFunctionName {
		return nil, fmt.Errorf("failed to parse function name: no opening parenthesis found")
	}

	fn.ReturnType = strings.Join(returnTypeParts, " ")

	// Parse parameters
	for i < len(tokens) {
		for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
			i++
		}

		if i >= len(tokens) || strings.Contains(tokens[i].Text, ")") || strings.Contains(tokens[i].Text, ";") {
			break
		}

		param := Parameter{}
		paramTypeParts := []string{}

		for i < len(tokens) {
			if strings.Contains(tokens[i].Text, ")") || strings.Contains(tokens[i].Text, ";") {
				break
			}
			if tokens[i].Text == "," {
				break
			}

			if tokens[i].Kind == "internalParam" {
				param.Name = tokens[i].Text
				i++
				break
			}

			if tokens[i].Kind == "identifier" {
				j := i + 1
				for j < len(tokens) && tokens[j].Kind == "text" && strings.TrimSpace(tokens[j].Text) == "" {
					j++
				}
				if j < len(tokens) && (tokens[j].Text == "," || strings.Contains(tokens[j].Text, ")") || strings.Contains(tokens[j].Text, ";")) {
					param.Name = tokens[i].Text
					i = j
					break
				}
			}

			text := tokens[i].Text
			if text != "" && strings.TrimSpace(text) != "" &&
				!strings.Contains(text, ";") && !strings.Contains(text, ")") &&
				!strings.Contains(text, "(") && !strings.Contains(text, ",") {
				paramTypeParts = append(paramTypeParts, text)
			}
			i++
		}

		param.Type = strings.Join(paramTypeParts, " ")
		// Don't add void or empty parameters
		if param.Type != "" && param.Type != "void" {
			fn.Parameters = append(fn.Parameters, param)
		}

		if i < len(tokens) && tokens[i].Text == "," {
			i++
		}
	}

	// Skip functions with colons (ObjC selectors)
	if strings.Contains(fn.Name, ":") {
		return nil, fmt.Errorf("skipping ObjC selector: %s", fn.Name)
	}

	if fn.Name == "" {
		return nil, fmt.Errorf("failed to parse function name")
	}

	fn.ReturnType = strings.TrimRight(fn.ReturnType, ";)")
	fn.ReturnType = strings.TrimSpace(fn.ReturnType)
	// Remove trailing function name from return type if it leaked in
	if strings.HasSuffix(fn.ReturnType, " "+fn.Name) {
		fn.ReturnType = strings.TrimSuffix(fn.ReturnType, " "+fn.Name)
		fn.ReturnType = strings.TrimSpace(fn.ReturnType)
	}

	return fn, nil
}

// ParseClassDeclaration parses an Objective-C class declaration from tokens.
func ParseClassDeclaration(tokens []appledocs.Token) *ParsedClass {
	cls := &ParsedClass{}

	// Look for @interface keyword
	interfaceIdx := -1
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Kind == "keyword" && tokens[i].Text == "@interface" {
			interfaceIdx = i
			break
		}
	}

	if interfaceIdx == -1 {
		// No @interface keyword found, try to extract class name from first identifier
		for i := 0; i < len(tokens); i++ {
			if tokens[i].Kind == "identifier" {
				cls.Name = tokens[i].Text
				// Look for superclass after colon
				for j := i + 1; j < len(tokens); j++ {
					if tokens[j].Text == ":" {
						for k := j + 1; k < len(tokens); k++ {
							if tokens[k].Kind == "identifier" || tokens[k].Kind == "typeIdentifier" {
								cls.SuperClass = tokens[k].Text
								break
							}
						}
						break
					}
				}
				break
			}
		}
		if cls.Name == "" {
			return nil
		}
		return cls
	}

	// Found @interface, parse properly
	i := interfaceIdx + 1
	for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
		i++
	}

	if i < len(tokens) && tokens[i].Kind == "identifier" {
		cls.Name = tokens[i].Text
		i++

		// Look for superclass
		for i < len(tokens) {
			if tokens[i].Kind == "text" && strings.Contains(tokens[i].Text, ":") {
				i++
				break
			}
			i++
		}

		// Find superclass name
		for i < len(tokens) {
			if tokens[i].Kind == "identifier" || tokens[i].Kind == "typeIdentifier" {
				cls.SuperClass = tokens[i].Text
				break
			}
			if tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) != "" {
				break
			}
			i++
		}
	}

	if cls.Name == "" {
		return nil
	}

	return cls
}

// ParseProtocolDeclaration parses an Objective-C protocol declaration from tokens.
func ParseProtocolDeclaration(tokens []appledocs.Token) *ParsedProtocol {
	proto := &ParsedProtocol{}

	// Look for @protocol keyword
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Kind == "keyword" && tokens[i].Text == "@protocol" {
			// Skip whitespace
			i++
			for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
				i++
			}

			if i < len(tokens) && tokens[i].Kind == "identifier" {
				proto.Name = tokens[i].Text
				return proto
			}
		}
	}

	// No @protocol keyword found, try to find first identifier
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Kind == "identifier" {
			proto.Name = tokens[i].Text
			break
		}
	}

	if proto.Name == "" {
		return nil
	}

	return proto
}
