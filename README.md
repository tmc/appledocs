# appledocs

Simple Go package for accessing Apple documentation JSON files.

## Philosophy: Radical Simplicity

This package provides a **minimalist** API for reading Apple's documentation. Instead of generating thousands of types, we embrace Go's `map[string]interface{}` with thoughtful helper functions.

## Installation

```bash
go get github.com/tmc/appledocs
```

## Quick Start

```go
import "github.com/tmc/appledocs"

// Open docs
fsys, _ := appledocs.Open("output/tutorials/data/documentation")

// Load Foundation
foundation, _ := appledocs.LoadMap(fsys, "Foundation.json")
fmt.Println(appledocs.Title(foundation))  // "Foundation"

// Load NSString
nsstring, _ := appledocs.LoadMap(fsys, "Foundation/NSString.json")
fmt.Println(appledocs.SymbolKind(nsstring))  // "class"
fmt.Println(appledocs.ExternalID(nsstring))  // "c:objc(cs)NSString"
```

## Why Simple Maps?

Auto-generating types from the JSON creates 4,000+ types from just 100 files. That's absurd.

The JSON structure is consistent enough that `map[string]interface{}` + helpers works great:
- **50 lines** instead of 50,000 lines of generated code
- Fast builds, no code generation
- Type safety where it matters (the helpers)

See [RADICAL_SIMPLICITY.md](RADICAL_SIMPLICITY.md) for details.

## API

### Core
- `Open(path)` - Open docs directory
- `LoadMap(fsys, path)` - Load JSON as map

### Helpers
- `Title(doc)`, `Kind(doc)`, `SymbolKind(doc)`, `ExternalID(doc)`
- `Platforms(doc)`, `Modules(doc)`, `References(doc)`, `Metadata(doc)`
- `GetString(doc, "path", "to", "field")` - Navigate nested maps

See [appledocs.go](appledocs.go) for full API.

## Packages

- `github.com/tmc/appledocs` - Simple map-based API (recommended)
- `github.com/tmc/appledocs/reader` - Typed API with more structure
- `github.com/tmc/appledocs/cmd/appledocs` - Crawler and generator tools
