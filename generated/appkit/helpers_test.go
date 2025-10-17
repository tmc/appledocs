package appkit_test

import (
	"fmt"

	"github.com/tmc/appledocs/generated/appkit"
)

// ExampleSharedApplication demonstrates getting the NSApplication singleton.
func ExampleSharedApplication() {
	app := appkit.SharedApplication()
	fmt.Printf("Got application instance: %T\n", app)

	// Output:
	// Got application instance: appkit.Application
}

// ExampleApplication_SetActivationPolicy demonstrates setting activation policy.
func ExampleApplication_SetActivationPolicy() {
	app := appkit.SharedApplication()

	// Set to regular app (shows in Dock)
	app.SetActivationPolicy(appkit.ActivationPolicyRegular)

	fmt.Println("Set activation policy to Regular")
	fmt.Println("Available policies:")
	fmt.Println("- ActivationPolicyRegular: Normal app with Dock icon")
	fmt.Println("- ActivationPolicyAccessory: No Dock icon")
	fmt.Println("- ActivationPolicyProhibited: Cannot activate")

	// Output:
	// Set activation policy to Regular
	// Available policies:
	// - ActivationPolicyRegular: Normal app with Dock icon
	// - ActivationPolicyAccessory: No Dock icon
	// - ActivationPolicyProhibited: Cannot activate
}

// ExampleWindow_SetTitle demonstrates setting a window's title.
func ExampleWindow_SetTitle() {
	window := appkit.NewWindow()

	window.SetTitle("My Application Window")

	fmt.Println("Window title set successfully")

	// Output:
	// Window title set successfully
}

// ExampleWindow_ContentView demonstrates getting a window's content view.
func ExampleWindow_ContentView() {
	window := appkit.NewWindow()

	contentView := window.ContentView()

	fmt.Printf("Got content view: %T\n", contentView)

	// Output:
	// Got content view: appkit.View
}

// ExampleView_AddSubviewTyped demonstrates adding a subview type-safely.
func ExampleView_AddSubviewTyped() {
	contentView := appkit.NewWindow().ContentView()
	subview := appkit.NewView()

	contentView.AddSubviewTyped(subview)

	fmt.Println("Added subview successfully")

	// Output:
	// Added subview successfully
}

// ExampleView_SetFrameRect demonstrates setting a view's frame.
func ExampleView_SetFrameRect() {
	view := appkit.NewView()

	// Set frame to (x:10, y:20, width:100, height:50)
	view.SetFrameRect(10, 20, 100, 50)

	fmt.Println("Set view frame successfully")

	// Output:
	// Set view frame successfully
}

// ExampleButton_SetTitleString demonstrates setting a button's title.
func ExampleButton_SetTitleString() {
	button := appkit.NewButton()

	button.SetTitleString("Click Me!")

	fmt.Println("Button title set")

	// Output:
	// Button title set
}

// ExampleButton_SetButtonType demonstrates setting button type with constants.
func ExampleButton_SetButtonType() {
	button := appkit.NewButton()

	// Set to momentary light button (standard push button)
	button.SetButtonType(appkit.ButtonTypeMomentaryLight)
	button.SetBezelStyle(appkit.BezelStyleRounded)

	fmt.Println("Button type and style configured")
	fmt.Println("Available types: MomentaryLight, PushOnPushOff, Toggle, Switch, Radio")
	fmt.Println("Available styles: Rounded, RegularSquare, Circular, HelpButton, etc.")

	// Output:
	// Button type and style configured
	// Available types: MomentaryLight, PushOnPushOff, Toggle, Switch, Radio
	// Available styles: Rounded, RegularSquare, Circular, HelpButton, etc.
}

// ExampleControl_SetStringValue demonstrates setting a control's string value.
func ExampleControl_SetStringValue() {
	control := appkit.NewControl()

	// Set the control's string value
	control.SetStringValue("Hello, Control!")

	fmt.Println("Control string value set")

	// Output:
	// Control string value set
}

// ExampleControl_SetEditable demonstrates making a control editable or read-only.
func ExampleControl_SetEditable() {
	control := appkit.NewControl()

	// Make the control read-only
	control.SetEditable(false)

	fmt.Println("Control set to read-only")

	// Output:
	// Control set to read-only
}

// ExampleControl_SetBordered demonstrates setting whether a control has a border.
func ExampleControl_SetBordered() {
	control := appkit.NewControl()

	// Remove the border
	control.SetBordered(false)

	fmt.Println("Control border removed")

	// Output:
	// Control border removed
}

// ExampleControl_SetAlignment demonstrates setting text alignment in a control.
func ExampleControl_SetAlignment() {
	control := appkit.NewControl()

	// Center-align the text
	control.SetAlignment(appkit.TextAlignmentCenter)

	fmt.Println("Text alignment set to center")
	fmt.Println("Available alignments:")
	fmt.Println("- TextAlignmentLeft: Left-aligned")
	fmt.Println("- TextAlignmentCenter: Center-aligned")
	fmt.Println("- TextAlignmentRight: Right-aligned")
	fmt.Println("- TextAlignmentJustified: Justified")
	fmt.Println("- TextAlignmentNatural: Natural (system default)")

	// Output:
	// Text alignment set to center
	// Available alignments:
	// - TextAlignmentLeft: Left-aligned
	// - TextAlignmentCenter: Center-aligned
	// - TextAlignmentRight: Right-aligned
	// - TextAlignmentJustified: Justified
	// - TextAlignmentNatural: Natural (system default)
}

// ExampleTextField_SetDrawsBackground demonstrates controlling background drawing.
func ExampleTextField_SetDrawsBackground() {
	// Note: This demonstrates the API but we can't actually create a TextField
	// in this example because TextField constructors aren't generated yet.
	// In a real application, you would create a TextField and use it like:
	//
	// textField := appkit.NewTextField()
	// textField.SetDrawsBackground(false)

	fmt.Println("TextField background control is available")
	fmt.Println("Use SetDrawsBackground(false) to make background transparent")

	// Output:
	// TextField background control is available
	// Use SetDrawsBackground(false) to make background transparent
}

// ExampleRunApp demonstrates how to use RunApp to create a simple macOS application.
// This example shows the minimal code needed to create and display a window.
//
// Note: This example will block when run interactively. In automated tests, it's
// typically skipped or run with a timeout.
func ExampleRunApp() {
	// This example demonstrates the RunApp API but doesn't actually run
	// to avoid blocking tests. In a real application, this would be in main():
	//
	// func main() {
	//     appkit.RunApp(func(app appkit.Application) {
	//         window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(...)
	//         window.MakeKeyAndOrderFront(0)
	//     })
	// }

	fmt.Println("RunApp provides a simple way to create macOS applications")
	fmt.Println("It handles all the boilerplate:")
	fmt.Println("- Thread locking")
	fmt.Println("- Getting shared application")
	fmt.Println("- Setting activation policy")
	fmt.Println("- Finishing launch")
	fmt.Println("- Running the event loop")

	// Output:
	// RunApp provides a simple way to create macOS applications
	// It handles all the boilerplate:
	// - Thread locking
	// - Getting shared application
	// - Setting activation policy
	// - Finishing launch
	// - Running the event loop
}
