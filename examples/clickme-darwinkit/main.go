package main

import (
	"flag"
	"fmt"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	darwinkitObjc "github.com/progrium/darwinkit/objc"
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
	app := appkit.Application_SharedApplication()
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)

	// Create window
	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		foundation.Rect{
			Origin: foundation.Point{X: 100, Y: 100},
			Size:   foundation.Size{Width: 400, Height: 300},
		},
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable,
		appkit.BackingStoreBuffered,
		false,
	)
	window.SetTitle("Hello from DarwinKit!")

	// Create counter label
	counterLabel = appkit.NewTextFieldWithFrame(foundation.Rect{
		Origin: foundation.Point{X: 100, Y: 200},
		Size:   foundation.Size{Width: 200, Height: 40},
	})
	counterLabel.SetStringValue("Clicks: 0")
	counterLabel.SetEditable(false)
	counterLabel.SetBordered(false)
	counterLabel.SetDrawsBackground(false)
	counterLabel.SetAlignment(appkit.TextAlignmentCenter)
	window.ContentView().AddSubview(counterLabel)

	// Create button handler
	buttonHandler := createButtonHandler()

	// Create button with target and action using constructor
	button := appkit.Button_ButtonWithTitleTargetAction(
		"Click Me!",
		darwinkitObjc.ObjectFrom(unsafe.Pointer(buttonHandler)),
		darwinkitObjc.Sel("buttonClicked:"),
	)

	// Set button frame
	button.SetFrameOrigin(foundation.Point{X: 150, Y: 120})
	button.SetFrameSize(foundation.Size{Width: 100, Height: 40})

	// Add button to window's content view
	window.ContentView().AddSubview(button)

	// Make window key and bring to front
	window.MakeKeyAndOrderFront(nil)

	// Activate ignoring other apps
	app.ActivateIgnoringOtherApps(true)

	// In E2E mode, terminate immediately
	if *e2e {
		fmt.Println("E2E mode: terminating immediately")
		app.Terminate(nil)
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
