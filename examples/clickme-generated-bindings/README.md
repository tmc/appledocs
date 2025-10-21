# clickme-generated-bindings

A minimal (88-line) AppKit application demonstrating usage of generated Go bindings for macOS frameworks.

## Overview

This example shows how to:
- Create an NSApplication using generated bindings
- Build a simple window with UI controls
- Handle button clicks with Objective-C target/action pattern
- Run automated E2E tests using `PerformClick()`

**Total code: 88 lines** - fully inline, single main function, comprehensive E2E testing

## Features

- ✅ Uses generated AppKit bindings (no manual objc.Send in main code)
- ✅ Single code path for both interactive and E2E modes
- ✅ Actual functional testing (clicks button, verifies counter updates)
- ✅ Clean, inline code - easy to read top-to-bottom
- ✅ Automatic string conversion (pass Go strings directly)

## Building

```bash
go build
```

## Running

### Interactive Mode
```bash
go run .
# or
./clickme-generated-bindings
```
Creates a window with a "Click Me!" button. Each click increments a counter.

### E2E Test Mode
```bash
go run . -e2e
# or
./clickme-generated-bindings -e2e
```
Runs automated test that clicks the button twice and verifies the counter updates correctly.

## Key Features of Generated Bindings

1. **Type Safety**: All AppKit types are properly wrapped in Go structs
2. **Automatic String Conversion**: Pass Go strings directly, no manual NSString conversion needed
3. **Convenience Constructors**: `NewWindowWithFrame()`, `NewButtonWithFrame()`
4. **Type-Safe Methods**: `SetTitle()`, `SetTitleString()`, `AddSubviewTyped()`
5. **No Manual Framework Loading**: Bindings handle framework initialization

## Code Highlights

### Automatic String Conversion
```go
// No manual createNSString() needed - just pass Go strings directly!
window.SetTitle("Hello from Generated Bindings!")
counterLabel.SetStringValue("Clicks: 0")
counterLabel.SetStringValue(fmt.Sprintf("Clicks: %d", clickCount))
```

### Type-Safe Constructors
```go
// Type-safe constructor with proper parameter types
window := appkit.NewWindowWithFrame(100, 100, 400, 300,
    appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable)

// Convenience constructor for TextField
counterLabel = appkit.NewTextFieldWithFrame(100, 200, 200, 40)
```

### Button Handler
```go
func createButtonHandler() objc.ID {
    buttonClicked := func(self objc.ID, _cmd objc.SEL, sender objc.ID) {
        clickCount++
        // Automatic string conversion - just use Go strings!
        counterLabel.SetStringValue(fmt.Sprintf("Clicks: %d", clickCount))
    }
    // Register custom Objective-C class...
}
```

## Comparison

This is the generated bindings version of the clickme example. Compare with:
- `../clickme-purego/` - Pure purego/objc version (manual string conversion, lower-level)
- `../clickme-darwinkit/` - DarwinKit version (external dependency, different API)

See `../CLICKME_COMPARISON.md` for a detailed comparison.

## Key Advantages

1. **Automatic String Conversion**: Pass Go strings directly, no manual NSString creation
2. **Type Safety**: Proper Go types for all parameters and return values
3. **API Alignment**: Matches Apple's official documentation
4. **Local Dependency**: Part of this project, no external packages needed
5. **IDE Support**: Full autocomplete and type checking
