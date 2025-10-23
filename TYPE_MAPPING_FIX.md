# Type Mapping Fix Summary

## Issue

Two critical type mapping bugs were causing incorrect Go function signatures:

### 1. Float Return Type Bug
**Problem**: C `float` return types mapped to `unsafe.Pointer`

**Example**:
```go
// BEFORE (incorrect)
func CGColorGetContentHeadroom(color ColorRef) unsafe.Pointer

// AFTER (correct)
func CGColorGetContentHeadroom(color ColorRef) float32
```

**Root Cause**:
- `mapCTypeToGoWithFramework("float", "CoreGraphics")` called `occ2go.MapCTypeToGo` → `"float32"`
- Then passed `"float32"` through `mapObjCTypeToGo` which called `MapCTypeToGo` again
- `MapCTypeToGo("float32", ...)` returned `"unsafe.Pointer"` (unrecognized C type)

### 2. Pointer Parameter Bug
**Problem**: `const CGFloat *` parameters mapped to `unsafe.Pointer`

**Example**:
```go
// BEFORE (incorrect)
func CGColorCreate(space ColorSpaceRef, components unsafe.Pointer) ColorRef

// AFTER (correct)
func CGColorCreate(space ColorSpaceRef, components []float64) ColorRef
```

**Root Cause**:
- All pointer types (`*`) were blanket-mapped to `unsafe.Pointer`
- No special handling for pointer-to-primitive types

## Solution

### Fix 1: Prevent Double-Processing of Primitives
**File**: `cmd/generate-framework-bindings/funcs_types.go`

Added check in `mapCTypeToGoWithFramework` to return primitives and slices directly:

```go
// If the result is a Go primitive type or slice, return directly
goPrimitives := map[string]bool{
    "string": true,
    "int": true, "float32": true, "float64": true,
    // ... etc
}
if goPrimitives[goType] || strings.HasPrefix(goType, "[]") {
    return goType  // Don't reprocess
}
```

### Fix 2: Map Pointer-to-Primitives to Slices
**File**: `occ2go/typemap.go`

Enhanced pointer handling to recognize primitive patterns:

```go
case strings.Contains(cType, "*"):
    baseType := strings.TrimSpace(strings.ReplaceAll(
        strings.ReplaceAll(cType, "const", ""), "*", ""))
    switch baseType {
    case "CGFloat":    return "[]float64"
    case "float":      return "[]float32"
    case "double":     return "[]float64"
    case "int":        return "[]int"
    case "uint32_t":   return "[]uint32"
    case "uint64_t":   return "[]uint64"
    default:           return "unsafe.Pointer"
    }
```

## Impact

### Affected Functions (Sample)

**Float return types fixed**:
- `CGColorGetContentHeadroom(ColorRef) float32`
- `CGGradientGetContentHeadroom(GradientRef) float32`
- `CGImageGetContentHeadroom(ImageRef) float32`
- `CGImageCalculateContentAverageLightLevel(ImageRef) float32`

**Pointer parameters fixed to slices**:
- `CGColorCreate(ColorSpaceRef, []float64) ColorRef`
- `CGColorCreateWithPattern(ColorSpaceRef, PatternRef, []float64) ColorRef`
- `CGContextSetFillColor(ContextRef, []float64)`
- `CGContextSetStrokeColor(ContextRef, []float64)`
- `CGContextSetFillPattern(ContextRef, PatternRef, []float64)`
- `CGContextSetStrokePattern(ContextRef, PatternRef, []float64)`
- `CGFunctionCreate(unsafe.Pointer, uintptr, []float64, uintptr, []float64, unsafe.Pointer) FunctionRef`
- `CGGradientCreateWithColorComponents(ColorSpaceRef, []float64, []float64, uintptr) GradientRef`
- `CGGradientCreateWithColors(ColorSpaceRef, unsafe.Pointer, []float64) GradientRef`
- `CGColorSpaceCreateICCBased(uintptr, []float64, DataProviderRef, ColorSpaceRef) ColorSpaceRef`

**Other pointer types fixed**:
- `CGFontGetGlyphAdvances(FontRef, unsafe.Pointer, uintptr, []int) bool`
- `CGPDFDocumentGetVersion(PDFDocumentRef, []int, []int)`
- `CGGetActiveDisplayList(uint32, unsafe.Pointer, []uint32) unsafe.Pointer`

### Total Impact
- 10+ CoreGraphics functions now have correct signatures
- More idiomatic Go code (slices instead of unsafe.Pointer)
- Better type safety
- Easier to use without manual unsafe conversions

## Testing

Created comprehensive signature verification tests in `color_create_manual_test.go`:

```bash
$ cd generated/coregraphics && go test -v -run TestManual
=== RUN   TestManualCGColorCreateSignature
  ✓ Parameter 1: []float64 (correct slice type!)
  ✓ CGColorCreate has correct signature: func(ColorSpaceRef, []float64) ColorRef
--- PASS: TestManualCGColorCreateSignature (0.00s)

=== RUN   TestManualCGColorGetContentHeadroomSignature
  ✓ Return type: float32 (correct float32, not unsafe.Pointer!)
  ✓ CGColorGetContentHeadroom has correct signature: func(ColorRef) float32
--- PASS: TestManualCGColorGetContentHeadroomSignature (0.00s)

=== RUN   TestManualCGColorCreateWithContentHeadroomSignature
  ✓ Parameter 0 (headroom): float32 (correct float32!)
--- PASS: TestManualCGColorCreateWithContentHeadroomSignature (0.00s)

=== RUN   TestManualSliceParameterFunctions
  ✓ CGColorCreateWithPattern parameter 2 is []float64
  ✓ CGContextSetFillColor parameter 1 is []float64
  ✓ CGContextSetStrokeColor parameter 1 is []float64
--- PASS: TestManualSliceParameterFunctions (0.00s)

PASS
ok      github.com/tmc/appledocs/generated/coregraphics 0.525s
```

## Files Changed

- `occ2go/typemap.go` - Enhanced pointer-to-primitive mapping
- `cmd/generate-framework-bindings/funcs_types.go` - Prevent double-processing
- `generated/coregraphics/*.gen.go` - Regenerated with correct types
- `generated/coregraphics/color_create_manual_test.go` - Verification tests

## Commit

```
commit b3b73dc8bc
fix(typemap): Map pointer-to-primitive C types to Go slices
```
