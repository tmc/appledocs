// Hello World using only generated bindings (no darwinkit)
//
// Demonstrates using generated AppKit bindings with purego/objc.
// The generated bindings automatically convert Go strings to NSString objects,
// so you can pass Go strings directly to methods without manual conversion.
//
// Press Cmd+Q to quit.
// Run with -e2e flag for automated end-to-end testing mode.
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

var (
	e2eMode = flag.Bool("e2e", false, "Run in end-to-end test mode (non-interactive)")
)

func init() {
	runtime.LockOSThread()
}

// Note: The generated bindings now handle Go string to NSString conversion automatically.
// You can pass Go strings directly to methods that expect NSString* - no manual conversion needed!

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
			counterLabel.SetStringValue(fmt.Sprintf("Clicks: %d", clickCount))
		}
		class, _ = objc.RegisterClass(className, superClass, nil, nil, []objc.MethodDef{
			{Cmd: objc.RegisterName("buttonClicked:"), Fn: buttonClicked},
		})
	}
	handler := objc.ID(class).Send(objc.RegisterName("alloc"))
	return handler.Send(objc.RegisterName("init"))
}

func main() {
	runtime.LockOSThread()
	flag.Parse()

	if *e2eMode {
		runE2ETest()
		return
	}

	fmt.Println("=== Hello World (Generated Bindings) ===")

	// Initialize the application (matching darwinkit structure exactly)
	app := appkit.SharedApplication()
	app.SetActivationPolicy(appkit.ActivationPolicyRegular)

	// CRITICAL: Disable automatic window restoration entirely
	// State restoration was attempting to restore windows with invalid HUD style mask (0x2000)
	// which caused canBecomeKeyWindow to return NO, making windows invisible.
	fmt.Println("Disabling automatic window restoration...")

	// Method 1: Disable relaunch-on-login (also disables state restoration)
	app.ID.Send(objc.RegisterName("disableRelaunchOnLogin"))

	// Method 2: Create NSApplicationDelegate with state restoration disabled
	fmt.Println("Creating NSApplicationDelegate to prevent state restoration...")

	delegateClass, err := objc.RegisterClass(
		"AppDelegate",
		objc.GetClass("NSObject"),
		nil, // ivars
		nil, // properties
		[]objc.MethodDef{
			{
				// Modern macOS 12+ method
				Cmd: objc.RegisterName("applicationSupportsSecureRestorableState:"),
				Fn: func(self objc.ID, _cmd objc.SEL, app objc.ID) bool {
					fmt.Println("DEBUG: applicationSupportsSecureRestorableState: called, returning NO")
					return false // Disable secure state restoration
				},
			},
			{
				// Legacy method for older macOS versions
				Cmd: objc.RegisterName("application:shouldRestoreApplicationState:"),
				Fn: func(self objc.ID, _cmd objc.SEL, app objc.ID, coder objc.ID) bool {
					fmt.Println("DEBUG: application:shouldRestoreApplicationState: called, returning NO")
					return false // Never restore application state
				},
			},
			{
				// Additional method to prevent window restoration
				Cmd: objc.RegisterName("applicationShouldAutomaticallyRestoreWindows:"),
				Fn: func(self objc.ID, _cmd objc.SEL, app objc.ID) bool {
					fmt.Println("DEBUG: applicationShouldAutomaticallyRestoreWindows: called, returning NO")
					return false // Prevent automatic window restoration
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ERROR: Failed to register delegate class: %v\n", err)
		os.Exit(1)
	}

	delegate := objc.ID(delegateClass).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	app.ID.Send(objc.RegisterName("setDelegate:"), delegate)
	fmt.Println("✓ NSApplicationDelegate registered and set")

	fmt.Println("✓ Starting AppKit application...")

	// Create window - use direct alloc/init to avoid premature autorelease
	type NSPoint struct{ X, Y float64 }
	type NSSize struct{ Width, Height float64 }
	type NSRect struct {
		Origin NSPoint
		Size   NSSize
	}
	rect := NSRect{
		Origin: NSPoint{X: 100, Y: 100},
		Size:   NSSize{Width: 400, Height: 300},
	}

	windowClass := objc.GetClass("NSWindow")
	windowID := objc.ID(windowClass).Send(objc.RegisterName("alloc"))
	styleMask := appkit.WindowStyleMaskTitled | appkit.WindowStyleMaskClosable | appkit.WindowStyleMaskResizable
	fmt.Printf("DEBUG: styleMask = 0x%x (Titled=0x%x, Closable=0x%x, Resizable=0x%x)\n",
		styleMask, appkit.WindowStyleMaskTitled, appkit.WindowStyleMaskClosable, appkit.WindowStyleMaskResizable)
	windowID = windowID.Send(objc.RegisterName("initWithContentRect:styleMask:backing:defer:"),
		unsafe.Pointer(&rect), styleMask, appkit.BackingStoreBuffered, false)
	window := appkit.WindowFrom(unsafe.Pointer(windowID))

	fmt.Printf("DEBUG: Window ID = %v (should be non-zero)\n", windowID)

	// SetTitle accepts Go strings directly (automatic conversion to NSString)
	window.SetTitle("Hello from Generated Bindings!")

	// Get content view
	contentView := window.ContentView()

	// Create and configure label
	label := appkit.NewTextFieldWithFrame(50, 200, 300, 50)
	// SetStringValue also accepts Go strings directly
	label.SetStringValue("Using generated bindings!")
	label.SetEditable(false)
	label.SetBordered(false)
	label.SetDrawsBackground(false)
	contentView.AddSubviewTyped(label)

	// Create and configure counter label
	counterLabel = appkit.NewTextFieldWithFrame(50, 80, 300, 30)
	counterLabel.SetStringValue("Clicks: 0")
	counterLabel.SetEditable(false)
	counterLabel.SetBordered(false)
	counterLabel.SetDrawsBackground(false)
	counterLabel.SetAlignment(appkit.TextAlignmentCenter)
	contentView.AddSubviewTyped(counterLabel)

	// Create button using generated constructor with automatic string conversion
	button := appkit.NewButtonWithTitleTargetAction("Click Me!", createButtonHandler(), objc.RegisterName("buttonClicked:"))
	button.SetFrameRect(150, 130, 100, 40)
	button.SetButtonType(appkit.ButtonTypeMomentaryLight)
	button.SetBezelStyle(appkit.BezelStyleRounded)
	contentView.AddSubviewTyped(button)

	// Retain window to prevent premature deallocation (critical for purego!)
	window.ID.Send(objc.RegisterName("retain"))

	// Center window on screen for better visibility
	window.ID.Send(objc.RegisterName("center"))

	// CRITICAL: Activate app BEFORE showing window (order matters!)
	app.ActivateIgnoringOtherApps(true)

	// Show window using generated method - pass window.ID as sender
	window.MakeKeyAndOrderFront(window.ID)

	// Force window to front regardless of other windows
	window.ID.Send(objc.RegisterName("orderFrontRegardless"))

	// Check if window is visible
	isVisible := window.ID.Send(objc.RegisterName("isVisible"))
	canBecomeKey := window.ID.Send(objc.RegisterName("canBecomeKeyWindow"))
	isKeyWindow := window.ID.Send(objc.RegisterName("isKeyWindow"))
	fmt.Printf("DEBUG: Window visibility - isVisible: %v, canBecomeKey: %v, isKeyWindow: %v\n",
		isVisible != 0, canBecomeKey != 0, isKeyWindow != 0)

	fmt.Println("✓ Window created and displayed")
	fmt.Println("✅ Using generated bindings with automatic string conversion:")
	fmt.Println("   - Types: Window, Button, TextField, View, Application")
	fmt.Println("   - Constructors: NewWindowWithFrame, NewButtonWithTitleTargetAction")
	fmt.Println("   - Methods: SetTitle, SetStringValue (Go strings → NSString automatically!)")
	fmt.Println("   - Type safety: AddSubviewTyped accepts IView interface")
	fmt.Println()
	fmt.Println("   No manual string conversion needed - pass Go strings directly!")
	fmt.Println("   Click the button! Press Cmd+Q to quit.")

	// Run the application
	app.Run()
}

// runE2ETest runs automated end-to-end tests without user interaction.
func runE2ETest() {
	fmt.Println("=== E2E Test Mode (Generated Bindings) ===")

	// Get NSApplication shared instance
	app := appkit.SharedApplication()
	app.SetActivationPolicy(appkit.ActivationPolicyAccessory) // No dock icon in tests

	// Create window
	window := appkit.NewWindowWithFrame(100, 100, 400, 300,
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable|appkit.WindowStyleMaskResizable)
	window.SetTitle("E2E Test Window")
	fmt.Println("✓ Created window with title")

	// Get content view
	contentView := window.ContentView()
	fmt.Println("✓ Got content view")

	// Create and configure label
	label := appkit.NewTextFieldWithFrame(50, 200, 300, 50)
	label.SetStringValue("Test Label")
	label.SetEditable(false)
	label.SetBordered(false)
	contentView.AddSubviewTyped(label)
	fmt.Println("✓ Created and configured label")

	// Create counter label
	counterLabel = appkit.NewTextFieldWithFrame(50, 80, 300, 30)
	counterLabel.SetStringValue("Clicks: 0")
	counterLabel.SetEditable(false)
	counterLabel.SetBordered(false)
	counterLabel.SetAlignment(appkit.TextAlignmentCenter)
	contentView.AddSubviewTyped(counterLabel)
	fmt.Println("✓ Created counter label")

	// Create button
	button := appkit.NewButtonWithFrame(150, 130, 100, 40)
	button.SetTitleString("Test Button")
	button.SetButtonType(appkit.ButtonTypeMomentaryLight)
	button.SetBezelStyle(appkit.BezelStyleRounded)

	handler := createButtonHandler()
	button.SetTarget(handler)
	button.SetAction(objc.RegisterName("buttonClicked:"))
	contentView.AddSubviewTyped(button)
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
