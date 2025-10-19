# Window Visibility Bug Fix - State Restoration Issue

## Problem

macOS State Restoration was attempting to restore windows with an invalid HUD style mask (0x2000), which caused `canBecomeKeyWindow` to return NO, making windows invisible.

### Symptoms

```
(AppKit) [com.apple.AppKit:Window] NSWindow does not support HUD styleMask 0x2000; use NSPanel instead.
(AppKit) [com.apple.AppKit:Window] Warning: -[NSWindow makeKeyWindow] called on <NSWindow: 0x...> which returned NO from -[NSWindow canBecomeKeyWindow].
```

## Root Cause

When macOS State Restoration is enabled and the app previously had windows open, macOS tries to restore those windows automatically. However, the restoration process was incorrectly setting the HUD style mask (0x2000) on NSWindow objects, which is only valid for NSPanel.

## Solution

We implemented a **multi-layered defense** against unwanted state restoration:

### Method 1: Disable Relaunch-on-Login (Primary Fix)

```go
app.ID.Send(objc.RegisterName("disableRelaunchOnLogin"))
```

This single call completely disables both:
- Automatic relaunch on login
- Automatic window state restoration

This is the **primary fix** that resolved the HUD style mask error.

### Method 2: NSApplicationDelegate (Defense in Depth)

We also registered a custom NSApplicationDelegate with three methods to prevent state restoration:

```go
delegateClass, err := objc.RegisterClass(
    "AppDelegate",
    objc.GetClass("NSObject"),
    nil, // ivars
    nil, // properties
    []objc.MethodDef{
        {
            // Modern macOS 12+ method
            Cmd: objc.RegisterName("applicationSupportsSecureRestorableState:"),
            Fn: func(self objc.ID, _cmd objc.SEL, app objc.ID) bool {
                return false // Disable secure state restoration
            },
        },
        {
            // Legacy method for older macOS versions
            Cmd: objc.RegisterName("application:shouldRestoreApplicationState:"),
            Fn: func(self objc.ID, _cmd objc.SEL, app objc.ID, coder objc.ID) bool {
                return false // Never restore application state
            },
        },
        {
            // Prevent automatic window restoration
            Cmd: objc.RegisterName("applicationShouldAutomaticallyRestoreWindows:"),
            Fn: func(self objc.ID, _cmd objc.SEL, app objc.ID) bool {
                return false // Prevent automatic window restoration
            },
        },
    },
)
```

### Method 3: Per-Window Restoration Disable

For individual windows, we also set:

```go
window.ID.Send(objc.RegisterName("setRestorable:"), false)
```

## Results

After implementing this fix:

✅ **NO HUD/RESTORE ERRORS** in system logs
✅ E2E tests pass
✅ Windows become key properly
✅ Windows are visible and responsive

## Implementation Location

- **File**: `/Volumes/tmc/go/src/github.com/tmc/appledocs/examples/helloworld-generated-bindings/main.go`
- **Lines**: 67-121 (delegate creation and state restoration disabling)

## Key Learnings

1. **Order Matters**: `disableRelaunchOnLogin()` must be called after `SharedApplication()` but before window creation
2. **Defense in Depth**: Multiple layers of protection ensure the fix works across different macOS versions
3. **Delegate Methods**: While the delegate methods provide additional protection, `disableRelaunchOnLogin()` is the primary fix
4. **purego-Specific**: This issue is specific to purego-based bindings; CGO-based frameworks like darwinkit may handle state restoration differently

## Testing

Run the example:
```bash
./helloworld-generated-bindings
```

Check system logs for errors:
```bash
log show --last 1m --predicate 'processImagePath CONTAINS "helloworld"' | grep -iE "hud|0x2000|canBecomeKey"
```

Expected result: **No output** (no errors)

## References

- [NSApplication disableRelaunchOnLogin](https://developer.apple.com/documentation/appkit/nsapplication/1428747-disablerelaunchonlogin)
- [NSApplicationDelegate State Restoration](https://developer.apple.com/documentation/appkit/nsapplicationdelegate/app_state_restoration)
- [Window Restoration Guide](https://developer.apple.com/library/archive/documentation/General/Conceptual/MOSXAppProgrammingGuide/CoreAppDesign/CoreAppDesign.html#//apple_ref/doc/uid/TP40010543-CH3-SW35)
