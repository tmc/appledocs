// Package main implements the Apple framework Go bindings generator.
//
// # Overview
//
// The generate-framework-bindings tool automatically generates comprehensive,
// type-safe Go bindings for Apple frameworks (AppKit, Foundation, CoreGraphics, etc.)
// using Apple's official documentation JSON as the source of truth.
//
// # Code Generation Pipeline
//
// The generation process flows through six stages:
//
// 1. Discovery & Loading
//   - Framework discovery with pattern matching (discoverFrameworks)
//   - Cross-framework type registry building (buildCrossFrameworkTypeRegistry)
//
// 2. Document Processing
//   - Symbol extraction from API collections (extractSymbolsFromAPICollections)
//   - Document parsing by symbol type (c:@F@ = function, c:objc(cs) = class, etc.)
//
// 3. Enum Value Enrichment
//   - Extract actual numeric values from macOS SDK headers (enrichEnumValues)
//   - Uses extract-enum-values tool via clang preprocessor
//
// 4. Type Resolution
//   - Framework-aware type mapping (mapObjCTypeToGo)
//   - Cross-framework type resolution (resolveType)
//   - Hierarchy: local types → registry → stdlib → fallback
//
// 5. Generation Preparation
//   - Reference type extraction
//   - Method grouping and constructor naming
//   - Inheritance chain resolution (SortClassesByDependency)
//   - Property override merging
//
// 6. Template Execution
//   - Load template archive (templates.txtar)
//   - Apply variant overlays if specified
//   - Two-tier template function registration:
//   - Core utilities (templateFuncs) - string manipulation, formatting
//   - Generator methods (GeneratorFuncs.Funcs()) - state-dependent operations
//   - Single execution path via GenerateTxtarFromModule()
//
// # Type System
//
// Types are resolved through framework-aware mappings:
//
//	Primitive types: BOOL → bool, NSInteger → int, CGFloat → float64
//	Object types:    id → objc.ID, NSString* → string, NSArray* → []T
//	Geometry types:  NSRect → foundation.Rect (in Foundation)
//	                 NSRect → coregraphics.CGRect (in AppKit)
//	Reference types: CGContextRef → coregraphics.CGContextRef
//
// # Generated Code Structure
//
// Each framework generates:
//
//	doc.gen.go         - Package documentation
//	types.gen.go       - Type definitions and enums
//	enums.gen.go       - Enumeration constants with SDK values
//	functions.gen.go   - C function wrappers
//	<class>.gen.go     - Individual class implementations
//	protocol.gen.go    - Protocol definitions
//	*_test.gen.go      - Test examples
//	objc/sel.go        - Cached selector lookups (~50x faster)
//
// # Usage
//
// Generate a single framework:
//
//	generate-framework-bindings -framework Foundation -output generated
//
// Pattern matching:
//
//	generate-framework-bindings -framework '^Core.*' -output generated
//
// With options:
//
//	generate-framework-bindings -framework AppKit \
//	    -output generated \
//	    -with-ref-methods \
//	    -generate-tests \
//	    -v
//
// # Architecture
//
// Key components:
//
//	main.go                - Entry point and pipeline orchestration
//	funcs.go               - Template function registration
//	funcs_core.go          - Core string utilities and template helpers
//	funcs_naming.go        - Name conversion (ObjC to Go conventions)
//	funcs_types.go         - Type mapping and resolution
//	funcs_methods.go       - Method processing and formatting
//	funcs_imports.go       - Import resolution and management
//	funcs_properties.go    - Property accessor generation
//	funcs_templates.go     - Template data preparation
//	funcs_class_helpers.go - Class-level helper functions
//	funcs_registry.go      - Type registry management
//	typemapping.go         - Objective-C to Go type mappings
//	templates.txtar        - Code generation templates
//	occ2go/                - Documentation parser library
//
// The funcs.go file has been refactored into focused modules for better
// maintainability and clarity. Each module handles a specific aspect of
// the code generation process.
//
// For detailed design documentation, see README.md.
package main
