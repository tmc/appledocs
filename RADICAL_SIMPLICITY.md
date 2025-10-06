# Radical Simplicity: Data-Driven Design

## What the Data Actually Shows

After analyzing 500 files:

### Document Structure (100% consistent)
```
Every document has exactly these fields:
- identifier {interfaceLanguage, url}
- kind: "symbol" | "article"
- metadata {role, title, modules, ...}
- hierarchy {paths: [][]string}
- references: map[string]Reference
- schemaVersion {major, minor, patch}
- legalNotices {copyright, privacyPolicy, termsOfUse}
```

### Symbol Kinds (from data)
```
Top symbol kinds:
- property (207)
- var (119)
- method (46)
- struct (37)
- init (16)
- class (16)
- enum (11)
- protocol (4)
- module (4)
```

## The Simplest Possible Design

### 1. Just Use `map[string]interface{}` + Helpers

```go
// appledocs.go - THE ENTIRE API
package appledocs

import (
    "encoding/json"
    "io/fs"
    "os"
)

// Open returns an fs.FS for the docs directory
func Open(path string) (fs.FS, error) {
    return os.DirFS(path), nil
}

// Load unmarshals a JSON file
func Load(fsys fs.FS, path string, v interface{}) error {
    data, err := fs.ReadFile(fsys, path)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, v)
}

// LoadMap is a convenience for loading as map
func LoadMap(fsys fs.FS, path string) (map[string]interface{}, error) {
    var m map[string]interface{}
    err := Load(fsys, path, &m)
    return m, err
}

// GetString safely gets a string from nested map
func GetString(m map[string]interface{}, path ...string) string {
    current := m
    for i, key := range path {
        if i == len(path)-1 {
            if v, ok := current[key].(string); ok {
                return v
            }
            return ""
        }
        if next, ok := current[key].(map[string]interface{}); ok {
            current = next
        } else {
            return ""
        }
    }
    return ""
}

// Helper accessors
func Kind(m map[string]interface{}) string {
    return GetString(m, "kind")
}

func SymbolKind(m map[string]interface{}) string {
    return GetString(m, "metadata", "symbolKind")
}

func Title(m map[string]interface{}) string {
    return GetString(m, "metadata", "title")
}

func ExternalID(m map[string]interface{}) string {
    return GetString(m, "metadata", "externalID")
}
```

### 2. Usage (DarwinKit)

```go
import "github.com/tmc/appledocs"

fsys, _ := appledocs.Open("output/tutorials/data/documentation")

// Load Foundation framework
foundation, _ := appledocs.LoadMap(fsys, "Foundation.json")
title := appledocs.Title(foundation) // "Foundation"

// Load NSString class
nsstring, _ := appledocs.LoadMap(fsys, "Foundation/NSString.json")
kind := appledocs.SymbolKind(nsstring) // "class"
extID := appledocs.ExternalID(nsstring) // "c:objc(cs)NSString"

// References are just the map
refs := nsstring["references"].(map[string]interface{})
for id, ref := range refs {
    refMap := ref.(map[string]interface{})
    if refMap["role"] == "symbol" {
        // This is a method/property
    }
}
```

That's it. **50 lines total.**

## Why This Works

1. **JSON is self-describing** - We don't need 4,000 types
2. **Structure is consistent** - Same fields, same nesting
3. **Callers know what they want** - DarwinKit knows it wants `metadata.externalID`
4. **Go's type assertions work fine** - `m["references"].(map[string]interface{})`
5. **No code generation needed** - No build step, no maintenance

## What We Lose vs. 4,000 Types

- Compile-time type safety for field access
- IDE autocomplete for field names
- Type errors become runtime panics

## What We Gain

- **50 lines** instead of 50,000 lines of generated code
- No build step
- No type explosion
- Dead simple to understand
- Easy to extend (just add more helper functions)

## Optional: Add a Few Key Types

If you want *some* type safety for the most common operations:

```go
// types.go - OPTIONAL, only for common patterns
type Metadata struct {
    Title      string   `json:"title"`
    Role       string   `json:"role"`
    SymbolKind string   `json:"symbolKind,omitempty"`
    ExternalID string   `json:"externalID,omitempty"`
    Modules    []Module `json:"modules"`
}

type Module struct {
    Name string `json:"name"`
}

// Extract metadata from any doc
func GetMetadata(m map[string]interface{}) (*Metadata, error) {
    data, _ := json.Marshal(m["metadata"])
    var meta Metadata
    err := json.Unmarshal(data, &meta)
    return &meta, err
}
```

Now you have type safety where it matters, flexibility everywhere else.

## The Radical Insight

**The JSON schema IS the documentation.** We don't need to replicate it in Go types. We just need:

1. A way to read files (`fs.FS`)
2. A way to parse JSON (`json.Unmarshal`)
3. A few helpers for common paths (`GetString(m, "metadata", "title")`)

Everything else is **overengineering**.

## Recommendation

Ship this as the initial version:
- `appledocs.go` - 50 lines
- `reader/` package - DELETE (or keep as examples)
- `types/types_generated.go` - DELETE
- `cmd/appledocs/gentypes.go` - KEEP (for analysis only)

Done. Simple. Maintainable. Sufficient.
