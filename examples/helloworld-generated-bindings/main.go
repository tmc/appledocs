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
	titleStr := objc.ID(objc.GetClass("NSString")).Send(
		objc.RegisterName("stringWithUTF8String:"),
		"Hello from Generated Bindings!",
	)
	window.Send(objc.RegisterName("setTitle:"), titleStr)

	// Create and set delegate to handle window close
	delegate := createAppDelegate()
	window.Send(objc.RegisterName("setDelegate:"), delegate)

	// Create button
	buttonClass := objc.GetClass("NSButton")
	button := objc.ID(buttonClass).Send(objc.RegisterName("alloc"))

	buttonFrame := NSRect{
		Origin: NSPoint{X: 150, Y: 130},
		Size:   NSSize{Width: 100, Height: 40},
	}

	button = button.Send(objc.RegisterName("initWithFrame:"), buttonFrame)

	// Set button title
	buttonTitle := objc.ID(objc.GetClass("NSString")).Send(
		objc.RegisterName("stringWithUTF8String:"),
		"Click Me!",
	)
	button.Send(objc.RegisterName("setTitle:"), buttonTitle)

	// Set button type (NSMomentaryLight = 0)
	button.Send(objc.RegisterName("setButtonType:"), 0)

	// Set bezel style (NSRoundedBezelStyle = 1)
	button.Send(objc.RegisterName("setBezelStyle:"), 1)

	// Get content view and add button
	contentView := window.Send(objc.RegisterName("contentView"))
	contentView.Send(objc.RegisterName("addSubview:"), button)

	// Add a label
	textFieldClass := objc.GetClass("NSTextField")
	label := objc.ID(textFieldClass).Send(objc.RegisterName("alloc"))

	labelFrame := NSRect{
		Origin: NSPoint{X: 50, Y: 200},
		Size:   NSSize{Width: 300, Height: 50},
	}

	label = label.Send(objc.RegisterName("initWithFrame:"), labelFrame)

	labelText := objc.ID(objc.GetClass("NSString")).Send(
		objc.RegisterName("stringWithUTF8String:"),
		"This uses only generated bindings!",
	)
	label.Send(objc.RegisterName("setStringValue:"), labelText)
	label.Send(objc.RegisterName("setEditable:"), false)
	label.Send(objc.RegisterName("setBordered:"), false)
	label.Send(objc.RegisterName("setBackgroundColor:"), 0) // nil/transparent

	contentView.Send(objc.RegisterName("addSubview:"), label)

	// Show window
	window.Send(objc.RegisterName("makeKeyAndOrderFront:"), 0)

	// Activate app
	app.Send(objc.RegisterName("activateIgnoringOtherApps:"), true)

	fmt.Println("✅ Window created using only generated bindings")
	fmt.Println("   No darwinkit dependency!")
	fmt.Println("   Using:")
	fmt.Println("   - purego/objc for Objective-C runtime")
	fmt.Println("   - Generated type definitions")
	fmt.Println("   Press Cmd+Q to quit\n")

	// Run event loop
	app.Send(objc.RegisterName("run"))
}
