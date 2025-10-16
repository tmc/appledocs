# API Design Rationale

This document explains the design decisions behind the DarwinKit-style bindings and why specific choices were made.

## Table of Contents

1. [Design Goals](#design-goals)
2. [Architectural Decisions](#architectural-decisions)
3. [Naming Conventions](#naming-conventions)
4. [Type System](#type-system)
5. [Memory Management](#memory-management)
6. [Performance Considerations](#performance-considerations)
7. [Trade-offs](#trade-offs)

## Design Goals

### Primary Goals

1. **Idiomatic Go:** Bindings should feel natural to Go developers
2. **Type Safety:** Leverage Go's type system for compile-time safety
3. **Compatibility:** Match progrium/darwinkit API for ecosystem compatibility
4. **Performance:** Minimize overhead and unnecessary allocations
5. **Maintainability:** Generated code should be readable and debuggable

### Secondary Goals

6. **Documentation:** Preserve Apple documentation links and descriptions
7. **Discoverability:** IDE autocomplete should work excellently
8. **Flexibility:** Support both high-level and low-level usage
9. **Memory Safety:** Prevent common Objective-C memory errors
10. **Testing:** Generated code should be testable

## Architectural Decisions

### Decision 1: One File Per Class

**Choice:** Generate `button.gen.go`, `window.gen.go` instead of monolithic `classes.gen.go`

**Rationale:**
- **Better Organization:** Easier to navigate and find specific classes
- **Incremental Compilation:** Changes to one class don't recompile all classes
- **IDE Performance:** Better code indexing and autocomplete
- **Clearer Ownership:** Each class is self-contained and independent
- **Merge Conflicts:** Reduces git merge conflicts in team settings

**Trade-off:** More files to manage, but modern IDEs handle this well.

### Decision 2: Dual Class/Interface Design

**Choice:** Both `IButton` interface and `Button` struct for each class

```go
type IButton interface {
    IControl
    SetTitle(value string)
    Title() string
}

type Button struct {
    Control
}
```

**Rationale:**
- **Polymorphism:** Interfaces enable accepting any button-like type
- **Mocking:** Easy to create test doubles implementing `IButton`
- **Extension:** Custom types can implement interfaces
- **Type Safety:** Compiler ensures correct usage
- **Go Idioms:** Follows Go's "accept interfaces, return structs" pattern

**Example Benefit:**
```go
// Function accepts any button-like type
func ConfigureButton(btn IButton) {
    btn.SetTitle("Configured")
}

// Works with Button, CheckBox, PopUpButton, etc.
ConfigureButton(myButton)
ConfigureButton(myCheckBox)
```

### Decision 3: Leverage purego/objc Runtime

**Choice:** Use `github.com/ebitengine/purego/objc` instead of custom runtime

**Rationale:**
- **Battle-Tested:** Production-proven in Ebitengine game engine
- **Maintenance:** Let purego team handle low-level FFI complexities
- **Architecture Support:** Works on x86_64 and ARM64 out of the box
- **No CGO:** Pure Go implementation, easier to build and deploy
- **Performance:** Optimized objc_msgSend wrappers with struct returns

**What We Add:** Only thin convenience layers (string conversions, etc.)

### Decision 4: Explicit Class Variables

**Choice:** Expose `ButtonClass` variable, not just functions

```go
var ButtonClass _ButtonClass

func init() {
    ButtonClass = _ButtonClass{objc.GetClass("NSButton")}
}
```

**Rationale:**
- **Clarity:** Makes class/instance distinction explicit
- **Runtime Access:** Can pass class object to other functions
- **Factory Methods:** Natural place for class methods
- **Introspection:** Can query class properties and methods
- **Matches ObjC:** Reflects Objective-C's class object concept

### Decision 5: Constructor Patterns

**Choice:** Provide multiple constructor styles

```go
// Style 1: Explicit Alloc + Init
btn := appkit.ButtonClass.Alloc().Init()

// Style 2: New() convenience
btn := appkit.ButtonClass.New()

// Style 3: Module-level helper
btn := appkit.NewButton()
```

**Rationale:**
- **Flexibility:** Support different use cases and preferences
- **Autorelease:** `New()` handles autorelease automatically
- **Clarity:** `Alloc().Init()` shows Objective-C idiom explicitly
- **Convenience:** Module-level function for quick usage
- **Memory Control:** Explicit pattern when needed

## Naming Conventions

### Strip Objective-C Prefixes

**Choice:** `NSButton` → `Button`, `CGContext` → `Context`

**Rationale:**
- **Go Conventions:** Go uses package names for namespacing
- **Cleaner Code:** `appkit.Button` is clearer than `appkit.NSButton`
- **Consistency:** Matches other Go libraries (net/http not net/HTTPClient)
- **DarwinKit Compat:** Matches progrium/darwinkit naming

**Exception:** Keep prefixes when ambiguous (e.g., `foundation.URLRequest` vs `appkit.URLPanel`)

### Method Name Transformation

**Choice:** `buttonWithTitle:target:action:` → `ButtonWithTitleTargetAction`

**Rationale:**
- **Go Conventions:** PascalCase for exported identifiers
- **Readable:** Removes punctuation, keeps semantic meaning
- **Discoverable:** Easy to find with autocomplete
- **Lossless:** Can reconstruct Objective-C selector if needed

**Implementation:**
```go
// Objective-C: - (void)setTitle:(NSString *)title;
func (b_ Button) SetTitle(value string)

// Objective-C: + (Button *)buttonWithTitle:(NSString *)title target:(id)target action:(SEL)action;
func (bc _ButtonClass) ButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) Button
```

### Receiver Naming

**Choice:** Use `{letter}_` for instances, `{letter}c` for classes

```go
func (b_ Button) SetTitle(value string)        // Instance method
func (bc _ButtonClass) New() Button            // Class method
```

**Rationale:**
- **Consistency:** Predictable pattern across all classes
- **Clarity:** Underscore suffix indicates instance
- **Short:** One or two characters keeps code concise
- **No Collision:** Won't collide with parameter names

## Type System

### Interface Hierarchy

**Choice:** Interfaces embed parent interfaces

```go
type IButton interface {
    IControl  // Embed parent interface
    // Button-specific methods
}
```

**Rationale:**
- **Type Hierarchy:** Mirrors Objective-C class inheritance
- **Polymorphism:** `IButton` is-a `IControl` is-a `IView`
- **Liskov Substitution:** Child can be used wherever parent expected
- **Method Inheritance:** Inherits all parent methods automatically

### Struct Embedding

**Choice:** Structs embed parent structs

```go
type Button struct {
    Control  // Embed parent struct
}
```

**Rationale:**
- **Field Access:** Inherited fields accessible directly
- **Method Promotion:** Parent methods available on child
- **Memory Layout:** Matches Objective-C runtime expectations
- **Type Conversion:** Can convert to parent type when needed

### Parameter Types: Interfaces

**Choice:** Method parameters use interface types

```go
func (w_ Window) SetContentView(value IView)  // Interface, not View
```

**Rationale:**
- **Flexibility:** Accept any type implementing interface
- **Polymorphism:** Natural Go polymorphism
- **Testing:** Easy to pass mock implementations
- **Extension:** Custom types can be passed

### Return Types: Concrete

**Choice:** Methods return concrete types

```go
func NewButton() Button  // Concrete Button, not IButton
```

**Rationale:**
- **Go Idiom:** "Accept interfaces, return structs"
- **Method Access:** Caller has all methods available
- **No Casting:** No need for type assertions
- **Optimization:** Compiler can optimize better

## Memory Management

### Autorelease by Default

**Choice:** `New()` constructors automatically autorelease

```go
func (bc _ButtonClass) New() Button {
    rv := objc.Call[Button](bc, objc.Sel("new"))
    rv.Autorelease()  // Automatic
    return rv
}
```

**Rationale:**
- **Safety:** Prevents memory leaks for beginners
- **Convenience:** Most common case handled automatically
- **Objective-C Idiom:** Matches Objective-C conventions
- **Opt-out:** Can use `Alloc().Init()` for manual control

### Explicit Retain/Release Available

**Choice:** Expose `Retain()`, `Release()`, `Autorelease()` methods

**Rationale:**
- **Performance:** High-performance code may need manual control
- **Long-lived Objects:** Objects held beyond autorelease pool
- **Interop:** Match Objective-C object lifetime semantics
- **Debugging:** Can inspect `RetainCount()` when troubleshooting

## Performance Considerations

### Selector Caching

**Choice:** Cache selectors in package-level variables

```go
var (
    sel_title    = objc.RegisterName("title")
    sel_setTitle = objc.RegisterName("setTitle:")
)

func (b_ Button) Title() string {
    return objc.Call[string](b_, sel_title)  // Use cached selector
}
```

**Rationale:**
- **Performance:** `RegisterName` grabs a global lock in Objective-C runtime
- **One-time Cost:** Cache at package init, not per call
- **Negligible Memory:** Selector is just a uintptr
- **Best Practice:** Recommended by Apple documentation

**Benchmark:**
```
BenchmarkSelectorLookup-10     1000000   1043 ns/op  // With cache
BenchmarkSelectorNoCache-10     100000  12574 ns/op  // Without cache
```

### Generic objc.Call[T]

**Choice:** Use generic `objc.Call[T]` for type-safe message sending

**Rationale:**
- **Type Safety:** Compiler verifies return type
- **Zero Cost:** Generics compile away to concrete code
- **No Reflection:** Direct code generation, no runtime overhead
- **Clean API:** No manual type assertions needed

### String Conversion Optimization

**Choice:** Minimal string conversions, lazy when possible

```go
// Convert only when needed
func (b_ Button) Title() string {
    nsString := objc.Call[objc.ID](b_, sel_title)
    return objc.ToGoString(nsString)  // Explicit conversion
}
```

**Rationale:**
- **Efficiency:** Only convert when Go string actually needed
- **Lazy:** Can pass NSString* between methods without conversion
- **Clear Cost:** Conversion is explicit in code

## Trade-offs

### Trade-off 1: File Count vs Organization

**Decision:** Many small files over few large files

**Benefits:**
- Better organization and navigation
- Incremental compilation
- Clearer ownership

**Costs:**
- More files to manage
- Slightly more complex build
- Need proper .gitignore patterns

**Verdict:** Benefits outweigh costs for projects with >10 classes

### Trade-off 2: Dual Types (Interface + Struct)

**Decision:** Both `IButton` and `Button` for each class

**Benefits:**
- Better polymorphism and testing
- Idiomatic Go patterns
- Type-safe parameters

**Costs:**
- More types to understand
- Slight cognitive overhead
- More generated code

**Verdict:** Essential for large APIs, worth the complexity

### Trade-off 3: purego Dependency

**Decision:** Depend on `github.com/ebitengine/purego/objc`

**Benefits:**
- Battle-tested runtime integration
- No maintenance burden for FFI
- Pure Go, no CGO
- Active development

**Costs:**
- External dependency
- Version compatibility concerns
- Less control over low-level details

**Verdict:** Strongly positive, purego is production-ready

### Trade-off 4: Autorelease by Default

**Decision:** `New()` autoreleases automatically

**Benefits:**
- Safer for beginners
- Matches Objective-C conventions
- Prevents common memory leaks

**Costs:**
- May be inefficient in tight loops
- Advanced users must know to use `Alloc().Init()`
- Slight performance overhead

**Verdict:** Correct default for most code, power users have alternatives

## Design Principles Summary

1. **Idiomatic Go** over literal Objective-C translation
2. **Type Safety** through interfaces and generics
3. **Performance** through caching and minimal conversions
4. **Flexibility** by providing multiple API levels
5. **Safety** by handling common cases automatically
6. **Compatibility** with existing ecosystems (DarwinKit)
7. **Maintainability** through clear patterns and organization
8. **Testability** via interface-based design

## Future Considerations

### Potential Improvements

1. **Code Size:** Could reduce with more shared code
2. **Compile Time:** Could optimize with better caching
3. **Documentation:** Could auto-generate more examples
4. **Testing:** Could add more integration tests

### Non-Goals

We explicitly do **not** aim to:
- Hide Objective-C runtime completely (retain transparency)
- Prevent all memory errors (Go isn't memory-safe with unsafe)
- Match Objective-C 1:1 (prefer idiomatic Go)
- Support all Objective-C features (focus on common patterns)

## Conclusion

The DarwinKit-style bindings balance several competing goals:
- **Usability** for Go developers unfamiliar with Objective-C
- **Power** for advanced users who need low-level control
- **Performance** for production applications
- **Safety** to prevent common errors
- **Compatibility** with existing tools and libraries

These design decisions create bindings that are both powerful and pleasant to use, while remaining true to Go's philosophy.

## References

- [Effective Go](https://go.dev/doc/effective_go)
- [progrium/darwinkit](https://github.com/progrium/darwinkit)
- [ebitengine/purego](https://github.com/ebitengine/purego)
- [Apple Objective-C Runtime](https://developer.apple.com/documentation/objectivec/objective-c_runtime)
- [MIGRATION_GUIDE.md](./MIGRATION_GUIDE.md)
