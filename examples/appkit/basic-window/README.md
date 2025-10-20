# AppKit Basic Window Example

This example demonstrates creating and displaying a basic window using AppKit in Go.

## What it demonstrates

- Getting the shared NSApplication instance
- Creating an NSWindow with frame and style
- Configuring window properties (title, visibility)
- Querying window state
- Basic window operations (center, show, close)
- Proper main thread handling for AppKit

## Running the example

```bash
go run main.go
```

The window will appear for 3 seconds and then close automatically.

## Key Concepts

### NSApplication

NSApplication is the central controller for macOS applications. Every AppKit app needs one:

```go
// Get shared application instance
sharedAppSel := objc.RegisterName("sharedApplication")
appClass := objc.GetClass("NSApplication")
appID := objc.ID(appClass).Send(sharedAppSel)
```

### NSWindow Creation

Create a window with frame, style, and backing store:

```go
windowRect := coregraphics.CGRect{
    Origin: coregraphics.CGPoint{X: 100, Y: 100},
    Size:   coregraphics.CGSize{Width: 400, Height: 300},
}

styleMask := appkit.WindowStyleMaskTitled |
             appkit.WindowStyleMaskClosable |
             appkit.WindowStyleMaskResizable

window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
    windowRect,
    styleMask,
    appkit.BackingStoreBuffered,
    false,
)
```

### Window Style Masks

Common style mask flags:
- `WindowStyleMaskTitled` - Window has a title bar
- `WindowStyleMaskClosable` - Window can be closed
- `WindowStyleMaskMiniaturizable` - Window can be minimized
- `WindowStyleMaskResizable` - Window can be resized
- `WindowStyleMaskFullScreen` - Window supports fullscreen
- `WindowStyleMaskBorderless` - Window has no border

Combine with bitwise OR: `mask1 | mask2 | mask3`

### Backing Store Types

- `BackingStoreBuffered` - Standard buffered backing (use this)
- `BackingStoreRetained` - Legacy, deprecated
- `BackingStoreNonretained` - Legacy, deprecated

### Showing the Window

```go
// Set title
setTitleSel := objc.RegisterName("setTitle:")
window.ID.Send(setTitleSel, titleString)

// Show window and make it key (active)
makeKeyAndOrderFrontSel := objc.RegisterName("makeKeyAndOrderFront:")
window.ID.Send(makeKeyAndOrderFrontSel, nil)
```

## Main Thread Requirement

**CRITICAL:** AppKit MUST run on the main thread!

```go
func main() {
    runtime.LockOSThread()  // Lock to main thread
    defer runtime.UnlockOSThread()

    // AppKit code here...
}
```

Without this, you'll get crashes or undefined behavior.

## Window Coordinate System

macOS uses a bottom-left origin coordinate system:
- Origin (0,0) is at the bottom-left of the screen
- X increases to the right
- Y increases upward

```
(0, screenHeight) ──────────── (screenWidth, screenHeight)
      │                                  │
      │        Screen                    │
      │                                  │
(0, 0) ──────────────────────── (screenWidth, 0)
```

## Event Loop

In a real application, you must run the event loop:

```go
// Start the event loop (this blocks)
runSel := objc.RegisterName("run")
appID.Send(runSel)
```

This example doesn't run the event loop, so it just shows the window briefly and exits.

## Common Window Operations

### Centering

```go
centerSel := objc.RegisterName("center")
window.ID.Send(centerSel)
```

### Moving

```go
setFrameDisplaySel := objc.RegisterName("setFrame:display:")
window.ID.Send(setFrameDisplaySel, newFrame, true)
```

### Resizing

```go
setContentSizeSel := objc.RegisterName("setContentSize:")
window.ID.Send(setContentSizeSel, newSize)
```

### Minimizing

```go
minimizeSel := objc.RegisterName("miniaturize:")
window.ID.Send(minimizeSel, nil)
```

### Closing

```go
closeSel := objc.RegisterName("close")
window.ID.Send(closeSel)
```

## Window Levels

Windows can be placed at different levels:
- `NSNormalWindowLevel` (0) - Normal windows
- `NSFloatingWindowLevel` - Floating palette windows
- `NSStatusWindowLevel` - Status/menu bar windows
- `NSPopUpMenuWindowLevel` - Pop-up menus
- `NSModalPanelWindowLevel` - Modal dialogs
- `NSMainMenuWindowLevel` - Main menu bar

## Use Cases

### Document Windows

```go
// Standard document window
window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
    frame,
    appkit.WindowStyleMaskTitled |
    appkit.WindowStyleMaskClosable |
    appkit.WindowStyleMaskMiniaturizable |
    appkit.WindowStyleMaskResizable,
    appkit.BackingStoreBuffered,
    false,
)
```

### Utility Windows

```go
// Floating utility window
window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
    frame,
    appkit.WindowStyleMaskTitled |
    appkit.WindowStyleMaskClosable |
    appkit.WindowStyleMaskUtilityWindow,
    appkit.BackingStoreBuffered,
    false,
)
```

### Borderless Windows

```go
// Custom borderless window
window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
    frame,
    appkit.WindowStyleMaskBorderless,
    appkit.BackingStoreBuffered,
    false,
)
```

## Real Application Structure

A complete AppKit application would include:

1. **Application Delegate**:
   - Handle app lifecycle events
   - Manage application state

2. **Window Delegate**:
   - Handle window events (close, resize, etc.)
   - Implement window behavior

3. **Event Loop**:
   - Run NSApplication event loop
   - Process user input

4. **Views and Controls**:
   - Add NSView subclasses to window
   - Add buttons, text fields, etc.

5. **Menu Bar**:
   - Create application menu
   - Add menu items

## Next Steps

After understanding basic windows, explore:
- Adding views (`NSView`)
- Adding controls (`NSButton`, `NSTextField`)
- Handling events (mouse, keyboard)
- Creating menus (`NSMenu`)
- Working with delegates

## References

- [NSWindow Documentation](https://developer.apple.com/documentation/appkit/nswindow)
- [NSApplication Documentation](https://developer.apple.com/documentation/appkit/nsapplication)
- [AppKit App Architecture](https://developer.apple.com/documentation/appkit/app_and_environment)
- [Window Programming Guide](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/WinPanel/Introduction.html)
