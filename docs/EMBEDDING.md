# Embedding Apple Documentation in Go

## Overview

This document describes how to make the Apple documentation data available for import in Go programs using `embed.FS`.

## Data Structure

The crawled documentation lives in:
- `output/tutorials/data/documentation/<Framework>/` - Framework-specific JSON files
- Total size: ~2.0 GB
- Key frameworks:
  - AppKit: 40MB, 719 JSON files
  - Foundation: 56MB
  - UIKit: 58MB
  - CoreGraphics: 5.8MB

## Embedding Challenges

### 1. Size Constraint
- **Problem**: Embedding 2GB of JSON directly would create massive binaries
- **Solution**: Create separate sub-packages per framework or allow selective embedding

### 2. Build Time
- **Problem**: Go embed directive processes all files at build time
- **Solution**: Use lazy loading or external data package pattern

## Proposed Solutions

### Option A: Separate Framework Packages (Recommended)

Create individual packages for each framework:

```
appledocs/
  data/
    appkit/
      doc.go       // package appkit; import _ "embed"
      embedded.go  // //go:embed *.json
    foundation/
      doc.go
      embedded.go
    ...
```

**Pros:**
- Users only import what they need
- Smaller binaries
- Faster builds

**Cons:**
- More packages to maintain
- Requires organization

### Option B: Single Embed with FS Interface

Create a single package with all data:

```go
package appledocs

import "embed"

//go:embed output/tutorials/data/documentation/**/*.json
var Docs embed.FS
```

**Pros:**
- Simple API
- All data in one place

**Cons:**
- 2GB+ binary size
- Slow build times
- Not practical for most use cases

### Option C: External Data with Download Helper

Don't embed, but provide helpers to download/cache locally:

```go
package appledocs

import (
  "io/fs"
  "os"
)

func OpenDocs(cacheDir string) (fs.FS, error) {
  // Check if data exists locally
  // If not, download on-demand
  // Return fs.FS interface
}
```

**Pros:**
- Zero binary size impact
- Can update docs without rebuilding
- Flexible caching

**Cons:**
- Requires network access (first time)
- More complex

### Option D: Hybrid - Small Embed + On-Demand Download

Embed critical/common symbols, download full docs on-demand:

```go
package appledocs

import "embed"

//go:embed data/index/**/*.json
var Index embed.FS  // Small index files only

func DownloadFramework(name string, dir string) error {
  // Download full framework data
}
```

**Pros:**
- Fast common case
- Full data available when needed
- Reasonable binary size

**Cons:**
- More complex
- Requires maintaining index

## Recommended Approach for DarwinKit

For DarwinKit's code generation use case, **Option C** (External Data) is recommended:

1. Keep appledocs as a separate tool for downloading/updating docs
2. Create a thin Go package that provides an FS interface to locally cached docs
3. DarwinKit's code generator imports this package and uses the FS interface
4. Developers run `appledocs` once to populate the cache

Example usage in DarwinKit:

```go
import "github.com/tmc/appledocs/reader"

func main() {
  docsFS, err := reader.Open("~/.cache/appledocs")
  if err != nil {
    log.Fatal("Run 'appledocs crawl' first to download docs")
  }

  // Read framework data
  data, _ := fs.ReadFile(docsFS, "AppKit/NSTextList.json")
  // Generate code...
}
```

## Implementation Plan

1. Create `reader` package with FS interface
2. Support both embedded (small subset) and external (full data) modes
3. Add helpers for common queries (get class, get methods, etc.)
4. Document integration with DarwinKit

## Current Implementation Status

### ✅ Completed

1. **Reader Package** (`reader/`)
   - `reader.go` - fs.FS interface to access Apple docs from disk
   - `query.go` - Helper functions for common queries
   - `doc.go` - Package documentation
   - `reader_test.go` - Tests
   - Basic type definitions for Document, Metadata, Platform, etc.

2. **Type Generator** (`cmd/appledocs/gentypes.go`)
   - Automatically generates Go types from JSON schemas
   - Scans corpus and discovers all field types
   - Generates code with proper json tags
   - Run via: `appledocs -mode=gentypes`

3. **Main Tool Integration**
   - Added `gentypes` mode to main appledocs CLI
   - Flags: `-gentypes-output`, `-gentypes-max-files`

### 🚧 Current Limitations

The auto-generated types have some issues:
- Too many intermediate types (4280 types from 100 files)
- Extensive use of `map[string]interface{}` instead of proper structs
- Overly nested type definitions

### 📋 Next Steps

1. **Improve Type Generator**
   - Better type inference for nested objects
   - Merge similar schemas to reduce duplication
   - Use the existing manually-defined types in `reader/reader.go` as templates

2. **Create Root Package API**
   - Move types from `reader/` to root `appledocs.go`
   - Add `Load(path string, v interface{})` helper
   - Add query functions: `GetMethods()`, `GetProperties()`, etc.

3. **DarwinKit Integration**
   - Update DarwinKit's `generate/tools/appledocs.go` to use this package
   - Remove duplicate type definitions
   - Simplify code generator

## File Structure (Actual)

```
github.com/tmc/appledocs/
  cmd/appledocs/              # CLI tool
    main.go                   # Main entry point with gentypes mode
    gentypes.go               # Type generator implementation
    markdown.go               # Markdown generation
    html.go                   # HTML generation
    validation.go             # Data validation

  reader/                     # Library package for Go imports
    reader.go                 # fs.FS interface + basic types
    query.go                  # Helper query functions
    doc.go                    # Package documentation
    reader_test.go            # Tests
    example_test.go           # Example usage
    README.md                 # Usage documentation

  types/                      # Auto-generated types
    types_generated.go        # Generated by gentypes (53K+ lines)

  output/                     # Crawled documentation (gitignored)
    tutorials/
      data/
        documentation/
          AppKit/             # Framework JSON files
          Foundation/
          ...
```
