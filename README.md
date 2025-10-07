# appledocs

Go package for programmatic access to Apple's documentation.

## Installation

```bash
go get github.com/tmc/appledocs
```

## Quick Start

```go
import "github.com/tmc/appledocs"

// Open docs directory
fsys, _ := appledocs.Open("output/tutorials/data/documentation")

// Load a symbol with typed API (recommended)
doc, _ := appledocs.GetSymbol(fsys, "Foundation/NSString")
fmt.Println(doc.Metadata.Title)        // NSString
fmt.Println(doc.Metadata.SymbolKind)   // class
fmt.Println(doc.Metadata.ExternalID)   // c:objc(cs)NSString

// Iterate methods - all typed, no casts!
for id, ref := range doc.References {
    if ref.Role == "symbol" && ref.SymbolKind == "method" {
        fmt.Println(ref.Title)
    }
}
```

## Two APIs

### Typed API (Recommended)

Clean, type-safe access to Apple documentation:

```go
// List frameworks
frameworks, _ := appledocs.ListFrameworks(fsys)

// Load symbol
doc, _ := appledocs.GetSymbol(fsys, "Foundation/NSString")

// Access fields (compile-time checked!)
title := doc.Metadata.Title
kind := doc.Metadata.SymbolKind
platforms := doc.Metadata.Platforms

// Process references
for id, ref := range doc.References {
    if ref.Role == "symbol" {
        // ref.Title, ref.Fragments all typed
    }
}
```

**Use when:**
- Writing production code
- You want compile-time safety
- You value clean, maintainable code

### Map API (For Flexibility)

Access any field, handle edge cases:

```go
// Load as map
raw, _ := appledocs.LoadMap(fsys, "Foundation/NSString.json")

// Use helpers
title := appledocs.Title(raw)
kind := appledocs.SymbolKind(raw)

// Navigate arbitrary paths
customField := appledocs.GetString(raw, "custom", "nested", "field")
```

**Use when:**
- Exploring unknown structure
- Accessing undocumented fields
- Writing quick scripts

## API Reference

### Typed API

```go
// Opening docs
fsys, err := Open(path string) (*FS, error)

// Framework queries
ListFrameworks(fsys *FS) ([]string, error)
GetFramework(fsys *FS, name string) (*Document, error)
GetFrameworkInfo(fsys *FS, name string) (*FrameworkInfo, error)

// Symbol queries
GetSymbol(fsys *FS, path string) (*Document, error)
GetSymbolByURL(fsys *FS, url string) (*Document, error)
GetSymbolInfo(fsys *FS, path string) (*SymbolInfo, error)
ListSymbols(fsys *FS, framework string) ([]string, error)
SearchSymbols(fsys *FS, framework, query string) ([]string, error)

// Utilities
IsFramework(name string) bool
FrameworkName(path string) string
```

### Map API

```go
// Core
LoadMap(fsys fs.FS, path string) (map[string]interface{}, error)
Load(fsys fs.FS, path string, v interface{}) error

// Navigation
GetString(m map, path ...string) string
GetInt(m map, path ...string) int
GetBool(m map, path ...string) bool
GetMap(m map, path ...string) map[string]interface{}
GetArray(m map, path ...string) []interface{}

// Document helpers
Title(m) string
Kind(m) string
SymbolKind(m) string
ExternalID(m) string
URL(m) string
InterfaceLanguage(m) string

// Collections
References(m) map[string]interface{}
Metadata(m) map[string]interface{}
Platforms(m) []interface{}
Modules(m) []interface{}

// Path utilities
SymbolPath(framework, symbol) string
```

## Types

Core types for typed API:

```go
type Document struct {
    Identifier Identifier
    Kind       string
    Metadata   Metadata
    Abstract   []InlineContent
    Hierarchy  Hierarchy
    References map[string]Reference
    // ...
}

type Metadata struct {
    ExternalID string
    Title      string
    SymbolKind string
    Platforms  []Platform
    Modules    []Module
    Fragments  []Fragment
    // ...
}

type Reference struct {
    Identifier string
    Kind       string
    Role       string
    Title      string
    SymbolKind string
    Fragments  []Fragment
    // ...
}
```

See [types.go](types.go) for complete type definitions.

## Examples

