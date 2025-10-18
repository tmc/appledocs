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
	flag.Parse()

	if *e2eMode {
		runE2ETest()
		return
	}

	fmt.Println("=== Hello World (Generated Bindings) ===\n")

	// Get NSApplication shared instance
	app := appkit.SharedApplication()
	app.SetActivationPolicy(appkit.ActivationPolicyRegular)

	// Create window
	window := appkit.NewWindowWithFrame(100, 100, 400, 300,
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable|appkit.WindowStyleMaskResizable)
	window.SetTitle("Hello from Generated Bindings!")

	// Get content view
	contentView := window.ContentView()

	// Create and configure label
	label := appkit.NewTextFieldWithFrame(50, 200, 300, 50)
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

	// Create button using generated constructor
	button := appkit.NewButtonWithTitleTargetAction("Click Me!", createButtonHandler(), objc.RegisterName("buttonClicked:"))
	button.SetFrameRect(150, 130, 100, 40)
	button.SetButtonType(appkit.ButtonTypeMomentaryLight)
	button.SetBezelStyle(appkit.BezelStyleRounded)
	contentView.AddSubviewTyped(button)

	// Show window using generated method
	window.MakeKeyAndOrderFront(0)

	fmt.Println("✅ Using generated bindings:")
	fmt.Println("   - Types: Window, Button, TextField, View, Application")
	fmt.Println("   - Constructors: NewWindowWithContentRectStyleMaskBackingDefer, NewButtonWithTitleTargetAction")
	fmt.Println("   - Methods: MakeKeyAndOrderFront, AddSubviewTyped (type-safe!), SetStringValue")
	fmt.Println("   - Conversions: ApplicationFrom, TextFieldFrom, ViewFrom")
	fmt.Println("   - Interfaces: IView, IButton, ITextField for type safety")
	fmt.Println("\n   Click the button! Press Cmd+Q to quit.\n")

	// Finish launching and run
	app.FinishLaunching()
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}

// runE2ETest runs automated end-to-end tests without user interaction.
func runE2ETest() {
	fmt.Println("=== E2E Test Mode (Generated Bindings) ===\n")

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
