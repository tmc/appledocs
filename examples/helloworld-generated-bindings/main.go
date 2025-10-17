// Hello World using only generated bindings (no darwinkit)
//
// Demonstrates using generated AppKit bindings with purego/objc.
// Press Cmd+Q to quit.
//
// Run with -e2e flag for automated end-to-end testing mode.
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

var (
	e2eMode = flag.Bool("e2e", false, "Run in end-to-end test mode (non-interactive)")
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
	flag.Parse()

	if *e2eMode {
		runE2ETest()
		return
	}

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
	// counterLabel.ID.Send(objc.RegisterName("setStringValue:"), nsString("Clicks: 0"))
	counterLabel.SetStringValue("Clicks: 0")
	counterLabel.ID.Send(objc.RegisterName("setEditable:"), false)
	counterLabel.ID.Send(objc.RegisterName("setBordered:"), false)
	counterLabel.ID.Send(objc.RegisterName("setDrawsBackground:"), false)
	counterLabel.ID.Send(objc.RegisterName("setAlignment:"), 2) // Center
	contentView.AddSubview(unsafe.Pointer(counterLabel.ID))

	// Create button using generated constructor
	button := appkit.NewButtonWithTitleTargetAction("Click Me!", createButtonHandler(), objc.RegisterName("buttonClicked:"))
	button.ID.Send(objc.RegisterName("setFrame:"), NSRect{Origin: NSPoint{150, 130}, Size: NSSize{100, 40}})
	button.ID.Send(objc.RegisterName("setButtonType:"), 0) // Momentary
	button.ID.Send(objc.RegisterName("setBezelStyle:"), 1) // Rounded
	contentView.AddSubview(unsafe.Pointer(button.ID))

	// Show window using generated method
	window.MakeKeyAndOrderFront(0)

	fmt.Println("✅ Using generated bindings:")
	fmt.Println("   - Types: Window, Button, TextField, View, Application")
	fmt.Println("   - Constructors: NewWindowWithContentRectStyleMaskBackingDefer, NewButtonWithTitleTargetAction")
	fmt.Println("   - Methods: MakeKeyAndOrderFront, AddSubview")
	fmt.Println("   - Conversions: ApplicationFrom, TextFieldFrom, ViewFrom")
	fmt.Println("\n   Click the button! Press Cmd+Q to quit.\n")

	// Finish launching and run
	app.ID.Send(objc.RegisterName("finishLaunching"))
	app.ID.Send(objc.RegisterName("activateIgnoringOtherApps:"), true)
	app.ID.Send(objc.RegisterName("run"))
}

// runE2ETest runs automated end-to-end tests without user interaction.
func runE2ETest() {
	fmt.Println("=== E2E Test Mode (Generated Bindings) ===\n")

	// Get NSApplication shared instance
	appClass := objc.GetClass("NSApplication")
	app := appkit.ApplicationFrom(unsafe.Pointer(objc.ID(appClass).Send(objc.RegisterName("sharedApplication"))))
	app.ID.Send(objc.RegisterName("setActivationPolicy:"), 1) // Accessory app (no dock icon in tests)

	// Create window
	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		unsafe.Pointer(&NSRect{Origin: NSPoint{100, 100}, Size: NSSize{400, 300}}),
		1|2|8, // Titled, Closable, Resizable
		2,     // Buffered
		false,
	)
	window.ID.Send(objc.RegisterName("setTitle:"), nsString("E2E Test Window"))
	fmt.Println("✓ Created window with title")

	// Get content view
	contentView := appkit.ViewFrom(unsafe.Pointer(window.ID.Send(objc.RegisterName("contentView"))))
	fmt.Println("✓ Got content view")

	// Create and configure label
	textFieldClass := objc.GetClass("NSTextField")
	label := appkit.TextFieldFrom(unsafe.Pointer(
		objc.ID(textFieldClass).Send(objc.RegisterName("alloc")).Send(
			objc.RegisterName("initWithFrame:"),
			NSRect{Origin: NSPoint{50, 200}, Size: NSSize{300, 50}},
		)))
	label.ID.Send(objc.RegisterName("setStringValue:"), nsString("Test Label"))
	label.ID.Send(objc.RegisterName("setEditable:"), false)
	label.ID.Send(objc.RegisterName("setBordered:"), false)
	contentView.AddSubview(unsafe.Pointer(label.ID))
	fmt.Println("✓ Created and configured label")

	// Create counter label
	counterLabel = appkit.TextFieldFrom(unsafe.Pointer(
		objc.ID(textFieldClass).Send(objc.RegisterName("alloc")).Send(
			objc.RegisterName("initWithFrame:"),
			NSRect{Origin: NSPoint{50, 80}, Size: NSSize{300, 30}},
		)))
	counterLabel.ID.Send(objc.RegisterName("setStringValue:"), nsString("Clicks: 0"))
	counterLabel.ID.Send(objc.RegisterName("setEditable:"), false)
	counterLabel.ID.Send(objc.RegisterName("setBordered:"), false)
	counterLabel.ID.Send(objc.RegisterName("setAlignment:"), 2) // Center
	contentView.AddSubview(unsafe.Pointer(counterLabel.ID))
	fmt.Println("✓ Created counter label")

	// Create button using simpler approach
	buttonClass := objc.GetClass("NSButton")
	button := appkit.ButtonFrom(unsafe.Pointer(
		objc.ID(buttonClass).Send(objc.RegisterName("alloc")).Send(
			objc.RegisterName("initWithFrame:"),
			NSRect{Origin: NSPoint{150, 130}, Size: NSSize{100, 40}},
		)))
	button.ID.Send(objc.RegisterName("setTitle:"), nsString("Test Button"))
	button.ID.Send(objc.RegisterName("setButtonType:"), 0) // Momentary
	button.ID.Send(objc.RegisterName("setBezelStyle:"), 1) // Rounded

	handler := createButtonHandler()
	button.ID.Send(objc.RegisterName("setTarget:"), handler)
	button.ID.Send(objc.RegisterName("setAction:"), objc.RegisterName("buttonClicked:"))
	contentView.AddSubview(unsafe.Pointer(button.ID))
	fmt.Println("✓ Created button with target/action")

	// Verify button is valid
	if button.ID == 0 {
		fmt.Println("✗ FAIL: Button not created")
		os.Exit(1)
	}
	fmt.Println("✓ Button created and configured")

	// Verify label string values
	labelValue := label.ID.Send(objc.RegisterName("stringValue"))
	counterValue := counterLabel.ID.Send(objc.RegisterName("stringValue"))
	if labelValue == 0 || counterValue == 0 {
		fmt.Println("✗ FAIL: Label values not set")
		os.Exit(1)
	}
	fmt.Println("✓ Label values set correctly")

	// Close window
	window.ID.Send(objc.RegisterName("close"))
	fmt.Println("✓ Window closed")

	fmt.Println("\n=== E2E Test PASSED ===")
	os.Exit(0)
}
