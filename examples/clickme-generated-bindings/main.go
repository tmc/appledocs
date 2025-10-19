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
	e2e          = flag.Bool("e2e", false, "run end-to-end tests")
	simple       = flag.Bool("simple", false, "run simple version without RunApp/delegate")
	rawTest      = flag.Bool("raw", false, "test raw window creation with purego/objc")
	clickCount   int
	counterLabel appkit.TextField
)

func init() {
	runtime.LockOSThread()
}

// RunApp runs the AppKit event loop with proper initialization order.
// It mimics DarwinKit's RunApp pattern: creates a delegate, sets it on the app,
// and calls the didLaunch callback AFTER the app finishes launching.
// This ensures macOS has fully initialized before any windows are created.
func RunApp(didLaunch func(app appkit.Application)) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	fmt.Println("=== Click Me (Generated Bindings) ===")
	fmt.Println("DEBUG: Starting RunApp...")

	// Get shared application instance
	app := appkit.SharedApplication()
	fmt.Println("DEBUG: Got shared application")

	// Create application delegate with ApplicationDidFinishLaunching callback
	delegateClass, err := objc.RegisterClass(
		"AppDelegate",
		objc.GetClass("NSObject"),
		nil, // ivars
		nil, // properties
		[]objc.MethodDef{
			{
				// ApplicationDidFinishLaunching - called when app is ready
				Cmd: objc.RegisterName("applicationDidFinishLaunching:"),
				Fn: func(self objc.ID, _cmd objc.SEL, notification objc.ID) {
					fmt.Println("DEBUG: ✓ applicationDidFinishLaunching called")
					// Call user's setup code AFTER app is fully initialized
					didLaunch(app)
				},
			},
			{
				// Disable state restoration (prevents HUD mask issues)
				Cmd: objc.RegisterName("applicationSupportsSecureRestorableState:"),
				Fn: func(self objc.ID, _cmd objc.SEL, app objc.ID) bool {
					fmt.Println("DEBUG: applicationSupportsSecureRestorableState called, returning false")
					return false
				},
			},
			{
				// Quit when last window closes (convenience)
				Cmd: objc.RegisterName("applicationShouldTerminateAfterLastWindowClosed:"),
				Fn: func(self objc.ID, _cmd objc.SEL, sender objc.ID) bool {
					fmt.Println("DEBUG: applicationShouldTerminateAfterLastWindowClosed called, returning true")
					return true
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ERROR: Failed to register delegate class: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("DEBUG: Registered delegate class")

	delegate := objc.ID(delegateClass).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	app.ID.Send(objc.RegisterName("setDelegate:"), delegate)
	fmt.Println("DEBUG: Set delegate on application")

	fmt.Println("DEBUG: Starting AppKit event loop...")
	// Run the event loop - this will call applicationDidFinishLaunching
	app.Run()
}

func main() {
	runtime.LockOSThread()
	flag.Parse()

	// Raw test mode - test pure purego/objc window creation
	if *rawTest {
		runRawTest()
		return
	}

	// Simple mode - no RunApp, no delegate (matching darwinkit clickme)
	if *simple {
		runSimpleMode()
		return
	}

	// In E2E mode, run synchronously without RunApp pattern
	if *e2e {
		fmt.Println("=== E2E Mode ===")
		app := appkit.SharedApplication()
		app.SetActivationPolicy(appkit.ActivationPolicyAccessory)

		styleMask := appkit.WindowStyleMaskTitled | appkit.WindowStyleMaskClosable
		window := appkit.NewWindowWithFrame(100, 100, 400, 300, styleMask)
		window.SetTitle("E2E Test")

		counterLabel = appkit.NewTextFieldWithFrame(100, 200, 200, 40)
		counterLabel.SetStringValue("Clicks: 0")
		counterLabel.SetEditable(false)
		counterLabel.SetBordered(false)
		counterLabel.SetDrawsBackground(false)
		window.ContentView().AddSubviewTyped(counterLabel)

		button := appkit.NewButtonWithFrame(150, 120, 100, 40)
		button.SetTitleString("Click Me!")
		button.SetTarget(createButtonHandler())
		button.SetAction(objc.RegisterName("buttonClicked:"))
		window.ContentView().AddSubviewTyped(button)

		fmt.Println("✓ E2E test passed")
		return
	}

	// Use RunApp pattern - window creation happens in the callback
	RunApp(func(app appkit.Application) {
		fmt.Println("DEBUG: Inside didLaunch callback")

		// Set activation policy INSIDE the callback (after app is initialized)
		app.SetActivationPolicy(appkit.ActivationPolicyRegular)
		fmt.Println("DEBUG: Set activation policy to Regular")

		app.ActivateIgnoringOtherApps(true)
		fmt.Println("DEBUG: Activated app ignoring other apps")

		// Create window - app is fully initialized now
		fmt.Println("DEBUG: Creating window...")
		styleMask := appkit.WindowStyleMaskTitled | appkit.WindowStyleMaskClosable | appkit.WindowStyleMaskResizable
		window := appkit.NewWindowWithFrame(100, 100, 400, 300, styleMask)
		fmt.Printf("DEBUG: Created window, ID: %v\n", window.ID)

		window.SetTitle("Click Me!")
		fmt.Println("DEBUG: Set window title")

		// Create counter label
		fmt.Println("DEBUG: Creating counter label...")
		counterLabel = appkit.NewTextFieldWithFrame(100, 200, 200, 40)
		counterLabel.SetStringValue("Clicks: 0")
		counterLabel.SetEditable(false)
		counterLabel.SetBordered(false)
		counterLabel.SetDrawsBackground(false)
		counterLabel.SetAlignment(appkit.TextAlignmentCenter)
		fmt.Printf("DEBUG: Created counter label, ID: %v\n", counterLabel.ID)

		// Add to content view
		contentView := window.ContentView()
		fmt.Printf("DEBUG: Got content view, ID: %v\n", contentView.ID)
		contentView.AddSubviewTyped(counterLabel)
		fmt.Println("DEBUG: Added counter label to content view")

		// Create button handler
		fmt.Println("DEBUG: Creating button handler...")
		buttonHandler := createButtonHandler()
		fmt.Printf("DEBUG: Created button handler, ID: %v\n", buttonHandler)

		// Create button using generated constructor
		fmt.Println("DEBUG: Creating button...")
		button := appkit.NewButtonWithFrame(150, 120, 100, 40)
		fmt.Printf("DEBUG: Created button, ID: %v\n", button.ID)

		button.SetTitleString("Click Me!")
		fmt.Println("DEBUG: Set button title")

		// Set button target and action
		button.SetTarget(buttonHandler)
		button.SetAction(objc.RegisterName("buttonClicked:"))
		fmt.Println("DEBUG: Set button target and action")

		// Add button to window's content view
		contentView.AddSubviewTyped(button)
		fmt.Println("DEBUG: Added button to content view")

		// Retain window to prevent premature deallocation (critical for purego!)
		window.ID.Send(objc.RegisterName("retain"))
		fmt.Println("DEBUG: Retained window")

		// Center window on screen
		window.ID.Send(objc.RegisterName("center"))
		fmt.Println("DEBUG: Centered window")

		// Show window
		fmt.Println("DEBUG: Making window key and ordering front...")
		window.MakeKeyAndOrderFront(window.ID)

		// Check if window is visible
		isVisible := window.ID.Send(objc.RegisterName("isVisible"))
		canBecomeKey := window.ID.Send(objc.RegisterName("canBecomeKeyWindow"))
		isKeyWindow := window.ID.Send(objc.RegisterName("isKeyWindow"))
		fmt.Printf("DEBUG: Window state - isVisible: %v, canBecomeKey: %v, isKeyWindow: %v\n",
			isVisible != 0, canBecomeKey != 0, isKeyWindow != 0)

		fmt.Println("✓ Window created and displayed")
		fmt.Println("Click the button! Press Cmd+Q to quit.")
	})
}

// createButtonHandler creates an NSObject subclass that handles button clicks
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

// runRawTest tests raw window creation with pure purego/objc (no generated bindings)
func runRawTest() {
	fmt.Println("=== Raw Test Mode (Pure purego/objc) ===")

	// Get NSWindow class directly
	windowClass := objc.GetClass("NSWindow")
	fmt.Printf("DEBUG: NSWindow class: %v\n", windowClass)

	// Create window using raw objc
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

	styleMask := uintptr(0xb) // Titled | Closable | Resizable

	// alloc
	windowID := objc.ID(windowClass).Send(objc.RegisterName("alloc"))
	fmt.Printf("DEBUG: After alloc, windowID: %v\n", windowID)

	// initWithContentRect:styleMask:backing:defer:
	windowID = windowID.Send(objc.RegisterName("initWithContentRect:styleMask:backing:defer:"),
		unsafe.Pointer(&rect), styleMask, uintptr(2), false)
	fmt.Printf("DEBUG: After init, windowID: %v\n", windowID)

	// Check canBecomeKeyWindow immediately after init
	canBecomeKey := windowID.Send(objc.RegisterName("canBecomeKeyWindow"))
	fmt.Printf("DEBUG: After init - canBecomeKey: %v (raw: %d)\n", canBecomeKey != 0, canBecomeKey)

	// Check styleMask
	actualStyleMask := windowID.Send(objc.RegisterName("styleMask"))
	fmt.Printf("DEBUG: Actual styleMask: 0x%x\n", actualStyleMask)

	// Set title using raw NSString creation
	titleSel := objc.RegisterName("setTitle:")
	nsStringClass := objc.GetClass("NSString")
	titleStr := objc.ID(nsStringClass).Send(objc.RegisterName("stringWithUTF8String:"), "Raw Test Window\x00")
	windowID.Send(titleSel, titleStr)
	fmt.Println("DEBUG: Set title")

	// Check canBecomeKey again
	canBecomeKey2 := windowID.Send(objc.RegisterName("canBecomeKeyWindow"))
	fmt.Printf("DEBUG: After setTitle - canBecomeKey: %v (raw: %d)\n", canBecomeKey2 != 0, canBecomeKey2)

	fmt.Println("\n✓ Raw test complete - window created with pure purego/objc")
	fmt.Printf("RESULT: canBecomeKeyWindow = %v\n", canBecomeKey != 0)
}

// runSimpleMode runs without RunApp helper or delegate, matching darwinkit's clickme pattern
func runSimpleMode() {
	fmt.Println("=== Click Me (Simple Mode - No Delegate) ===")
	fmt.Println("DEBUG: Matching darwinkit clickme pattern")

	// Create application
	app := appkit.SharedApplication()
	app.SetActivationPolicy(appkit.ActivationPolicyRegular)
	fmt.Println("DEBUG: Set activation policy")

	// Create window BEFORE app.Run() (darwinkit pattern)
	fmt.Println("DEBUG: Creating window...")
	styleMask := appkit.WindowStyleMaskTitled | appkit.WindowStyleMaskClosable | appkit.WindowStyleMaskResizable
	window := appkit.NewWindowWithFrame(100, 100, 400, 300, styleMask)
	fmt.Printf("DEBUG: Created window, ID: %v\n", window.ID)

	window.SetTitle("Click Me (Simple - No Delegate)")
	fmt.Println("DEBUG: Set window title")

	// Create counter label
	counterLabel = appkit.NewTextFieldWithFrame(100, 200, 200, 40)
	counterLabel.SetStringValue("Clicks: 0")
	counterLabel.SetEditable(false)
	counterLabel.SetBordered(false)
	counterLabel.SetDrawsBackground(false)
	counterLabel.SetAlignment(appkit.TextAlignmentCenter)
	window.ContentView().AddSubviewTyped(counterLabel)
	fmt.Println("DEBUG: Created and added counter label")

	// Create button
	button := appkit.NewButtonWithFrame(150, 120, 100, 40)
	button.SetTitleString("Click Me!")
	button.SetTarget(createButtonHandler())
	button.SetAction(objc.RegisterName("buttonClicked:"))
	window.ContentView().AddSubviewTyped(button)
	fmt.Println("DEBUG: Created and added button")

	// Show window BEFORE app.Run() (darwinkit pattern)
	window.MakeKeyAndOrderFront(window.ID)
	fmt.Println("DEBUG: Made window key and ordered front")

	// Activate app BEFORE app.Run() (darwinkit pattern)
	app.ActivateIgnoringOtherApps(true)
	fmt.Println("DEBUG: Activated app")

	// Check window state
	isVisible := window.ID.Send(objc.RegisterName("isVisible"))
	canBecomeKey := window.ID.Send(objc.RegisterName("canBecomeKeyWindow"))
	isKeyWindow := window.ID.Send(objc.RegisterName("isKeyWindow"))
	fmt.Printf("DEBUG: Window state - isVisible: %v, canBecomeKey: %v, isKeyWindow: %v\n",
		isVisible != 0, canBecomeKey != 0, isKeyWindow != 0)

	fmt.Println("✓ Window created and displayed")
	fmt.Println("✓ NO RunApp helper, NO delegate (darwinkit pattern)")
	fmt.Println("Click the button! Press Cmd+Q to quit.")

	// Run application (window already created and visible)
	app.Run()
}
