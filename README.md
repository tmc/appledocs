# appledocs

[![Build Status](https://img.shields.io/badge/build-100%25-brightgreen)](FRAMEWORK_COVERAGE.md)
[![Frameworks](https://img.shields.io/badge/frameworks-69%2F270-blue)](FRAMEWORK_COVERAGE.md)
[![Platform](https://img.shields.io/badge/platform-macOS-lightgrey)](PLATFORM_RESTRICTIONS.md)
[![Go Version](https://img.shields.io/badge/go-1.24.1%2B-00ADD8)](go.mod)

Go package for programmatic access to Apple's documentation and framework bindings.

## Table of Contents

- [Framework Bindings](#framework-bindings) - Go bindings for 69 macOS frameworks
- [Documentation API](#documentation-api) - Parse Apple's documentation JSON
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Examples](#examples)
- [Swift Bindings](#swift-binding-generation)

## Installation

### For Documentation Parsing

```bash
go get github.com/tmc/appledocs
```

### For Framework Bindings

```bash
# Import specific frameworks
import "github.com/tmc/appledocs/generated/appkit"
import "github.com/tmc/appledocs/generated/foundation"
import "github.com/tmc/appledocs/generated/metal"
```

See [FRAMEWORK_COVERAGE.md](FRAMEWORK_COVERAGE.md) for all available frameworks.

## Documentation API

### Quick Start

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

## Framework Bindings

This project generates comprehensive, type-safe Go bindings for Apple frameworks using [purego](https://github.com/ebitengine/purego) for cgo-free Objective-C interop.

### Architecture Overview

The binding generator transforms Apple's official documentation into idiomatic Go code through a multi-layered pipeline:

```
┌─────────────────────────────────────────────────────────────────┐
│                   Apple Documentation JSON                       │
│              (~/.appledocs/cache/*.json files)                   │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Discovery & Loading Phase                     │
│  • Framework discovery (pattern matching, regex)                │
│  • Cross-framework type registry (existing bindings)            │
│  • Symbol extraction from API collections                       │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Objective-C Parsing Phase                     │
│  • Token analysis (Objective-C/Swift declarations)              │
│  • Symbol-specific parsers (classes, methods, properties)       │
│  • Type declaration extraction (enums, protocols, typedefs)     │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                     Type Resolution Phase                        │
│  • Framework-aware type mapping (ObjC → Go)                     │
│  • Cross-framework type lookups                                 │
│  • Geometry types, reference types, class types                 │
│  • Dependency hierarchy enforcement                             │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                   Generation Preparation Phase                   │
│  • Dependency sorting (topological order)                       │
│  • Parent stub generation (missing superclasses)                │
│  • Property override merging (manual corrections)               │
│  • Method signature processing                                  │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Template Execution Phase                      │
│  • Two-tier function registration:                              │
│    - Core utilities (string manipulation, formatting)           │
│    - Generator methods (state-dependent operations)             │
│  • Module generation (txtar archive format)                     │
│  • Per-symbol file generation (classes, protocols, enums)       │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                       Generated Go Code                          │
│  • Type-safe wrappers (classes, methods, properties)            │
│  • Automatic memory management (Autorelease)                    │
│  • Cached selectors (performance optimization)                  │
│  • Example tests (usage demonstrations)                         │
└─────────────────────────────────────────────────────────────────┘
```

### Key Components

#### Documentation Processing
- **Symbol Extraction**: Discovers C functions, Objective-C classes, protocols, and enums from Apple's JSON documentation
- **API Collections**: Synthesizes missing symbols from `-api.json` collection files
- **Type Registry**: Builds cross-framework type index from previously generated bindings

#### Parsing & Type Resolution
- **Objective-C Parser**: Token-based analysis of class declarations, method signatures, and property definitions
- **Type Mapping**: Framework-aware conversion of Objective-C types to idiomatic Go equivalents
- **Hierarchy Management**: Prevents circular dependencies through framework-level dependency tracking

#### Code Generation
- **Template-Driven**: Uses Go text/template with custom helper functions for consistent, maintainable output
- **Memory Safety**: Automatic Autorelease() injection for Objective-C reference counting
- **Performance**: Selector caching via compile-time lookup tables

### Coverage

- **69 frameworks** with complete Go bindings (see [FRAMEWORK_COVERAGE.md](FRAMEWORK_COVERAGE.md))
- **100% build success rate** (68/68 buildable frameworks)
- **270 total macOS frameworks** available for binding generation
- **Latest additions**: Contacts, ContactsUI

### Notable Frameworks

**UI & Graphics**: AppKit, QuartzCore, Metal, MetalKit, CoreGraphics, CoreImage, WebKit
**Media**: AVFoundation, CoreAudio, CoreVideo, ImageIO
**ML & Vision**: CoreML, Vision, CreateML, SoundAnalysis
**System**: Foundation, Security, FileProvider, SystemExtensions
**Cloud & Data**: CloudKit, CoreData, Contacts

See [FRAMEWORK_COVERAGE.md](FRAMEWORK_COVERAGE.md) for complete list with statistics.

### Platform Support

While 270+ frameworks are documented by Apple, only macOS frameworks can currently be generated. iOS, watchOS, tvOS, and other platform-specific frameworks require platform-specific SDKs. See [PLATFORM_RESTRICTIONS.md](PLATFORM_RESTRICTIONS.md) for details.

## Swift Binding Generation

In addition to Objective-C framework bindings, this project includes experimental tooling for generating Go bindings from Swift framework extensions:

- **Parse** .swiftinterface files with SwiftSyntax
- **Generate** Swift @_cdecl wrappers
- **Create** idiomatic Go packages with type safety

See [SWIFT_BINDINGS.md](SWIFT_BINDINGS.md) for complete workflow and examples.

### Example: Photos Framework

```go
import "github.com/tmc/appledocs/photosframework"

// Type-safe, idiomatic Go API for Swift extensions
library := photosframework.SharedPhotoLibrary()
defer library.Release()

// Access Swift-only APIs (not available in Objective-C!)
iter := fetchResult.Iterator()  // Swift Sequence protocol
for change := iter.Next(); change != nil; change = iter.Next() {
    // Process persistent changes
}
```

**Benefits:**
- ✅ No cgo required (uses purego)
- ✅ Type-safe Go wrappers
- ✅ Automatic memory management
- ✅ Access to Swift-only framework extensions

## Related Projects

- [DarwinKit](https://github.com/progrium/darwinkit) - Go bindings for macOS frameworks (uses this package)

## License

MIT
