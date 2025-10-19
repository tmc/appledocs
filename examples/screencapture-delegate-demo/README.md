# ScreenCaptureKit Delegate API Demo

This example demonstrates the type-safe delegate API for ScreenCaptureKit's `SCStreamOutput` protocol.

## Overview

Instead of manually registering Objective-C classes and handling low-level callbacks, the generated bindings now provide high-level delegate helpers that:

- Provide type-safe parameters (e.g., `screencapturekit.SCStream` instead of `objc.ID`)
- Eliminate selector name construction
- Remove class registration boilerplate
- Offer cleaner, more idiomatic Go code

## Running the Demo

```bash
cd examples/screencapture-delegate-demo
go run main.go
```

## What This Demo Shows

### 1. Interface-Based Delegate (Recommended)

Create a struct that implements the `SCStreamOutputHandler` interface:

```go
type MyHandler struct {
    frameCount int
}

func (h *MyHandler) StreamDidOutputSampleBuffer(stream screencapturekit.SCStream, buf uintptr, typ int) {
    h.frameCount++
    fmt.Printf("Frame %d\n", h.frameCount)
}

delegate := screencapturekit.NewSCStreamOutputDelegate(&MyHandler{})
```

**Best for:**
- Complex delegates with state
- Delegates that need to implement multiple methods
- Reusable delegate logic

### 2. Simple Function-Based Delegate

For quick, simple callbacks:

```go
delegate := screencapturekit.NewSimpleSCStreamOutputDelegate(
    func(stream screencapturekit.SCStream, buf uintptr, typ int) {
        fmt.Printf("Received frame\n")
    },
)
```

**Best for:**
- Simple, stateless callbacks
- Quick prototyping
- One-off delegate uses

### 3. Custom Logic Example

The frame recorder example shows how to build complex state machines:

```go
type FrameRecorder struct {
    maxFrames  int
    frameCount int
    frames     []uintptr
}

func (r *FrameRecorder) StreamDidOutputSampleBuffer(...) {
    r.frameCount++
    if r.frameCount <= r.maxFrames {
        r.frames = append(r.frames, sampleBuffer)
    }
}

recorder := &FrameRecorder{maxFrames: 10}
delegate := screencapturekit.NewSCStreamOutputDelegate(recorder)
```

## API Comparison

### Old Way (Manual)

```go
frameCount := 0
streamDidOutputSampleBuffer := func(self objc.ID, cmd objc.SEL, stream objc.ID, buf uintptr, typ int) {
    frameCount++
    fmt.Printf("Frame %d\n", frameCount)
}

class, _ := objc.RegisterClass(
    "MyDelegate",
    objc.GetClass("NSObject"),
    []*objc.Protocol{screencapturekit.SCStreamOutputProtocol},
    nil,
    []objc.MethodDef{{
        Cmd: objc.RegisterName("stream:didOutputSampleBuffer:ofType:"),
        Fn:  streamDidOutputSampleBuffer,
    }},
)
delegate := objc.ID(class).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
```

### New Way (Type-Safe)

```go
frameCount := 0
delegate := screencapturekit.NewSimpleSCStreamOutputDelegate(
    func(stream screencapturekit.SCStream, buf uintptr, typ int) {
        frameCount++
        fmt.Printf("Frame %d\n", frameCount)
    },
)
```

## Benefits

✓ **Type Safety** - Parameters use proper types (`SCStream` instead of `objc.ID`)
✓ **No Manual Selectors** - Selector names are handled internally
✓ **No Boilerplate** - Automatic class registration and initialization
✓ **Cleaner Code** - More idiomatic Go patterns
✓ **Flexible** - Can still use `objc.RegisterClass` for advanced cases

## Implementation Details

The delegate helpers are implemented in `generated/screencapturekit/helpers.go`:

- `SCStreamOutputHandler` interface - Defines the delegate methods
- `NewSCStreamOutputDelegate()` - Creates a delegate from a handler interface
- `NewSimpleSCStreamOutputDelegate()` - Creates a delegate from a simple function

These helpers automatically:
1. Generate unique class names to avoid conflicts
2. Register the class with the SCStreamOutput protocol
3. Create and initialize an instance
4. Convert objc.ID parameters to type-safe Go types

## See Also

- [ScreenCaptureKit Documentation](https://developer.apple.com/documentation/screencapturekit/)
- [Project CLAUDE.md](../../CLAUDE.md) - Detailed documentation on delegate patterns
- [Session 044C](https://github.com/tmc/appledocs/issues/044C) - Original async delegate implementation work
