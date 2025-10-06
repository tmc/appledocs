package appledocs

// Document represents an Apple documentation JSON file.
type Document struct {
	// Identifier uniquely identifies this documentation node
	Identifier Identifier `json:"identifier"`

	// Kind describes the type of documentation (symbol, article, etc)
	Kind string `json:"kind"`

	// Metadata contains framework and platform information
	Metadata Metadata `json:"metadata"`

	// Abstract provides a brief description
	Abstract []InlineContent `json:"abstract,omitempty"`

	// Hierarchy shows the documentation tree structure
	Hierarchy Hierarchy `json:"hierarchy,omitempty"`

	// References contains referenced symbols and topics
	References map[string]Reference `json:"references,omitempty"`

	// TopicSections organizes related topics
	TopicSections []TopicSection `json:"topicSections,omitempty"`

	// PrimaryContentSections contains the main documentation content
	PrimaryContentSections []ContentSection `json:"primaryContentSections,omitempty"`
}

// Identifier uniquely identifies a documentation node.
type Identifier struct {
	// InterfaceLanguage is typically "swift" or "occ"
	InterfaceLanguage string `json:"interfaceLanguage"`

	// URL is the documentation URL (e.g., "doc://com.apple.foundation/documentation/Foundation")
	URL string `json:"url"`
}

// Metadata contains framework and platform information.
type Metadata struct {
	// ExternalID is the framework or symbol identifier
	ExternalID string `json:"externalID,omitempty"`

	// Title is the display name
	Title string `json:"title"`

	// Role describes the documentation role (collection, symbol, etc)
	Role string `json:"role,omitempty"`

	// RoleHeading provides a human-readable role description
	RoleHeading string `json:"roleHeading,omitempty"`

	// SymbolKind describes the symbol type (module, class, struct, etc)
	SymbolKind string `json:"symbolKind,omitempty"`

	// Modules lists the modules this documentation belongs to
	Modules []Module `json:"modules,omitempty"`

	// Platforms lists supported platforms
	Platforms []Platform `json:"platforms,omitempty"`

	// Fragments contains syntax elements for code display
	Fragments []Fragment `json:"fragments,omitempty"`
}

// Module represents a framework or module.
type Module struct {
	Name string `json:"name"`
}

// Platform describes platform availability.
type Platform struct {
	Name         string `json:"name"`
	IntroducedAt string `json:"introducedAt,omitempty"`
	DeprecatedAt string `json:"deprecatedAt,omitempty"`
	Beta         bool   `json:"beta"`
	Deprecated   bool   `json:"deprecated,omitempty"`
	Unavailable  bool   `json:"unavailable,omitempty"`
}

// Fragment represents a syntax element.
type Fragment struct {
	Kind string `json:"kind"`
	Text string `json:"text"`
}

// InlineContent represents inline documentation content.
type InlineContent struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

// Hierarchy describes the documentation tree structure.
type Hierarchy struct {
	Paths [][]string `json:"paths"`
}

// Reference represents a referenced documentation node.
type Reference struct {
	Identifier string          `json:"identifier"`
	Kind       string          `json:"kind"`
	Role       string          `json:"role,omitempty"`
	Title      string          `json:"title"`
	Type       string          `json:"type"`
	URL        string          `json:"url"`
	Abstract   []InlineContent `json:"abstract,omitempty"`
	Fragments  []Fragment      `json:"fragments,omitempty"`
	SymbolKind string          `json:"symbolKind,omitempty"` // Added for easier filtering
}

// TopicSection organizes related topics.
type TopicSection struct {
	Title       string   `json:"title,omitempty"`
	Anchor      string   `json:"anchor,omitempty"`
	Identifiers []string `json:"identifiers"`
	Generated   bool     `json:"generated,omitempty"`
}

// ContentSection contains documentation content.
type ContentSection struct {
	Kind    string        `json:"kind"`
	Content []interface{} `json:"content,omitempty"`
}
