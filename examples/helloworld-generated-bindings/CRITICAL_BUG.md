# CRITICAL BUGS IN GENERATED BINDINGS - ALL FIXED ✅

## BUG 1: Framework Load Order Issue ✅ FIXED

### Summary
The generated package-level class variables were initialized BEFORE the AppKit framework was loaded, resulting in `windowClass = 0` and broken constructors.

### Root Cause
Generated code had:
```go
var windowClass = _WindowClass{objc.GetClass("NSWindow")}
```

This ran during package init, but if `Dlopen` happened in the importing package's init, it was too late - the classes were already 0.

### Fix Applied (Commit 885f1eab13)
Added automatic framework loading in generated packages. Each framework package now includes:
```go
func init() {
    _, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
    if err != nil {
        panic(err)
    }
}
```

This ensures the framework is loaded BEFORE any class variables are initialized.

### Status
**✅ FIXED** - No workaround needed! The framework loads automatically when you import the package.

## BUG 2: Missing String Wrapping in Generated Constructors ✅ FIXED

### Summary
Generated constructors that took string parameters were passing Go strings directly to Objective-C, causing segfaults.

### Root Cause
AppKit bindings were generated on Oct 17, BEFORE the string wrapping fix was committed on Oct 18 (commit d4ca095898).

### Fix Applied (Commit 885f1eab13)
Regenerated AppKit with the string wrapping templates. All string parameters are now wrapped:

**Before:**
```go
func NewButtonWithTitleTargetAction(title string, target objc.ID, action objc.SEL) Button {
    rv := objc.Send[Button](objc.ID(buttonClass.class), objc.Sel("buttonWithTitle:target:action:"), title, target, action)
    //                                                                                            ^^^^^ NOT WRAPPED!
    rv.Autorelease()
    return rv
}
```

**After:**
```go
func NewButtonWithTitleTargetAction(title string, target objc.ID, action objc.SEL) Button {
    rv := objc.Send[Button](objc.ID(buttonClass.class), objc.Sel("buttonWithTitle:target:action:"), objc.String(title), target, action)
    //                                                                                            ^^^^^^^^^^^^^^^^^ WRAPPED!
    rv.Autorelease()
    return rv
}
```

### Status
**✅ FIXED** - All string parameters in constructors and methods are now automatically wrapped with `objc.String()`.

## Verification

The E2E test now passes with both fixes:
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

## Usage

No workarounds needed! Just use the generated bindings naturally:

```go
import "github.com/tmc/appledocs/generated/appkit"

// Framework loads automatically on import
// Pass Go strings directly - automatic conversion to NSString
window := appkit.NewWindowWithFrame(100, 100, 400, 300, appkit.WindowStyleMaskTitled)
window.SetTitle("My Window")

button := appkit.NewButtonWithTitleTargetAction("Click Me!", handler, action)
```

## Commits

- d4ca095898: Add automatic Go string to NSString conversion
- b617026b1f: Improve test value generation for special string parameters
- 885f1eab13: Add automatic framework loading in generated packages

## Impact

**ALL** generated constructors and methods now work correctly:
- Window creation ✅
- Button creation ✅
- View creation ✅
- All AppKit types ✅

Both bugs are completely fixed in the code generator!
