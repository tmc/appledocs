# Click Me - Pure Purego Version

This example demonstrates creating a simple macOS window with a clickable button using pure purego/objc without any high-level bindings.

## Features

- Creates a window with title "Hello from Purego!"
- Displays a "Click Me!" button
- Shows a counter label that updates each time the button is clicked
- Uses only `github.com/ebitengine/purego` - no other dependencies
- Demonstrates custom Objective-C class registration for button handling

## Building

```bash
go build
```

## Running

```bash
./clickme-purego
```

Click the button and watch the counter increment!

## E2E Testing

```bash
./clickme-purego -e2e
```

## Code Highlights

### Manual String Conversion
```go
func createNSString(s string) objc.ID {
    nsStringClass := objc.GetClass("NSString")
    str := objc.ID(nsStringClass).Send(objc.RegisterName("alloc"))
    return str.Send(objc.RegisterName("initWithUTF8String:"), objc.RegisterName(s))
}
```

### Custom Button Handler
```go
func createButtonHandler() objc.ID {
    // Register a new Objective-C class with a buttonClicked: method
    class, _ = objc.RegisterClass("ButtonHandler", superClass, nil, nil, []objc.MethodDef{
        {Cmd: objc.RegisterName("buttonClicked:"), Fn: buttonClicked},
    })
    // ...
}
```

## Comparison

This is the pure purego version of the clickme example. Compare with:
- `../clickme-darwinkit/` - DarwinKit version (high-level Go API)
- `../clickme-generated-bindings/` - Generated bindings version (type-safe, auto string conversion)

See `../CLICKME_COMPARISON.md` for a detailed comparison.

## Learning Value

This example is excellent for:
- Understanding how Objective-C runtime works
- Learning the low-level details of AppKit
- Debugging framework behavior
- Implementing features not yet available in bindings
