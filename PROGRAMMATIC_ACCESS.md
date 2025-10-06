# Programmatic Access to Apple Documentation

## Yes! This provides full programmatic access from Go.

The package enables you to:

1. **Load any documentation file**
2. **Extract metadata programmatically**
3. **Navigate the structure**
4. **Generate code from it**

## Quick Example

```go
import "github.com/tmc/appledocs"

// Open docs
fsys, _ := appledocs.Open("output/tutorials/data/documentation")

// Load NSString class
doc, _ := appledocs.LoadMap(fsys, "Foundation/NSString.json")

// Extract data programmatically
title := appledocs.Title(doc)          // "NSString"
kind := appledocs.SymbolKind(doc)      // "class"
extID := appledocs.ExternalID(doc)     // "c:objc(cs)NSString"
platforms := appledocs.Platforms(doc)  // Array of platform info

// Process references
refs := appledocs.References(doc)
for id, ref := range refs {
    refMap := ref.(map[string]interface{})
    if refMap["role"] == "symbol" {
        // This is a method/property - process it
        title := refMap["title"].(string)
        // Generate bindings...
    }
}
```

## What You Can Do Programmatically

### 1. Enumerate Frameworks

```go
fsys, _ := appledocs.Open("docs/")

// List all frameworks
entries, _ := fs.ReadDir(fsys, ".")
for _, e := range entries {
    if appledocs.IsFramework(e.Name()) {
        framework := appledocs.FrameworkName(e.Name())
        // Process framework
    }
}
```

### 2. Extract Class Information

```go
doc, _ := appledocs.LoadMap(fsys, "Foundation/NSString.json")

className := appledocs.Title(doc)        // NSString
symbolKind := appledocs.SymbolKind(doc)  // class
externalID := appledocs.ExternalID(doc)  // c:objc(cs)NSString
url := appledocs.URL(doc)                // doc://...
```

### 3. Get Platform Availability

```go
platforms := appledocs.Platforms(doc)
for _, p := range platforms {
    pmap := p.(map[string]interface{})
    name := pmap["name"].(string)        // "iOS"
    introduced := pmap["introducedAt"]   // "2.0"
    deprecated := pmap["deprecated"]     // false
}
```

### 4. Process Methods/Properties

```go
refs := appledocs.References(doc)

for refID, refData := range refs {
    ref := refData.(map[string]interface{})

    if ref["role"] == "symbol" {
        title := ref["title"].(string)

        // Check if it has declaration info
        if fragments, ok := ref["fragments"]; ok {
            // Parse method signature from fragments
            // Generate bindings
        }
    }
}
```

### 5. Navigate Arbitrary Fields

```go
// Type-safe accessors for common paths
title := appledocs.GetString(doc, "metadata", "title")
version := appledocs.GetInt(doc, "schemaVersion", "major")
beta := appledocs.GetBool(doc, "metadata", "platforms", "beta")

// Get nested maps/arrays
metadata := appledocs.GetMap(doc, "metadata")
platforms := appledocs.GetArray(doc, "metadata", "platforms")
```

## Real-World Usage: Code Generation

Here's how DarwinKit (or similar tools) would use this:

```go
package main

import (
    "github.com/tmc/appledocs"
    "io/fs"
)

func generateBindings(docsPath string) error {
    fsys, _ := appledocs.Open(docsPath)

    // 1. Find all frameworks
    frameworks := []string{}
    entries, _ := fs.ReadDir(fsys, ".")
    for _, e := range entries {
        if appledocs.IsFramework(e.Name()) {
            frameworks = append(frameworks, appledocs.FrameworkName(e.Name()))
        }
    }

    // 2. Process each framework
    for _, fw := range frameworks {
        // Load framework doc
        fwDoc, _ := appledocs.LoadMap(fsys, fw+".json")

        // Get classes from references
        refs := appledocs.References(fwDoc)
        for refID, refData := range refs {
            ref := refData.(map[string]interface{})

            if ref["role"] == "symbol" {
                symbolKind := ref["symbolKind"]

                if symbolKind == "class" {
                    // Load full class doc
                    className := ref["title"].(string)
                    classPath := appledocs.SymbolPath(fw, className)
                    classDoc, _ := appledocs.LoadMap(fsys, classPath)

                    // Generate Go bindings
                    generateGoClass(classDoc)
                }
            }
        }
    }

    return nil
}

func generateGoClass(doc map[string]interface{}) {
    className := appledocs.Title(doc)
    externalID := appledocs.ExternalID(doc)

    // Extract methods from references
    refs := appledocs.References(doc)
    for _, refData := range refs {
        ref := refData.(map[string]interface{})
        if ref["role"] == "symbol" {
            // Generate method binding
            generateMethod(className, ref)
        }
    }
}

func generateMethod(className string, ref map[string]interface{}) {
    // Parse fragments to extract signature
    // Generate Go wrapper code
    // ...
}
```

## Performance Characteristics

- **Lazy loading**: Files read on-demand via `fs.FS`
- **No parsing overhead**: Direct JSON unmarshaling
- **Memory efficient**: Don't need to load entire corpus
- **Fast queries**: Direct map access, no ORM overhead

## Two APIs Available

### 1. Simple Map API (Recommended)

```go
import "github.com/tmc/appledocs"

doc, _ := appledocs.LoadMap(fsys, "Foundation/NSString.json")
title := appledocs.Title(doc)
```

**Pros:**
- 240 lines of code
- Maximum flexibility
- No build dependencies

**Cons:**
- Runtime type assertions
- No compile-time field checking

### 2. Typed API (Optional)

```go
import "github.com/tmc/appledocs/reader"

doc, _ := reader.GetSymbol(fsys, "Foundation/NSString")
title := doc.Metadata.Title  // Compile-time checked!
```

**Pros:**
- Compile-time type safety
- IDE autocomplete
- Structured types

**Cons:**
- More code (~1000 lines)
- Less flexible for edge cases

## Live Demo Output

```
=== Programmatic Access Demo ===

1. Enumerate frameworks:
   - ARKit
   - AVFAudio
   - AVFoundation
   ... (376 total)

2. Load class metadata:
   Title: NSString
   Kind: class
   External ID: c:objc(cs)NSString

3. Platform availability:
   - iOS (since 2.0)
   - iPadOS (since 2.0)
   - Mac Catalyst (since 13.0)
   ... (7 total)

4. Analyze references:
   symbol: 242
   article: 1
   collection: 1

✓ All data is programmatically accessible!
```

## Summary

**Yes, this provides complete programmatic access to Apple documentation from Go.**

You can:
- ✅ Enumerate all frameworks
- ✅ Load any symbol/class/method
- ✅ Extract metadata (names, IDs, signatures)
- ✅ Navigate relationships (superclasses, protocols)
- ✅ Check platform availability
- ✅ Generate code from it

It's designed specifically for code generation tools like DarwinKit.
