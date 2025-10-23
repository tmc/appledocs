package occ2go

import (
	"fmt"
	"os"
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
	// IMPORTANT: Check for methods and properties BEFORE checking for classes,
	// since method/property IDs also start with "c:objc(cs)"
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

	case strings.HasPrefix(externalID, "c:objc(cs)") && strings.Contains(externalID, "(im)"):
		// Objective-C instance method
		return nil, nil, nil, fmt.Errorf("instance method (use ParseMethod): %s", externalID)

	case strings.HasPrefix(externalID, "c:objc(cs)") && strings.Contains(externalID, "(cm)"):
		// Objective-C class method
		return nil, nil, nil, fmt.Errorf("class method (use ParseMethod): %s", externalID)

	case strings.HasPrefix(externalID, "c:objc(cs)") && (strings.Contains(externalID, "(py)") || strings.Contains(externalID, "(cpy)")):
		// Objective-C property (instance or class property)
		return nil, nil, nil, fmt.Errorf("property (use ParseProperty): %s", externalID)

	case strings.HasPrefix(externalID, "c:objc(cs)"):
		// Objective-C class (must come AFTER method/property checks)
		cls := ParseClassDeclaration(tokens)
		if cls == nil {
			return nil, nil, nil, fmt.Errorf("failed to parse class declaration")
		}
		cls.Availability = availability
		cls.DocURL = docURL
		cls.Abstract = abstract
		cls.Overview = ExtractOverview(doc)

		// Extract superclass from relationshipsSections if not found in tokens
		// This is the authoritative source for inheritance information
		if superClass := ExtractSuperClass(doc); superClass != "" {
			cls.SuperClass = superClass
		}

		return nil, cls, nil, nil

	case strings.HasPrefix(externalID, "c:objc(pl)"):
		// Objective-C protocol
		proto := ParseProtocolDeclaration(tokens)
		if proto == nil {
			return nil, nil, nil, fmt.Errorf("failed to parse protocol declaration")
		}
		proto.Availability = availability
		proto.DocURL = docURL
		proto.Abstract = abstract
		return nil, nil, proto, nil

	case strings.HasPrefix(externalID, "c:@E@"):
		// Objective-C enum type (c:@E@NSWindowStyleMask)
		// Note: Individual enum cases have different external IDs (c:@NSTitledWindowMask)
		// and should be handled by ParseEnumCase instead
		enum := ParseEnumDeclaration(tokens)
		if enum == nil {
			return nil, nil, nil, fmt.Errorf("failed to parse enum declaration")
		}
		enum.Availability = availability
		enum.DocURL = docURL
		enum.Abstract = abstract
		return nil, nil, nil, fmt.Errorf("enum type parsed (use ParseEnum): %s", externalID)

	case strings.HasPrefix(externalID, "c:@T@"),
		strings.HasPrefix(externalID, "c:@SA@"),
		strings.HasPrefix(externalID, "c:@UA@"),
		strings.HasPrefix(externalID, "c:objc(cy)"),
		strings.HasPrefix(externalID, "s:"),
		strings.HasPrefix(externalID, "doc:"):
		// Known but unsupported types:
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

// ExtractOverview extracts the overview description from primaryContentSections.
// It looks for content sections with an "overview" anchor and extracts the paragraph text.
func ExtractOverview(doc *appledocs.Document) string {
	if doc == nil {
		return ""
	}

	for _, section := range doc.PrimaryContentSections {
		if section.Kind != "content" {
			continue
		}

		// Look for overview heading and subsequent paragraphs
		inOverview := false
		var overviewParts []string

		for _, contentItem := range section.Content {
			// Type assert to map to access fields
			contentMap, ok := contentItem.(map[string]interface{})
			if !ok {
				continue
			}

			contentType, _ := contentMap["type"].(string)

			// Check if this is the overview or discussion heading
			if contentType == "heading" {
				anchor, _ := contentMap["anchor"].(string)
				if anchor == "overview" || anchor == "Discussion" {
					inOverview = true
					continue
				}
				// If we're in the overview section and hit another heading, stop
				if inOverview {
					break
				}
			}

			if !inOverview {
				continue
			}

			// Extract paragraph text
			if contentType == "paragraph" {
				if inlineContent, ok := contentMap["inlineContent"].([]interface{}); ok {
					for _, inline := range inlineContent {
						if inlineMap, ok := inline.(map[string]interface{}); ok {
							if text, ok := inlineMap["text"].(string); ok && text != "" {
								overviewParts = append(overviewParts, text)
							}
						}
					}
				}
			}

			// Extract list items
			if contentType == "unorderedList" || contentType == "orderedList" {
				if items, ok := contentMap["items"].([]interface{}); ok {
					for _, item := range items {
						if itemMap, ok := item.(map[string]interface{}); ok {
							if content, ok := itemMap["content"].([]interface{}); ok {
								for _, c := range content {
									if cMap, ok := c.(map[string]interface{}); ok {
										if inlineContent, ok := cMap["inlineContent"].([]interface{}); ok {
											var itemText []string
											for _, inline := range inlineContent {
												if inlineMap, ok := inline.(map[string]interface{}); ok {
													if text, ok := inlineMap["text"].(string); ok && text != "" {
														itemText = append(itemText, text)
													}
												}
											}
											if len(itemText) > 0 {
												overviewParts = append(overviewParts, strings.Join(itemText, " "))
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

		if len(overviewParts) > 0 {
			return strings.Join(overviewParts, " ")
		}
	}

	return ""
}

// ExtractSuperClass extracts the superclass name from relationshipsSections.
// This parses the "inheritsFrom" relationship to determine the parent class.
// Returns the class name without the "NS" prefix (e.g., "NSURLSessionTask" -> "URLSessionTask").
func ExtractSuperClass(doc *appledocs.Document) string {
	if doc == nil {
		return ""
	}

	// Look through relationshipsSections for "inheritsFrom" type
	for _, section := range doc.RelationshipsSections {
		if section.Type == "inheritsFrom" && len(section.Identifiers) > 0 {
			// Extract class name from identifier
			// Format: "doc://com.externally.resolved.symbol/c:objc(cs)NSURLSessionTask"
			identifier := section.Identifiers[0]

			// Look for c:objc(cs) prefix which indicates an Objective-C class
			if idx := strings.Index(identifier, "c:objc(cs)"); idx != -1 {
				className := identifier[idx+len("c:objc(cs)"):]
				// Remove any trailing content after parentheses (method/property markers)
				if parenIdx := strings.Index(className, "("); parenIdx != -1 {
					className = className[:parenIdx]
				}
				return className
			}
		}
	}

	return ""
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

// ParseMethod parses an Objective-C method declaration from a document.
// External IDs:
//   - c:objc(cs)NSButton(im)initWithFrame: (instance method)
//   - c:objc(cs)NSButton(cm)buttonWithTitle:target:action: (class method)
func ParseMethod(doc *appledocs.Document) (*ParsedMethod, error) {
	if doc == nil {
		return nil, fmt.Errorf("document is nil")
	}

	externalID := doc.Metadata.ExternalID
	availability := ExtractAvailability(doc.Metadata.Platforms)
	docURL := ConvertDocURLToWeb(doc.Identifier.URL)
	abstract := ExtractAbstract(doc.Abstract)

	// Determine if it's a class or instance method
	isClassMethod := strings.Contains(externalID, "(cm)")
	if !isClassMethod && !strings.Contains(externalID, "(im)") {
		return nil, fmt.Errorf("not a method: %s", externalID)
	}

	// Get Objective-C variant tokens
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
		return nil, fmt.Errorf("no declaration found for %s", externalID)
	}

	method, err := ParseMethodDeclaration(tokens, isClassMethod)
	if err != nil {
		return nil, err
	}

	method.Availability = availability
	method.DocURL = docURL
	method.Abstract = abstract
	method.IsInitializer = (doc.Metadata.SymbolKind == "init")

	return method, nil
}

// ParseMethodDeclaration parses an Objective-C method declaration from tokens.
// Objective-C method syntax:
//   - (ReturnType)methodName:(Type1)param1 withArg:(Type2)param2
//   - (ReturnType)classMethod:(Type)param
func ParseMethodDeclaration(tokens []appledocs.Token, isClassMethod bool) (*ParsedMethod, error) {
	method := &ParsedMethod{
		IsClassMethod: isClassMethod,
		Parameters:    []Parameter{},
	}

	i := 0

	// Skip leading +/- if present
	if i < len(tokens) && (tokens[i].Text == "+" || tokens[i].Text == "-") {
		i++
		for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
			i++
		}
	}

	// Parse return type in parentheses
	if i < len(tokens) && strings.Contains(tokens[i].Text, "(") {
		// Skip opening paren
		for i < len(tokens) && !strings.Contains(tokens[i].Text, "(") {
			i++
		}
		if i < len(tokens) {
			i++
		}

		// Collect return type until closing paren
		returnTypeParts := []string{}
		preciseID := ""
		for i < len(tokens) && !strings.Contains(tokens[i].Text, ")") {
			if tokens[i].Kind == "typeIdentifier" && tokens[i].PreciseIdentifier != "" {
				preciseID = tokens[i].PreciseIdentifier
			}
			if tokens[i].Text != "" && strings.TrimSpace(tokens[i].Text) != "" {
				returnTypeParts = append(returnTypeParts, tokens[i].Text)
			}
			i++
		}
		// Use preciseIdentifier for enum/typedef if available
		if preciseID != "" && (strings.HasPrefix(preciseID, "c:@E@") || strings.HasPrefix(preciseID, "c:@T@")) {
			parts := strings.Split(preciseID, "@")
			if len(parts) >= 3 {
				returnTypeParts = []string{parts[len(parts)-1]}
			}
		}
		method.ReturnType = strings.Join(returnTypeParts, " ")

		// Skip closing paren
		if i < len(tokens) && strings.Contains(tokens[i].Text, ")") {
			i++
		}

		// Skip whitespace
		for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
			i++
		}
	}

	// Parse selector and parameters
	selectorParts := []string{}
	paramNames := []string{}

	for i < len(tokens) {
		// Look for selector component (identifier followed by colon)
		if i < len(tokens) && tokens[i].Kind == "identifier" {
			selectorPart := tokens[i].Text
			i++

			// Check if the identifier already includes the colon (e.g., "buttonWithTitle:")
			hasColon := strings.HasSuffix(selectorPart, ":")

			if !hasColon {
				// Check if followed by colon as a separate token
				for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
					i++
				}

				if i < len(tokens) && strings.HasPrefix(tokens[i].Text, ":") {
					hasColon = true
					selectorPart += ":"
					i++
				}
			}

			selectorParts = append(selectorParts, selectorPart)

			// If there was a colon, parse the parameter
			if hasColon {
				// Skip whitespace
				for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
					i++
				}

				// Parse parameter type in parentheses
				param := Parameter{}
				if i < len(tokens) && strings.Contains(tokens[i].Text, "(") {
					// Skip opening paren
					for i < len(tokens) && !strings.Contains(tokens[i].Text, "(") {
						i++
					}
					if i < len(tokens) {
						i++
					}

					// Collect parameter type
					// Need to handle nested parentheses for blocks like: void (^)(NSModalResponse)
					paramTypeParts := []string{}
					parenDepth := 0
					preciseID := ""

					for i < len(tokens) {
						text := tokens[i].Text

						// Capture preciseIdentifier for enum/typedef types
						if tokens[i].Kind == "typeIdentifier" && tokens[i].PreciseIdentifier != "" {
							preciseID = tokens[i].PreciseIdentifier
						}

						// Count opening and closing parens in this token
						for _, ch := range text {
							if ch == '(' {
								parenDepth++
							} else if ch == ')' {
								parenDepth--
								// If we go negative, we've hit the closing paren for the parameter type
								if parenDepth < 0 {
									break
								}
							}
						}

						// If we've closed all parens, check if there's content before the closing paren
						if parenDepth < 0 {
							// Extract any content before the closing paren
							if idx := strings.Index(text, ")"); idx > 0 {
								beforeParen := strings.TrimSpace(text[:idx])
								if beforeParen != "" {
									paramTypeParts = append(paramTypeParts, beforeParen)
								}
							}
							break
						}

						// Add token text if non-empty
						if text != "" && strings.TrimSpace(text) != "" {
							paramTypeParts = append(paramTypeParts, text)
						}

						i++
					}

					// Use preciseIdentifier for enum/typedef if available
					if preciseID != "" && (strings.HasPrefix(preciseID, "c:@E@") || strings.HasPrefix(preciseID, "c:@T@")) {
						parts := strings.Split(preciseID, "@")
						if len(parts) >= 3 {
							paramTypeParts = []string{parts[len(parts)-1]}
						}
					}

					param.Type = strings.Join(paramTypeParts, " ")

					// Skip closing paren
					if i < len(tokens) && strings.Contains(tokens[i].Text, ")") {
						i++
					}

					// Skip whitespace
					for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
						i++
					}

					// Get parameter name
					if i < len(tokens) && tokens[i].Kind == "internalParam" {
						param.Name = tokens[i].Text
						i++
					} else if i < len(tokens) && tokens[i].Kind == "identifier" {
						param.Name = tokens[i].Text
						i++
					}

					method.Parameters = append(method.Parameters, param)
					paramNames = append(paramNames, param.Name)
				}
			}

			// Skip whitespace
			for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
				i++
			}

			// Check for end of method declaration
			if i < len(tokens) && (strings.Contains(tokens[i].Text, ";") || tokens[i].Kind == "keyword") {
				break
			}
		} else {
			// Skip non-identifier tokens
			i++
		}

		// Safety check to avoid infinite loops
		if i >= len(tokens) {
			break
		}
	}

	method.Selector = strings.Join(selectorParts, "")

	// Generate Go method name from selector
	method.Name = SelectorToGoName(method.Selector)

	if method.Selector == "" {
		return nil, fmt.Errorf("failed to parse method selector")
	}

	return method, nil
}

// SelectorToGoName converts an Objective-C selector to a Go method name.
// Examples:
//   - "initWithFrame:" -> "InitWithFrame"
//   - "buttonWithTitle:target:action:" -> "ButtonWithTitleTargetAction"
//   - "title" -> "Title"
func SelectorToGoName(selector string) string {
	// Remove trailing colon if present
	selector = strings.TrimSuffix(selector, ":")

	// Split by colon
	parts := strings.Split(selector, ":")

	// Capitalize each part
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}

	return strings.Join(parts, "")
}

// ParseProperty parses an Objective-C property declaration from a document.
// External ID: c:objc(cs)NSButton(py)title
func ParseProperty(doc *appledocs.Document) (*ParsedProperty, error) {
	if doc == nil {
		return nil, fmt.Errorf("document is nil")
	}

	externalID := doc.Metadata.ExternalID
	availability := ExtractAvailability(doc.Metadata.Platforms)
	docURL := ConvertDocURLToWeb(doc.Identifier.URL)
	abstract := ExtractAbstract(doc.Abstract)

	hasPy := strings.Contains(externalID, "(py)")
	hasCpy := strings.Contains(externalID, "(cpy)")
	if os.Getenv("DEBUG_PARSER") != "" {
		fmt.Fprintf(os.Stderr, "DEBUG: Property check %s: (py)=%v (cpy)=%v\n", externalID, hasPy, hasCpy)
	}
	if !hasPy && !hasCpy {
		return nil, fmt.Errorf("not a property: %s", externalID)
	}

	// Get Objective-C variant tokens
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
		return nil, fmt.Errorf("no declaration found for %s", externalID)
	}

	property, err := ParsePropertyDeclaration(tokens)
	if err != nil {
		return nil, err
	}

	// Mark as class property if external ID contains (cpy)
	property.IsClassProperty = strings.Contains(externalID, "(cpy)")

	property.Availability = availability
	property.DocURL = docURL
	property.Abstract = abstract

	return property, nil
}

// parseSwiftPropertyDeclaration parses a Swift property declaration.
// Swift property syntax: [class] var name: Type { get [set] }
func parseSwiftPropertyDeclaration(tokens []appledocs.Token) (*ParsedProperty, error) {
	if os.Getenv("DEBUG_PARSER") != "" {
		fmt.Fprintf(os.Stderr, "DEBUG: parseSwiftPropertyDeclaration called with %d tokens\n", len(tokens))
		if len(tokens) > 0 {
			fmt.Fprintf(os.Stderr, "DEBUG: First token: kind=%s text=%s\n", tokens[0].Kind, tokens[0].Text)
		}
	}

	property := &ParsedProperty{
		Attributes: []string{},
	}

	i := 0

	// Check for 'class' keyword (indicates class property)
	if i < len(tokens) && tokens[i].Kind == "keyword" && tokens[i].Text == "class" {
		// This is a class property, skip to 'var'
		i++
		// Skip whitespace
		for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
			i++
		}
	}

	// Expect 'var' or 'let' keyword
	if i >= len(tokens) || tokens[i].Kind != "keyword" {
		return nil, fmt.Errorf("expected var or let keyword in Swift property")
	}
	if tokens[i].Text != "var" && tokens[i].Text != "let" {
		return nil, fmt.Errorf("expected var or let, got %s", tokens[i].Text)
	}

	// 'let' is read-only
	if tokens[i].Text == "let" {
		property.Attributes = append(property.Attributes, "readonly")
	}

	i++ // Skip 'var' or 'let'

	// Skip whitespace
	for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
		i++
	}

	// Parse property name (identifier)
	if i >= len(tokens) || tokens[i].Kind != "identifier" {
		return nil, fmt.Errorf("expected property name identifier")
	}
	property.Name = tokens[i].Text
	i++

	// Skip whitespace and colon
	for i < len(tokens) && (tokens[i].Kind == "text" || tokens[i].Text == ":") {
		i++
	}

	// Parse property type (typeIdentifier or keyword)
	typeParts := []string{}
	var preciseID string
	for i < len(tokens) && (tokens[i].Kind == "typeIdentifier" || tokens[i].Kind == "keyword") {
		typeParts = append(typeParts, tokens[i].Text)
		// Capture preciseIdentifier from LAST typeIdentifier token
		// For nested types like "CKQueryOperation.Cursor", we want the last one (CKQueryCursor)
		// not the first one (CKQueryOperation)
		if tokens[i].Kind == "typeIdentifier" && tokens[i].PreciseIdentifier != "" {
			preciseID = tokens[i].PreciseIdentifier
		}
		i++
		// Skip whitespace and dots between type parts
		for i < len(tokens) && (tokens[i].Kind == "text" || tokens[i].Text == ".") && strings.TrimSpace(tokens[i].Text) != "" && tokens[i].Text != "{" {
			if tokens[i].Text == "." {
				typeParts = append(typeParts, tokens[i].Text)
			}
			i++
		}
		// Skip pure whitespace
		for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
			i++
		}
	}

	if len(typeParts) == 0 {
		return nil, fmt.Errorf("no type found for property %s", property.Name)
	}

	property.Type = strings.Join(typeParts, " ")

	// Map Swift types to ObjC types for proper code generation
	objcType := property.Type

	// Check if this is an ObjC class type using preciseIdentifier
	if preciseID != "" && strings.HasPrefix(preciseID, "c:objc(cs)") {
		// Extract ObjC class name: c:objc(cs)NSAttributedString -> NSAttributedString *
		className := strings.TrimPrefix(preciseID, "c:objc(cs)")
		objcType = className + " *"
	} else {
		// Map Swift primitive types to ObjC types
		switch property.Type {
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

	// Debug output for property type mapping
	if os.Getenv("DEBUG_PARSER") == "1" {
		fmt.Fprintf(os.Stderr, "DEBUG parseSwiftPropertyDeclaration: %s\n", property.Name)
		fmt.Fprintf(os.Stderr, "  Swift Type: %s\n", property.Type)
		fmt.Fprintf(os.Stderr, "  ObjC Type: %s\n", property.ObjCType)
		fmt.Fprintf(os.Stderr, "  PreciseID: %s\n", preciseID)
	}

	// Check for { get } or { get set } to determine readonly
	// Look for opening brace
	for i < len(tokens) && tokens[i].Text != "{" {
		i++
	}
	if i < len(tokens) {
		i++ // Skip {
		// Look for 'get' and 'set' keywords
		hasGet := false
		hasSet := false
		for i < len(tokens) && tokens[i].Text != "}" {
			if tokens[i].Kind == "keyword" {
				if tokens[i].Text == "get" {
					hasGet = true
				} else if tokens[i].Text == "set" {
					hasSet = true
				}
			}
			i++
		}
		// If only 'get' or neither, it's readonly
		if hasGet && !hasSet {
			if !contains(property.Attributes, "readonly") {
				property.Attributes = append(property.Attributes, "readonly")
			}
		}
	}

	return property, nil
}

// contains checks if a string slice contains a value
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// ParsePropertyDeclaration parses an Objective-C property declaration from tokens.
// Property syntax: @property (attributes) Type name;
func ParsePropertyDeclaration(tokens []appledocs.Token) (*ParsedProperty, error) {
	property := &ParsedProperty{
		Attributes: []string{},
	}

	// Check if this is a Swift property declaration (class var / var / let)
	// Pattern: [@attributes] [class] var name: Type { get }
	// Need to skip attributes to find the property keyword
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Kind == "keyword" {
			if tokens[i].Text == "class" || tokens[i].Text == "var" || tokens[i].Text == "let" {
				return parseSwiftPropertyDeclaration(tokens)
			}
			// If we hit another keyword (like @property), it's not Swift
			if tokens[i].Text == "@property" {
				break
			}
		}
	}

	i := 0

	// Look for @property keyword (Objective-C)
	for i < len(tokens) && !(tokens[i].Kind == "keyword" && tokens[i].Text == "@property") {
		i++
	}

	if i >= len(tokens) {
		return nil, fmt.Errorf("@property keyword not found")
	}

	i++ // Skip @property

	// Skip whitespace
	for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
		i++
	}

	// Parse attributes in parentheses
	if i < len(tokens) && strings.Contains(tokens[i].Text, "(") {
		// Skip opening paren
		for i < len(tokens) && !strings.Contains(tokens[i].Text, "(") {
			i++
		}
		if i < len(tokens) {
			i++
		}

		// Collect attributes until closing paren
		for i < len(tokens) && !strings.Contains(tokens[i].Text, ")") {
			if tokens[i].Kind == "keyword" || tokens[i].Kind == "identifier" {
				property.Attributes = append(property.Attributes, tokens[i].Text)
			}
			i++
		}

		// Skip closing paren
		if i < len(tokens) && strings.Contains(tokens[i].Text, ")") {
			i++
		}
	}

	// Skip whitespace
	for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
		i++
	}

	// Parse property type
	typeParts := []string{}
	for i < len(tokens) && tokens[i].Kind != "identifier" {
		if tokens[i].Kind == "typeIdentifier" || tokens[i].Kind == "keyword" {
			typeParts = append(typeParts, tokens[i].Text)
		} else if tokens[i].Kind == "text" {
			// Include all text tokens that are part of the type
			// This includes generic type parameters like <NSButton *>, pointer markers *, etc.
			// Skip only pure whitespace
			text := strings.TrimSpace(tokens[i].Text)
			if text != "" {
				typeParts = append(typeParts, text)
			}
		}
		i++
	}

	property.Type = strings.Join(typeParts, " ")
	// Store the original Objective-C type for proper type mapping
	property.ObjCType = property.Type

	// Parse property name
	if i < len(tokens) && tokens[i].Kind == "identifier" {
		property.Name = tokens[i].Text
	}

	if property.Name == "" {
		return nil, fmt.Errorf("failed to parse property name")
	}

	return property, nil
}

