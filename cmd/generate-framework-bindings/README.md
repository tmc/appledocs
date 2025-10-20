# Framework Bindings Generator - Improvements

## Overview

Enhanced code generation pipeline for creating Go bindings from Apple's documentation JSON files.

## Key Improvements 

### 1. Error Handling
- `ErrorCollector` accumulates all errors with context
- Detailed error reporting by stage/file/symbol
- Strict mode option (-strict)

### 2. Enhanced Parsing
- **Blocks** - Objective-C closures
- **Methods** - Instance and class methods
- **Properties** - With attributes
- **Categories** - Extensions
- **Enums** - Constants (partial)

### 3. Better Documentation
- Full API doc extraction from JSON
- Proper GoDoc comments
- Parameter/return descriptions
- Code examples preserved

### 4. New Output Files
- `methods.gen.go` - Objective-C methods
- `properties.gen.go` - Property declarations
- `blocks.gen.go` - Block type definitions

## Files

- `main.go` - Original generator
- `main_improved.go` - Enhanced version
- `errors.go` - Error handling
- `parsers.go` - Enhanced parsers
- `docgen.go` - Documentation extraction

## Usage

```bash
# Basic
go run ./cmd/generate-framework-bindings/*.go -framework Foundation

# Verbose with error details
go run ./cmd/generate-framework-bindings/*.go -framework CoreGraphics -verbose

# Strict mode (fail on errors)
go run ./cmd/generate-framework-bindings/*.go -framework AppKit -strict

# EXPERIMENTAL: Swift interop mode (generates bindings for calling Swift via purego)
go run ./cmd/generate-framework-bindings/*.go -framework MySwiftFramework -swift-interop
```

## Swift Interop (Experimental)

The `-swift-interop` flag enables experimental Swift interop mode. This generates bindings for calling Swift code using purego instead of cgo.

**Key Requirements:**
- Swift functions must use `@_cdecl` attribute for C compatibility
- Only simple types supported (primitives, pointers, C strings)
- For complex types, use Swift → ObjC bridge

**Example:**
See `examples/swift-interop/` for a complete working example.

## Migration

To use improved generator:
```bash
cp cmd/generate-framework-bindings/main.go cmd/generate-framework-bindings/main_original.go
mv cmd/generate-framework-bindings/main_improved.go cmd/generate-framework-bindings/main.go
```

See `/tmp/generator-analysis.md` for detailed analysis.

## Code Generation Flow

### Class Generation (`class.gen.go`)

The following diagram illustrates how a single class file is generated from Apple's documentation JSON:

```mermaid
flowchart TD
    Start([Generate class.gen.go]) --> FetchDocs[Fetch Documentation<br/>appledocs.Symbols fsys, framework]

    FetchDocs --> ParseClass[Parse Class Definition<br/>occ2go.ParseClassFromJSON]

    ParseClass --> CollectMembers{Collect Members}

    CollectMembers --> Properties[Parse Properties<br/>- Name, Type, Attributes<br/>- Read-only vs Read-write]
    CollectMembers --> Methods[Parse Methods<br/>- Instance methods<br/>- Class methods<br/>- Selectors & Parameters]
    CollectMembers --> Protocols[Parse Protocols<br/>- Protocol conformance]

    Properties --> PrepProps[prepareProperties<br/>- Filter duplicates<br/>- Map types<br/>- Generate accessors]

    Methods --> PrepMethods[prepareInstanceMethods<br/>- Deduplicate by selector<br/>- Map parameter types<br/>- Map return types]

    PrepProps --> TypeMapping[Type Mapping<br/>mapObjCTypeToGo]
    PrepMethods --> TypeMapping

    TypeMapping --> FrameworkCheck{Framework-specific?}
    FrameworkCheck -->|Yes| Qualify[Qualify with package<br/>CGRect → coregraphics.CGRect]
    FrameworkCheck -->|No| Direct[Direct mapping<br/>id → objc.ID]

    Qualify --> CollectImports[Collect Required Imports<br/>getClassRequiredImports]
    Direct --> CollectImports

    CollectImports --> Template[Render Template<br/>templates.txtar: class.gen.go]

    Template --> GenSections{Generate Sections}

    GenSections --> Header[Package & Imports<br/>- Package declaration<br/>- Import statements<br/>- Sync.Once for class singleton]

    GenSections --> ClassDef[Class Definition<br/>- Type alias<br/>- Interface definition<br/>- Struct embedding objectivec.Object]

    GenSections --> Constructors[Constructors<br/>- Alloc/Init pattern<br/>- Convenience constructors<br/>- Autorelease calls]

    GenSections --> MethodsOut[Instance Methods<br/>{{range prepareInstanceMethods}}<br/>- Selector calls via objc.Send<br/>- Type-safe parameters<br/>- Return value handling]

    GenSections --> PropsOut[Property Accessors<br/>{{range prepareProperties}}<br/>- Getter methods<br/>- Setter methods if read-write<br/>- Proper naming conventions]

    Header --> WriteFile[Write to generated/framework/<br/>classname.gen.go]
    ClassDef --> WriteFile
    Constructors --> WriteFile
    MethodsOut --> WriteFile
    PropsOut --> WriteFile

    WriteFile --> TestGen[Generate Test File<br/>classname.gen_test.go]

    TestGen --> End([Complete])

    style Start fill:#e1f5e1
    style End fill:#e1f5e1
    style TypeMapping fill:#fff4e1
    style Template fill:#e1f0ff
    style WriteFile fill:#ffe1f0
```

### Key Functions

1. **prepareInstanceMethods(methods []*ParsedMethod)**
   - Deduplicates methods by selector
   - Maps Objective-C types to Go types
   - Generates method signatures

2. **prepareProperties(properties []*ParsedProperty)**
   - Filters duplicate properties
   - Generates getter/setter pairs
   - Handles read-only attributes

3. **mapObjCTypeToGo(objcType, framework)**
   - Framework-aware type mapping
   - Handles cross-framework references
   - Qualifies types with package names

4. **getClassRequiredImports(class, framework)**
   - Scans methods and properties
   - Extracts package-qualified types
   - Returns sorted import list

### Current Known Issues

⚠️ **Duplicate Property/Method Generation** (appledocs-223)
- Some properties are generated twice in output
- Root cause: Properties may be appearing both in property list AND method list
- Affects: CloudKit (CreationDate, RecordID), MetalKit (Zone), others
- Status: Under investigation

### Template Structure

The `templates.txtar` file contains multiple templates:

```
-- class.gen.go --
{{- if .Methods}}
    {{range prepareInstanceMethods .Methods}}
    // Generate method
    {{end}}
{{end}}

{{- if .Properties}}
    {{range prepareProperties .Properties}}
    // Generate getter/setter
    {{end}}
{{end}}
```
