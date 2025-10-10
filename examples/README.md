# appledocs Examples

This directory contains example programs demonstrating various use cases for the appledocs package.

## Prerequisites

First, download the Apple documentation:

```bash
cd ../
make download  # or run: ./cmd/appledocs/appledocs -mode download
```

This will create `output/tutorials/data/documentation/` containing the JSON files.

## Examples

### 1. List Frameworks

List all available frameworks in the documentation.

```bash
cd list-frameworks
go run main.go

# Show details
go run main.go -v
```

**Output:**
```
Found 376 frameworks:

  ARKit                AVFAudio             AVFoundation         AVKit
  Accelerate           Accessibility        ActivityKit          AdServices
  ...
```

### 2. List Classes

Find all classes, protocols, structs, and enums in a framework.

```bash
cd list-classes
go run main.go -framework Foundation

# Filter by type
go run main.go -framework UIKit -kind class
go run main.go -framework Swift -kind protocol
```

**Output:**
```
Symbols in Foundation:

Class (127):
  - NSArray
  - NSData
  - NSDictionary
  - NSString
  ...

Protocol (45):
  - NSCoding
  - NSCopying
  ...
```

### 3. Platform Analysis

Analyze platform availability across a framework.

```bash
cd platform-analysis
go run main.go -framework Foundation

# Show deprecated symbols
go run main.go -framework UIKit -deprecated
```

**Output:**
```
Platform Analysis: Foundation

Framework available on:
  - iOS: since 2.0
  - macOS: since 10.0
  - tvOS: since 9.0
  ...

Symbol Availability by Platform:
Platform                Total   Deprecated     Beta
----------------------------------------------------
iOS                      1234           45        2
macOS                    1256           52        1
...
```

### 4. Search Symbols

Search for symbols across all frameworks.

```bash
cd search-symbols
go run main.go -query string

# Limit to specific framework
go run main.go -query button -framework UIKit

# Limit results
go run main.go -query view -max 10
```

**Output:**
```
Searching for 'string'...

Foundation (8 matches):
  - NSString (class) - A static, plain-text Unicode string object
  - NSMutableString (class) - A dynamic plain-text Unicode string
  - NSAttributedString (class) - A string with associated attributes
  ...

Swift (12 matches):
  - String (structure) - A Unicode string value
  - Substring (structure) - A slice of a string
  ...
```

### 5. Extract Methods

Extract and display method signatures for a specific class.

```bash
cd extract-methods
go run main.go -class Foundation/NSString

# Filter by type
go run main.go -class Foundation/NSArray -type instance
go run main.go -class UIKit/UIView -type property
```

**Output:**
```
Methods for NSString

Instance Methods (42):
  - (NSString *)stringByAppendingString:(NSString *)aString
  - (NSString *)substringFromIndex:(NSUInteger)from
  - (NSComparisonResult)compare:(NSString *)string
  ...

Class Methods (15):
  + (instancetype)stringWithFormat:(NSString *)format, ...
  + (instancetype)stringWithUTF8String:(const char *)nullTerminatedCString
  ...

Properties (8):
  - length
  - UTF8String
  - uppercaseString
  ...
```

### 6. Code Generation Demo

Demonstrate how to generate Go bindings from Apple documentation (like DarwinKit).

```bash
cd codegen-demo
go run main.go -framework Foundation -class NSString

# Generate for all classes in a framework
go run main.go -framework Foundation -output foundation_gen.go
```

**Output:**
```go
// Code generated from Apple Documentation for Foundation
// DO NOT EDIT

package foundation

// NSString - A static, plain-text Unicode string object
//
// External ID: c:objc(cs)NSString
type NSString struct {
    // Internal pointer to Objective-C object
    ptr unsafe.Pointer
}

// Properties:
//   - length
//   - UTF8String
//   - uppercaseString

// Methods:
//   - (NSString *)stringByAppendingString:(NSString *)aString
//   - (NSString *)substringFromIndex:(NSUInteger)from
...
```

### 7. Cross-Reference (New!)

Navigate between Swift and Objective-C documentation variants using the cross-reference API.

```bash
cd cross-reference
go run main.go
```

**Features:**
- Get cross-reference information between Swift and ObjC
- Check language variant availability
- Extract declarations for specific languages
- Transform documents between language variants

**Output:**
```
=== NSString Cross-Reference ===
Symbol: NSString
Kind: class

Swift Declaration:
  class NSString

Objective-C Declaration:
  @interface NSString : NSObject

=== Language Variant Availability ===
Foundation/NSArray            : Swift=true ObjC=true
Foundation/NSDate             : Swift=true ObjC=true
Foundation/NSURL              : Swift=true ObjC=true
```

### 8. CoreGraphics Methods (Work in Progress)

Demonstrate method-style CoreGraphics API combined with cross-reference support.

```bash
cd coregraphics-methods
go run main.go
```

**Features:**
- Method-style CoreGraphics API (e.g., `ctx.MoveToPoint(x, y)`)
- Creating bitmap contexts
- Drawing with method-style wrappers
- Cross-referencing CoreGraphics functions with their Swift/ObjC equivalents

**Example API:**
```go
// Create context
ctx := cg.CGBitmapContextCreate(...)

// Use method-style API - much more ergonomic!
ctx.BeginPath()
ctx.MoveToPoint(50, 50)
ctx.AddLineToPoint(150, 50)
ctx.FillPath()

// Compare with documentation
doc, _ := appledocs.GetSymbol(fsys, "CoreGraphics/CGContextMoveToPoint")
ref := appledocs.GetCrossReference(doc)
fmt.Printf("Swift: %s\n", ref.SwiftDeclaration)
fmt.Printf("ObjC:  %s\n", ref.ObjCDeclaration)
fmt.Printf("Go:    ctx.MoveToPoint(x, y)\n")
```

## Building and Installing

To install an example as a command-line tool:

```bash
cd list-frameworks
go install

# Now you can run it from anywhere
list-frameworks -v
```

## Use Cases

These examples demonstrate patterns useful for:

1. **Documentation Exploration** - Browse and search Apple's documentation programmatically
2. **API Analysis** - Analyze platform availability, deprecation patterns, etc.
3. **Code Generation** - Generate language bindings (like DarwinKit does for Go)
4. **Migration Tools** - Find deprecated APIs and suggest replacements
5. **IDE Integration** - Build autocomplete, hover information, etc.
6. **Static Analysis** - Check if your code uses unavailable APIs for target platforms

## Tips

- All examples accept `-docs` flag to specify documentation location
- Use `-h` flag to see all options for each example
- Examples use both the typed API (for clean code) and map API (for flexibility)
- Check the source code to see different usage patterns

## Next Steps

- Modify examples to suit your needs
- Combine techniques (e.g., search + method extraction)
- Build your own tools using the appledocs package
- See the main [README](../README.md) for full API documentation