// ParseEnumDeclaration parses an Objective-C enum type declaration from tokens.
// Enum type syntax: typedef NS_ENUM(NSUInteger, NSWindowStyleMask) { ... }
//                   typedef NS_OPTIONS(NSUInteger, NSEventModifierFlags) { ... }
// Note: This parses the enum type definition, not individual enum cases.
// Individual cases are parsed by ParseEnumCase.
func ParseEnumDeclaration(tokens []appledocs.Token) *ParsedEnum {
	enum := &ParsedEnum{
		Cases: []*ParsedEnumCase{},
	}

	i := 0

	// Try modern syntax first: enum NSBackingStoreType : NSUInteger
	// Look for "enum" keyword
	for i < len(tokens) && !(tokens[i].Kind == "keyword" && tokens[i].Text == "enum") {
		i++
	}

	if i >= len(tokens) {
		// No enum keyword found
		return nil
	}

	i++ // Skip "enum"

	// Skip whitespace
	for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
		i++
	}

	// Get enum name (e.g., NSBackingStoreType)
	if i < len(tokens) && (tokens[i].Kind == "typeIdentifier" || tokens[i].Kind == "identifier") {
		enum.Name = tokens[i].Text
		i++
	} else {
		return nil
	}

	// Skip whitespace
	for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
		i++
	}

	// Look for colon
	if i < len(tokens) && strings.Contains(tokens[i].Text, ":") {
		i++ // Skip colon
	} else {
		// No colon found - might be a simple enum without base type
		return enum
	}

	// Skip whitespace
	for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
		i++
	}

	// Get base type (e.g., NSUInteger, NSInteger, UInt)
	if i < len(tokens) && (tokens[i].Kind == "typeIdentifier" || tokens[i].Kind == "identifier") {
		enum.BaseType = tokens[i].Text

		// Determine if this is an options-style (bitfield) enum
		// NSUInteger typically indicates NS_OPTIONS, NSInteger typically indicates NS_ENUM
		// For now, we'll infer based on enum name patterns (contains "Options", "Mask", etc.)
		lowerName := strings.ToLower(enum.Name)
		if strings.Contains(lowerName, "options") || strings.Contains(lowerName, "mask") {
			enum.IsOptions = true
		}
	}

	return enum
}

