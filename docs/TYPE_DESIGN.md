# Apple Documentation Type Design

## Overview

This document describes how we generalized the Apple documentation types from the JSON schema analysis.

## Methodology

1. **Schema Analysis** - Analyzed 500 sample JSON files to discover:
   - Field frequency (which fields appear in what % of documents)
   - Field types (what Go types map to each JSON field)
   - Nested structures and their patterns
   - Array element types

2. **Type Generalization Strategy**:
   - Include all fields present in >90% of documents as required fields
   - Include fields present in 50-90% of documents with `omitempty`
   - Include symbol-specific fields conditionally
   - Use concrete types where patterns are consistent
   - Use `interface{}` or `map[string]interface{}` only when truly polymorphic

## Key Findings from Schema Analysis

### Core Fields (100% frequency)
These fields appear in every Apple documentation JSON file:

```
hierarchy              - Document navigation breadcrumbs
identifier             - Unique doc:// URL identifier
kind                   - Document type ("symbol", "article", etc.)
legalNotices           - Copyright information
metadata               - Framework/platform metadata
references             - Referenced symbols/topics
schemaVersion          - JSON schema version
sections               - Content sections
variants               - Language variants (Swift/ObjC)
```

### Common Metadata Structure

The `metadata` object has these fields in >75% of documents:
- `modules` - Array of framework names (100%)
- `role` - Document role (100%)
- `title` - Display title (100%)
- `roleHeading` - Human-readable role (99%)
- `platforms` - Platform availability (90%)
- `externalID` - Symbol identifier like `c:objc(cs)NSView` (77%)
- `symbolKind` - Symbol type: class, func, var, etc. (77%)
- `fragments` - Syntax highlighting tokens (76%)

### Content Patterns

**Primary Content Sections** (96% of documents):
- Kind: "declarations", "content", "parameters", etc.
- Contains code declarations with syntax highlighting
- Flexible content array for rich documentation

**Topic Sections** (33% of documents):
- Organizes methods/properties into logical groups
- Common in class/struct documentation
- Contains arrays of doc:// URLs to related symbols

**See Also Sections** (87% of documents):
- Related documentation links
- Cross-references to similar APIs

## Type Design Decisions

### 1. Concrete Types Over Maps

**Bad** (auto-generated approach):
```go
type Document struct {
    Metadata map[string]interface{} `json:"metadata,omitempty"`
}
```

**Good** (our approach):
```go
type Document struct {
    Metadata Metadata `json:"metadata"`
}

type Metadata struct {
    Title   string     `json:"title"`
    Role    string     `json:"role"`
    Modules []Module   `json:"modules"`
    // ... more concrete fields
}
```

**Rationale**: Concrete types provide:
- Type safety
- Auto-completion in IDEs
- Self-documenting code
- Compile-time checking

### 2. Required vs Optional Fields

**Core fields** (no `omitempty`):
- Present in 100% of documents
- Critical for basic functionality

**Optional fields** (`omitempty`):
- Present in <100% of documents
- Symbol-specific or content-dependent

**Example**:
```go
type Metadata struct {
    Title   string   `json:"title"`            // Always present
    Role    string   `json:"role"`             // Always present

    Platforms []Platform `json:"platforms,omitempty"` // Only for symbols
    ExternalID string    `json:"externalID,omitempty"` // Only for symbols
}
```

### 3. Handling Polymorphism

Some fields are truly polymorphic and require flexibility:

**ContentBlock** - Can be text, code, links, or nested content:
```go
type ContentBlock struct {
    Type       string         `json:"type"` // Discriminator
    Text       string         `json:"text,omitempty"`
    Code       string         `json:"code,omitempty"`
    InlineContent []ContentBlock `json:"inlineContent,omitempty"`
}
```

**ContentSection** - Flexible content array:
```go
type ContentSection struct {
    Kind    string        `json:"kind"` // "declarations", "content", etc.
    Content []interface{} `json:"content,omitempty"` // Truly polymorphic
}
```

### 4. Fragment Structure for Code Display

Apple uses a token-based syntax highlighting system:

```go
type Fragment struct {
    Kind   string `json:"kind"` // "keyword", "identifier", "text"
    Text   string `json:"text"` // The actual text
    Identifier string `json:"identifier,omitempty"` // Reference link
}
```

Example from NSView declaration:
```json
{
  "tokens": [
    {"kind": "keyword", "text": "class"},
    {"kind": "text", "text": " "},
    {"kind": "identifier", "text": "NSView"}
  ]
}
```

### 5. Platform Availability

Consistent structure across all symbols:

```go
type Platform struct {
    Name         string `json:"name"` // "macOS", "iOS", etc.
    Beta         bool   `json:"beta"`
    IntroducedAt string `json:"introducedAt,omitempty"` // "10.0"
    DeprecatedAt string `json:"deprecatedAt,omitempty"`
}
```

## Statistics

From 500 file analysis:
- **72,407** unique field paths discovered
- **60** fields appear in >90% of documents
- **105** fields appear in >50% of documents

Our generalized types cover the 105 most common fields while maintaining type safety.

## Comparison: Auto-Generated vs Hand-Crafted

### Auto-Generated (gentypes)
- 4,280 types from 100 files
- 53,647 lines of code
- Extensive use of `map[string]interface{}`
- Hard to use, poor IDE support

### Hand-Crafted (types.go)
- ~25 core types
- ~350 lines of code
- Concrete types with proper structure
- Excellent IDE support and type safety

## Usage Example

```go
import (
    "github.com/tmc/appledocs"
    "github.com/tmc/appledocs/reader"
)

// Read a document
fs, _ := reader.Open("output/tutorials/data/documentation")
data, _ := fs.ReadFile("AppKit/NSView.json")

// Unmarshal into typed structure
var doc appledocs.Document
json.Unmarshal(data, &doc)

// Type-safe access
fmt.Println(doc.Metadata.Title)        // "NSView"
fmt.Println(doc.Metadata.SymbolKind)   // "class"
fmt.Println(doc.Metadata.ExternalID)   // "c:objc(cs)NSView"

// Access platforms
for _, platform := range doc.Metadata.Platforms {
    fmt.Printf("%s: introduced %s\n", platform.Name, platform.IntroducedAt)
}

// Get related symbols
for _, section := range doc.TopicSections {
    fmt.Printf("Section: %s\n", section.Title)
    for _, id := range section.Identifiers {
        // id is a doc:// URL to another symbol
        fmt.Println("  -", id)
    }
}
```

## Next Steps

1. **Add Helper Methods**:
   ```go
   func (d *Document) GetMethods() []string
   func (d *Document) GetProperties() []string
   func (d *Document) IsDeprecated() bool
   ```

2. **Query Functions**:
   ```go
   func (m *Metadata) SupportsPlatform(name string) bool
   func (m *Metadata) MinimumVersion(platform string) string
   ```

3. **DarwinKit Integration**:
   - Replace custom JSON parsing with these types
   - Use concrete types instead of `map[string]interface{}`
   - Simplify code generation logic
