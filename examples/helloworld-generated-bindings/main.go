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
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/appkit"
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
func NewButton(frame NSRect) appkit.Button {
	button := appkit.ButtonClass.Alloc()
	button.ID = button.ID.Send(objc.RegisterName("initWithFrame:"), frame)
	return button
}

// NewTextField creates a new NSTextField with the given frame
func NewTextField(frame NSRect) appkit.TextField {
	textField := appkit.TextFieldClass.Alloc()
	textField.ID = textField.ID.Send(objc.RegisterName("initWithFrame:"), frame)
	return textField
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
	counterLabel appkit.TextField
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
			app := appkit.SharedApplication()
			app.Terminate(0)
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

			// Update label using typed method
			counterLabel.SetStringValue(fmt.Sprintf("Button clicks: %d", clickCount))
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

	// Get NSApplication shared instance using typed method
	app := appkit.SharedApplication()

	// Set activation policy to regular (makes it a proper app with dock icon)
	app.SetActivationPolicy(0) // NSApplicationActivationPolicyRegular = 0

	// Create window
	window := appkit.WindowClass.Alloc()

	frame := NSRect{
		Origin: NSPoint{X: 100, Y: 100},
		Size:   NSSize{Width: 400, Height: 300},
	}

	// Initialize window with frame
	// initWithContentRect:styleMask:backing:defer:
	window.ID = window.ID.Send(
		objc.RegisterName("initWithContentRect:styleMask:backing:defer:"),
		frame,
		1|2|8, // NSTitledWindowMask(1) | NSClosableWindowMask(2) | NSResizableWindowMask(8)
		2,     // NSBackingStoreBuffered
		false,
	)

	// Set window title using typed method
	window.SetTitle("Hello from Generated Bindings!")

	// Create and set delegate to handle window close
	delegate := createAppDelegate()
	window.SetDelegate(delegate)

	// Get content view using typed method
	contentView := appkit.ViewFrom(window.ContentView())

	// Add a label using typed methods
	label := NewTextField(NSRect{
		Origin: NSPoint{X: 50, Y: 200},
		Size:   NSSize{Width: 300, Height: 50},
	})
	label.SetStringValue("This uses only generated bindings!")
	label.SetEditable(false)
	label.SetBordered(false)
	label.SetBackgroundColor(nil) // transparent

	contentView.AddSubview(unsafe.Pointer(label.ID))

	// Add counter label using typed methods
	counterLabel = NewTextField(NSRect{
		Origin: NSPoint{X: 50, Y: 80},
		Size:   NSSize{Width: 300, Height: 30},
	})
	counterLabel.SetStringValue("Button clicks: 0")
	counterLabel.SetEditable(false)
	counterLabel.SetBordered(false)
	counterLabel.SetBackgroundColor(nil) // transparent
	counterLabel.SetAlignment(2)         // NSTextAlignmentCenter = 2

	contentView.AddSubview(unsafe.Pointer(counterLabel.ID))

	// Create button using NewButton helper
	button := NewButton(NSRect{
		Origin: NSPoint{X: 150, Y: 130},
		Size:   NSSize{Width: 100, Height: 40},
	})

	// Configure button using typed methods
	button.SetTitle("Click Me!")
	button.SetButtonType(0) // NSMomentaryLight = 0
	button.SetBezelStyle(1) // NSRoundedBezelStyle = 1

	// Create button handler and set as target
	buttonHandler := createButtonHandler()
	button.SetTarget(buttonHandler)
	button.SetAction(objc.RegisterName("buttonClicked:"))

	// Add button to window
	contentView.AddSubview(unsafe.Pointer(button.ID))

	// Show window
	window.MakeKeyAndOrderFront(0)

	// Activate app using typed method
	app.ActivateIgnoringOtherApps(true)

	fmt.Println("✅ Window created using typed generated bindings")
	fmt.Println("   No darwinkit dependency!")
	fmt.Println("   Using:")
	fmt.Println("   - Typed methods from generated bindings")
	fmt.Println("   - purego/objc for Objective-C runtime")
	fmt.Println("   Click the button to see the counter increment!")
	fmt.Println("   Close window or press Cmd+Q to quit\n")

	// Run event loop using typed method
	app.Run()
}
