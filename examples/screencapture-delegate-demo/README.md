# ScreenCaptureKit Delegate API Demo

Demonstrates type-safe delegate helpers for ScreenCaptureKit's `SCStreamOutput` protocol.

## Quick Start

```bash
cd examples/screencapture-delegate-demo
go run main.go
```

## Two Approaches

### Interface-Based (stateful handlers)

```go
type FrameCounter struct {
    count int
}

func (f *FrameCounter) StreamDidOutputSampleBuffer(
    stream screencapturekit.SCStream, buf uintptr, typ int) {
    f.count++
    fmt.Printf("Frame %d\n", f.count)
}

delegate := screencapturekit.NewSCStreamOutputDelegate(&FrameCounter{})
```

### Function-Based (simple callbacks)

```go
delegate := screencapturekit.NewSimpleSCStreamOutputDelegate(
    func(stream screencapturekit.SCStream, buf uintptr, typ int) {
        fmt.Printf("Frame received\n")
    },
)
```

## Benefits

- Type-safe parameters (`SCStream` vs `objc.ID`)
- No manual selector registration
- No `objc.RegisterClass` boilerplate
- Clean, idiomatic Go code

## See Also

- [ScreenCaptureKit Docs](https://developer.apple.com/documentation/screencapturekit/)
- [Delegate Pattern Details](../../CLAUDE.md)