### List All Frameworks

```go
fsys, _ := appledocs.Open("docs/")
frameworks, _ := appledocs.ListFrameworks(fsys)
for _, fw := range frameworks {
    fmt.Println(fw)
}
```

### Find All Classes in a Framework

```go
doc, _ := appledocs.GetFramework(fsys, "Foundation")
for id, ref := range doc.References {
    if ref.Role == "symbol" && ref.SymbolKind == "class" {
        fmt.Println(ref.Title)
    }
}
```

### Extract Method Signatures

```go
doc, _ := appledocs.GetSymbol(fsys, "Foundation/NSString")
for id, ref := range doc.References {
    if ref.Role == "symbol" && ref.SymbolKind == "method" {
        // ref.Fragments contains the signature parts
        for _, frag := range ref.Fragments {
            fmt.Printf("%s", frag.Text)
        }
        fmt.Println()
    }
}
```

### Check Platform Availability

```go
doc, _ := appledocs.GetSymbol(fsys, "Foundation/NSString")
for _, platform := range doc.Metadata.Platforms {
    fmt.Printf("%s: introduced %s\n",
        platform.Name, platform.IntroducedAt)
    if platform.Deprecated {
        fmt.Printf("  (deprecated in %s)\n", platform.DeprecatedAt)
    }
}
```

## Use Cases

### Code Generation (DarwinKit)

Generate Go bindings from Objective-C classes:

```go
frameworks, _ := appledocs.ListFrameworks(fsys)
for _, fw := range frameworks {
    doc, _ := appledocs.GetFramework(fsys, fw)
    for id, ref := range doc.References {
        if ref.SymbolKind == "class" {
            generateGoClass(ref.Title, doc)
        }
    }
}
```

### Documentation Search

Search for symbols across frameworks:

```go
matches, _ := appledocs.SearchSymbols(fsys, "Foundation", "string")
for _, match := range matches {
    info, _ := appledocs.GetSymbolInfo(fsys, "Foundation/"+match)
    fmt.Printf("%s: %s\n", info.Title, info.Abstract)
}
```

### API Analysis

Analyze platform support across a framework:

```go
symbols, _ := appledocs.ListSymbols(fsys, "Foundation")
iosCount := 0
for _, symbol := range symbols {
    doc, _ := appledocs.GetSymbol(fsys, "Foundation/"+symbol)
    for _, p := range doc.Metadata.Platforms {
        if p.Name == "iOS" {
            iosCount++
            break
        }
    }
}
```

## Design Philosophy

This package provides both:

1. **Typed API** (406 lines) - Clean usage, compile-time safety
2. **Map API** (232 lines) - Maximum flexibility

Total: ~650 lines vs 50,000+ lines of auto-generated code.

See [RADICAL_SIMPLICITY.md](RADICAL_SIMPLICITY.md) for design rationale.

## Performance

- **Lazy loading**: Files read on-demand via fs.FS
- **No parsing overhead**: Direct JSON unmarshaling
- **Memory efficient**: Don't load entire corpus
- **Fast queries**: Direct map/struct access

## Distribution

The Apple documentation is ~2GB of JSON. We provide multiple distribution options:

### Option 1: Local Documentation (Recommended for Development)

Download docs manually:

```bash
make download  # Downloads to output/tutorials/data/documentation
```

Use in your code:

```go
fsys, _ := appledocs.Open("output/tutorials/data/documentation")
```

### Option 2: Embedded Data Module (Coming Soon)

Import versioned documentation as a Go module:

```bash
go get github.com/tmc/appledocs-data/v17@latest
```

```go
import data "github.com/tmc/appledocs-data/v17"

fsys := data.FS()  // iOS 17 / macOS 14 docs
```

**Note**: Binary size increases by 200-400MB

### Option 3: On-Demand Fetching (Planned)

Download specific frameworks as needed:

```go
import "github.com/tmc/appledocs-fetch"

fsys, _ := fetch.Framework("Foundation", "17.0")  // Caches locally
```

See [DISTRIBUTION.md](DISTRIBUTION.md) for details on module structure and versioning strategy.

## Related Projects

- [DarwinKit](https://github.com/progrium/darwinkit) - Go bindings for macOS frameworks (uses this package)

## License

MIT
