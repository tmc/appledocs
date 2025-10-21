package occ2go

import "sort"

// ParsedFunction represents a C function declaration extracted from Apple documentation.
type ParsedFunction struct {
	Name         string
	ReturnType   string
	Parameters   []Parameter
	Comment      string
	Availability Availability
	DocURL       string // Apple documentation URL
	Abstract     string // Short description from Apple docs
}

// ParsedClass represents an Objective-C class declaration.
type ParsedClass struct {
	Name         string
	SuperClass   string
	Methods      []*ParsedMethod   // Instance and class methods
	Properties   []*ParsedProperty // Properties
	Comment      string
	Availability Availability
	DocURL       string
	Abstract     string
	Overview     string // Extended description from documentation overview section
}

// HasInitMethods returns true if the class has at least one documented init method.
// Init methods either have selectors that start with "init" (e.g., "init", "initWithFrame:")
// or are marked as initializers in the documentation (IsInitializer == true).
// Classes without init methods cannot be instantiated directly and should not have
// New*() constructors generated.
func (c *ParsedClass) HasInitMethods() bool {
	for _, method := range c.Methods {
		if method.IsInitializer {
			return true
		}
		if len(method.Selector) >= 4 && method.Selector[:4] == "init" {
			return true
		}
	}
	return false
}

// ParsedMethod represents an Objective-C method declaration.
type ParsedMethod struct {
	Name          string // Go-style method name (e.g., "InitWithFrame")
	Selector      string // Objective-C selector (e.g., "initWithFrame:")
	IsClassMethod bool   // true for class methods (+), false for instance methods (-)
	IsInitializer bool   // true if this is an initializer (symbolKind == "init" in docs)
	ReturnType    string // Objective-C return type
	Parameters    []Parameter
	Comment       string
	Availability  Availability
	DocURL        string
	Abstract      string
}

// ParsedProperty represents an Objective-C property declaration.
type ParsedProperty struct {
	Name            string
	Type            string      // Swift/Go type (e.g., "String?", "string", "unsafe.Pointer")
	ObjCType        string      // Original Objective-C type (e.g., "NSString *") for proper type mapping
	Attributes      []string    // e.g., "readonly", "nonatomic", "strong"
	IsClassProperty bool        // true for class properties (static/class var in Swift)
	Comment         string
	Availability    Availability
	DocURL          string
	Abstract        string
}

// ParsedProtocol represents an Objective-C protocol declaration.
type ParsedProtocol struct {
	Name         string
	Comment      string
	Availability Availability
	DocURL       string
	Abstract     string
}

// ParsedEnum represents an Objective-C enum type declaration.
type ParsedEnum struct {
	Name         string
	BaseType     string // Underlying type (e.g., "NSUInteger", "NSInteger")
	Cases        []*ParsedEnumCase
	Comment      string
	Availability Availability
	DocURL       string
	Abstract     string
	IsOptions    bool // true for NS_OPTIONS (bitfield), false for NS_ENUM
}

// ParsedEnumCase represents a single enum constant/case.
type ParsedEnumCase struct {
	Name         string
	Value        string // The numeric value or expression (can be int literal or expression)
	IntValue     int    // The resolved integer value (populated by extract-enum-values tool)
	Comment      string
	Availability Availability
	DocURL       string
	Abstract     string
}

// ParsedTypedef represents a C typedef declaration (e.g., typedef int CIFormat).
type ParsedTypedef struct {
	Name         string
	BaseType     string // The underlying type (e.g., "int", "CGRect (^)(int, CGRect)")
	Comment      string
	Availability Availability
	DocURL       string
	Abstract     string
	IsTypedEnum  bool     // true for NS_TYPED_ENUM (has associated constants)
	Constants    []string // List of constant names (e.g., ["kCIFormatARGB8", "kCIFormatBGRA8"])
}

// ParsedConstant represents an extern const declaration (e.g., CORE_IMAGE_EXPORT const CIFormat kCIFormatARGB8).
type ParsedConstant struct {
	Name         string
	Type         string // The type (e.g., "CIFormat")
	Comment      string
	Availability Availability
	DocURL       string
	Abstract     string
}

// Parameter represents a function parameter.
type Parameter struct {
	Name string
	Type string
}

// Availability contains version information for API availability across platforms.
// Uses maps for flexibility - handles new platforms without code changes.
type Availability struct {
	// IntroducedAt maps platform name to version string (e.g., "macOS" -> "10.14")
	IntroducedAt map[string]string

	// DeprecatedAt maps platform name to deprecation version
	DeprecatedAt map[string]string

	Beta bool
}

// IsEmpty returns true if no version information is available.
func (a *Availability) IsEmpty() bool {
	return len(a.IntroducedAt) == 0 && len(a.DeprecatedAt) == 0
}

// Platforms returns a sorted list of platforms with availability info.
func (a *Availability) Platforms() []string {
	platforms := make([]string, 0, len(a.IntroducedAt))
	for p := range a.IntroducedAt {
		platforms = append(platforms, p)
	}
	sort.Strings(platforms)
	return platforms
}
