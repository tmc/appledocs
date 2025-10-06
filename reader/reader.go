package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// FS provides filesystem access to Apple documentation.
type FS struct {
	root string
	fsys fs.FS
}

// Open creates a new FS rooted at the given directory.
// The directory should contain Apple documentation JSON files.
func Open(root string) (*FS, error) {
	info, err := os.Stat(root)
	if err != nil {
		return nil, fmt.Errorf("open documentation: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("open documentation: %s is not a directory", root)
	}

	return &FS{
		root: root,
		fsys: os.DirFS(root),
	}, nil
}

// Open implements fs.FS.
func (f *FS) Open(name string) (fs.File, error) {
	return f.fsys.Open(name)
}

// ReadFile reads the named file from the documentation filesystem.
func (f *FS) ReadFile(name string) ([]byte, error) {
	file, err := f.fsys.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}

// Stat returns file info for the named file.
func (f *FS) Stat(name string) (fs.FileInfo, error) {
	return fs.Stat(f.fsys, name)
}

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

// ReadDocument reads and parses a documentation JSON file.
func (f *FS) ReadDocument(path string) (*Document, error) {
	data, err := f.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read document %s: %w", path, err)
	}

	var doc Document
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("parse document %s: %w", path, err)
	}

	return &doc, nil
}

// IsFramework returns true if the given name appears to be a framework.
// Framework files are JSON files at the root level.
func IsFramework(name string) bool {
	// Frameworks are .json files at the root (no directory separator)
	return strings.HasSuffix(name, ".json") && !strings.Contains(name, string(filepath.Separator))
}

// FrameworkName extracts the framework name from a path.
// For "Foundation.json" returns "Foundation".
// For "Foundation/NSString.json" returns "Foundation".
func FrameworkName(path string) string {
	// Remove .json suffix
	name := strings.TrimSuffix(path, ".json")
	// Take first path component
	if idx := strings.Index(name, string(filepath.Separator)); idx != -1 {
		return name[:idx]
	}
	return name
}
