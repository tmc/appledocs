# Best Practices for DarwinKit-Style Bindings

This guide provides best practices, patterns, and anti-patterns for using the generated DarwinKit-style bindings effectively.

## Table of Contents

1. [Memory Management](#memory-management)
2. [Type Usage](#type-usage)
3. [Error Handling](#error-handling)
4. [Concurrency](#concurrency)
5. [Performance](#performance)
6. [Testing](#testing)
7. [Common Pitfalls](#common-pitfalls)

## Memory Management

### ✅ DO: Use New() Constructors for Simple Cases

```go
// Good: Automatic memory management
btn := appkit.NewButton()
btn.SetTitle("Click Me")
// No explicit cleanup needed, autoreleased
```

### ✅ DO: Use Autorelease Pools for Tight Loops

```go
// Good: Explicit pool for intensive operations
for i := 0; i < 10000; i++ {
    pool := foundation.NewAutoreleasePool()

    // Create many temporary objects
    str := foundation.String_StringWithFormat("Item %d", i)
    processString(str)

    pool.Drain()  // Release temporary objects
}
```

### ❌ DON'T: Forget to Drain Autorelease Pools

```go
// Bad: Memory leak in long-running loops
for i := 0; i < 10000; i++ {
    pool := foundation.NewAutoreleasePool()
    // ... work ...
    // Missing pool.Drain()!
}
```

### ✅ DO: Retain Long-Lived Objects

```go
// Good: Explicitly retain objects you'll hold
type MyController struct {
    window appkit.Window
}

func (mc *MyController) SetWindow(w appkit.Window) {
    if mc.window.Ptr() != nil {
        mc.window.Release()
    }
    mc.window = w.Retain()  // Retain for ownership
}

func (mc *MyController) Close() {
    if mc.window.Ptr() != nil {
        mc.window.Release()
        mc.window = appkit.Window{}
    }
}
```

### ❌ DON'T: Over-Release Objects

```go
// Bad: Will cause crashes
btn := appkit.NewButton()
btn.Release()  // ❌ NewButton() already handles memory
// btn is now invalid!
```

## Type Usage

### ✅ DO: Accept Interfaces, Return Structs

```go
// Good: Flexible input, concrete output
func ConfigureView(view appkit.IView) appkit.View {
    view.SetHidden(false)
    view.SetNeedsDisplay(true)
    return appkit.ViewFrom(view.Ptr())
}
```

### ✅ DO: Use Interface Types for Parameters

```go
// Good: Works with Button, CheckBox, RadioButton, etc.
func StyleButton(btn appkit.IButton) {
    btn.SetBordered(true)
    btn.SetBezelStyle(appkit.BezelStyleRounded)
}
```

### ❌ DON'T: Use Concrete Types for Parameters

```go
// Bad: Only works with exact Button type
func StyleButton(btn appkit.Button) {
    // Won't work with CheckBox or other button subtypes
}
```

### ✅ DO: Check for Nil Before Using Objects

```go
// Good: Safe nil handling
func ProcessWindow(window appkit.Window) {
    if window.Ptr() == nil {
        log.Println("Window is nil")
        return
    }
    window.SetTitle("Processed")
}
```

### ✅ DO: Use Type Assertions Safely

```go
// Good: Safe conversion with check
if btn, ok := control.(appkit.IButton); ok {
    btn.SetTitle("As Button")
} else {
    log.Println("Not a button")
}
```

## Error Handling

### ✅ DO: Check Return Values

```go
// Good: Check for error conditions
window := appkit.WindowClass.Alloc().InitWithContentRect(...)
if window.Ptr() == nil {
    return fmt.Errorf("failed to create window")
}
```

### ✅ DO: Use Defer for Cleanup

```go
// Good: Ensure cleanup happens
func ProcessFile(path string) error {
    data := foundation.Data_DataWithContentsOfFile(path)
    if data.Ptr() == nil {
        return fmt.Errorf("failed to load file")
    }
    defer data.Release()  // Cleanup guaranteed

    // Process data...
    return nil
}
```

### ✅ DO: Validate Objective-C Responses

```go
// Good: Check for valid responses
delegate := app.Delegate()
if delegate.Ptr() == nil {
    log.Println("No delegate set")
} else if delegate.RespondsToSelector(objc.RegisterName("applicationWillTerminate:")) {
    // Safe to send message
}
```

## Concurrency

### ❌ DON'T: Access UI from Background Goroutines

```go
// Bad: Will crash on macOS
go func() {
    btn := appkit.NewButton()  // ❌ UI objects must be on main thread
    btn.SetTitle("Hello")
}()
```

### ✅ DO: Use Main Thread for UI Operations

```go
// Good: Dispatch UI work to main thread
go func() {
    // Background work
    result := computeExpensiveResult()

    // Update UI on main thread
    foundation.MainQueue().Async(func() {
        btn.SetTitle(result)
    })
}()
```

### ✅ DO: Protect Shared State

```go
// Good: Mutex for concurrent access
type SafeController struct {
    mu     sync.Mutex
    window appkit.Window
}

func (sc *SafeController) Window() appkit.Window {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    return sc.window
}

func (sc *SafeController) SetWindow(w appkit.Window) {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    if sc.window.Ptr() != nil {
        sc.window.Release()
    }
    sc.window = w.Retain()
}
```

## Performance

### ✅ DO: Cache Selectors

```go
// Good: Cache selectors at package level
var (
    sel_setTitle = objc.RegisterName("setTitle:")
    sel_title    = objc.RegisterName("title")
)

func updateButton(btn appkit.Button) {
    // Use cached selectors
    objc.Call[objc.Void](btn, sel_setTitle, "New Title")
}
```

### ❌ DON'T: Look Up Selectors Repeatedly

```go
// Bad: Expensive repeated lookups
for i := 0; i < 1000; i++ {
    sel := objc.RegisterName("setTitle:")  // ❌ Slow!
    objc.Call[objc.Void](btn, sel, "Title")
}
```

### ✅ DO: Batch UI Updates

```go
// Good: Minimize layout passes
view.BeginUpdate()  // If available
view.SetFrame(newFrame)
view.SetHidden(false)
view.SetAlphaValue(1.0)
view.EndUpdate()
```

### ✅ DO: Reuse Objects When Possible

```go
// Good: Reuse formatter
formatter := foundation.DateFormatter_New()
formatter.SetDateFormat("yyyy-MM-dd")

for _, date := range dates {
    str := formatter.StringFromDate(date)
    processString(str)
}
formatter.Release()
```

### ❌ DON'T: Create Unnecessary Temporary Objects

```go
// Bad: Creating NSString for each iteration
for i := 0; i < 10000; i++ {
    str := foundation.String_StringWithFormat("Item %d", i)  // ❌ Expensive
    // Better to use Go string formatting when possible
}

// Good: Use Go primitives where applicable
for i := 0; i < 10000; i++ {
    str := fmt.Sprintf("Item %d", i)  // ✅ Faster
}
```

## Testing

### ✅ DO: Use Interfaces for Mocking

```go
// Good: Mockable interface
type ButtonLike interface {
    appkit.IButton
    SetTitle(string)
    Title() string
}

type MockButton struct {
    title string
}

func (mb *MockButton) SetTitle(t string) {
    mb.title = t
}

func (mb *MockButton) Title() string {
    return mb.title
}

// Test uses mock
func TestConfigureButton(t *testing.T) {
    mock := &MockButton{}
    ConfigureButton(mock)
    if mock.Title() != "Configured" {
        t.Error("Button not configured")
    }
}
```

### ✅ DO: Test Memory Management

```go
// Good: Verify retain counts
func TestWindowRetention(t *testing.T) {
    window := appkit.NewWindow()
    initialCount := window.RetainCount()

    controller := NewController()
    controller.SetWindow(window)

    if window.RetainCount() != initialCount+1 {
        t.Error("Window not retained")
    }

    controller.Close()

    if window.RetainCount() != initialCount {
        t.Error("Window not released")
    }
}
```

### ✅ DO: Use Test Fixtures

```go
// Good: Reusable test setup
type AppKitTestFixture struct {
    app      appkit.Application
    window   appkit.Window
    delegate *appkit.ApplicationDelegate
}

func NewAppKitTestFixture() *AppKitTestFixture {
    return &AppKitTestFixture{
        app:    appkit.Application_SharedApplication(),
        window: createTestWindow(),
        delegate: &appkit.ApplicationDelegate{},
    }
}

func (f *AppKitTestFixture) Cleanup() {
    if f.window.Ptr() != nil {
        f.window.Close()
        f.window = appkit.Window{}
    }
}

func TestWithFixture(t *testing.T) {
    fixture := NewAppKitTestFixture()
    defer fixture.Cleanup()

    // Test using fixture...
}
```

## Common Pitfalls

### Pitfall 1: Not Understanding Autorelease

```go
// ❌ Problem: Over-retaining autoreleased objects
func createButton() appkit.Button {
    btn := appkit.NewButton()  // Already autoreleased
    btn.Retain()               // ❌ Unnecessary, will leak
    return btn
}

// ✅ Solution: Trust autorelease
func createButton() appkit.Button {
    return appkit.NewButton()  // Correct, autoreleased
}

// ✅ Alternative: Explicit control
func createManagedButton() appkit.Button {
    btn := appkit.ButtonClass.Alloc().Init()  // Not autoreleased
    // Caller must Release() when done
    return btn
}
```

### Pitfall 2: Mixing Class and Instance Methods

```go
// ❌ Problem: Calling class method on instance
btn := appkit.NewButton()
btn.ButtonWithTitle("Hello")  // ❌ Compile error, ButtonWithTitle is class method

// ✅ Solution: Use correct receiver
btn := appkit.Button_ButtonWithTitle("Hello", nil, nil)  // Class method
btn.SetTitle("New Title")  // Instance method
```

### Pitfall 3: Forgetting Thread Safety

```go
// ❌ Problem: Concurrent UI updates
var window appkit.Window

func initWindow() {
    window = appkit.NewWindow()
}

func updateWindow() {
    go func() {
        window.SetTitle("Updated")  // ❌ Not thread-safe!
    }()
}

// ✅ Solution: Serialize UI updates
func updateWindow() {
    foundation.MainQueue().Async(func() {
        window.SetTitle("Updated")  // ✅ On main thread
    })
}
```

### Pitfall 4: Not Checking Nil Pointers

```go
// ❌ Problem: Assuming objects are valid
func processWindow(w appkit.Window) {
    w.SetTitle("Title")  // ❌ May crash if w is nil
}

// ✅ Solution: Always validate
func processWindow(w appkit.Window) {
    if w.Ptr() == nil {
        return
    }
    w.SetTitle("Title")
}
```

### Pitfall 5: Incorrect Type Conversions

```go
// ❌ Problem: Unsafe type assertions
func getButton(view appkit.View) appkit.Button {
    return view.(appkit.Button)  // ❌ May panic
}

// ✅ Solution: Safe conversions
func getButton(view appkit.IView) (appkit.Button, bool) {
    if btn, ok := view.(appkit.IButton); ok {
        return appkit.ButtonFrom(btn.Ptr()), true
    }
    return appkit.Button{}, false
}
```

## Quick Reference Checklist

**Memory Management:**
- [ ] Use `New()` for simple cases
- [ ] Use `Alloc().Init()` for explicit control
- [ ] Drain autorelease pools in loops
- [ ] Retain long-lived objects
- [ ] Release what you retain

**Type Safety:**
- [ ] Accept interfaces, return structs
- [ ] Check for nil pointers
- [ ] Use safe type assertions
- [ ] Validate Objective-C responses

**Concurrency:**
- [ ] UI operations on main thread
- [ ] Protect shared state with mutexes
- [ ] Use `foundation.MainQueue()` for UI updates

**Performance:**
- [ ] Cache selectors
- [ ] Reuse expensive objects
- [ ] Batch UI updates
- [ ] Profile before optimizing

**Testing:**
- [ ] Mock using interfaces
- [ ] Test memory management
- [ ] Use test fixtures
- [ ] Clean up after tests

## Additional Resources

- [MIGRATION_GUIDE.md](./MIGRATION_GUIDE.md) - Migrating from old bindings
- [API_DESIGN_RATIONALE.md](./API_DESIGN_RATIONALE.md) - Design decisions explained
- [Apple Memory Management Guide](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/MemoryMgmt/)
- [Effective Go](https://go.dev/doc/effective_go)

## Getting Help

If you encounter issues:
1. Check this guide for common pitfalls
2. Review example code in `/examples/`
3. Search existing issues on GitHub
4. Open a new issue with minimal reproduction

Remember: When in doubt, follow Go idioms and trust the autorelease pool!
