// AppKit Window Demo - Creates a simple macOS window with a button
package main

import (
	"fmt"

	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

func main() {
	// Initialize the application
	app := appkit.Application_SharedApplication()
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)

	fmt.Println("✓ Starting AppKit application...")
	fmt.Println("  A window will appear - close it to exit")

	// Create the main window directly (no delegate for simplicity)
	window := createMainWindow()
	window.MakeKeyAndOrderFront(nil)

	fmt.Println("✓ Window created and displayed")

	// Activate the app
	app.ActivateIgnoringOtherApps(true)

	// Run the application
	app.Run()
}

func createMainWindow() appkit.Window {
	// Window frame: x, y, width, height
	frame := foundation.Rect{
		Origin: foundation.Point{X: 100, Y: 100},
		Size:   foundation.Size{Width: 400, Height: 300},
	}

	// Window style mask
	styleMask := appkit.WindowStyleMaskTitled |
		appkit.WindowStyleMaskClosable |
		appkit.WindowStyleMaskMiniaturizable |
		appkit.WindowStyleMaskResizable

	// Create window
	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		frame,
		styleMask,
		appkit.BackingStoreBuffered,
		false,
	)

	window.SetTitle("AppKit Demo - DarwinKit Go Bindings")
	window.Center()

	// Create a button
	buttonFrame := foundation.Rect{
		Origin: foundation.Point{X: 150, Y: 120},
		Size:   foundation.Size{Width: 100, Height: 40},
	}

	button := appkit.NewButtonWithFrame(buttonFrame)
	button.SetTitle("Click Me!")
	button.SetBezelStyle(appkit.BezelStyleRounded)

	// Set button action
	button.SetTarget(button)
	button.SetAction(objc.Sel("performClick:"))

	// Create a custom button that handles clicks
	clickButton := appkit.NewButtonWithFrame(buttonFrame)
	clickButton.SetTitle("Click Me!")
	clickButton.SetBezelStyle(appkit.BezelStyleRounded)

	// Add button to window
	contentView := window.ContentView()
	contentView.AddSubview(clickButton)

	// Create a label
	labelFrame := foundation.Rect{
		Origin: foundation.Point{X: 50, Y: 200},
		Size:   foundation.Size{Width: 300, Height: 30},
	}

	label := appkit.NewTextFieldWithFrame(labelFrame)
	label.SetStringValue("Welcome to AppKit with DarwinKit!")
	label.SetEditable(false)
	label.SetBezeled(false)
	label.SetDrawsBackground(false)
	label.SetAlignment(appkit.TextAlignmentCenter)

	contentView.AddSubview(label)

	// Add click counter label
	counterFrame := foundation.Rect{
		Origin: foundation.Point{X: 50, Y: 80},
		Size:   foundation.Size{Width: 300, Height: 30},
	}

	counterLabel := appkit.NewTextFieldWithFrame(counterFrame)
	counterLabel.SetStringValue("Button clicks: 0")
	counterLabel.SetEditable(false)
	counterLabel.SetBezeled(false)
	counterLabel.SetDrawsBackground(false)
	counterLabel.SetAlignment(appkit.TextAlignmentCenter)

	contentView.AddSubview(counterLabel)

	return window
}
