// Hello World using darwinkit
//
// This example demonstrates using darwinkit for AppKit bindings,
// creating a simple window with a button and click counter.
package main

import (
	"fmt"

	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
)

var (
	clickCount   int
	counterLabel appkit.TextField
)

func main() {
	fmt.Println("=== Hello World (Darwinkit) ===\n")

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
	window.SetTitle("Hello from Darwinkit!")

	// Create label
	label := appkit.NewTextField()
	label.SetStringValue("This uses darwinkit!")
	label.SetFrameOrigin(foundation.Point{X: 50, Y: 200})
	label.SetFrameSize(foundation.Size{Width: 300, Height: 50})
	label.SetEditable(false)
	label.SetBordered(false)
	label.SetBackgroundColor(nil)

	window.ContentView().AddSubview(label)

	// Create counter label
	counterLabel = appkit.NewTextField()
	counterLabel.SetStringValue("Button clicks: 0")
	counterLabel.SetFrameOrigin(foundation.Point{X: 50, Y: 80})
	counterLabel.SetFrameSize(foundation.Size{Width: 300, Height: 30})
	counterLabel.SetEditable(false)
	counterLabel.SetBordered(false)
	counterLabel.SetBackgroundColor(nil)
	counterLabel.SetAlignment(appkit.TextAlignmentCenter)

	window.ContentView().AddSubview(counterLabel)

	// Create button
	button := appkit.NewButtonWithTitle("Click Me!")
	button.SetFrameOrigin(foundation.Point{X: 150, Y: 130})
	button.SetFrameSize(foundation.Size{Width: 100, Height: 40})
	button.SetButtonType(appkit.ButtonTypeMomentaryLight)
	button.SetBezelStyle(appkit.BezelStyleRounded)

	// Set button action
	button.SetTarget(button)
	button.SetAction(appkit.Sel("buttonClicked:"))

	// Register button click handler
	appkit.AddMethod(button, appkit.Sel("buttonClicked:"), func() {
		clickCount++
		fmt.Printf("Button clicked! Count: %d\n", clickCount)
		counterLabel.SetStringValue(fmt.Sprintf("Button clicks: %d", clickCount))
	})

	window.ContentView().AddSubview(button)

	// Show window
	window.MakeKeyAndOrderFront(nil)
	app.ActivateIgnoringOtherApps(true)

	fmt.Println("✅ Window created using darwinkit")
	fmt.Println("   Click the button to see the counter increment!")
	fmt.Println("   Close window or press Cmd+Q to quit\n")

	// Run application
	app.Run()
}
