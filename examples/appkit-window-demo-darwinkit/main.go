// AppKit Window Demo - Creates a simple macOS window with a button that increments a counter
package main

import (
	"fmt"

	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
)

// Global counter and label
var (
	clickCount   int
	counterLabel appkit.TextField
)

func main() {
	// Initialize the application
	app := appkit.Application_SharedApplication()
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)

	fmt.Println("✓ Starting AppKit application...")
	fmt.Println("  A window will appear - click the button to increment the counter")

	// Create the main window
	window := createMainWindow()
	window.MakeKeyAndOrderFront(nil)

	fmt.Println("✓ Window created and displayed")
	fmt.Println("✓ Click the button to see the counter increment")

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

	window.SetTitle("AppKit Demo - Counter Button")
	window.Center()

	// Get content view
	contentView := window.ContentView()

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

	// Add click counter label (store in global for access in handler)
	counterFrame := foundation.Rect{
		Origin: foundation.Point{X: 50, Y: 80},
		Size:   foundation.Size{Width: 300, Height: 30},
	}

	counterLabel = appkit.NewTextFieldWithFrame(counterFrame)
	counterLabel.SetStringValue("Button clicks: 0")
	counterLabel.SetEditable(false)
	counterLabel.SetBezeled(false)
	counterLabel.SetDrawsBackground(false)
	counterLabel.SetAlignment(appkit.TextAlignmentCenter)

	contentView.AddSubview(counterLabel)

	// Create button
	buttonFrame := foundation.Rect{
		Origin: foundation.Point{X: 150, Y: 120},
		Size:   foundation.Size{Width: 100, Height: 40},
	}

	button := appkit.NewButtonWithFrame(buttonFrame)
	button.SetTitle("Click Me!")
	button.SetBezelStyle(appkit.BezelStyleRounded)
	button.SetButtonType(appkit.ButtonTypeMomentaryPushIn)

	// Add button to window
	contentView.AddSubview(button)

	// Use NSTimer to periodically check if button is being pressed
	// This approach polls the button state - a workaround for target/action complexity
	lastHighlighted := false

	timer := foundation.Timer_TimerWithTimeIntervalRepeatsBlock(
		0.05, // Check every 50ms
		true,  // repeats
		func(timer foundation.Timer) {
			// Check if button is currently highlighted (being clicked)
			isHighlighted := button.IsHighlighted()

			// Detect the transition from highlighted to not highlighted (button release)
			if lastHighlighted && !isHighlighted {
				clickCount++
				counterLabel.SetStringValue(fmt.Sprintf("Button clicks: %d", clickCount))
				fmt.Printf("Button clicked! Count: %d\n", clickCount)
			}

			lastHighlighted = isHighlighted
		},
	)

	// Add timer to run loop so it actually fires
	foundation.RunLoop_CurrentRunLoop().AddTimerForMode(timer, foundation.RunLoopCommonModes)

	return window
}
