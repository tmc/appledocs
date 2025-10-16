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
}

// ParsedMethod represents an Objective-C method declaration.
type ParsedMethod struct {
	Name          string   // Go-style method name (e.g., "InitWithFrame")
	Selector      string   // Objective-C selector (e.g., "initWithFrame:")
	IsClassMethod bool     // true for class methods (+), false for instance methods (-)
	ReturnType    string   // Objective-C return type
	Parameters    []Parameter
	Comment       string
	Availability  Availability
	DocURL        string
	Abstract      string
}

// ParsedProperty represents an Objective-C property declaration.
type ParsedProperty struct {
	Name         string
	Type         string
	Attributes   []string // e.g., "readonly", "nonatomic", "strong"
	Comment      string
	Availability Availability
	DocURL       string
	Abstract     string
}

// ParsedProtocol represents an Objective-C protocol declaration.
type ParsedProtocol struct {
	Name         string
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
