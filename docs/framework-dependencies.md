# Framework Dependencies and Layering

## Investigation Summary

**Bead:** appledocs-589
**Date:** 2025-10-24
**Tool:** cmd/analyze-framework-deps

## Key Findings

### Circular Dependencies Are the Norm

Analysis of 189 Apple frameworks shows that **circular dependencies are ubiquitous** across the entire framework ecosystem:

- **Vision ↔ CoreImage**: VNBarcodeObservation uses CIBarcodeDescriptor
- **Foundation ↔ UIKit**: Extensive cross-references
- **CoreGraphics ↔ Metal**: Mutual dependencies
- **AVFoundation ↔ CoreImage**: Video processing types
- **All frameworks** end up in "Layer 99" (circular dependency group)

### Why This Matters

The original assumption was that frameworks follow a strict hierarchy (Foundation at level 1, AppKit at level 3, etc.). The data shows this is incorrect:

1. Frameworks reference types from multiple other frameworks
2. These references often form cycles
3. Apple handles this at the source level using protocols and opaque types

### Current Implementation Issues

The hardcoded `getFrameworkLevel()` function in `funcs_types.go:654` attempts to enforce a hierarchy that doesn't exist in practice. This causes:

- Incorrect import skipping decisions
- Build failures when the assumed hierarchy conflicts with reality
- Maintenance burden as frameworks evolve

## Recommendations

### 1. Simplify Import Logic

The package-based check at `funcs_imports.go:277` is **correct and sufficient**:

```go
if typePackage != currentPackage && typePackage != "objc" && typePackage != "objectivec" {
    // Skip import - code generator will use objc.IObject
    continue
}
```

This approach:
- ✅ Prevents circular imports
- ✅ Works with any dependency pattern
- ✅ Doesn't require maintaining framework levels
- ✅ Matches Apple's actual architecture

### 2. Remove Framework Level Checks

The hierarchy violation checks (lines 288-310 in `funcs_imports.go`) should be removed:

```go
// REMOVE THIS:
currentLevel := getFrameworkLevel(strings.ToLower(framework))
targetLevel := getFrameworkLevel(targetFramework)
if currentLevel >= 0 && targetLevel > currentLevel {
    shouldSkip = true
}
```

These checks don't reflect reality and add complexity without benefit.

### 3. Update Documentation

The architecture principle should be:

> **Cross-Framework Types**: When a type belongs to a different framework, always use the interface type (`objc.IObject`) rather than importing the concrete type. This prevents circular dependencies and matches Apple's protocol-based design.

## How Apple Handles This

Apple's frameworks use several patterns to manage circular dependencies:

1. **Protocols**: Define behavior contracts without concrete types
2. **Opaque Types**: Use `id` or protocol types at API boundaries
3. **Forward Declarations**: Declare classes without importing headers
4. **Runtime Binding**: Use dynamic dispatch and runtime type checking

Our bindings mirror this by converting cross-framework concrete types to `objc.IObject` at code generation time.

## Tool: analyze-framework-deps

Located in `cmd/analyze-framework-deps/`, this tool:

1. Scans Apple documentation in `~/.appledocs/cache/`
2. Extracts cross-framework type references from:
   - `preciseIdentifier` fields in tokens
   - `doc://` identifiers in relationships
   - Type fragments in declarations
3. Builds a dependency graph
4. Attempts topological sorting (which fails due to cycles)

### Usage

```bash
cd cmd/analyze-framework-deps
go build
./analyze-framework-deps
```

### Output Format

```
Framework Dependencies:
Vision -> [Foundation CoreImage CoreML AVFoundation]
CoreImage -> [QuartzCore com.apple.coreimage Foundation UIKit Metal AVFoundation CoreML]
...

Framework Layers:
Layer 99 (circular deps): [Vision CoreImage Foundation UIKit ...]
```

## Conclusion

The investigation confirms that:

1. ❌ Strict framework hierarchy doesn't exist
2. ✅ Circular dependencies are intentional and pervasive
3. ✅ Simple package-based import skipping is correct
4. ✅ Interface types (`objc.IObject`) are the right abstraction

The hardcoded `getFrameworkLevel()` function should be removed in favor of the simpler package-based approach.
