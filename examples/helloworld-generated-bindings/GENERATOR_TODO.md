# Generator Fix - COMPLETED ✅

## Issue: Package-Level Class Variables Initialize Before Framework Load

### Problem (RESOLVED)

Even though commit 885f1eab13 added `init()` function to load frameworks, the generated code had:

```go
// In window.gen.go (and all other class files) - OLD CODE
var windowClass = _WindowClass{objc.GetClass("NSWindow")}
```

This line ran during package initialization, which happened **BEFORE** the `init()` function in doc.gen.go ran. Result: `windowClass.class = 0`.

### Fix Applied (Commit 651ab8a3f6)

The generator now uses **lazy initialization** with `sync.Once` for all class variables:

```go
var (
	windowClass     _WindowClass
	windowClassOnce sync.Once
)

func getWindowClass() _WindowClass {
	windowClassOnce.Do(func() {
		windowClass = _WindowClass{objc.GetClass("NSWindow")}
	})
	return windowClass
}

// All constructors now use lazy getter:
func NewWindowWithFrame(...) Window {
	instance := getWindowClass().Alloc()  // ✅ Lazy init!
	...
}
```

### Verification

E2E test passes with all fixes:
```bash
$ ./helloworld-generated-bindings -e2e
=== E2E Test Mode (Generated Bindings) ===
✓ Created window with title
✓ Got content view
✓ Created and configured label
✓ Created counter label
✓ Created button with target/action
✓ Button created and configured
✓ Label values set correctly
✓ Window closed

=== E2E Test PASSED ===
```

### Complete Fix Trilogy ✅

All three critical issues are now resolved:

1. **Framework Loading (885f1eab13)**: Automatic Dlopen in init()
2. **Lazy Initialization (651ab8a3f6)**: sync.Once for class variables
3. **String Conversion (d4ca095898)**: Automatic objc.String() wrapping

### Files Updated

- `cmd/generate-framework-bindings/templates.txtar`:
  - class.gen.go template - Added lazy init pattern
  - classes.gen.go template - Added lazy init pattern
- All generated `**/class*.gen.go` files regenerated with lazy init

### Impact

**ALL** generated frameworks now work correctly:
- ✅ Framework loads automatically on import
- ✅ Classes initialize lazily after framework load
- ✅ Constructors return valid objects (ID ≠ 0)
- ✅ String parameters convert automatically
- ✅ No workarounds needed!

### Status

**✅ COMPLETED** - No further action needed. Generator produces correct code.
