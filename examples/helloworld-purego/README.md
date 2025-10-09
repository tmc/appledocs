# HelloWorld Purego Example

Pure Go port of the darwinkit helloworld example using purego and objc runtime.

## What This Example Demonstrates

- Loading Foundation, AppKit, and WebKit frameworks without cgo
- Creating and configuring an NSApplication
- Creating a WKWebView and loading a URL
- Creating an NSWindow and displaying it
- Custom NSApplicationDelegate using objc.RegisterClass
- Handling window close events
- E2E testing support

## Building and Running

```bash
go build
./helloworld-purego
```

For E2E testing (terminates immediately):

```bash
./helloworld-purego -e2e
```

## Key Translation Patterns from darwinkit

### 1. Framework Loading

**darwinkit** (implicit via cgo):
```go
import "github.com/progrium/darwinkit/macos/webkit"
```

**purego** (explicit):
```go
purego.Dlopen("/System/Library/Frameworks/WebKit.framework/WebKit", ...)
```

### 2. NSRect Structures

**darwinkit**:
```go
frame := foundation.Rect{Size: foundation.Size{1440, 900}}
```

**purego**:
```go
frame := NSRect{
    Origin: NSPoint{X: 0, Y: 0},
    Size:   NSSize{Width: 1440, Height: 900},
}
```

### 3. Method Calls

**darwinkit** (high-level wrappers):
```go
app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
app.ActivateIgnoringOtherApps(true)
```

**purego** (objc runtime):
```go
sharedApp.Send(objc.RegisterName("setActivationPolicy:"), 0)
sharedApp.Send(objc.RegisterName("activateIgnoringOtherApps:"), true)
```

### 4. Custom Delegates

**darwinkit**:
```go
delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
    return true
})
```

**purego**:
```go
class, err := objc.RegisterClass(
    "AppDelegate",
    objc.GetClass("NSObject"),
    nil,
    nil,
    []objc.MethodDef{
        {
            Cmd: objc.RegisterName("applicationShouldTerminateAfterLastWindowClosed:"),
            Fn:  func(self objc.ID, _cmd objc.SEL, sender objc.ID) bool { return true },
        },
    },
)
```

## Architecture Differences

- **darwinkit**: Uses cgo with high-level Go wrappers for Objective-C classes
- **purego**: Uses FFI to call Objective-C runtime directly without cgo
- **Translation effort**: Significant - requires understanding objc runtime message sending

## See Also

- Original: `/Volumes/tmc/go/src/github.com/progrium/darwinkit/macos/_examples/helloworld`
- Purego window example: `/Volumes/tmc/go/src/github.com/ebitengine/purego/examples/window`
- Foundation working example: `../foundation-working`
