// Hello World using only generated bindings (no darwinkit)
//
// Demonstrates using generated AppKit bindings with purego/objc.
// Press Cmd+Q to quit.
package main

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

func init() {
	runtime.LockOSThread()
	_, err := purego.Dlopen("/System/Library/Frameworks/AppKit.framework/AppKit", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

// Geometry types (would come from foundation package)
type NSPoint struct{ X, Y float64 }
type NSSize struct{ Width, Height float64 }
type NSRect struct {
	Origin NSPoint
	Size   NSSize
}

// Helper to create NSString
func nsString(s string) objc.ID {
	return objc.ID(objc.GetClass("NSString")).Send(objc.RegisterName("stringWithUTF8String:"), s)
}

var (
	clickCount   int
	counterLabel appkit.TextField
)

func createButtonHandler() objc.ID {
	className := "ButtonHandler"
	class := objc.GetClass(className)
	if class == 0 {
		superClass := objc.GetClass("NSObject")
		buttonClicked := func(self objc.ID, _cmd objc.SEL, sender objc.ID) {
			clickCount++
			fmt.Printf("Button clicked! Count: %d\n", clickCount)
			counterLabel.ID.Send(objc.RegisterName("setStringValue:"), nsString(fmt.Sprintf("Clicks: %d", clickCount)))
		}
		class, _ = objc.RegisterClass(className, superClass, nil, nil, []objc.MethodDef{
			{Cmd: objc.RegisterName("buttonClicked:"), Fn: buttonClicked},
		})
	}
	handler := objc.ID(class).Send(objc.RegisterName("alloc"))
	return handler.Send(objc.RegisterName("init"))
}

func main() {
	fmt.Println("=== Hello World (Generated Bindings) ===\n")

	// Get NSApplication shared instance (using generated ApplicationFrom)
	appClass := objc.GetClass("NSApplication")
	app := appkit.ApplicationFrom(unsafe.Pointer(objc.ID(appClass).Send(objc.RegisterName("sharedApplication"))))
	app.ID.Send(objc.RegisterName("setActivationPolicy:"), 0) // Regular app

	// Create window using generated constructor
	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		unsafe.Pointer(&NSRect{Origin: NSPoint{100, 100}, Size: NSSize{400, 300}}),
		1|2|8, // Titled, Closable, Resizable
		2,     // Buffered
		false,
	)
	window.ID.Send(objc.RegisterName("setTitle:"), nsString("Hello from Generated Bindings!"))

	// Get content view (using generated ViewFrom)
	contentView := appkit.ViewFrom(unsafe.Pointer(window.ID.Send(objc.RegisterName("contentView"))))

	// Create and configure label
	textFieldClass := objc.GetClass("NSTextField")
	label := appkit.TextFieldFrom(unsafe.Pointer(
		objc.ID(textFieldClass).Send(objc.RegisterName("alloc")).Send(
			objc.RegisterName("initWithFrame:"),
			NSRect{Origin: NSPoint{50, 200}, Size: NSSize{300, 50}},
		)))
	label.ID.Send(objc.RegisterName("setStringValue:"), nsString("Using generated bindings!"))
	label.ID.Send(objc.RegisterName("setEditable:"), false)
	label.ID.Send(objc.RegisterName("setBordered:"), false)
	label.ID.Send(objc.RegisterName("setDrawsBackground:"), false)
	contentView.AddSubview(unsafe.Pointer(label.ID))

	// Create and configure counter label
	counterLabel = appkit.TextFieldFrom(unsafe.Pointer(
		objc.ID(textFieldClass).Send(objc.RegisterName("alloc")).Send(
			objc.RegisterName("initWithFrame:"),
			NSRect{Origin: NSPoint{50, 80}, Size: NSSize{300, 30}},
		)))
	counterLabel.ID.Send(objc.RegisterName("setStringValue:"), nsString("Clicks: 0"))
	counterLabel.ID.Send(objc.RegisterName("setEditable:"), false)
	counterLabel.ID.Send(objc.RegisterName("setBordered:"), false)
	counterLabel.ID.Send(objc.RegisterName("setDrawsBackground:"), false)
	counterLabel.ID.Send(objc.RegisterName("setAlignment:"), 2) // Center
	contentView.AddSubview(unsafe.Pointer(counterLabel.ID))

	// Create button with title, target, and action using generated constructor
	buttonHandler := createButtonHandler()
	button := appkit.NewButtonWithTitleTargetAction(
		"Click Me!",
		buttonHandler,
		objc.RegisterName("buttonClicked:"),
	)

	// Set button frame and style
	buttonFrame := NSRect{
		Origin: NSPoint{X: 150, Y: 130},
		Size:   NSSize{Width: 100, Height: 40},
	}
	button.ID.Send(objc.RegisterName("setFrame:"), buttonFrame)
	button.ID.Send(objc.RegisterName("setButtonType:"), 0) // NSMomentaryLight = 0
	button.ID.Send(objc.RegisterName("setBezelStyle:"), 1) // NSRoundedBezelStyle = 1

	// Add button to window
	contentView.AddSubview(unsafe.Pointer(button.ID))

	// Show window - using generated method!
	window.MakeKeyAndOrderFront(0)

	fmt.Println("✅ Window created using generated bindings")
	fmt.Println("   No darwinkit dependency!")
	fmt.Println("   Using:")
	fmt.Println("   - Generated types (Button, Window, TextField, Application, View)")
	fmt.Println("   - Generated constructors (NewWindowWithContentRectStyleMaskBackingDefer)")
	fmt.Println("   - Generated methods (AddSubview, MakeKeyAndOrderFront)")
	fmt.Println("   - *From constructors (ButtonFrom, TextFieldFrom, ViewFrom, ApplicationFrom)")
	fmt.Println("   - purego/objc for remaining calls")
	fmt.Println("   Click the button to see the counter increment!")
	fmt.Println("   Close window or press Cmd+Q to quit\n")

	// Finish launching and activate app
	app.ID.Send(objc.RegisterName("finishLaunching"))
	app.ID.Send(objc.RegisterName("activateIgnoringOtherApps:"), true)

	// Run event loop
	app.ID.Send(objc.RegisterName("run"))
}
