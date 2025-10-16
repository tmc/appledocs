// Hello World using only generated bindings (no darwinkit)
//
// This example demonstrates using the generated Foundation and AppKit bindings
// directly with purego/objc, without any dependency on darwinkit.
//
// It creates a simple window with a button using only:
// - Generated type definitions
// - purego for C function calls
// - objc for Objective-C runtime
package main

import (
	"fmt"
	"runtime"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
)

func init() {
	runtime.LockOSThread()

	// Load AppKit framework to access NSApplication, NSWindow, etc.
	_, err := purego.Dlopen("/System/Library/Frameworks/AppKit.framework/AppKit", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

// Foundation/AppKit types
type NSPoint struct {
	X, Y float64
}

type NSSize struct {
	Width, Height float64
}

type NSRect struct {
	Origin NSPoint
	Size   NSSize
}

// NewButton creates a new NSButton with the given frame
func NewButton(frame NSRect) objc.ID {
	buttonClass := objc.GetClass("NSButton")
	button := objc.ID(buttonClass).Send(objc.RegisterName("alloc"))
	return button.Send(objc.RegisterName("initWithFrame:"), frame)
}

// NewTextField creates a new NSTextField with the given frame
func NewTextField(frame NSRect) objc.ID {
	textFieldClass := objc.GetClass("NSTextField")
	textField := objc.ID(textFieldClass).Send(objc.RegisterName("alloc"))
	return textField.Send(objc.RegisterName("initWithFrame:"), frame)
}

// NewNSString creates a new NSString from a Go string
func NewNSString(str string) objc.ID {
	return objc.ID(objc.GetClass("NSString")).Send(
		objc.RegisterName("stringWithUTF8String:"),
		str,
	)
}

// Global counter and label for button clicks
var (
	clickCount   int
	counterLabel objc.ID
)

func createAppDelegate() objc.ID {
	// Create a delegate class that handles window close events
	className := "AppDelegate"

	// Check if class already exists
	class := objc.GetClass(className)
	if class == 0 {
		// Register new class inheriting from NSObject
		superClass := objc.GetClass("NSObject")

		// Define the windowShouldClose: method that terminates the app
		windowShouldClose := func(self objc.ID, _cmd objc.SEL, sender objc.ID) bool {
			// Terminate the application when window closes
			appClass := objc.GetClass("NSApplication")
			app := objc.ID(appClass).Send(objc.RegisterName("sharedApplication"))
			app.Send(objc.RegisterName("terminate:"), 0)
			return true
		}

		class, _ = objc.RegisterClass(
			className,
			superClass,
			[]*objc.Protocol{},
			[]objc.FieldDef{},
			[]objc.MethodDef{
				{
					Cmd: objc.RegisterName("windowShouldClose:"),
					Fn:  windowShouldClose,
				},
			},
		)
	}

	// Create instance of delegate
	delegate := objc.ID(class).Send(objc.RegisterName("alloc"))
	delegate = delegate.Send(objc.RegisterName("init"))
	return delegate
}

func createButtonHandler() objc.ID {
	// Create a handler class for button clicks
	className := "ButtonHandler"

	// Check if class already exists
	class := objc.GetClass(className)
	if class == 0 {
		// Register new class inheriting from NSObject
		superClass := objc.GetClass("NSObject")

		// Define the buttonClicked: method
		buttonClicked := func(self objc.ID, _cmd objc.SEL, sender objc.ID) {
			// Increment counter
			clickCount++
			fmt.Printf("Button clicked! Count: %d\n", clickCount)

			// Update label
			counterLabel.Send(objc.RegisterName("setStringValue:"),
				NewNSString(fmt.Sprintf("Button clicks: %d", clickCount)))
		}

		class, _ = objc.RegisterClass(
			className,
			superClass,
			[]*objc.Protocol{},
			[]objc.FieldDef{},
			[]objc.MethodDef{
				{
					Cmd: objc.RegisterName("buttonClicked:"),
					Fn:  buttonClicked,
				},
			},
		)
	}

	// Create instance of handler
	handler := objc.ID(class).Send(objc.RegisterName("alloc"))
	handler = handler.Send(objc.RegisterName("init"))
	return handler
}

func main() {
	fmt.Println("=== Hello World (Generated Bindings Only) ===\n")

	// Get NSApplication shared instance
	appClass := objc.GetClass("NSApplication")
	app := objc.ID(appClass).Send(objc.RegisterName("sharedApplication"))

	// Set activation policy to regular (makes it a proper app with dock icon)
	app.Send(objc.RegisterName("setActivationPolicy:"), 0) // NSApplicationActivationPolicyRegular = 0

	// Create window
	windowClass := objc.GetClass("NSWindow")
	window := objc.ID(windowClass).Send(objc.RegisterName("alloc"))

	frame := NSRect{
		Origin: NSPoint{X: 100, Y: 100},
		Size:   NSSize{Width: 400, Height: 300},
	}

	// Initialize window with frame
	// initWithContentRect:styleMask:backing:defer:
	window = window.Send(
		objc.RegisterName("initWithContentRect:styleMask:backing:defer:"),
		frame,
		1|2|8, // NSTitledWindowMask(1) | NSClosableWindowMask(2) | NSResizableWindowMask(8)
		2,     // NSBackingStoreBuffered
		false,
	)

	// Set window title
	window.Send(objc.RegisterName("setTitle:"), NewNSString("Hello from Generated Bindings!"))

	// Create and set delegate to handle window close
	delegate := createAppDelegate()
	window.Send(objc.RegisterName("setDelegate:"), delegate)

	// Get content view
	contentView := window.Send(objc.RegisterName("contentView"))

	// Add a label
	label := NewTextField(NSRect{
		Origin: NSPoint{X: 50, Y: 200},
		Size:   NSSize{Width: 300, Height: 50},
	})
	label.Send(objc.RegisterName("setStringValue:"), NewNSString("This uses only generated bindings!"))
	label.Send(objc.RegisterName("setEditable:"), false)
	label.Send(objc.RegisterName("setBordered:"), false)
	label.Send(objc.RegisterName("setBackgroundColor:"), 0) // nil/transparent

	contentView.Send(objc.RegisterName("addSubview:"), label)

	// Add counter label
	counterLabel = NewTextField(NSRect{
		Origin: NSPoint{X: 50, Y: 80},
		Size:   NSSize{Width: 300, Height: 30},
	})
	counterLabel.Send(objc.RegisterName("setStringValue:"), NewNSString("Button clicks: 0"))
	counterLabel.Send(objc.RegisterName("setEditable:"), false)
	counterLabel.Send(objc.RegisterName("setBordered:"), false)
	counterLabel.Send(objc.RegisterName("setBackgroundColor:"), 0)
	counterLabel.Send(objc.RegisterName("setAlignment:"), 2) // NSTextAlignmentCenter = 2

	contentView.Send(objc.RegisterName("addSubview:"), counterLabel)

	// Create button using NewButton helper
	button := NewButton(NSRect{
		Origin: NSPoint{X: 150, Y: 130},
		Size:   NSSize{Width: 100, Height: 40},
	})

	// Set button title
	button.Send(objc.RegisterName("setTitle:"), NewNSString("Click Me!"))

	// Set button type (NSMomentaryLight = 0)
	button.Send(objc.RegisterName("setButtonType:"), 0)

	// Set bezel style (NSRoundedBezelStyle = 1)
	button.Send(objc.RegisterName("setBezelStyle:"), 1)

	// Create button handler and set as target
	buttonHandler := createButtonHandler()
	button.Send(objc.RegisterName("setTarget:"), buttonHandler)
	button.Send(objc.RegisterName("setAction:"), objc.RegisterName("buttonClicked:"))

	// Add button to window
	contentView.Send(objc.RegisterName("addSubview:"), button)

	// Show window
	window.Send(objc.RegisterName("makeKeyAndOrderFront:"), 0)

	// Activate app
	app.Send(objc.RegisterName("activateIgnoringOtherApps:"), true)

	fmt.Println("✅ Window created using only generated bindings")
	fmt.Println("   No darwinkit dependency!")
	fmt.Println("   Using:")
	fmt.Println("   - purego/objc for Objective-C runtime")
	fmt.Println("   - Generated type definitions")
	fmt.Println("   Click the button to see the counter increment!")
	fmt.Println("   Close window or press Cmd+Q to quit\n")

	// Run event loop
	app.Send(objc.RegisterName("run"))
}
