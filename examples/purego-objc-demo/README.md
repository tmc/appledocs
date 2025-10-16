# Purego/ObjC Direct Usage Demo

This example demonstrates using `github.com/ebitengine/purego/objc` directly to interact with Objective-C frameworks, without needing a custom runtime layer.

## Key Findings

### What purego/objc Provides (Already Production-Ready)

1. **Core Types**: `objc.ID`, `objc.Class`, `objc.SEL`, `objc.IMP`
2. **Message Sending**:
   - `id.Send(sel, args...)` - returns `objc.ID`
   - `objc.Send[T](id, sel, args...)` - generic typed returns
   - `objc.SendSuper[T](id, sel, args...)` - super calls
3. **Runtime Operations**:
   - `objc.GetClass(name)` - class lookup
   - `objc.RegisterName(selector)` - selector registration (caches internally)
   - `objc.RegisterClass(...)` - custom class creation
4. **Low-Level FFI**: All objc_msgSend variants, struct returns, proper ABIs

### What We Need to Add

For darwinkit-style ergonomics, we only need thin wrappers:

1. **objc/conversions.go**: Type conversion helpers
   - `ToNSString(s string) objc.ID`
   - `ToGoString(id objc.ID) string`
   - `ToNSArray(slice []any) objc.ID`
   - `ToGoSlice[T](id objc.ID) []T`

2. **objc/doc.go**: Package documentation

3. **Templates**: Generate method wrappers using purego/objc directly

## Revised Architecture

```
Generated Bindings (generated/frameworks/appkit/)
├── import "github.com/ebitengine/purego/objc"  ← Direct import
├── classes.gen.go     ← Interfaces + Structs
└── functions.gen.go   ← C functions (existing)

Thin Helpers (objc/)
├── conversions.go     ← String/Array/Dict conversions
└── doc.go            ← Package docs
```

## Example Generated Code (Future)

```go
// Button interface wraps NSButton
type Button interface {
    objc.ID  // Embed objc.ID for Send() method

    // Title returns the button's title
    Title() string

    // SetTitle sets the button's title
    SetTitle(title string)
}

type button struct {
    objc.ID
}

var (
    ButtonClass = objc.GetClass("NSButton")
    sel_title = objc.RegisterName("title")
    sel_setTitle = objc.RegisterName("setTitle:")
)

func (b button) Title() string {
    nsString := b.Send(sel_title)
    return ToGoString(nsString)  // From objc/conversions.go
}

func (b button) SetTitle(title string) {
    nsString := ToNSString(title)  // From objc/conversions.go
    b.Send(sel_setTitle, nsString)
}
```

## Benefits

- ✅ No custom FFI layer to maintain
- ✅ Battle-tested purego implementation
- ✅ Minimal wrapper code
- ✅ Compatible with purego ecosystem
- ✅ Focus on ergonomics, not plumbing
