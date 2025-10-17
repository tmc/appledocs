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
			str := objc.ID(objc.GetClass("NSString")).Send(
				objc.RegisterName("stringWithUTF8String:"),
				fmt.Sprintf("Button clicks: %d", clickCount),
			)
			counterLabel.ID.Send(objc.RegisterName("setStringValue:"), str)
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
	app := appkit.ApplicationFrom(unsafe.Pointer(
		objc.ID(appClass).Send(objc.RegisterName("sharedApplication")),
	))

	// Set activation policy to regular (makes it a proper app with dock icon)
	app.ID.Send(objc.RegisterName("setActivationPolicy:"), 0) // NSApplicationActivationPolicyRegular = 0

	// Create window using generated constructor
	frame := NSRect{
		Origin: NSPoint{X: 100, Y: 100},
		Size:   NSSize{Width: 400, Height: 300},
	}
	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		unsafe.Pointer(&frame),
		1|2|8, // NSTitledWindowMask(1) | NSClosableWindowMask(2) | NSResizableWindowMask(8)
		2,     // NSBackingStoreBuffered
		false,
	)

	// Set window title
	title := objc.ID(objc.GetClass("NSString")).Send(
		objc.RegisterName("stringWithUTF8String:"),
		"Hello from Generated Bindings!",
	)
	window.ID.Send(objc.RegisterName("setTitle:"), title)

	// Create and set delegate to handle window close
	delegate := createAppDelegate()
	window.ID.Send(objc.RegisterName("setDelegate:"), delegate)

	// Get content view
	contentViewPtr := window.ID.Send(objc.RegisterName("contentView"))
	contentView := appkit.ViewFrom(unsafe.Pointer(contentViewPtr))

	// Add a label
	labelClass := objc.GetClass("NSTextField")
	label := appkit.TextFieldFrom(unsafe.Pointer(
		objc.ID(labelClass).Send(objc.RegisterName("alloc")).Send(
			objc.RegisterName("initWithFrame:"),
			NSRect{
				Origin: NSPoint{X: 50, Y: 200},
				Size:   NSSize{Width: 300, Height: 50},
			},
		),
	))
	labelStr := objc.ID(objc.GetClass("NSString")).Send(
		objc.RegisterName("stringWithUTF8String:"),
		"This uses only generated bindings!",
	)
	label.ID.Send(objc.RegisterName("setStringValue:"), labelStr)
	label.ID.Send(objc.RegisterName("setEditable:"), false)
	label.ID.Send(objc.RegisterName("setBordered:"), false)
	label.ID.Send(objc.RegisterName("setDrawsBackground:"), false) // transparent

	contentView.AddSubview(unsafe.Pointer(label.ID))

	// Add counter label
	counterLabel = appkit.TextFieldFrom(unsafe.Pointer(
		objc.ID(labelClass).Send(objc.RegisterName("alloc")).Send(
			objc.RegisterName("initWithFrame:"),
			NSRect{
				Origin: NSPoint{X: 50, Y: 80},
				Size:   NSSize{Width: 300, Height: 30},
			},
		),
	))
	counterStr := objc.ID(objc.GetClass("NSString")).Send(
		objc.RegisterName("stringWithUTF8String:"),
		"Button clicks: 0",
	)
	counterLabel.ID.Send(objc.RegisterName("setStringValue:"), counterStr)
	counterLabel.ID.Send(objc.RegisterName("setEditable:"), false)
	counterLabel.ID.Send(objc.RegisterName("setBordered:"), false)
	counterLabel.ID.Send(objc.RegisterName("setDrawsBackground:"), false) // transparent
	counterLabel.ID.Send(objc.RegisterName("setAlignment:"), 2)            // NSTextAlignmentCenter = 2

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
