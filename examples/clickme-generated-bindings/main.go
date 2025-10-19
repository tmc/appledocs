package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

var (
	e2e          = flag.Bool("e2e", false, "run end-to-end tests")
	clickCount   int
	counterLabel appkit.TextField
)

func init() {
	runtime.LockOSThread()
}

func main() {
	flag.Parse()

	// Create application
	app := appkit.SharedApplication()
	app.SetActivationPolicy(appkit.ActivationPolicyRegular)

	// Create window using generated constructor
	window := appkit.NewWindowWithFrame(100, 100, 400, 300,
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable)

	// SetTitle accepts Go strings directly (automatic conversion to NSString)
	window.SetTitle("Hello from Generated Bindings!")

	// Create counter label
	counterLabel = appkit.NewTextFieldWithFrame(100, 200, 200, 40)
	counterLabel.SetStringValue("Clicks: 0")
	counterLabel.SetEditable(false)
	counterLabel.SetBordered(false)
	counterLabel.SetDrawsBackground(false)
	counterLabel.SetAlignment(appkit.TextAlignmentCenter)
	window.ContentView().AddSubviewTyped(counterLabel)

	// Create button handler
	buttonHandler := createButtonHandler()

	// Create button using generated constructor
	// Note: Generated bindings automatically convert Go strings to NSString
	button := appkit.NewButtonWithFrame(150, 120, 100, 40)
	button.SetTitleString("Click Me!")

	// Set button target and action
	button.SetTarget(buttonHandler)
	button.SetAction(objc.RegisterName("buttonClicked:"))

	// Add button to window's content view
	window.ContentView().AddSubviewTyped(button)

	// Retain window to prevent premature deallocation (critical for purego!)
	window.ID.Send(objc.RegisterName("retain"))

	// Center window on screen for better visibility
	window.ID.Send(objc.RegisterName("center"))

	// CRITICAL: Activate app BEFORE showing window (order matters!)
	app.ActivateIgnoringOtherApps(true)

	// Make window key and bring to front
	window.MakeKeyAndOrderFront(window.ID)

	// Force window to front regardless of other windows
	window.ID.Send(objc.RegisterName("orderFrontRegardless"))

	// In E2E mode, terminate immediately
	if *e2e {
		fmt.Println("E2E mode: terminating immediately")
		app.Terminate(0)
		return
	}

	// Run the application event loop
	app.Run()
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
