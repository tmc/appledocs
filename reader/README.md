# reader

Package `reader` provides an `io/fs.FS` interface to Apple documentation JSON files.

## Overview

The reader package allows access to downloaded Apple documentation without embedding the entire 2GB+ documentation archive into binaries. It provides both low-level filesystem access and high-level query functions for working with frameworks, symbols, and documentation metadata.

## Installation

```bash
go get github.com/tmc/appledocs/reader
```

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/tmc/appledocs/reader"
)

func main() {
    // Open the documentation filesystem
    fsys, err := reader.Open("output/tutorials/data/documentation")
    if err != nil {
        log.Fatal(err)
    }

    // List available frameworks
    frameworks, err := reader.ListFrameworks(fsys)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d frameworks\n", len(frameworks))

    // Get framework information
    info, err := reader.GetFrameworkInfo(fsys, "Foundation")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Framework: %s\n", info.Name)
    fmt.Printf("Abstract: %s\n", info.Abstract)
}
```

### Reading Framework Documentation

```go
// Get the full framework document
doc, err := reader.GetFramework(fsys, "Foundation")
if err != nil {
    log.Fatal(err)
}

fmt.Println(doc.Metadata.Title)
fmt.Println(doc.Abstract[0].Text)
```

### Working with Symbols

```go
// List all symbols in a framework
symbols, err := reader.ListSymbols(fsys, "Foundation")
if err != nil {
    log.Fatal(err)
}

// Get a specific symbol
doc, err := reader.GetSymbol(fsys, "Foundation/NSString")
if err != nil {
    log.Fatal(err)
}

// Search for symbols
matches, err := reader.SearchSymbols(fsys, "Foundation", "Array")
if err != nil {
    log.Fatal(err)
}
```

### Using as fs.FS

The reader implements `io/fs.FS`, so it can be used with any code that accepts a filesystem interface:

```go
fsys, err := reader.Open("output/tutorials/data/documentation")
if err != nil {
    log.Fatal(err)
}

// Use with standard library functions
file, err := fsys.Open("Foundation.json")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

// Use with fs package functions
entries, err := fs.ReadDir(fsys, "Foundation")
if err != nil {
    log.Fatal(err)
}
```

## API Reference

### Core Types

- **`FS`**: Provides filesystem access to Apple documentation
- **`Document`**: Represents an Apple documentation JSON file
- **`Metadata`**: Contains framework and platform information
- **`FrameworkInfo`**: High-level framework information
- **`SymbolInfo`**: High-level symbol information

### Key Functions

#### Opening Documentation

```go
func Open(root string) (*FS, error)
```

Opens a documentation filesystem rooted at the given directory.

#### Framework Operations

```go
func ListFrameworks(fsys *FS) ([]string, error)
func GetFramework(fsys *FS, name string) (*Document, error)
func GetFrameworkInfo(fsys *FS, name string) (*FrameworkInfo, error)
```

#### Symbol Operations

```go
func ListSymbols(fsys *FS, framework string) ([]string, error)
func GetSymbol(fsys *FS, path string) (*Document, error)
func GetSymbolInfo(fsys *FS, path string) (*SymbolInfo, error)
func SearchSymbols(fsys *FS, framework, query string) ([]string, error)
```

#### URL Operations

```go
func GetSymbolByURL(fsys *FS, url string) (*Document, error)
```

Retrieves a symbol by its documentation URL (e.g., `doc://com.apple.foundation/documentation/Foundation/NSString`).

## Directory Structure

The reader expects a directory structure matching Apple's documentation layout:

```
documentation/
├── Foundation.json
├── Foundation/
│   ├── NSString.json
│   └── NSString/
│       └── ...
├── AppKit.json
└── AppKit/
    └── ...
```

Framework-level documentation is stored in JSON files at the root level, while symbols and nested types are organized in subdirectories.

## Use Cases

### Code Generation

The reader package is designed to support code generators like DarwinKit:

```go
fsys, err := reader.Open("output/tutorials/data/documentation")
if err != nil {
    log.Fatal(err)
}

// Enumerate all frameworks and their symbols
frameworks, _ := reader.ListFrameworks(fsys)
for _, fw := range frameworks {
    info, _ := reader.GetFrameworkInfo(fsys, fw)
    symbols, _ := reader.ListSymbols(fsys, fw)

    // Generate bindings for each symbol
    for _, symbol := range symbols {
        doc, _ := reader.GetSymbol(fsys, fw + "/" + symbol)
        // Generate code based on doc.Metadata...
    }
}
```

### Documentation Tools

Build documentation viewers, search tools, or static site generators:

```go
// Search across all frameworks
frameworks, _ := reader.ListFrameworks(fsys)
for _, fw := range frameworks {
    matches, _ := reader.SearchSymbols(fsys, fw, "Button")
    for _, match := range matches {
        info, _ := reader.GetSymbolInfo(fsys, fw + "/" + match)
        fmt.Printf("%s.%s: %s\n", fw, match, info.Abstract)
    }
}
```

## Performance

The reader package uses lazy loading and does not load the entire documentation tree into memory. Files are read on-demand using the standard `io/fs` interfaces, making it suitable for processing large documentation sets.

## Testing

Run the tests:

```bash
go test ./reader/...
```

The tests require the documentation to be downloaded first. Run from the project root.

## License

Same as the parent project.