// ParseSwiftEnumDeclaration parses a Swift enum/struct declaration.
// Swift option set syntax: struct NSWindowStyleMask : OptionSet
func ParseSwiftEnumDeclaration(tokens []appledocs.Token) *ParsedEnum {
	enum := &ParsedEnum{
		Cases:    []*ParsedEnumCase{},
		BaseType: "UInt", // Swift default for OptionSet
	}

	i := 0

	// Look for struct keyword
	for i < len(tokens) && !(tokens[i].Kind == "keyword" && tokens[i].Text == "struct") {
		i++
	}

	if i >= len(tokens) {
		return nil
	}

	i++ // Skip struct

	// Skip whitespace
	for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
		i++
	}

	// Get enum name
	if i < len(tokens) && tokens[i].Kind == "identifier" {
		enum.Name = tokens[i].Text
		i++
	} else {
		return nil
	}

	// Check for : OptionSet to determine if it's a bitfield
	for i < len(tokens) {
		if tokens[i].Text == ":" {
			i++
			for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
				i++
			}
			if i < len(tokens) && tokens[i].Kind == "typeIdentifier" && tokens[i].Text == "OptionSet" {
				enum.IsOptions = true
			}
			break
		}
		i++
	}

	return enum
}

// ParseEnumCase parses an individual enum constant/case from a document.
// External ID format: c:@NSTitledWindowMask (without c:@E@ prefix)
// Token pattern: static const NSWindowStyleMask NSTitledWindowMask;
//           or: static var titled: NSWindow.StyleMask { get }
func ParseEnumCase(doc *appledocs.Document) (*ParsedEnumCase, error) {
	if doc == nil {
		return nil, fmt.Errorf("document is nil")
	}

	externalID := doc.Metadata.ExternalID
	availability := ExtractAvailability(doc.Metadata.Platforms)
	docURL := ConvertDocURLToWeb(doc.Identifier.URL)
	abstract := ExtractAbstract(doc.Abstract)

	// If abstract is empty, try to extract from Discussion section
	if abstract == "" {
		abstract = ExtractOverview(doc)
	}

	// Modern enum cases have external IDs like: c:@E@NSWindowCollectionBehavior@NSWindowCollectionBehaviorDefault
	// Enum type definitions have 3 parts: c:@E@EnumName
	// Enum cases have 4+ parts: c:@E@EnumName@CaseName
	if strings.HasPrefix(externalID, "c:@E@") {
		parts := strings.Split(externalID, "@")
		if len(parts) == 3 {
			// This is an enum type, not a case
			return nil, fmt.Errorf("enum type, not case: %s", externalID)
		}
		// If len(parts) >= 4, this is an enum case, continue parsing
	}

	// Get Objective-C variant tokens
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
		return nil, fmt.Errorf("no declaration found for %s", externalID)
	}

	enumCase, err := ParseEnumCaseDeclaration(tokens)
	if err != nil {
		return nil, err
	}

	enumCase.Availability = availability
	enumCase.DocURL = docURL
	enumCase.Abstract = abstract

	return enumCase, nil
}

