# ClickMe Purego Example

Simple GUI application demonstrating button creation using pure Go with purego.

## What This Example Demonstrates

- Creating a simple macOS window
- Creating and positioning an NSButton
- Adding views to a window's content view
- Simpler alternative to the helloworld example

## Building and Running

```bash
go build
./clickme-purego
```

For E2E testing:

```bash
./clickme-purego -e2e
```

## Key Translation from darwinkit

### Button Creation

**darwinkit**:
```go
button := appkit.NewButtonWithTitle("Click Me!")
button.SetFrameOrigin(foundation.Point{X: 150, Y: 120})
button.SetFrameSize(foundation.Size{Width: 100, Height: 40})
window.ContentView().AddSubview(button)
```

**purego**:
```go
button := objc.ID(objc.GetClass("NSButton")).Send(objc.RegisterName("alloc"))
buttonFrame := NSRect{
    Origin: NSPoint{X: 150, Y: 120},
    Size:   NSSize{Width: 100, Height: 40},
}
button = button.Send(objc.RegisterName("initWithFrame:"), buttonFrame)
buttonTitle := createNSString("Click Me!")
button.Send(objc.RegisterName("setTitle:"), buttonTitle)
contentView := window.Send(objc.RegisterName("contentView"))
objc.ID(contentView).Send(objc.RegisterName("addSubview:"), button)
```

## See Also

- Original: `/Volumes/tmc/go/src/github.com/progrium/darwinkit/macos/_examples/clickme`
- HelloWorld purego: `../helloworld-purego`
