# AppledDocs Architecture

## Overview

The appledocs project has two main layers:

1. **Root Package (`appledocs`)** - Public API for reading Apple documentation
2. **Code Generator (`cmd/generate-framework-bindings`)** - Generates Go bindings from documentation

## Package Structure

```
appledocs/
├── types.go         # Core types (Document, Metadata, Token, etc)
├── fs.go            # Filesystem API (Open, ReadDocument, ListFrameworks)
├── maps.go          # Map-based convenience API (LoadMap, GetString, etc)
├── doc.go           # Package documentation
│
├── cmd/generate-framework-bindings/
│   └── main.go      # CLI tool for generating bindings
│
└── internal/generator/
    ├── types.go     # Generator-specific types (ParsedFunction, ParsedClass)
    ├── parser.go    # Parses JSON using appledocs.Document
    ├── codegen.go   # Generates Go code
    └── templates.go # Code templates
```

## How It Works

### 1. Root Package API (types.go, fs.go, maps.go)

The root package provides **two approaches** for reading Apple documentation:

#### Structured API (types.go + fs.go)
```go
// Open documentation filesystem
fsys, _ := appledocs.Open("~/.appledocs/cache/...")

// Read with full type safety
doc, _ := fsys.ReadDocument("Foundation/NSString.json")
title := doc.Metadata.Title
kind := doc.Metadata.SymbolKind
platforms := doc.Metadata.Platforms
```

**Benefits:**
- Full type safety
- IDE autocomplete
- Compile-time validation
- Used by the generator

#### Map-based API (maps.go)
```go
// Read as map for flexibility
doc, _ := appledocs.LoadMap(fsys, "Foundation/NSString.json")
title := appledocs.Title(doc)
kind := appledocs.SymbolKind(doc)
```

**Benefits:**
- Simpler for one-off scripts
- No need to know full structure
- Graceful handling of missing fields

### 2. Code Generator (internal/generator/)

The generator **uses the structured API** from the root package:

```go
// parser.go
func ParseJSONFile(path string) (*ParsedFunction, *ParsedClass, *ParsedProtocol, error) {
    var doc appledocs.Document  // ✓ Uses root package type
    json.Unmarshal(data, &doc)

    externalID := doc.Metadata.ExternalID
    tokens := getObjectiveCVariant(&doc)
    // ... parse tokens into generator types
}
```

**Flow:**
1. `main.go` walks framework directory
2. Calls `parser.ParseJSONFile()` for each JSON
3. Parser uses `appledocs.Document` to read JSON
4. Extracts Objective-C declarations from tokens
5. Converts to generator types (`ParsedFunction`, `ParsedClass`)
6. `codegen.go` generates Go code from parsed types

## Key Design Decisions

### Why Two APIs?

1. **Structured (types.go)** - For code generation
   - Needs reliability and type safety
   - Used by the generator internally
   - Validates structure at compile time

2. **Map-based (maps.go)** - For exploration
   - Simpler for scripts and tools
   - Easier to experiment with
   - No need to update types when schema changes

### Why Internal Generator Package?

The `internal/generator/` package is separate because:
- Generator types (`ParsedFunction`, `ParsedClass`) are **not** the same as `appledocs.Document`
- Generator needs to transform documentation into Go bindings
- Keeps parsing logic separate from API design

### Current Usage Status

✅ **Generator uses root package correctly**
- Uses `appledocs.Document` struct
- Accesses via typed fields
- Proper separation of concerns

✅ **Root package is stable**
- Well-defined types
- Both APIs (structured + map-based) work
- Used in production by generator

## Example: Full Pipeline

```go
// 1. User runs generator
$ go run cmd/generate-framework-bindings/main.go -framework=AppKit

// 2. main.go walks AppKit directory
filepath.Walk(frameworkDir, ...)

// 3. For each JSON, calls parser
fn, cls, proto, err := generator.ParseJSONFile(path)

// 4. Parser uses appledocs.Document
var doc appledocs.Document
json.Unmarshal(data, &doc)

// 5. Extracts Objective-C tokens
tokens := getObjectiveCVariant(&doc)

// 6. Parses into generator types
fn := ParseFunctionDeclaration(tokens)
cls := ParseClassDeclaration(tokens)

// 7. Generates Go code
generator.Generate(opts, functions, classes, protocols)

// 8. Outputs to generated/frameworks/appkit/
├── doc.go
├── classes.gen.go
├── functions.gen.go
└── types.gen.go
```

## Current State Summary

| Component | Status | Notes |
|-----------|--------|-------|
| Root types (types.go) | ✅ Production | Used by generator |
| FS API (fs.go) | ✅ Production | Full filesystem access |
| Map API (maps.go) | ✅ Production | Convenience helpers |
| Generator parser | ✅ Uses root package | Properly integrated |
| Generator codegen | ✅ Working | Generates AppKit, CoreGraphics, etc |
| Generated code | ✅ Compiles | All frameworks build successfully |

## Next Steps for Methods

To generate method bindings (like DarwinKit), the generator needs to:

1. **Walk class subdirectories** - Each class has methods in subdirectory
   ```
   AppKit/NSAlert/           <- Class directory
   ├── runModal().json       <- Method
   ├── setMessageText(_:).json
   └── ...
   ```

2. **Parse method JSON files** - Already have `appledocs.Document` type
   - Extract method name from tokens
   - Parse parameter types
   - Generate objc.Call[] wrappers

3. **Enhance codegen** - Add method generation
   ```go
   func (a Alert) RunModal() ModalResponse {
       return objc.Call[ModalResponse](a, objc.Sel("runModal"))
   }
   ```

The foundation is solid - just need to extend the generator to walk method subdirectories and parse method JSON files using the existing `appledocs.Document` type.
