# Click Me - DarwinKit Version

This example demonstrates creating a simple macOS window with a clickable button using the [DarwinKit](https://github.com/progrium/darwinkit) library.

## Features

- Creates a window with title "Hello from DarwinKit!"
- Displays a "Click Me!" button
- Shows a counter label that updates each time the button is clicked
- Uses DarwinKit's high-level, type-safe Go API
- Demonstrates custom button handler with target/action pattern

## Building

```bash
go build
```

## Running

```bash
./clickme-darwinkit
```

## E2E Testing

```bash
./clickme-darwinkit -e2e
```

## Code Highlights

### Type-Safe Label Creation
```go
counterLabel = appkit.NewTextFieldWithFrame(foundation.Rect{
    Origin: foundation.Point{X: 100, Y: 200},
    Size:   foundation.Size{Width: 200, Height: 40},
})
counterLabel.SetStringValue("Clicks: 0")
```

### Button with Target/Action
```go
button := appkit.Button_ButtonWithTitleTargetAction(
    "Click Me!",
    darwinkitObjc.ObjectFrom(unsafe.Pointer(buttonHandler)),
    darwinkitObjc.Sel("buttonClicked:"),
)
```

## Comparison

This is the DarwinKit version of the clickme example. Compare with:
- `../clickme-purego/` - Pure purego/objc version (low-level, manual)
- `../clickme-generated-bindings/` - Generated bindings version (type-safe, local)

See `../CLICKME_COMPARISON.md` for a detailed comparison