// ParseEnumCaseDeclaration parses an enum constant declaration from tokens.
// Modern format (ObjC variant in docs): NSBackingStoreBuffered
// Old format: static const NSWindowStyleMask NSTitledWindowMask;
// Swift format: static var titled: NSWindow.StyleMask { get }
func ParseEnumCaseDeclaration(tokens []appledocs.Token) (*ParsedEnumCase, error) {
	enumCase := &ParsedEnumCase{}

	i := 0

	// Modern Apple docs just have a single identifier token for ObjC enum cases
	// Try this first before the old formats
	if len(tokens) == 1 && tokens[0].Kind == "identifier" {
		enumCase.Name = tokens[0].Text
		return enumCase, nil
	}

	// Skip static keyword if present
	if i < len(tokens) && tokens[i].Kind == "keyword" && tokens[i].Text == "static" {
		i++
		for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
			i++
		}
	}

	// Check for const (Objective-C) or var (Swift)
	if i < len(tokens) && tokens[i].Kind == "keyword" {
		if tokens[i].Text == "const" {
			// Objective-C: static const TypeName CaseName;
			i++

			// Skip whitespace
			for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
				i++
			}

			// Skip type name
			if i < len(tokens) && (tokens[i].Kind == "typeIdentifier" || tokens[i].Kind == "identifier") {
				i++
			}

			// Skip whitespace
			for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
				i++
			}

			// Get case name
			if i < len(tokens) && tokens[i].Kind == "identifier" {
				enumCase.Name = tokens[i].Text
			}

		} else if tokens[i].Text == "var" {
			// Swift: static var caseName: TypeName { get }
			i++

			// Skip whitespace
			for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
				i++
			}

			// Get case name
			if i < len(tokens) && tokens[i].Kind == "identifier" {
				enumCase.Name = tokens[i].Text
			}
		}
	} else if i < len(tokens) && (tokens[i].Kind == "identifier" || tokens[i].Kind == "typeIdentifier") {
		// Fallback: just an identifier (modern Apple docs format)
		enumCase.Name = tokens[i].Text
	}

	// Note: We don't try to extract the numeric value from tokens because:
	// 1. Apple docs often don't include the actual value in the declaration
	// 2. Values are better extracted from documentation or inferred
	// The generator will handle value assignment based on order and context

	if enumCase.Name == "" {
		return nil, fmt.Errorf("failed to parse enum case name")
	}

	return enumCase, nil
}
