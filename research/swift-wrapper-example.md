# Swift Wrapper Library Pattern for Go Bindings

## Problem

Pure Swift frameworks (Charts, SwiftUI) don't expose C-compatible APIs directly.

## Solution

Create a Swift wrapper library that exports @_cdecl functions, then bind to that from Go.

## Example: Wrapping a Swift Framework

### Step 1: Create Swift Wrapper

```swift
// ChartsWrapper.swift
import Charts
import Foundation

// Wrapper for Swift Chart API
@_cdecl("charts_create_bar_chart")
public func createBarChart() -> UnsafeMutableRawPointer? {
    // In a real implementation, you'd create the chart
    // and return an opaque pointer
    return nil
}

@_cdecl("charts_add_data_point")
public func addDataPoint(_ chart: UnsafeMutableRawPointer, _ x: Double, _ y: Double) {
    // Implementation
}

@_cdecl("charts_render")
public func renderChart(_ chart: UnsafeMutableRawPointer) -> UnsafeMutableRawPointer? {
    // Return rendered image data
    return nil
}

@_cdecl("charts_free")
public func freeChart(_ chart: UnsafeMutableRawPointer) {
    // Clean up
}
```

### Step 2: Compile Wrapper

```bash
# Compile as dynamic library
swiftc -emit-library \
  -o libchartswrapper.dylib \
  -emit-module \
  -module-name ChartsWrapper \
  ChartsWrapper.swift

# Or as static library for distribution
swiftc -emit-library \
  -static \
  -o libchartswrapper.a \
  ChartsWrapper.swift
```

### Step 3: Generate Go Bindings

```go
// chartswrapper.go
package charts

import (
    "github.com/ebitengine/purego"
    "unsafe"
)

var (
    lib uintptr

    chartsCreateBarChart func() uintptr
    chartsAddDataPoint   func(uintptr, float64, float64)
    chartsRender         func(uintptr) uintptr
    chartsFree           func(uintptr)
)

func init() {
    var err error
    lib, err = purego.Dlopen("libchartswrapper.dylib", purego.RTLD_NOW|purego.RTLD_GLOBAL)
    if err != nil {
        panic(err)
    }

    purego.RegisterLibFunc(&chartsCreateBarChart, lib, "charts_create_bar_chart")
    purego.RegisterLibFunc(&chartsAddDataPoint, lib, "charts_add_data_point")
    purego.RegisterLibFunc(&chartsRender, lib, "charts_render")
    purego.RegisterLibFunc(&chartsFree, lib, "charts_free")
}

type BarChart struct {
    ptr uintptr
}

func NewBarChart() *BarChart {
    ptr := chartsCreateBarChart()
    return &BarChart{ptr: ptr}
}

func (c *BarChart) AddDataPoint(x, y float64) {
    chartsAddDataPoint(c.ptr, x, y)
}

func (c *BarChart) Render() []byte {
    // Implementation
    return nil
}

func (c *BarChart) Free() {
    chartsFree(c.ptr)
}
```

### Step 4: Use from Go

```go
package main

import "github.com/tmc/appledocs/generated/charts"

func main() {
    chart := charts.NewBarChart()
    defer chart.Free()

    chart.AddDataPoint(1.0, 10.0)
    chart.AddDataPoint(2.0, 20.0)
    chart.AddDataPoint(3.0, 15.0)

    imageData := chart.Render()
    // Use imageData
}
```

## Automation Potential

We could build a tool that:

1. Parses Swift framework headers (.swiftinterface files)
2. Generates @_cdecl wrapper functions automatically
3. Compiles the wrapper library
4. Generates Go bindings for the wrapper

This would be similar to SWIG but for Swift→Go.

## Distribution

For frameworks like Charts:

1. Create `appledocs-swift-wrappers` repository
2. Provide pre-compiled wrapper dylibs
3. Go bindings import the wrapper
4. Users download both the bindings and the wrapper library

## Trade-offs

**Pros:**
- ✅ Enables access to pure Swift frameworks
- ✅ Type-safe Swift code (wrapper is in Swift)
- ✅ Go bindings use proven purego pattern
- ✅ No CGo required

**Cons:**
- ❌ Extra build step (compile Swift wrapper)
- ❌ Maintenance overhead (wrapper needs updates)
- ❌ Distribution complexity (ship dylibs)
- ❌ Manual mapping of Swift APIs

## Recommendation

Use this pattern for:
- High-value Swift frameworks (Charts, SwiftUI components)
- Frameworks without Objective-C APIs
- When official C API is unavailable

Don't use for:
- Frameworks with existing Objective-C APIs
- Low-level frameworks (use direct binding)
- Rapidly changing APIs (maintenance burden)
