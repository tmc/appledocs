# ScreenCaptureKit Example - Generated Bindings

This example demonstrates using ScreenCaptureKit with generated Go bindings from appledocs.

## Features

- ✅ Explicit ScreenCaptureKit framework loading via `purego.Dlopen()`
- ✅ Async block-based completion handlers using `objc.NewBlock()`
- ✅ SCShareableContent enumeration (displays and windows)
- ✅ Proper error handling with localized error messages
- ✅ CoreGraphics fallback for display capture
- ✅ Framework loading diagnostics with `-test-loading` flag

## Requirements

- macOS 12.3+ (ScreenCaptureKit was introduced in Monterey)
- Go 1.24+
- **Screen Recording permission** (see below)

## Building

```bash
go build
```

## Running

### E2E Test Mode (No TCC required)

```bash
./screencapturekit-generated-bindings -e2e
```

This uses CoreGraphics for display capture and doesn't require Screen Recording permission.

### Interactive Mode (Requires TCC)

```bash
./screencapturekit-generated-bindings
```

This attempts to use SCShareableContent to enumerate displays and windows.

### Framework Loading Test

```bash
./screencapturekit-generated-bindings -test-loading
```

Verifies that ScreenCaptureKit framework loads correctly and all classes are available.

## Granting Screen Recording Permission

ScreenCaptureKit requires Screen Recording permission. There are two ways to grant it:

### Option 1: Manual (Recommended)

1. Run the example once - it will fail with permission error
2. Open **System Settings** → **Privacy & Security** → **Screen Recording**
3. Enable permission for the `screencapturekit-generated-bindings` executable
4. Run the example again

### Option 2: Using tccutil (Resets permission prompt)

```bash
tccutil reset ScreenCapture
```

Then run the example - macOS will show the permission prompt.

## Implementation Details

### Framework Loading

ScreenCaptureKit must be explicitly loaded before use:

```go
func init() {
    runtime.LockOSThread()

    // Explicitly load ScreenCaptureKit framework
    _, err := purego.Dlopen(
        "/System/Library/Frameworks/ScreenCaptureKit.framework/ScreenCaptureKit",
        purego.RTLD_NOW|purego.RTLD_GLOBAL)
    if err != nil {
        // Framework not available, will be handled at runtime
    }
}
```

### Async Completion Handlers

ScreenCaptureKit uses async APIs with Objective-C blocks. Use `objc.NewBlock()`:

```go
completionBlock := objc.NewBlock(
    func(block objc.Block, content objc.ID, error objc.ID) {
        defer func() { done <- true }()

        if error != 0 {
            // Handle error
            desc := error.Send(objc.RegisterName("localizedDescription"))
            errStr := objc.Send[string](desc, objc.RegisterName("UTF8String"))
            fmt.Printf("Error: %s\n", errStr)
            return
        }

        // Process content...
    },
)
defer completionBlock.Release()

// Call async method
sel := objc.RegisterName("getShareableContentWithCompletionHandler:")
objc.ID(shareableContentClass).Send(sel, completionBlock)

// Wait for completion
select {
case <-done:
    // Success
case <-time.After(5 * time.Second):
    // Timeout
}
```

### Error Handling

The example demonstrates proper error message extraction from NSError:

```go
if error != 0 {
    desc := error.Send(objc.RegisterName("localizedDescription"))
    if desc != 0 {
        errStr := objc.Send[string](desc, objc.RegisterName("UTF8String"))
        fmt.Printf("Error: %s\n", errStr)
    }
}
```

## Troubleshooting

### "SCShareableContent class not found"

**Cause**: ScreenCaptureKit framework not loaded
**Solution**: The example now loads it explicitly in `init()`. If you still see this, your macOS version may be older than 12.3.

### "The user declined TCCs for application, window, display capture"

**Cause**: Screen Recording permission not granted
**Solution**: Grant permission via System Settings → Privacy & Security → Screen Recording

### Classes load but still get permission error

This is expected! The code is working correctly. You just need to grant the TCC permission.

## Generated Bindings

This example uses bindings generated from Apple's ScreenCaptureKit documentation:

- `SCShareableContent` - System content enumeration
- `SCDisplay` - Display information
- `SCWindow` - Window information
- `SCRunningApplication` - Application information
- `SCContentFilter` - Content filtering
- `SCStream` - Real-time screen capture
- `SCStreamConfiguration` - Stream settings
- `SCScreenshotManager` - Screenshot capture

All bindings are in `../../generated/screencapturekit/`.

## References

- [ScreenCaptureKit Documentation](https://developer.apple.com/documentation/screencapturekit)
- [purego](https://github.com/ebitengine/purego)
- [objc.NewBlock](https://pkg.go.dev/github.com/ebitengine/purego/objc#NewBlock)
