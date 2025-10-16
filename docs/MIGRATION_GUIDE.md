# Migration Guide: DarwinKit-Style Bindings

This guide helps you migrate from function-style bindings to the new DarwinKit-style object-oriented bindings.

## Table of Contents

1. [Overview](#overview)
2. [Key Differences](#key-differences)
3. [Migration Examples](#migration-examples)
4. [Memory Management](#memory-management)
5. [Type Safety](#type-safety)
6. [Common Patterns](#common-patterns)
7. [Troubleshooting](#troubleshooting)

## Overview

The appledocs generator now supports two binding styles:

### Function-Style (Original)
```go
// Old style: C-like functions
btn := appkit.NSButton_ButtonWithTitle(title, target, action)
appkit.NSButton_SetTitle(btn, "New Title")
```

### DarwinKit-Style (New)
```go
// New style: Object-oriented methods
btn := appkit.Button_ButtonWithTitle(title, target, action)
btn.SetTitle("New Title")
```

**Key Benefits of DarwinKit-Style:**
- ✅ Idiomatic Go with methods on types
- ✅ Better IDE autocomplete and type safety
- ✅ Matches progrium/darwinkit API for compatibility
- ✅ Easier to read and maintain
- ✅ Proper inheritance chain through interfaces
- ✅ One file per class (better organization)

## Key Differences

### 1. Naming Conventions

| Old (Function-Style) | New (DarwinKit-Style) |
|---------------------|----------------------|
| `NSButton` | `Button` (prefix stripped) |
| `appkit.NSButton_ButtonWithTitle()` | `appkit.Button_ButtonWithTitle()` |
| `appkit.NSButton_SetTitle(btn, title)` | `btn.SetTitle(title)` |
| Package: all in `classes.gen.go` | Package: `button.gen.go`, `window.gen.go`, etc. |

### 2. Import Statements

```go
// Old
import "github.com/tmc/appledocs/generated/frameworks/appkit"

// New (same import, different API)
import (
    "github.com/tmc/appledocs/generated/frameworks/appkit"
    "github.com/ebitengine/purego/objc"  // For objc.ID, objc.SEL if needed
)
```

### 3. Type System

```go
// Old: Functions taking objc.ID
btn := appkit.NSButton_Alloc()  // Returns objc.ID
appkit.NSButton_Init(btn)        // Takes objc.ID

// New: Methods on concrete types
btn := appkit.ButtonClass.Alloc()  // Returns appkit.Button
btn = btn.Init()                    // Method on Button type
// Or simply:
btn := appkit.NewButton()           // Convenience constructor
```

### 4. Class Variables

```go
// Old: No class variables
// (call functions directly)

// New: Explicit class variables
var btn appkit.Button
appkit.ButtonClass.New()  // Access class through ButtonClass variable
```

## Migration Examples

### Example 1: Creating a Button

**Before (Function-Style):**
```go
btn := appkit.NSButton_Alloc()
btn = appkit.NSButton_Init(btn)
appkit.NSButton_SetTitle(btn, "Click Me")
appkit.NSButton_SetFrame(btn, frame)
```

**After (DarwinKit-Style):**
```go
btn := appkit.NewButton()  // Alloc + Init + Autorelease
btn.SetTitle("Click Me")
btn.SetFrame(frame)
```

### Example 2: Factory Methods

**Before (Function-Style):**
```go
btn := appkit.NSButton_ButtonWithTitle("Click Me", nil, objc.RegisterName("action:"))
```

**After (DarwinKit-Style):**
```go
// Option 1: Class method
btn := appkit.ButtonClass.ButtonWithTitle("Click Me", nil, objc.RegisterName("action:"))

// Option 2: Module-level convenience
btn := appkit.Button_ButtonWithTitle("Click Me", nil, objc.RegisterName("action:"))
```

### Example 3: Window Creation

**Before (Function-Style):**
```go
rect := foundation.NSMakeRect(0, 0, 800, 600)
window := appkit.NSWindow_Alloc()
window = appkit.NSWindow_InitWithContentRect(
    window, rect,
    appkit.NSWindowStyleMaskTitled|appkit.NSWindowStyleMaskClosable,
    appkit.NSBackingStoreBuffered,
    false,
)
appkit.NSWindow_SetTitle(window, "My Window")
appkit.NSWindow_MakeKeyAndOrderFront(window, nil)
```

**After (DarwinKit-Style):**
```go
rect := foundation.MakeRect(0, 0, 800, 600)
window := appkit.WindowClass.Alloc().InitWithContentRect(
    rect,
    appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable,
    appkit.BackingStoreBuffered,
    false,
)
window.SetTitle("My Window")
window.MakeKeyAndOrderFront(nil)
```

### Example 4: Delegate Pattern

**Before (Function-Style):**
```go
// Manual delegate object creation was complex
delegate := createCustomDelegate()
appkit.NSApplication_SetDelegate(app, delegate)
```

**After (DarwinKit-Style):**
```go
// Use delegate builder pattern
delegate := &appkit.ApplicationDelegate{}
delegate.SetApplicationDidFinishLaunching(func(notification foundation.Notification) {
    // Handle app launch
})
delegate.SetApplicationWillTerminate(func(notification foundation.Notification) {
    // Cleanup
})
app.SetDelegate(delegate)
```

### Example 5: Property Access

**Before (Function-Style):**
```go
title := appkit.NSButton_Title(btn)
appkit.NSButton_SetTitle(btn, "New Title")

enabled := appkit.NSButton_IsEnabled(btn)
appkit.NSButton_SetEnabled(btn, true)
```

**After (DarwinKit-Style):**
```go
title := btn.Title()
btn.SetTitle("New Title")

enabled := btn.IsEnabled()
btn.SetEnabled(true)
```

## Memory Management

### Autorelease Pools

**Before (Function-Style):**
```go
// Manual autorelease pool management
pool := objc.AutoreleasePoolCreate()
defer objc.AutoreleasePoolDrain(pool)

// ... code ...
```

**After (DarwinKit-Style):**
```go
// Built-in autorelease on New() constructors
btn := appkit.NewButton()  // Automatically autoreleased

// Or explicit control
btn := appkit.ButtonClass.Alloc().Init()
btn.Autorelease()  // Manual autorelease if needed
```

### Retain/Release

Both styles support explicit memory management:

```go
// Both old and new
obj.Retain()   // Increment retain count
obj.Release()  // Decrement retain count
count := obj.RetainCount()  // Check retain count
```

**Best Practice:** Let the Objective-C autorelease pool handle memory management unless you have specific requirements.

## Type Safety

### Interface Types for Parameters

**Key Improvement:** DarwinKit-style uses interface types for parameters, enabling better polymorphism.

```go
// DarwinKit-style
type IButton interface {
    IControl
    SetTitle(value string)
    Title() string
}

// Accept any type implementing IButton
func ConfigureButton(btn appkit.IButton) {
    btn.SetTitle("Configured")
}

// Works with Button, CheckBox, RadioButton, etc.
ConfigureButton(myButton)
ConfigureButton(myCheckBox)
```

### Type Hierarchy

```go
// Interface hierarchy mirrors Objective-C class hierarchy
type IButton interface {
    IControl  // Inherits Control methods
}

type IControl interface {
    IView  // Inherits View methods
}

type IView interface {
    IResponder  // Inherits Responder methods
}

// Concrete type embeds parent
type Button struct {
    Control  // Embeds parent struct
}
```

## Common Patterns

### Pattern 1: Application Lifecycle

**DarwinKit-Style:**
```go
package main

import (
    "github.com/tmc/appledocs/generated/frameworks/appkit"
    "github.com/tmc/appledocs/generated/frameworks/foundation"
)

func main() {
    app := appkit.Application_SharedApplication()

    delegate := &appkit.ApplicationDelegate{}
    delegate.SetApplicationDidFinishLaunching(func(notification foundation.Notification) {
        // Setup window
        window := createMainWindow()
        window.MakeKeyAndOrderFront(nil)
    })

    app.SetDelegate(delegate)
    app.Run()
}

func createMainWindow() appkit.Window {
    rect := foundation.MakeRect(0, 0, 800, 600)
    window := appkit.WindowClass.Alloc().InitWithContentRect(
        rect,
        appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable|appkit.WindowStyleMaskResizable,
        appkit.BackingStoreBuffered,
        false,
    )
    window.SetTitle("My App")
    window.Center()
    return window
}
```

### Pattern 2: Custom Views

**DarwinKit-Style:**
```go
type CustomView struct {
    appkit.View
    // Custom fields
}

func NewCustomView(frame foundation.Rect) CustomView {
    view := appkit.ViewClass.Alloc().InitWithFrame(frame)
    return CustomView{View: view}
}

func (cv CustomView) DrawRect(dirtyRect foundation.Rect) {
    // Custom drawing code
    // Call super if needed: cv.View.DrawRect(dirtyRect)
}
```

### Pattern 3: Event Handling

**DarwinKit-Style:**
```go
btn := appkit.NewButton()
btn.SetTarget(nil)
btn.SetAction(objc.RegisterName("buttonClicked:"))

// Or use button with closure
btn = appkit.Button_ButtonWithTitle("Click", nil, objc.RegisterName("action:"))
```

## Troubleshooting

### Issue: "Type X does not implement interface Y"

**Solution:** Make sure you're using the interface type, not the concrete type:

```go
// Wrong
func SetView(view appkit.View) { }

// Correct
func SetView(view appkit.IView) { }
```

### Issue: "Cannot find class name"

**Cause:** DarwinKit-style strips the Objective-C prefix.

```go
// Old
appkit.NSButton_New()

// New
appkit.Button_New()  // or appkit.NewButton()
```

### Issue: "Method not found on type"

**Cause:** Using wrong receiver type.

```go
// Wrong - calling class method on instance
btn := appkit.NewButton()
btn.ButtonWithTitle("Hello")  // ❌ ButtonWithTitle is a class method

// Correct
btn := appkit.Button_ButtonWithTitle("Hello", nil, nil)  // ✅
```

### Issue: "Import cycle detected"

**Cause:** One file per class can create import cycles if not careful.

**Solution:** Use forward references and interfaces:

```go
// In button.gen.go
type IButton interface {
    IControl  // Reference parent interface
    // ...
}
```

## Migration Checklist

- [ ] Update import statements (same package, different API)
- [ ] Replace `NSButton` with `Button` (strip NS/CG/CF prefix)
- [ ] Convert function calls to method calls
- [ ] Use interface types (`IButton`) for parameters
- [ ] Replace `Alloc()+Init()` with `New()` constructors
- [ ] Update delegate patterns to use builder style
- [ ] Use concrete types for return values
- [ ] Test compilation and runtime behavior
- [ ] Update documentation and examples

## Next Steps

1. **Start Small:** Migrate one file at a time
2. **Test Thoroughly:** Ensure runtime behavior matches
3. **Use Types:** Leverage Go's type system with interfaces
4. **Read Examples:** Check `examples/` directory for patterns
5. **Consult Docs:** See [API_DESIGN_RATIONALE.md](./API_DESIGN_RATIONALE.md)

## Getting Help

- **Issues:** https://github.com/tmc/appledocs/issues
- **Examples:** `/examples/` directory
- **Reference:** DarwinKit project for additional patterns
- **Documentation:** Apple's Objective-C documentation remains relevant

## Compatibility

Both binding styles are generated and can coexist:
- Function-style: Use `-variant base` (default)
- DarwinKit-style: Use `-variant darwinkit`

Choose the style that fits your project best, or migrate gradually.
